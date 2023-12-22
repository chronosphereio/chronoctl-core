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

package pagination

import "github.com/chronosphereio/chronoctl-core/src/types"

var _ types.Object = &Result{}

// Token represents a pagination token.
type Token string

// ResultKind represents the result object kind.
const ResultKind = "PageInfo"

// Result is a printable object that shows the pagination kind.
// Implements the types.Object interface as a noop.
type Result struct {
	Kind          string `json:"kind"`
	Message       string `json:"message"`
	NextPageToken Token  `json:"next_page_token"`
	FullCommand   string `json:"full_command"`
}

// Type returns an empty TypeMeta.
func (r *Result) Type() types.TypeMeta { return types.TypeMeta{} }

// Description returns an empty description.
func (r *Result) Description() string { return "" }

// Identifier returns an empty identifier.
func (r *Result) Identifier() string { return "" }
