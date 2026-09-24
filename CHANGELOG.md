# Releases

## Unreleased

### Added
* Add support for resource `v1/config/CommandCenterGroup` and the `command-center-groups` commands. Command center groups are promoted from the unstable API to v1: a group names a primary SLO and a list of related SLOs by slug. The v1 resource does not carry the deprecated `group_slo_reference` field of the unstable resource; use `primary_slo_reference`.

## v1.35.0

### Added
* Add a `dashboards:importFromClassic` endpoint to `v1/config`, which converts classic dashboard JSON into a standard Dashboard and creates it, or updates an existing dashboard in place when `update_if_exists` is set. Available to library consumers only.
* Add a `dashboards:validate` endpoint to `unstable/config`, which validates dashboard JSON without persisting it. Available to library consumers only.
* Add support for resource `v1/config/MetricNameActiveSeriesLimit` and the `metric-name-active-series-limits` commands. A metric name active series limit caps the active time series allowed for a single metric name.
* Add `description` to `v1/config/Monitor`, an optional Markdown description rendered in the Chronosphere app.
* Add a `label_matchers` filter to listing `v1/config/Monitor`, and a matching `--label-matchers` flag to `monitors list`. Each entry is `name=value` or `name!=value`; all entries must hold for a monitor to be returned.
* Add `recurrence` to `v1/config/MutingRule`, which mutes on a repeating weekly schedule instead of continuously.
* Add `GCP_US_LOS_ANGELES` location to `v1/config/SyntheticTest`.
* Add `primary_slo_reference` and `related_slo_references` to `unstable/config/CommandCenterGroup`. Available to library consumers only.

### Removed
* Remove `unstable/config/MetricNameActiveSeriesLimit` from the generated `configunstable` client, replaced by `v1/config/MetricNameActiveSeriesLimit` and the `metric-name-active-series-limits` commands. The unstable endpoints no longer exist in the public API. Available to library consumers only. **This is a breaking change**.
* Rename the generated `configv1` monitor schedule model types (`ScheduleWeeklySchedule` becomes `MonitorScheduleWeeklySchedule`, `ScheduleDayTimeRange` becomes `ScheduleScheduleDayTimeRange`). Available to library consumers only. **This is a breaking change**.
* Remove `dns_server` and `dns_server_port` from the `v1/config/SyntheticTest` DNS test configuration. The API no longer accepts a custom nameserver, so drop the fields from any `synthetic-tests` YAML that sets them. **This is a breaking change**.
* Remove the `dashboards:createFromClassic` endpoint from `unstable/config`, replaced by `v1/config` `dashboards:importFromClassic`. Available to library consumers only. **This is a breaking change**.

## v1.34.0

### Added
* Add `PUT` and `DELETE` HTTP methods to `v1/config/SyntheticTest` HTTP test configuration.
* Add support for resource `unstable/config/QueryResourcePools` and the `query-resource-pools` commands. Query resource pools group automated metrics query sources (monitors, recording rules, SLOs, and service accounts) into pools with per-pool datapoints-read rate limits.

### Changed
* Rewrite field documentation across `v1/config/SyntheticTest` models for clarity; no functional changes.

## v1.33.0

### Added
* Add support for resource `v1/config/SyntheticTest` and the `synthetic-tests` commands. Synthetic tests are promoted from the unstable API to v1: configures HTTP, DNS, TCP, and TLS probes that run from a set of locations on a schedule, with per-type assertions, retry config, and optional alerting on failures.
* Add `DATADOG` connection type and `datadog` config to `v1/config/ExternalConnection`.
* Add a `dashboards:createFromClassic` endpoint to `unstable/config`, which converts raw classic (Grafana) dashboard JSON and creates the result as a standard Dashboard in one call.
* Add `severity` to `v1/state` rule evaluations, distinguishing hard evaluation failures from non-fatal warnings, and an `include_warnings` query parameter on the rule evaluation list endpoint. Both are available to library consumers only; the `rule-evaluations list` command does not expose them.
* Add support for resource `v1/config/CloudIntegration` and the `cloud-integrations` commands. Cloud integrations are promoted from the unstable API to v1: configures ingestion of metrics from an external provider through an external connection, with a `provider_config` object selecting what to ingest and `metric_labels` attached to ingested metrics.
* Add `slo_type` and `latency_threshold_nanos` to the `v1/config/SLO` SLI. `slo_type` records whether the SLO measures endpoint availability or latency, and `latency_threshold_nanos` is required for latency SLOs; both record intent only, and query generation is unchanged.
* Add a `name_contains` filter to listing `v1/config/Monitor`, and a matching `--name-contains` flag to `monitors list`. The filter is case-insensitive.
* Add a `collection` reference and an update-time `dry_run` flag to `unstable/config/Notebook`. Available to library consumers only.

