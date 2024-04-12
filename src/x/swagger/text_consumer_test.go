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

package swagger_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/client/operations"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
	"github.com/chronosphereio/chronoctl-core/src/x/swagger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTextConsumer(t *testing.T) {
	tests := []struct {
		name         string
		inputData    interface{}
		expectedData interface{}
		body         string
		expectedErr  error
	}{
		{
			name:         "Test 500 response body with unexpected data type",
			inputData:    "",
			expectedData: "",
			body:         "500 Internal Server Error",
			expectedErr:  fmt.Errorf("received an unexpected response type as text, response body: 500 Internal Server Error"),
		},
		{
			name:      "Test 500 response body with a RuntimError",
			inputData: &models.APIError{},
			expectedData: &models.APIError{
				Code:    2,
				Message: "500 Internal Server Error",
			},
			body:        "500 Internal Server Error",
			expectedErr: nil,
		},
		{
			name:      "Test json response body with a RuntimError",
			inputData: &models.APIError{},
			expectedData: &models.APIError{
				Code:    2,
				Message: `{"code": 16, "message":"InternalServer", "details":"An internal server error has occured"}`,
			},
			body:        `{"code": 16, "message":"InternalServer", "details":"An internal server error has occured"}`,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			consumer := swagger.TextConsumer()
			input := tt.inputData
			reader := bytes.NewBuffer([]byte(tt.body))
			err := consumer.Consume(reader, input)
			require.Equal(t, err, tt.expectedErr)
			assert.Equal(t, input, tt.expectedData)
		})
	}
}

func TestTextConsumerDummyServer(t *testing.T) {
	tests := []struct {
		name          string
		mimeType      string
		content       string
		message       string
		noResp        bool
		unexpectedErr bool
	}{
		{
			name:     "text mime type",
			mimeType: "text",
			content:  "500 internal server error",
			message:  "500 internal server error",
		},
		{
			name:     "text/html mime type with json",
			mimeType: "text/html",
			content:  `{"code":2,"message":"unknown error"}`,
			message:  `{"code":2,"message":"unknown error"}`,
		},
		{
			name:     "text/html mime type",
			mimeType: "text/html",
			content:  "500 internal server error",
			message:  "500 internal server error",
		},
		{
			name:     "unknown mime type",
			mimeType: "unkown",
			content:  "500 internal server error",
			message:  "500 internal server error",
		},
		{
			// This test case should hit the JSONConsumer type and succeed.
			name:     "application/json mime type with well formed json",
			mimeType: "application/json",
			content:  `{"code":2,"message":"unknown error"}`,
			message:  "unknown error",
		},
		{
			// Regression test for ch-16822.
			name:          "no response should not panic",
			noResp:        true,
			unexpectedErr: true,
			message:       "EOF",
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// setup a dummy server.
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.noResp {
					panic(http.ErrAbortHandler)
				}

				w.Header().Set("Content-Type", tt.mimeType)
				w.WriteHeader(500)
				_, err := w.Write([]byte(tt.content))
				require.NoError(t, err)
			}))
			defer server.Close()

			clientFlags := client.Flags{
				APIToken:  "fake",
				APIUrl:    server.URL,
				AllowHTTP: true,
			}

			client, err := clientFlags.ConfigV1Client()
			require.NoError(t, err)

			_, err = client.ListBuckets(&operations.ListBucketsParams{
				Context: ctx,
			})

			bucketDefault, ok := err.(*operations.ListBucketsInternalServerError)
			if tt.unexpectedErr {
				require.False(t, ok)
				assert.Contains(t, err.Error(), tt.message)
				return
			}
			assert.Equal(t, bucketDefault.Code(), 500)
			fmt.Println(bucketDefault.GetPayload())
			assert.Equal(t, bucketDefault.GetPayload().Code, int32(2))
			assert.Equal(t, bucketDefault.GetPayload().Message, tt.message)
		})
	}
}
