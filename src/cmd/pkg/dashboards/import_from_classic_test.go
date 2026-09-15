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

package dashboards

import (
	"bytes"
	"errors"
	"os"
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

func TestImportFromClassicRun(t *testing.T) {
	tests := []struct {
		name       string
		opts       func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions
		wantStdout string
		wantStderr func(t *testing.T, stderr string)
		wantErr    string
	}{
		{
			name: "happy path",
			opts: func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions {
				classicJSON := readTestClassicDashboard(t)
				o := newImportFromClassicOptions()
				o.classicDashboardJSON = classicJSON
				o.collectionSlug = "my-collection"
				o.name = "API Latency Overview"
				o.slug = "api-latency-overview"
				o.updateIfExists = true

				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ImportDashboardFromClassic(gomock.Any()).DoAndReturn(
					func(params *config_v1.ImportDashboardFromClassicParams, _ ...config_v1.ClientOption) (*config_v1.ImportDashboardFromClassicOK, error) {
						require.Equal(t, classicJSON, params.Body.ClassicDashboardJSON)
						require.Equal(t, "my-collection", params.Body.CollectionSlug)
						require.Equal(t, "API Latency Overview", params.Body.Name)
						require.Equal(t, "api-latency-overview", params.Body.Slug)
						require.False(t, params.Body.DryRun)
						require.True(t, params.Body.UpdateIfExists)
						return &config_v1.ImportDashboardFromClassicOK{
							Payload: &models.Configv1ImportDashboardFromClassicResponse{
								Dashboard: &models.Configv1Dashboard{
									Name:           "API Latency Overview",
									Slug:           "api-latency-overview",
									CollectionSlug: "my-collection",
								},
							},
						}, nil
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
				assert.Contains(t, stderr, `Dashboard with slug "api-latency-overview" imported successfully`)
			},
		},
		{
			name: "dry run",
			opts: func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions {
				o := newImportFromClassicOptions()
				o.classicDashboardJSON = readTestClassicDashboard(t)
				o.collectionSlug = "my-collection"
				o.dryRunFlags.DryRun = true

				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ImportDashboardFromClassic(gomock.Any()).DoAndReturn(
					func(params *config_v1.ImportDashboardFromClassicParams, _ ...config_v1.ClientOption) (*config_v1.ImportDashboardFromClassicOK, error) {
						require.True(t, params.Body.DryRun)
						return &config_v1.ImportDashboardFromClassicOK{
							Payload: &models.Configv1ImportDashboardFromClassicResponse{
								Dashboard: &models.Configv1Dashboard{
									Name:           "API Latency Overview",
									Slug:           "api-latency-overview",
									CollectionSlug: "my-collection",
								},
							},
						}, nil
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
				assert.Contains(t, stderr, "--dry-run is set")
				assert.Contains(t, stderr, "Dashboard is valid and can be imported")
			},
		},
		{
			name: "unsupported features",
			opts: func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions {
				o := newImportFromClassicOptions()
				o.classicDashboardJSON = readTestClassicDashboard(t)
				o.collectionSlug = "my-collection"

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
				assert.Contains(t, stderr, "[LEVEL_UNSUPPORTED] unknown-variable: variable type is not supported (3 occurrences; panels: a, b)")
				assert.Contains(t, stderr, "[LEVEL_UNSUPPORTED] legacy-alert: panel-level alerting is not carried over\n")
			},
		},
		{
			name: "api error",
			opts: func(t *testing.T, ctrl *gomock.Controller) *importFromClassicOptions {
				o := newImportFromClassicOptions()
				o.classicDashboardJSON = readTestClassicDashboard(t)
				o.collectionSlug = "my-collection"

				cli := mocks.NewMockClientService(ctrl)
				cli.EXPECT().ImportDashboardFromClassic(gomock.Any()).Return(nil, errors.New("boom"))
				o.client = cli
				return o
			},
			wantErr: "boom",
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
		name    string
		opts    func(*testing.T, *importFromClassicOptions)
		wantErr string
	}{
		{
			name: "missing collection slug",
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
			name: "missing filename",
			opts: func(t *testing.T, o *importFromClassicOptions) {
				o.collectionSlug = "my-collection"
			},
			wantErr: "--filename is required",
		},
		{
			name: "unreadable file",
			opts: func(t *testing.T, o *importFromClassicOptions) {
				o.collectionSlug = "my-collection"
				o.fileFlags.Filename = "testdata/does_not_exist.json"
			},
			wantErr: "could not open file",
		},
		{
			name: "happy path",
			opts: func(t *testing.T, o *importFromClassicOptions) {
				o.collectionSlug = "my-collection"
				o.fileFlags.Filename = testClassicDashboardPath
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
			assert.Equal(t, readTestClassicDashboard(t), o.classicDashboardJSON)
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

		for _, flagName := range []string{"filename", "collection-slug", "name", "slug", "update-if-exists", "dry-run", "output"} {
			assert.NotNil(t, importCmd.Flags().Lookup(flagName), "expected flag %q", flagName)
		}
	})

	t.Run("errors when dashboards command is missing", func(t *testing.T) {
		root := &cobra.Command{Use: "chronoctl"}
		require.Error(t, AddCommandsTo(root))
	})
}
