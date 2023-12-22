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
	"time"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/parsing"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/timesutil"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"

	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

// Creator is the interface for creating entities using CLI flags.
// All custom logic for parsing flags must go into CreateObject.
type Creator[T any] interface {
	// AddFlags registers the flags on the command. Note that we cannot
	// set any of the shorthand commands to be required, as we want
	// to allow `-f` usage.
	AddFlags(*cobra.Command)
	// GenerateSpec parses the flag inputs and returns a spec to create the object.
	GenerateSpec() (T, error)
	// IsEmpty determines if the feature flags have been used.
	// If they aren't, then we go ahead and use the file input.
	IsEmpty() bool
}

type mutingRuleFlags struct {
	name           string
	slug           string
	comment        string
	startsAt       string
	endsAt         string
	mutingMatchers string
}

// NewMutingRule returns a new shorthand create for muting rules
func NewMutingRule() Creator[*models.Configv1MutingRule] {
	return &mutingRuleFlags{}
}

func (m *mutingRuleFlags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&m.name, "name", "", "Name for the muting rule.")
	cmd.Flags().StringVar(&m.slug, "slug", "", "Slug for the muting rule.")
	cmd.Flags().StringVar(&m.comment, "comment", "", "Additional information on the muting rule.")
	cmd.Flags().StringVar(&m.startsAt, "starts-at", time.Now().Format(time.RFC3339), "Time for muting rule to take effect, given in RFC3339 timestamp format. Relative durations are also supported, e.g. +15m, +1h30m, +1d5h15m.")
	cmd.Flags().StringVar(&m.endsAt, "ends-at", "", "Time for muting rule to end effect, given in RFC3339 timestamp format. Relative durations are also supported, e.g. +15m, +1h30m, +1d5h15m.")
	cmd.Flags().StringVar(&m.mutingMatchers, "match-labels", "", "Monitor or series labels to match on.  Each matcher is a string representation of\\nlabel pairs in Prometheus query format and regular expressions are supported: `{key=\\\"value\\\"}`, `{a=\\\"aa\\\",b=\\\"bb\\\"}`,\\n`{c=~\\\"this has a \\\\\\\"quoted string\\\\\\\"\\\"}`.\",\n")
}

func (m *mutingRuleFlags) GenerateSpec() (*models.Configv1MutingRule, error) {
	startsAtRFC3339, err := timesutil.ParseTimestamp(time.Now(), m.startsAt)
	if err != nil {
		return nil, err
	}
	startsAt, err := strfmt.ParseDateTime(startsAtRFC3339)
	if err != nil {
		return nil, err
	}

	endsAtRFC339, err := timesutil.ParseTimestamp(time.Now(), m.endsAt)
	if err != nil {
		return nil, err
	}
	endsAt, err := strfmt.ParseDateTime(endsAtRFC339)
	if err != nil {
		return nil, err
	}

	mutingMatchers, err := parsing.FromPromLabelsToMutingMatchers(m.mutingMatchers)
	if err != nil {
		return nil, err
	}

	return &models.Configv1MutingRule{
		Name:          m.name,
		Slug:          m.slug,
		Comment:       m.comment,
		StartsAt:      startsAt,
		EndsAt:        endsAt,
		LabelMatchers: mutingMatchers,
	}, nil
}

func (m *mutingRuleFlags) IsEmpty() bool {
	// We explicitly ignore starts at as it we default to now.
	return m.endsAt == "" &&
		m.name == "" &&
		m.slug == "" &&
		m.comment == "" &&
		m.mutingMatchers == ""
}
