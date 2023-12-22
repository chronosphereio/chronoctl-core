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

// Package swagger contains generated code from the Swagger spec.
package swagger

import (
	"encoding/json"
	"io"

	"github.com/go-openapi/runtime"
)

// JSONConsumer is a fork of https://github.com/go-openapi/runtime/blob/master/json.go
func JSONConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		dec := json.NewDecoder(reader)
		dec.UseNumber() // preserve number formats

		if err := dec.Decode(data); err != nil {
			return err
		}

		if ider, ok := reader.(RequestIDer); ok {
			appendRequestIDToSwaggerError(data, ider.RequestID())
		}
		return nil
	})
}
