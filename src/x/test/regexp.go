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
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	replacementPatterns = map[string]*regexp.Regexp{
		// UUID is any version UUID
		`$uuid`: regexp.MustCompile(`(` +
			`([a-f0-9]{32,36})|` +
			`([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})|` +
			`([A-F0-9]{8}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{12})` +
			`)`),

		// Datetime is an ISO8601 date time string with optional decimal part
		`$datetime`: regexp.MustCompile(`(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(\.[0-9]+)?(Z)?`),
	}
)

// RegexpEscapedPatternWithTypes returns a regexp pattern with specific
// type aliases replaced with their regexp patterns and all other characters
// escaped.
// The supported types are:
//   - string: $uuid
//     pattern: "?[a-zA-Z0-9-]{32,36}"?
func RegexpEscapedPatternWithTypes(t *testing.T, input string) *regexp.Regexp {
	replacerInput := make([]string, 0, len(replacementPatterns)*2)
	for k, v := range replacementPatterns {
		replacerInput = append(replacerInput, fmt.Sprintf("\\%s", k))
		replacerInput = append(replacerInput, v.String())
	}
	replacer := strings.NewReplacer(replacerInput...)

	escapedWithTypesReplaced := replacer.Replace(regexp.QuoteMeta(input))
	re, err := regexp.Compile(escapedWithTypesReplaced)
	require.NoError(t, err)

	return re
}
