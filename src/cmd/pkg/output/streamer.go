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

package output

import (
	"fmt"
	"io"
	"strings"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/sliceutil"
	"github.com/chronosphereio/chronoctl-core/src/types"
	"github.com/spf13/cobra"

	"github.com/fatih/color"
)

// DeprecatedTextColor is the color used for deprecation messages.
var DeprecatedTextColor = color.New(color.FgRed)

// Streamer is a callback for processing an individual entity.
// If an error is returned, the caller should stop processing results.
type Streamer[T types.Object] func(T) error

// NewWriteObjectStreamer streams objects to the io.Writer provided.
func NewWriteObjectStreamer[T types.Object](manager *WriterManager) Streamer[T] {
	return func(e T) error {
		if err := manager.WriteObject(e); err != nil {
			return err
		}
		return nil
	}
}

// ColorPrinter prints to a given io.Writer with color.
type ColorPrinter struct {
	w io.Writer
	c *color.Color
}

// NewColorPrinter returns a printer that ensures all output is the given color.
func NewColorPrinter(w io.Writer, c *color.Color) ColorPrinter {
	return ColorPrinter{
		w: w,
		c: c,
	}
}

// Println prints a line.
func (c ColorPrinter) Println(a ...any) (n int, err error) {
	return c.c.Fprintln(c.w, a...)
}

// Printf prints a formatted string.
func (c ColorPrinter) Printf(format string, a ...any) (n int, err error) {
	return c.c.Fprintf(c.w, format, a...)
}

// NewStderrPrinter returns a printer that ensures all output
// is the color blue.
func NewStderrPrinter(cmd *cobra.Command) ColorPrinter {
	stderr := cmd.ErrOrStderr()
	c := color.New(color.FgBlue)
	return ColorPrinter{
		w: stderr,
		c: c,
	}
}

// Deprecated prints a deprecation message to the given command's stderr
// with a list of alternative commands.
func Deprecated(cmd *cobra.Command, alternatives ...string) {
	p := ColorPrinter{
		w: cmd.ErrOrStderr(),
		c: DeprecatedTextColor,
	}

	altCommands := sliceutil.Map(alternatives, func(s string) string {
		return fmt.Sprintf("`chronoctl %s`", s)
	})
	altStr := strings.Join(altCommands, ", ")

	if len(altCommands) > 1 {
		altStr = "one of: " + altStr
	}

	p.Printf("Command `%s` is deprecated, use %s\n", cmd.CommandPath(), altStr) //nolint:errcheck
}
