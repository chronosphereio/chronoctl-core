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

func init() { types.MustRegisterObject(DatasetTypeMeta, &Dataset{}) }

var _ types.Object = &Dataset{}

var DatasetTypeMeta = types.TypeMeta{
	APIVersion: "v1/config",
	Kind:       "Dataset",
}

type Dataset struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.Configv1Dataset `json:"spec"`
}

func NewDataset(spec *models.Configv1Dataset) *Dataset {
	return &Dataset{
		TypeMeta: DatasetTypeMeta,
		Spec:     spec,
	}
}

func (e *Dataset) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *Dataset) Identifier() string {
	return e.Spec.Slug
}

func CreateDataset(
	ctx context.Context,
	client config_v1.ClientService,
	entity *Dataset,
	dryRun bool,
) (*Dataset, error) {
	res, err := client.CreateDataset(&config_v1.CreateDatasetParams{
		Context: ctx,
		Body: &models.Configv1CreateDatasetRequest{
			DryRun:  dryRun,
			Dataset: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewDataset(res.Payload.Dataset), nil
}

func newDatasetCreateCmd() *cobra.Command {
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
	short = "Creates a single Dataset."

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

			var dataset *Dataset
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			dataset, err = types.MustDecodeSingleObject[*Dataset](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullDataset, err := CreateDataset(ctx, client, dataset, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("Dataset is valid and can be created")
				return nil
			}
			stderr.Printf("Dataset with slug %q created successfully\n", fullDataset.Spec.Slug)

			if err := outputFlags.WriteObject(fullDataset, cmd.OutOrStdout()); err != nil {
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

func GetDataset(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) (*Dataset, error) {
	res, err := client.ReadDataset(&config_v1.ReadDatasetParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewDataset(res.GetPayload().Dataset), nil
}

func newDatasetReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single Dataset by slug"
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
			entity, err := GetDataset(ctx, client, args[0])
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

func UpdateDataset(
	ctx context.Context,
	client config_v1.ClientService,
	entity *Dataset,
	opts UpdateOptions,
) (*Dataset, error) {
	res, err := client.UpdateDataset(&config_v1.UpdateDatasetParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: &models.ConfigV1UpdateDatasetBody{
			CreateIfMissing: opts.CreateIfMissing,
			DryRun:          opts.DryRun,
			Dataset:         entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewDataset(res.Payload.Dataset), nil
}

func newDatasetUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing Dataset.",
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

			dataset, err := types.MustDecodeSingleObject[*Dataset](file, permissiveParsing)
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

			fullDataset, err := UpdateDataset(ctx, client, dataset, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("Dataset is valid and can be updated")
				return nil
			}
			stderr.Printf("Dataset with slug %q applied successfully\n", fullDataset.Spec.Slug)

			if err := outputFlags.WriteObject(fullDataset, cmd.OutOrStdout()); err != nil {
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
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the Dataset if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteDataset(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) error {
	_, err := client.DeleteDataset(&config_v1.DeleteDatasetParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newDatasetDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single Dataset by slug",
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

			res, err := client.DeleteDataset(&config_v1.DeleteDatasetParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted Dataset with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type DatasetListOpts struct {
	// Limit represents that maximum number of items we wish to return.
	Limit int
	// PageToken is the pagination token we want to start our request at.
	PageToken string
	// PageMaxSize is the maximum page size to use when making List calls.
	PageMaxSize int
	Names       []string
	Slugs       []string
	Type        string
}

func (r *DatasetListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any Dataset with a matching name in the given list (and matches all other filters) is returned.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any Dataset with a matching slug in the given list (and matches all other filters) is returned.")
	var emptyType string
	flags.StringVar(&r.Type, "type", emptyType, "Custom filtering option: list filtered down to a specific telemetry type.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListDatasets(
	ctx context.Context,
	client config_v1.ClientService,
	streamer output.Streamer[*Dataset],
	opts DatasetListOpts,
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
		res, err := client.ListDatasets(&config_v1.ListDatasetsParams{
			Context:     ctx,
			PageToken:   &nextToken,
			PageMaxSize: ptr.Int64(int64(pageMaxSize)),
			Names:       opts.Names,
			Slugs:       opts.Slugs,
			Type:        &opts.Type,
		})
		if err != nil {
			return pagination.Token(""), clienterror.Wrap(err)
		}

		for _, v := range res.Payload.Datasets {
			if err := streamer(NewDataset(v)); err != nil {
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
			MaxPageSize: len(res.Payload.Datasets),
		})
	}
}

func newDatasetListCmd() *cobra.Command {
	var listOptions DatasetListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all Datasets and applies optional filters",
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

			streamer := output.NewWriteObjectStreamer[*Dataset](writer)
			nextToken, err := ListDatasets(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional datasets. To view more, use the next page token or run the full command.",
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

const DatasetScaffoldYAML = `api_version: v1/config
kind: Dataset
spec:
    # Required name of the Dataset. May be modified after the Dataset is created.
    name: <string>
    # Unique identifier of the Dataset. If slug is not provided, one will be generated based of the name field. Cannot be modified after the Dataset is created.
    slug: <string>
    # Optional longer-form description of the dataset.
    description: <string>
    configuration:
        trace_dataset:
            match_criteria:
                # Each SpanFilter object represents all conditions that need to be true on
                # the same span for the span to be considered matching the SpanFilter. If
                # 'span_count' is used, the number of spans within the trace that match the
                # SpanFilter needs to be within [min, max]. Multiple SpanFilters can be used,
                # and each can be satisfied by any number of spans within the trace.
                span:
                    - # Matches the tags of the candidate.
                      tags:
                        - # The key (or name) of the span tag that is inspected by this filter.
                          key: <string>
                          numeric_value:
                            # The filter value used in comparison against match candidates.
                            value: <number>
                            comparison: <EQUAL|NOT_EQUAL|GREATER_THAN|GREATER_THAN_OR_EQUAL|LESS_THAN|LESS_THAN_OR_EQUAL>
                          value:
                            # The value the filter compares to the target trace or span field.
                            value: <string>
                            # Values the filter tests against when using IN or NOT_IN match type.
                            in_values:
                                - <string>
                            match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION|IN|NOT_IN>
                      duration:
                        # Minimum duration, in seconds, required for a span or trace to match.
                        min_secs: <number>
                        # Maximum duration, in seconds, required for a span or trace to match.
                        max_secs: <number>
                      error:
                        # The value the filter compares to the target trace or span field.
                        value: <true|false>
                      match_type: <INCLUDE|EXCLUDE>
                      operation:
                        # The value the filter compares to the target trace or span field.
                        value: <string>
                        # Values the filter tests against when using IN or NOT_IN match type.
                        in_values:
                            - <string>
                        match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION|IN|NOT_IN>
                      parent_operation:
                        # The value the filter compares to the target trace or span field.
                        value: <string>
                        # Values the filter tests against when using IN or NOT_IN match type.
                        in_values:
                            - <string>
                        match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION|IN|NOT_IN>
                      parent_service:
                        # The value the filter compares to the target trace or span field.
                        value: <string>
                        # Values the filter tests against when using IN or NOT_IN match type.
                        in_values:
                            - <string>
                        match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION|IN|NOT_IN>
                      service:
                        # The value the filter compares to the target trace or span field.
                        value: <string>
                        # Values the filter tests against when using IN or NOT_IN match type.
                        in_values:
                            - <string>
                        match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION|IN|NOT_IN>
                      span_count:
                        # Minimum number of spans that must match a SpanFilter (inclusive).
                        min: <integer>
                        # Maximum number of spans that must match a SpanFilter (inclusive).
                        max: <integer>
                trace:
                    duration:
                        # Minimum duration, in seconds, required for a span or trace to match.
                        min_secs: <number>
                        # Maximum duration, in seconds, required for a span or trace to match.
                        max_secs: <number>
                    error:
                        # The value the filter compares to the target trace or span field.
                        value: <true|false>
        type: <TRACES>
`

func newDatasetScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), DatasetScaffoldYAML)
		},
	}
	return cmd
}

func NewDatasetCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "datasets",
		GroupID: groups.Config.ID,
		Short:   "All commands for Datasets",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newDatasetCreateCmd(),
		newDatasetReadCmd(),
		newDatasetUpdateCmd(),
		newDatasetDeleteCmd(),
		newDatasetListCmd(),
		newDatasetScaffoldCmd(),
	)
	return root
}
