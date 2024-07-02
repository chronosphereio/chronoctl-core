package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/env"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/token"

	"github.com/stretchr/testify/require"
)

const (
	tenantName     = "faketenant"
	validSessionID = "fake-session"
)

type fakeBrowser struct {
	t          *testing.T
	loginPath  string
	dataToSend []byte
}

func newFakeBrowser(t *testing.T, loginPath string, dataToSend []byte) *fakeBrowser {
	return &fakeBrowser{
		t:          t,
		loginPath:  loginPath,
		dataToSend: dataToSend,
	}
}

func (b *fakeBrowser) open(targetURL string) error {
	b.t.Helper()
	// Validate that the browser is opening to a correct URL
	validURL := url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s.chronosphere.io", tenantName),
		Path:   b.loginPath,
	}
	query := validURL.Query()
	query.Set("token_callback", "localhost:PORT")
	validURL.RawQuery = query.Encode()

	// Replace the port with `PORT` since port is arbitrary based on the `localhost:0` listener
	// %3A is the URL escaped `:` character
	r := regexp.MustCompile("localhost%3A[0-9]+")
	require.Equal(b.t, validURL.String(), r.ReplaceAllString(targetURL, "localhost%3APORT"))

	// Simulate sending data from the browser to the CLI.
	// Done async to simulate the browser returning nil here immediately once the page has been opened.
	go func() {
		parsedTargetURL, err := url.Parse(targetURL)
		require.NoError(b.t, err)

		callback := url.URL{
			Scheme: "http",
			Host:   parsedTargetURL.Query().Get("token_callback"),
		}

		// Send an OPTIONS preflight request as a real browser will send the same request to determine if
		// the cross-origin request is allowed
		client := http.Client{}
		preflight, err := http.NewRequest(http.MethodOptions, callback.String(), nil)
		require.NoError(b.t, err)

		// A real browser will set these headers to let the server know what type of cross-origin request to expect
		preflight.Header.Set("Access-Control-Request-Method", http.MethodPost)
		preflight.Header.Set("Access-Control-Request-Headers", "content-type")
		origin := &url.URL{Scheme: parsedTargetURL.Scheme, Host: parsedTargetURL.Host}
		preflight.Header.Set("Origin", origin.String())
		resp, err := client.Do(preflight)
		require.NoError(b.t, err)
		require.Equal(b.t, http.StatusOK, resp.StatusCode)

		// Send the POST request with the session data
		request, err := http.NewRequest(http.MethodPost, callback.String(), bytes.NewReader(b.dataToSend))
		require.NoError(b.t, err)
		request.Header.Set("Origin", origin.String())
		response, err := client.Do(request)
		require.NoError(b.t, err)
		require.Equal(b.t, http.StatusOK, response.StatusCode)
	}()

	return nil
}

// noDataBrowserOpen simulates a browser that can't send anything to the terminal, i.e. Safari
func noDataBrowserOpen(string) error {
	return nil
}

