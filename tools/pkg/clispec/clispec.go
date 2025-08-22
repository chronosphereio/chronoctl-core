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
	"strings"

	"github.com/go-openapi/inflect"
	"github.com/go-openapi/spec"

	"github.com/chronosphereio/chronoctl-core/src/thirdparty/yaml"
)

const (
	commandNameCreate = "create"
	commandNameGet    = "get"
	commandNameUpdate = "update"
	commandNameDelete = "delete"
	commandNameList   = "list"
)

// ParameterPlace is a string that represents the location of a parameter in a request
type ParameterPlace string

const (
	// PlacePath represents a parameter that is in the path of a request (e.g. /foo/{bar})
	PlacePath ParameterPlace = "path"
	// PlaceQuery represents a parameter that is in the query string of a request (e.g. ?foo=bar)
	PlaceQuery ParameterPlace = "query"
	// PlaceBody represents a parameter that is in the body of a request
	PlaceBody ParameterPlace = "body"
)

// Spec is the top-level struct that represents the Chronosphere CLI spec
type Spec struct {
	Name     string             `json:"name,omitempty"`
	BasePath string             `json:"basePath,omitempty"`
	Entities map[string]*Entity `json:"entities,omitempty"`
}

func (s *Spec) postProcessEntities() {
	for _, entity := range s.Entities {
		entity.processSingleton()
	}
}

// Entity represents a Chronosphere entity exposed by the public API (e.g. a "monitor")
type Entity struct {
	Name   string   `json:"name,omitempty"`
	Create *Command `json:"create,omitempty"`
	Get    *Command `json:"get,omitempty"`
	Update *Command `json:"update,omitempty"`
	Delete *Command `json:"delete,omitempty"`
	List   *Command `json:"list,omitempty"`
	Type   *Type    `json:"type,omitempty"`

	IsSingleton bool
}

// Command represents a Chronosphere CLI command
type Command struct {
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	APISummary  string       `json:"apiSummary,omitempty"`
	Path        string       `json:"path,omitempty"`
	Verb        string       `json:"verb,omitempty"`
	Parameters  []*Parameter `json:"parameters,omitempty"`
	RequestBody *RequestBody `json:"requestBody,omitempty"`
	Response    *Response    `json:"response,omitempty"`
}

// SupportsDryRun determines in the endpoint has support for dry-run.
func (c *Command) SupportsDryRun() bool {
	// n.b. we check the parameters for Delete endpoints.
	for _, parameter := range c.Parameters {
		if parameter.Name == "dry_run" {
			return true
		}
	}

	if c.RequestBody == nil || c.RequestBody.Schema == nil {
		return false
	}

	for propertyName := range c.RequestBody.Schema.Properties {
		if propertyName == "dry_run" {
			return true
		}
	}

	return false
}

// processSingleton shuffles around commands.
// The heuristic we're using is
// 1. There is no parameterized path - no read with slug
// 2. There is a  list path.
// 3. Update is parameterized.
func (e *Entity) processSingleton() {
	updateNotParamerterized := e.Update != nil && !isParameterizedPath(e.Update.Path)
	if e.Get == nil && e.List != nil && updateNotParamerterized {
		e.IsSingleton = true
		e.Get, e.List = e.List, e.Get
		fmt.Println("detected singleton for", e.Name)
	}
}

// IsNotSingleton determines if the entity is not a singleton. Helper for templates.
func (e *Entity) IsNotSingleton() bool {
	return !e.IsSingleton
}

func (e *Entity) getEntitySchema() spec.Schema {
	if properties, ok := e.Create.RequestBody.Schema.Properties[e.singletonName()]; ok {
		return properties
	}
	name := strings.ToLower(inflect.Underscore(inflect.Singularize(e.Type.Kind)))
	return e.Create.RequestBody.Schema.Properties[name]
}

func (e *Entity) SupportsForceDelete() bool {
	if e.Delete != nil {
		for _, parameter := range e.Delete.Parameters {
			if parameter.Name == "force_delete" {
				return true
			}
		}
	}
	return false
}

func (e *Entity) singletonName() string {
	return strings.ToLower(inflect.Underscore(inflect.Pluralize(e.Type.Kind)))
}

// ScaffoldYAML returns a valid config/v1 object of the entity.
func (e *Entity) ScaffoldYAML() (string, error) {
	root := yaml.Node{
		Kind: yaml.MappingNode,
	}

	root.Content = append(root.Content,
		newKeyNode("api_version", "" /* description */),
		&yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: e.Type.APIVersion,
		},

		newKeyNode("kind", "" /* description */),
		&yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: e.scaffoldingKindName(),
		},

		newKeyNode("spec", "" /* description */),
		newNode(e.getEntitySchema()),
	)

	b, err := yaml.Marshal(root)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (e *Entity) scaffoldingKindName() string {
	entityKind := e.Type.Kind
	if entityKind == "Slo" {
		return "SLO"
	}
	// Set of singleton entity kinds that are pluralized.
	// By default, singletons are assumed to be singular, but a few
	// singletons are naturally plural.
	pluralSingletons := map[string]bool{
		"ResourcePool":          true,
		"TraceTailSamplingRule": true,
	}
	if pluralSingletons[entityKind] {
		return inflect.Pluralize(entityKind)
	}
	return entityKind
}

// Parameter represents a parameter to a Chronosphere API that must be provided by the user
type Parameter struct {
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Place       ParameterPlace `json:"place,omitempty"`
	GoType      string         `json:"goType,omitempty"`
	Required    bool           `json:"required,omitempty"`
}

// FlagType returns the function name we expect to use for the given parameter type.
func (p *Parameter) FlagType() string {
	var builder strings.Builder
	isSlice := strings.HasPrefix(p.GoType, "[]")
	t := strings.TrimPrefix(p.GoType, "[]")
	t = inflect.Capitalize(t)

	builder.WriteString(t)
	if isSlice {
		builder.WriteString("Slice")
	}
	builder.WriteString("Var")
	return builder.String()
}

// SwaggerName returns the name of the field in the Swagger generated List parameters.
func (p *Parameter) SwaggerName() string {
	name := inflect.Camelize(p.Name)

	// Camelize doesn't handle acronyms, and uses Json instead of JSON.
	replaces := [][2]string{
		{"Json", "JSON"},
		{"Id", "ID"},
		{"Ui", "UI"},
		{"Slo", "SLO"},
	}
	for _, r := range replaces {
		if strings.HasSuffix(name, r[0]) {
			name = name[:len(name)-len(r[0])] + r[1]
		}
		if strings.HasPrefix(name, r[0]) {
			name = r[1] + name[len(r[0]):]
		}
	}
	return name
}

// Response represents the response from a Chronosphere API request
type Response struct {
	IsArray bool  `json:"isArray,omitempty"`
	Schema  *Type `json:"schema,omitempty"`
}

// RequestBody represents the request body of a Chronosphere API request
type RequestBody struct {
	Required bool  `json:"required,omitempty"`
	Schema   *Type `json:"schema,omitempty"`
}

// Type represents a type for the input/output files that chronoctl writes
type Type struct {
	APIVersion string                `json:"apiVersion,omitempty"`
	Kind       string                `json:"kind,omitempty"`
	Properties spec.SchemaProperties `json:"properties,omitempty"`
}
