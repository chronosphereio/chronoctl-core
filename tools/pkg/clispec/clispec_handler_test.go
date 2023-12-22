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
