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
