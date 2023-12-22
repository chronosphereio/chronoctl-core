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

// Calculation contains all the info needed to calculate a request's page size.
type Calculation struct {
	// GotItems is how many items we currently have received.
	GotItems int
	// MaxItems is how many items we want in total.
	MaxItems int
	// MaxPageSize is the maximum number of items we want to query for per request.
	MaxPageSize int
}

// CalculatePageSize calculates the desired page size for paginated queries.
func CalculatePageSize(c Calculation) int {
	// if no limit was provided, always fetch the maximum items per page
	if c.MaxItems <= 0 {
		return c.MaxPageSize
	}

	remaining := c.MaxItems - c.GotItems
	// if we have more than a single page of results to fetch, request a full page
	if remaining > c.MaxPageSize {
		return c.MaxPageSize
	}

	// request exactly the number of items left so the returned pagination token is correct
	return remaining
}
