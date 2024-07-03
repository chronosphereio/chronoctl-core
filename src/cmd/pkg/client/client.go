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

// Package client implements client flags for chronoctl CLI commands and gives commands
// access to client objects that can be used to make API calls.
package client

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/auth"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/env"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/token"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/transport"
	config_unstable "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/client/operations"
	config_v1 "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/client/operations"
	state_unstable "github.com/chronosphereio/chronoctl-core/src/generated/swagger/stateunstable/client/operations"
	state_v1 "github.com/chronosphereio/chronoctl-core/src/generated/swagger/statev1/client/operations"
)

const (
	// TestChronosphereURL is the URL of the Chronosphere API used for testing.
	TestChronosphereURL = "https://myorg.chronosphere.io/api"
	// TestBadURL is a bad URL used for testing.
	TestBadURL = "://bad.domain.io"

	defaultTimeoutSeconds = 30
	apiURLFormat          = "https://%s.chronosphere.io%s"
)

// Clients is a list of clients our generated CLI needs access to.
type Clients struct {
	ConfigV1       config_v1.ClientService
	ConfigUnstable config_unstable.ClientService
}

// Flags composes common flags used to configure the Chronosphere API client.
type Flags struct {
	APIToken           string
	APITokenFilename   string
	OrgName            string
	TimeoutSeconds     int
	APIUrl             string // dev flag
	AllowHTTP          bool   // dev flag
	InsecureSkipVerify bool   // dev flag
	TokenStoreDir      string // dev flag
}

// NewClientFlags returns a new Flags object.
func NewClientFlags() *Flags {
	return &Flags{}
}

// ConfigUnstableClient creates a new client to hit config/unstable APIs. This client should only be used in hidden commands
// as these APIs are subject to change.
func (f *Flags) ConfigUnstableClient() (config_unstable.ClientService, error) {
	transport, err := f.Transport(transport.ComponentChronoctlUnstable, "")
	if err != nil {
		return nil, fmt.Errorf("could not construct Chronosphere config unstable API client: %v", err)
	}

	return config_unstable.New(transport, strfmt.Default), nil
}

// StateUnstableClient creates a new client to hit state/unstable APIs. This client should only be used in hidden commands
// as these APIs are subject to change.
func (f *Flags) StateUnstableClient() (state_unstable.ClientService, error) {
	transport, err := f.Transport(transport.ComponentChronoctlUnstable, "")
	if err != nil {
		return nil, fmt.Errorf("could not construct Chronosphere state unstable API client: %v", err)
	}

	return state_unstable.New(transport, strfmt.Default), nil
}

// StateV1Client creates a new client to hit statev1 APIs.
func (f *Flags) StateV1Client() (state_v1.ClientService, error) {
	transport, err := f.Transport(transport.ComponentChronoctl, "")
	if err != nil {
		return nil, fmt.Errorf("could not construct Chronosphere state v1 API client: %v", err)
	}

	return state_v1.New(transport, strfmt.Default), nil
}

// ConfigV1Client creates a new client to hit configv1 APIs.
func (f *Flags) ConfigV1Client() (config_v1.ClientService, error) {
	transport, err := f.Transport(transport.ComponentChronoctl, "")
	if err != nil {
		return nil, fmt.Errorf("could not construct Chronosphere config v1 API client: %v", err)
	}

	return config_v1.New(transport, strfmt.Default), nil
}

// Transport returns a new transport for the given api base path.
func (f *Flags) Transport(component transport.Component, basePath string) (*httptransport.Runtime, error) {
	apiURL, err := f.getAPIURL(basePath)
	if err != nil {
		return nil, err
	}
	apiToken, err := f.getAPIToken(apiURL)
	if err != nil {
		return nil, err
	}

	transport, err := transport.New(transport.RuntimeConfig{
		Component:          component,
		APIToken:           apiToken,
		APIUrl:             apiURL,
		InsecureSkipVerify: f.InsecureSkipVerify,
		AllowHTTP:          f.AllowHTTP,
		DefaultBasePath:    basePath,
		EntityNamespace:    f.getEntityNamespace(),
	})
	if err != nil {
		return nil, err
	}

	return transport, nil
}

