package dashboards

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/clienterror"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/file"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/groups"
	"github.com/chronosphereio/chronoctl-core/src/generated/cli/configunstable"
	"github.com/chronosphereio/chronoctl-core/src/generated/cli/configv1"
	config_unstable "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/models"
	"github.com/chronosphereio/chronoctl-core/src/thirdparty/yaml"
	"github.com/chronosphereio/chronoctl-core/src/types"
)

func newValidateCommand() *cobra.Command {
	o := newValidateOptions()

	cmd := &cobra.Command{
		Use:     "validate -f <file>",
		GroupID: groups.Commands.ID,
		Short:   "Validates a dashboard without persisting it.",
		Long: `Validates a dashboard against the API without creating or updating anything.

-f accepts either of two inputs, or "-" to read one from stdin:

  * Raw dashboard JSON: the JSON a Dashboard's dashboard_json field holds.
  * A chronoctl Dashboard manifest in YAML or JSON, such as the output of
    "chronoctl dashboards read <slug>". Only its spec.dashboard_json is
    validated; the other manifest fields are not.

Input with a top-level api_version is treated as a manifest; anything else
is sent as raw dashboard JSON.

When the API accepts the dashboard the command prints "Dashboard is valid"
and exits 0. When it does not, the command exits non-zero and the error
carries the validation failures reported by the API.`,
		Example: `# Validate raw dashboard JSON.
chronoctl unstable dashboards validate -f dashboard.json

# Validate an edited copy of an existing dashboard.
chronoctl dashboards read my-dashboard > dashboard.yml
# ... edit dashboard.yml ...
chronoctl unstable dashboards validate -f dashboard.yml

# Validate from stdin.
cat dashboard.json | chronoctl unstable dashboards validate -f -`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.validate(); err != nil {
				return err
			}
			return o.run(cmd.OutOrStdout())
		},
	}

	o.addFlags(cmd)
	return cmd
}

type validateOptions struct {
	clientFlags *client.Flags
	fileFlags   *file.Flags

	dashboardJSON string

	client config_unstable.ClientService
}

func newValidateOptions() *validateOptions {
	return &validateOptions{
		clientFlags: client.NewClientFlags(),
		fileFlags:   file.NewFlags(true /* required */),
	}
}

func (o *validateOptions) addFlags(cmd *cobra.Command) {
	o.clientFlags.AddFlags(cmd)
	o.fileFlags.AddFlags(cmd)
}

func (o *validateOptions) validate() error {
	if o.fileFlags.Filename == "" {
		return errors.New("--filename is required")
	}
	f, err := o.fileFlags.File()
	if err != nil {
		return err
	}
	if f != os.Stdin {
		defer f.Close() //nolint:errcheck
	}
	content, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("could not read dashboard file: %v", err)
	}
	if len(content) == 0 {
		return errors.New("dashboard file is empty")
	}
	dashboardJSON, err := dashboardJSONFromInput(content)
	if err != nil {
		return err
	}
	o.dashboardJSON = dashboardJSON

	if o.client == nil {
		c, err := o.clientFlags.ConfigUnstableClient()
		if err != nil {
			return err
		}
		o.client = c
	}

	return nil
}

func (o *validateOptions) run(w io.Writer) error {
	ctx, cancel := context.WithTimeout(context.Background(), o.clientFlags.Timeout())
	defer cancel()

	_, err := o.client.ValidateDashboard(&config_unstable.ValidateDashboardParams{
		Context: ctx,
		Body: &models.ConfigunstableValidateDashboardRequest{
			DashboardJSON: o.dashboardJSON,
		},
	})
	if err != nil {
		return clienterror.Wrap(err)
	}

	_, err = fmt.Fprintln(w, "Dashboard is valid")
	return err
}

var jsonStart = regexp.MustCompile(`^\s*\{`)

// dashboardJSONFromInput returns the raw dashboard JSON to validate. Input
// with a top-level api_version is a chronoctl manifest and must hold a single
// Dashboard, whose spec.dashboard_json is returned. Anything else is raw
// dashboard JSON and is returned verbatim.
func dashboardJSONFromInput(content []byte) (string, error) {
	meta, err := peekTypeMeta(content)
	if err != nil {
		return "", err
	}

	if meta.APIVersion == "" {
		if !json.Valid(content) {
			return "", errors.New("input is not valid JSON; expected raw dashboard JSON or a chronoctl Dashboard manifest with api_version and kind")
		}
		return string(content), nil
	}

	obj, err := types.MustDecodeSingleObject[types.Object](bytes.NewReader(content), false /* permissiveParsing */)
	if err != nil {
		return "", err
	}

	var dashboardJSON string
	switch d := obj.(type) {
	case *configv1.Dashboard:
		if d.Spec != nil {
			dashboardJSON = d.Spec.DashboardJSON
		}
	case *configunstable.Dashboard:
		if d.Spec != nil {
			dashboardJSON = d.Spec.DashboardJSON
		}
	default:
		return "", fmt.Errorf("expected a Dashboard manifest, got %s", obj.Type())
	}
	if dashboardJSON == "" {
		return "", errors.New("manifest has no spec.dashboard_json")
	}
	return dashboardJSON, nil
}

// peekTypeMeta decodes only api_version and kind, choosing the decoder the
// same way types.Decode does. JSON that does not decode into a TypeMeta is
// still candidate raw dashboard JSON, so only YAML parse failures are errors.
func peekTypeMeta(content []byte) (types.TypeMeta, error) {
	var meta types.TypeMeta
	if jsonStart.Match(content) {
		_ = json.Unmarshal(content, &meta)
		return meta, nil
	}
	if err := yaml.Unmarshal(content, &meta); err != nil {
		return meta, fmt.Errorf("could not parse input as YAML: %w", err)
	}
	return meta, nil
}
