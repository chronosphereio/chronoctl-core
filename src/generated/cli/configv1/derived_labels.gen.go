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

func init() { types.MustRegisterObject(DerivedLabelTypeMeta, &DerivedLabel{}) }

var _ types.Object = &DerivedLabel{}

var DerivedLabelTypeMeta = types.TypeMeta{
	APIVersion: "v1/config",
	Kind:       "DerivedLabel",
}

type DerivedLabel struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.Configv1DerivedLabel `json:"spec"`
}

func NewDerivedLabel(spec *models.Configv1DerivedLabel) *DerivedLabel {
	return &DerivedLabel{
		TypeMeta: DerivedLabelTypeMeta,
		Spec:     spec,
	}
}

func (e *DerivedLabel) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *DerivedLabel) Identifier() string {
	return e.Spec.Slug
}

func CreateDerivedLabel(
	ctx context.Context,
	client config_v1.ClientService,
	entity *DerivedLabel,
	dryRun bool,
) (*DerivedLabel, error) {
	res, err := client.CreateDerivedLabel(&config_v1.CreateDerivedLabelParams{
		Context: ctx,
		Body: &models.Configv1CreateDerivedLabelRequest{
			DryRun:       dryRun,
			DerivedLabel: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewDerivedLabel(res.Payload.DerivedLabel), nil
}

func newDerivedLabelCreateCmd() *cobra.Command {
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
	short = "Creates a single DerivedLabel."

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

			var derivedLabel *DerivedLabel
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			derivedLabel, err = types.MustDecodeSingleObject[*DerivedLabel](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullDerivedLabel, err := CreateDerivedLabel(ctx, client, derivedLabel, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("DerivedLabel is valid and can be created")
				return nil
			}
			stderr.Printf("DerivedLabel with slug %q created successfully\n", fullDerivedLabel.Spec.Slug)

			if err := outputFlags.WriteObject(fullDerivedLabel, cmd.OutOrStdout()); err != nil {
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

func GetDerivedLabel(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) (*DerivedLabel, error) {
	res, err := client.ReadDerivedLabel(&config_v1.ReadDerivedLabelParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewDerivedLabel(res.GetPayload().DerivedLabel), nil
}

func newDerivedLabelReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single DerivedLabel by slug"
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
			entity, err := GetDerivedLabel(ctx, client, args[0])
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

func UpdateDerivedLabel(
	ctx context.Context,
	client config_v1.ClientService,
	entity *DerivedLabel,
	opts UpdateOptions,
) (*DerivedLabel, error) {
	res, err := client.UpdateDerivedLabel(&config_v1.UpdateDerivedLabelParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: config_v1.UpdateDerivedLabelBody{
			CreateIfMissing: opts.CreateIfMissing,
			DryRun:          opts.DryRun,
			DerivedLabel:    entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewDerivedLabel(res.Payload.DerivedLabel), nil
}

func newDerivedLabelUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing DerivedLabel.",
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

			derivedLabel, err := types.MustDecodeSingleObject[*DerivedLabel](file, permissiveParsing)
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

			fullDerivedLabel, err := UpdateDerivedLabel(ctx, client, derivedLabel, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("DerivedLabel is valid and can be updated")
				return nil
			}
			stderr.Printf("DerivedLabel with slug %q applied successfully\n", fullDerivedLabel.Spec.Slug)

			if err := outputFlags.WriteObject(fullDerivedLabel, cmd.OutOrStdout()); err != nil {
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
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the DerivedLabel if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteDerivedLabel(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) error {
	_, err := client.DeleteDerivedLabel(&config_v1.DeleteDerivedLabelParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newDerivedLabelDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single DerivedLabel by slug",
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

			res, err := client.DeleteDerivedLabel(&config_v1.DeleteDerivedLabelParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted DerivedLabel with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type DerivedLabelListOpts struct {
	// Limit represents that maximum number of items we wish to return.
	Limit int
	// PageToken is the pagination token we want to start our request at.
	PageToken string
	// PageMaxSize is the maximum page size to use when making List calls.
	PageMaxSize int
	Names       []string
	Slugs       []string
}

func (r *DerivedLabelListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any DerivedLabel with a matching name in the given list (and matches all other filters) is returned.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any DerivedLabel with a matching slug in the given list (and matches all other filters) is returned.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListDerivedLabels(
	ctx context.Context,
	client config_v1.ClientService,
	streamer output.Streamer[*DerivedLabel],
	opts DerivedLabelListOpts,
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
		res, err := client.ListDerivedLabels(&config_v1.ListDerivedLabelsParams{
			Context:     ctx,
			PageToken:   &nextToken,
			PageMaxSize: ptr.Int64(int64(pageMaxSize)),
			Names:       opts.Names,
			Slugs:       opts.Slugs,
		})
		if err != nil {
			return pagination.Token(""), clienterror.Wrap(err)
		}

		for _, v := range res.Payload.DerivedLabels {
			if err := streamer(NewDerivedLabel(v)); err != nil {
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
			MaxPageSize: len(res.Payload.DerivedLabels),
		})
	}
}

func newDerivedLabelListCmd() *cobra.Command {
	var listOptions DerivedLabelListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all DerivedLabels and applies optional filters",
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

			streamer := output.NewWriteObjectStreamer[*DerivedLabel](writer)
			nextToken, err := ListDerivedLabels(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional derivedLabels. To view more, use the next page token or run the full command.",
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

const DerivedLabelScaffoldYAML = `api_version: v1/config
kind: DerivedLabel
spec:
    # Required name of the DerivedLabel. May be modified after the DerivedLabel is created.
    name: <string>
    # Unique identifier of the DerivedLabel. If slug is not provided, one will be generated based of the name field. Cannot be modified after the DerivedLabel is created.
    slug: <string>
    # Name of the derived label. It needs to be unique across the system.
    label_name: <string>
    # Optional description of the derived label.
    description: <string>
    existing_label_policy: <KEEP|OVERRIDE>
    metric_label:
        constructed_label:
            value_definitions:
                - value: <string>
                  filters:
                    - # Name of the label to match.
                      name: <string>
                      # Glob value of the label to match.
                      value_glob: <string>
        mapping_label:
            name_mappings:
                - filters:
                    - # Name of the label to match.
                      name: <string>
                      # Glob value of the label to match.
                      value_glob: <string>
                  # The actual label ingested on the time series
                  source_label: <string>
                  # These value mappings apply to just the name mapping they belong to.
                  value_mappings:
                    - # Defines the source label values that should be mapped into the given target_value.
                      source_value_globs:
                        - <string>
                      # The value that source_value_globs are mapped into.
                      # For example, given this mapping:
                      # '''yaml
                      # value_mappings:
                      #  - source_value_globs:
                      #      - Cat
                      #      - CAT
                      #    target_value: cat
                      # '''
                      # This indicates that the target value 'cat' maps to the source label's values 'Cat' and 'CAT'.
                      target_value: <string>
            # These value mappings apply to the whole mapping label.
            # If there's no name_mappings, these value mappings apply to the label that exists on the metric.
            value_mappings:
                - # Defines the source label values that should be mapped into the given target_value.
                  source_value_globs:
                    - <string>
                  # The value that source_value_globs are mapped into.
                  # For example, given this mapping:
                  # '''yaml
                  # value_mappings:
                  #  - source_value_globs:
                  #      - Cat
                  #      - CAT
                  #    target_value: cat
                  # '''
                  # This indicates that the target value 'cat' maps to the source label's values 'Cat' and 'CAT'.
                  target_value: <string>
`

func newDerivedLabelScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), DerivedLabelScaffoldYAML)
		},
	}
	return cmd
}

func NewDerivedLabelCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "derived-labels",
		GroupID: groups.Config.ID,
		Short:   "All commands for DerivedLabels",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newDerivedLabelCreateCmd(),
		newDerivedLabelReadCmd(),
		newDerivedLabelUpdateCmd(),
		newDerivedLabelDeleteCmd(),
		newDerivedLabelListCmd(),
		newDerivedLabelScaffoldCmd(),
	)
	return root
}