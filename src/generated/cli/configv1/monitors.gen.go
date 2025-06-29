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

func init() { types.MustRegisterObject(MonitorTypeMeta, &Monitor{}) }

var _ types.Object = &Monitor{}

var MonitorTypeMeta = types.TypeMeta{
	APIVersion: "v1/config",
	Kind:       "Monitor",
}

type Monitor struct {
	types.TypeMeta `json:",inline"`
	Spec           *models.Configv1Monitor `json:"spec"`
}

func NewMonitor(spec *models.Configv1Monitor) *Monitor {
	return &Monitor{
		TypeMeta: MonitorTypeMeta,
		Spec:     spec,
	}
}

func (e *Monitor) Description() string {
	return types.TypeDescription(e, "name", e.Spec.Name, "slug", e.Spec.Slug)
}

func (e *Monitor) Identifier() string {
	return e.Spec.Slug
}

func CreateMonitor(
	ctx context.Context,
	client config_v1.ClientService,
	entity *Monitor,
	dryRun bool,
) (*Monitor, error) {
	res, err := client.CreateMonitor(&config_v1.CreateMonitorParams{
		Context: ctx,
		Body: &models.Configv1CreateMonitorRequest{
			DryRun:  dryRun,
			Monitor: entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewMonitor(res.Payload.Monitor), nil
}

func newMonitorCreateCmd() *cobra.Command {
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
	short = "Creates a single Monitor."

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

			var monitor *Monitor
			file, err := fileFlags.File()
			if err != nil {
				return err
			}
			defer file.Close() //nolint:errcheck
			monitor, err = types.MustDecodeSingleObject[*Monitor](file, permissiveParsing)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("--dry-run is set")
			}
			fullMonitor, err := CreateMonitor(ctx, client, monitor, dryRunFlags.DryRun)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("Monitor is valid and can be created")
				return nil
			}
			stderr.Printf("Monitor with slug %q created successfully\n", fullMonitor.Spec.Slug)

			if err := outputFlags.WriteObject(fullMonitor, cmd.OutOrStdout()); err != nil {
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

func GetMonitor(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) (*Monitor, error) {
	res, err := client.ReadMonitor(&config_v1.ReadMonitorParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}
	return NewMonitor(res.GetPayload().Monitor), nil
}

func newMonitorReadCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())
	var (
		short string
		use   string
		args  cobra.PositionalArgs
	)
	short = "Reads a single Monitor by slug"
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
			entity, err := GetMonitor(ctx, client, args[0])
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

func UpdateMonitor(
	ctx context.Context,
	client config_v1.ClientService,
	entity *Monitor,
	opts UpdateOptions,
) (*Monitor, error) {
	res, err := client.UpdateMonitor(&config_v1.UpdateMonitorParams{
		Context: ctx,
		Slug:    entity.Spec.Slug,
		Body: &models.ConfigV1UpdateMonitorBody{
			CreateIfMissing: opts.CreateIfMissing,
			DryRun:          opts.DryRun,
			Monitor:         entity.Spec,
		},
	})
	if err != nil {
		return nil, clienterror.Wrap(err)
	}

	return NewMonitor(res.Payload.Monitor), nil
}

func newMonitorUpdateCmd() *cobra.Command {
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
		Short:   "Updates an existing Monitor.",
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

			monitor, err := types.MustDecodeSingleObject[*Monitor](file, permissiveParsing)
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

			fullMonitor, err := UpdateMonitor(ctx, client, monitor, updateOpts)
			if err != nil {
				return err
			}

			if dryRunFlags.DryRun {
				stderr.Println("Monitor is valid and can be updated")
				return nil
			}
			stderr.Printf("Monitor with slug %q applied successfully\n", fullMonitor.Spec.Slug)

			if err := outputFlags.WriteObject(fullMonitor, cmd.OutOrStdout()); err != nil {
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
	cmd.Flags().BoolVar(&createIfMissing, "create-if-missing", false, "If set, creates the Monitor if it does not already exist. Defaults to false.")

	return cmd
}

func DeleteMonitor(
	ctx context.Context,
	client config_v1.ClientService,
	slug string,
) error {
	_, err := client.DeleteMonitor(&config_v1.DeleteMonitorParams{
		Context: ctx,
		Slug:    slug,
	})
	if err != nil {
		return clienterror.Wrap(err)
	}
	return nil
}

func newMonitorDeleteCmd() *cobra.Command {
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags(output.WithoutOutputDirectory(), output.WithoutCreateFilePerObject())

	cmd := &cobra.Command{
		Use:     "delete <slug>",
		GroupID: groups.Commands.ID,
		Short:   "Deletes a single Monitor by slug",
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

			res, err := client.DeleteMonitor(&config_v1.DeleteMonitorParams{
				Context: ctx,
				Slug:    args[0],
			})
			if err != nil {
				return clienterror.Wrap(err)
			}
			_ = res
			fmt.Fprintf(cmd.OutOrStdout(), "deleted Monitor with slug %q\n", args[0])

			return nil
		},
	}
	clientFlags.AddFlags(cmd)
	outputFlags.AddFlags(cmd)
	return cmd
}

type MonitorListOpts struct {
	// Limit represents that maximum number of items we wish to return.
	Limit int
	// PageToken is the pagination token we want to start our request at.
	PageToken string
	// PageMaxSize is the maximum page size to use when making List calls.
	PageMaxSize     int
	BucketSlugs     []string
	CollectionSlugs []string
	Names           []string
	Slugs           []string
	TeamSlugs       []string
}

func (r *MonitorListOpts) registerFlags(flags *flag.FlagSet) {
	var emptyBucketSlugs []string
	flags.StringSliceVar(&r.BucketSlugs, "bucket-slugs", emptyBucketSlugs, "Filters results by bucket_slug, where any Monitor with a matching bucket_slug in the given list (and matches all other filters) is returned.")
	var emptyCollectionSlugs []string
	flags.StringSliceVar(&r.CollectionSlugs, "collection-slugs", emptyCollectionSlugs, "Filters results by collection_slug, where any Monitor with a matching collection_slug in the given list (and matches all other filters) is returned.")
	var emptyNames []string
	flags.StringSliceVar(&r.Names, "names", emptyNames, "Filters results by name, where any Monitor with a matching name in the given list (and matches all other filters) is returned.")
	var emptySlugs []string
	flags.StringSliceVar(&r.Slugs, "slugs", emptySlugs, "Filters results by slug, where any Monitor with a matching slug in the given list (and matches all other filters) is returned.")
	var emptyTeamSlugs []string
	flags.StringSliceVar(&r.TeamSlugs, "team-slugs", emptyTeamSlugs, "Filter returned monitors by the teams that own the collections that they belong to.")
	flags.IntVar(&r.Limit, "limit", 0, "maximum number of items to return")
	flags.IntVar(&r.PageMaxSize, "page-max-size", 0, "maximum page size")
	flags.StringVar(&r.PageToken, "page-token", "", "begins listing items at the start of the pagination token")
}

func ListMonitors(
	ctx context.Context,
	client config_v1.ClientService,
	streamer output.Streamer[*Monitor],
	opts MonitorListOpts,
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
		res, err := client.ListMonitors(&config_v1.ListMonitorsParams{
			Context:         ctx,
			PageToken:       &nextToken,
			PageMaxSize:     ptr.Int64(int64(pageMaxSize)),
			BucketSlugs:     opts.BucketSlugs,
			CollectionSlugs: opts.CollectionSlugs,
			Names:           opts.Names,
			Slugs:           opts.Slugs,
			TeamSlugs:       opts.TeamSlugs,
		})
		if err != nil {
			return pagination.Token(""), clienterror.Wrap(err)
		}

		for _, v := range res.Payload.Monitors {
			if err := streamer(NewMonitor(v)); err != nil {
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
			MaxPageSize: len(res.Payload.Monitors),
		})
	}
}

func newMonitorListCmd() *cobra.Command {
	var listOptions MonitorListOpts
	clientFlags := client.NewClientFlags()
	outputFlags := output.NewFlags()

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all Monitors and applies optional filters",
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

			streamer := output.NewWriteObjectStreamer[*Monitor](writer)
			nextToken, err := ListMonitors(ctx, client, streamer, listOptions)
			if err != nil {
				return err
			}

			if nextToken != "" {
				nextPage := pagination.Result{
					Kind:          pagination.ResultKind,
					Message:       "There are additional monitors. To view more, use the next page token or run the full command.",
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

const MonitorScaffoldYAML = `api_version: v1/config
kind: Monitor
spec:
    # Unique identifier of the Monitor. If a 'slug' isn't provided, one will be generated based of the 'name' field. You can't modify this field after the Monitor is created.
    slug: <string>
    # Required. Name of the Monitor. You can modify this value after the Monitor is created.
    name: <string>
    # Slug of the bucket the monitor belongs to. Required if 'collection_slug' isn't
    # set.
    bucket_slug: <string>
    # Slug of the collection the monitor belongs to. Required if 'bucket_slug' isn't
    # set.
    collection_slug: <string>
    # Required. Labels to include in notifications generated by this monitor, and can
    # be used to route alerts with notification overrides.
    labels:
        key_1: <string>
    # Annotations are visible in notifications generated by this monitor.
    # They can be be templated with labels from notifications.
    annotations:
        key_1: <string>
    # Specifies the notification policy used to route alerts generated by the monitor.
    # If omitted, the notification policy is inherited from the monitor.
    notification_policy_slug: <string>
    # Specifies how often alerts are evaluated. Default: '60s'.
    interval_secs: <integer>
    # PromQL query to evaluate for the alert. If set, no other queries can be set.
    prometheus_query: <string>
    # Graphite query to evaluate for the alert. If set, no other queries can be set.
    graphite_query: <string>
    # Logging query to evaluate for the alert. If set, no other queries can be set.
    logging_query: <string>
    collection:
        slug: <string>
        # Type values must match entitiespb.Collection.CollectionType.
        type: <SIMPLE|SERVICE>
    schedule:
        # The timezone of the time ranges.
        timezone: <string>
        weekly_schedule:
            friday:
                # The time ranges that the monitor is active on this day. Required if 'active'
                # is set to 'ONLY_DURING_RANGES'. Otherwise, this field must be empty.
                ranges:
                    - # Start time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      start_hh_mm: <string>
                      # End time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      end_hh_mm: <string>
                active: <ALL_DAY|ONLY_DURING_RANGES|NEVER>
            monday:
                # The time ranges that the monitor is active on this day. Required if 'active'
                # is set to 'ONLY_DURING_RANGES'. Otherwise, this field must be empty.
                ranges:
                    - # Start time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      start_hh_mm: <string>
                      # End time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      end_hh_mm: <string>
                active: <ALL_DAY|ONLY_DURING_RANGES|NEVER>
            saturday:
                # The time ranges that the monitor is active on this day. Required if 'active'
                # is set to 'ONLY_DURING_RANGES'. Otherwise, this field must be empty.
                ranges:
                    - # Start time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      start_hh_mm: <string>
                      # End time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      end_hh_mm: <string>
                active: <ALL_DAY|ONLY_DURING_RANGES|NEVER>
            sunday:
                # The time ranges that the monitor is active on this day. Required if 'active'
                # is set to 'ONLY_DURING_RANGES'. Otherwise, this field must be empty.
                ranges:
                    - # Start time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      start_hh_mm: <string>
                      # End time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      end_hh_mm: <string>
                active: <ALL_DAY|ONLY_DURING_RANGES|NEVER>
            thursday:
                # The time ranges that the monitor is active on this day. Required if 'active'
                # is set to 'ONLY_DURING_RANGES'. Otherwise, this field must be empty.
                ranges:
                    - # Start time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      start_hh_mm: <string>
                      # End time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      end_hh_mm: <string>
                active: <ALL_DAY|ONLY_DURING_RANGES|NEVER>
            tuesday:
                # The time ranges that the monitor is active on this day. Required if 'active'
                # is set to 'ONLY_DURING_RANGES'. Otherwise, this field must be empty.
                ranges:
                    - # Start time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      start_hh_mm: <string>
                      # End time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      end_hh_mm: <string>
                active: <ALL_DAY|ONLY_DURING_RANGES|NEVER>
            wednesday:
                # The time ranges that the monitor is active on this day. Required if 'active'
                # is set to 'ONLY_DURING_RANGES'. Otherwise, this field must be empty.
                ranges:
                    - # Start time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      start_hh_mm: <string>
                      # End time in the in format '"<hour>:<minute>"'. For example, '"15:30"'.
                      end_hh_mm: <string>
                active: <ALL_DAY|ONLY_DURING_RANGES|NEVER>
    # Conditions evaluated against each queried series to determine the severity of each series.
    series_conditions:
        # Optional. Specifies a list of overrides to use for series having matching
        # labels. Each override defines labels that potentially match a series' labels.
        # If one or more overrides match a series, the severity conditions of the first
        # matching override are used instead of the defaults.

        # Cannot be used if 'graphite_query' is set.
        overrides:
            - # Set of matchers on a series' labels. If all labels match, then the conditions
              # defined in this override are used.
              label_matchers:
                - # name always matches against an exact label name.
                  name: <string>
                  # value matches against a label value based on the configured type.
                  value: <string>
                  type: <EXACT|REGEX>
              severity_conditions:
                critical:
                    # List of conditions to evaluate against a series. Only one condition must
                    # match to assign a severity to a signal.
                    conditions:
                        - # Required. The value to compare to the metric value using the 'op' operation.
                          value: <number>
                          # Amount of time the query needs to fail the condition check before an alert is
                          # triggered. Must be an integer. Accepts one of 's' (seconds), 'm' (minutes), or
                          # 'h' (hours) as units.
                          sustain_secs: <integer>
                          # Amount of time the query needs to no longer fire before resolving. Must be an
                          # integer. Accepts one of 's' (seconds), 'm' (minutes), or 'h' (hours) as units.
                          resolve_sustain_secs: <integer>
                          op: <GEQ|GT|LEQ|LT|EQ|NEQ|EXISTS|NOT_EXISTS>
                warn:
                    # List of conditions to evaluate against a series. Only one condition must
                    # match to assign a severity to a signal.
                    conditions:
                        - # Required. The value to compare to the metric value using the 'op' operation.
                          value: <number>
                          # Amount of time the query needs to fail the condition check before an alert is
                          # triggered. Must be an integer. Accepts one of 's' (seconds), 'm' (minutes), or
                          # 'h' (hours) as units.
                          sustain_secs: <integer>
                          # Amount of time the query needs to no longer fire before resolving. Must be an
                          # integer. Accepts one of 's' (seconds), 'm' (minutes), or 'h' (hours) as units.
                          resolve_sustain_secs: <integer>
                          op: <GEQ|GT|LEQ|LT|EQ|NEQ|EXISTS|NOT_EXISTS>
        defaults:
            critical:
                # List of conditions to evaluate against a series. Only one condition must
                # match to assign a severity to a signal.
                conditions:
                    - # Required. The value to compare to the metric value using the 'op' operation.
                      value: <number>
                      # Amount of time the query needs to fail the condition check before an alert is
                      # triggered. Must be an integer. Accepts one of 's' (seconds), 'm' (minutes), or
                      # 'h' (hours) as units.
                      sustain_secs: <integer>
                      # Amount of time the query needs to no longer fire before resolving. Must be an
                      # integer. Accepts one of 's' (seconds), 'm' (minutes), or 'h' (hours) as units.
                      resolve_sustain_secs: <integer>
                      op: <GEQ|GT|LEQ|LT|EQ|NEQ|EXISTS|NOT_EXISTS>
            warn:
                # List of conditions to evaluate against a series. Only one condition must
                # match to assign a severity to a signal.
                conditions:
                    - # Required. The value to compare to the metric value using the 'op' operation.
                      value: <number>
                      # Amount of time the query needs to fail the condition check before an alert is
                      # triggered. Must be an integer. Accepts one of 's' (seconds), 'm' (minutes), or
                      # 'h' (hours) as units.
                      sustain_secs: <integer>
                      # Amount of time the query needs to no longer fire before resolving. Must be an
                      # integer. Accepts one of 's' (seconds), 'm' (minutes), or 'h' (hours) as units.
                      resolve_sustain_secs: <integer>
                      op: <GEQ|GT|LEQ|LT|EQ|NEQ|EXISTS|NOT_EXISTS>
    # SignalGrouping defines how the set of series from the query are split into signals.
    signal_grouping:
        # Set of label names used to split series into signals. Each unique combination
        # of labels result in its own signal. For example, if 'label_names' is
        # '["service", "code"]', then all series including labels
        # '{service="foo",code="404"}' will be grouped together in the same signal.

        # Cannot be used if 'graphite_query' is set.
        label_names:
            - <string>
        # If set to 'true', each series will have its own signal. Cannot be used with
        # 'label_names'.
        signal_per_series: <true|false>
`

func newMonitorScaffoldCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		GroupID: groups.Commands.ID,
		Short:   "Scaffolds a complete object with placeholder values",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), MonitorScaffoldYAML)
		},
	}
	return cmd
}

func NewMonitorCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "monitors",
		GroupID: groups.Config.ID,
		Short:   "All commands for Monitors",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newMonitorCreateCmd(),
		newMonitorReadCmd(),
		newMonitorUpdateCmd(),
		newMonitorDeleteCmd(),
		newMonitorListCmd(),
		newMonitorScaffoldCmd(),
	)
	return root
}
