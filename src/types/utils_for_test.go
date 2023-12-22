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

// testObject is an internal API object that is used across multiple test files.
type testObject struct {
	TypeMeta `json:",inline"`

	Name string `json:"name"`
}

var _ Object = (*testObject)(nil)

// Description returns a friendly description of the object.
func (t *testObject) Description() string {
	return TypeDescription(t)
}

// Identifier returns a blank string
func (t *testObject) Identifier() string {
	return ""
}
