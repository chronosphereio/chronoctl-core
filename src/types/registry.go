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
	"errors"
	"fmt"
	"reflect"
)

var (
	errKindMissing       = errors.New("cannot register an Object that doesn't have a Kind")
	errAPIVersionMissing = errors.New("cannot register an Object that doesn't have an API version")
	errInvalidType       = errors.New("can only register types that are pointers to structs")
)

// ObjectRegistry is a map from the type information of an object to it's actual Go type.
type ObjectRegistry map[TypeMeta]reflect.Type

// Registry contains the type information for all objects in the Chronosphere API.
var Registry = make(ObjectRegistry)

// MustRegisterObject registers a Chronosphere object. `meta` must define the object's Kind
// and API version, be a pointer to a struct, and there must not be a different object with
// the same Kind and API version in Registry already. If any of these conditions aren't met
// the function will panic. MustRegisterObject is intended to be called in init functions to
// register all valid types at startup.
func MustRegisterObject(meta TypeMeta, obj Object) {
	if err := Registry.registerType(meta, obj); err != nil {
		panic(fmt.Sprintf("Unable to register Object: %v.", err))
	}
}

// registerType registers the specified type and object in the ObjectRegistry.
func (r ObjectRegistry) registerType(meta TypeMeta, obj Object) error {
	if meta.Kind == "" {
		return errKindMissing
	}

	if meta.APIVersion == "" {
		return errAPIVersionMissing
	}

	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		return errInvalidType
	}
	t = t.Elem()
	if t.Kind() != reflect.Struct {
		return errInvalidType
	}

	if oldT, found := r[meta]; found && oldT != t {
		return fmt.Errorf(
			"Double registration of different types for %v: old=%v.%v, new=%v.%v",
			meta, oldT.PkgPath(), oldT.Name(), t.PkgPath(), t.Name(),
		)
	}
	r[meta] = t

	return nil
}

func (r ObjectRegistry) typeForObject(obj Object) (TypeMeta, bool) {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for k, v := range r {
		if t == v {
			return k, true
		}
	}

	return TypeMeta{}, false
}
