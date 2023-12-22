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

// Package timesutil provides utilities for parsing and validating RFC3339 timestamps
package timesutil

import (
	"fmt"
	"time"

	"github.com/karrick/tparse/v2"
)

// ParseTimestamp validates and returns an RFC3339 timestamp. If a invalid RFC3339 is passed in,
// we return an error. If a relative duration is passed in, e.g. -15m, it returns the RFC3339
// timestamp relative to that bit.
func ParseTimestamp(now time.Time, ts string) (string, error) {
	if ts == "" {
		return "", nil
	}

	_, err := time.Parse(time.RFC3339, ts)
	if err == nil {
		return ts, nil
	}

	parsedDuration, err := tparse.AbsoluteDuration(time.Time{}, ts)
	if err == nil {
		if parsedDuration == 0 {
			return "", nil
		}
		return now.Add(parsedDuration).Format(time.RFC3339), nil

	}

	// N.B. The time.Parse cmderrors are particularly useful for users.
	return "", fmt.Errorf("unable to parse \"%s\" as either an RFC3339 timestamp or relative duration", ts)
}
