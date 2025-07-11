// Code generated by chronogen; DO NOT EDIT
package configv1

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
	config_v1 "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
	"github.com/chronosphereio/chronoctl-core/src/types"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

func init() {
	types.MustRegisterObject(TraceJaegerRemoteSamplingStrategyTypeMeta, &TraceJaegerRemoteSamplingStrategy{})
}

var _ types.Object = &TraceJaegerRemoteSamplingStrategy{}

var TraceJaegerRemoteSamplingStrategyTypeMeta = types.TypeMeta{
	APIVersion: "v1/config",
	Kind:       "TraceJaegerRemoteSamplingStrategy",
}

type TraceJaegerRemoteSamplingStrategy struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.Configv1TraceJaegerRemoteSamplingStrategy `json:"spec"`
}

func NewTraceJaegerRemoteSamplingStrategy(spec *models.Configv1TraceJaegerRemoteSamplingStrategy) *TraceJaegerRemoteSamplingStrategy {
	return &TraceJaegerRemoteSamplingStrategy{
		TypeMeta: TraceJaegerRemoteSamplingStrategyTypeMeta,
		Spec:     spec,
	}
}

func (e *TraceJaegerRemoteSamplingStrategy) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *TraceJaegerRemoteSamplingStrategy) Identifier() string {
	return e.Spec.Slug
}

