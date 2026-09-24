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

const testDashboardPath = "testdata/dashboard.json"

func readTestDashboard(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile(testDashboardPath)
	require.NoError(t, err)
	return string(b)
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

func TestValidateValidate(t *testing.T) {
	tests := []struct {
		name    string
		opts    func(t *testing.T, o *validateOptions)
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
			assert.Equal(t, readTestDashboard(t), o.dashboardJSON)
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
