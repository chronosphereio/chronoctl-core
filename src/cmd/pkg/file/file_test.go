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

package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testAlertsFile = "./testdata/monitor.yml"
	notExistsFile  = "./testdata/does_not_exist.yml"
)

func TestFileFlagsFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  string
	}{
		{
			name:     "file exists",
			filename: testAlertsFile,
		},
		{
			name:     "stdin",
			filename: "-",
		},
		{
			name:     "file doesn't exist",
			filename: notExistsFile,
			wantErr:  "could not open file: open ./testdata/does_not_exist.yml: no such file or directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flags := Flags{Filename: tt.filename}
			f, err := flags.File()
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)

			if tt.filename == "-" {
				assert.Equal(t, os.Stdin, f)
			}
		})
	}
}
