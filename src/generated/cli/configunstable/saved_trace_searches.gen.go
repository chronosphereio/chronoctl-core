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

func init() { types.MustRegisterObject(SavedTraceSearchTypeMeta, &SavedTraceSearch{}) }

var _ types.Object = &SavedTraceSearch{}

var SavedTraceSearchTypeMeta = types.TypeMeta{
	APIVersion: "unstable/config",
	Kind:       "SavedTraceSearch",
}

type SavedTraceSearch struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.ConfigunstableSavedTraceSearch `json:"spec"`
}

func NewSavedTraceSearch(spec *models.ConfigunstableSavedTraceSearch) *SavedTraceSearch {
	return &SavedTraceSearch{
		TypeMeta: SavedTraceSearchTypeMeta,
		Spec:     spec,
	}
}

func (e *SavedTraceSearch) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *SavedTraceSearch) Identifier() string {
	return e.Spec.Slug
}

func CreateSavedTraceSearch(
	ctx context.Context,
	client config_unstable.ClientService,
	entity *SavedTraceSearch,
	dryRun bool,
) (*SavedTraceSearch, error) {
	res, err := client.CreateSavedTraceSearch(&config_unstable.CreateSavedTraceSearchParams{
		Context: ctx,
		Body: &models.ConfigunstableCreateSavedTraceSearchRequest{
			SavedTraceSearch: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewSavedTraceSearch(res.Payload.SavedTraceSearch), nil
}

func newSavedTraceSearchCreateCmd() *cobra.Command {
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
	short = "Creates a single SavedTraceSearch."

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

			var savedTraceSearch *SavedTraceSearch
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			savedTraceSearch, err = types.MustDecodeSingleObject[*SavedTraceSearch](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullSavedTraceSearch, err := CreateSavedTraceSearch(ctx, client, savedTraceSearch, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("SavedTraceSearch is valid and can be created")
				return nil
			}
			stderr.Printf("SavedTraceSearch with slug %q created successfully\n", fullSavedTraceSearch.Spec.Slug)

			if err := outputFlags.WriteObject(fullSavedTraceSearch, cmd.OutOrStdout()); err != nil {
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

func GetSavedTraceSearch(
	ctx context.Context,
	client config_unstable.ClientService,
	slug string,
) (*SavedTraceSearch, error) {
	res, err := client.ReadSavedTraceSearch(&config_unstable.ReadSavedTraceSearchParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewSavedTraceSearch(res.GetPayload().SavedTraceSearch), nil
}

func newSavedTraceSearchReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single SavedTraceSearch by slug"
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
			entity, err := GetSavedTraceSearch(ctx, client, args[0])
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

func UpdateSavedTraceSearch(
	ctx context.Context,
	client config_unstable.ClientService,
	entity *SavedTraceSearch,
	opts UpdateOptions,
) (*SavedTraceSearch, error) {
	res, err := client.UpdateSavedTraceSearch(&config_unstable.UpdateSavedTraceSearchParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: config_unstable.UpdateSavedTraceSearchBody{
			CreateIfMissing:  opts.CreateIfMissing,
			SavedTraceSearch: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewSavedTraceSearch(res.Payload.SavedTraceSearch), nil
}

func newSavedTraceSearchUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing SavedTraceSearch.",
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

			savedTraceSearch, err := types.MustDecodeSingleObject[*SavedTraceSearch](file, permissiveParsing)
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

			fullSavedTraceSearch, err := UpdateSavedTraceSearch(ctx, client, savedTraceSearch, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("SavedTraceSearch is valid and can be updated")
				return nil
			}
			stderr.Printf("SavedTraceSearch with slug %q applied successfully\n", fullSavedTraceSearch.Spec.Slug)

			if err := outputFlags.WriteObject(fullSavedTraceSearch, cmd.OutOrStdout()); err != nil {
				return err
			}
			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	fileFlags.AddFlags(cmd)
	cmd.Flags().BoolVar(&permissiveParsing, "no-strict", false, "If set, manifests with unknown fields are allowed. Defaults to false.")
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the SavedTraceSearch if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteSavedTraceSearch(
	ctx context.Context,
	client config_unstable.ClientService,
	slug string,
) error {
	_, err := client.DeleteSavedTraceSearch(&config_unstable.DeleteSavedTraceSearchParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newSavedTraceSearchDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single SavedTraceSearch by slug",
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

			res, err := client.DeleteSavedTraceSearch(&config_unstable.DeleteSavedTraceSearchParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted SavedTraceSearch with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type SavedTraceSearchListOpts struct {
	// Limit represents that maximum number of items we wish to return.
	Limit int
	// PageToken is the pagination token we want to start our request at.
	PageToken string
	// PageMaxSize is the maximum page size to use when making List calls.
	PageMaxSize int
	CreatedBy   string
	Names       []string
	Slugs       []string
}

func (r *SavedTraceSearchListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyCreatedBy string
	flags.StringVar(&r.CreatedBy, "created-by", emptyCreatedBy, "Optional filter: creator email address.")
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any SavedTraceSearch with a matching name in the given list (and matches all other filters) is returned.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any SavedTraceSearch with a matching slug in the given list (and matches all other filters) is returned.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListSavedTraceSearches(
	ctx context.Context,
	client config_unstable.ClientService,
	streamer output.Streamer[*SavedTraceSearch],
	opts SavedTraceSearchListOpts,
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
		res, err := client.ListSavedTraceSearches(&config_unstable.ListSavedTraceSearchesParams{
			Context:     ctx,
			PageToken:   &nextToken,
			PageMaxSize: ptr.Int64(int64(pageMaxSize)),
			CreatedBy:   &opts.CreatedBy,
			Names:       opts.Names,
			Slugs:       opts.Slugs,
		})
		if err != nil {
			return pagination.Token(""), clienterror.Wrap(err)
		}

		for _, v := range res.Payload.SavedTraceSearches {
			if err := streamer(NewSavedTraceSearch(v)); err != nil {
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
			MaxPageSize: len(res.Payload.SavedTraceSearches),
		})
	}
}

func newSavedTraceSearchListCmd() *cobra.Command {
	var listOptions SavedTraceSearchListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all SavedTraceSearches and applies optional filters",
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

			streamer := output.NewWriteObjectStreamer[*SavedTraceSearch](writer)
			nextToken, err := ListSavedTraceSearches(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional savedTraceSearches. To view more, use the next page token or run the full command.",
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

const SavedTraceSearchScaffoldYAML = `api_version: unstable/config
kind: SavedTraceSearch
spec:
    # Required name of the SavedTraceSearch. May be modified after the SavedTraceSearch is created.
    name: <string>
    # Unique identifier of the SavedTraceSearch. If slug is not provided, one will be generated based of the name field. Cannot be modified after the SavedTraceSearch is created.
    slug: <string>
    comparison:
        criteria:
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
                        match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION>
                  duration:
                    # Minimum duration, in seconds, required for a span or trace to match.
                    min_secs: <integer>
                    # Maximum duration, in seconds, required for a span or trace to match.
                    max_secs: <integer>
                  error:
                    # The value the filter compares to the target trace or span field.
                    value: <true|false>
                  match_type: <INCLUDE|EXCLUDE>
                  operation:
                    # The value the filter compares to the target trace or span field.
                    value: <string>
                    match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION>
                  parent_operation:
                    # The value the filter compares to the target trace or span field.
                    value: <string>
                    match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION>
                  parent_service:
                    # The value the filter compares to the target trace or span field.
                    value: <string>
                    match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION>
                  service:
                    # The value the filter compares to the target trace or span field.
                    value: <string>
                    match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION>
                  span_count:
                    # Minimum number of spans that must match a SpanFilter (inclusive).
                    min: <integer>
                    # Maximum number of spans that must match a SpanFilter (inclusive).
                    max: <integer>
            trace:
                duration:
                    # Minimum duration, in seconds, required for a span or trace to match.
                    min_secs: <integer>
                    # Maximum duration, in seconds, required for a span or trace to match.
                    max_secs: <integer>
                error:
                    # The value the filter compares to the target trace or span field.
                    value: <true|false>
        time:
            between:
                # Start time of the search interval.
                min_time: <date-time>
                # End time of the search interval.
                max_time: <date-time>
            close_to:
                # Time around which the search will performed.
                time: <date-time>
            relative:
                # The duration, in seconds, from now to the beginning of the search interval.
                start_relative_offset_secs: <integer>
                # The duration, in seconds, from now to the end of the search interval.
                end_relative_offset_secs: <integer>
    search:
        criteria:
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
                        match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION>
                  duration:
                    # Minimum duration, in seconds, required for a span or trace to match.
                    min_secs: <integer>
                    # Maximum duration, in seconds, required for a span or trace to match.
                    max_secs: <integer>
                  error:
                    # The value the filter compares to the target trace or span field.
                    value: <true|false>
                  match_type: <INCLUDE|EXCLUDE>
                  operation:
                    # The value the filter compares to the target trace or span field.
                    value: <string>
                    match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION>
                  parent_operation:
                    # The value the filter compares to the target trace or span field.
                    value: <string>
                    match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION>
                  parent_service:
                    # The value the filter compares to the target trace or span field.
                    value: <string>
                    match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION>
                  service:
                    # The value the filter compares to the target trace or span field.
                    value: <string>
                    match: <EXACT|REGEX|EXACT_NEGATION|REGEX_NEGATION>
                  span_count:
                    # Minimum number of spans that must match a SpanFilter (inclusive).
                    min: <integer>
                    # Maximum number of spans that must match a SpanFilter (inclusive).
                    max: <integer>
            trace:
                duration:
                    # Minimum duration, in seconds, required for a span or trace to match.
                    min_secs: <integer>
                    # Maximum duration, in seconds, required for a span or trace to match.
                    max_secs: <integer>
                error:
                    # The value the filter compares to the target trace or span field.
                    value: <true|false>
        time:
            between:
                # Start time of the search interval.
                min_time: <date-time>
                # End time of the search interval.
                max_time: <date-time>
            close_to:
                # Time around which the search will performed.
                time: <date-time>
            relative:
                # The duration, in seconds, from now to the beginning of the search interval.
                start_relative_offset_secs: <integer>
                # The duration, in seconds, from now to the end of the search interval.
                end_relative_offset_secs: <integer>
`

func newSavedTraceSearchScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), SavedTraceSearchScaffoldYAML)
		},
	}
	return cmd
}

func NewSavedTraceSearchCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "saved-trace-searches",
		GroupID: groups.Config.ID,
		Short:   "All commands for SavedTraceSearches",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newSavedTraceSearchCreateCmd(),
		newSavedTraceSearchReadCmd(),
		newSavedTraceSearchUpdateCmd(),
		newSavedTraceSearchDeleteCmd(),
		newSavedTraceSearchListCmd(),
		newSavedTraceSearchScaffoldCmd(),
	)
	return root
}
