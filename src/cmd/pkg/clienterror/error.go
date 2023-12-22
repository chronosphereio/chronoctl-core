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

package clienterror

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/pkg/errors"

	configunstablemodels "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/models"
	configv1models "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
)

// Error contains the structured error data from a server to the API client.
type Error struct {
	cause error

	Operation string `json:"operation,omitempty"`
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
}

var _ error = &Error{}

// Error returns the error string
func (e *Error) Error() string {
	return fmt.Sprintf("%s failed with HTTP code %d: %s", e.Operation, e.Code, e.Message)
}

// Cause is the error returned by the Swagger client.
func (e *Error) Cause() error {
	return e.cause
}

// Wrap presents errors from the Swagger-generated client in a more user-friendly format.
// This returns a wrapped Error when possible.
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	code, ok := getCode(err)
	if !ok {
		return errors.WithStack(err)
	}
	msg, ok := getMessage(err)
	if !ok {
		return errors.WithStack(err)
	}
	return errors.WithStack(&Error{
		cause:     err,
		Operation: getOperation(err),
		Code:      code,
		Message:   msg,
	})
}

type codeError interface {
	Code() int
}

func getCode(err error) (int, bool) {
	if e, ok := err.(codeError); ok {
		return e.Code(), true
	}
	return 0, false
}

// Structured errors:

type configV1APIError interface {
	GetPayload() *configv1models.APIError
}

type configUnstableAPIError interface {
	GetPayload() *configunstablemodels.APIError
}

// Generic errors:
type configV1GenericError interface {
	GetPayload() configv1models.GenericError
}

type configUnstableGenericError interface {
	GetPayload() configunstablemodels.GenericError
}

func getMessage(err error) (string, bool) {
	if e, ok := err.(configV1APIError); ok {
		return e.GetPayload().Message, true
	}
	if e, ok := err.(configUnstableAPIError); ok {
		return e.GetPayload().Message, true
	}
	if e, ok := err.(configV1GenericError); ok {
		return tryMarshal(e.GetPayload())
	}
	if e, ok := err.(configUnstableGenericError); ok {
		return tryMarshal(e.GetPayload())
	}

	return "", false
}

// Hack to parse the verb + URL out of a Swagger error string, e.g
// "[POST /api/v1/config/notifiers][400] createNotifierBadRequest"
var errOperationRegex = regexp.MustCompile(`^\[(\w+) ([^\]]+)\]`)

func getOperation(err error) string {
	matches := errOperationRegex.FindStringSubmatch(err.Error())
	if len(matches) == 3 {
		return fmt.Sprintf("%s %s", matches[1], matches[2])
	}
	return err.Error()
}

func tryMarshal(v any) (string, bool) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", false
	}
	return string(b), true
}

// StatusCode returns the HTTP status code of an error, if there is one.
func StatusCode(err error) int {
	if err == nil {
		return 0
	}

	err = Wrap(err)
	var clientErr *Error
	if errors.As(err, &clientErr) {
		return clientErr.Code
	}

	return 0
}

// IsNotFound A simple helper to get a status code for an error and see if it's a not found status code
func IsNotFound(err error) bool {
	return StatusCode(err) == http.StatusNotFound
}

func getAPIErrorCode(err error) (code int32, ok bool) {
	if p, ok := err.(configV1APIError); ok {
		return p.GetPayload().Code, true
	}
	if p, ok := err.(configUnstableAPIError); ok {
		return p.GetPayload().Code, true
	}
	return 0, false
}

const entityValidationFailedCode = 400008

// IsEntityValidationFailed checks if the entity validation failed.
func IsEntityValidationFailed(err error) bool {
	errorCode, ok := getAPIErrorCode(err)
	return ok && errorCode == entityValidationFailedCode
}
