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
	"context"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/chronosphereio/chronoctl-core/src/cmd/cli"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/output"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/pagination"
	state_unstable "github.com/chronosphereio/chronoctl-core/src/generated/swagger/stateunstable/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/stateunstable/models"
)

func newListCommand() *cobra.Command {
	opts := newListOptions()

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all rule evaluations",
		Example: `# List all rule evaluations.
chronoctl rule-evaluations list`,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.originalCommandString = cli.BuildCommandString(cmd, []string{"next-token"} /* ignoreFlags */)

			if err := opts.validate(); err != nil {
				return err
			}
			if err := opts.runByMetric(os.Stdout); err != nil {
				return err
			}
			return nil
		},
	}

	opts.addFlags(cmd)

	return cmd
}

type listOptions struct {
	outputFlags *output.Flags
	clientFlags *client.Flags

	maxItems              int64
	nextToken             string
	originalCommandString string

	client state_unstable.ClientService
}

func newListOptions() *listOptions {
	return &listOptions{
		clientFlags: client.NewClientFlags(),
		outputFlags: output.NewFlags(),
	}
}

func (o *listOptions) validate() error {
	client, err := o.clientFlags.StateUnstableClient()
	if err != nil {
		return err
	}
	o.client = client

	return nil
}

func (o *listOptions) addFlags(cmd *cobra.Command) {
	o.clientFlags.AddFlags(cmd)
	o.outputFlags.AddFlags(cmd)

	cmd.Flags().Int64Var(&o.maxItems, "max-items", 0, "Maximum number of rule evaluations to return. If omitted, all rule evaluations will be returned.")
	cmd.Flags().StringVar(&o.nextToken, "next-token", "", "Pagination token to use")
}

func (o *listOptions) runByMetric(w io.Writer) error {
	defer o.outputFlags.Close(w) //nolint:errcheck

	ctx, cancel := context.WithTimeout(context.Background(), o.clientFlags.Timeout())
	defer cancel()

	byMetric, nextToken, err := o.queryByMetric(ctx)
	if err != nil {
		return err
	}
	for _, ruleEvaluation := range byMetric {
		if err := o.outputFlags.WriteObject(ruleEvaluation, w); err != nil {
			return err
		}
	}

	if nextToken != "" {
		o.outputFlags.WriteAfterClose(fmt.Sprintf("\nThere are additional metric usages. To continue getting more, run: %v --next-token %v\n", o.originalCommandString, nextToken))
	}

	return nil
}

func (o *listOptions) runByLabel(w io.Writer) error {
	defer o.outputFlags.Close(w) //nolint:errcheck

	ctx, cancel := context.WithTimeout(context.Background(), o.clientFlags.Timeout())
	defer cancel()

	byMetric, nextToken, err := o.queryByLabel(ctx)
	if err != nil {
		return err
	}
	for _, ruleEvaluation := range byMetric {
		if err := o.outputFlags.WriteObject(ruleEvaluation, w); err != nil {
			return err
		}
	}

	if nextToken != "" {
		o.outputFlags.WriteAfterClose(fmt.Sprintf("\nThere are additional metric usages. To continue getting more, run: %v --next-token %v\n", o.originalCommandString, nextToken))
	}

	return nil
}

func (o *listOptions) queryByMetric(ctx context.Context) (usages []*models.StateunstableMetricUsageByMetricName, token string, _ error) {
	return pagination.List(
		pagination.Page{
			Size:  o.maxItems,
			Token: o.nextToken,
		},
		func(p pagination.Page) (items []*models.StateunstableMetricUsageByMetricName, token string, err error) {
			resp, err := o.client.ListMetricUsagesByMetricName(&state_unstable.ListMetricUsagesByMetricNameParams{
				Context:     ctx,
				PageMaxSize: &p.Size,
				PageToken:   &p.Token,
			})
			if err != nil {
				return nil, "", err
			}
			t := ""
			if page := resp.GetPayload().Page; page != nil {
				t = page.NextToken
			}
			return resp.GetPayload().Usages, t, nil
		})
}

func (o *listOptions) queryByLabel(ctx context.Context) (usages []*models.StateunstableMetricUsageByLabelName, token string, _ error) {
	return pagination.List(
		pagination.Page{
			Size:  o.maxItems,
			Token: o.nextToken,
		},
		func(p pagination.Page) (items []*models.StateunstableMetricUsageByLabelName, token string, err error) {
			resp, err := o.client.ListMetricUsagesByLabelName(&state_unstable.ListMetricUsagesByLabelNameParams{
				Context:     ctx,
				PageMaxSize: &p.Size,
				PageToken:   &p.Token,
			})
			if err != nil {
				return nil, "", err
			}
			t := ""
			if page := resp.GetPayload().Page; page != nil {
				t = page.NextToken
			}
			return resp.GetPayload().Usages, t, nil
		})
}
