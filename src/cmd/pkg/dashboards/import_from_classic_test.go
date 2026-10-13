package dashboards

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	config_v1 "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/mocks"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
)

const testClassicDashboardPath = "testdata/classic_dashboard.json"

func readTestClassicDashboard(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile(testClassicDashboardPath)
	require.NoError(t, err)
	return string(b)
}

// writeTempInput writes content to a file in a test temp dir and returns its
// path, for driving validate() the way -f does.
func writeTempInput(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "input")
	require.NoError(t, os.WriteFile(path, []byte(content), 0o600))
	return path
}

func importOK(dashboard *models.Configv1Dashboard) *config_v1.ImportDashboardFromClassicOK {
	return &config_v1.ImportDashboardFromClassicOK{
		Payload: &models.Configv1ImportDashboardFromClassicResponse{Dashboard: dashboard},
	}
}

func TestImportFromClassicRun(t *testing.T) {
	tests := []struct {
		name       string
		opts       func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions
		wantStdout string
		wantStderr func(t *testing.T, stderr string)
		wantErr    string
	}{
		{
			name: "request body",
			opts: func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions {
				body := &models.Configv1ImportDashboardFromClassicRequest{
					ClassicDashboardJSON: readTestClassicDashboard(t),
					CollectionSlug:       "my-collection",
					Name:                 "API Latency Overview",
					Slug:                 "api-latency-overview",
					UpdateIfExists:       true,
				}
				o := newImportFromClassicOptions()
				o.request = &importRequest{source: "-f grafana.json", body: body}

				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ImportDashboardFromClassic(gomock.Any()).DoAndReturn(
					func(params *config_v1.ImportDashboardFromClassicParams, _ ...config_v1.ClientOption) (*config_v1.ImportDashboardFromClassicOK, error) {
						require.Equal(t, body, params.Body)
						return importOK(&models.Configv1Dashboard{
							Name:           "API Latency Overview",
							Slug:           "api-latency-overview",
							CollectionSlug: "my-collection",
						}), nil
					})
				o.client = cli
				return o
			},
			wantStdout: `api_version: v1/config
kind: Dashboard
spec:
  slug: api-latency-overview
  name: API Latency Overview
  collection_slug: my-collection
`,
			wantStderr: func(t *testing.T, stderr string) {
				assert.Equal(t, "Imported -f grafana.json as Dashboard \"api-latency-overview\"\n", stderr)
			},
		},
		{
			name: "dry run",
			opts: func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions {
				o := newImportFromClassicOptions()
				o.request = &importRequest{
					source: "-f grafana.json",
					body: &models.Configv1ImportDashboardFromClassicRequest{
						ClassicDashboardJSON: readTestClassicDashboard(t),
						CollectionSlug:       "my-collection",
						DryRun:               true,
					},
				}

				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ImportDashboardFromClassic(gomock.Any()).DoAndReturn(
					func(params *config_v1.ImportDashboardFromClassicParams, _ ...config_v1.ClientOption) (*config_v1.ImportDashboardFromClassicOK, error) {
						require.True(t, params.Body.DryRun)
						return importOK(&models.Configv1Dashboard{
							Name:           "API Latency Overview",
							Slug:           "api-latency-overview",
							CollectionSlug: "my-collection",
						}), nil
					})
				o.client = cli
				return o
			},
			wantStdout: `api_version: v1/config
kind: Dashboard
spec:
  slug: api-latency-overview
  name: API Latency Overview
  collection_slug: my-collection
`,
			wantStderr: func(t *testing.T, stderr string) {
				assert.Equal(t, "[dry run] Imported -f grafana.json as Dashboard \"api-latency-overview\"\n", stderr)
			},
		},
		{
			name: "unsupported features",
			opts: func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions {
				o := newImportFromClassicOptions()
				o.request = &importRequest{
					source: "-f grafana.json",
					body: &models.Configv1ImportDashboardFromClassicRequest{
						ClassicDashboardJSON: readTestClassicDashboard(t),
						CollectionSlug:       "my-collection",
					},
				}

				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ImportDashboardFromClassic(gomock.Any()).Return(&config_v1.ImportDashboardFromClassicOK{
					Payload: &models.Configv1ImportDashboardFromClassicResponse{
						Dashboard: &models.Configv1Dashboard{
							Name: "API Latency Overview",
							Slug: "api-latency-overview",
						},
						UnsupportedFeatures: []*models.ImportDashboardFromClassicResponseUnsupportedFeature{
							{
								Kind:        "unknown-variable",
								Message:     "variable type is not supported",
								Level:       models.UnsupportedFeatureLevelLEVELUNSUPPORTED,
								Occurrences: 3,
								PanelKeys:   []string{"a", "b"},
							},
							{
								Kind:    "legacy-alert",
								Message: "panel-level alerting is not carried over",
								Level:   models.UnsupportedFeatureLevelLEVELUNSUPPORTED,
							},
						},
					},
				}, nil)
				o.client = cli
				return o
			},
			wantStdout: `api_version: v1/config
kind: Dashboard
spec:
  slug: api-latency-overview
  name: API Latency Overview
`,
			wantStderr: func(t *testing.T, stderr string) {
				assert.Contains(t, stderr, "Imported -f grafana.json as Dashboard \"api-latency-overview\"\n")
				assert.Contains(t, stderr, "[LEVEL_UNSUPPORTED] unknown-variable: variable type is not supported (3 occurrences; panels: a, b)")
				assert.Contains(t, stderr, "[LEVEL_UNSUPPORTED] legacy-alert: panel-level alerting is not carried over\n")
			},
		},
		{
			name: "api error",
			opts: func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions {
				o := newImportFromClassicOptions()
				o.request = &importRequest{
					source: "-f grafana.json",
					body: &models.Configv1ImportDashboardFromClassicRequest{
						ClassicDashboardJSON: readTestClassicDashboard(t),
						CollectionSlug:       "my-collection",
					},
				}

				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ImportDashboardFromClassic(gomock.Any()).Return(nil, errors.New("boom"))
				o.client = cli
				return o
			},
			wantErr: "boom",
		},
		{
			name: "wrapped API error",
			opts: func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions {
				o := newImportFromClassicOptions()
				o.request = &importRequest{
					source: `ClassicDashboard "xyz"`,
					body: &models.Configv1ImportDashboardFromClassicRequest{
						ClassicDashboardJSON: readTestClassicDashboard(t),
						CollectionSlug:       "my-collection",
						Slug:                 "xyz",
					},
				}

				apiErr := config_v1.NewImportDashboardFromClassicDefault(409)
				apiErr.Payload = &models.APIError{Message: `Classic dashboard with slug "xyz" already exists`}
				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ImportDashboardFromClassic(gomock.Any()).Return(nil, apiErr)
				o.client = cli
				return o
			},
			wantErr: `POST /api/v1/config/dashboards:importFromClassic failed with HTTP code 409: Classic dashboard with slug "xyz" already exists`,
		},
		{
			name: "no dashboard in response",
			opts: func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions {
				o := newImportFromClassicOptions()
				o.request = &importRequest{
					source: "-f grafana.json",
					body: &models.Configv1ImportDashboardFromClassicRequest{
						ClassicDashboardJSON: readTestClassicDashboard(t),
						CollectionSlug:       "my-collection",
					},
				}

				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ImportDashboardFromClassic(gomock.Any()).Return(importOK(nil), nil)
				o.client = cli
				return o
			},
			wantErr: "API returned no dashboard",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var stdout, stderr bytes.Buffer
			o := tt.opts(t, ctrl)
			o.stderr = &stderr

			err := o.run(&stdout)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				assert.Empty(t, stdout.String())
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantStdout, stdout.String())
			if tt.wantStderr != nil {
				tt.wantStderr(t, stderr.String())
			}
		})
	}
}

