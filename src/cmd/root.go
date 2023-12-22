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
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/groups"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/ruleevaluations"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/unstable"
	"github.com/chronosphereio/chronoctl-core/src/generated/cli/configv1"
	"github.com/spf13/cobra"
)

// Options contain all the options for creating a new chronoctl command.
type Options struct {
	ApplyOptions ApplyOptions
}

// New returns a new chronoctl command.
func New(options Options) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "chronoctl",
		Short: "chronoctl is used to manage resources in the Chronosphere API",
		Long:  "chronoctl is used to manage resources in the Chronosphere API.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// If command parsing works, let's silence usage so errors RunE errors
			// don't display usage (adding unnecessary noise to the output).
			cmd.SilenceUsage = true
		},
	}

	// AddGroup will display help text for the groups in the order that they
	// were added so it is important that Deprecated is last.
	cmd.AddGroup(groups.Commands, groups.Config, groups.State)
	cmd.AddCommand(newVersionCommand())
	cmd.AddCommand(NewApplyCommand(options.ApplyOptions))
	cmd.AddCommand(unstable.NewCommand())
	cmd.AddCommand(ruleevaluations.NewCommand())
	configv1.AddCommandsTo(cmd)

	return cmd, nil
}
