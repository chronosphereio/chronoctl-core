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

func init() { types.MustRegisterObject(ServiceTypeMeta, &Service{}) }

var _ types.Object = &Service{}

var ServiceTypeMeta = types.TypeMeta{
	APIVersion: "unstable/config",
	Kind:       "Service",
}

type Service struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.ConfigunstableService `json:"spec"`
}

func NewService(spec *models.ConfigunstableService) *Service {
	return &Service{
		TypeMeta: ServiceTypeMeta,
		Spec:     spec,
	}
}

func (e *Service) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *Service) Identifier() string {
	return e.Spec.Slug
}

func CreateService(
	ctx context.Context,
	client config_unstable.ClientService,
	entity *Service,
	dryRun bool,
) (*Service, error) {
	res, err := client.CreateService(&config_unstable.CreateServiceParams{
		Context: ctx,
		Body: &models.ConfigunstableCreateServiceRequest{
			DryRun:  dryRun,
			Service: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewService(res.Payload.Service), nil
}

func newServiceCreateCmd() *cobra.Command {
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
	short = "Creates a single Service."

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

			var service *Service
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			service, err = types.MustDecodeSingleObject[*Service](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullService, err := CreateService(ctx, client, service, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("Service is valid and can be created")
				return nil
			}
			stderr.Printf("Service with slug %q created successfully\n", fullService.Spec.Slug)

			if err := outputFlags.WriteObject(fullService, cmd.OutOrStdout()); err != nil {
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

func GetService(
	ctx context.Context,
	client config_unstable.ClientService,
	slug string,
) (*Service, error) {
	res, err := client.ReadService(&config_unstable.ReadServiceParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewService(res.GetPayload().Service), nil
}

func newServiceReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single Service by slug"
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
			entity, err := GetService(ctx, client, args[0])
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

func UpdateService(
	ctx context.Context,
	client config_unstable.ClientService,
	entity *Service,
	opts UpdateOptions,
) (*Service, error) {
	res, err := client.UpdateService(&config_unstable.UpdateServiceParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: config_unstable.UpdateServiceBody{
			CreateIfMissing: opts.CreateIfMissing,
			DryRun:          opts.DryRun,
			Service:         entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewService(res.Payload.Service), nil
}

func newServiceUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing Service.",
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

			service, err := types.MustDecodeSingleObject[*Service](file, permissiveParsing)
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

			fullService, err := UpdateService(ctx, client, service, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("Service is valid and can be updated")
				return nil
			}
			stderr.Printf("Service with slug %q applied successfully\n", fullService.Spec.Slug)

			if err := outputFlags.WriteObject(fullService, cmd.OutOrStdout()); err != nil {
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
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the Service if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteService(
	ctx context.Context,
	client config_unstable.ClientService,
	slug string,
) error {
	_, err := client.DeleteService(&config_unstable.DeleteServiceParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newServiceDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single Service by slug",
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

			res, err := client.DeleteService(&config_unstable.DeleteServiceParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted Service with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type ServiceListOpts struct {
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

func (r *ServiceListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any Service with a matching name in the given list (and matches all other filters) is returned.")
	var emptyNotificationPolicySlugs []string
	flags.StringSliceVar(&r.NotificationPolicySlugs, "notification-policy-slugs", emptyNotificationPolicySlugs, "Get services that directly reference notifications policies by the referenced policy slugs.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any Service with a matching slug in the given list (and matches all other filters) is returned.")
	var emptyTeamSlugs []string
	flags.StringSliceVar(&r.TeamSlugs, "team-slugs", emptyTeamSlugs, "Filters results by team_slug, where any Service with a matching team_slug in the given list (and matches all other filters) is returned.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListServices(
	ctx context.Context,
	client config_unstable.ClientService,
	streamer output.Streamer[*Service],
	opts ServiceListOpts,
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
		res, err := client.ListServices(&config_unstable.ListServicesParams{
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

		for _, v := range res.Payload.Services {
			if err := streamer(NewService(v)); err != nil {
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
			MaxPageSize: len(res.Payload.Services),
		})
	}
}

func newServiceListCmd() *cobra.Command {
	var listOptions ServiceListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all Services and applies optional filters",
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

			streamer := output.NewWriteObjectStreamer[*Service](writer)
			nextToken, err := ListServices(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional services. To view more, use the next page token or run the full command.",
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

const ServiceScaffoldYAML = `api_version: unstable/config
kind: Service
spec:
    # Unique identifier of the Service. If slug is not provided, one will be generated based of the name field. Cannot be modified after the Service is created.
    slug: <string>
    # Required name of the Service. May be modified after the Service is created.
    name: <string>
    # Required telemetry name of the service.
    derived_name: <string>
    # Required slug of the team the service collection belongs to.
    team_slug: <string>
    # Optional description of the service collection.
    description: <string>
    # Slug of the notification policy used by default for monitors in this service collection.
    # This is optional if the collection does not contain monitors or all of its monitors explicitly reference a policy.
    # This does not override the policy used when a monitor explicitly references a policy.
    notification_policy_slug: <string>
`

func newServiceScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), ServiceScaffoldYAML)
		},
	}
	return cmd
}

func NewServiceCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "services",
		GroupID: groups.Config.ID,
		Short:   "All commands for Services",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newServiceCreateCmd(),
		newServiceReadCmd(),
		newServiceUpdateCmd(),
		newServiceDeleteCmd(),
		newServiceListCmd(),
		newServiceScaffoldCmd(),
	)
	return root
}