### Changed
* Rename the generated `configv1` enum type `FilterOperator` to `PartitionFilterOperator` in `v1/config/ConsumptionConfig` partition filter conditions (enum values are unchanged). Available to library consumers only.
* Replace the `configunstable` client's `SyncPrometheus` error payload type `GooglerpcStatus` with `APIError`, and remove the generated `GooglerpcStatus` and `ProtobufAny` models. Available to library consumers only. **This is a breaking change**.

### Removed
* Remove support for resources `v1/config/LogScaleAlert` and `v1/config/LogScaleAction`, along with the `log-scale-alerts` and `log-scale-actions` commands. These endpoints no longer exist in the public API. **This is a breaking change**.
* Remove `unstable/config/SyntheticTest` from the generated `configunstable` client, replaced by `v1/config/SyntheticTest` and the `synthetic-tests` commands. The unstable endpoints no longer exist in the public API. Available to library consumers only; there were never unstable `synthetic-tests` commands. **This is a breaking change**.
* Remove `unstable/config/CloudIntegration` from the generated `configunstable` client, replaced by `v1/config/CloudIntegration` and the `cloud-integrations` commands. The unstable endpoints no longer exist in the public API. Available to library consumers only; there were never unstable `cloud-integrations` commands. **This is a breaking change**.

## v1.32.0

### Added
* Add `disable_repeat` field to `v1/config/NotificationPolicy` routes. When `true`, alerts are notified once and on state changes but are never re-sent on the repeat interval.
* Add `SENDGRID` connection type and `sendgrid` config to `v1/config/ExternalConnection`.
* Add `resolve_sustain_for_no_data` field to `v1/config/Monitor` series conditions, controlling how a firing condition resolves once its series stops returning data.
* Add `new_series_delay_secs` field to `v1/config/Monitor` conditions, suppressing alerts for a newly observed series until the delay has elapsed since the series was first seen.
* Add `CLOUDFLARE` provider type and `cloudflare` config to `unstable/config/CloudIntegration`.

## v1.31.0

### Added
* Add `TRACE_PROCESSED_BYTES`, `TRACE_PERSISTED_BYTES`, and `TRACE_ALL` resource groups to `v1/config/ConsumptionBudget`.
* Add `ALL` resource group to `v1/config/ConsumptionBudget`.

## v1.30.0

### Added
* Add `MONGODB_ATLAS` connection type and `mongodb_atlas` config to `v1/config/ExternalConnection`.
* Add `LOG_ALL` resource group to `v1/config/ConsumptionBudget`.

### Changed
* Replace the `trace_filter` field with `trace_span_filters` in `v1/config/ConsumptionConfig` partition filter conditions.

## v1.29.0

### Added
* Add `PARENT_SERVICE`, `PARENT_OPERATION`, `ROOT_SERVICE`, and `ROOT_OPERATION` group-by key types to `v1/config/TraceMetricsRule`.

### Removed
* Remove the deprecated `sku_group` field from `v1/config/ConsumptionBudget` thresholds. Use `resource_group` instead.

## v1.28.0

### Added
* Add `ROLLING_30_MINUTE_VOLUME` threshold type to `v1/config/ConsumptionBudget`.
* Add `CLOUDFLARE` connection type and `cloudflare` config to `v1/config/ExternalConnection`.
* Add `GCP_US_LOS_ANGELES` location to `unstable/config/SyntheticTest`, replacing `AWS_US_EAST_1`.
* Add `resource_group` field to `v1/config/ConsumptionBudget` thresholds as an alias of the deprecated `sku_group`. Regenerated the `configv1` swagger client/models from the updated public spec, which renames the generated enum type `ConsumptionBudgetSKUGroup` to `ConsumptionBudgetResourceGroup` (enum values are unchanged).

## v1.27.0

### Changed
* Changed the singular `routing_key` field in the VictorOps destination in `v1/config/NotificationPolicy` to be plural `routing_keys`

## v1.26.0

Added:
* Add `metric_filters` and `trace_filter` fields to `v1/config/ConsumptionConfig` partition filter conditions.
* Add `METRIC_PERSISTED_SERIES` and `METRIC_ALL` SKU groups and `CREDITS` unit to `v1/config/ConsumptionBudget`.
* Add `parse_field` rule type to `v1/config/LogControlConfig`.
* Add `grok_parser` to `v1/config/LogIngestConfig`.

Internal:
* Upgrade to Go 1.26.

## v1.25.0

Added:
* Add support for resource `v1/config/ExternalConnection`. Manages external connections for notification targets including Slack, PagerDuty, Webhook, VictorOps, and OpsGenie.
* Update `v1/config/NotificationPolicy` to allow configuration of notification routing to an external connection.

## v1.24.0

Added:
* Add `notification_template` field to `v1/config/Monitor` for customizing notification title and description.

## v1.23.0

Added:
* Add `resolve_value` field to `v1/config/Monitor` conditions. Allows configuring a different threshold for resolving an alert.

## v1.22.0

