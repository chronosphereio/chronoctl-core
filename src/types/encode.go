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

package types

import (
	"bytes"
	"fmt"

	"github.com/chronosphereio/chronoctl-core/src/thirdparty/yaml"
)

// EncodeYAML marshals an object as YAML.
func EncodeYAML(v interface{}) ([]byte, error) {
	var node yaml.Node
	if err := node.Encode(v); err != nil {
		return nil, fmt.Errorf("encoding to YAML: %w", err)
	}

	removeEmptyLists(&node)

	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)

	if err := encoder.Encode(node); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// removeEmptyLists mutates a YAML node tree, removing keys from documents where
// the value is an empty list.
func removeEmptyLists(n *yaml.Node) {
	newContent := make([]*yaml.Node, 0, len(n.Content))
	for _, child := range n.Content {
		if n.Kind != yaml.MappingNode {
			newContent = append(newContent, child)
			removeEmptyLists(child)
		} else if child.Kind == yaml.SequenceNode && len(child.Content) == 0 {
			newContent = newContent[0 : len(newContent)-1]
		} else {
			newContent = append(newContent, child)
			removeEmptyLists(child)
		}
	}
	n.Content = newContent
}
