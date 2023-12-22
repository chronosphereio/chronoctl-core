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

import "github.com/spf13/cobra"

// StrictModeFlags indicated whether or not to allow permissive parsing.
type StrictModeFlags struct {
	PermissiveParsing bool
}

// NewStrictModeFlags returns a new StrictModeFlags set.
func NewStrictModeFlags() *StrictModeFlags {
	return &StrictModeFlags{}
}

// AddFlags registers the strict mode vars onto a command.
func (c *StrictModeFlags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&c.PermissiveParsing, "no-strict", false, "If set, manifests with unknown fields are not rejected. Defaults to false.")
}