// TestImportFromClassicManifest drives validate() and run() together, as the
// command does, with a manifest read from -f.
func TestImportFromClassicManifest(t *testing.T) {
	const classicManifest = `api_version: v1/config
kind: ClassicDashboard
spec:
  slug: xyz-native
  name: API Latency Overview
  created_at: "2026-01-02T03:04:05.000Z"
  updated_at: "2026-01-02T03:04:05.000Z"
  collection_slug: team-a
  dashboard_json: '{"title":"API Latency Overview","panels":[]}'
`
	const wantStdout = `api_version: v1/config
kind: Dashboard
spec:
  slug: xyz-native
  name: API Latency Overview
  collection_slug: team-a
`

	tests := []struct {
		name       string
		input      string
		dryRun     bool
		mock       func(t *testing.T, cli *mocks.MockClientService)
		wantStdout string
		wantStderr string
		wantErr    string
	}{
		{
			name:  "manifest maps to the request",
			input: classicManifest,
			mock: func(t *testing.T, cli *mocks.MockClientService) {
				cli.EXPECT().ImportDashboardFromClassic(gomock.Any()).DoAndReturn(
					func(params *config_v1.ImportDashboardFromClassicParams, _ ...config_v1.ClientOption) (*config_v1.ImportDashboardFromClassicOK, error) {
						require.Equal(t, &models.Configv1ImportDashboardFromClassicRequest{
							ClassicDashboardJSON: `{"title":"API Latency Overview","panels":[]}`,
							CollectionSlug:       "team-a",
							Name:                 "API Latency Overview",
							Slug:                 "xyz-native",
						}, params.Body)
						return importOK(&models.Configv1Dashboard{
							Name:           "API Latency Overview",
							Slug:           "xyz-native",
							CollectionSlug: "team-a",
						}), nil
					})
			},
			wantStdout: wantStdout,
			wantStderr: "Imported ClassicDashboard \"xyz-native\" as Dashboard \"xyz-native\"\n",
		},
		{
			name:   "dry run",
			input:  classicManifest,
			dryRun: true,
			mock: func(t *testing.T, cli *mocks.MockClientService) {
				cli.EXPECT().ImportDashboardFromClassic(gomock.Any()).DoAndReturn(
					func(params *config_v1.ImportDashboardFromClassicParams, _ ...config_v1.ClientOption) (*config_v1.ImportDashboardFromClassicOK, error) {
						require.True(t, params.Body.DryRun)
						return importOK(&models.Configv1Dashboard{
							Name:           "API Latency Overview",
							Slug:           "xyz-native",
							CollectionSlug: "team-a",
						}), nil
					})
			},
			wantStdout: wantStdout,
			wantStderr: "[dry run] Imported ClassicDashboard \"xyz-native\" as Dashboard \"xyz-native\"\n",
		},
		{
			name:    "multi-document input makes no API call",
			input:   classicManifest + "---\n" + classicManifest,
			mock:    func(t *testing.T, cli *mocks.MockClientService) {},
			wantErr: "input contains more than one item",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			cli := mocks.NewMockClientService(ctrl)
			tt.mock(t, cli)

			var stdout, stderr bytes.Buffer
			o := newImportFromClassicOptions()
			o.client = cli
			o.stderr = &stderr
			o.fileFlags.Filename = writeTempInput(t, tt.input)
			o.dryRunFlags.DryRun = tt.dryRun

			err := o.validate()
			if err == nil {
				err = o.run(&stdout)
			}
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				assert.Empty(t, stdout.String())
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantStdout, stdout.String())
			assert.Equal(t, tt.wantStderr, stderr.String())
		})
	}
}

