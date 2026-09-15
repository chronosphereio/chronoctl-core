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
	"strings"

	"github.com/spf13/cobra"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/clienterror"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/dry"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/file"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/groups"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/output"
	"github.com/chronosphereio/chronoctl-core/src/generated/cli/configv1"
	config_v1 "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
)

func newImportFromClassicCommand() *cobra.Command {
	o := newImportFromClassicOptions()

	cmd := &cobra.Command{
		Use:     "import-from-classic -f <file> --collection-slug <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Imports a classic (Grafana) dashboard JSON file as a Dashboard.",
		Long: `Imports a classic (Grafana) dashboard JSON file as a Dashboard.

-f takes raw classic dashboard JSON -- the JSON a classic (Grafana) dashboard
exports as -- NOT a chronoctl manifest. The API converts it and the resulting
v1/config Dashboard manifest is printed on stdout, including on --dry-run,
where it is a preview of what would be created rather than a persisted
result.

Classic dashboard features that could not be translated during the
conversion are reported on stderr as warnings. They do not affect the
command's exit code.`,
		Example: `# Import a classic dashboard JSON file into a collection.
chronoctl dashboards import-from-classic -f grafana.json --collection-slug my-collection

# Preview the conversion without persisting anything.
chronoctl dashboards import-from-classic -f grafana.json --collection-slug my-collection --dry-run

# Replace an existing dashboard that already has the given slug.
chronoctl dashboards import-from-classic -f grafana.json --collection-slug my-collection --slug my-dashboard --update-if-exists`,
		RunE: func(cmd *cobra.Command, args []string) error {
			o.stderr = cmd.ErrOrStderr()
			if err := o.validate(); err != nil {
				return err
			}
			return o.run(cmd.OutOrStdout())
		},
	}

	o.addFlags(cmd)
	return cmd
}

type importFromClassicOptions struct {
	clientFlags *client.Flags
	outputFlags *output.Flags
	dryRunFlags *dry.Flags
	fileFlags   *file.Flags

	collectionSlug string
	name           string
	slug           string
	updateIfExists bool

	classicDashboardJSON string

	stderr io.Writer

	client config_v1.ClientService
}

func newImportFromClassicOptions() *importFromClassicOptions {
	return &importFromClassicOptions{
		clientFlags: client.NewClientFlags(),
		outputFlags: output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject()),
		dryRunFlags: dry.NewFlags(),
		fileFlags:   file.NewFlags(true /* required */),
		stderr:      os.Stderr,
	}
}

func (o *importFromClassicOptions) addFlags(cmd *cobra.Command) {
	o.dryRunFlags.AddFlags(cmd)
	o.clientFlags.AddFlags(cmd)
	o.outputFlags.AddFlags(cmd)
	o.fileFlags.AddFlags(cmd)

	cmd.Flags().StringVar(&o.collectionSlug, "collection-slug", "", "Slug of the collection the imported dashboard belongs to.")
	cmd.MarkFlagRequired("collection-slug") //nolint:errcheck,gosec

	cmd.Flags().StringVar(&o.name, "name", "", "Name for the new dashboard. Defaults to the classic dashboard's title when unset.")
	cmd.Flags().StringVar(&o.slug, "slug", "", "Slug for the new dashboard. Generated from the name when unset. Required when --update-if-exists is set.")
	cmd.Flags().BoolVar(&o.updateIfExists, "update-if-exists", false, "If true, replaces an existing dashboard with the same slug instead of failing.")
}

func (o *importFromClassicOptions) validate() error {
	if err := o.outputFlags.Validate(); err != nil {
		return err
	}

	if o.collectionSlug == "" {
		return errors.New("--collection-slug is required")
	}

	if o.updateIfExists && o.slug == "" {
		return errors.New("--slug is required when --update-if-exists is set")
	}

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
		return fmt.Errorf("could not read classic dashboard file: %v", err)
	}
	if len(content) == 0 {
		return errors.New("classic dashboard file is empty")
	}
	o.classicDashboardJSON = string(content)

	if o.client == nil {
		c, err := o.clientFlags.ConfigV1Client()
		if err != nil {
			return err
		}
		o.client = c
	}

	return nil
}

func (o *importFromClassicOptions) run(w io.Writer) error {
	defer o.outputFlags.Close(w) //nolint:errcheck

	ctx, cancel := context.WithTimeout(context.Background(), o.clientFlags.Timeout())
	defer cancel()

	if o.dryRunFlags.DryRun {
		o.warnf("--dry-run is set\n")
	}

	resp, err := o.client.ImportDashboardFromClassic(&config_v1.ImportDashboardFromClassicParams{
		Context: ctx,
		Body: &models.Configv1ImportDashboardFromClassicRequest{
			ClassicDashboardJSON: o.classicDashboardJSON,
			CollectionSlug:       o.collectionSlug,
			Name:                 o.name,
			Slug:                 o.slug,
			DryRun:               o.dryRunFlags.DryRun,
			UpdateIfExists:       o.updateIfExists,
		},
	})
	if err != nil {
		return clienterror.Wrap(err)
	}

	if resp.Payload.Dashboard == nil {
		return errors.New("API returned no dashboard")
	}

	if o.dryRunFlags.DryRun {
		o.warnf("Dashboard is valid and can be imported\n")
	} else {
		o.warnf("Dashboard with slug %q imported successfully\n", resp.Payload.Dashboard.Slug)
	}

	for _, feature := range resp.Payload.UnsupportedFeatures {
		o.warnf("%s\n", formatUnsupportedFeature(feature))
	}

	return o.outputFlags.WriteObject(configv1.NewDashboard(resp.Payload.Dashboard), w)
}

// warnf writes status and warning lines to stderr. Like the rest of the CLI,
// a failed stderr write is not worth failing the command over.
func (o *importFromClassicOptions) warnf(format string, args ...any) {
	fmt.Fprintf(o.stderr, format, args...) //nolint:errcheck
}

// formatUnsupportedFeature renders a single unsupported-feature entry as a
// one-line stderr warning.
func formatUnsupportedFeature(f *models.ImportDashboardFromClassicResponseUnsupportedFeature) string {
	msg := fmt.Sprintf("[%s] %s: %s", f.Level, f.Kind, f.Message)
	if f.Detail != "" && f.Detail != f.Message {
		msg += fmt.Sprintf(" -- %s", f.Detail)
	}

	var parts []string
	if f.Occurrences > 0 {
		unit := "occurrence"
		if f.Occurrences != 1 {
			unit += "s"
		}
		parts = append(parts, fmt.Sprintf("%d %s", f.Occurrences, unit))
	}
	if len(f.PanelKeys) > 0 {
		parts = append(parts, fmt.Sprintf("panels: %s", strings.Join(f.PanelKeys, ", ")))
	}
	if len(parts) > 0 {
		msg += fmt.Sprintf(" (%s)", strings.Join(parts, "; "))
	}

	return msg
}
