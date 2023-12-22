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

	"github.com/chronosphereio/chronoctl-core/src/thirdparty/yaml"
	"github.com/go-openapi/spec"
)

// newKeyNode returns a new yaml node for the object key.
//
// We must attach the documentation string as a head comment on the key node,
// otherwise the comment will render incorrectly if attached to the value node.
func newKeyNode(key string, description string) *yaml.Node {
	return &yaml.Node{
		HeadComment: strings.ReplaceAll(description, "`", "'"), // we replace backticks with single quote to stop escaping.
		Kind:        yaml.ScalarNode,
		Value:       key,
	}
}

// newNode returns a node for a scalar value or builds an object node
// recursively.
func newNode(schema spec.Schema) *yaml.Node {
	// if a node has allOf set, it indicates that there is either a nested object
	// or that the value type of the key is an enum.
	// Note that this is the only time a type won't be set.
	if len(schema.AllOf) > 0 {
		return newNode(schema.AllOf[0])
	} else if len(schema.AllOf) > 1 {
		panic(fmt.Sprintf("invariant violated: more than one all_of found in schema: %+v", schema))
	}

	if len(schema.Type) == 0 {
		panic(fmt.Sprintf("invariant violated: schema type is empty: %+v", schema))
	}

	valueType := schema.Type[0]
	switch valueType {
	case "boolean", "string", "number", "integer":
		return newScalarNode(schema)
	case "array":
		return newArrayNode(schema)
	case "object":
		return newObjectNode(schema)
	default:
		// We capture the only defined spec types above:
		// https://swagger.io/docs/specification/data-models/data-types/
		// anything not that values does not conform to the spec.
		panic(fmt.Sprintf("unknown schema type value %q", valueType))
	}
}

func newArrayNode(schema spec.Schema) *yaml.Node {
	node := &yaml.Node{
		Kind: yaml.SequenceNode,
	}
	if schema.Items == nil {
		return node
	}
	if schema.Items.Schema != nil {
		node.Content = append(node.Content, newNode(*schema.Items.Schema))
	}
	for _, schema := range schema.Items.Schemas {
		node.Content = append(node.Content, newNode(schema))
	}
	return node
}

func newObjectNode(schema spec.Schema) *yaml.Node {
	// There are two kinds of objects: additionalProps which defines a map[string]T.
	// Then theres those that implement Properties, which is a nested object.
	// See https://goswagger.io/use/models/schemas.html#mapping-patterns for how this maps.

	// When the schema properties are set, it means that this is an object node that needs
	// to be recursed through.
	if len(schema.Properties) > 0 {
		return newSchemaPropertiesNode(schema.Properties)
	}

	if props := schema.AdditionalProperties; props != nil {
		if propSchema := props.Schema; propSchema != nil {
			value := getScalarValue(propSchema)
			return &yaml.Node{
				Kind: yaml.MappingNode,
				Content: []*yaml.Node{
					newKeyNode("key_1", ""),
					&yaml.Node{
						Kind:  yaml.ScalarNode,
						Value: value,
					},
				},
			}
		}
	}

	return &yaml.Node{}
}

// newSchemaPropertiesNode iterates through a set of <key, schema>,
// building the underlying object nodes.
func newSchemaPropertiesNode(properties spec.SchemaProperties) *yaml.Node {
	root := yaml.Node{
		Kind: yaml.MappingNode,
	}

	for _, item := range properties.ToOrderedSchemaItems() {
		if item.Schema.ReadOnly {
			continue
		}
		node := newNode(item.Schema)
		root.Content = append(root.Content, newKeyNode(item.Name, item.Schema.Description))
		root.Content = append(root.Content, node)
	}
	return &root
}

func newScalarNode(schema spec.Schema) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: getScalarValue(&schema),
	}
}

func getScalarValue(schema *spec.Schema) string {
	valueType := schema.Type[0]
	switch valueType {
	case "boolean":
		return "<true|false>"
	case "string":
		if schema.Enum != nil {
			enumValues := []string{}
			for _, value := range schema.Enum {
				enumValues = append(enumValues, value.(string))
			}
			return fmt.Sprintf("<%s>", strings.Join(enumValues, "|"))
		}
		if schema.Format != "" {
			return fmt.Sprintf("<%s>", schema.Format)
		}
		return "<string>"
	case "number":
		return "<number>"
	case "integer":
		return "<integer>"
	}
	return ""
}
