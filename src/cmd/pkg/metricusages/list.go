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

type listOptions[T any] struct {
	outputFlags *output.Flags
	clientFlags *client.Flags

	maxItems              int64
	nextToken             string
	originalCommandString string

	client state_unstable.ClientService

	listFn listFn[T]
}

type listFn[T any] func(ctx context.Context, client state_unstable.ClientService, p pagination.Page) (items []T, token string, err error)

func newListOptions[T any](listFn listFn[T]) *listOptions[T] {
	return &listOptions[T]{
		clientFlags: client.NewClientFlags(),
		outputFlags: output.NewFlags(),
		listFn:      listFn,
	}
}

func (o *listOptions[T]) buildCmd(short, example string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   short,
		Example: example,
		RunE: func(cmd *cobra.Command, args []string) error {
			o.originalCommandString = cli.BuildCommandString(cmd, []string{"next-token"} /* ignoreFlags */)

			if err := o.validate(); err != nil {
				return err
			}
			if err := o.run(os.Stdout); err != nil {
				return err
			}
			return nil
		},
	}
	o.addFlags(cmd)
	return cmd
}

func (o *listOptions[T]) validate() error {
	client, err := o.clientFlags.StateUnstableClient()
	if err != nil {
		return err
	}
	o.client = client

	return nil
}

func (o *listOptions[T]) addFlags(cmd *cobra.Command) {
	o.clientFlags.AddFlags(cmd)
	o.outputFlags.AddFlags(cmd)

	cmd.Flags().Int64Var(&o.maxItems, "max-items", 0, "Maximum number of rule evaluations to return. If omitted, all rule evaluations will be returned.")
	cmd.Flags().StringVar(&o.nextToken, "next-token", "", "Pagination token to use")
}

func (o *listOptions[T]) run(w io.Writer) error {
	defer o.outputFlags.Close(w) //nolint:errcheck

	ctx, cancel := context.WithTimeout(context.Background(), o.clientFlags.Timeout())
	defer cancel()

	byMetric, nextToken, err := o.query(ctx)
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

func (o *listOptions[T]) query(ctx context.Context) (usages []T, token string, _ error) {
	return pagination.List(
		pagination.Page{
			Size:  o.maxItems,
			Token: o.nextToken,
		},
		func(p pagination.Page) (items []T, token string, err error) {
			return o.listFn(ctx, o.client, p)
		})
}

func listMetricUsageByMetricName(
	ctx context.Context,
	client state_unstable.ClientService,
	p pagination.Page,
) (items []*models.StateunstableMetricUsageByMetricName, token string, err error) {
	resp, err := client.ListMetricUsagesByMetricName(&state_unstable.ListMetricUsagesByMetricNameParams{
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
}

func listMetricUsageByLabelName(
	ctx context.Context,
	client state_unstable.ClientService,
	p pagination.Page,
) (items []*models.StateunstableMetricUsageByLabelName, token string, err error) {
	resp, err := client.ListMetricUsagesByLabelName(&state_unstable.ListMetricUsagesByLabelNameParams{
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
}
