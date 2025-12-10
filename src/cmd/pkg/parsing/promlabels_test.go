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

package parsing

import (
	"testing"

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromPromLabelsToMutingRuleLabelMatchers(t *testing.T) {
	testCases := []struct {
		expr     string
		expected []*models.Configv1MutingRuleLabelMatcher
		err      string
	}{
		{
			expr:     ``,
			expected: nil,
		},
		{
			expr:     `{}`,
			expected: nil,
		},
		{
			expr: `{foo="bar"}`,
			expected: []*models.Configv1MutingRuleLabelMatcher{
				{Name: "foo", Value: "bar", Type: models.Configv1MutingRuleLabelMatcherMatcherTypeEXACT},
			},
		},
		{
			expr: `{"foo.😎.bar"="baz", "a=b=c"!="def"}`,
			expected: []*models.Configv1MutingRuleLabelMatcher{
				{Name: "foo.😎.bar", Value: "baz", Type: models.Configv1MutingRuleLabelMatcherMatcherTypeEXACT},
				{Name: "a=b=c", Value: "def", Type: models.Configv1MutingRuleLabelMatcherMatcherTypeNOTEXACT},
			},
		},
		{
			expr: `{foo="a\"b=c"}`,
			expected: []*models.Configv1MutingRuleLabelMatcher{
				{Name: "foo", Value: `a"b=c`, Type: models.Configv1MutingRuleLabelMatcherMatcherTypeEXACT},
			},
		},
		{
			expr: `{a="aa",b="bb"}`,
			expected: []*models.Configv1MutingRuleLabelMatcher{
				{Name: "a", Value: "aa", Type: models.Configv1MutingRuleLabelMatcherMatcherTypeEXACT},
				{Name: "b", Value: "bb", Type: models.Configv1MutingRuleLabelMatcherMatcherTypeEXACT},
			},
		},
		{
			expr: `{`,
			err:  "invalid labels string: 1:2: unexpected token \"<EOF>\" (expected \"}\")",
		},
		{
			expr: `{foo`,
			err:  "invalid labels string: 1:2: unexpected token \"foo\" (expected \"}\")",
		},
		{
			expr: `{foo=`,
			err:  "invalid labels string: 1:6: unexpected token \"<EOF>\" (expected <string>)",
		},
		{
			expr: `{foo=}`,
			err:  "invalid labels string: 1:6: unexpected token \"}\" (expected <string>)",
		},
		{
			expr: `{foo="bar}`,
			err:  "invalid labels string: 1:11: literal not terminated",
		},
		{
			expr: `{foo=~"bar"}`,
			expected: []*models.Configv1MutingRuleLabelMatcher{
				{Name: "foo", Value: "bar", Type: models.Configv1MutingRuleLabelMatcherMatcherTypeREGEX},
			},
		},
		{
			expr: `{foo=~"bar",foo2="bar2"}`,
			expected: []*models.Configv1MutingRuleLabelMatcher{
				{Name: "foo", Value: "bar", Type: models.Configv1MutingRuleLabelMatcherMatcherTypeREGEX},
				{Name: "foo2", Value: "bar2", Type: models.Configv1MutingRuleLabelMatcherMatcherTypeEXACT},
			},
		},
		{
			expr: `{foo!="bar"}`,
			expected: []*models.Configv1MutingRuleLabelMatcher{
				{Name: "foo", Value: "bar", Type: models.Configv1MutingRuleLabelMatcherMatcherTypeNOTEXACT},
			},
		},
		{
			expr: `{foo!~"bar"}`,
			expected: []*models.Configv1MutingRuleLabelMatcher{
				{Name: "foo", Value: "bar", Type: models.Configv1MutingRuleLabelMatcherMatcherTypeNOTREGEXP},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.expr, func(t *testing.T) {
			labels, err := FromPromLabelsToMutingMatchers(tt.expr)
			if tt.err != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, labels)
			}
		})
	}
}
