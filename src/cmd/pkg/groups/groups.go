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

// Groups is a package for shared CLI groups.
package groups

import "github.com/spf13/cobra"

var (
	// Commands group is a catch all group for commands.
	Commands = &cobra.Group{
		Title: "Commands",
		ID:    "commands",
	}
	// Config group is for commands that interact with the config API
	Config = &cobra.Group{
		Title: "Chronosphere Configuration",
		ID:    "config",
	}
	// State group is for commands that interact with the state API
	State = &cobra.Group{
		Title: "Chronosphere State",
		ID:    "state",
	}
	// Deprecated group is for commands that are deprecated
	Deprecated = &cobra.Group{
		Title: "Deprecated Commands",
		ID:    "deprecated",
	}
)
