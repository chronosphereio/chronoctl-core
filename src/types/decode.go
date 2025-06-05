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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/chronosphereio/chronoctl-core/src/thirdparty/yaml"
)

var (
	// isJSONRegexp determines if a string might be a JSON object.
	isJSONRegexp = regexp.MustCompile(`^\s*({|\[)`)

	// isJSONRegexp determines if a string might be a JSON object.
	isJSONArrayRegexp = regexp.MustCompile(`^\s*(\[)`)

	// yamlDocSuffix matches the strings ending with an empty document.
	yamlDocSuffix = regexp.MustCompile(`\n---(\n)*$`)
)

// MustDecodeSingleObject reads a single object and returns an error if there are more than one
// object in the input stream.
func MustDecodeSingleObject[T Object](r io.Reader, permissiveParsing bool) (T, error) {
	var empty T
	objs, err := decodeWithRegistry(r, Registry, permissiveParsing)
	if err != nil {
		return empty, err
	}
	if len(objs) == 0 {
		return empty, fmt.Errorf("input does not contain any items")
	}
	if len(objs) > 1 {
		return empty, fmt.Errorf("input contains more than one item")
	}
	castedObj, ok := objs[0].(T)
	if !ok {
		return empty, WrongObjectErr(empty, objs[0])
	}
	return castedObj, nil
}

// Decode decodes objects from `r` until it encounters an EOF. It uses the API registry to
// discover the types of the objects it decodes.
func Decode(r io.Reader, permissiveParsing bool) ([]Object, error) {
	return decodeWithRegistry(r, Registry, permissiveParsing)
}

func decodeWithRegistry(r io.Reader, reg ObjectRegistry, permissiveParsing bool) ([]Object, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}

	if len(data) == 0 {
		return nil, nil
	}

	isJSON := isJSONRegexp.Match(data)
	if !isJSON {
		// If the data ends with an empty document ("---" followed by nothing),
		// remove the last empty document.
		// This prevents the parser from decoding an empty string into a struct.
		data = yamlDocSuffix.ReplaceAll(data, nil)
	} else if isJSONArrayRegexp.Match(data) {
		data, err = convertJSONToJSONL(data)
		if err != nil {
			return nil, err
		}
	}

	type decoder interface {
		Decode(v interface{}) error
	}

	newDecoder := func(strict bool) decoder {
		if isJSON {
			jsonDecoder := newJSONDecoder(data)
			if strict {
				jsonDecoder.decoder.DisallowUnknownFields()
			}
			return jsonDecoder
		}
		yamlDecoder := yaml.NewDecoder(bytes.NewBuffer(data))
		yamlDecoder.KnownFields(strict)
		return yamlDecoder
	}

	// Create two decoders for the input to deal with polymorphic input.
	// The first decoder unmarshals docs to TypeMeta to determine the full Go struct type,
	// then the second decoder unmarshals into the final struct.
	// The second decoder is configurable (via the --strict tag, which defaults to true) to reject YAML/JSON with invalid fields
	var objs []Object
	typeDecoder := newDecoder(false)
	objDecoder := newDecoder(!permissiveParsing)
	for {
		// Unmarshal the next document into TypeMeta
		var meta TypeMeta
		if err := typeDecoder.Decode(&meta); err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("unable to decode object: %w", err)
		}

		// Determine the full Go struct to use for the type
		obj, err := newObjectFromMeta(reg, meta)
		if err != nil {
			return nil, err
		}

		// Unmarshal the *same* document into the full struct
		if err := objDecoder.Decode(obj); err != nil {
			return nil, fmt.Errorf("unable to decode object: %w", err)
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

func convertJSONToJSONL(data []byte) ([]byte, error) {
	var messages []json.RawMessage
	if err := json.Unmarshal(data, &messages); err != nil {
		return nil, err
	}

	var (
		out     = bytes.NewBuffer(data)
		encoder = json.NewEncoder(out)
	)
	out.Reset()
	for _, message := range messages {
		if err := encoder.Encode(message); err != nil {
			return nil, err
		}
	}
	return out.Bytes(), nil
}

// newJSONDecoder returns a decoder to unmarshal the given JSON data.
// The decoer wraps a json.Decoder, including line numbers in syntax parsing errors.
func newJSONDecoder(data []byte) *jsonDecoder {
	return &jsonDecoder{
		data:    data,
		decoder: json.NewDecoder(bytes.NewReader(data)),
	}
}

type jsonDecoder struct {
	data    []byte
	decoder *json.Decoder
}

func (j *jsonDecoder) Decode(v interface{}) error {
	err := j.decoder.Decode(v)

	if syntaxErr, ok := err.(*json.SyntaxError); ok {
		data, readErr := io.ReadAll(j.decoder.Buffered())
		if readErr != nil {
			log.Printf("reading stream failed: %v", readErr)
		}
		js := string(data)

		// If contents from io.Reader are not complete,
		// use the original raw data to prevent panic
		if int64(len(js)) <= syntaxErr.Offset {
			js = string(j.data)
		}

		start := strings.LastIndex(js[:syntaxErr.Offset], "\n") + 1
		line := strings.Count(js[:start], "\n")

		return fmt.Errorf("json: line %d: %w", line, syntaxErr)
	}

	return err
}

func newObjectFromMeta(reg ObjectRegistry, meta TypeMeta) (Object, error) {
	t, ok := reg[meta]
	if !ok {
		return nil, fmt.Errorf(
			"no registered type found for object with type with Kind: %v and API version: %v",
			meta.Kind, meta.APIVersion,
		)
	}

	return reflect.New(t).Interface().(Object), nil //nolint:errcheck
}
