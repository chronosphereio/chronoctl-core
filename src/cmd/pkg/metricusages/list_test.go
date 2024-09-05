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

package metricusages

import (
	"bytes"
	"testing"

	state_unstable "github.com/chronosphereio/chronoctl-core/src/generated/swagger/stateunstable/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/stateunstable/mocks"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/stateunstable/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListMetricUsagesByMetricName(t *testing.T) {
	newFn := func(name string) *models.StateunstableMetricUsageByMetricName {
		return &models.StateunstableMetricUsageByMetricName{
			MetricName: name,
			Usage: &models.StateunstableMetricUsage{
				TotalReferences:      1,
				TotalQueryExecutions: 2,
				TotalUniqueUsers:     3,
				UtilityScore:         4,
				ReferenceCountsByType: &models.MetricUsageReferenceCountsByType{
					Dashboards:       1,
					Monitors:         2,
					RecordingRules:   3,
					DropRules:        4,
					AggregationRules: 5,
				},
				QueryExecutionCountsByType: &models.MetricUsageQueryExecutionCountsByType{
					Explorer:  1,
					Dashboard: 2,
					External:  3,
				},
			},
			Cardinality: 1,
			Dpps:        2,
		}
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	o := newListOptions(listMetricUsageByMetricName)
	cli := mocks.NewMockClientService(ctrl)
	o.client = cli

	var buf bytes.Buffer

	// Nil result.
	cli.EXPECT().ListMetricUsagesByMetricName(gomock.Any()).Return(&state_unstable.ListMetricUsagesByMetricNameOK{
		Payload: &models.StateunstableListMetricUsagesByMetricNameResponse{
			Usages: nil,
		},
	}, nil)
	require.NoError(t, o.run(&buf))
	require.Empty(t, buf.String())
	buf.Reset()

	// All results.
	cli.EXPECT().ListMetricUsagesByMetricName(gomock.Any()).Return(&state_unstable.ListMetricUsagesByMetricNameOK{
		Payload: &models.StateunstableListMetricUsagesByMetricNameResponse{
			Usages: []*models.StateunstableMetricUsageByMetricName{
				newFn("metric-a"),
				newFn("metric-b"),
				newFn("metric-c"),
			},
		},
	}, nil)
	require.NoError(t, o.run(&buf))
	expected := `metric_name: metric-a
usage:
  total_references: 1
  total_query_executions: 2
  total_unique_users: 3
  utility_score: 4
  reference_counts_by_type:
    dashboards: 1
    monitors: 2
    recording_rules: 3
    drop_rules: 4
    aggregation_rules: 5
  query_execution_counts_by_type:
    explorer: 1
    dashboard: 2
    external: 3
cardinality: 1
dpps: 2
---
metric_name: metric-b
usage:
  total_references: 1
  total_query_executions: 2
  total_unique_users: 3
  utility_score: 4
  reference_counts_by_type:
    dashboards: 1
    monitors: 2
    recording_rules: 3
    drop_rules: 4
    aggregation_rules: 5
  query_execution_counts_by_type:
    explorer: 1
    dashboard: 2
    external: 3
cardinality: 1
dpps: 2
---
metric_name: metric-c
usage:
  total_references: 1
  total_query_executions: 2
  total_unique_users: 3
  utility_score: 4
  reference_counts_by_type:
    dashboards: 1
    monitors: 2
    recording_rules: 3
    drop_rules: 4
    aggregation_rules: 5
  query_execution_counts_by_type:
    explorer: 1
    dashboard: 2
    external: 3
cardinality: 1
dpps: 2
`
	assert.Equal(t, expected, buf.String())
	buf.Reset()
}
