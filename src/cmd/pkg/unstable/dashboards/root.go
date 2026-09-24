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

// Package dashboards contains hand-written commands that augment the
// generated `dashboards` command tree under `unstable`.
package dashboards

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AddCommandsTo attaches the hand-written dashboards commands to the
// generated `dashboards` subcommand of root.
func AddCommandsTo(root *cobra.Command) error {
	for _, c := range root.Commands() {
		if c.Name() == "dashboards" {
			c.AddCommand(newValidateCommand())
			return nil
		}
	}
	return fmt.Errorf("could not find %q command to attach hand-written dashboards commands to", "dashboards")
}
