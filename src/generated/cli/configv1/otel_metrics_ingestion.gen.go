// Code generated by chronogen; DO NOT EDIT
package configv1

import (
	"context"
	"fmt"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/clienterror"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/dry"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/file"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/groups"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/output"
	config_v1 "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
	"github.com/chronosphereio/chronoctl-core/src/types"
	"github.com/spf13/cobra"
)

func init() { types.MustRegisterObject(OtelMetricsIngestionTypeMeta, &OtelMetricsIngestion{}) }

var _ types.Object = &OtelMetricsIngestion{}

var OtelMetricsIngestionTypeMeta = types.TypeMeta{
	APIVersion: "v1/config",
	Kind:       "OtelMetricsIngestion",
}

type OtelMetricsIngestion struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.Configv1OtelMetricsIngestion `json:"spec"`
}

func NewOtelMetricsIngestion(spec *models.Configv1OtelMetricsIngestion) *OtelMetricsIngestion {
	return &OtelMetricsIngestion{
		TypeMeta: OtelMetricsIngestionTypeMeta,
		Spec:     spec,
	}
}

func (e *OtelMetricsIngestion) Description() string {
	return types.TypeDescription(e)
}

func (e *OtelMetricsIngestion) Identifier() string {
	return "OtelMetricsIngestion"
}

func CreateOtelMetricsIngestion(
	ctx context.Context,
	client config_v1.ClientService,
	entity *OtelMetricsIngestion,
	dryRun bool,
) (*OtelMetricsIngestion, error) {
	res, err := client.CreateOtelMetricsIngestion(&config_v1.CreateOtelMetricsIngestionParams{
		Context: ctx,
		Body: &models.Configv1CreateOtelMetricsIngestionRequest{
			DryRun:               dryRun,
			OtelMetricsIngestion: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewOtelMetricsIngestion(res.Payload.OtelMetricsIngestion), nil
}

func newOtelMetricsIngestionCreateCmd() *cobra.Command {
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
	short = "Creates a single OtelMetricsIngestion."

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

			var otelMetricsIngestion *OtelMetricsIngestion
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			otelMetricsIngestion, err = types.MustDecodeSingleObject[*OtelMetricsIngestion](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullOtelMetricsIngestion, err := CreateOtelMetricsIngestion(ctx, client, otelMetricsIngestion, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("OtelMetricsIngestion is valid and can be created")
				return nil
			}
			stderr.Printf("OtelMetricsIngestion created successfully\n")

			if err := outputFlags.WriteObject(fullOtelMetricsIngestion, cmd.OutOrStdout()); err != nil {
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

func GetOtelMetricsIngestion(
	ctx context.Context,
	client config_v1.ClientService,
) (*OtelMetricsIngestion, error) {
	res, err := client.ReadOtelMetricsIngestion(&config_v1.ReadOtelMetricsIngestionParams{
		Context: ctx,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewOtelMetricsIngestion(res.GetPayload().OtelMetricsIngestion), nil
}

func newOtelMetricsIngestionReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a OtelMetricsIngestion singleton"
	use = "read"

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
			entity, err := GetOtelMetricsIngestion(ctx, client)
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

func UpdateOtelMetricsIngestion(
	ctx context.Context,
	client config_v1.ClientService,
	entity *OtelMetricsIngestion,
	opts UpdateOptions,
) (*OtelMetricsIngestion, error) {
	res, err := client.UpdateOtelMetricsIngestion(&config_v1.UpdateOtelMetricsIngestionParams{
		Context: ctx,
		Body: &models.Configv1UpdateOtelMetricsIngestionRequest{
			CreateIfMissing:      opts.CreateIfMissing,
			DryRun:               opts.DryRun,
			OtelMetricsIngestion: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewOtelMetricsIngestion(res.Payload.OtelMetricsIngestion), nil
}

func newOtelMetricsIngestionUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing OtelMetricsIngestion.",
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

			otelMetricsIngestion, err := types.MustDecodeSingleObject[*OtelMetricsIngestion](file, permissiveParsing)
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

			fullOtelMetricsIngestion, err := UpdateOtelMetricsIngestion(ctx, client, otelMetricsIngestion, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("OtelMetricsIngestion is valid and can be updated")
				return nil
			}
			stderr.Printf("OtelMetricsIngestion applied successfully\n")

			if err := outputFlags.WriteObject(fullOtelMetricsIngestion, cmd.OutOrStdout()); err != nil {
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
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the OtelMetricsIngestion if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteOtelMetricsIngestion(
	ctx context.Context,
	client config_v1.ClientService,
) error {
	_, err := client.DeleteOtelMetricsIngestion(&config_v1.DeleteOtelMetricsIngestionParams{
		Context: ctx,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newOtelMetricsIngestionDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete",
		GroupID: groups.Commands.ID,
		Short:   "Deletes the OtelMetricsIngestion singleton",
		Args:    cobra.NoArgs,
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

			res, err := client.DeleteOtelMetricsIngestion(&config_v1.DeleteOtelMetricsIngestionParams{
				Context: ctx,
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted OtelMetricsIngestion")
			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

const OtelMetricsIngestionScaffoldYAML = `api_version: v1/config
kind: OtelMetricsIngestion
spec:
    resource_attributes:
        # Do not copy any resource attribute whose key exactly matches one of the
        # strings in this list.
        exclude_keys:
            - <string>
        # Generate a target_info time series with labels derived from resource
        # attributes. The filter_mode and exclude_keys settings apply in the same way as
        # for the "flatten" operation. The default is false.
        generate_target_info: <true|false>
        filter_mode: <APPEND_DEFAULT_EXCLUDE_KEYS|CUSTOM_EXCLUDE_KEYS>
        flatten_mode: <MERGE|OVERWRITE|IGNORE>
`

func newOtelMetricsIngestionScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), OtelMetricsIngestionScaffoldYAML)
		},
	}
	return cmd
}

func NewOtelMetricsIngestionCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "otel-metrics-ingestion",
		GroupID: groups.Config.ID,
		Short:   "All commands for OtelMetricsIngestion",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newOtelMetricsIngestionCreateCmd(),
		newOtelMetricsIngestionReadCmd(),
		newOtelMetricsIngestionUpdateCmd(),
		newOtelMetricsIngestionDeleteCmd(),
		newOtelMetricsIngestionScaffoldCmd(),
	)
	return root
}
