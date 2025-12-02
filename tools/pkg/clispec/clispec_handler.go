// Copyright 2023 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clispec

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/chronosphereio/chronoctl-core/tools/pkg/specscanner"
	"github.com/go-openapi/inflect"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/go-swagger/go-swagger/generator"
	"github.com/pkg/errors"
)

type cliSpecGen struct {
	spec              *Spec
	currentPath       string
	currentEntity     *Entity
	currentCommand    *Command
	currentAPIVersion string
	inflectionRuleset *inflect.Ruleset
}

var (
	swaggerTypeToGoType = map[string]string{
		"integer": "int",
		"boolean": "bool",
		"number":  "float64",
	}
)

// Parse will parse the swagger spec at the given path and return a *cli.Spec
func Parse(path string) (*Spec, error) {
	path, err := withAutoXOrder(path)
	if err != nil {
		return nil, err
	}

	swaggerSpec, err := loads.Spec(path)
	if err != nil {
		return nil, err
	}

	gen := newCliSpecGen()
	err = specscanner.Scan(swaggerSpec.Raw(), gen, nil)
	if err != nil {
		return nil, err
	}

	return gen.spec, nil
}

func newCliSpecGen() *cliSpecGen {
	ruleset := inflect.NewDefaultRuleset()
	// special case so the 's' isn't taken off when singularizing
	ruleset.AddUncountable("sync-prometheus")
	return &cliSpecGen{inflectionRuleset: ruleset}
}

func (c *cliSpecGen) StartScan(apiSpec *spec.Swagger) error {
	c.spec = &Spec{
		BasePath: "api/",
		Entities: map[string]*Entity{},
	}

	if apiSpec.BasePath != "" {
		c.spec.BasePath = apiSpec.BasePath
	}

	return nil
}

func (c *cliSpecGen) EndScan() error {
	c.spec.postProcessEntities()
	return nil
}

func (c *cliSpecGen) getEntity(path string) (*Entity, error) {
	pathParts := splitPath(path)
	// we expect the path to be in the format /api/<version>/<entity> or /api/<version>/<entity>/<id>
	if len(pathParts) < 4 {
		return nil, errors.Errorf("unexpected path format: %s", path)
	}

	var entityName string

	// Handle service attributes paths - both nested and flat
	if (len(pathParts) >= 6 && pathParts[3] == "services" && pathParts[5] == "attributes") ||
		pathParts[3] == "service-attributes" {
		entityName = "service-attribute"
	} else {
		entityName = pathParts[3]
	}

	if _, ok := c.spec.Entities[entityName]; !ok {
		entity := &Entity{
			Name: entityName,
			Type: &Type{
				Kind: c.inflectionRuleset.Typeify(entityName),
			},
		}
		
		// Set the appropriate slug field based on entity type
		if entityName == "service-attribute" {
			entity.EntityLinkedSingletonSlug = "ServiceSlug"
		}
		
		c.spec.Entities[entityName] = entity
	}

	return c.spec.Entities[entityName], nil
}

func (c *cliSpecGen) StartPath(path string, pathSpec spec.PathItem) error {
	c.currentPath = path

	entity, err := c.getEntity(c.currentPath)
	if err != nil {
		return err
	}

	pathParts := splitPath(path)
	// we expect the path to be in the format /api/<version>/<entity> or /api/<version>/<entity>/<id>
	if len(pathParts) < 4 {
		return errors.Errorf("unexpected path format: %s", path)
	}
	entity.Type.APIVersion = strings.Join(pathParts[1:3], "/")
	return nil
}

func (c *cliSpecGen) EndPath() error {
	c.currentPath = ""
	return nil
}

func (c *cliSpecGen) StartOp(op string, opSpec *spec.Operation) error {
	entity, err := c.getEntity(c.currentPath)
	if err != nil {
		return err
	}

	c.currentCommand = &Command{
		APISummary: cleanSummary(opSpec.Summary),
		Verb:       op,
		Path:       c.currentPath,
		Parameters: []*Parameter{},
	}

	switch op {
	case http.MethodGet:
		if isParameterizedPath(c.currentPath) {
			c.currentCommand.Name = commandNameGet
			c.currentCommand.Description = fmt.Sprintf("Get %s resource", entity.Name)
			entity.Get = c.currentCommand
		} else {
			c.currentCommand.Name = commandNameList
			c.currentCommand.Description = fmt.Sprintf("List %s resources", entity.Name)
			entity.List = c.currentCommand
		}
	case http.MethodPost:
		c.currentCommand.Name = commandNameCreate
		c.currentCommand.Description = fmt.Sprintf("Create %s resource", entity.Name)
		entity.Create = c.currentCommand
	case http.MethodPut:
		c.currentCommand.Name = commandNameUpdate
		c.currentCommand.Description = fmt.Sprintf("Apply %s resource", entity.Name)
		entity.Update = c.currentCommand
	case http.MethodDelete:
		c.currentCommand.Name = commandNameDelete
		c.currentCommand.Description = fmt.Sprintf("Delete %s resource", entity.Name)
		entity.Delete = c.currentCommand
	default:
		return errors.Errorf("unknown operation: %s", op)
	}

	return nil
}

