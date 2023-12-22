// Copyright 2023 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/dry"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/file"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/groups"
	"github.com/chronosphereio/chronoctl-core/src/generated/cli/configunstable"
	"github.com/chronosphereio/chronoctl-core/src/generated/cli/configv1"
	config_unstable "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/client/operations"
	config_v1 "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/types"

	"github.com/spf13/cobra"
)

// ErrNoApplyMapping allows an Applier to indicate that no mapping was found
// for the given object.
var ErrNoApplyMapping = errors.New("no apply mapping found")

// Applier represents a function that can apply an object.
// May return ErrNoApplyMapping.
type Applier func(o *ApplyCommand, obj types.Object, stdout io.Writer, stderr io.Writer) error

// Flagger is a interface for adding a flag to a cobra command.
type Flagger interface {
	AddFlags(cmd *cobra.Command)
}

// ApplyOptions contains all options for the apply command.
type ApplyOptions struct {
	Applier         Applier
	AdditionalFlags Flagger
}

// ApplyCommand is a command that allows an object to be upserted.
type ApplyCommand struct {
	ClientFlags *client.Flags

	UnstableClient config_unstable.ClientService
	ConfigV1Client config_v1.ClientService

	FileFlags *file.Flags
	File      io.ReadCloser

	DryRunFlags     *dry.Flags
	StrictModeFlags *StrictModeFlags

	Applier Applier
}

// NewApply returns a new ApplyCommand
func NewApply(options ApplyOptions) *ApplyCommand {
	applier := DefaultApplier
	if options.Applier != nil {
		applier = options.Applier
	}
	return &ApplyCommand{
		Applier:         applier,
		ClientFlags:     client.NewClientFlags(),
		FileFlags:       file.NewFlags(true),
		DryRunFlags:     dry.NewFlags(),
		StrictModeFlags: NewStrictModeFlags(),
	}
}

// NewApplyCommand returns a new cobra command for apply.
func NewApplyCommand(options ApplyOptions) *cobra.Command {
	o := NewApply(options)
	cmd := &cobra.Command{
		Use:   "apply -f [file]",
		Short: "Apply multiple resources by filename",
		Long:  `Apply multiple resources by filename. The resources are upserted, so resources found with a matching slug are updated, otherwise they will be created. JSON and YAML formats are accepted.`,
		Example: `# Apply the resources in resources.yml.
chronoctl apply -f resources.yml

# Dry run to validate the resources in resources.yml can be applied
chronoctl apply -f resources.yml -d

# Apply the resources passed into stdin.
chronoctl apply -f -`,
		Args:    cobra.NoArgs,
		GroupID: groups.Commands.ID,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.validate(); err != nil {
				return err
			}
			if err := o.Run(cmd.OutOrStdout(), cmd.ErrOrStderr()); err != nil {
				return err
			}
			return nil
		},
	}

	o.ClientFlags.AddFlags(cmd)
	o.FileFlags.AddFlags(cmd)
	o.DryRunFlags.AddFlags(cmd)
	o.StrictModeFlags.AddFlags(cmd)
	if options.AdditionalFlags != nil {
		options.AdditionalFlags.AddFlags(cmd)
	}

	return cmd
}

func (o *ApplyCommand) validate() error {
	configV1Client, err := o.ClientFlags.ConfigV1Client()
	if err != nil {
		return err
	}
	o.ConfigV1Client = configV1Client

	unstableClient, err := o.ClientFlags.ConfigUnstableClient()
	if err != nil {
		return err
	}
	o.UnstableClient = unstableClient

	file, err := o.FileFlags.File()
	if err != nil {
		return err
	}

	o.File = file

	return nil
}

// Run runs the apply command.
func (o *ApplyCommand) Run(stdout io.Writer, stderr io.Writer) error {
	// Close will only return an error if Close is called multiple times. We're only calling
	// it once in the defer function and since its called after the run has completed it won't
	// impact the result of the method so we can ignore it.
	defer o.File.Close() //nolint:errcheck

	objs, err := types.Decode(o.File, o.StrictModeFlags.PermissiveParsing)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		dryRun := o.DryRunFlags.DryRun
		dryRunPrefix := ""
		if dryRun {
			dryRunPrefix = "[dry run] "
		}
		fmt.Fprintf(stdout, "%sApplying %s ", dryRunPrefix, obj.Description())

		if err := o.Applier(o, obj, stdout, stderr); err != nil {
			if errors.Is(err, ErrNoApplyMapping) {
				// no matching apply found
				return fmt.Errorf("unsupported object type: %v", obj.Type())
			} else {
				return fmt.Errorf("error applying resource %s: %v", obj.Description(), err)
			}
		}
	}

	return nil
}

// DefaultApplier applies configv1 and configunstable resources.
func DefaultApplier(o *ApplyCommand, obj types.Object, out, err io.Writer) error {
	ctx, cancel := context.WithTimeout(context.Background(), o.ClientFlags.Timeout())
	defer cancel()

	objType := obj.Type()

	clients := client.Clients{
		ConfigV1:       o.ConfigV1Client,
		ConfigUnstable: o.UnstableClient,
	}

	apply, ok := lookup(objType, configv1.ApplyMappings(), configunstable.ApplyMappings())
	if !ok {
		return ErrNoApplyMapping
	}

	if err := apply(ctx, clients, obj, o.DryRunFlags.DryRun); err != nil {
		return err
	}

	fmt.Fprintf(out, "(applied successfully)\n")
	return nil
}

func lookup[K comparable, V any](k K, maps ...map[K]V) (V, bool) {
	for _, m := range maps {
		if v, ok := m[k]; ok {
			return v, true
		}
	}

	var zero V
	return zero, false
}