func TestFormatUnsupportedFeature(t *testing.T) {
	tests := []struct {
		name     string
		feature  *models.ImportDashboardFromClassicResponseUnsupportedFeature
		expected string
	}{
		{
			name: "no occurrences or panels",
			feature: &models.ImportDashboardFromClassicResponseUnsupportedFeature{
				Kind:    "unknown-variable",
				Message: "variable type is not supported",
				Level:   models.UnsupportedFeatureLevelLEVELUNSUPPORTED,
			},
			expected: "[LEVEL_UNSUPPORTED] unknown-variable: variable type is not supported",
		},
		{
			name: "singular occurrence",
			feature: &models.ImportDashboardFromClassicResponseUnsupportedFeature{
				Kind:        "unknown-variable",
				Message:     "variable type is not supported",
				Level:       models.UnsupportedFeatureLevelLEVELUNSUPPORTED,
				Occurrences: 1,
			},
			expected: "[LEVEL_UNSUPPORTED] unknown-variable: variable type is not supported (1 occurrence)",
		},
		{
			name: "plural occurrences",
			feature: &models.ImportDashboardFromClassicResponseUnsupportedFeature{
				Kind:        "unknown-variable",
				Message:     "variable type is not supported",
				Level:       models.UnsupportedFeatureLevelLEVELUNSUPPORTED,
				Occurrences: 3,
			},
			expected: "[LEVEL_UNSUPPORTED] unknown-variable: variable type is not supported (3 occurrences)",
		},
		{
			name: "panels only",
			feature: &models.ImportDashboardFromClassicResponseUnsupportedFeature{
				Kind:      "unknown-variable",
				Message:   "variable type is not supported",
				Level:     models.UnsupportedFeatureLevelLEVELUNSUPPORTED,
				PanelKeys: []string{"a", "b"},
			},
			expected: "[LEVEL_UNSUPPORTED] unknown-variable: variable type is not supported (panels: a, b)",
		},
		{
			name: "occurrences and panels",
			feature: &models.ImportDashboardFromClassicResponseUnsupportedFeature{
				Kind:        "unknown-variable",
				Message:     "variable type is not supported",
				Level:       models.UnsupportedFeatureLevelLEVELUNSUPPORTED,
				Occurrences: 3,
				PanelKeys:   []string{"a", "b"},
			},
			expected: "[LEVEL_UNSUPPORTED] unknown-variable: variable type is not supported (3 occurrences; panels: a, b)",
		},
		{
			name: "with detail differing from message",
			feature: &models.ImportDashboardFromClassicResponseUnsupportedFeature{
				Kind:    "unknown-variable",
				Message: "variable type is not supported",
				Detail:  "variable \"foo\" uses type \"custom\"",
				Level:   models.UnsupportedFeatureLevelLEVELUNSUPPORTED,
			},
			expected: `[LEVEL_UNSUPPORTED] unknown-variable: variable type is not supported -- variable "foo" uses type "custom"`,
		},
		{
			name: "detail equal to message is not duplicated",
			feature: &models.ImportDashboardFromClassicResponseUnsupportedFeature{
				Kind:    "unknown-variable",
				Message: "variable type is not supported",
				Detail:  "variable type is not supported",
				Level:   models.UnsupportedFeatureLevelLEVELUNSUPPORTED,
			},
			expected: "[LEVEL_UNSUPPORTED] unknown-variable: variable type is not supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, formatUnsupportedFeature(tt.feature))
		})
	}
}

