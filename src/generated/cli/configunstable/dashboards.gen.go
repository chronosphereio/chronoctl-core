// Code generated by chronogen; DO NOT EDIT
package configunstable

import (
	"context"
	"fmt"

	"github.com/chronosphereio/chronoctl-core/src/cmd/cli"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/clienterror"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/dry"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/file"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/groups"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/output"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/pagination"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/ptr"
	config_unstable "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/models"
	"github.com/chronosphereio/chronoctl-core/src/types"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

func init() { types.MustRegisterObject(DashboardTypeMeta, &Dashboard{}) }

var _ types.Object = &Dashboard{}

var DashboardTypeMeta = types.TypeMeta{
	APIVersion: "unstable/config",
	Kind:       "Dashboard",
}

type Dashboard struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.ConfigunstableDashboard `json:"spec"`
}

func NewDashboard(spec *models.ConfigunstableDashboard) *Dashboard {
	return &Dashboard{
		TypeMeta: DashboardTypeMeta,
		Spec:     spec,
	}
}

func (e *Dashboard) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *Dashboard) Identifier() string {
	return e.Spec.Slug
}

func CreateDashboard(
	ctx context.Context,
	client config_unstable.ClientService,
	entity *Dashboard,
	dryRun bool,
) (*Dashboard, error) {
	res, err := client.CreateDashboard(&config_unstable.CreateDashboardParams{
		Context: ctx,
		Body: &models.ConfigunstableCreateDashboardRequest{
			Dashboard: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewDashboard(res.Payload.Dashboard), nil
}

func newDashboardCreateCmd() *cobra.Command {
	var (
		permissiveParsing bool
		dryRunFlags       = dry.NewFlags()
		clientFlags       = client.NewClientFlags()
		outputFlags       = output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
		fileFlags         = file.NewFlags(true /* required */)
	)

	var (
		use   string
		short string
	)
	use = "create -f <file>"
	short = "Creates a single Dashboard."

	cmd := &cobra.Command{
		Use:     use,
		GroupID: groups.Commands.ID,
		Short:   short,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), clientFlags.Timeout())
			defer cancel()
			if err := outputFlags.Validate(); err != nil {
				return err
			}
			defer outputFlags.Close(cmd.OutOrStdout())
			stderr := output.NewStderrPrinter(cmd)

			client, err := clientFlags.ConfigUnstableClient()
			if err != nil {
				return err
			}

			var dashboard *Dashboard
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			dashboard, err = types.MustDecodeSingleObject[*Dashboard](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullDashboard, err := CreateDashboard(ctx, client, dashboard, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("Dashboard is valid and can be created")
				return nil
			}
			stderr.Printf("Dashboard with slug %q created successfully\n", fullDashboard.Spec.Slug)

			if err := outputFlags.WriteObject(fullDashboard, cmd.OutOrStdout()); err != nil {
				return err
			}
			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	fileFlags.AddFlags(cmd)
	cmd.Flags().BoolVar(&permissiveParsing, "no-strict", false, "If set, manifests with unknown fields are not rejected. Defaults to false.")

	return cmd
}

func GetDashboard(
	ctx context.Context,
	client config_unstable.ClientService,
	slug string,
) (*Dashboard, error) {
	res, err := client.ReadDashboard(&config_unstable.ReadDashboardParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewDashboard(res.GetPayload().Dashboard), nil
}

func newDashboardReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single Dashboard by slug"
	use = "read <slug>"
	args = cobra.ExactArgs(1)

	cmd := &cobra.Command{
		Use:     use,
		GroupID: groups.Commands.ID,
		Short:   short,
		Args:    args,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), clientFlags.Timeout())
			defer cancel()
			if err := outputFlags.Validate(); err != nil {
				return err
			}
			defer outputFlags.Close(cmd.OutOrStdout())

			client, err := clientFlags.ConfigUnstableClient()
			if err != nil {
				return err
			}
			entity, err := GetDashboard(ctx, client, args[0])
			if err != nil {
				return err
			}
			if err := outputFlags.WriteObject(entity, cmd.OutOrStdout()); err != nil {
				return err
			}
			return nil
		},
	}

	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)

	return cmd
}