Added:
* Add `execution_mode` field to `v1/config/RecordingRule`.
* Pass the value of the `CHRONOSPHERE_ACTOR` environment variable in header `Chronosphere-Actor` header on API requests. The value is recorded as change metadata in Version History.
* Add support for resource `v1/config/LogRetentionConfig`.
* Add `SIGNAL_NOT_EXISTS` condition op to `v1/config/Monitor`. Signal no longer exists. Behavior varies by alert type: single monitor alerts trigger if all series are missing; signal-based alerts trigger if all series in the signal go missing; alert-on-every-series alerts trigger if any individual series disappears.

## v1.20.0

Added:
* Add `service` to `chronosphere_log_ingest_config` resource. This field replaces the `primary_key` field, which is deprecated. **This is a breaking change**.
* Add `ROLLING_1_DAY_VOLUME` and `ROLLING_7_DAY_VOLUME` thresholds to `chronosphere_consumption_budget` resource.

Deprecated:
* Deprecate the `primary_key` field in the `chronosphere_log_ingest_config` resource in favor of the `service` field. This field can no longer be set.

## v1.18.0

Added:
* Add `ROLLING_1_DAY_VOLUME` and `ROLLING_7_DAY_VOLUME` thresholds to `v1/config/ConsumptionBudget`.
* Update Prom label grammar to support UTF-8 label names.

## v1.17.0

Added:
* Add `skip_on_conflict` field to `v1/config/RollupRule`.
* Adds new resource for service attributes. 

## v1.16.0

* Adds new resource `v1/config/AzureMetricsIntegration`
* Adds `emit_metrics` and `replace_fields` control rule types to `chronosphere_log_control_config` resource.
* Adds `field_normalization` to `chronosphere_log_ingest_config` resource.

## v1.15.0

* Update parser resource in `v1/config/LogIngestConfig`
* Adds unstable resource `unstable/config/AzureMetricsIntegration`
* Fix ConsumptionConfig, LogAllocationConfig, LogControlConfig, and TraceBehaviorConfig kinds in scaffold command
* Adds new resource `v1/config/ConsumptionConfig` 
* Adds new resource `v1/config/ConsumptionBudget` 

## v1.14.0

Added:
* Add support for resource `v1/config/LogIngestConfig`
* Add support for TimeSlice SLOs

## v1.13.0

Added:
* Move SLOs stable public API in `v1/config/Slos`

## v1.12.0

Added:
* Adds support for pool PriorityThresholds in `v1/config/ResourcePools`
* Adds `chronoctl auth whoami` which prints the current user
* Adds `chronoctl auth logout` which logs out an authenticated chronoctl session

## v1.11.0

Added:
* Adds support for labels to `v1/config/Dashboard`
* Adds new unstable resource `unstable/config/SLO`
* add `is_root_span` to trace search span filter in entities:
  * `v1/config/TraceMetricRules`
  * `v1/config/TraceTailSamplingRules`
  * `v1/config/Dataset` trace datasets
* Adds `preview_behavior_assignments` to `v1/config/TraceBehaviorConfig` for previewing changes to trace behavior assignments

## v1.10.1

No changes

## v1.10.0

### Added
* Adds new resource `v1/config/LogScaleAction`
* Adds new resource `v1/config/LogScaleAlert`
* Adds new field `pool.allocation.fixed_values` to `v1/config/ResourcePools`
* Adds new state command `metric-usages-by-metric-name`
* Adds new state command `metric-usages-by-label-name`

## v1.9.0

### Added
* Adds new resource `v1/config/OtelMetricsIngestion`
* Adds support for logs in `v1/config/Dataset`

## v1.8.0

* New command: auth

## v1.7.0

### Added
* Adds new resource `v1/config/TraceBehaviorConfig`
* Adds new field `logging_query` to `v1/config/Monitor`
* Adds new field `group_by` to `v1/config/NotificationPolicy`

## v1.6.0

## Added
* New config: `v1/config/Datasets`
* New config: `v1/config/GcpMetricsIntegration`
* Adds support for entity namespaces

## v1.5.0

### Added
* New config: gcp-metrics-integrations

## v1.4.0

### Added
* Adds new field `graphite_label_policy` to `v1/config/RollupRule`

## v1.3.0

### Added
* Adds new field `drop_nan_value` to `v1/config/DropRule`
* Adds new metric type `DELTA_EXPONENTIAL_HISTOGRAM` to `v1/config/RollupRule`.

### Deprecated
* Unused field `label_replace` has been removed from `v1/config/RollupRule`.

## v1.2.0

### Added
- New config: trace-tail-sampling-rules

## v1.1.0

### Added
 - New config: trace-jaeger-remote-sampling-strategies.
 - New config: trace-metric-rules.
 - Support for new aggregation type, `COUNT_SAMPLES`.

### Fixed
 - Removed read-only fields from scaffold output for `resource-pools`.

### Misc
 - Upgraded participle library to v2.1.0
 - Added goreleaser support for releasing binaries

## v1.0.0

Initial release for chronoctl-core.