func TestImportFromClassicValidate(t *testing.T) {
	tests := []struct {
		name        string
		opts        func(*testing.T, *importFromClassicOptions)
		wantErr     string
		wantRequest func(t *testing.T) *importRequest
	}{
		{
			name: "missing filename",
			opts: func(t *testing.T, o *importFromClassicOptions) {
				o.collectionSlug = "my-collection"
			},
			wantErr: "--filename is required",
		},
		{
			name: "empty input",
			opts: func(t *testing.T, o *importFromClassicOptions) {
				o.collectionSlug = "my-collection"
				o.fileFlags.Filename = writeTempInput(t, "")
			},
			wantErr: "input is empty",
		},
		{
			name: "missing collection slug for raw JSON",
			opts: func(t *testing.T, o *importFromClassicOptions) {
				o.fileFlags.Filename = testClassicDashboardPath
			},
			wantErr: "--collection-slug is required",
		},
		{
			name: "update-if-exists without slug",
			opts: func(t *testing.T, o *importFromClassicOptions) {
				o.fileFlags.Filename = testClassicDashboardPath
				o.collectionSlug = "my-collection"
				o.updateIfExists = true
			},
			wantErr: "--slug is required when --update-if-exists is set",
		},
		{
			name: "missing file",
			opts: func(t *testing.T, o *importFromClassicOptions) {
				o.collectionSlug = "my-collection"
				o.fileFlags.Filename = "testdata/does_not_exist.json"
			},
			wantErr: "could not open file",
		},
		{
			name: "unreadable input",
			opts: func(t *testing.T, o *importFromClassicOptions) {
				o.collectionSlug = "my-collection"
				o.fileFlags.Filename = t.TempDir()
			},
			wantErr: "could not read input",
		},
		{
			name: "happy path",
			opts: func(t *testing.T, o *importFromClassicOptions) {
				o.collectionSlug = "my-collection"
				o.name = "API Latency Overview"
				o.slug = "api-latency-overview"
				o.updateIfExists = true
				o.dryRunFlags.DryRun = true
				o.fileFlags.Filename = testClassicDashboardPath
			},
			wantRequest: func(t *testing.T) *importRequest {
				return &importRequest{
					source: "-f " + testClassicDashboardPath,
					body: &models.Configv1ImportDashboardFromClassicRequest{
						ClassicDashboardJSON: readTestClassicDashboard(t),
						CollectionSlug:       "my-collection",
						Name:                 "API Latency Overview",
						Slug:                 "api-latency-overview",
						DryRun:               true,
						UpdateIfExists:       true,
					},
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			o := newImportFromClassicOptions()
			o.client = mocks.NewMockClientService(ctrl)
			tt.opts(t, o)

			err := o.validate()
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantRequest(t), o.request)
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

		var importCmd *cobra.Command
		for _, c := range dashboardsCmd.Commands() {
			if c.Name() == "import-from-classic" {
				importCmd = c
			}
		}
		require.NotNil(t, importCmd)

		for _, flagName := range []string{"filename", "collection-slug", "name", "slug", "update-if-exists", "dry-run", "no-strict", "output"} {
			assert.NotNil(t, importCmd.Flags().Lookup(flagName), "expected flag %q", flagName)
		}
	})

	t.Run("errors when dashboards command is missing", func(t *testing.T) {
		root := &cobra.Command{Use: "chronoctl"}
		require.Error(t, AddCommandsTo(root))
	})
}
