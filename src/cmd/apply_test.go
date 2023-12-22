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
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testToken           = "token"
	TestChronosphereURL = "https://myorg.chronosphere.io/api"
	TestBadURL          = "://bad.domain.io"
)

func TestApplyOptionsValidate(t *testing.T) {
	tests := []struct {
		name    string
		opts    *ApplyCommand
		wantErr string
	}{
		{
			name: "valid options",
			opts: func() *ApplyCommand {
				o := NewApply(ApplyOptions{})
				o.ClientFlags.APIToken = testToken
				o.ClientFlags.APIUrl = TestChronosphereURL
				o.FileFlags.Filename = "-"
				return o
			}(),
		},
		{
			name: "cannot construct client",
			opts: func() *ApplyCommand {
				o := NewApply(ApplyOptions{})
				o.ClientFlags.APIToken = testToken
				o.ClientFlags.APIUrl = TestBadURL
				o.FileFlags.Filename = "-"
				return o
			}(),
			wantErr: `could not construct Chronosphere config v1 API client: unable to parse URL of Chronosphere URL: parse "://bad.domain.io": missing protocol scheme`,
		},
		{
			name: "file does not exist",
			opts: func() *ApplyCommand {
				o := NewApply(ApplyOptions{})
				o.ClientFlags.APIToken = testToken
				o.ClientFlags.APIUrl = TestChronosphereURL
				o.FileFlags.Filename = "./testdata/does_not_exist.yml"
				return o
			}(),
			wantErr: "could not open file: open ./testdata/does_not_exist.yml: no such file or directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
		})
	}
}
