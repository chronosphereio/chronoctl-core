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

func init() { types.MustRegisterObject(LogScaleAlertTypeMeta, &LogScaleAlert{}) }

var _ types.Object = &LogScaleAlert{}

var LogScaleAlertTypeMeta = types.TypeMeta{
	APIVersion: "v1/config",
	Kind:       "LogScaleAlert",
}

type LogScaleAlert struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.Configv1LogScaleAlert `json:"spec"`
}

func NewLogScaleAlert(spec *models.Configv1LogScaleAlert) *LogScaleAlert {
	return &LogScaleAlert{
		TypeMeta: LogScaleAlertTypeMeta,
		Spec:     spec,
	}
}

func (e *LogScaleAlert) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *LogScaleAlert) Identifier() string {
	return e.Spec.Slug
}

func CreateLogScaleAlert(
	ctx context.Context,
	client config_v1.ClientService,
	entity *LogScaleAlert,
	dryRun bool,
) (*LogScaleAlert, error) {
	res, err := client.CreateLogScaleAlert(&config_v1.CreateLogScaleAlertParams{
		Context: ctx,
		Body: &models.Configv1CreateLogScaleAlertRequest{
			DryRun:        dryRun,
			LogScaleAlert: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewLogScaleAlert(res.Payload.LogScaleAlert), nil
}

func newLogScaleAlertCreateCmd() *cobra.Command {
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
	short = "Creates a single LogScaleAlert."

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

			var logScaleAlert *LogScaleAlert
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			logScaleAlert, err = types.MustDecodeSingleObject[*LogScaleAlert](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullLogScaleAlert, err := CreateLogScaleAlert(ctx, client, logScaleAlert, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("LogScaleAlert is valid and can be created")
				return nil
			}
			stderr.Printf("LogScaleAlert with slug %q created successfully\n", fullLogScaleAlert.Spec.Slug)

			if err := outputFlags.WriteObject(fullLogScaleAlert, cmd.OutOrStdout()); err != nil {
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

func GetLogScaleAlert(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) (*LogScaleAlert, error) {
	res, err := client.ReadLogScaleAlert(&config_v1.ReadLogScaleAlertParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewLogScaleAlert(res.GetPayload().LogScaleAlert), nil
}

func newLogScaleAlertReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single LogScaleAlert by slug"
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
			entity, err := GetLogScaleAlert(ctx, client, args[0])
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

func UpdateLogScaleAlert(
	ctx context.Context,
	client config_v1.ClientService,
	entity *LogScaleAlert,
	opts UpdateOptions,
) (*LogScaleAlert, error) {
	res, err := client.UpdateLogScaleAlert(&config_v1.UpdateLogScaleAlertParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: &models.ConfigV1UpdateLogScaleAlertBody{
			CreateIfMissing: opts.CreateIfMissing,
			DryRun:          opts.DryRun,
			LogScaleAlert:   entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewLogScaleAlert(res.Payload.LogScaleAlert), nil
}

func newLogScaleAlertUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing LogScaleAlert.",
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

			logScaleAlert, err := types.MustDecodeSingleObject[*LogScaleAlert](file, permissiveParsing)
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

			fullLogScaleAlert, err := UpdateLogScaleAlert(ctx, client, logScaleAlert, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("LogScaleAlert is valid and can be updated")
				return nil
			}
			stderr.Printf("LogScaleAlert with slug %q applied successfully\n", fullLogScaleAlert.Spec.Slug)

			if err := outputFlags.WriteObject(fullLogScaleAlert, cmd.OutOrStdout()); err != nil {
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
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the LogScaleAlert if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteLogScaleAlert(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) error {
	_, err := client.DeleteLogScaleAlert(&config_v1.DeleteLogScaleAlertParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newLogScaleAlertDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single LogScaleAlert by slug",
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

			res, err := client.DeleteLogScaleAlert(&config_v1.DeleteLogScaleAlertParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted LogScaleAlert with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type LogScaleAlertListOpts struct {
	// Limit represents that maximum number of items we wish to return.
	Limit int
	// PageToken is the pagination token we want to start our request at.
	PageToken string
	// PageMaxSize is the maximum page size to use when making List calls.
	PageMaxSize int
	Names       []string
	Slugs       []string
}

func (r *LogScaleAlertListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any LogScaleAlert with a matching name in the given list (and matches all other filters) is returned.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any LogScaleAlert with a matching slug in the given list (and matches all other filters) is returned.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListLogScaleAlerts(
	ctx context.Context,
	client config_v1.ClientService,
	streamer output.Streamer[*LogScaleAlert],
	opts LogScaleAlertListOpts,
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
		res, err := client.ListLogScaleAlerts(&config_v1.ListLogScaleAlertsParams{
			Context:     ctx,
			PageToken:   &nextToken,
			PageMaxSize: ptr.Int64(int64(pageMaxSize)),
			Names:       opts.Names,
			Slugs:       opts.Slugs,
		})
		if err != nil {
			return pagination.Token(""), clienterror.Wrap(err)
		}

		for _, v := range res.Payload.LogScaleAlerts {
			if err := streamer(NewLogScaleAlert(v)); err != nil {
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
			MaxPageSize: len(res.Payload.LogScaleAlerts),
		})
	}
}

func newLogScaleAlertListCmd() *cobra.Command {
	var listOptions LogScaleAlertListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all LogScaleAlerts and applies optional filters",
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

			streamer := output.NewWriteObjectStreamer[*LogScaleAlert](writer)
			nextToken, err := ListLogScaleAlerts(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional logScaleAlerts. To view more, use the next page token or run the full command.",
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

const LogScaleAlertScaffoldYAML = `api_version: v1/config
kind: LogScaleAlert
spec:
    # Unique identifier of the LogScaleAlert. If slug is not provided, one will be generated based of the name field. Cannot be modified after the LogScaleAlert is created.
    slug: <string>
    # Name of LogScale repository the alerts belongs to. Required.
    repository: <string>
    # Name of the alert.
    name: <string>
    # Description of the alert.
    description: <string>
    # Flag indicating whether the alert is disabled.
    disabled: <true|false>
    # LogScale query to execute.
    log_scale_query: <string>
    # Lookback window used for an alert's evaluation.
    # If this is set to 86400 seconds (24 hours), only the events from the last 24 hours will be considered when the alert query is run.
    time_window_secs: <integer>
    # Throttle time in seconds. The alert is triggered at most once per throttle period.
    throttle_secs: <integer>
    # Field to throttle on. Optional.
    throttle_field: <string>
    # Slugs of LogScale actions that will receive the alerts. When the value is empty
    # this alert won't trigger. Optional.
    log_scale_action_slugs:
        - <string>
    # Tags attached to the alert.
    tags:
        - <string>
    # Email of the user that the alert runs on behalf of. Required.
    run_as_user: <string>
    alert_type: <STANDARD|FILTER>
`

func newLogScaleAlertScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), LogScaleAlertScaffoldYAML)
		},
	}
	return cmd
}

func NewLogScaleAlertCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "log-scale-alerts",
		GroupID: groups.Config.ID,
		Short:   "All commands for LogScaleAlerts",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newLogScaleAlertCreateCmd(),
		newLogScaleAlertReadCmd(),
		newLogScaleAlertUpdateCmd(),
		newLogScaleAlertDeleteCmd(),
		newLogScaleAlertListCmd(),
		newLogScaleAlertScaffoldCmd(),
	)
	return root
}