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

// Package unstable implements the unstable cobra command
package unstable

import (
	"github.com/chronosphereio/chronoctl-core/src/generated/cli/configunstable"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra command for the unstable namespace of commands
func NewCommand() *cobra.Command {
	root := &cobra.Command{
		Use:    "unstable",
		Hidden: true,
	}
	root.AddGroup(&cobra.Group{Title: "Configuration Entities", ID: "config"})

	configunstable.AddCommandsTo(root)
	return root
}
