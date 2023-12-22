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

package shorthand

import (
	"testing"
	"time"

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"
)

func TestMutingRule(t *testing.T) {
	require.True(t, NewMutingRule().IsEmpty())

	start := time.Unix(1, 0).UTC()
	end := time.Unix(2, 0).UTC()

	tests := []struct {
		name     string
		in       mutingRuleFlags
		expected *models.Configv1MutingRule
		wantErr  string
	}{
		{
			name: "valid object",
			in: mutingRuleFlags{
				name:           "my name",
				slug:           "my-slug",
				comment:        "we made a rule",
				startsAt:       start.Format(time.RFC3339),
				endsAt:         end.Format(time.RFC3339),
				mutingMatchers: `{key=~"value"}`,
			},
			expected: &models.Configv1MutingRule{
				Name:     "my name",
				Slug:     "my-slug",
				Comment:  "we made a rule",
				StartsAt: strfmt.DateTime(start),
				EndsAt:   strfmt.DateTime(end),
				LabelMatchers: []*models.Configv1MutingRuleLabelMatcher{
					{
						Type:  models.Configv1MutingRuleLabelMatcherMatcherTypeREGEX,
						Name:  "key",
						Value: "value",
					},
				},
			},
		},
		{
			name: "invalid startsAt",
			in: mutingRuleFlags{
				startsAt: "bad time",
			},
			wantErr: `unable to parse "bad time" as either an RFC3339 timestamp or relative duration`,
		},
		{
			name: "invalid endsAt",
			in: mutingRuleFlags{
				endsAt: "bad time",
			},
			wantErr: `unable to parse "bad time" as either an RFC3339 timestamp or relative duration`,
		},
		{
			name: "invalid matcher",
			in: mutingRuleFlags{
				mutingMatchers: `badmatcher`,
			},
			wantErr: `unexpected token "badmatcher"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec, err := tt.in.GenerateSpec()
			if tt.wantErr == "" {
				require.NoError(t, err)
				require.Equal(t, tt.expected, spec)
			} else {
				require.ErrorContains(t, err, tt.wantErr)
			}
		})
	}
}
