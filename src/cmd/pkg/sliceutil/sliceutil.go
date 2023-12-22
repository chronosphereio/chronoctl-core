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

// Package sliceutil contains utility functions for working with slices.
package sliceutil

// Map converts a slice of one object into another using a function
func Map[X any, Y any](xs []X, f func(X) Y) []Y {
	if xs == nil {
		return nil
	}
	ys := make([]Y, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// MapErr cnverts a slice of one object into another using a function that returns an error.
func MapErr[X any, Y any](xs []X, f func(X) (Y, error)) ([]Y, error) {
	if xs == nil {
		return nil, nil
	}
	ys := make([]Y, 0, len(xs))
	for _, x := range xs {
		y, err := f(x)
		if err != nil {
			return nil, err
		}
		ys = append(ys, y)
	}
	return ys, nil
}
