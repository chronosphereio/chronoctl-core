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

package types

import "fmt"

// WrongObjectErr returns an error indicating that the object the CLI is formatting
// is not what was expected.
func WrongObjectErr(want, actual Object) error {
	return fmt.Errorf("object type is wrong, expected %v, found %v",
		userTypeForObject(want),
		userTypeForObject(actual),
	)
}

// userTypeForObject uses the TypeMeta for the object's type using the
// registry. It doesn't use obj.Type() since obj may be nil.
func userTypeForObject(obj Object) string {
	if t, ok := Registry.typeForObject(obj); ok {
		return t.String()
	}
	// Fallback to the Go type if we can't find the object in the registry.
	return fmt.Sprintf("%T", obj)
}
