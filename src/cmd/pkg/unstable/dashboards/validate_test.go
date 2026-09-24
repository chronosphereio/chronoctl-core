package dashboards

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	config_unstable "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/mocks"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/models"
)

const (
	testDashboardPath         = "testdata/dashboard.json"
	testManifestPath          = "testdata/dashboard_manifest.yml"
	testManifestDashboardJSON = `{"kind":"Dashboard","metadata":{"name":"api-latency-overview"},"spec":{"duration":"1h","panels":{},"layouts":[]}}`
)

func readTestDashboard(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile(testDashboardPath)
	require.NoError(t, err)
	return string(b)
}

func readFile(t *testing.T, path string) []byte {
	t.Helper()
	b, err := os.ReadFile(path)
	require.NoError(t, err)
	return b
}

func TestValidateRun(t *testing.T) {
	tests := []struct {
		name       string
		opts       func(t *testing.T, ctrl *gomock.Controller) *validateOptions
		wantStdout string
		wantErr    string
	}{
		{
			name: "happy path",
			opts: func(t *testing.T, ctrl *gomock.Controller) *validateOptions {
				dashboardJSON := readTestDashboard(t)
				o := newValidateOptions()
				o.dashboardJSON = dashboardJSON

				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ValidateDashboard(gomock.Any()).DoAndReturn(
					func(params *config_unstable.ValidateDashboardParams, _ ...config_unstable.ClientOption) (*config_unstable.ValidateDashboardOK, error) {
						require.Equal(t, dashboardJSON, params.Body.DashboardJSON)
						return &config_unstable.ValidateDashboardOK{}, nil
					})
				o.client = cli
				return o
			},
			wantStdout: "Dashboard is valid\n",
		},
		{
			name: "api error",
			opts: func(t *testing.T, ctrl *gomock.Controller) *validateOptions {
				o := newValidateOptions()
				o.dashboardJSON = readTestDashboard(t)

				resp := config_unstable.NewValidateDashboardDefault(400)
				resp.Payload = &models.APIError{Code: 400, Message: "dashboard_json: spec.panels: required"}

				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ValidateDashboard(gomock.Any()).Return(nil, resp)
				o.client = cli
				return o
			},
			wantErr: "spec.panels: required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var stdout bytes.Buffer
			o := tt.opts(t, ctrl)

			err := o.run(&stdout)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "HTTP code 400")
				assert.Contains(t, err.Error(), tt.wantErr)
				assert.Empty(t, stdout.String())
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantStdout, stdout.String())
		})
	}
}

