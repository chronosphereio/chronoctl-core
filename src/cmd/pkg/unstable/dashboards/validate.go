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

package dashboards

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/clienterror"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/file"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/groups"
	config_unstable "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/models"
)

func newValidateCommand() *cobra.Command {
	o := newValidateOptions()

	cmd := &cobra.Command{
		Use:     "validate -f <file>",
		GroupID: groups.Commands.ID,
		Short:   "Validates raw dashboard JSON without persisting it.",
		Long: `Validates raw dashboard JSON against the API without creating or updating
anything.

-f takes the raw JSON of a dashboard -- the same JSON a Dashboard's
dashboard_json field holds -- NOT a chronoctl manifest. Supplying "-" reads
from stdin.

When the API accepts the dashboard the command prints "Dashboard is valid"
and exits 0. When it does not, the command exits non-zero and the error
carries the validation failures reported by the API.`,
		Example: `# Validate a dashboard JSON file.
chronoctl unstable dashboards validate -f dashboard.json

# Validate dashboard JSON from stdin.
cat dashboard.json | chronoctl unstable dashboards validate -f -`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.validate(); err != nil {
				return err
			}
			return o.run(cmd.OutOrStdout())
		},
	}

	o.addFlags(cmd)
	return cmd
}

type validateOptions struct {
	clientFlags *client.Flags
	fileFlags   *file.Flags

	dashboardJSON string

	client config_unstable.ClientService
}

func newValidateOptions() *validateOptions {
	return &validateOptions{
		clientFlags: client.NewClientFlags(),
		fileFlags:   file.NewFlags(true /* required */),
	}
}

func (o *validateOptions) addFlags(cmd *cobra.Command) {
	o.clientFlags.AddFlags(cmd)
	o.fileFlags.AddFlags(cmd)
}

func (o *validateOptions) validate() error {
	if o.fileFlags.Filename == "" {
		return errors.New("--filename is required")
	}
	f, err := o.fileFlags.File()
	if err != nil {
		return err
	}
	if f != os.Stdin {
		defer f.Close() //nolint:errcheck
	}
	content, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("could not read dashboard file: %v", err)
	}
	if len(content) == 0 {
		return errors.New("dashboard file is empty")
	}
	o.dashboardJSON = string(content)

	if o.client == nil {
		c, err := o.clientFlags.ConfigUnstableClient()
		if err != nil {
			return err
		}
		o.client = c
	}

	return nil
}

func (o *validateOptions) run(w io.Writer) error {
	ctx, cancel := context.WithTimeout(context.Background(), o.clientFlags.Timeout())
	defer cancel()

	_, err := o.client.ValidateDashboard(&config_unstable.ValidateDashboardParams{
		Context: ctx,
		Body: &models.ConfigunstableValidateDashboardRequest{
			DashboardJSON: o.dashboardJSON,
		},
	})
	if err != nil {
		return clienterror.Wrap(err)
	}

	_, err = fmt.Fprintln(w, "Dashboard is valid")
	return err
}
