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
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"
)

const codeUnknown = 2

// TextConsumer creates a new text consumer.
//
// Specifically we attempt to deal with the case here where nginx returns a 500
// that is HTML, so we can't unmarshal into our RuntimeError object. In the case
// HTML is returned, we set the error's message field to the HTML body.
func TextConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		if reader == nil {
			return fmt.Errorf("unexpected response, missing response body")
		}

		b, err := ioutil.ReadAll(reader)
		if err != nil {
			return errors.Wrap(err, "could not read unexpected response body")
		}

		if !setSwaggerErrorMessage(data, string(b)) {
			// We don't know what to do with our response body.
			return fmt.Errorf("received an unexpected response type as text, response body: %s", b)
		}

		if ider, ok := reader.(RequestIDer); ok {
			appendRequestIDToSwaggerError(data, ider.RequestID())
		}
		return nil
	})
}

// setSwaggerErrorMessage attempts to set a message received from a TextConsumer
// into a swagger Runtime/API Error.
func setSwaggerErrorMessage(data any, message string) bool {
	modelType := reflect.TypeOf(data)
	val := reflect.ValueOf(data)
	if !strings.Contains(modelType.String(), "APIError") && !strings.Contains(modelType.String(), "RuntimeError") {
		return false
	}

	messageField := val.Elem().FieldByName("Message")
	if !messageField.IsValid() {
		return false
	}
	messageField.SetString(message)

	codeField := val.Elem().FieldByName("Code")
	if !codeField.IsValid() {
		return false
	}
	codeField.SetInt(codeUnknown)

	return true
}
