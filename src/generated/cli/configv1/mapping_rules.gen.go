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

func init() { types.MustRegisterObject(MappingRuleTypeMeta, &MappingRule{}) }

var _ types.Object = &MappingRule{}

var MappingRuleTypeMeta = types.TypeMeta{
	APIVersion: "v1/config",
	Kind:       "MappingRule",
}

type MappingRule struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.Configv1MappingRule `json:"spec"`
}

func NewMappingRule(spec *models.Configv1MappingRule) *MappingRule {
	return &MappingRule{
		TypeMeta: MappingRuleTypeMeta,
		Spec:     spec,
	}
}

func (e *MappingRule) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *MappingRule) Identifier() string {
	return e.Spec.Slug
}

func CreateMappingRule(
	ctx context.Context,
	client config_v1.ClientService,
	entity *MappingRule,
	dryRun bool,
) (*MappingRule, error) {
	res, err := client.CreateMappingRule(&config_v1.CreateMappingRuleParams{
		Context: ctx,
		Body: &models.Configv1CreateMappingRuleRequest{
			DryRun:      dryRun,
			MappingRule: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewMappingRule(res.Payload.MappingRule), nil
}

func newMappingRuleCreateCmd() *cobra.Command {
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
	short = "Creates a single MappingRule."

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

			var mappingRule *MappingRule
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			mappingRule, err = types.MustDecodeSingleObject[*MappingRule](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullMappingRule, err := CreateMappingRule(ctx, client, mappingRule, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("MappingRule is valid and can be created")
				return nil
			}
			stderr.Printf("MappingRule with slug %q created successfully\n", fullMappingRule.Spec.Slug)

			if err := outputFlags.WriteObject(fullMappingRule, cmd.OutOrStdout()); err != nil {
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

func GetMappingRule(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) (*MappingRule, error) {
	res, err := client.ReadMappingRule(&config_v1.ReadMappingRuleParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewMappingRule(res.GetPayload().MappingRule), nil
}

func newMappingRuleReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single MappingRule by slug"
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
			entity, err := GetMappingRule(ctx, client, args[0])
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

func UpdateMappingRule(
	ctx context.Context,
	client config_v1.ClientService,
	entity *MappingRule,
	opts UpdateOptions,
) (*MappingRule, error) {
	res, err := client.UpdateMappingRule(&config_v1.UpdateMappingRuleParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: &models.ConfigV1UpdateMappingRuleBody{
			CreateIfMissing: opts.CreateIfMissing,
			DryRun:          opts.DryRun,
			MappingRule:     entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewMappingRule(res.Payload.MappingRule), nil
}

func newMappingRuleUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing MappingRule.",
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

			mappingRule, err := types.MustDecodeSingleObject[*MappingRule](file, permissiveParsing)
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

			fullMappingRule, err := UpdateMappingRule(ctx, client, mappingRule, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("MappingRule is valid and can be updated")
				return nil
			}
			stderr.Printf("MappingRule with slug %q applied successfully\n", fullMappingRule.Spec.Slug)

			if err := outputFlags.WriteObject(fullMappingRule, cmd.OutOrStdout()); err != nil {
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
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the MappingRule if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteMappingRule(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) error {
	_, err := client.DeleteMappingRule(&config_v1.DeleteMappingRuleParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newMappingRuleDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single MappingRule by slug",
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

			res, err := client.DeleteMappingRule(&config_v1.DeleteMappingRuleParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted MappingRule with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type MappingRuleListOpts struct {
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

func (r *MappingRuleListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyBucketSlugs []string
	flags.StringSliceVar(&r.BucketSlugs, "bucket-slugs", emptyBucketSlugs, "Filters results by bucket_slug, where any MappingRule with a matching bucket_slug in the given list (and matches all other filters) is returned.")
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any MappingRule with a matching name in the given list (and matches all other filters) is returned.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any MappingRule with a matching slug in the given list (and matches all other filters) is returned.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListMappingRules(
	ctx context.Context,
	client config_v1.ClientService,
	streamer output.Streamer[*MappingRule],
	opts MappingRuleListOpts,
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
		res, err := client.ListMappingRules(&config_v1.ListMappingRulesParams{
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

		for _, v := range res.Payload.MappingRules {
			if err := streamer(NewMappingRule(v)); err != nil {
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
			MaxPageSize: len(res.Payload.MappingRules),
		})
	}
}

func newMappingRuleListCmd() *cobra.Command {
	var listOptions MappingRuleListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all MappingRules and applies optional filters",
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

			streamer := output.NewWriteObjectStreamer[*MappingRule](writer)
			nextToken, err := ListMappingRules(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional mappingRules. To view more, use the next page token or run the full command.",
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

const MappingRuleScaffoldYAML = `api_version: v1/config
kind: MappingRule
spec:
    # Unique identifier of the MappingRule. If slug is not provided, one will be generated based of the name field. Cannot be modified after the MappingRule is created.
    slug: <string>
    # Required name of the MappingRule. May be modified after the MappingRule is created.
    name: <string>
    # Required slug of the bucket the MappingRule belongs to.
    bucket_slug: <string>
    # Required filters that determine to which metrics to apply the rule.
    filters:
        - # Name of the label to match.
          name: <string>
          # Glob value of the label to match.
          value_glob: <string>
    # Whether to drop the given set of metrics. If set, then the aggregation
    # policy can't be set.
    drop: <true|false>
    aggregation_policy:
        # Interval between aggregated data points, equivalent to the resolution
        # field in storage policy. If set, then the storage_policy field can't be
        # set.
        interval: <string>
        # Deprecated: This field is no longer supported.
        drop_timestamp: <true|false>
        aggregation: <LAST|MIN|MAX|MEAN|MEDIAN|COUNT|SUM|SUMSQ|STDEV|P10|P20|P30|P40|P50|P60|P70|P80|P90|P95|P99|P999|P9999|P25|P75|COUNT_SAMPLES|HISTOGRAM>
        storage_policy:
            # Required resolution of the aggregated metrics.
            resolution: <string>
            # Required retention of the aggregated metrics.
            retention: <string>
    #  - ENABLED: ENABLED rules are applied. Rules default to ENABLED.
    #  - PREVIEW: PREVIEW rules are not applied, but shaping impact stats
    # for them rule are recorded.
    mode: <ENABLED|PREVIEW>
`

func newMappingRuleScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), MappingRuleScaffoldYAML)
		},
	}
	return cmd
}

func NewMappingRuleCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "mapping-rules",
		GroupID: groups.Config.ID,
		Short:   "All commands for MappingRules",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newMappingRuleCreateCmd(),
		newMappingRuleReadCmd(),
		newMappingRuleUpdateCmd(),
		newMappingRuleDeleteCmd(),
		newMappingRuleListCmd(),
		newMappingRuleScaffoldCmd(),
	)
	return root
}
