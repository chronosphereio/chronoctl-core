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
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeYAML(t *testing.T) {
	type Author struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	type Book struct {
		FullTitle string    `json:"full_title"`      // "omitempty" is omitted, so this is always marshaled
		Authors   []*Author `json:"authors"`         // Empty lists should be excluded from output
		Pages     int       `yaml:"pages,omitempty"` // Both json and yaml tags can be used
	}

	type Thing struct {
		Thing  *Thing   `yaml:"thing,omitempty"`
		Things []*Thing `yaml:"things,omitempty"`
		Expr   string   `yaml:"expr,omitempty"`
	}

	testCases := []struct {
		input  interface{}
		output string
	}{
		{
			input: Book{
				FullTitle: "Riot",
				Authors:   []*Author{{FirstName: "Isaac", LastName: "Julien"}},
				Pages:     123,
			},
			output: `full_title: Riot
authors:
  - first_name: Isaac
    last_name: Julien
pages: 123
`,
		},
		{
			input: Book{},
			output: `full_title: ""
`,
		},
		// The following tests patches against this issue https://github.com/go-yaml/yaml/issues/643
		{
			input: Thing{
				Thing: &Thing{
					Thing: &Thing{
						Things: []*Thing{
							{
								Expr: "\nsum(rate(rocksdb_block_cache_hits{job=\"crdb\",cluster=\"$cluster\",instance=~\"$node\"}[$rate_interval]))/(sum(rate(rocksdb_block_cache_hits{job=\"crdb\",cluster=\"$cluster\",instance=~\"$node\"}[$rate_interval])) + sum(rate(rocksdb_block_cache_misses{job=\"crdb\",cluster=\"$cluster\",instance=~\"$node\"}[$rate_interval])))",
							},
						},
					},
				},
			},
			output: `thing:
  thing:
    things:
      - expr: "\nsum(rate(rocksdb_block_cache_hits{job=\"crdb\",cluster=\"$cluster\",instance=~\"$node\"}[$rate_interval]))/(sum(rate(rocksdb_block_cache_hits{job=\"crdb\",cluster=\"$cluster\",instance=~\"$node\"}[$rate_interval])) + sum(rate(rocksdb_block_cache_misses{job=\"crdb\",cluster=\"$cluster\",instance=~\"$node\"}[$rate_interval])))"
`,
		},
	}

	for _, tt := range testCases {
		t.Run(spew.Sdump(tt.input), func(t *testing.T) {
			out, err := EncodeYAML(tt.input)
			require.NoError(t, err)

			assert.Equal(t, tt.output, string(out))
		})
	}
}
