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

//nolint:errcheck
package output

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		flags   *Flags
		wantErr string
	}{
		{
			name:  "valid default flags",
			flags: NewFlags(),
		},
		{
			name: "valid custom flags",
			flags: func() *Flags {
				o := NewFlags()
				o.Format = JSONFormat
				return o
			}(),
		},
		{
			name: "invalid output Format",
			flags: func() *Flags {
				o := NewFlags()
				o.Format = "invalid"
				return o
			}(),
			wantErr: "invalid output format 'invalid', must be one of json|jsonl|yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.flags.Validate()
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestWrite(t *testing.T) {
	tests := []struct {
		name       string
		flags      *Flags
		input      interface{}
		multiInput []any
		want       string
		wantErr    string
	}{
		{
			name:  "output in YAML",
			flags: NewFlags(),
			input: &testGoodObject{
				Foo: "foo",
				Bar: 42,
				Baz: map[string]int{
					"a": 1,
				},
			},
			want: `foo: foo
bar: 42
baz:
  a: 1
`,
		},
		{
			name: "output in JSON",
			flags: func() *Flags {
				o := NewFlags()
				o.Format = JSONFormat
				return o
			}(),
			input: &testGoodObject{
				Foo: "foo",
				Bar: 42,
				Baz: map[string]int{
					"a": 1,
				},
			},
			want: `[
  {
    "foo": "foo",
    "bar": 42,
    "baz": {
      "a": 1
    }
  }
]
`,
		},
		{
			name: "output in JSONL",
			flags: func() *Flags {
				o := NewFlags()
				o.Format = JSONLFormat
				return o
			}(),
			input: &testGoodObject{
				Foo: "foo",
				Bar: 42,
				Baz: map[string]int{
					"a": 1,
				},
			},
			want: `{
  "foo": "foo",
  "bar": 42,
  "baz": {
    "a": 1
  }
}
`,
		},
		{
			name: "output multiple objects in JSON ",
			flags: func() *Flags {
				o := NewFlags()
				o.Format = JSONLFormat
				return o
			}(),
			multiInput: []any{
				&testGoodObject{
					Foo: "foo",
					Bar: 42,
					Baz: map[string]int{
						"a": 1,
					},
				},
				&testGoodObject{
					Foo: "bar",
					Bar: 100,
					Baz: map[string]int{
						"c": 2,
					},
				},
			},
			want: `{
  "foo": "foo",
  "bar": 42,
  "baz": {
    "a": 1
  }
}
{
  "foo": "bar",
  "bar": 100,
  "baz": {
    "c": 2
  }
}
`,
		},
		{
			name:  "unable to unmarshal to YAML",
			flags: NewFlags(),
			input: &testBadObject{},
			want: `{}
`,
		},
		{
			name: "unable to unmarshal to JSON",
			flags: func() *Flags {
				o := NewFlags()
				o.Format = JSONFormat
				return o
			}(),
			input:   &testBadObject{},
			wantErr: "unable to marshal object to JSON: json: error calling MarshalJSON for type *output.testBadObject: test error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			var err error
			if tt.multiInput != nil {
				for _, o := range tt.multiInput {
					if err = tt.flags.WriteObject(o, &buf); err != nil {
						break
					}
				}
			} else {
				err = tt.flags.WriteObject(tt.input, &buf)
			}
			tt.flags.Close(&buf)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

type testGoodObject struct {
	Foo string         `json:"foo"`
	Bar int            `json:"bar"`
	Baz map[string]int `json:"baz"`
}

type testBadObject struct{}

func (o *testBadObject) MarshalJSON() ([]byte, error) {
	return nil, errors.New("test error")
}
