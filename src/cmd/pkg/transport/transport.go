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

// Package transport provides methods to create an HTTP transport that can communicate with
// the Chronosphere API.
package transport

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"

	httpruntime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/buildinfo"
	xswagger "github.com/chronosphereio/chronoctl-core/src/x/swagger"
)

const (
	apiURLFormat   = "https://%s.chronosphere.io%s"
	apiTokenHeader = "api-token"

	// OrgNameEnvVar is the name of the environment variable that can be used to set the
	// Chronosphere organization name.
	OrgNameEnvVar = "CHRONOSPHERE_ORG_NAME"
)

// Component is a value that indicates the part of the CLI that is invoking an
// API. This is used to set the User-Agent header when making requests to the Chronosphere API.
type Component string

var (
	// ComponentChronoctl is the value for the stable chronoctl component
	ComponentChronoctl Component = "chrono-ctl"
	// ComponentChronoctlUnstable is the value for the unstable chronoctl component
	ComponentChronoctlUnstable Component = "unstable-chrono-ctl"
)

// RuntimeConfig is a struct that contains the configuration for creating a new HTTP transport
type RuntimeConfig struct {
	Component          Component
	OrgName            string
	APIToken           string
	APIUrl             string
	InsecureSkipVerify bool
	AllowHTTP          bool
	DefaultBasePath    string
	EntityNamespace    string
}

// New creates a new HTTP transport that can communicate with the Chronosphere API.
func New(config RuntimeConfig) (*httptransport.Runtime, error) {
	if config.APIUrl == "" {
		if config.OrgName == "" {
			if config.OrgName = os.Getenv(OrgNameEnvVar); config.OrgName == "" {
				return nil, errors.New("organization must be provided as a flag or via " + OrgNameEnvVar + " environment variable when the API URL isn't set")
			}
		}
		config.APIUrl = fmt.Sprintf(apiURLFormat, config.OrgName, config.DefaultBasePath)
	}

	apiURL, err := url.Parse(config.APIUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL of Chronosphere URL: %v", err)
	}

	var schemes []string
	switch apiURL.Scheme {
	case "https":
		schemes = []string{"https"}
	case "http":
		if !config.AllowHTTP {
			return nil, errors.New("the scheme of the API URL is HTTP but the --allow-http flag was not set")
		}
		schemes = []string{"http"}
	default:
		// If the client didn't specify a scheme in the URL just use the API client's default
		// by passing an empty slice to the transport constructor.
	}

	transport := httptransport.New(apiURL.Host, apiURL.Path, schemes)

	if config.InsecureSkipVerify {
		tlsOpts := httptransport.TLSClientOptions{InsecureSkipVerify: true}
		tlsTransport, err := httptransport.TLSTransport(tlsOpts)
		if err != nil {
			return nil, fmt.Errorf("unable to construct TLS transport: %v", err)
		}
		transport.Transport = tlsTransport
	}

	transport.DefaultAuthentication = httptransport.APIKeyAuth(apiTokenHeader, "header", config.APIToken)

	transport.Transport = xswagger.WithRequestIDTrailerTransport(
		withCustomHeaders(config.Component, config.EntityNamespace, transport.Transport),
	)
	transport.Consumers[httpruntime.JSONMime] = xswagger.JSONConsumer()
	transport.Consumers[httpruntime.HTMLMime] = xswagger.TextConsumer()
	transport.Consumers[httpruntime.TextMime] = xswagger.TextConsumer()
	transport.Consumers["*/*"] = xswagger.TextConsumer() // backup, default consumer.

	return transport, nil
}

const userAgentHeader = "User-Agent"

// CustomHeaderTransport is a RoundTripper that adds a custom headres to all requests
// for example: User-Agent, and Chrono-Entity-Namespace
type CustomHeaderTransport struct {
	agent           string
	entityNamespace string
	Rt              http.RoundTripper
}

func withCustomHeaders(component Component, entityNamespace string, rt http.RoundTripper) http.RoundTripper {
	if rt == nil {
		rt = http.DefaultTransport
	}

	return CustomHeaderTransport{
		Rt: rt,
		agent: fmt.Sprintf("%s/%v-%v (%s; %s; %s)",
			component,
			buildinfo.Version,
			buildinfo.SHA,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH,
		),
		entityNamespace: entityNamespace,
	}
}

// RoundTrip implements the RoundTripper interface.
func (c CustomHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(userAgentHeader, c.agent)
	if c.entityNamespace != "" {
		req.Header.Set("Chrono-Entity-Namespace", c.entityNamespace)
	}
	return c.Rt.RoundTrip(req)
}
