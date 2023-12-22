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
	"os"
	"strings"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"

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

	// ChronosphereOrgNameKey is the environment variable that specifies the Chronosphere customer organization
	ChronosphereOrgNameKey = "CHRONOSPHERE_ORG_NAME"
	// ChronosphereOrgKey is the environment variable that specifies the Chronosphere customer organization
	ChronosphereOrgKey = "CHRONOSPHERE_ORG" // fallback for CHRONOSPHERE_ORG_NAME
	// ChronosphereAPITokenKey is the environment variable that specifies the Chronosphere API token
	ChronosphereAPITokenKey = "CHRONOSPHERE_API_TOKEN"
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
	apiToken, err := f.getAPIToken()
	if err != nil {
		return nil, err
	}

	transport, err := transport.New(transport.RuntimeConfig{
		Component:          component,
		OrgName:            f.OrgName,
		APIToken:           apiToken,
		APIUrl:             f.APIUrl,
		InsecureSkipVerify: f.InsecureSkipVerify,
		AllowHTTP:          f.AllowHTTP,
		DefaultBasePath:    basePath,
	})
	if err != nil {
		return nil, err
	}

	return transport, nil
}

// AddFlags adds client flags to a Cobra command.
func (f *Flags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.APIToken, "api-token", "", "The client API token used to authenticate to user. Mutally exclusive with --api-token-filename. If both --api-token and --api-token-filename are unset, the "+ChronosphereAPITokenKey+" environment variable is used.")
	cmd.Flags().StringVar(&f.APITokenFilename, "api-token-filename", "", "A file containing the API token used for authentication. Mutally exclusive with --api-token. If both --api-token and --api-token-filename are unset, the "+ChronosphereAPITokenKey+" environment variable is used.")
	cmd.Flags().StringVar(&f.OrgName, "org-name", "", "The name of your team's Chronosphere organization. Defaults to "+ChronosphereOrgNameKey+" environment variable.")
	cmd.Flags().IntVar(&f.TimeoutSeconds, "timeout", defaultTimeoutSeconds, "The timeout of the request in seconds, defaults to 30 seconds")

	cmd.Flags().StringVar(&f.APIUrl, "api-url", f.APIUrl, "The URL of the Chronosphere API. Defaults to https://<organization>.chronosphere.io/api.")
	cmd.Flags().MarkHidden("api-url") //nolint:errcheck

	cmd.Flags().BoolVar(&f.AllowHTTP, "allow-http", false, "Allow connections to the API over HTTP.")
	cmd.Flags().MarkHidden("allow-http") //nolint:errcheck

	cmd.Flags().BoolVar(&f.InsecureSkipVerify, "insecure-skip-verify", false, "If true, TLS accepts any certificate presented by the server and any host name in that certificate. This should be used only for testing..")
	cmd.Flags().MarkHidden("insecure-skip-verify") //nolint:errcheck
}

// Timeout returns the value of the timeout flag as a time.Duration.
func (f *Flags) Timeout() time.Duration {
	return time.Duration(f.TimeoutSeconds) * time.Second
}

func (f *Flags) getAPIToken() (string, error) {
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

	if key := os.Getenv(ChronosphereAPITokenKey); key != "" {
		return key, nil
	}

	return "", errors.New("client API token must be provided via --api-token, --api-token-filename, or " + ChronosphereAPITokenKey + " environment variable")
}
