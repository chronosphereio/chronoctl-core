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

// package file contains commands related to file operations.
package file

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Flags contains the info needed for reading in a file for a command.
type Flags struct {
	Required bool
	Filename string
}

// NewFlags returns a new Flags
func NewFlags(required bool) *Flags {
	return &Flags{Required: required}
}

// AddFlags registers the flags on the supplied command.
func (f *Flags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Filename, "filename", "f", "", `Name of the file containing the configuration to apply. Supplying "-" will read from stdin.`)

	if f.Required {
		// An error here is only possible if we don't have a flag named `filename`, but we set it
		// above and so can ignore the error.
		cmd.MarkFlagRequired("filename") //nolint:errcheck, gosec
	}
}

// File returns an opened file based on the filename. If `-` is provided, we return stdin.
func (f *Flags) File() (*os.File, error) {
	if !f.Required && f.Filename == "" {
		return nil, nil
	}

	if f.Filename == "-" {
		return os.Stdin, nil
	}

	file, err := os.Open(f.Filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	return file, nil
}
