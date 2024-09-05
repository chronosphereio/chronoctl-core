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

// Package metricusages contains all commands related to metric usages
package metricusages

import (
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/groups"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra command for managing rule evaluations
func NewCommands() []*cobra.Command {
	return []*cobra.Command{
		newCommandByMetricName(),
		newCommandByLabelName(),
	}
}

func newCommandByMetricName() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "metric-usages-by-metric-name",
		GroupID: groups.State.ID,
		Short:   "All commands related to metric usages by metric name",
		Example: `# List all metric usages by metric name.
chronoctl metric-usages-by-metric-name list`,
	}
	cmd.AddCommand(newListOptions(listMetricUsageByMetricName).buildCmd(
		"List all metric usages by metric name.",
		`# List all metric usages by metric name.
chronoctl metric-usages-by-metric-name list`))
	return cmd
}

func newCommandByLabelName() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "metric-usages-by-label-name",
		GroupID: groups.State.ID,
		Short:   "All commands related to metric usages by label name",
		Example: `# List all metric usages by label name.
chronoctl metric-usages-by-label-name list`,
	}
	cmd.AddCommand(newListOptions(listMetricUsageByLabelName).buildCmd(
		"List all metric usages by metric name.",
		`# List all metric usages by metric name.
chronoctl metric-usages-by-metric-name list`))
	return cmd
}
