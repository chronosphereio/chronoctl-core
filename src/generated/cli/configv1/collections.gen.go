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

func init() { types.MustRegisterObject(CollectionTypeMeta, &Collection{}) }

var _ types.Object = &Collection{}

var CollectionTypeMeta = types.TypeMeta{
	APIVersion: "v1/config",
	Kind:       "Collection",
}

type Collection struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.Configv1Collection `json:"spec"`
}

func NewCollection(spec *models.Configv1Collection) *Collection {
	return &Collection{
		TypeMeta: CollectionTypeMeta,
		Spec:     spec,
	}
}

func (e *Collection) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *Collection) Identifier() string {
	return e.Spec.Slug
}

func CreateCollection(
	ctx context.Context,
	client config_v1.ClientService,
	entity *Collection,
	dryRun bool,
) (*Collection, error) {
	res, err := client.CreateCollection(&config_v1.CreateCollectionParams{
		Context: ctx,
		Body: &models.Configv1CreateCollectionRequest{
			DryRun:     dryRun,
			Collection: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewCollection(res.Payload.Collection), nil
}

func newCollectionCreateCmd() *cobra.Command {
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
	short = "Creates a single Collection."

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

			var collection *Collection
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			collection, err = types.MustDecodeSingleObject[*Collection](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullCollection, err := CreateCollection(ctx, client, collection, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("Collection is valid and can be created")
				return nil
			}
			stderr.Printf("Collection with slug %q created successfully\n", fullCollection.Spec.Slug)

			if err := outputFlags.WriteObject(fullCollection, cmd.OutOrStdout()); err != nil {
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

func GetCollection(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) (*Collection, error) {
	res, err := client.ReadCollection(&config_v1.ReadCollectionParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewCollection(res.GetPayload().Collection), nil
}

func newCollectionReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single Collection by slug"
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
			entity, err := GetCollection(ctx, client, args[0])
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

func UpdateCollection(
	ctx context.Context,
	client config_v1.ClientService,
	entity *Collection,
	opts UpdateOptions,
) (*Collection, error) {
	res, err := client.UpdateCollection(&config_v1.UpdateCollectionParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: &models.ConfigV1UpdateCollectionBody{
			CreateIfMissing: opts.CreateIfMissing,
			DryRun:          opts.DryRun,
			Collection:      entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewCollection(res.Payload.Collection), nil
}

func newCollectionUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing Collection.",
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

			collection, err := types.MustDecodeSingleObject[*Collection](file, permissiveParsing)
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

			fullCollection, err := UpdateCollection(ctx, client, collection, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("Collection is valid and can be updated")
				return nil
			}
			stderr.Printf("Collection with slug %q applied successfully\n", fullCollection.Spec.Slug)

			if err := outputFlags.WriteObject(fullCollection, cmd.OutOrStdout()); err != nil {
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
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the Collection if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteCollection(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) error {
	_, err := client.DeleteCollection(&config_v1.DeleteCollectionParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newCollectionDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single Collection by slug",
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

			res, err := client.DeleteCollection(&config_v1.DeleteCollectionParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted Collection with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type CollectionListOpts struct {
	// Limit represents that maximum number of items we wish to return.
	Limit int
	// PageToken is the pagination token we want to start our request at.
	PageToken string
	// PageMaxSize is the maximum page size to use when making List calls.
	PageMaxSize             int
	Names                   []string
	NotificationPolicySlugs []string
	Slugs                   []string
	TeamSlugs               []string
}

func (r *CollectionListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any Collection with a matching name in the given list (and matches all other filters) is returned.")
	var emptyNotificationPolicySlugs []string
	flags.StringSliceVar(&r.NotificationPolicySlugs, "notification-policy-slugs", emptyNotificationPolicySlugs, "Get collections that directly reference notifications policies by the referenced policy slugs.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any Collection with a matching slug in the given list (and matches all other filters) is returned.")
	var emptyTeamSlugs []string
	flags.StringSliceVar(&r.TeamSlugs, "team-slugs", emptyTeamSlugs, "Filters results by team_slug, where any Collection with a matching team_slug in the given list (and matches all other filters) is returned.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListCollections(
	ctx context.Context,
	client config_v1.ClientService,
	streamer output.Streamer[*Collection],
	opts CollectionListOpts,
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
		res, err := client.ListCollections(&config_v1.ListCollectionsParams{
			Context:                 ctx,
			PageToken:               &nextToken,
			PageMaxSize:             ptr.Int64(int64(pageMaxSize)),
			Names:                   opts.Names,
			NotificationPolicySlugs: opts.NotificationPolicySlugs,
			Slugs:                   opts.Slugs,
			TeamSlugs:               opts.TeamSlugs,
		})
		if err != nil {
			return pagination.Token(""), clienterror.Wrap(err)
		}

		for _, v := range res.Payload.Collections {
			if err := streamer(NewCollection(v)); err != nil {
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
			MaxPageSize: len(res.Payload.Collections),
		})
	}
}

func newCollectionListCmd() *cobra.Command {
	var listOptions CollectionListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all Collections and applies optional filters",
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

			streamer := output.NewWriteObjectStreamer[*Collection](writer)
			nextToken, err := ListCollections(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional collections. To view more, use the next page token or run the full command.",
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

const CollectionScaffoldYAML = `api_version: v1/config
kind: Collection
spec:
    # Unique identifier of the Collection. If slug is not provided, one will be generated based of the name field. Cannot be modified after the Collection is created.
    slug: <string>
    # Required name of the Collection. May be modified after the Collection is created.
    name: <string>
    # Required slug of the team the collection belongs to.
    team_slug: <string>
    # Optional description of the collection.
    description: <string>
    # Slug of the notification policy used by default for monitors in this collection.
    # This is optional if the collection does not contain monitors or all of its monitors explicitly reference a policy.
    # This does not override the policy used when a monitor explicitly references a policy.
    notification_policy_slug: <string>
`

func newCollectionScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), CollectionScaffoldYAML)
		},
	}
	return cmd
}

func NewCollectionCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "collections",
		GroupID: groups.Config.ID,
		Short:   "All commands for Collections",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newCollectionCreateCmd(),
		newCollectionReadCmd(),
		newCollectionUpdateCmd(),
		newCollectionDeleteCmd(),
		newCollectionListCmd(),
		newCollectionScaffoldCmd(),
	)
	return root
}
