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

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/chronosphereio/chronoctl-core/tools/pkg/clispec"
	"github.com/go-openapi/inflect"
)

// pkgMetadata takes an entity's API type and version (e.g., "v1/config") and
// generates package/type names (e.g.., "ConfigV1, Configv1, configv1, config_v1).
type pkgMetadata struct {
	Type       string
	APIVersion string
}

func newPkgMetadata(entity *clispec.Entity) pkgMetadata {
	parts := strings.Split(entity.Type.APIVersion, "/")
	if len(parts) != 2 {
		log.Fatalf("invalid field entity.Type.APIVersion, expected only two fields: %+v", entity)
	}
	return pkgMetadata{
		APIVersion: parts[0],
		Type:       parts[1],
	}
}

func (p pkgMetadata) Package() string {
	return fmt.Sprintf("%s%s", p.Type, p.APIVersion)
}

func (p pkgMetadata) Client() string {
	return fmt.Sprintf("%s_%s", p.Type, p.APIVersion)
}

func (p pkgMetadata) Swagger() string {
	return inflect.Camelize(fmt.Sprintf("%s_%s", p.Type, p.APIVersion))
}

func (p pkgMetadata) Model() string {
	return (fmt.Sprintf("%s%s", inflect.Camelize(p.Type), p.APIVersion))
}
