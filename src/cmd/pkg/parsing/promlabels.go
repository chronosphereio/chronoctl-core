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

// Package parsing contains methods for parsing strings into various types.
package parsing

import (
	"fmt"

	"github.com/alecthomas/participle/v2"

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"

	"github.com/pkg/errors"
)

var (
	promLabelsParser = participle.MustBuild[PromLabels](participle.Unquote())
)

// PromLabelParser returns a parser for PromLabels
func PromLabelParser() *participle.Parser[PromLabels] {
	return promLabelsParser
}

// FromPromLabelsToMutingMatchers converts a string representing labels in Prometheus format to a list of
// muting matchers.
func FromPromLabelsToMutingMatchers(s string) ([]*models.Configv1MutingRuleLabelMatcher, error) {
	if s == "" {
		return nil, nil
	}

	labelsHolder, err := promLabelsParser.ParseString("", s)
	if err != nil {
		return nil, errors.Wrap(err, "invalid labels string")
	}

	if len(labelsHolder.Labels) == 0 {
		return nil, nil
	}

	matchers := make([]*models.Configv1MutingRuleLabelMatcher, len(labelsHolder.Labels))
	for i, matcher := range labelsHolder.Labels {
		matcherType, err := OpToMutingMatcherType(matcher.Op)
		if err != nil {
			return nil, err
		}
		matchers[i] = &models.Configv1MutingRuleLabelMatcher{
			Name:  matcher.Name,
			Value: matcher.Value,
			Type:  matcherType,
		}
	}

	return matchers, nil
}

// PromLabels represents a list of key-value labels surrounded by brackets.
type PromLabels struct {
	Labels []PromLabel `"{" ( @@ ","? )* "}"` //nolint:govet
}

// PromLabel represents a single key-value pair.
type PromLabel struct {
	Name  string `@Ident`                                     //nolint:govet
	Op    string `@(("=" "~") | "=" | ("!" "=") | ("!" "~"))` //nolint:govet
	Value string `@String`                                    //nolint:govet
}

// OpToMutingMatcherType converts an op string to a entities MutingMatcherType
func OpToMutingMatcherType(op string) (models.Configv1MutingRuleLabelMatcherMatcherType, error) {
	switch op {
	case "=":
		return models.Configv1MutingRuleLabelMatcherMatcherTypeEXACT, nil
	case "=~":
		return models.Configv1MutingRuleLabelMatcherMatcherTypeREGEX, nil
	case "!=":
		return models.Configv1MutingRuleLabelMatcherMatcherTypeNOTEXACT, nil
	case "!~":
		return models.Configv1MutingRuleLabelMatcherMatcherTypeNOTREGEXP, nil
	default:
		return models.Configv1MutingRuleLabelMatcherMatcherType(""), fmt.Errorf("invalid matcher %q", op)
	}
}