func TestAuthLogin(t *testing.T) {
	tests := []struct {
		name                   string
		flags                  []string
		envChronosphereOrgName string
		envChronosphereOrg     string
		openFunc               browserOpenFunc
		dataFromStdin          string
		defaultOrg             string
		wantSessionID          string
		wantDefaultTenant      bool
		wantErr                error
	}{
		{
			name:  "valid data from browser",
			flags: []string{"--org-name", tenantName},
			openFunc: func() browserOpenFunc {
				browserRespBytes, err := json.Marshal(loginResponse{
					Token:  validSessionID,
					Tenant: tenantName,
				})
				require.NoError(t, err)
				fb := newFakeBrowser(t, defaultLoginPath, browserRespBytes)
				return fb.open
			}(),
			wantSessionID:     validSessionID,
			wantDefaultTenant: true,
		},
		{
			name:  "valid data from browser, login path override",
			flags: []string{"--org-name", tenantName, "--login-path", "/otherloginpath"},
			openFunc: func() browserOpenFunc {
				browserRespBytes, err := json.Marshal(loginResponse{
					Token:  validSessionID,
					Tenant: tenantName,
				})
				require.NoError(t, err)
				fb := newFakeBrowser(t, "/otherloginpath", browserRespBytes)
				return fb.open
			}(),
			wantSessionID:     validSessionID,
			wantDefaultTenant: true,
		},
		{
			name:                   "valid data from browser, CHRONOSPHERE_ORG_NAME set",
			envChronosphereOrgName: tenantName,
			openFunc: func() browserOpenFunc {
				browserRespBytes, err := json.Marshal(loginResponse{
					Token:  validSessionID,
					Tenant: tenantName,
				})
				require.NoError(t, err)
				fb := newFakeBrowser(t, defaultLoginPath, browserRespBytes)
				return fb.open
			}(),
			wantSessionID:     validSessionID,
			wantDefaultTenant: true,
		},
		{
			name:               "valid data from browser, CHRONOSPHERE_ORG set",
			envChronosphereOrg: tenantName,
			openFunc: func() browserOpenFunc {
				browserRespBytes, err := json.Marshal(loginResponse{
					Token:  validSessionID,
					Tenant: tenantName,
				})
				require.NoError(t, err)
				fb := newFakeBrowser(t, defaultLoginPath, browserRespBytes)
				return fb.open
			}(),
			wantSessionID:     validSessionID,
			wantDefaultTenant: true,
		},
		{
			name:  "browser returns token with wrong tenant",
			flags: []string{"--org-name", tenantName},
			openFunc: func() browserOpenFunc {
				browserRespBytes, err := json.Marshal(loginResponse{
					Token:  validSessionID,
					Tenant: "wrong-tenant",
				})
				require.NoError(t, err)
				fb := newFakeBrowser(t, defaultLoginPath, browserRespBytes)
				return fb.open
			}(),
			wantErr: &wrongTenantError{expected: tenantName, actual: "wrong-tenant"},
		},
		{
			name:     "data from stdin",
			flags:    []string{"--org-name", tenantName},
			openFunc: noDataBrowserOpen,
			// data pasted from stdin includes newlines and the `;` delimiter since stdin is limited to 1024 characters
			// in most terminals
			dataFromStdin:     "validsession\n1234\n5678;",
			wantSessionID:     "validsession12345678",
			wantDefaultTenant: true,
		},
		{
			name:                   "data from stdin, CHRONOSPHERE_ORG_NAME set",
			envChronosphereOrgName: tenantName,
			openFunc:               noDataBrowserOpen,
			// data pasted from stdin includes newlines and the `;` delimiter since stdin is limited to 1024 characters
			// in most terminals
			dataFromStdin:     "validsession\n1234\n5678;",
			wantSessionID:     "validsession12345678",
			wantDefaultTenant: true,
		},
		{
			name:               "data from stdin, CHRONOSPHERE_ORG set",
			envChronosphereOrg: tenantName,
			openFunc:           noDataBrowserOpen,
			// data pasted from stdin includes newlines and the `;` delimiter since stdin is limited to 1024 characters
			// in most terminals
			dataFromStdin:     "validsession\n1234\n5678;",
			wantSessionID:     "validsession12345678",
			wantDefaultTenant: true,
		},
		{
			name:          "data from stdin not delimited with `;` fails",
			flags:         []string{"--org-name", tenantName},
			openFunc:      noDataBrowserOpen,
			dataFromStdin: "validsession\n1234\n5678",
			wantErr:       io.EOF,
		},
		{
			name:  "don't set defaults",
			flags: []string{"--org-name", tenantName, "--skip-set-default-org", "true"},
			openFunc: func() browserOpenFunc {
				browserRespBytes, err := json.Marshal(loginResponse{
					Token:  validSessionID,
					Tenant: tenantName,
				})
				require.NoError(t, err)
				fb := newFakeBrowser(t, defaultLoginPath, browserRespBytes)
				return fb.open
			}(),
			wantSessionID: validSessionID,
		},
		{
			name:       "login without org-name or env falls back to default org",
			defaultOrg: tenantName,
			openFunc: func() browserOpenFunc {
				browserRespBytes, err := json.Marshal(loginResponse{
					Token:  validSessionID,
					Tenant: tenantName,
				})
				require.NoError(t, err)
				fb := newFakeBrowser(t, defaultLoginPath, browserRespBytes)
				return fb.open
			}(),
			wantSessionID:     validSessionID,
			wantDefaultTenant: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			store := token.NewFileStore(tmpDir)
			if tt.defaultOrg != "" {
				require.NoError(t, store.Put(defaultOrgPath, token.Token{
					Value:  []byte(tt.defaultOrg),
					Expiry: time.Now().Add(time.Hour * 24 * 365),
				}))
			}

			c := subcommand{store: store}
			cmd := c.newAuthLoginCmd(tt.openFunc)
			reader, writer := io.Pipe()
			cmd.SetIn(reader)

			// Send input to pipe asynchronously, as Write will block until reader.Read is called
			if tt.dataFromStdin != "" {
				go func() {
					_, err := writer.Write([]byte(tt.dataFromStdin))
					require.NoError(t, err)
					require.NoError(t, writer.Close())
				}()
			}

			t.Setenv(env.ChronosphereOrgNameKey, tt.envChronosphereOrgName)
			t.Setenv(env.ChronosphereOrgKey, tt.envChronosphereOrg)
			cmd.SetArgs(tt.flags)
			err := cmd.Execute()
			if tt.wantErr != nil {
				require.ErrorContains(t, err, tt.wantErr.Error())
				return
			}
			require.NoError(t, err)

			sessionToken, err := store.Get(tenantName)
			require.NoError(t, err)
			require.Equal(t, tt.wantSessionID, string(sessionToken.Value))

			// Verify that the correct org is set as default if applicable
			defaultOrg, err := store.Get(defaultOrgPath)
			if tt.wantDefaultTenant {
				require.NoError(t, err)
				require.Equal(t, tenantName, string(defaultOrg.Value))
			} else {
				require.ErrorIs(t, err, token.ErrNotExist)
			}
		})
	}
}