func UpdateDashboard(
	ctx context.Context,
	client config_unstable.ClientService,
	entity *Dashboard,
	opts UpdateOptions,
) (*Dashboard, error) {
	res, err := client.UpdateDashboard(&config_unstable.UpdateDashboardParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: config_unstable.UpdateDashboardBody{
			CreateIfMissing: opts.CreateIfMissing,
			Dashboard:       entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewDashboard(res.Payload.Dashboard), nil
}

func newDashboardUpdateCmd() *cobra.Command {
	var (
		permissiveParsing bool
		createIfMissing   bool
		dryRunFlags       = dry.NewFlags()
		clientFlags       = client.NewClientFlags()
		outputFlags       = output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
		fileFlags         = file.NewFlags(true /* required */)
	)

	cmd := &cobra.Command{
		Use:     "update -f <filename>",
		GroupID: groups.Commands.ID,
		Short:   "Updates an existing Dashboard.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), clientFlags.Timeout())
			defer cancel()
			if err := outputFlags.Validate(); err != nil {
				return err
			}
			defer outputFlags.Close(cmd.OutOrStdout())
			stderr := output.NewStderrPrinter(cmd)

			client, err := clientFlags.ConfigUnstableClient()
			if err != nil {
				return err
			}

			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck

			dashboard, err := types.MustDecodeSingleObject[*Dashboard](file, permissiveParsing)
			if err != nil {
				return err
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRunFlags.DryRun,
				CreateIfMissing: createIfMissing,
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set, update not persisted")
			}

			fullDashboard, err := UpdateDashboard(ctx, client, dashboard, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("Dashboard is valid and can be updated")
				return nil
			}
			stderr.Printf("Dashboard with slug %q applied successfully\n", fullDashboard.Spec.Slug)

			if err := outputFlags.WriteObject(fullDashboard, cmd.OutOrStdout()); err != nil {
				return err
			}
			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	fileFlags.AddFlags(cmd)
	cmd.Flags().BoolVar(&permissiveParsing, "no-strict", false, "If set, manifests with unknown fields are allowed. Defaults to false.")
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the Dashboard if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteDashboard(
	ctx context.Context,
	client config_unstable.ClientService,
	slug string,
) error {
	_, err := client.DeleteDashboard(&config_unstable.DeleteDashboardParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newDashboardDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single Dashboard by slug",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), clientFlags.Timeout())
			defer cancel()
			if err := outputFlags.Validate(); err != nil {
				return err
			}
			defer outputFlags.Close(cmd.OutOrStdout())

			client, err := clientFlags.ConfigUnstableClient()
			if err != nil {
				return err
			}

			res, err := client.DeleteDashboard(&config_unstable.DeleteDashboardParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted Dashboard with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type DashboardListOpts struct {
	// Limit represents that maximum number of items we wish to return.
	Limit int
	// PageToken is the pagination token we want to start our request at.
	PageToken string
	// PageMaxSize is the maximum page size to use when making List calls.
	PageMaxSize          int
	CollectionSlugs      []string
	IncludeDashboardJSON bool
	Names                []string
	Slugs                []string
}

func (r *DashboardListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyCollectionSlugs []string
	flags.StringSliceVar(&r.CollectionSlugs, "collection-slugs", emptyCollectionSlugs, "Filters results by collection_slug, where any Dashboard with a matching collection_slug in the given list (and matches all other filters) is returned.")
	var emptyIncludeDashboardJSON bool
	flags.BoolVar(&r.IncludeDashboardJSON, "include-dashboard-json", emptyIncludeDashboardJSON, "Optional flag to populate the dashboard_json of the returned dashboards. By default, dashboard_json will be left empty.")
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any Dashboard with a matching name in the given list (and matches all other filters) is returned.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any Dashboard with a matching slug in the given list (and matches all other filters) is returned.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListDashboards(
	ctx context.Context,
	client config_unstable.ClientService,
	streamer output.Streamer[*Dashboard],
	opts DashboardListOpts,
) (pagination.Token, error) {
	var (
		gotItems    = 0
		nextToken   = opts.PageToken
		pageMaxSize = opts.PageMaxSize
	)

	// Use the limit if it's set, and smaller than a set page size.
	if opts.Limit > 0 && (opts.Limit < pageMaxSize || pageMaxSize == 0) {
		pageMaxSize = opts.Limit
	}

	for {
		res, err := client.ListDashboards(&config_unstable.ListDashboardsParams{
			Context:              ctx,
			PageToken:            &nextToken,
			PageMaxSize:          ptr.Int64(int64(pageMaxSize)),
			CollectionSlugs:      opts.CollectionSlugs,
			IncludeDashboardJSON: &opts.IncludeDashboardJSON,
			Names:                opts.Names,
			Slugs:                opts.Slugs,
		})
		if err != nil {
			return pagination.Token(""), clienterror.Wrap(err)
		}

		for _, v := range res.Payload.Dashboards {
			if err := streamer(NewDashboard(v)); err != nil {
				return pagination.Token(""), err
			}
			gotItems++
		}

		nextToken = res.Payload.Page.NextToken
		if nextToken == "" {
			return pagination.Token(""), nil
		}

		if opts.Limit > 0 && gotItems >= opts.Limit {
			return pagination.Token(nextToken), nil
		}

		pageMaxSize = pagination.CalculatePageSize(pagination.Calculation{
			GotItems:    gotItems,
			MaxItems:    opts.Limit,
			MaxPageSize: len(res.Payload.Dashboards),
		})
	}
}

func newDashboardListCmd() *cobra.Command {
	var listOptions DashboardListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all Dashboards and applies optional filters",
		GroupID: groups.Commands.ID,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), clientFlags.Timeout())
			defer cancel()
			if err := outputFlags.Validate(); err != nil {
				return err
			}

			writer, err := outputFlags.NewWriterManager(cmd.OutOrStdout())
			if err != nil {
				return err
			}
			defer writer.Close()

			client, err := clientFlags.ConfigUnstableClient()
			if err != nil {
				return err
			}

			streamer := output.NewWriteObjectStreamer[*Dashboard](writer)
			nextToken, err := ListDashboards(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional dashboards. To view more, use the next page token or run the full command.",
					NextPageToken: nextToken,
					FullCommand: fmt.Sprintf("%s --page-token %q",
						cli.BuildCommandString(cmd, []string{"page-token"}),
						nextToken),
				}
				if err := outputFlags.WriteObject(nextPage, cmd.OutOrStdout()); err != nil {
					return err
				}
			}

			return nil
		},
	}

	listOptions.registerFlags(cmd.Flags())
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)

	return cmd
}

const DashboardScaffoldYAML = `api_version: unstable/config
kind: Dashboard
spec:
    # Unique identifier of the Dashboard. If slug is not provided, one will be generated based of the name field. Cannot be modified after the Dashboard is created.
    slug: <string>
    # Required name of the Dashboard. May be modified after the Dashboard is created.
    name: <string>
    # Required slug of the collection the dashboard belongs to.
    collection_slug: <string>
    # Required raw JSON of the dashboard.
    dashboard_json: <string>
`

func newDashboardScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), DashboardScaffoldYAML)
		},
	}
	return cmd
}

func NewDashboardCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "dashboards",
		GroupID: groups.Config.ID,
		Short:   "All commands for Dashboards",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newDashboardCreateCmd(),
		newDashboardReadCmd(),
		newDashboardUpdateCmd(),
		newDashboardDeleteCmd(),
		newDashboardListCmd(),
		newDashboardScaffoldCmd(),
	)
	return root
}
