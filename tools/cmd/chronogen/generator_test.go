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

package main

import (
	"os/exec"
	"testing"

	xtest "github.com/chronosphereio/chronoctl-core/src/x/test"
	"github.com/chronosphereio/chronoctl-core/tools/pkg/clispec"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	const (
		cliSpecPath = "../../pkg/specscanner/testdata/service.swagger.json"
		testCLIDir  = "./testdata/testcli"
	)
	cliSpec, err := clispec.Parse(cliSpecPath)
	require.NoError(t, err)
	files, err := generate("configv1", cliSpec)
	require.NoError(t, err)
	for _, file := range files {
		xtest.MustEqualFile(t, xtest.GoldenFilePath(t, testCLIDir, file.name+".go"), file.content)
	}

	// Validate that the golden files compile correctly.
	cmd := exec.Command("go", "build", ".")
	cmd.Dir = testCLIDir
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, "go build failed, could not compile the generated client. Output:\n %s", string(out))
}
