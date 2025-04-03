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

package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/env"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/token"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/transport"
	api "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/x/swagger"
)

func TestClientFlagsClient(t *testing.T) {
	tests := []struct {
		name    string
		flags   *Flags
		wantErr string
	}{
		{
			name: "valid flags",
			flags: &Flags{
				APIUrl:   TestChronosphereURL,
				APIToken: "token",
			},
		},
		{
			name: "invalid flags",
			flags: &Flags{
				APIUrl: TestBadURL,
			},
			wantErr: `unable to parse URL of Chronosphere URL: parse "://bad.domain.io": missing protocol scheme`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.flags.ConfigV1Client()
			if tt.wantErr != "" {
				assert.Error(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

type transportTestData struct {
	name      string
	component transport.Component
}

func TestClientFlagsTransport(t *testing.T) {
	tests := []struct {
		name                   string
		flags                  *Flags
		apiToken               string
		orgName                string
		wantErr                string
		wantHost               string
		wantBasePath           string
		wantInsecureSkipVerify bool
	}{
		{
			name: "all necessary flags specified",
			flags: &Flags{
				APIUrl:        TestChronosphereURL,
				APIToken:      "token",
				TokenStoreDir: t.TempDir(),
			},
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "fall back to org name flag for organization name",
			flags: &Flags{
				APIToken:      "token",
				OrgName:       "specialorg",
				TokenStoreDir: t.TempDir(),
			},
			wantHost:     "specialorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "use HTTP",
			flags: &Flags{
				APIUrl:        "http://localhost:8000/",
				APIToken:      "token",
				AllowHTTP:     true,
				TokenStoreDir: t.TempDir(),
			},
			wantHost:     "localhost:8000",
			wantBasePath: "/",
		},
		{
			name: "TLS insecure skip verify enabled",
			flags: &Flags{
				APIUrl:             TestChronosphereURL,
				APIToken:           "token",
				InsecureSkipVerify: true,
				TokenStoreDir:      t.TempDir(),
			},
			wantHost:               "myorg.chronosphere.io",
			wantBasePath:           "/api",
			wantInsecureSkipVerify: true,
		},
		{
			name: "--api-token and --api-token-filename are exclusive",
			flags: &Flags{
				APIUrl:           TestChronosphereURL,
				APIToken:         "token-param",
				APITokenFilename: "./testdata/token",
				TokenStoreDir:    t.TempDir(),
			},

			wantErr: "only one of --api-token and --api-token-filename can be set",
		},
		{
			name: "--api-token overrides environment variable",
			flags: &Flags{
				APIUrl:        TestChronosphereURL,
				APIToken:      "token-param",
				TokenStoreDir: t.TempDir(),
			},

			apiToken:     "token-param",
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "--api-token-filename parameter overrides environment variable",
			flags: &Flags{
				APIUrl:           TestChronosphereURL,
				APITokenFilename: "./testdata/token",
				TokenStoreDir:    t.TempDir(),
			},

			apiToken:     "token-from-file",
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "invalid --api-token-filename parameter returns error",
			flags: &Flags{
				APIUrl:           TestChronosphereURL,
				APITokenFilename: "./testdata/no_token",
				TokenStoreDir:    t.TempDir(),
			},

			wantErr: "reading api token from file ./testdata/no_token: open ./testdata/no_token: no such file or directory",
		},
		{
			name: "fall back to environment variable for API token",
			flags: &Flags{
				APIUrl:        TestChronosphereURL,
				TokenStoreDir: t.TempDir(),
			},
			apiToken:     "token",
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "token and store not set",
			flags: &Flags{
				APIUrl:        TestChronosphereURL,
				TokenStoreDir: t.TempDir(),
			},
			wantErr: "client API token must be provided via --api-token, --api-token-filename, CHRONOSPHERE_API_TOKEN environment variable, or by authenticating with 'auth login'",
		},
		{
			name: "token not set falls back to token store",
			flags: &Flags{
				APIUrl: TestChronosphereURL,
				TokenStoreDir: func() string {
					dir := t.TempDir()
					// Populate store directory with a token
					store := token.NewFileStore(dir)
					require.NoError(t, store.Put("myorg", token.Token{
						Value:  []byte("token"),
						Expiry: time.Now().Add(time.Hour),
					}))
					return dir
				}(),
			},
			apiToken:     "token",
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "organization and default org not set",
			flags: &Flags{
				APIToken:      "token",
				TokenStoreDir: t.TempDir(),
			},
			wantErr: "organization must be provided as a flag, via the CHRONOSPHERE_ORG_NAME environment variable, or by setting a default org when the API URL isn't set",
		},
		{
			name: "organization not set falls back to default",
			flags: &Flags{
				APIToken: "token",
				TokenStoreDir: func() string {
					dir := t.TempDir()
					// Populate store directory with a token
					store := token.NewFileStore(dir)
					require.NoError(t, store.Put("default-org", token.Token{
						Value:  []byte("myorg"),
						Expiry: time.Now().Add(time.Hour),
					}))
					return dir
				}(),
			},
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "invalid URL",
			flags: &Flags{
				APIUrl:        TestBadURL,
				APIToken:      "token",
				TokenStoreDir: t.TempDir(),
			},
			wantErr: `unable to parse URL of Chronosphere URL: parse "://bad.domain.io": missing protocol scheme`,
		},
		{
			name: "URL scheme is HTTP but HTTP wasn't allowed",
			flags: &Flags{
				APIUrl:        "http://localhost:8000",
				APIToken:      "token",
				TokenStoreDir: t.TempDir(),
			},
			wantErr: "the scheme of the API URL is HTTP but the --allow-http flag was not set",
		},
	}

	transports := []transportTestData{
		{
			name:      "api",
			component: transport.ComponentChronoctl,
		},
		{
			name:      "configunstable",
			component: transport.ComponentChronoctlUnstable,
		},
	}

	for _, tt := range tests {
		for _, tp := range transports {
			t.Run(fmt.Sprintf("%s %s", tt.name, tp.name), func(t *testing.T) {
				t.Setenv(env.ChronosphereAPITokenKey, tt.apiToken)
				t.Setenv(env.ChronosphereOrgNameKey, tt.orgName)

				tp, err := tt.flags.Transport(tp.component, "/api")
				if tt.wantErr != "" {
					assert.ErrorContains(t, err, tt.wantErr)
					return
				}

				require.NoError(t, err)
				assert.Equal(t, tt.wantHost, tp.Host)
				assert.Equal(t, tt.wantBasePath, tp.BasePath)
				assert.NotNil(t, tp.DefaultAuthentication)

				if tt.wantInsecureSkipVerify {
					httpTransport := tp.Transport.(swagger.RequestIDTrailerTransport).RT.(transport.CustomHeaderTransport).Rt.(*http.Transport) //nolint:errcheck
					assert.Equal(t, tt.wantInsecureSkipVerify, httpTransport.TLSClientConfig.InsecureSkipVerify)
				}
			})
		}
	}
}

func TestClientFlagsNewRequest(t *testing.T) {
	type envVars struct {
		token string
		org   string
	}
	tests := []struct {
		name         string
		flags        *Flags
		env          envVars
		basePath     string
		wantToken    string
		wantErr      string
		wantHost     string
		wantBasePath string
	}{
		{
			name: "all necessary flags specified",
			flags: &Flags{
				APIUrl:   TestChronosphereURL,
				APIToken: "token",
			},
			basePath:     "/api",
			wantToken:    "token",
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "--api-token and --api-token-filename are exclusive",
			flags: &Flags{
				APIUrl:           TestChronosphereURL,
				APIToken:         "token-param",
				APITokenFilename: "./testdata/token",
			},
			basePath: "/api",
			wantErr:  "only one of --api-token and --api-token-filename can be set",
		},
		{
			name: "--api-token overrides environment variable",
			flags: &Flags{
				APIUrl:   TestChronosphereURL,
				APIToken: "token-param",
			},
			basePath:     "/api",
			wantToken:    "token-param",
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "--api-token-filename parameter overrides environment variable",
			flags: &Flags{
				APIUrl:           TestChronosphereURL,
				APITokenFilename: "./testdata/token",
			},
			env: envVars{
				token: "bad-token",
			},
			basePath:     "/api",
			wantToken:    "token-from-file",
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "invalid --api-token-filename parameter returns error",
			flags: &Flags{
				APIUrl:           TestChronosphereURL,
				APITokenFilename: "./testdata/no_token",
			},
			basePath: "/api",
			wantErr:  "reading api token from file ./testdata/no_token: open ./testdata/no_token: no such file or directory",
		},
		{
			name: "fall back to environment variable for API token",
			flags: &Flags{
				APIUrl: TestChronosphereURL,
			},
			env:          envVars{token: "token"},
			basePath:     "/api",
			wantToken:    "token",
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "token and store not set",
			flags: &Flags{
				APIUrl: TestChronosphereURL,
			},
			basePath: "/api",
			wantErr:  "client API token must be provided via --api-token, --api-token-filename, CHRONOSPHERE_API_TOKEN environment variable, or by authenticating with 'auth login'",
		},
		{
			name: "token not set falls back to token store",
			flags: &Flags{
				APIUrl: TestChronosphereURL,
				TokenStoreDir: func() string {
					dir := t.TempDir()
					// Populate store directory with a token
					store := token.NewFileStore(dir)
					require.NoError(t, store.Put("myorg", token.Token{
						Value:  []byte("token"),
						Expiry: time.Now().Add(time.Hour),
					}))
					return dir
				}(),
			},
			basePath:     "/api",
			wantToken:    "token",
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "organization and default org not set",
			flags: &Flags{
				APIToken: "token",
			},
			basePath: "/api",
			wantErr:  "organization must be provided as a flag, via the CHRONOSPHERE_ORG_NAME environment variable, or by setting a default org when the API URL isn't set",
		},
		{
			name: "fall back to org name flag for organization name",
			flags: &Flags{
				APIToken: "token",
				OrgName:  "specialorg",
			},
			basePath:     "/api",
			wantToken:    "token",
			wantHost:     "specialorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "organization not set falls back to default",
			flags: &Flags{
				APIToken: "token",
				TokenStoreDir: func() string {
					dir := t.TempDir()
					// Populate store directory with a token
					store := token.NewFileStore(dir)
					require.NoError(t, store.Put("default-org", token.Token{
						Value:  []byte("myorg"),
						Expiry: time.Now().Add(time.Hour),
					}))
					return dir
				}(),
			},
			basePath:     "/api",
			wantToken:    "token",
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/api",
		},
		{
			name: "api url overrides the base path",
			flags: &Flags{
				APIUrl:   "https://myorg.chronosphere.io/notthebasepath",
				APIToken: "token",
			},
			basePath:     "/api",
			wantToken:    "token",
			wantHost:     "myorg.chronosphere.io",
			wantBasePath: "/notthebasepath",
		},
		{
			name: "invalid URL",
			flags: &Flags{
				APIUrl:   TestBadURL,
				APIToken: "token",
			},
			basePath: "/api",
			wantErr:  `parse "://bad.domain.io": missing protocol scheme`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(env.ChronosphereAPITokenKey, tt.env.token)
			t.Setenv(env.ChronosphereOrgNameKey, tt.env.org)

			if tt.flags.TokenStoreDir == "" {
				tt.flags.TokenStoreDir = t.TempDir()
			}
			req, err := tt.flags.NewRequest(http.MethodGet, tt.basePath, nil /* body */)
			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantHost, req.Host)
			assert.Equal(t, tt.wantBasePath, req.URL.Path)
			assert.Equal(t, tt.wantToken, req.Header.Get("API-Token"))
		})
	}
}

type RoundTripperFunc func(*http.Request) (*http.Response, error)

func (fn RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

func TestRequestIDTransport(t *testing.T) {
	t.Skip("FIXME: Request ID headers only work on external API errors")
	flags := &Flags{
		APIToken:  "<not real>",
		APIUrl:    "127.0.0.1",
		AllowHTTP: true,
	}
	client, err := flags.ConfigV1Client()
	require.NoError(t, err)

	httpClient := &http.Client{
		Transport: swagger.WithRequestIDTrailerTransport(
			RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
				res := &http.Response{
					Status:     "400 Bad Request",
					StatusCode: 400,
					Header: http.Header{
						"Chrono-Request-Id": []string{"abc123"},
					},
					Body: ioutil.NopCloser(strings.NewReader(`{"code": 12, "message": "hi"}`)),
				}
				return res, nil
			}),
		),
	}

	_, err = client.ReadBucket(&api.ReadBucketParams{
		HTTPClient: httpClient,
	})
	require.Error(t, err)
}
