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
		Use:     "import-from-classic -f <file>",
		GroupID: groups.Commands.ID,
		Short:   "Imports a classic (Grafana) dashboard as a Dashboard.",
		Long: `Imports a classic (Grafana) dashboard as a Dashboard.

-f takes either of:

  * Raw Grafana dashboard JSON, the bare dashboard object a classic dashboard
    exports as. It is sent to the API verbatim, and --collection-slug is
    required.
  * One ClassicDashboard or GrafanaDashboard manifest, such as the output of
    "chronoctl classic-dashboards read <slug>". Its spec.dashboard_json,
    spec.name, spec.slug and collection fill the request.

Each run imports one dashboard. Multi-document input, such as "list" output,
is rejected before the API is called; import one "read" at a time instead.

For a manifest, --collection-slug, --name and --slug override the manifest's
collection, spec.name and spec.slug. Without --collection-slug, the collection
comes from spec.collection_slug, or from spec.collection when it is a simple
collection. A manifest that belongs to a bucket, or to a SERVICE or USER
collection, needs --collection-slug.

A dashboard cannot reuse the slug of a classic dashboard that still exists, so
piping "classic-dashboards read" output in unchanged fails with a conflict.
Set a new spec.slug or pass --slug, or delete spec.slug to have a slug
generated from the name.

The resulting v1/config Dashboard manifest is printed on stdout, ready for
"chronoctl apply -f -". On --dry-run it is a preview of what would be created
rather than a persisted result. A status line and any classic dashboard
features the conversion could not carry over are written to stderr.
Unsupported features do not affect the command's exit code.`,
		Example: `# Import a Grafana dashboard JSON file into a collection.
chronoctl dashboards import-from-classic -f grafana.json --collection-slug team-a

# Import a classic dashboard under a new slug, since a dashboard cannot reuse
# the slug of a classic dashboard that still exists.
chronoctl classic-dashboards read xyz \
  | yq '.spec.slug = "xyz-native"' \
  | chronoctl dashboards import-from-classic -f -

# Preview a Grafana dashboard's conversion into another collection, letting
# the API generate the slug from the name.
chronoctl grafana-dashboards read xyz \
  | yq 'del(.spec.slug)' \
  | chronoctl dashboards import-from-classic -f - --collection-slug migrated --dry-run

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

	collectionSlug    string
	name              string
	slug              string
	updateIfExists    bool
	permissiveParsing bool

	request *importRequest

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

	cmd.Flags().StringVar(&o.collectionSlug, "collection-slug", "", "Slug of the collection the imported dashboard belongs to. Required for raw JSON input; for a manifest, overrides its collection.")
	cmd.Flags().StringVar(&o.name, "name", "", "Name for the new dashboard. Overrides a manifest's spec.name. Defaults to the classic dashboard's title when neither is set.")
	cmd.Flags().StringVar(&o.slug, "slug", "", "Slug for the new dashboard. Overrides a manifest's spec.slug. Generated from the name when neither is set. A slug is required when --update-if-exists is set.")
	cmd.Flags().BoolVar(&o.updateIfExists, "update-if-exists", false, "If true, replaces an existing dashboard with the same slug instead of failing.")
	cmd.Flags().BoolVar(&o.permissiveParsing, "no-strict", false, "If set, manifests with unknown fields are allowed. Defaults to false.")
}

func (o *importFromClassicOptions) validate() error {
	if err := o.outputFlags.Validate(); err != nil {
		return err
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
		return fmt.Errorf("could not read input: %v", err)
	}
	o.request, err = newImportRequest(content, importFlags{
		filename:          o.fileFlags.Filename,
		collectionSlug:    o.collectionSlug,
		name:              o.name,
		slug:              o.slug,
		dryRun:            o.dryRunFlags.DryRun,
		updateIfExists:    o.updateIfExists,
		permissiveParsing: o.permissiveParsing,
	})
	if err != nil {
		return err
	}

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

	resp, err := o.client.ImportDashboardFromClassic(&config_v1.ImportDashboardFromClassicParams{
		Context: ctx,
		Body:    o.request.body,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}

	if resp.Payload.Dashboard == nil {
		return errors.New("API returned no dashboard")
	}

	// Status goes to stderr because stdout carries only the manifest.
	var prefix string
	if o.request.body.DryRun {
		prefix = "[dry run] "
	}
	o.warnf("%sImported %s as Dashboard %q\n", prefix, o.request.source, resp.Payload.Dashboard.Slug)

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