func TestDashboardJSONFromInput(t *testing.T) {
	tests := []struct {
		name    string
		input   func(t *testing.T) []byte
		want    string
		wantErr string
	}{
		{
			name:  "raw json",
			input: func(t *testing.T) []byte { return readFile(t, testDashboardPath) },
			want:  readTestDashboard(t),
		},
		{
			name:  "raw json with surrounding whitespace",
			input: func(t *testing.T) []byte { return []byte("\n  {\"kind\":\"Dashboard\"}\n") },
			want:  "\n  {\"kind\":\"Dashboard\"}\n",
		},
		{
			name:  "yaml manifest",
			input: func(t *testing.T) []byte { return readFile(t, testManifestPath) },
			want:  testManifestDashboardJSON,
		},
		{
			name: "json manifest for the unstable api",
			input: func(t *testing.T) []byte {
				return []byte(`{"api_version":"unstable/config","kind":"Dashboard","spec":{"slug":"x","dashboard_json":"{\"kind\":\"Dashboard\"}"}}`)
			},
			want: `{"kind":"Dashboard"}`,
		},
		{
			name: "manifest of another kind",
			input: func(t *testing.T) []byte {
				return []byte("api_version: v1/config\nkind: Monitor\nspec:\n  slug: m\n")
			},
			wantErr: "expected a Dashboard manifest",
		},
		{
			name: "manifest with unknown api version",
			input: func(t *testing.T) []byte {
				return []byte("api_version: v2/config\nkind: Dashboard\n")
			},
			wantErr: "no registered type",
		},
		{
			name: "manifest without dashboard_json",
			input: func(t *testing.T) []byte {
				return []byte("api_version: v1/config\nkind: Dashboard\nspec:\n  slug: x\n")
			},
			wantErr: "no spec.dashboard_json",
		},
		{
			name: "manifest with unknown field",
			input: func(t *testing.T) []byte {
				return []byte("api_version: v1/config\nkind: Dashboard\nspec:\n  dashbord_json: '{}'\n")
			},
			wantErr: "dashbord_json",
		},
		{
			name: "multiple manifests",
			input: func(t *testing.T) []byte {
				return []byte("api_version: v1/config\nkind: Dashboard\nspec:\n  dashboard_json: '{}'\n---\napi_version: v1/config\nkind: Dashboard\nspec:\n  dashboard_json: '{}'\n")
			},
			wantErr: "more than one item",
		},
		{
			name:    "yaml that is not a manifest",
			input:   func(t *testing.T) []byte { return []byte("foo: bar\n") },
			wantErr: "not valid JSON",
		},
		{
			name:    "invalid json",
			input:   func(t *testing.T) []byte { return []byte("{\"kind\": ") },
			wantErr: "not valid JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dashboardJSONFromInput(tt.input(t))
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestValidateValidate(t *testing.T) {
	tests := []struct {
		name    string
		opts    func(t *testing.T, o *validateOptions)
		want    string
		wantErr string
	}{
		{
			name:    "missing filename",
			opts:    func(t *testing.T, o *validateOptions) {},
			wantErr: "--filename is required",
		},
		{
			name: "unreadable file",
			opts: func(t *testing.T, o *validateOptions) {
				o.fileFlags.Filename = "testdata/does_not_exist.json"
			},
			wantErr: "could not open file",
		},
		{
			name: "empty file",
			opts: func(t *testing.T, o *validateOptions) {
				emptyFile := filepath.Join(t.TempDir(), "empty.json")
				require.NoError(t, os.WriteFile(emptyFile, nil, 0o600))
				o.fileFlags.Filename = emptyFile
			},
			wantErr: "dashboard file is empty",
		},
		{
			name: "happy path",
			opts: func(t *testing.T, o *validateOptions) {
				o.fileFlags.Filename = testDashboardPath
			},
			want: readTestDashboard(t),
		},
		{
			name: "manifest file",
			opts: func(t *testing.T, o *validateOptions) {
				o.fileFlags.Filename = testManifestPath
			},
			want: testManifestDashboardJSON,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			o := newValidateOptions()
			o.client = mocks.NewMockClientService(ctrl)
			tt.opts(t, o)

			err := o.validate()
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, o.dashboardJSON)
		})
	}
}

func TestAddCommandsTo(t *testing.T) {
	t.Run("attaches to the dashboards command", func(t *testing.T) {
		root := &cobra.Command{Use: "chronoctl"}
		root.AddCommand(&cobra.Command{Use: "dashboards"})

		require.NoError(t, AddCommandsTo(root))

		var dashboardsCmd *cobra.Command
		for _, c := range root.Commands() {
			if c.Name() == "dashboards" {
				dashboardsCmd = c
			}
		}
		require.NotNil(t, dashboardsCmd)

		var validateCmd *cobra.Command
		for _, c := range dashboardsCmd.Commands() {
			if c.Name() == "validate" {
				validateCmd = c
			}
		}
		require.NotNil(t, validateCmd)

		assert.NotNil(t, validateCmd.Flags().Lookup("filename"))
		assert.NotNil(t, validateCmd.Flags().ShorthandLookup("f"))
	})

	t.Run("errors when dashboards command is missing", func(t *testing.T) {
		root := &cobra.Command{Use: "chronoctl"}
		require.Error(t, AddCommandsTo(root))
	})
}
