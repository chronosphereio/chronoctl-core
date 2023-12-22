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

package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVersionOptionsValidate(t *testing.T) {
	tests := []struct {
		name    string
		opts    *versionOptions
		wantErr string
	}{
		{
			name: "valid options",
			opts: newVersionOptions(),
		},
		{
			name: "invalid options",
			opts: func() *versionOptions {
				o := newVersionOptions()
				o.outputFlags.Format = "invalid"
				return o
			}(),
			wantErr: "invalid output format 'invalid', must be one of json|jsonl|yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestVersionOptionsRun(t *testing.T) {
	tests := []struct {
		name    string
		opts    *versionOptions
		input   versionInfo
		want    string
		wantErr string
	}{
		{
			name: "valid input",
			opts: newVersionOptions(),
			input: versionInfo{
				Revision:  "123456",
				BuildDate: "2019-10-01",
				GoVersion: "1.15",
				Version:   "0.3.0",
			},
			want: `Version: 0.3.0
Revision: "123456"
BuildDate: "2019-10-01"
GoVersion: "1.15"
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tt.opts.run(tt.input, &buf)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}
