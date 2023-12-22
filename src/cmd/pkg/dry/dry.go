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

// Package dry implements the dry run flag for chronoctl commands.
package dry

import (
	"github.com/spf13/cobra"
)

// Flags is a struct that contains a dry run flag that is used by CLI commands
// to get user input
type Flags struct {
	DryRun bool
}

// NewFlags returns a new DryRunFlags struct
func NewFlags() *Flags {
	return &Flags{}
}

// AddFlags adds the dry run flag to a cobra command
func (d *Flags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&d.DryRun, "dry-run", "d", false, "If true, directs the command to only perform a dry run.")

	cmd.Flags().BoolVar(&d.DryRun, "dryrun", false, "")
	cmd.Flags().MarkHidden("dryrun") //nolint:errcheck
}
