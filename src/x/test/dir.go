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

// Package xtest contains test helpers for Chronoctl.
package xtest

import (
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertDirEqual verifies the contents of files in the specified directories match.
// Note: It does not support recursing into subdirectoreis, and assumes a flat list of files.
func AssertDirEqual(t testing.TB, expectedDir, actualDir string) {
	t.Helper()

	if *updateGoldenFiles {
		updateGoldenDir(t, expectedDir, actualDir)
		return
	}

	expectedNames := readDirNames(t, expectedDir)
	actualNames := readDirNames(t, actualDir)

	assert.Equal(t, expectedNames, actualNames, "list of files mismatched")

	for _, expectedName := range expectedNames {
		expectedContent, err := os.ReadFile(filepath.Join(expectedDir, expectedName))
		require.NoError(t, err, "failed to read expected file")

		actualContent, err := os.ReadFile(filepath.Join(actualDir, expectedName))
		require.NoError(t, err, "failed to read actual file")

		assert.Equal(t, string(expectedContent), string(actualContent), "contents mismatch for %v", expectedName)
	}
}

func updateGoldenDir(t testing.TB, expectedDir, actualDir string) {
	t.Helper()

	actualNames := readDirNames(t, actualDir)
	for _, name := range actualNames {
		content, err := os.ReadFile(filepath.Join(actualDir, name))
		require.NoError(t, err, "failed to read actual file")

		expectedName := filepath.Join(expectedDir, name)
		err = os.WriteFile(expectedName, content, 0o0666)
		require.NoError(t, err, "failed to write expected file")
	}
}

func readDirNames(t testing.TB, dir string) []string {
	t.Helper()

	files, err := os.ReadDir(dir)
	require.NoError(t, err)

	names := make([]string, len(files))
	for i, f := range files {
		names[i] = f.Name()
	}
	sort.Strings(names)
	return names
}
