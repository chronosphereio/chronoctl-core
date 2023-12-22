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

// Package types contains all of the type metadata for entities output by chronoctl
package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	inValidChars = regexp.MustCompile("[^a-zA-Z0-9]+")
	separator    = "_"
)

// TypeMeta describes an object in the Chronosphere API. Note that the json tags for the
// object use snake case instead of camel case for field names to match the convention
// used by the API.
type TypeMeta struct {
	// APIVersion is the versioned schema of this representation of an object.
	APIVersion string `json:"api_version,omitempty"`

	// Kind is the name of the resource this object represents in the Chronosphere API.
	Kind string `json:"kind,omitempty"`
}

// Type returns itself
func (t TypeMeta) Type() TypeMeta { return t }

// String returns a string representation of the type metadata.
func (t TypeMeta) String() string {
	return fmt.Sprintf("%s/%s", t.APIVersion, t.Kind)
}

// Object is the interface that all types in the Chronosphere API must fulfill.
type Object interface {
	Type() TypeMeta
	Description() string
	Identifier() string
}

// TypeDescription creates a friendly identifying description of an object based on its
// king and a list of key-value pairs.
// See Description for examples.
func TypeDescription(obj Object, idKeyPairs ...string) string {
	return Description(obj.Type().Kind, idKeyPairs...)
}

// Description creates a friendly identifying description of an object based on its
// king and a list of key-value pairs.
// The idKeyPairs must have an even number of elements, although values can be empty.
// Example: TypeDescription("Foo") will result in "Foo{}"
// Example: TypeDescription("Foo", "id", "xyz", "name", "thing") will result in "Foo{id="xyz", name="thing"}"
// Example: TypeDescription("Foo", "id", "", "name", "thing") will result in "Foo{name="thing"}"
func Description(typeKind string, idKeyPairs ...string) string {
	if len(idKeyPairs) == 0 {
		return typeKind + "{}"
	}

	pairs := make([]string, 0, len(idKeyPairs)/2)
	for i := 0; i < len(idKeyPairs); i += 2 {
		if idKeyPairs[i+1] == "" {
			continue
		}

		pairs = append(pairs, fmt.Sprintf("%s=%s", idKeyPairs[i], strconv.Quote(idKeyPairs[i+1])))
	}

	if len(pairs) == 0 {
		return typeKind + "{}"
	}

	return fmt.Sprintf("%s{%s}", typeKind, strings.Join(pairs, ", "))
}

// DisplayName creates a name for a type based on its metadata.
func DisplayName(typeKind string, identifier string) string {
	formattedTypeKind := strings.ToLower(strings.ReplaceAll(typeKind, "-", "_"))
	if identifier == "" {
		return formattedTypeKind
	}
	formattedIdentifier := inValidChars.ReplaceAllString(strings.ToLower(identifier), separator)

	return formattedTypeKind + separator + formattedIdentifier
}
