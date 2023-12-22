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

package specscanner

import (
	"fmt"
	"testing"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	swaggerSpec, err := loads.Spec("testdata/service.swagger.json")
	require.NoError(t, err)

	tests := []struct {
		name              string
		spec              []byte
		expectedPaths     []string
		expectedOps       []string
		expectedResponses []string
		expectedParams    []string
		wantErr           string
	}{
		{
			name:    "invalid spec",
			spec:    []byte("{invalid:json}"),
			wantErr: "invalid character 'i' looking for beginning of object key string",
		},
		{
			name: "valid spec",
			spec: swaggerSpec.Raw(),
			expectedPaths: []string{
				"/api/v1/config/monitors/{slug}",
				"/api/v1/config/teams",
				"/api/v1/config/teams/{slug}",
				"/api/v1/config/monitors",
			},
			expectedOps: []string{
				"GET /api/v1/config/monitors/{slug}",
				"PUT /api/v1/config/monitors/{slug}",
				"DELETE /api/v1/config/monitors/{slug}",
				"GET /api/v1/config/teams",
				"POST /api/v1/config/teams",
				"GET /api/v1/config/teams/{slug}",
				"PUT /api/v1/config/teams/{slug}",
				"DELETE /api/v1/config/teams/{slug}",
				"GET /api/v1/config/monitors",
				"POST /api/v1/config/monitors",
			},
			expectedResponses: []string{
				"200 A successful response.",
				"404 Cannot delete the Monitor because the slug does not exist.",
				"500 An unexpected error response.",
				"500 An unexpected error response.",
				"200 A successful response.",
				"404 Cannot read the Monitor because the slug does not exist.",
				"200 A successful response containing the updated Monitor.",
				"400 Cannot update the Monitor because the request is invalid.",
				"404 Cannot update the Monitor because the slug does not exist.",
				"409 Cannot update the Monitor because there is a conflict with an existing Monitor.",
				"500 An unexpected error response.",
				"500 An unexpected error response.",
				"200 A successful response.",
				"500 An unexpected error response.",
				"200 A successful response containing the created Team.",
				"400 Cannot create the Team because the request is invalid.",
				"409 Cannot create the Team because there is a conflict with an existing Team.",
				"200 A successful response containing the updated Team.",
				"400 Cannot update the Team because the request is invalid.",
				"404 Cannot update the Team because the slug does not exist.",
				"409 Cannot update the Team because there is a conflict with an existing Team.",
				"500 An unexpected error response.",
				"200 A successful response.",
				"404 Cannot delete the Team because the slug does not exist.",
				"500 An unexpected error response.",
				"200 A successful response.",
				"404 Cannot read the Team because the slug does not exist.",
				"500 An unexpected error response.",
				"500 An unexpected error response.",
				"200 A successful response.",
				"200 A successful response containing the created Monitor.",
				"400 Cannot create the Monitor because the request is invalid.",
				"409 Cannot create the Monitor because there is a conflict with an existing Monitor.",
				"500 An unexpected error response.",
			},
			expectedParams: []string{
				"GET /api/v1/config/teams/{slug}:slug",
				"PUT /api/v1/config/teams/{slug}:slug",
				"PUT /api/v1/config/teams/{slug}:body",
				"DELETE /api/v1/config/teams/{slug}:slug",
				"GET /api/v1/config/monitors:page.max_size",
				"GET /api/v1/config/monitors:page.token",
				"GET /api/v1/config/monitors:slugs",
				"GET /api/v1/config/monitors:bucket_slugs",
				"GET /api/v1/config/monitors:collection_slugs",
				"GET /api/v1/config/monitors:names",
				"GET /api/v1/config/monitors:team_slugs",
				"POST /api/v1/config/monitors:body",
				"PUT /api/v1/config/monitors/{slug}:slug",
				"PUT /api/v1/config/monitors/{slug}:body",
				"DELETE /api/v1/config/monitors/{slug}:slug",
				"GET /api/v1/config/monitors/{slug}:slug",
				"GET /api/v1/config/teams:page.max_size",
				"GET /api/v1/config/teams:page.token",
				"GET /api/v1/config/teams:slugs",
				"GET /api/v1/config/teams:names",
				"POST /api/v1/config/teams:body",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := &testHandler{}
			err := Scan(tt.spec, th, nil)
			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.True(t, th.scanStarted)
				assert.True(t, th.scanEnded)
				assert.Equal(t, th.pathsStarted, th.pathsEnded)
				assert.Equal(t, th.opsStarted, th.opsEnded)
				assert.Equal(t, th.respsStarted, th.respsEnded)
				assert.Equal(t, th.paramsStarted, th.paramsEnded)

				// assert that everything was read properly
				assert.ElementsMatch(t, tt.expectedPaths, th.pathsStarted, "expected paths differs from actual")
				assert.ElementsMatch(t, tt.expectedOps, th.opsStarted, "expected ops differs from actual")
				assert.ElementsMatch(t, tt.expectedResponses, th.respsStarted, "expected responses differs from actual")
				assert.ElementsMatch(t, tt.expectedParams, th.paramsStarted, "expected params differs from actual")
			}
		})

	}
}

// Define a test implementation of the Handler interface for testing purposes
type testHandler struct {
	scanStarted   bool
	scanEnded     bool
	pathsStarted  []string
	pathsEnded    []string
	opsStarted    []string
	opsEnded      []string
	respsStarted  []string
	respsEnded    []string
	paramsStarted []string
	paramsEnded   []string
}

func (th *testHandler) StartScan(apiSpec *spec.Swagger) error {
	th.scanStarted = true
	return nil
}

func (th *testHandler) EndScan() error {
	th.scanEnded = true
	return nil
}

func (th *testHandler) StartPath(path string, pathSpec spec.PathItem) error {
	th.pathsStarted = append(th.pathsStarted, path)
	return nil
}

func (th *testHandler) EndPath() error {
	th.pathsEnded = append(th.pathsEnded, th.pathsStarted[len(th.pathsStarted)-1])
	return nil
}

func (th *testHandler) StartOp(op string, opSpec *spec.Operation) error {
	currentPath := th.pathsStarted[len(th.pathsStarted)-1]
	th.opsStarted = append(th.opsStarted, fmt.Sprintf("%s %s", op, currentPath))
	return nil
}

func (th *testHandler) EndOp() error {
	th.opsEnded = append(th.opsEnded, th.opsStarted[len(th.opsStarted)-1])
	return nil
}

func (th *testHandler) StartResponse(statusCode int, respSpec *spec.Response) error {
	th.respsStarted = append(th.respsStarted, fmt.Sprintf("%d %s", statusCode, respSpec.Description))
	return nil
}

func (th *testHandler) EndResponse() error {
	th.respsEnded = append(th.respsEnded, th.respsStarted[len(th.respsStarted)-1])
	return nil
}

func (th *testHandler) StartParam(paramSpec *spec.Parameter) error {
	th.paramsStarted = append(th.paramsStarted, fmt.Sprintf("%s:%s", th.opsStarted[len(th.opsStarted)-1], paramSpec.Name))
	return nil
}

func (th *testHandler) EndParam() error {
	th.paramsEnded = append(th.paramsEnded, th.paramsStarted[len(th.paramsStarted)-1])
	return nil
}
