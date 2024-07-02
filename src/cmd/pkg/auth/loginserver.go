package auth

import (
	"bufio"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/token"

	"github.com/pkg/errors"
)

type loginServer struct {
	loginPath         string
	org               string
	skipSetDefaultOrg bool
	store             *token.Store
	openFunc          browserOpenFunc
}

type loginResponse struct {
	Token  string `json:"token"`
	Tenant string `json:"tenant"`
}

type wrongTenantError struct {
	expected string
	actual   string
}

func (w *wrongTenantError) Error() string {
	return fmt.Sprintf("expected tenant %q but got %q", w.expected, w.actual)
}

func newLoginServer(store *token.Store, openFunc browserOpenFunc, opts *loginOpts) (*loginServer, error) {
	var defaultOrg string
	if defaultOrgToken, err := store.Get(defaultOrgPath); err == nil {
		defaultOrg = string(defaultOrgToken.Value)
	}

	orgName := cmp.Or(
		opts.orgName,
		os.Getenv(client.ChronosphereOrgNameKey),
		os.Getenv(client.ChronosphereOrgKey),
		defaultOrg,
	)
	if orgName == "" {
		return nil, errors.New("could not determine Chronosphere org")
	}

	return &loginServer{
		loginPath:         opts.loginPath,
		org:               orgName,
		skipSetDefaultOrg: opts.skipSetDefaultOrg,
		store:             store,
		openFunc:          openFunc,
	}, nil
}

func (s *loginServer) login(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
	// Create channels that the server will send response or error on
	respCh := make(chan loginResponse)
	errCh := make(chan error)

	originURL := url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s.chronosphere.io", s.org),
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(s.makeHandler(originURL.String(), respCh, errCh)))
	// Add a /health endpoint so that we can ensure that the server is up and running before we redirect to the browser
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	srv := &http.Server{
		Handler: mux,
	}
	defer srv.Shutdown(ctx) // nolint:errcheck

	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return errors.WithStack(err)
	}

	// Start the server in a goroutine so that we don't block on the server
	go func() {
		errCh <- errors.WithStack(srv.Serve(listener))
	}()

	// Make sure the server has started before proceeding
	if err := s.checkHealth(listener, 5); err != nil {
		return errors.WithStack(err)
	}

	// Start another goroutine to allow the user to provide the session ID via stdin if their browser is unable
	// to send it directly to the server
	go func() {
		// input is delimited by the `;` character as the token may be spread across several lines.
		// This is due to a terminal limitation where lines of input can only be so long (e.g. 1024 characters on MacOS)
		const sessionDelimiter = ';'
		fmt.Fprintln(stderr, "Waiting for input from the browser. You can also input the session id via stdin.") //nolint:errcheck
		r := bufio.NewReader(stdin)
		input, err := r.ReadString(sessionDelimiter)
		if err != nil {
			errCh <- errors.WithMessage(err, "could not get input from stdin")
			return
		}
		// input may contain newlines as the maximum terminal input length per-line is generally 1024
		input = strings.ReplaceAll(input, "\n", "")
		respCh <- loginResponse{
			Token:  strings.TrimSuffix(input, string(sessionDelimiter)),
			Tenant: s.org,
		}
	}()

	// Open the browser to the CLI login URL
	authURL := originURL.JoinPath(s.loginPath)
	urlParams := authURL.Query()
	// We can't use listener.Addr().String() here since it produces an address of the form `[::]:PORT` which we're not
	// allowed to send to under our CSP, which explicitly requires the form `localhost:PORT`
	urlParams.Set("token_callback", fmt.Sprintf("localhost:%d", listener.Addr().(*net.TCPAddr).Port))
	authURL.RawQuery = urlParams.Encode()
	if s.openFunc(authURL.String()) != nil {
		return errors.WithStack(err)
	}

	// Wait for a token or error response from goroutines
	select {
	case r := <-respCh:
		if r.Tenant != s.org {
			return &wrongTenantError{expected: s.org, actual: r.Tenant}
		}
		if err := s.store.Put(s.org, token.Token{
			Value:  []byte(r.Token),
			Expiry: time.Now().Add(24 * time.Hour),
		}); err != nil {
			return errors.WithStack(err)
		}
		fmt.Fprintln(stdout, "Successfully added session id to local credentials store.") //nolint:errcheck
		if s.skipSetDefaultOrg {
			return nil
		}
		if err := setDefaultOrg(s.store, s.org); err != nil {
			fmt.Fprintf(stderr, "Failed to set default organization: %s\n", err)
		}
		return nil
	case err := <-errCh:
		return errors.WithStack(err)
	case <-ctx.Done():
		return errors.WithStack(ctx.Err())
	}
}

func (s *loginServer) makeHandler(origin string, respCh chan loginResponse, errCh chan error) func(http.ResponseWriter, *http.Request) {
	var headers string
	return func(w http.ResponseWriter, r *http.Request) {
		// Verify that the origin is what we expect so that we don't accept a token from the wrong tenant
		if o := r.Header.Get("Origin"); o != origin {
			w.WriteHeader(http.StatusForbidden)
			errCh <- errors.Errorf("unexpected origin: %s", o)
			return
		}

		// Respond to OPTIONS preflight requests that provide metadata on the subsequent cross-origin request
		if r.Method == http.MethodOptions {
			// Verify that the method is POST, as that's the only method we expect from the browser
			if m := r.Header.Get("Access-Control-Request-Method"); m != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				errCh <- errors.Errorf("unexpected method: %s", m)
				return
			}

			// Set headers so that the OPTIONS response, and later, the POST response, allows only the headers explicitly
			// requested by the client in the preflight request
			headers = r.Header.Get("Access-Control-Request-Headers")
			s.setCORSHeaders(w, origin, headers)
			w.WriteHeader(http.StatusOK)
			return
		}

		// Only allow POST requests, as that's the only way we'll receive the loginResponse data
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			errCh <- errors.New("request method not allowed")
			return
		}

		// Set CORS headers to properly accept the request. If these aren't set, the browser will show that the
		// request failed, even if we've successfully received the loginResponse data.
		s.setCORSHeaders(w, origin, headers)

		// Read request body
		data, err := io.ReadAll(r.Body)
		defer r.Body.Close() //nolint:errcheck
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errCh <- errors.WithMessage(err, "error reading request body")
			return
		}

		// Handle the received data
		var resp loginResponse
		err = json.Unmarshal(data, &resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errCh <- errors.WithStack(err)
			return
		}
		respCh <- resp
	}
}

func (s *loginServer) setCORSHeaders(w http.ResponseWriter, origin, headers string) {
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", http.MethodPost)
	w.Header().Set("Access-Control-Allow-Headers", headers)
}

// checkHealth will attempt to call the server's health check endpoint, waiting twice as long each attempt
// before failing after `retries` attempts
func (s *loginServer) checkHealth(l net.Listener, retries int) error {
	backoff := 50 * time.Millisecond
	for range retries {
		healthURL := url.URL{
			Scheme: "http",
			Host:   l.Addr().String(),
			Path:   "/health",
		}
		resp, err := http.Get(healthURL.String())
		if err != nil || resp.StatusCode != http.StatusOK {
			time.Sleep(backoff)
			backoff *= 2
			continue
		}
		return nil
	}
	return errors.New("could not start server to receive session ID")
}
