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

// package main is the entrypoint for the chronocgen utility. This utility
// takes a CLI spec based on a swagger spec and generates CLI commands
// for interacting with the Chronosphere API.
package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/chronosphereio/chronoctl-core/tools/cmd/chronogen/clitemplates"
	"github.com/chronosphereio/chronoctl-core/tools/pkg/clispec"
	"github.com/go-openapi/inflect"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

var (
	paramDenyList = map[string]any{
		"page.max_size": true,
		"page.token":    true,
	}
	templateFuncs = map[string]any{
		"singularize":     inflect.Singularize,
		"lowerCamel":      lowerCamel,
		"stripNewlines":   stripNewlines,
		"shouldReference": shouldReference,
		"dasherize":       dasherize,
	}
	cliTmpl    = template.Must(template.New("cli").Funcs(templateFuncs).Parse(clitemplates.CLI))
	entityTmpl = newEntityTmpl()
)

func newEntityTmpl() *template.Template {
	tmpl := template.Must(template.New("entity").Funcs(templateFuncs).Parse(clitemplates.Entity))
	template.Must(tmpl.Parse(clitemplates.Create))
	template.Must(tmpl.Parse(clitemplates.Read))
	template.Must(tmpl.Parse(clitemplates.Update))
	template.Must(tmpl.Parse(clitemplates.Delete))
	template.Must(tmpl.Parse(clitemplates.List))
	template.Must(tmpl.Parse(clitemplates.Scaffold))
	return tmpl
}

type generatedFile struct {
	name    string
	content string
}

type mainSpec struct {
	Entities []entitySpec
	Pkg      pkgMetadata
	PkgName  string
}

func newMainSpec(entities map[string]entitySpec, pkgName string) mainSpec {
	entSlice := maps.Values(entities)
	sort.Slice(entSlice, func(i, j int) bool {
		return entSlice[i].Name() < entSlice[j].Name()
	})
	return mainSpec{
		Entities: entSlice,
		PkgName:  pkgName,
	}
}

type entitySpec struct {
	Entity                  *clispec.Entity
	Pkg                     pkgMetadata
	PkgName                 string
	ScaffoldYAML            string
	SupportsShorthandCreate bool
}

// Name returns the singularized name.
func (e entitySpec) Name() string {
	if e.Entity.IsSingleton {
		// Singletons shouldn't have singular/plural applied.
		return upperCamel(e.Entity.Name)
	}
	if e.Entity.Name == "slos" {
		return "SLO"
	}
	return inflect.Singularize(upperCamel(e.Entity.Name))
}

func (e entitySpec) ModelName() string {
	name := e.Name()
	// ClassicDashboards use the GrafanaDashboard type, so we need to customize
	// the model name.
	if name == "ClassicDashboard" {
		name = "GrafanaDashboard"
	}
	return fmt.Sprintf("%s%s", e.Pkg.Model(), name)
}

// NameP returns the pluralized name.
func (e entitySpec) NameP() string {
	if e.Entity.IsSingleton {
		// Singletons shouldn't have singular/plural applied.
		return upperCamel(e.Entity.Name)
	}
	if e.Entity.Name == "slos" {
		return "SLOs"
	}
	return inflect.Pluralize(upperCamel(e.Entity.Name))
}

// PayloadNameP returns the pluralized name for the entity when generating code
// for accessing API response Payloads. In some rare cases (SLOs) the pluralized
// name is in agreement to our typical pluralization rules.
func (e entitySpec) PayloadNameP() string {
	if e.Entity.Name == "slos" {
		return "Slos"
	}
	return e.NameP()
}

// FileFlagRequired marks the file flags as required if shorthand is not supported.
func (e entitySpec) FileFlagRequired() bool {
	return !e.SupportsShorthandCreate
}

// IncludeListCommand determines whether a list command should be included for
// the entity
func (e entitySpec) IncludeListCommand() bool {
	return e.Entity.IsNotSingleton() && e.Entity.List != nil
}

// ListParameters strips out any of the page params.
func (e entitySpec) ListParameters() []*clispec.Parameter {
	originalParams := e.Entity.List.Parameters
	var filteredParams []*clispec.Parameter
	for _, p := range originalParams {
		if _, ok := paramDenyList[p.Name]; ok {
			continue
		}
		filteredParams = append(filteredParams, p)
	}
	return filteredParams
}

func generate(pkgName string, spec *clispec.Spec) ([]generatedFile, error) {
	var files []generatedFile
	specs := make(map[string]entitySpec, len(spec.Entities))
	allowList := entityAllowList()

	for k, e := range spec.Entities {
		if _, ok := allowList[e.Name]; len(allowList) > 0 && !ok {
			continue
		}

		scaffoldYAML, err := e.ScaffoldYAML()
		if err != nil {
			return nil, errors.Wrap(err, e.Name)
		}

		// Basically what's going on is both classic dashboards and grafana dashboards use
		// the same underlying ref: configv1GrafanaDashboard. The scaffold comments come
		// from the spec comments, aka what's on configv1GrafanaDashboard so as a result
		// its references GrafanaDashboard.
		if e.Name == "classic-dashboards" {
			scaffoldYAML = strings.ReplaceAll(scaffoldYAML, "GrafanaDashboard", "ClassicDashboard")
		}

		var b bytes.Buffer
		entSpec := entitySpec{
			Entity:                  e,
			Pkg:                     newPkgMetadata(e),
			PkgName:                 pkgName,
			ScaffoldYAML:            scaffoldYAML,
			SupportsShorthandCreate: shorthandCreateSupport[e.Name],
		}
		specs[k] = entSpec
		if err := entityTmpl.Execute(&b, entSpec); err != nil {
			return nil, fmt.Errorf("execute template for %v: %v", e.Name, err)
		}
		files = append(files, generatedFile{
			name:    strings.ReplaceAll(entSpec.Entity.Name, "-", "_"),
			content: b.String(),
		})
	}

	var b bytes.Buffer
	if err := cliTmpl.Execute(&b, newMainSpec(specs, pkgName)); err != nil {
		return nil, fmt.Errorf("execute main template: %v", err)
	}
	files = append(files, generatedFile{
		name:    "cli",
		content: b.String(),
	})
	return files, nil
}

func stripNewlines(str string) string {
	return strings.ReplaceAll(str, "\n", " ")
}

// shouldReference determines whether the param's gotype requires an `&` reference
// in the ListParameters.
func shouldReference(param *clispec.Parameter) bool {
	return !strings.HasPrefix(param.GoType, "[]")
}

// upperCamel turns a string separated by - or _ into CamelCase (first letter is uppercase).
func upperCamel(str string) string {
	return inflect.Camelize(str)
}

func dasherize(str string) string {
	return strings.ReplaceAll(str, "_", "-")
}

// lowerCamel turns a string separated by - or _ into camelCase (first letter is lowercase).
func lowerCamel(str string) string {
	return inflect.CamelizeDownFirst(str)
}

func entityAllowList() map[string]any {
	m := map[string]any{}
	if *allowedEntities == "" {
		return m
	}
	for _, e := range strings.Split(*allowedEntities, ",") {
		m[e] = true
	}
	return m
}
