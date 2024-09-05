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

package ruleevaluations

import (
	"bytes"
	"testing"

	state_v1 "github.com/chronosphereio/chronoctl-core/src/generated/swagger/statev1/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/statev1/mocks"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/statev1/models"
	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListRuleEvaluations(t *testing.T) {
	base := models.Statev1RuleEvaluation{
		RuleSlug:   "slug-one",
		RuleType:   models.RuleEvaluationRuleTypeMONITOR,
		Count:      5,
		Message:    "error-one",
		DetectedAt: strfmt.NewDateTime(),
	}

	tests := []struct {
		name     string
		opts     func(*gomock.Controller) *listOptions
		expected string
		err      error
	}{
		{
			name: "get all rule evaluations",
			opts: func(ctrl *gomock.Controller) *listOptions {
				o := newListOptions()
				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ListRuleEvaluations(gomock.Any()).Return(&state_v1.ListRuleEvaluationsOK{
					Payload: &models.Statev1ListRuleEvaluationsResponse{
						RuleEvaluations: []*models.Statev1RuleEvaluation{
							createTestRuleEvaluation(base),
							createTestRuleEvaluation(base, withRuleSlug("slug-two"), withCount(20), withRuleType(models.RuleEvaluationRuleTypeRECORDING)),
						},
					},
				}, nil)
				o.client = cli
				return o
			},
			expected: `rule_slug: slug-one
rule_type: MONITOR
detected_at: "1970-01-01T00:00:00.000Z"
count: 5
message: error-one
---
rule_slug: slug-two
rule_type: RECORDING
detected_at: "1970-01-01T00:00:00.000Z"
count: 20
message: error-one
`,
		},
		{
			name: "get all rule evaluations with pagination",
			opts: func(ctrl *gomock.Controller) *listOptions {
				o := newListOptions()
				cli := mocks.NewMockClientService(ctrl)

				first := cli.EXPECT().ListRuleEvaluations(gomock.Any()).Return(&state_v1.ListRuleEvaluationsOK{
					Payload: &models.Statev1ListRuleEvaluationsResponse{
						RuleEvaluations: []*models.Statev1RuleEvaluation{
							createTestRuleEvaluation(base),
							createTestRuleEvaluation(base, withRuleSlug("slug-two"), withCount(20), withRuleType(models.RuleEvaluationRuleTypeRECORDING)),
						},
						Page: &models.Configv1PageResult{ // indication of next page
							NextToken: "next-page",
						},
					},
				}, nil)

				second := cli.EXPECT().ListRuleEvaluations(gomock.Any()).DoAndReturn(func(params *state_v1.ListRuleEvaluationsParams, _ ...state_v1.ClientOption) (*state_v1.ListRuleEvaluationsOK, error) {
					require.NotEmpty(t, params.PageToken)
					return &state_v1.ListRuleEvaluationsOK{
						Payload: &models.Statev1ListRuleEvaluationsResponse{
							RuleEvaluations: []*models.Statev1RuleEvaluation{
								createTestRuleEvaluation(base, withRuleSlug("slug-three")),
							},
						},
					}, nil
				})

				gomock.InOrder(first, second)
				o.client = cli
				return o
			},
			expected: `rule_slug: slug-one
rule_type: MONITOR
detected_at: "1970-01-01T00:00:00.000Z"
count: 5
message: error-one
---
rule_slug: slug-two
rule_type: RECORDING
detected_at: "1970-01-01T00:00:00.000Z"
count: 20
message: error-one
---
rule_slug: slug-three
rule_type: MONITOR
detected_at: "1970-01-01T00:00:00.000Z"
count: 5
message: error-one
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var (
				buf  bytes.Buffer
				opts = tt.opts(ctrl)
			)
			err := opts.run(&buf)
			require.Equal(t, tt.err, err)
			assert.Equal(t, tt.expected, buf.String())
		})
	}
}

type modifierFn func(models.Statev1RuleEvaluation) models.Statev1RuleEvaluation

func withRuleSlug(slug string) modifierFn {
	return func(sre models.Statev1RuleEvaluation) models.Statev1RuleEvaluation {
		sre.RuleSlug = slug
		return sre
	}
}

func withRuleType(ruleType models.RuleEvaluationRuleType) modifierFn {
	return func(sre models.Statev1RuleEvaluation) models.Statev1RuleEvaluation {
		sre.RuleType = ruleType
		return sre
	}
}

func withCount(count int32) modifierFn {
	return func(sre models.Statev1RuleEvaluation) models.Statev1RuleEvaluation {
		sre.Count = count
		return sre
	}
}

func createTestRuleEvaluation(base models.Statev1RuleEvaluation, modifiers ...modifierFn) *models.Statev1RuleEvaluation {
	for _, modifierFn := range modifiers {
		base = modifierFn(base)
	}

	return &base
}
