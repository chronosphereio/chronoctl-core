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

func init() { types.MustRegisterObject(RecordingRuleTypeMeta, &RecordingRule{}) }

var _ types.Object = &RecordingRule{}

var RecordingRuleTypeMeta = types.TypeMeta{
	APIVersion: "v1/config",
	Kind:       "RecordingRule",
}

type RecordingRule struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.Configv1RecordingRule `json:"spec"`
}

func NewRecordingRule(spec *models.Configv1RecordingRule) *RecordingRule {
	return &RecordingRule{
		TypeMeta: RecordingRuleTypeMeta,
		Spec:     spec,
	}
}

func (e *RecordingRule) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *RecordingRule) Identifier() string {
	return e.Spec.Slug
}

func CreateRecordingRule(
	ctx context.Context,
	client config_v1.ClientService,
	entity *RecordingRule,
	dryRun bool,
) (*RecordingRule, error) {
	res, err := client.CreateRecordingRule(&config_v1.CreateRecordingRuleParams{
		Context: ctx,
		Body: &models.Configv1CreateRecordingRuleRequest{
			DryRun:        dryRun,
			RecordingRule: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewRecordingRule(res.Payload.RecordingRule), nil
}

func newRecordingRuleCreateCmd() *cobra.Command {
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
	short = "Creates a single RecordingRule."

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

			var recordingRule *RecordingRule
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			recordingRule, err = types.MustDecodeSingleObject[*RecordingRule](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullRecordingRule, err := CreateRecordingRule(ctx, client, recordingRule, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("RecordingRule is valid and can be created")
				return nil
			}
			stderr.Printf("RecordingRule with slug %q created successfully\n", fullRecordingRule.Spec.Slug)

			if err := outputFlags.WriteObject(fullRecordingRule, cmd.OutOrStdout()); err != nil {
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

func GetRecordingRule(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) (*RecordingRule, error) {
	res, err := client.ReadRecordingRule(&config_v1.ReadRecordingRuleParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewRecordingRule(res.GetPayload().RecordingRule), nil
}

func newRecordingRuleReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single RecordingRule by slug"
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
			entity, err := GetRecordingRule(ctx, client, args[0])
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

func UpdateRecordingRule(
	ctx context.Context,
	client config_v1.ClientService,
	entity *RecordingRule,
	opts UpdateOptions,
) (*RecordingRule, error) {
	res, err := client.UpdateRecordingRule(&config_v1.UpdateRecordingRuleParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: &models.ConfigV1UpdateRecordingRuleBody{
			CreateIfMissing: opts.CreateIfMissing,
			DryRun:          opts.DryRun,
			RecordingRule:   entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewRecordingRule(res.Payload.RecordingRule), nil
}

func newRecordingRuleUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing RecordingRule.",
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

			recordingRule, err := types.MustDecodeSingleObject[*RecordingRule](file, permissiveParsing)
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

			fullRecordingRule, err := UpdateRecordingRule(ctx, client, recordingRule, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("RecordingRule is valid and can be updated")
				return nil
			}
			stderr.Printf("RecordingRule with slug %q applied successfully\n", fullRecordingRule.Spec.Slug)

			if err := outputFlags.WriteObject(fullRecordingRule, cmd.OutOrStdout()); err != nil {
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
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the RecordingRule if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteRecordingRule(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) error {
	_, err := client.DeleteRecordingRule(&config_v1.DeleteRecordingRuleParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newRecordingRuleDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single RecordingRule by slug",
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

			res, err := client.DeleteRecordingRule(&config_v1.DeleteRecordingRuleParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted RecordingRule with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type RecordingRuleListOpts struct {
	// Limit represents that maximum number of items we wish to return.
	Limit int
	// PageToken is the pagination token we want to start our request at.
	PageToken string
	// PageMaxSize is the maximum page size to use when making List calls.
	PageMaxSize     int
	BucketSlugs     []string
	ExecutionGroups []string
	Names           []string
	Slugs           []string
}

func (r *RecordingRuleListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyBucketSlugs []string
	flags.StringSliceVar(&r.BucketSlugs, "bucket-slugs", emptyBucketSlugs, "The execution_groups filter cannot be used when a bucket_slug filter is provided.")
	var emptyExecutionGroups []string
	flags.StringSliceVar(&r.ExecutionGroups, "execution-groups", emptyExecutionGroups, "The bucket_slugs filter cannot be used when an execution_group filter is provided.")
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any RecordingRule with a matching name in the given list (and matches all other filters) is returned.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any RecordingRule with a matching slug in the given list (and matches all other filters) is returned.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListRecordingRules(
	ctx context.Context,
	client config_v1.ClientService,
	streamer output.Streamer[*RecordingRule],
	opts RecordingRuleListOpts,
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
		res, err := client.ListRecordingRules(&config_v1.ListRecordingRulesParams{
			Context:         ctx,
			PageToken:       &nextToken,
			PageMaxSize:     ptr.Int64(int64(pageMaxSize)),
			BucketSlugs:     opts.BucketSlugs,
			ExecutionGroups: opts.ExecutionGroups,
			Names:           opts.Names,
			Slugs:           opts.Slugs,
		})
		if err != nil {
			return pagination.Token(""), clienterror.Wrap(err)
		}

		for _, v := range res.Payload.RecordingRules {
			if err := streamer(NewRecordingRule(v)); err != nil {
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
			MaxPageSize: len(res.Payload.RecordingRules),
		})
	}
}

func newRecordingRuleListCmd() *cobra.Command {
	var listOptions RecordingRuleListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all RecordingRules and applies optional filters",
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

			streamer := output.NewWriteObjectStreamer[*RecordingRule](writer)
			nextToken, err := ListRecordingRules(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional recordingRules. To view more, use the next page token or run the full command.",
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

const RecordingRuleScaffoldYAML = `api_version: v1/config
kind: RecordingRule
spec:
    # Unique identifier of the RecordingRule. If a 'slug' isn't provided, one will be generated based of the 'name' field. You can't modify this field after the RecordingRule is created.
    slug: <string>
    # Required. Name of the RecordingRule. You can modify this value after the RecordingRule is created.
    name: <string>
    # Optional slug of the bucket the RecordingRule belongs to.
    bucket_slug: <string>
    # Optional interval for evaluating the recording rule.
    interval_secs: <integer>
    # The name of the time series to use for output, which must be a valid
    # metric name.
    metric_name: <string>
    # The PromQL expression to evaluate at the time of each evaluation cycle.
    # The result is recorded as a new time series with its metric name
    # defined by the metric_name (or name) field.
    prometheus_expr: <string>
    # Optional execution_group in which this rule is to be evaluated.
    # At least one of bucket_slug and execution_group must be set. If both are set, then they are expected to match.
    execution_group: <string>
    label_policy:
        # Labels to add or overwrite before storing the result.
        add:
            key_1: <string>
`

func newRecordingRuleScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), RecordingRuleScaffoldYAML)
		},
	}
	return cmd
}

func NewRecordingRuleCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "recording-rules",
		GroupID: groups.Config.ID,
		Short:   "All commands for RecordingRules",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newRecordingRuleCreateCmd(),
		newRecordingRuleReadCmd(),
		newRecordingRuleUpdateCmd(),
		newRecordingRuleDeleteCmd(),
		newRecordingRuleListCmd(),
		newRecordingRuleScaffoldCmd(),
	)
	return root
}
