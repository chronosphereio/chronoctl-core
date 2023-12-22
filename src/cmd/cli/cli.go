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

// Package cli provides utilities for use in the chronoctl CLI
package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// BuildCommandString rebuilds the supplied command into a string, flags included. ignoreFlags filters out
// any flags specified. For example, we can rebuild `chronoctl get dashboards --bucket my-bucket` and
// filter the api-token flag.
func BuildCommandString(cmd *cobra.Command, ignoreFlags []string) string {
	ignoreFlagsMap := make(map[string]struct{}, len(ignoreFlags))
	for _, flag := range ignoreFlags {
		ignoreFlagsMap[flag] = struct{}{}
	}
	name := cmd.Name()
	parents := []string{}
	// rebuild the parent command chain, e.g. the parent may be "preview", which has the parent "chronoctl"
	for parent := cmd.Parent(); parent != nil; parent = parent.Parent() {
		parents = append([]string{parent.Name()}, parents...)
	}
	flags := []string{}
	cmd.Flags().Visit(func(f *pflag.Flag) {
		if _, ok := ignoreFlagsMap[f.Name]; ok {
			return
		}
		value := f.Value.String()
		//turn a slice value into a comma seperated list.
		if slice, ok := f.Value.(pflag.SliceValue); ok {
			value = strings.Join(slice.GetSlice(), ",")
		}
		flags = append(flags, fmt.Sprintf("--%s %v", f.Name, value))
	})
	return fmt.Sprintf("%s %s %s", strings.Join(parents, " "), name, strings.Join(flags, " "))
}
