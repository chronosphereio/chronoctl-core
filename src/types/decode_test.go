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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		permissiveParsing bool
		want              []Object
		wantErr           string
	}{
		{
			name: "decode single object in YAML format",
			input: `
kind: Test
api_version: v1
name: my-test-object`,
			want: []Object{
				&testObject{
					TypeMeta: TypeMeta{
						Kind:       "Test",
						APIVersion: "v1",
					},
					Name: "my-test-object",
				},
			},
		},
		{
			name: "decode object in YAML format disallows unknown keys when strict mode is on",
			input: `
kind: Test
api_version: v1
name: my-test-object
id: abc-123`,
			wantErr: "unable to decode object: yaml: unmarshal errors:\n  line 5: field id not found in type types.testObject",
		},
		{
			name: "decode object in YAML format allows unknown keys when strict mode is off",
			input: `
kind: Test
api_version: v1
name: my-test-object
id: abc-123`,
			permissiveParsing: true,
			want: []Object{
				&testObject{
					TypeMeta: TypeMeta{
						Kind:       "Test",
						APIVersion: "v1",
					},
					Name: "my-test-object",
				},
			},
		},

		{
			name: "decode single object in JSON format",
			input: `
{
	"kind": "Test",
	"api_version": "v1",
	"name": "my-test-object"
}`,
			want: []Object{
				&testObject{
					TypeMeta: TypeMeta{
						Kind:       "Test",
						APIVersion: "v1",
					},
					Name: "my-test-object",
				},
			},
		}, {
			name: "decode multiple objects in JSONL format",
			input: `
{
	"kind": "Test",
	"api_version": "v1",
	"name": "my-test-object"
}
{
	"kind": "Test",
	"api_version": "v1",
	"name": "my-other-object"
}
`,
			want: []Object{
				&testObject{
					TypeMeta: TypeMeta{
						Kind:       "Test",
						APIVersion: "v1",
					},
					Name: "my-test-object",
				},
				&testObject{
					TypeMeta: TypeMeta{
						Kind:       "Test",
						APIVersion: "v1",
					},
					Name: "my-other-object",
				},
			},
		},
		{
			name: "decode multiple objects in JSON array format",
			input: `
[{
	"kind": "Test",
	"api_version": "v1",
	"name": "my-test-object"
},
{
	"kind": "Test",
	"api_version": "v1",
	"name": "my-other-object"
}]
`,
			want: []Object{
				&testObject{
					TypeMeta: TypeMeta{
						Kind:       "Test",
						APIVersion: "v1",
					},
					Name: "my-test-object",
				},
				&testObject{
					TypeMeta: TypeMeta{
						Kind:       "Test",
						APIVersion: "v1",
					},
					Name: "my-other-object",
				},
			},
		},
		{
			name: "decode in JSON format disallows unknown keys when strict mode is on",
			input: `
{
	"kind": "Test",
	"api_version": "v1",
	"name": "my-test-object",
	"id": 1234
}`,
			wantErr: "unable to decode object: json: unknown field \"id\"",
		},
		{
			name: "decode in JSON format allows unknown keys when strict mode is off",
			input: `
{
	"kind": "Test",
	"api_version": "v1",
	"name": "my-test-object",
	"id": 1234
}`,
			permissiveParsing: true,
			want: []Object{
				&testObject{
					TypeMeta: TypeMeta{
						Kind:       "Test",
						APIVersion: "v1",
					},
					Name: "my-test-object",
				},
			},
		},

		{
			name: "decode multiple objects",
			input: `
---
kind: Test
api_version: v1
name: my-first-object
---
kind: Test
api_version: v1
name: my-second-object
---`,
			want: []Object{
				&testObject{
					TypeMeta: TypeMeta{
						Kind:       "Test",
						APIVersion: "v1",
					},
					Name: "my-first-object",
				},
				&testObject{
					TypeMeta: TypeMeta{
						Kind:       "Test",
						APIVersion: "v1",
					},
					Name: "my-second-object",
				},
			},
		},
		{
			name: "JSON input is malformed",
			input: `
{
	"kind": "Test",
	"api_version": "v1",
	"name": "my-test-object"`,
			wantErr: "unable to decode object: unexpected EOF",
		},
		{
			name: "YAML input is malformed",
			input: `
kind: Test
api_version: v1
name: my-first-object
---
api_version: v1
meh`,
			wantErr: `unable to decode object: yaml: line 7: could not find expected ':'`,
		},
		{
			name: "object isn't registered",
			input: `
{
	"kind": "Test",
	"api_version": "v2",
	"name": "my-test-object"
}`,
			wantErr: "no registered type found for object with type with Kind: Test and API version: v2",
		},
	}

	registry := make(ObjectRegistry)
	o := &testObject{
		TypeMeta: TypeMeta{
			Kind:       "Test",
			APIVersion: "v1",
		},
	}

	require.NoError(t, registry.registerType(o.TypeMeta, o))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			objs, err := decodeWithRegistry(r, registry, tt.permissiveParsing)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, objs)
		})
	}
}

func TestConvertJSONToJSONL(t *testing.T) {
	input := `
[
  {"key": 1},
  {"key": 2}
]
`

	out, err := convertJSONToJSONL([]byte(input))
	require.NoError(t, err)
	assert.Equal(t, string(out), `{"key":1}
{"key":2}
`)
}
