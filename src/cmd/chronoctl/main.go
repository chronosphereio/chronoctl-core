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

// package main
package main

import (
	"os"

	"github.com/chronosphereio/chronoctl-core/src/cmd"
)

func main() {
	var opts cmd.Options
	cmd, err := cmd.New(opts)
	if err != nil {
		os.Stderr.WriteString(err.Error()) //nolint:errcheck
	}
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