func (c *cliSpecGen) EndOp() error {
	// sort parameters so we have a stable order
	sort.Slice(c.currentCommand.Parameters, func(i, j int) bool {
		return c.currentCommand.Parameters[i].Name < c.currentCommand.Parameters[j].Name
	})
	return nil
}

func (c *cliSpecGen) StartResponse(statusCode int, response *spec.Response) error {
	if statusCode != http.StatusOK {
		return nil
	}

	entity, err := c.getEntity(c.currentPath)
	if err != nil {
		return err
	}

	if response.Schema.Properties != nil {
		c.currentCommand.Response = &Response{
			Schema: &Type{
				APIVersion: entity.Type.APIVersion,
				Kind:       entity.Type.Kind,
			},
		}
		if _, hasPages := response.Schema.Properties["page"]; hasPages {
			c.currentCommand.Response.IsArray = true
		}
	}
	return nil
}

func (c *cliSpecGen) EndResponse() error {
	return nil
}

func (c *cliSpecGen) StartParam(param *spec.Parameter) error {
	entity, err := c.getEntity(c.currentPath)
	if err != nil {
		return err
	}

	// NB: Swagger (OpenAPI v2.x) sends the body as a parameter. OpenAPI 3
	// stores the body as its own property called `requestBody`. If we ever
	// upgrade to OpenAPI v3, we'll need to handle this differently.
	if param.In == string(PlaceBody) {
		c.currentCommand.RequestBody = &RequestBody{
			Required: param.Required,
			Schema: &Type{
				APIVersion: c.currentAPIVersion,
				Kind:       c.inflectionRuleset.Typeify(entity.Name),
				Properties: param.Schema.Properties,
			},
		}
	} else {
		parameterPlace, err := parameterPlace(param.In)
		if err != nil {
			return err
		}

		cliParam := &Parameter{
			Name:        param.Name,
			Description: strings.ReplaceAll(param.Description, "\n", " "),
			GoType:      param.Type,
			Place:       parameterPlace,
			Required:    param.Required,
		}

		if param.Type == "array" {
			arrayType := param.Items.Type
			if goType, ok := swaggerTypeToGoType[arrayType]; ok {
				arrayType = goType
			}
			cliParam.GoType = fmt.Sprintf("[]%s", arrayType)
		} else {
			if goType, ok := swaggerTypeToGoType[param.Type]; ok {
				cliParam.GoType = goType
			}
		}

		if param.Type == "" {
			cliParam.GoType = param.Schema.Type[0]
			if cliParam.GoType == "array" {
				cliParam.GoType = fmt.Sprintf("[]%s", param.Schema.Items.Schema.Type[0])
			}
		}

		c.currentCommand.Parameters = append(c.currentCommand.Parameters, cliParam)
	}

	return nil
}

func (c *cliSpecGen) EndParam() error {
	return nil
}

// A parameterized path is one that contains a path parameter, e.g. /foo/{bar}.
func isParameterizedPath(path string) bool {
	parts := strings.Split(path, "/")
	last := parts[len(parts)-1]
	return strings.HasPrefix(last, "{")
}

func parameterPlace(in string) (ParameterPlace, error) {
	switch in {
	case string(PlacePath):
		return PlacePath, nil
	case string(PlaceQuery):
		return PlaceQuery, nil
	case string(PlaceBody):
		return PlaceBody, nil
	}
	return "", errors.Errorf("unknown parameter location: %s", in)
}

// some of the swagger spec summaries have extra \n or asterisks as prefixes/suffixes.
// We attempted to clean that up here.
func cleanSummary(summary string) string {
	summary = strings.Replace(summary, "***\n", "", -1)
	summary = strings.Replace(summary, "\n***", "", -1)
	return summary
}

func splitPath(path string) []string {
	return strings.Split(strings.Trim(path, "/"), "/")
}

// withAutoXOrder wraps the swagger generator's WithAutoXOrder which
// panics instead of returning errors, so we convert those to errors.
func withAutoXOrder(path string) (newPath string, retErr error) {
	defer func() {
		if err := recover(); err != nil {
			newPath = ""
			retErr = errors.New(fmt.Sprint(err))
		}
	}()

	return generator.WithAutoXOrder(path), nil
}
