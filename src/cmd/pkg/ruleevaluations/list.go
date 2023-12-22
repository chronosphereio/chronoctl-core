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
	"context"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/chronosphereio/chronoctl-core/src/cmd/cli"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/output"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/pagination"
	state_v1 "github.com/chronosphereio/chronoctl-core/src/generated/swagger/statev1/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/statev1/models"
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
			if err := opts.run(os.Stdout); err != nil {
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

	client state_v1.ClientService
}

func newListOptions() *listOptions {
	return &listOptions{
		clientFlags: client.NewClientFlags(),
		outputFlags: output.NewFlags(),
	}
}

func (o *listOptions) validate() error {
	client, err := o.clientFlags.StateV1Client()
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

func (o *listOptions) run(w io.Writer) error {
	defer o.outputFlags.Close(w) //nolint:errcheck

	ctx, cancel := context.WithTimeout(context.Background(), o.clientFlags.Timeout())
	defer cancel()

	ruleEvaluations, nextToken, err := o.queryRuleEvaluations(ctx)
	if err != nil {
		return err
	}

	// output all results
	for _, ruleEvaluation := range ruleEvaluations {
		if err := o.outputFlags.WriteObject(ruleEvaluation, w); err != nil {
			return err
		}
	}

	if nextToken != "" {
		o.outputFlags.WriteAfterClose(fmt.Sprintf("\nThere are additional rule evaluations. To continue getting more, run: %v --next-token %v\n", o.originalCommandString, nextToken))
	}

	return nil
}

// queryRuleEvaluations fetches all rule evaluations given the input parameters
func (o *listOptions) queryRuleEvaluations(ctx context.Context) (evals []*models.Statev1RuleEvaluation, token string, _ error) {
	var (
		ruleEvaluations []*models.Statev1RuleEvaluation
		pageSize        = o.maxItems
		nextToken       = o.nextToken
	)

	// loop until we've fetched enough results
	for {
		resp, err := o.client.ListRuleEvaluations(&state_v1.ListRuleEvaluationsParams{
			Context:     ctx,
			PageMaxSize: &pageSize,
			PageToken:   &nextToken,
		})
		if err != nil {
			return nil, "", err
		}

		nextToken = getNextToken(resp)
		ruleEvaluations = append(ruleEvaluations, resp.GetPayload().RuleEvaluations...)

		if o.maxItems > 0 && len(ruleEvaluations) >= int(o.maxItems) {
			return ruleEvaluations, nextToken, nil
		}

		// no more results to fetch
		if nextToken == "" {
			return ruleEvaluations, nextToken, nil
		}

		pageSize = int64(pagination.CalculatePageSize(pagination.Calculation{
			GotItems:    len(ruleEvaluations),
			MaxItems:    int(o.maxItems),
			MaxPageSize: len(resp.GetPayload().RuleEvaluations),
		}))
	}
}

// getNextToken safely extracts the next token from the pagination result
func getNextToken(resp *state_v1.ListRuleEvaluationsOK) string {
	if resp.GetPayload().Page == nil {
		return ""
	}

	return resp.GetPayload().Page.NextToken
}
