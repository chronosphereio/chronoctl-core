# Releases

## Unreleased

* Added support for entity namespaces
* Added support for datasets
* unstable: Added support for otel-metrics-ingestion
* Promote gcp-metrics-integration to stable

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