func TestAuthSetDefaultOrg(t *testing.T) {
	tests := []struct {
		name               string
		args               []string
		existingDefaultOrg string
		wantErr            error
	}{
		{
			name: "no existing default set",
			args: []string{"fakeOrg"},
		},
		{
			name:               "existing default set",
			args:               []string{"fakeOrg"},
			existingDefaultOrg: "wrongOrg",
		},
		{
			name:    "empty org",
			args:    []string{},
			wantErr: errMustIncludeOrgName,
		},
		{
			name:    "multiple orgs",
			args:    []string{"fakeOrg1", "fakeOrg2"},
			wantErr: errMustIncludeOrgName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := token.NewFileStore(t.TempDir())
			if tt.existingDefaultOrg != "" {
				require.NoError(t, store.Put(defaultOrgPath, token.Token{
					Value:  []byte(tt.existingDefaultOrg),
					Expiry: time.Now().Add(time.Hour * 24 * 365),
				}))
			}

			c := subcommand{store: store}
			cmd := c.newSetDefaultOrgCmd()
			cmd.SetArgs(tt.args)
			err := cmd.Execute()
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)

			defaultOrg, err := store.Get(defaultOrgPath)
			require.NoError(t, err)
			require.Equal(t, tt.args[0], string(defaultOrg.Value))
		})
	}
}

func TestAuthPrintSessionID(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		store         *token.Store
		wantSessionID string
		wantErr       error
	}{
		{
			name: "valid session id",
			args: []string{tenantName},
			store: func() *token.Store {
				store := token.NewFileStore(t.TempDir())
				require.NoError(t, store.Put(tenantName, token.Token{
					Value:  []byte(validSessionID),
					Expiry: time.Now().Add(time.Hour * 24),
				}))
				return store
			}(),
			wantSessionID: validSessionID,
		},
		{
			name: "expired session id",
			args: []string{tenantName},
			store: func() *token.Store {
				store := token.NewFileStore(t.TempDir())
				require.NoError(t, store.Put(tenantName, token.Token{
					Value: []byte(validSessionID),
					// Token expired an hour ago
					Expiry: time.Now().Add(-time.Hour),
				}))
				return store
			}(),
			wantErr: token.ErrTokenExpired,
		},
		{
			name:    "session id doesn't exist",
			args:    []string{tenantName},
			store:   token.NewFileStore(t.TempDir()),
			wantErr: token.ErrNotExist,
		},
		{
			name:    "empty org fails",
			args:    []string{},
			store:   token.NewFileStore(t.TempDir()),
			wantErr: errMustIncludeOrgName,
		},
		{
			name: "multiple orgs fails",
			args: []string{"fakeOrg1", "fakeOrg2"},
			store: func() *token.Store {
				store := token.NewFileStore(t.TempDir())
				// Store contains both orgs, but we don't allow printing multiple orgs for simplicity
				require.NoError(t, store.Put("fakeOrg1", token.Token{
					Value:  []byte(validSessionID),
					Expiry: time.Now().Add(time.Hour * 24),
				}))
				require.NoError(t, store.Put("fakeOrg2", token.Token{
					Value:  []byte(validSessionID),
					Expiry: time.Now().Add(time.Hour * 24),
				}))
				return store
			}(),
			wantErr: errMustIncludeOrgName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := subcommand{store: tt.store}
			cmd := c.newPrintAccessTokenCmd()
			cmd.SetArgs(tt.args)
			stdout := &bytes.Buffer{}
			cmd.SetOut(stdout)

			err := cmd.Execute()
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantSessionID, strings.TrimSpace(stdout.String()))
		})
	}
}

