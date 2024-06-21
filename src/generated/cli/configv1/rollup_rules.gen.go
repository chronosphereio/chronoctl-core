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

func init() { types.MustRegisterObject(RollupRuleTypeMeta, &RollupRule{}) }

var _ types.Object = &RollupRule{}

var RollupRuleTypeMeta = types.TypeMeta{
	APIVersion: "v1/config",
	Kind:       "RollupRule",
}

type RollupRule struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.Configv1RollupRule `json:"spec"`
}

func NewRollupRule(spec *models.Configv1RollupRule) *RollupRule {
	return &RollupRule{
		TypeMeta: RollupRuleTypeMeta,
		Spec:     spec,
	}
}

func (e *RollupRule) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *RollupRule) Identifier() string {
	return e.Spec.Slug
}

func CreateRollupRule(
	ctx context.Context,
	client config_v1.ClientService,
	entity *RollupRule,
	dryRun bool,
) (*RollupRule, error) {
	res, err := client.CreateRollupRule(&config_v1.CreateRollupRuleParams{
		Context: ctx,
		Body: &models.Configv1CreateRollupRuleRequest{
			DryRun:     dryRun,
			RollupRule: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewRollupRule(res.Payload.RollupRule), nil
}

func newRollupRuleCreateCmd() *cobra.Command {
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
	short = "Creates a single RollupRule."

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

			var rollupRule *RollupRule
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			rollupRule, err = types.MustDecodeSingleObject[*RollupRule](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullRollupRule, err := CreateRollupRule(ctx, client, rollupRule, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("RollupRule is valid and can be created")
				return nil
			}
			stderr.Printf("RollupRule with slug %q created successfully\n", fullRollupRule.Spec.Slug)

			if err := outputFlags.WriteObject(fullRollupRule, cmd.OutOrStdout()); err != nil {
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

func GetRollupRule(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) (*RollupRule, error) {
	res, err := client.ReadRollupRule(&config_v1.ReadRollupRuleParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewRollupRule(res.GetPayload().RollupRule), nil
}

func newRollupRuleReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single RollupRule by slug"
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
			entity, err := GetRollupRule(ctx, client, args[0])
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

func UpdateRollupRule(
	ctx context.Context,
	client config_v1.ClientService,
	entity *RollupRule,
	opts UpdateOptions,
) (*RollupRule, error) {
	res, err := client.UpdateRollupRule(&config_v1.UpdateRollupRuleParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: &models.ConfigV1UpdateRollupRuleBody{
			CreateIfMissing: opts.CreateIfMissing,
			DryRun:          opts.DryRun,
			RollupRule:      entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewRollupRule(res.Payload.RollupRule), nil
}

func newRollupRuleUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing RollupRule.",
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

			rollupRule, err := types.MustDecodeSingleObject[*RollupRule](file, permissiveParsing)
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

			fullRollupRule, err := UpdateRollupRule(ctx, client, rollupRule, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("RollupRule is valid and can be updated")
				return nil
			}
			stderr.Printf("RollupRule with slug %q applied successfully\n", fullRollupRule.Spec.Slug)

			if err := outputFlags.WriteObject(fullRollupRule, cmd.OutOrStdout()); err != nil {
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
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the RollupRule if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteRollupRule(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) error {
	_, err := client.DeleteRollupRule(&config_v1.DeleteRollupRuleParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newRollupRuleDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single RollupRule by slug",
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

			res, err := client.DeleteRollupRule(&config_v1.DeleteRollupRuleParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted RollupRule with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type RollupRuleListOpts struct {
	// Limit represents that maximum number of items we wish to return.
	Limit int
	// PageToken is the pagination token we want to start our request at.
	PageToken string
	// PageMaxSize is the maximum page size to use when making List calls.
	PageMaxSize int
	BucketSlugs []string
	Names       []string
	Slugs       []string
}

func (r *RollupRuleListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyBucketSlugs []string
	flags.StringSliceVar(&r.BucketSlugs, "bucket-slugs", emptyBucketSlugs, "Filters results by bucket_slug, where any RollupRule with a matching bucket_slug in the given list (and matches all other filters) is returned.")
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any RollupRule with a matching name in the given list (and matches all other filters) is returned.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any RollupRule with a matching slug in the given list (and matches all other filters) is returned.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListRollupRules(
	ctx context.Context,
	client config_v1.ClientService,
	streamer output.Streamer[*RollupRule],
	opts RollupRuleListOpts,
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
		res, err := client.ListRollupRules(&config_v1.ListRollupRulesParams{
			Context:     ctx,
			PageToken:   &nextToken,
			PageMaxSize: ptr.Int64(int64(pageMaxSize)),
			BucketSlugs: opts.BucketSlugs,
			Names:       opts.Names,
			Slugs:       opts.Slugs,
		})
		if err != nil {
			return pagination.Token(""), clienterror.Wrap(err)
		}

		for _, v := range res.Payload.RollupRules {
			if err := streamer(NewRollupRule(v)); err != nil {
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
			MaxPageSize: len(res.Payload.RollupRules),
		})
	}
}

func newRollupRuleListCmd() *cobra.Command {
	var listOptions RollupRuleListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all RollupRules and applies optional filters",
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

			streamer := output.NewWriteObjectStreamer[*RollupRule](writer)
			nextToken, err := ListRollupRules(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional rollupRules. To view more, use the next page token or run the full command.",
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

const RollupRuleScaffoldYAML = `api_version: v1/config
kind: RollupRule
spec:
    # Unique identifier of the RollupRule. If slug is not provided, one will be generated based of the name field. Cannot be modified after the RollupRule is created.
    slug: <string>
    # Required name of the RollupRule. May be modified after the RollupRule is created.
    name: <string>
    # Required slug of the bucket the RollupRule belongs to.
    bucket_slug: <string>
    # Filters that determine to which metrics to apply the rule.
    filters:
        - # Name of the label to match.
          name: <string>
          # Glob value of the label to match.
          value_glob: <string>
    # Name of the new metric created as a result of the rollup.
    metric_name: <string>
    # Interval between aggregated data points, equivalent to the resolution
    # field in storage policy. If set, then the storage_policy field can't be
    # set.
    interval: <string>
    # Enables expansive label matching behavior for the provided filters and
    # label_policy.keep or graphite_label_policy.replace (if set). By default
    # (expansive_match=false), a series matches and aggregates only if each label
    # defined by filters and label_policy.keep or graphite_label_policy.replace
    # (respectively) exist in said series. Setting expansive_match=true removes
    # this restriction.
    expansive_match: <true|false>
    # Defines whether to add a '__rollup_type__' label in the new metric.
    add_metric_type_label: <true|false>
    # Defines whether to automatically generate drop rules for this rollup.
    drop_raw: <true|false>
    aggregation: <LAST|MIN|MAX|MEAN|MEDIAN|COUNT|SUM|SUMSQ|STDEV|P10|P20|P30|P40|P50|P60|P70|P80|P90|P95|P99|P999|P9999|P25|P75|COUNT_SAMPLES|HISTOGRAM>
    graphite_label_policy:
        # Required list of labels to replace. Useful for discarding
        # high-cardinality values while still preserving the original positions of
        # the Graphite metric.
        replace:
            - # Required name of the label whose value should be replaced. Only
              # '__gX__' labels are allowed (aka positional Graphite labels).
              name: <string>
              # Required new value of the replaced label.
              new_value: <string>
    # TODO: consolidate w/ RecordingRule.LabelPolicy once both of these
    #  entities implement the same label semantics.
    label_policy:
        # Labels that should be retained in the output metric. If set, then the
        # discard field must be empty.
        keep:
            - <string>
        # Labels that should be discarded in the output metric. If set, then the
        # keep field must be empty.
        discard:
            - <string>
    # Must keep this around for backwards compatibility because terraform will
    # still send this key w/ a null value.
    label_replace: null
    #  - CUMULATIVE_COUNTER: Alias of COUNTER.
    #  - DELTA_COUNTER: Alias of DELTA.
    metric_type: <COUNTER|GAUGE|DELTA|DISTRIBUTION|CUMULATIVE_EXPONENTIAL_HISTOGRAM|MEASUREMENT|CUMULATIVE_COUNTER|DELTA_COUNTER|DELTA_EXPONENTIAL_HISTOGRAM>
    mode: <ENABLED|PREVIEW>
    storage_policy:
        # Required resolution of the aggregated metrics.
        resolution: <string>
        # Required retention of the aggregated metrics.
        retention: <string>
`

func newRollupRuleScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), RollupRuleScaffoldYAML)
		},
	}
	return cmd
}

func NewRollupRuleCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "rollup-rules",
		GroupID: groups.Config.ID,
		Short:   "All commands for RollupRules",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newRollupRuleCreateCmd(),
		newRollupRuleReadCmd(),
		newRollupRuleUpdateCmd(),
		newRollupRuleDeleteCmd(),
		newRollupRuleListCmd(),
		newRollupRuleScaffoldCmd(),
	)
	return root
}