func CreateTraceJaegerRemoteSamplingStrategy(
	ctx context.Context,
	client config_v1.ClientService,
	entity *TraceJaegerRemoteSamplingStrategy,
	dryRun bool,
) (*TraceJaegerRemoteSamplingStrategy, error) {
	res, err := client.CreateTraceJaegerRemoteSamplingStrategy(&config_v1.CreateTraceJaegerRemoteSamplingStrategyParams{
		Context: ctx,
		Body: &models.Configv1CreateTraceJaegerRemoteSamplingStrategyRequest{
			DryRun:                            dryRun,
			TraceJaegerRemoteSamplingStrategy: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewTraceJaegerRemoteSamplingStrategy(res.Payload.TraceJaegerRemoteSamplingStrategy), nil
}

func newTraceJaegerRemoteSamplingStrategyCreateCmd() *cobra.Command {
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
	short = "Creates a single TraceJaegerRemoteSamplingStrategy."

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

			client, err := clientFlags.ConfigV1Client()
			if err != nil {
				return err
			}

			var traceJaegerRemoteSamplingStrategy *TraceJaegerRemoteSamplingStrategy
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			traceJaegerRemoteSamplingStrategy, err = types.MustDecodeSingleObject[*TraceJaegerRemoteSamplingStrategy](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullTraceJaegerRemoteSamplingStrategy, err := CreateTraceJaegerRemoteSamplingStrategy(ctx, client, traceJaegerRemoteSamplingStrategy, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("TraceJaegerRemoteSamplingStrategy is valid and can be created")
				return nil
			}
			stderr.Printf("TraceJaegerRemoteSamplingStrategy with slug %q created successfully\n", fullTraceJaegerRemoteSamplingStrategy.Spec.Slug)

			if err := outputFlags.WriteObject(fullTraceJaegerRemoteSamplingStrategy, cmd.OutOrStdout()); err != nil {
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

func GetTraceJaegerRemoteSamplingStrategy(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) (*TraceJaegerRemoteSamplingStrategy, error) {
	res, err := client.ReadTraceJaegerRemoteSamplingStrategy(&config_v1.ReadTraceJaegerRemoteSamplingStrategyParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewTraceJaegerRemoteSamplingStrategy(res.GetPayload().TraceJaegerRemoteSamplingStrategy), nil
}

func newTraceJaegerRemoteSamplingStrategyReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single TraceJaegerRemoteSamplingStrategy by slug"
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

			client, err := clientFlags.ConfigV1Client()
			if err != nil {
				return err
			}
			entity, err := GetTraceJaegerRemoteSamplingStrategy(ctx, client, args[0])
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

func UpdateTraceJaegerRemoteSamplingStrategy(
	ctx context.Context,
	client config_v1.ClientService,
	entity *TraceJaegerRemoteSamplingStrategy,
	opts UpdateOptions,
) (*TraceJaegerRemoteSamplingStrategy, error) {
	res, err := client.UpdateTraceJaegerRemoteSamplingStrategy(&config_v1.UpdateTraceJaegerRemoteSamplingStrategyParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: &models.ConfigV1UpdateTraceJaegerRemoteSamplingStrategyBody{
			CreateIfMissing:                   opts.CreateIfMissing,
			DryRun:                            opts.DryRun,
			TraceJaegerRemoteSamplingStrategy: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewTraceJaegerRemoteSamplingStrategy(res.Payload.TraceJaegerRemoteSamplingStrategy), nil
}

func newTraceJaegerRemoteSamplingStrategyUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing TraceJaegerRemoteSamplingStrategy.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), clientFlags.Timeout())
			defer cancel()
			if err := outputFlags.Validate(); err != nil {
				return err
			}
			defer outputFlags.Close(cmd.OutOrStdout())
			stderr := output.NewStderrPrinter(cmd)

			client, err := clientFlags.ConfigV1Client()
			if err != nil {
				return err
			}

			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck

			traceJaegerRemoteSamplingStrategy, err := types.MustDecodeSingleObject[*TraceJaegerRemoteSamplingStrategy](file, permissiveParsing)
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

			fullTraceJaegerRemoteSamplingStrategy, err := UpdateTraceJaegerRemoteSamplingStrategy(ctx, client, traceJaegerRemoteSamplingStrategy, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("TraceJaegerRemoteSamplingStrategy is valid and can be updated")
				return nil
			}
			stderr.Printf("TraceJaegerRemoteSamplingStrategy with slug %q applied successfully\n", fullTraceJaegerRemoteSamplingStrategy.Spec.Slug)

			if err := outputFlags.WriteObject(fullTraceJaegerRemoteSamplingStrategy, cmd.OutOrStdout()); err != nil {
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
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the TraceJaegerRemoteSamplingStrategy if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteTraceJaegerRemoteSamplingStrategy(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) error {
	_, err := client.DeleteTraceJaegerRemoteSamplingStrategy(&config_v1.DeleteTraceJaegerRemoteSamplingStrategyParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newTraceJaegerRemoteSamplingStrategyDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single TraceJaegerRemoteSamplingStrategy by slug",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), clientFlags.Timeout())
			defer cancel()
			if err := outputFlags.Validate(); err != nil {
				return err
			}
			defer outputFlags.Close(cmd.OutOrStdout())

			client, err := clientFlags.ConfigV1Client()
			if err != nil {
				return err
			}

			res, err := client.DeleteTraceJaegerRemoteSamplingStrategy(&config_v1.DeleteTraceJaegerRemoteSamplingStrategyParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted TraceJaegerRemoteSamplingStrategy with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type TraceJaegerRemoteSamplingStrategyListOpts struct {
	// Limit represents that maximum number of items we wish to return.
	Limit int
	// PageToken is the pagination token we want to start our request at.
	PageToken string
	// PageMaxSize is the maximum page size to use when making List calls.
	PageMaxSize           int
	NameOrServiceContains string
	Names                 []string
	ServiceNames          []string
	Slugs                 []string
}

func (r *TraceJaegerRemoteSamplingStrategyListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyNameOrServiceContains string
	flags.StringVar(&r.NameOrServiceContains, "name-or-service-contains", emptyNameOrServiceContains, "")
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any TraceJaegerRemoteSamplingStrategy with a matching name in the given list (and matches all other filters) is returned.")
	var emptyServiceNames []string
	flags.StringSliceVar(&r.ServiceNames, "service-names", emptyServiceNames, "Filters results by service_name, where any TraceJaegerRemoteSamplingStrategy with a matching service_name in the given list (and matches all other filters) is returned.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any TraceJaegerRemoteSamplingStrategy with a matching slug in the given list (and matches all other filters) is returned.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListTraceJaegerRemoteSamplingStrategies(
	ctx context.Context,
	client config_v1.ClientService,
	streamer output.Streamer[*TraceJaegerRemoteSamplingStrategy],
	opts TraceJaegerRemoteSamplingStrategyListOpts,
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
		res, err := client.ListTraceJaegerRemoteSamplingStrategies(&config_v1.ListTraceJaegerRemoteSamplingStrategiesParams{
			Context:               ctx,
			PageToken:             &nextToken,
			PageMaxSize:           ptr.Int64(int64(pageMaxSize)),
			NameOrServiceContains: &opts.NameOrServiceContains,
			Names:                 opts.Names,
			ServiceNames:          opts.ServiceNames,
			Slugs:                 opts.Slugs,
		})
		if err != nil {
			return pagination.Token(""), clienterror.Wrap(err)
		}

		for _, v := range res.Payload.TraceJaegerRemoteSamplingStrategies {
			if err := streamer(NewTraceJaegerRemoteSamplingStrategy(v)); err != nil {
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
			MaxPageSize: len(res.Payload.TraceJaegerRemoteSamplingStrategies),
		})
	}
}

func newTraceJaegerRemoteSamplingStrategyListCmd() *cobra.Command {
	var listOptions TraceJaegerRemoteSamplingStrategyListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all TraceJaegerRemoteSamplingStrategies and applies optional filters",
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

			client, err := clientFlags.ConfigV1Client()
			if err != nil {
				return err
			}

			streamer := output.NewWriteObjectStreamer[*TraceJaegerRemoteSamplingStrategy](writer)
			nextToken, err := ListTraceJaegerRemoteSamplingStrategies(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional traceJaegerRemoteSamplingStrategies. To view more, use the next page token or run the full command.",
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

const TraceJaegerRemoteSamplingStrategyScaffoldYAML = `api_version: v1/config
kind: TraceJaegerRemoteSamplingStrategy
spec:
    # Unique identifier of the TraceJaegerRemoteSamplingStrategy. If a 'slug' isn't provided, one will be generated based of the 'name' field. You can't modify this field after the TraceJaegerRemoteSamplingStrategy is created.
    slug: <string>
    # Required. Name of the TraceJaegerRemoteSamplingStrategy. You can modify this value after the TraceJaegerRemoteSamplingStrategy is created.
    name: <string>
    # The name of the service this sampling strategy applies to. This must match the slug and name fields.
    service_name: <string>
    applied_strategy:
        per_operation_strategies:
            # Defines the service-wide sampling probability (in the range [0, 1]) when specific operations are not matched.
            default_sampling_rate: <number>
            # Defines a minimum number of traces to send for ANY operation in the service, regardless of matching per operation strategy.
            default_lower_bound_traces_per_second: <number>
            # Defines a maximum number of traces to send for ANY operation in the service, regardless of matching per operation strategy.
            default_upper_bound_traces_per_second: <number>
            # Defines explicit operations-specific strategies that take precedence over the default sampling rate.
            per_operation_strategies:
                - # The operation to which this specific strategy should apply.
                  operation: <string>
                  probabilistic_sampling_strategy:
                    # Value in the range [0, 1] that defines the probability of sampling any trace.
                    sampling_rate: <number>
        probabilistic_strategy:
            # Value in the range [0, 1] that defines the probability of sampling any trace.
            sampling_rate: <number>
        rate_limiting_strategy:
            # Maximum number of traces to sample per second.
            max_traces_per_second: <integer>
`

func newTraceJaegerRemoteSamplingStrategyScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), TraceJaegerRemoteSamplingStrategyScaffoldYAML)
		},
	}
	return cmd
}

func NewTraceJaegerRemoteSamplingStrategyCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "trace-jaeger-remote-sampling-strategies",
		GroupID: groups.Config.ID,
		Short:   "All commands for TraceJaegerRemoteSamplingStrategies",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newTraceJaegerRemoteSamplingStrategyCreateCmd(),
		newTraceJaegerRemoteSamplingStrategyReadCmd(),
		newTraceJaegerRemoteSamplingStrategyUpdateCmd(),
		newTraceJaegerRemoteSamplingStrategyDeleteCmd(),
		newTraceJaegerRemoteSamplingStrategyListCmd(),
		newTraceJaegerRemoteSamplingStrategyScaffoldCmd(),
	)
	return root
}
