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
		Use:     "import-from-classic -f <filename>",
		GroupID: groups.Commands.ID,
		Short:   "Imports a classic dashboard as a Dashboard.",
		Long:    `Imports a classic dashboard as a Dashboard. -f accepts classic dashboard JSON or a single ClassicDashboard manifest, whose fields --collection-slug, --name and --slug override. The new Dashboard is printed to stdout.`,
		Example: `# Import classic dashboard JSON into a collection.
chronoctl dashboards import-from-classic -f dashboard.json --collection-slug my-collection

# Import an existing classic dashboard under a new slug.
chronoctl classic-dashboards read my-dashboard | yq '.spec.slug = "my-dashboard-v2"' | chronoctl dashboards import-from-classic -f -

# Preview the import without creating anything.
chronoctl dashboards import-from-classic -f dashboard.json --collection-slug my-collection --dry-run`,
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

	cmd.Flags().StringVar(&o.collectionSlug, "collection-slug", "", "Slug of the collection for the new dashboard. Required for classic dashboard JSON; overrides a manifest's collection.")
	cmd.Flags().StringVar(&o.name, "name", "", "Name of the new dashboard. Overrides a manifest's name; defaults to the classic dashboard's title.")
	cmd.Flags().StringVar(&o.slug, "slug", "", "Slug of the new dashboard. Overrides a manifest's slug; generated from the name if unset. Required with --update-if-exists.")
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
