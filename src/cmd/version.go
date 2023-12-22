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
	"io"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/buildinfo"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/output"
)

type versionInfo struct {
	Version   string `json:"Version"`
	Revision  string `json:"Revision"`
	BuildDate string `json:"BuildDate"`
	GoVersion string `json:"GoVersion"`
}

type versionOptions struct {
	outputFlags *output.Flags
}

func newVersionOptions() *versionOptions {
	return &versionOptions{
		outputFlags: output.NewFlags(),
	}
}

func newVersionCommand() *cobra.Command {
	o := newVersionOptions()

	cmd := &cobra.Command{
		Use:     "version",
		Short:   "Print the current version of chronoctl",
		Long:    "Print the current version of chronoctl.",
		Example: "chronoctl version",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.validate(); err != nil {
				return err
			}

			info := versionInfo{
				Version:   buildinfo.Version,
				Revision:  buildinfo.SHA,
				BuildDate: buildinfo.Date,
				GoVersion: runtime.Version(),
			}
			if err := o.run(info, cmd.OutOrStdout()); err != nil {
				return err
			}

			return nil
		},
	}

	o.outputFlags.AddFlags(cmd)

	return cmd
}

func (o *versionOptions) validate() error {
	return o.outputFlags.Validate()
}

func (o *versionOptions) run(v versionInfo, w io.Writer) error {
	return o.outputFlags.WriteObject(v, w)
}
