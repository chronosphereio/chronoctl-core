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

package swagger

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
)

const requestIDHeader = "Chrono-Request-Id"

// RequestIDer is an interface for types that have a request ID.
type RequestIDer interface {
	RequestID() string
}

// RequestIDBody wraps an io.ReadCloser, essentially an http body, with a request ID.
type RequestIDBody struct {
	io.ReadCloser
	requestID string
}

// RequestID returns the requestID for the request.
func (r RequestIDBody) RequestID() string {
	return r.requestID
}

// RequestIDTrailerTransport pulls out the request ID from response trailer and
// overwrites the (http.Response).Body to be a RequestIDBody.
type RequestIDTrailerTransport struct {
	RT http.RoundTripper
}

// RoundTrip implements the http.RoundTripper interface. It extracts the request ID
// from the response headers if it is available.
func (r RequestIDTrailerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := r.RT.RoundTrip(req)
	if err != nil {
		return res, err
	}

	if requestID := res.Header.Get(requestIDHeader); requestID != "" {
		res.Body = RequestIDBody{
			ReadCloser: res.Body,
			requestID:  requestID,
		}
	}
	return res, nil
}

// WithRequestIDTrailerTransport returns a new RequestIDTrailerTransport that wraps
// the given http.RoundTripper. If rt is nil, http.DefaultTransport is used.
func WithRequestIDTrailerTransport(rt http.RoundTripper) http.RoundTripper {
	r := RequestIDTrailerTransport{}
	if rt == nil {
		rt = http.DefaultTransport
	}
	r.RT = rt
	return r
}

func appendRequestIDToSwaggerError(data any, requestID string) {
	modelType := reflect.TypeOf(data)
	val := reflect.ValueOf(data)
	if !strings.Contains(modelType.String(), "APIError") && !strings.Contains(modelType.String(), "RuntimeError") {
		return
	}
	message := val.Elem().FieldByName("Message")
	if !message.IsValid() {
		return
	}
	message.SetString(fmt.Sprintf("%v (request_id=%s)", message.String(), requestID))
}
