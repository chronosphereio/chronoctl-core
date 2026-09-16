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
	"testing"

	"github.com/chronosphereio/chronoctl-core/src/thirdparty/yaml"
	xtest "github.com/chronosphereio/chronoctl-core/src/x/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCliSpecGen(t *testing.T) {
	cliSpec, err := Parse("../specscanner/testdata/service.swagger.json")
	require.NoError(t, err)

	// marshal the spec and compare it to a golden file
	yamlBytes, err := yaml.Marshal(cliSpec)
	assert.NoError(t, err)
	xtest.MustEqualFile(t, xtest.GoldenFilePath(t, "testdata", "expected_clispec.yaml"), string(yamlBytes))
}

// TestParseSkipsActionPaths verifies that a Google-style custom action path
// (e.g. /api/v1/config/monitors:bulkDisable, present in the shared test
// fixture) does not produce an entity. Before the fix, the colon-bearing
// path segment was taken verbatim as an entity name, which produced a
// schema-less entity that later panicked in ScaffoldYAML.
func TestParseSkipsActionPaths(t *testing.T) {
	cliSpec, err := Parse("../specscanner/testdata/service.swagger.json")
	require.NoError(t, err)

	gotNames := make([]string, 0, len(cliSpec.Entities))
	for name := range cliSpec.Entities {
		gotNames = append(gotNames, name)
	}
	assert.ElementsMatch(t, []string{"monitors", "teams"}, gotNames)

	for name := range cliSpec.Entities {
		assert.NotContains(t, name, ":", "entity name should never contain a colon")
	}
}

func TestIsActionPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "plain collection path",
			path: "/api/v1/config/monitors",
			want: false,
		},
		{
			name: "parameterized path",
			path: "/api/v1/config/monitors/{slug}",
			want: false,
		},
		{
			name: "custom action path",
			path: "/api/v1/config/dashboards:importFromClassic",
			want: true,
		},
		{
			name: "unstable custom action path",
			path: "/api/unstable/config/dashboards:createFromClassic",
			want: true,
		},
		{
			name: "hyphenated entity is not an action path",
			path: "/api/unstable/config/sync-prometheus",
			want: false,
		},
		{
			name: "too-short path",
			path: "/api/v1/config",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isActionPath(tt.path))
		})
	}
}
