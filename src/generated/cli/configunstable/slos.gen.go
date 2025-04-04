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

func init() { types.MustRegisterObject(SLOTypeMeta, &SLO{}) }

var _ types.Object = &SLO{}

var SLOTypeMeta = types.TypeMeta{
	APIVersion: "unstable/config",
	Kind:       "SLO",
}

type SLO struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.ConfigunstableSLO `json:"spec"`
}

func NewSLO(spec *models.ConfigunstableSLO) *SLO {
	return &SLO{
		TypeMeta: SLOTypeMeta,
		Spec:     spec,
	}
}

func (e *SLO) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *SLO) Identifier() string {
	return e.Spec.Slug
}

func CreateSLO(
	ctx context.Context,
	client config_unstable.ClientService,
	entity *SLO,
	dryRun bool,
) (*SLO, error) {
	res, err := client.CreateSLO(&config_unstable.CreateSLOParams{
		Context: ctx,
		Body: &models.ConfigunstableCreateSLORequest{
			DryRun: dryRun,
			SLO:    entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewSLO(res.Payload.SLO), nil
}

func newSLOCreateCmd() *cobra.Command {
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
	short = "Creates a single SLO."

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

			var sLO *SLO
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			sLO, err = types.MustDecodeSingleObject[*SLO](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullSLO, err := CreateSLO(ctx, client, sLO, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("SLO is valid and can be created")
				return nil
			}
			stderr.Printf("SLO with slug %q created successfully\n", fullSLO.Spec.Slug)

			if err := outputFlags.WriteObject(fullSLO, cmd.OutOrStdout()); err != nil {
				return err
			}
			return nil
		},
	}
	dryRunFlags.AddFlags(cmd)
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	fileFlags.AddFlags(cmd)
	cmd.Flags().BoolVar(&permissiveParsing, "no-strict", false, "If set, manifests with unknown fields are not rejected. Defaults to false.")

	return cmd
}

func GetSLO(
	ctx context.Context,
	client config_unstable.ClientService,
	slug string,
) (*SLO, error) {
	res, err := client.ReadSLO(&config_unstable.ReadSLOParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewSLO(res.GetPayload().SLO), nil
}

func newSLOReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single SLO by slug"
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
			entity, err := GetSLO(ctx, client, args[0])
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

func UpdateSLO(
	ctx context.Context,
	client config_unstable.ClientService,
	entity *SLO,
	opts UpdateOptions,
) (*SLO, error) {
	res, err := client.UpdateSLO(&config_unstable.UpdateSLOParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: &models.ConfigUnstableUpdateSLOBody{
			CreateIfMissing: opts.CreateIfMissing,
			DryRun:          opts.DryRun,
			SLO:             entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewSLO(res.Payload.SLO), nil
}

func newSLOUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing SLO.",
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

			sLO, err := types.MustDecodeSingleObject[*SLO](file, permissiveParsing)
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

			fullSLO, err := UpdateSLO(ctx, client, sLO, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("SLO is valid and can be updated")
				return nil
			}
			stderr.Printf("SLO with slug %q applied successfully\n", fullSLO.Spec.Slug)

			if err := outputFlags.WriteObject(fullSLO, cmd.OutOrStdout()); err != nil {
				return err
			}
			return nil
		},
	}
	dryRunFlags.AddFlags(cmd)
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	fileFlags.AddFlags(cmd)
	cmd.Flags().BoolVar(&permissiveParsing, "no-strict", false, "If set, manifests with unknown fields are allowed. Defaults to false.")
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the SLO if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteSLO(
	ctx context.Context,
	client config_unstable.ClientService,
	slug string,
) error {
	_, err := client.DeleteSLO(&config_unstable.DeleteSLOParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newSLODeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single SLO by slug",
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

			res, err := client.DeleteSLO(&config_unstable.DeleteSLOParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted SLO with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type SLOListOpts struct {
	// Limit represents that maximum number of items we wish to return.
	Limit int
	// PageToken is the pagination token we want to start our request at.
	PageToken string
	// PageMaxSize is the maximum page size to use when making List calls.
	PageMaxSize     int
	CollectionSlugs []string
	Names           []string
	ServiceSlugs    []string
	Slugs           []string
}

func (r *SLOListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyCollectionSlugs []string
	flags.StringSliceVar(&r.CollectionSlugs, "collection-slugs", emptyCollectionSlugs, "")
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any SLO with a matching name in the given list (and matches all other filters) is returned.")
	var emptyServiceSlugs []string
	flags.StringSliceVar(&r.ServiceSlugs, "service-slugs", emptyServiceSlugs, "")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any SLO with a matching slug in the given list (and matches all other filters) is returned.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListSLOs(
	ctx context.Context,
	client config_unstable.ClientService,
	streamer output.Streamer[*SLO],
	opts SLOListOpts,
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
		res, err := client.ListSLOs(&config_unstable.ListSLOsParams{
			Context:         ctx,
			PageToken:       &nextToken,
			PageMaxSize:     ptr.Int64(int64(pageMaxSize)),
			CollectionSlugs: opts.CollectionSlugs,
			Names:           opts.Names,
			ServiceSlugs:    opts.ServiceSlugs,
			Slugs:           opts.Slugs,
		})
		if err != nil {
			return pagination.Token(""), clienterror.Wrap(err)
		}

		for _, v := range res.Payload.Slos {
			if err := streamer(NewSLO(v)); err != nil {
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
			MaxPageSize: len(res.Payload.Slos),
		})
	}
}

func newSLOListCmd() *cobra.Command {
	var listOptions SLOListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all SLOs and applies optional filters",
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

			streamer := output.NewWriteObjectStreamer[*SLO](writer)
			nextToken, err := ListSLOs(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional sLOs. To view more, use the next page token or run the full command.",
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

const SLOScaffoldYAML = `api_version: unstable/config
kind: SLO
spec:
    # Unique identifier of the SLO. If slug is not provided, one will be generated based of the name field. Cannot be modified after the SLO is created.
    slug: <string>
    # Required name of the SLO. May be modified after the SLO is created.
    name: <string>
    description: <string>
    # Optional notification policy to explicitly apply to the generated monitors.
    # If this is not set then the team this SLO will belong to must have a
    # default notification policy
    notification_policy_slug: <string>
    # Annotations are visible in notifications generated by this SLO
    # They can be be templated with labels from notifications.
    annotations:
        key_1: <string>
    # Labels are visible in notifications generated by this SLO,
    # and can be used to route alerts with notification overrides.
    labels:
        key_1: <string>
    collection_ref:
        slug: <string>
        # Type values must match entitiespb.Collection.CollectionType.
        type: <SIMPLE|SERVICE>
    definition:
        # The SLO objective
        objective: <number>
        # This is deprecated.
        reporting_windows:
            - # duration as a string in the format "28d" or "24h", etc.
              duration: <string>
        # Provides the burn rate alert configuration for the SLO. If not provided the
        # default burn rates will be used. The configuration is only valid if the
        # enable_burn_rate_alerting flag is set to true.
        burn_rate_alerting_config:
            - window: <string>
              # budget defines the amount of errors that are allowed during a given
              # time window. It is a value between 0.0 and 100.0 exclusive.
              budget: <number>
              # Severity may only be one of: critical, warn.
              # This is also a string in the monitor schema
              severity: <string>
              # labels to attach to the burn rate alerts.
              labels:
                key_1: <string>
        # If true enables burn rate alerting.
        enable_burn_rate_alerting: <true|false>
        time_window:
            # duration as a string in the format "28d" or "24h", etc.
            duration: <string>
    # SignalGrouping defines how the set of series from the query are split into signals.
    signal_grouping:
        # Set of labels names used to split series into signals.
        # Each unique combination of labels will result in its own signal.
        # For example, if label_names is ["service", "code"] then all series including labels {service="foo",code="404"}
        # will be grouped together in the same signal.

        # Cannot be used if graphite_query is set.
        label_names:
            - <string>
        # If this is true, each series will have its own signal.
        # If this is true then label_names cannot be set.
        signal_per_series: <true|false>
    # SLI is a service level indicator. The contents of this proto are intended to
    # be used to generate queries necessary to measure SLOs.
    sli:
        # Indicates which Chronosphere service discovery template is used to build
        # the SLI queries
        lens_template_indicator: <string>
        # This overrides the equivalent ComponentDiscovery.MetricMetadata field as it
        # is common to change
        endpoint_label: <string>
        # Custom dimensions are used to configure additional labels to export from
        # the underlying queries. This feature provides a logical budget to group
        # unique combination of dimensions. For example, if you want to track a
        # budget per endpoint, add the endpoint label as a dimension. These dimensions
        # are provided on the top-level SLI so that queryful SLOs will receive them
        # in '.GroupBy'.
        custom_dimension_labels:
            - <string>
        # These are added to _every_ query and are intended to be used for things
        # like 'cluster!~"dev"'
        additional_promql_filters:
            - # Prometheus label name for the matcher
              name: <string>
              # Prometheus label value for the matcher
              value: <string>
              type: <MatchEqual|MatchRegexp|MatchNotEqual|MatchNotRegexp>
        custom_indicator:
            # A PromQL query that measures the number of "good" events for this SLI.
            # Either this or the bad_query_template must be set.
            good_query_template: <string>
            # A PromQL query that measures the number of "bad" events for this SLI.
            # Either this or the good_query_template must be set.
            bad_query_template: <string>
            # A PromQL query that measures the total number of events for this SLI.
            # This is required for all advanced SLIs.
            total_query_template: <string>
        # Configuration for an endpoint availability SLI.
        endpoint_availability:
            # the API endpoints to monitor in the SLO. If this is left empty then all
            # endpoints will be monitored.
            endpoints_monitored:
                - <string>
            # A list of result codes that indicate an unsuccessful event. Either this
            # or success_codes must be set.
            error_codes:
                - <string>
            # These are added to _every_ query and are intended to be used for things
            # like 'cluster!~"dev"'
            # deprecated: Use the top level SLI field
            additional_promql_filters:
                - # Prometheus label name for the matcher
                  name: <string>
                  # Prometheus label value for the matcher
                  value: <string>
                  type: <MatchEqual|MatchRegexp|MatchNotEqual|MatchNotRegexp>
        # Configuration for an endpoint latency SLI.
        endpoint_latency:
            # the API endpoints to monitor in the SLO. If this is left empty then all
            # endpoints will be monitored.
            endpoints_monitored:
                - <string>
            # The name of the histogram metric that measures latency.
            latency_bucket: <string>
            # These are added to _every_ query and are intended to be used for things
            # like 'cluster!~"dev"'
            # deprecated: Use the top level SLI field
            additional_promql_filters:
                - # Prometheus label name for the matcher
                  name: <string>
                  # Prometheus label value for the matcher
                  value: <string>
                  type: <MatchEqual|MatchRegexp|MatchNotEqual|MatchNotRegexp>
        source_service:
            slug: <string>
            # Type values must match entitiespb.Collection.CollectionType.
            type: <SIMPLE|SERVICE>
`

func newSLOScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), SLOScaffoldYAML)
		},
	}
	return cmd
}

func NewSLOCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "slos",
		GroupID: groups.Config.ID,
		Short:   "All commands for SLOs",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newSLOCreateCmd(),
		newSLOReadCmd(),
		newSLOUpdateCmd(),
		newSLODeleteCmd(),
		newSLOListCmd(),
		newSLOScaffoldCmd(),
	)
	return root
}
