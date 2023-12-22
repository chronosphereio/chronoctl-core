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
	"github.com/stretchr/testify/require"
)

func TestRegisterObject(t *testing.T) {
	tests := []struct {
		name     string
		obj      Object
		registry ObjectRegistry
		wantErr  string
	}{
		{
			name: "new valid object",
			obj: &testObject{
				TypeMeta: TypeMeta{
					Kind:       "Test",
					APIVersion: "v1",
				},
			},
			registry: make(ObjectRegistry),
		},
		{
			name: "object with same type exists already",
			obj: &testObject{
				TypeMeta: TypeMeta{
					Kind:       "Test",
					APIVersion: "v1",
				},
			},
			registry: func() ObjectRegistry {
				var (
					r = make(ObjectRegistry)
					o = &testObject{
						TypeMeta: TypeMeta{
							Kind:       "Test",
							APIVersion: "v1",
						},
					}
				)
				err := r.registerType(o.TypeMeta, o)
				if err != nil {
					panic(err)
				}
				return r
			}(),
		},
		{
			name: "Kind is missing",
			obj: &testObject{
				TypeMeta: TypeMeta{
					APIVersion: "v1",
				},
			},
			registry: make(ObjectRegistry),
			wantErr:  errKindMissing.Error(),
		},
		{
			name: "API version is missing",
			obj: &testObject{
				TypeMeta: TypeMeta{
					Kind: "Test",
				},
			},
			registry: make(ObjectRegistry),
			wantErr:  errAPIVersionMissing.Error(),
		},
		{
			name:     "object is not a pointer",
			obj:      nonPointerObject{},
			registry: make(ObjectRegistry),
			wantErr:  errInvalidType.Error(),
		},
		{
			name: "object is not a struct",
			obj: func() Object {
				o := nonStructObject(0)
				return &o
			}(),
			registry: make(ObjectRegistry),
			wantErr:  errInvalidType.Error(),
		},
		{
			name: "different object is already registered",
			obj: &testObject{
				TypeMeta: TypeMeta{
					Kind:       "Test",
					APIVersion: "v1",
				},
			},
			registry: func() ObjectRegistry {
				var (
					r = make(ObjectRegistry)
					o = &testObject2{
						TypeMeta: TypeMeta{
							Kind:       "Test",
							APIVersion: "v1",
						},
					}
				)
				err := r.registerType(o.TypeMeta, o)
				if err != nil {
					panic(err)
				}
				return r
			}(),
			wantErr: "Double registration of different types for v1/Test: old=github.com/chronosphereio/chronoctl-core/src/types.testObject2, new=github.com/chronosphereio/chronoctl-core/src/types.testObject",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.registry.registerType(tt.obj.Type(), tt.obj)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			_, ok := tt.registry[tt.obj.Type()]
			assert.True(t, ok)
		})
	}
}

func TestMustRegisterObjectPanics(t *testing.T) {
	obj1 := &testObject{
		TypeMeta: TypeMeta{
			Kind:       "Test",
			APIVersion: "v1",
		},
	}

	obj2 := &testObject2{
		TypeMeta: TypeMeta{
			Kind:       "Test",
			APIVersion: "v1",
		},
	}

	assert.NotPanics(t, func() { MustRegisterObject(obj1.TypeMeta, obj1) })
	assert.Panics(t, func() { MustRegisterObject(obj2.TypeMeta, obj2) })
}

type nonPointerObject struct{}

func (o nonPointerObject) Type() TypeMeta { return TypeMeta{Kind: "Fake", APIVersion: "v1"} }

func (o nonPointerObject) Description() string { return TypeDescription(o) }

func (o nonPointerObject) Identifier() string { return "" }

type nonStructObject int

func (o *nonStructObject) Type() TypeMeta { return TypeMeta{Kind: "Fake", APIVersion: "v1"} }

func (o *nonStructObject) Description() string { return TypeDescription(o) }

func (o *nonStructObject) Identifier() string { return "" }

type testObject2 struct {
	TypeMeta
}

func (o *testObject2) Description() string { return TypeDescription(o) }

func (o *testObject2) Identifier() string { return "" }
