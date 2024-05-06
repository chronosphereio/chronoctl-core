// Code generated by chronogen; DO NOT EDIT
package configv1

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/types"
)

// AddCommandsTo adds all entity subcommands to the given root command.
func AddCommandsTo(root *cobra.Command) {
	root.AddCommand(NewBucketCmd())
	root.AddCommand(NewClassicDashboardCmd())
	root.AddCommand(NewCollectionCmd())
	root.AddCommand(NewDashboardCmd())
	root.AddCommand(NewDatasetCmd())
	root.AddCommand(NewDerivedLabelCmd())
	root.AddCommand(NewDerivedMetricCmd())
	root.AddCommand(NewDropRuleCmd())
	root.AddCommand(NewGcpMetricsIntegrationCmd())
	root.AddCommand(NewGrafanaDashboardCmd())
	root.AddCommand(NewMappingRuleCmd())
	root.AddCommand(NewMonitorCmd())
	root.AddCommand(NewMutingRuleCmd())
	root.AddCommand(NewNotificationPolicyCmd())
	root.AddCommand(NewNotifierCmd())
	root.AddCommand(NewRecordingRuleCmd())
	root.AddCommand(NewResourcePoolsCmd())
	root.AddCommand(NewRollupRuleCmd())
	root.AddCommand(NewServiceAccountCmd())
	root.AddCommand(NewTeamCmd())
	root.AddCommand(NewTraceJaegerRemoteSamplingStrategyCmd())
	root.AddCommand(NewTraceMetricsRuleCmd())
	root.AddCommand(NewTraceTailSamplingRulesCmd())
}

func ApplyMappings() map[types.TypeMeta]func(context.Context, client.Clients, types.Object, bool) error {
	return map[types.TypeMeta]func(context.Context, client.Clients, types.Object, bool) error{
		BucketTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*Bucket)
			if !ok {
				return types.WrongObjectErr((&Bucket{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateBucket(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		ClassicDashboardTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*ClassicDashboard)
			if !ok {
				return types.WrongObjectErr((&ClassicDashboard{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateClassicDashboard(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		CollectionTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*Collection)
			if !ok {
				return types.WrongObjectErr((&Collection{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateCollection(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		DashboardTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*Dashboard)
			if !ok {
				return types.WrongObjectErr((&Dashboard{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateDashboard(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		DatasetTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*Dataset)
			if !ok {
				return types.WrongObjectErr((&Dataset{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateDataset(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		DerivedLabelTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*DerivedLabel)
			if !ok {
				return types.WrongObjectErr((&DerivedLabel{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateDerivedLabel(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		DerivedMetricTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*DerivedMetric)
			if !ok {
				return types.WrongObjectErr((&DerivedMetric{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateDerivedMetric(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		DropRuleTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*DropRule)
			if !ok {
				return types.WrongObjectErr((&DropRule{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateDropRule(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		GcpMetricsIntegrationTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*GcpMetricsIntegration)
			if !ok {
				return types.WrongObjectErr((&GcpMetricsIntegration{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateGcpMetricsIntegration(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		GrafanaDashboardTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*GrafanaDashboard)
			if !ok {
				return types.WrongObjectErr((&GrafanaDashboard{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateGrafanaDashboard(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		MappingRuleTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*MappingRule)
			if !ok {
				return types.WrongObjectErr((&MappingRule{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateMappingRule(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		MonitorTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*Monitor)
			if !ok {
				return types.WrongObjectErr((&Monitor{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateMonitor(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		MutingRuleTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*MutingRule)
			if !ok {
				return types.WrongObjectErr((&MutingRule{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateMutingRule(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		NotificationPolicyTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*NotificationPolicy)
			if !ok {
				return types.WrongObjectErr((&NotificationPolicy{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateNotificationPolicy(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		NotifierTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*Notifier)
			if !ok {
				return types.WrongObjectErr((&Notifier{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateNotifier(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		RecordingRuleTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*RecordingRule)
			if !ok {
				return types.WrongObjectErr((&RecordingRule{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateRecordingRule(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		ResourcePoolsTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*ResourcePools)
			if !ok {
				return types.WrongObjectErr((&ResourcePools{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateResourcePools(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		RollupRuleTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*RollupRule)
			if !ok {
				return types.WrongObjectErr((&RollupRule{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateRollupRule(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		TeamTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*Team)
			if !ok {
				return types.WrongObjectErr((&Team{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateTeam(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		TraceJaegerRemoteSamplingStrategyTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*TraceJaegerRemoteSamplingStrategy)
			if !ok {
				return types.WrongObjectErr((&TraceJaegerRemoteSamplingStrategy{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateTraceJaegerRemoteSamplingStrategy(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		TraceMetricsRuleTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*TraceMetricsRule)
			if !ok {
				return types.WrongObjectErr((&TraceMetricsRule{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateTraceMetricsRule(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
		TraceTailSamplingRulesTypeMeta: func(ctx context.Context, clients client.Clients, obj types.Object, dryRun bool) error {
			entity, ok := obj.(*TraceTailSamplingRules)
			if !ok {
				return types.WrongObjectErr((&TraceTailSamplingRules{}), obj)
			}

			updateOpts := UpdateOptions{
				DryRun:          dryRun,
				CreateIfMissing: true,
			}
			_, err := UpdateTraceTailSamplingRules(ctx, clients.ConfigV1, entity, updateOpts)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

// UpdateOptions represents the request level options for update.
type UpdateOptions struct {
	DryRun          bool
	CreateIfMissing bool
}
