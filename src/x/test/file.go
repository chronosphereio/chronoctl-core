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

package xtest

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/require"
)

var updateGoldenFiles = flag.Bool(
	"update-golden-files",
	os.Getenv("UPDATE_GOLDEN_FILES") == "1",
	"updates golden files containing expected outputs output of tests. Defaults to true if environment variable UPDATE_GOLDEN_FILES=1",
)

// MustReadFile reads and returns the file contents as string or fails the test
// if an error occurs.
func MustReadFile(t testing.TB, fileName string, msgAndArgs ...any) string {
	t.Helper()

	bytes, err := ioutil.ReadFile(path.Clean(fileName))
	require.NoError(t, err, msgAndArgs...)
	return string(bytes)
}

func writeGoldenFile(t *testing.T, updatedContent, filePath string) {
	t.Helper()

	// The directory may not exist, so create it first.
	require.NoError(t, os.MkdirAll(filepath.Dir(filePath), 0o777),
		"create directory for golden files failed")
	require.NoError(t, ioutil.WriteFile(path.Clean(filePath), []byte(updatedContent), 0o600),
		"failed to write golden file")

	err := ioutil.WriteFile(path.Clean(filePath), []byte(updatedContent+"\n"), 0o600)
	require.NoError(t, err)
}

func updateGoldenFile(t *testing.T, filePath, updatedData string) {
	t.Helper()

	// In our expected out results we have regexp patterns. Here we go through
	// and find anything that will match the regexp patterns and replace them
	// with our pre-defined keys.
	for k, v := range replacementPatterns {
		updatedData = v.ReplaceAllLiteralString(updatedData, k)
	}
	writeGoldenFile(t, updatedData, filePath)
}

func testFileEqual(t *testing.T, expectedFilePath, data string) {
	t.Helper()

	expectedData := MustReadFile(t, expectedFilePath, "couldn't read golden file %s. To create, run with env UPDATE_GOLDEN_FILES=1 or flag -update-golden-files", expectedFilePath)
	if strings.HasSuffix(expectedFilePath, ".json") {
		expectedData = JSONMarshalIndentedString(t, expectedData)
	} else {
		data += "\n"
	}
	pattern := RegexpEscapedPatternWithTypes(t, expectedData)

	match := pattern.MatchString(data)
	require.True(t, match,
		"Golden file mismatch %q. To update, run with env UPDATE_GOLDEN_FILES=1 or flag -update-golden-files. Diff:\n%s",
		expectedFilePath, diff(t, expectedData, data))
}

// MustEqualFile compares pass in data to a recorded file. If updateGoldenFiles
// flag is enabled, the recorded json is updated before the comparison.
// Ex: go test ./src/to/tests/... -update-golden-files
// or
// Ex: UPDATE_GOLDEN_FILES=1 go test ./src/to/tests/...
func MustEqualFile(t *testing.T, expectedFilePath, data string) {
	t.Helper()

	if *updateGoldenFiles {
		updateGoldenFile(t, expectedFilePath, data)
	}

	testFileEqual(t, expectedFilePath, data)
}

// StructMustEqualJSONFile compares passed in data to recorded json. If updateGoldenFiles
// flag is enabled, the recorded json is updated before the comparison.
// Ex: go test ./src/to/tests/... --update-golden-files
func StructMustEqualJSONFile(t *testing.T, expectedFilePath string, data any) {
	dataNormalized := JSONMarshalIndentedString(t, JSONMarshalIndentedValue(t, data))
	MustEqualFile(t, expectedFilePath, dataNormalized)
}

// GoldenFilePath constructs the file path for a golden test file based on the test name,
// test file, and id.
func GoldenFilePath(t *testing.T, root, id string) string {
	testName := strings.ReplaceAll(t.Name(), "/", "_")
	if id == "" {
		return fmt.Sprintf("%s/%s.json", root, testName)
	}

	suffix := ""
	if !strings.Contains(id, ".") {
		suffix = ".json"
	}

	return fmt.Sprintf("%s/%s-%s%s", root, testName, id, suffix)
}

func diff(t testing.TB, expected, actual string) string {
	diff, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(expected),
		B:        difflib.SplitLines(actual),
		FromFile: "Expected",
		FromDate: "",
		ToFile:   "Actual",
		ToDate:   "",
		Context:  1,
	})
	require.NoError(t, err)
	return diff
}
