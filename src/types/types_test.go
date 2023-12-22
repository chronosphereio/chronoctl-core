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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescription(t *testing.T) {
	fooObj := &testObject{
		TypeMeta: TypeMeta{Kind: "Foo"},
	}

	assert.Equal(t, `Foo{}`, TypeDescription(fooObj))
	assert.Equal(t, `Foo{}`, TypeDescription(fooObj, "id", ""))
	assert.Equal(t, `Foo{id="hello"}`, TypeDescription(fooObj, "id", "hello"))
	assert.Equal(t, `Foo{id="hello"}`, TypeDescription(fooObj, "id", "hello", "name", ""))
	assert.Equal(t, `Foo{name="what"}`, TypeDescription(fooObj, "id", "", "name", "what"))
	assert.Equal(t, `Foo{id="hello", name="world"}`, TypeDescription(fooObj, "id", "hello", "name", "world"))
}
