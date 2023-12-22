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

// Package ptr contains functions to create references from constant values,
// typically for representing optional values.
package ptr

import (
	"time"
)

// Wrap converts a value into a pointer.
func Wrap[T any](v T) *T {
	return &v
}

// Unwrap derefences a pointer, or returns the zero-value of T if the pointer is
// nil.
func Unwrap[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}

// String returns a string pointer.
func String(s string) *string { return &s }

// ToString is a convenience method for dereferencing a string pointer to a
// string. "" is returned if the pointer is nil.
func ToString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

// StringEqual compares two string pointers for equality.
func StringEqual(p1, p2 *string) bool {
	return p1 == p2 || (p1 != nil && p2 != nil && *p1 == *p2)
}

// Int returns an int pointer.
func Int(n int) *int { return &n }

// IntEqual compares two int pointers for equality.
func IntEqual(p1, p2 *int) bool {
	return p1 == p2 || (p1 != nil && p2 != nil && *p1 == *p2)
}

// UInt returns a uint pointer.
func UInt(n uint) *uint { return &n }

// ToUInt dereferences a uint pointer to a uint. If the pointer is nil, then returns 0.
func ToUInt(n *uint) uint {
	if n == nil {
		return 0
	}
	return *n
}

// Int32 returns an int32 pointer
func Int32(n int32) *int32 { return &n }

// Int32Equal compares two int32 pointers for equality.
func Int32Equal(p1, p2 *int32) bool {
	return p1 == p2 || (p1 != nil && p2 != nil && *p1 == *p2)
}

// Int64 returns an int64 pointer
func Int64(n int64) *int64 { return &n }

// Int64Equal compares two int64 pointers for equality.
func Int64Equal(p1, p2 *int64) bool {
	return p1 == p2 || (p1 != nil && p2 != nil && *p1 == *p2)
}

// Bool returns a bool pointer.
func Bool(b bool) *bool { return &b }

// BoolEqual compares two bool pointers for equality.
func BoolEqual(p1, p2 *bool) bool {
	return p1 == p2 || (p1 != nil && p2 != nil && *p1 == *p2)
}

// Float64 returns a float64 pointer.
func Float64(n float64) *float64 { return &n }

// Float64Equal compares two float64 pointers for equality.
func Float64Equal(p1, p2 *float64) bool {
	return p1 == p2 || (p1 != nil && p2 != nil && *p1 == *p2)
}

// Time returns a time pointer.
func Time(n time.Time) *time.Time { return &n }

// TimeEqual compares two Time pointers for equality
func TimeEqual(p1, p2 *time.Time) bool {
	return p1 == p2 || (p1 != nil && p2 != nil && *p1 == *p2)
}

// Duration returns a duration pointer.
func Duration(n time.Duration) *time.Duration { return &n }

// DurationEqual compares two Duration pointers for equality
func DurationEqual(p1, p2 *time.Duration) bool {
	return p1 == p2 || (p1 != nil && p2 != nil && *p1 == *p2)
}
