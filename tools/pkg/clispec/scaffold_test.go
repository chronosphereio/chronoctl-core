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
	"bytes"
	"strings"
	"testing"

	"github.com/chronosphereio/chronoctl-core/src/generated/cli/configv1"
	"github.com/chronosphereio/chronoctl-core/src/thirdparty/yaml"
	"github.com/stretchr/testify/require"
)

func TestScaffoldYAML(t *testing.T) {
	cliSpec, err := Parse("../specscanner/testdata/service.swagger.json")
	require.NoError(t, err)

	monitors := cliSpec.Entities["monitors"]
	rawYAML, err := monitors.ScaffoldYAML()
	require.NoError(t, err)

	rawYAML = strings.ReplaceAll(rawYAML, "<integer>", "0")
	rawYAML = strings.ReplaceAll(rawYAML, "<number>", "0")
	rawYAML = strings.ReplaceAll(rawYAML, "<true|false>", "false")

	roundtrippedMonitor := configv1.Monitor{}

	yamlDecoder := yaml.NewDecoder(bytes.NewBuffer([]byte(rawYAML)))
	yamlDecoder.KnownFields(true)
	require.NoError(t, yamlDecoder.Decode(&roundtrippedMonitor))
}
