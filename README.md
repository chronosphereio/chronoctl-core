# Chronoctl

Chronoctl is a command-line app that's a wrapper for accessing the Chronosphere API,
letting you manage resources with YAML or JSON. The app supports the `create`,
`read`, `update`, `delete`, and `list` operations as subcommands for each type of resource.

## Usage

`chronoctl COMMAND`

## Commands

| Command | Description |
|---------|-------------|
| `apply` | Apply multiple resources defined in a file |

## Configuration

Chronoctl provides commands to configure Chronosphere, each of which corresponds
to a resource type.

* `buckets`
* `classic-dashboards` (`grafana-dashboards`)
* `collections`
* `dashboards`
* `derived-labels`
* `derived-metrics`
* `drop-rules`
* `gcp-metrics-integrations`
* `mapping-rules`
* `monitors`
* `muting-rules`
* `notification-policies`
* `notifiers`
* `recording-rules`
* `resource-pools`
* `rollup-rules`
* `service-accounts`
* `teams`

These resource commands have subcommands to read and manipulate entities via
`create`, `read`, `update`, `delete`, and `list`.
The `scaffold` subcommand generates YAML with all supported fields and
comments with documentation.

For details about these resource types, see their
[documentation](https://docs.chronosphere.io).

## State

Chronoctl also provides a command to review the state of rule evaluations in
Chronosphere:

* `rule-evaluations`

## Additional commands

| Command | Description |
|---------|-------------|
| `completion` | Generate the autocompletion script for the specified shell |
| `help` | Help about any command |
| `version` | Print the current version of chronoctl |

## Building

To build chronoctl, you must have Go 1.21 installed. A download link can be found at
[go.dev](https://go.dev/doc/install).

Run `make chronoctl` to build the binary. It will be found under
`./bin/chronoctl`

## Backward compatibility

This repository guarantees backward compatibility between binary releases. This
repository is not intended for use as a library, and doesn't guarantee backward
compatibility of the code packages.

## Support

For any questions, bugs, or feature requests, please file a ticket through our standard support channels.