func TestAuthList(t *testing.T) {
	tests := []struct {
		name         string
		tokens       map[string]token.Token
		wantEntries  []listEntry
		wantNoOutput bool
	}{
		{
			name: "tokens with default",
			tokens: map[string]token.Token{
				defaultOrgPath: {
					Value:  []byte(tenantName),
					Expiry: time.Now().Add(time.Hour),
				},
				tenantName: {
					Value:  []byte("my-token"),
					Expiry: time.Now().Add(time.Hour),
				},
				"other-tenant": {
					Value:  []byte("my-other-token"),
					Expiry: time.Now().Add(time.Hour),
				},
				"expired-tenant": {
					Value:  []byte("expired-token"),
					Expiry: time.Now().Add(-time.Hour),
				},
			},
			wantEntries: []listEntry{
				{
					Organization: tenantName,
					Valid:        true,
					Default:      true,
				},
				{
					Organization: "other-tenant",
					Valid:        true,
					Default:      false,
				},
				{
					Organization: "expired-tenant",
					Valid:        false,
					Default:      false,
				},
			},
		},
		{
			name: "tokens without default",
			tokens: map[string]token.Token{
				tenantName: {
					Value:  []byte("my-token"),
					Expiry: time.Now().Add(time.Hour),
				},
				"other-tenant": {
					Value:  []byte("my-other-token"),
					Expiry: time.Now().Add(time.Hour),
				},
				"expired-tenant": {
					Value:  []byte("expired-token"),
					Expiry: time.Now().Add(-time.Hour),
				},
			},
			wantEntries: []listEntry{
				{
					Organization: tenantName,
					Valid:        true,
					Default:      false,
				},
				{
					Organization: "other-tenant",
					Valid:        true,
					Default:      false,
				},
				{
					Organization: "expired-tenant",
					Valid:        false,
					Default:      false,
				},
			},
		},
		{
			name: "expired default",
			tokens: map[string]token.Token{
				defaultOrgPath: {
					Value:  []byte(tenantName),
					Expiry: time.Now().Add(-time.Hour),
				},
				tenantName: {
					Value:  []byte("my-token"),
					Expiry: time.Now().Add(time.Hour),
				},
			},
			wantEntries: []listEntry{
				{
					Organization: tenantName,
					Valid:        true,
					Default:      false,
				},
			},
		},
		{
			name:         "no tokens",
			tokens:       map[string]token.Token{},
			wantNoOutput: true,
		},
		{
			name: "default set but no tokens",
			tokens: map[string]token.Token{
				defaultOrgPath: {
					Value:  []byte(tenantName),
					Expiry: time.Now().Add(time.Hour),
				},
			},
			wantNoOutput: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := token.NewFileStore(t.TempDir())
			for name, tok := range tt.tokens {
				require.NoError(t, store.Put(name, tok))
			}

			c := subcommand{store: store}
			cmd := c.newListCmd()
			stdout := &bytes.Buffer{}
			cmd.SetOut(stdout)
			cmd.SetArgs([]string{"-o", "json"})

			err := cmd.Execute()
			require.NoError(t, err)

			if tt.wantNoOutput {
				require.Empty(t, stdout.String())
				return
			}
			var gotEntries []listEntry
			require.NoError(t, json.Unmarshal(stdout.Bytes(), &gotEntries))
			require.ElementsMatch(t, tt.wantEntries, gotEntries)
		})
	}
}
