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

package buildinfo

// All of these vars are replaced at link time using ldflags.
var (
	// Version is a valid semantic version
	Version = "unknown"
	// SHA contains the git SHA of the commit
	SHA = "unknown"
	// Date is the date of the build
	Date = "unknown"
)
