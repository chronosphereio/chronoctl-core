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

package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStructA struct {
	A string
}

type testStructB struct {
	B string
}

func TestMap(t *testing.T) {
	tests := []struct {
		name string
		in   []testStructA
		out  []testStructB
	}{
		{
			name: "nil",
			in:   nil,
			out:  nil,
		},
		{
			name: "convert",
			in:   []testStructA{{A: "a"}, {A: "b"}, {A: "c"}},
			out:  []testStructB{{B: "a"}, {B: "b"}, {B: "c"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := Map(test.in, func(a testStructA) testStructB {
				return testStructB{B: a.A}
			})
			assert.Equal(t, test.out, b)
		})
	}
}

func TestMapErr(t *testing.T) {
	tests := []struct {
		name     string
		in       []testStructA
		out      []testStructB
		hasError bool
	}{
		{
			name:     "nil",
			in:       nil,
			out:      nil,
			hasError: false,
		},
		{
			name:     "convert",
			in:       []testStructA{{A: "a"}, {A: "b"}, {A: "c"}},
			out:      []testStructB{{B: "a"}, {B: "b"}, {B: "c"}},
			hasError: false,
		},
		{
			name:     "error",
			in:       []testStructA{{A: "a"}, {A: "b"}, {A: "c"}},
			out:      []testStructB{{B: "a"}, {B: "b"}, {B: "c"}},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b, err := MapErr(test.in, func(a testStructA) (testStructB, error) {
				if test.hasError {
					return testStructB{}, assert.AnError
				}
				return testStructB{B: a.A}, nil
			})
			if test.hasError {
				assert.Error(t, assert.AnError)
			} else {
				assert.Equal(t, test.out, b)
				assert.NoError(t, err)
			}
		})
	}
}