// AddFlags adds client flags to a Cobra command.
func (f *Flags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.APIToken, "api-token", "", "The client API token used to authenticate to user. Mutally exclusive with --api-token-filename. If both --api-token and --api-token-filename are unset, the "+env.ChronosphereAPITokenKey+" environment variable is used.")
	cmd.Flags().StringVar(&f.APITokenFilename, "api-token-filename", "", "A file containing the API token used for authentication. Mutally exclusive with --api-token. If both --api-token and --api-token-filename are unset, the "+env.ChronosphereAPITokenKey+" environment variable is used.")
	cmd.Flags().StringVar(&f.OrgName, "org-name", "", "The name of your team's Chronosphere organization. Defaults to "+env.ChronosphereOrgNameKey+" environment variable.")
	cmd.Flags().IntVar(&f.TimeoutSeconds, "timeout", defaultTimeoutSeconds, "The timeout of the request in seconds, defaults to 30 seconds")

	cmd.Flags().StringVar(&f.APIUrl, "api-url", f.APIUrl, "The URL of the Chronosphere API. Defaults to https://<organization>.chronosphere.io/api.")
	cmd.Flags().MarkHidden("api-url") //nolint:errcheck

	cmd.Flags().BoolVar(&f.AllowHTTP, "allow-http", false, "Allow connections to the API over HTTP.")
	cmd.Flags().MarkHidden("allow-http") //nolint:errcheck

	cmd.Flags().BoolVar(&f.InsecureSkipVerify, "insecure-skip-verify", false, "If true, TLS accepts any certificate presented by the server and any host name in that certificate. This should be used only for testing..")
	cmd.Flags().MarkHidden("insecure-skip-verify") //nolint:errcheck

	cmd.Flags().StringVar(&f.TokenStoreDir, "token-store-dir", "", "If set, overwrites the token store directory used to obtain a token for requests. This should be used only for testing..")
	cmd.Flags().MarkHidden("token-store-dir") //nolint:errcheck
}

// Timeout returns the value of the timeout flag as a time.Duration.
func (f *Flags) Timeout() time.Duration {
	return time.Duration(f.TimeoutSeconds) * time.Second
}

func (f *Flags) getEntityNamespace() string {
	if ns := os.Getenv(env.ChronosphereEntityNamespace); ns != "" {
		return ns
	}
	return ""
}

func (f *Flags) getAPIToken(apiURL string) (string, error) {
	if f.APIToken != "" && f.APITokenFilename != "" {
		return "", errors.New("only one of --api-token and --api-token-filename can be set")
	}

	if f.APIToken != "" {
		return f.APIToken, nil
	}

	if f.APITokenFilename != "" {
		b, err := os.ReadFile(f.APITokenFilename)
		if err != nil {
			return "", fmt.Errorf("reading api token from file %s: %w", f.APITokenFilename, err)
		}
		return strings.TrimSpace(string(b)), nil
	}

	if key := os.Getenv(env.ChronosphereAPITokenKey); key != "" {
		return key, nil
	}

	// Determine the org from the API URL to determine if we have a token stored for that org
	URL, err := url.Parse(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse API url %q: %w", apiURL, err)
	}
	orgSubdomain := strings.Split(URL.Hostname(), ".")[0]
	t, err := f.getTokenFromStore(orgSubdomain)
	if err != nil {
		return "", errors.Join(err, errors.New("client API token must be provided via --api-token, --api-token-filename, "+env.ChronosphereAPITokenKey+" environment variable, or by authenticating with 'auth login'"))
	}
	return t, nil
}

func (f *Flags) getTokenFromStore(org string) (string, error) {
	store, err := f.getStore()
	if err != nil {
		return "", err
	}
	t, err := store.Get(org)
	if err != nil {
		return "", fmt.Errorf("unable to get token for org %q from store: %v", org, err)
	}
	return string(t.Value), nil
}

func (f *Flags) getAPIURL(basePath string) (string, error) {
	if f.APIUrl != "" {
		return f.APIUrl, nil
	}
	if f.OrgName == "" {
		if f.OrgName = os.Getenv(env.ChronosphereOrgNameKey); f.OrgName == "" {
			defaultOrg, err := f.checkDefaultOrg()
			if err != nil {
				return "", errors.Join(err, errors.New("organization must be provided as a flag, via the "+env.ChronosphereOrgNameKey+" environment variable, or by setting a default org when the API URL isn't set"))
			}
			fmt.Fprintf(os.Stderr, "assuming default org %q\n", defaultOrg) //nolint:errcheck
			f.OrgName = defaultOrg
		}
	}
	return fmt.Sprintf(apiURLFormat, f.OrgName, basePath), nil
}

func (f *Flags) checkDefaultOrg() (string, error) {
	store, err := f.getStore()
	if err != nil {
		return "", err
	}
	defaultOrg, err := auth.GetDefaultOrg(store)
	if err != nil {
		return "", fmt.Errorf("unable to get default organization: %v", err)
	}
	return defaultOrg, nil
}

func (f *Flags) getStore() (*token.Store, error) {
	if f.TokenStoreDir != "" {
		return token.NewFileStore(f.TokenStoreDir), nil
	}
	store, err := auth.NewChronoctlStore()
	if err != nil {
		return nil, fmt.Errorf("unable to get chronoctl store: %v", err)
	}
	return store, nil
}
