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

package output

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/chronosphereio/chronoctl-core/src/types"
)

const (
	outputYAML    = "output.yaml"
	outputJSON    = "output.json"
	yamlExtension = ".yaml"
	jsonExtension = ".json"
)

var (
	outputFileNamesMap = map[string]string{
		JSONFormat: outputJSON,
		YAMLFormat: outputYAML,
	}

	extensionsMap = map[string]string{
		JSONFormat: jsonExtension,
		YAMLFormat: yamlExtension,
	}
)

// FileOperation is an interface that abstracts the file operations for the various
// writer implementations of chronoctl
type FileOperation interface {
	MkDirAll(string, os.FileMode) error
	OpenFile(string, int, os.FileMode) (io.WriteCloser, error)
}

// FileOps is the default implementation of FileOperation
type FileOps struct{}

// MkDirAll makes all directories in a path
func (FileOps) MkDirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// OpenFile opens a file
func (FileOps) OpenFile(name string, flag int, perm os.FileMode) (io.WriteCloser, error) {
	return os.OpenFile(name, flag, perm)
}

// WriterManager is a struct that manages the output of chronoctl
type WriterManager struct {
	W                io.Writer
	ObjectWriter     ObjectWriter
	FilePerObj       bool
	Dir              string
	Format           string
	closeFn          func() error
	FileOps          FileOperation
	objectWriterOpts ObjectWriterOpts
	opened           map[string]bool
}

// NewWriterManager creates a new WriterManager
func NewWriterManager(flags *Flags, stdout io.Writer, fileOps FileOperation) (*WriterManager, error) {
	manager := &WriterManager{
		ObjectWriter: flags.Writers[flags.Format],
		Dir:          flags.OutputDirectory,
		FilePerObj:   flags.CreateFilePerObject,
		Format:       flags.Format,
		W:            stdout,
		FileOps:      fileOps,
		objectWriterOpts: ObjectWriterOpts{
			multiObjects: !flags.CreateFilePerObject,
		},
		opened: make(map[string]bool),
	}

	if manager.Dir != "" {
		if err := manager.FileOps.MkDirAll(manager.Dir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("unable to make directory: %v", err)
		}
		if !manager.FilePerObj {
			filePath := filepath.Join(manager.Dir, outputFileNamesMap[manager.Format])
			writer, closer, err := manager.initializeOutputFile(filePath)
			if err != nil {
				return nil, err
			}
			manager.closeFn = closer
			manager.W = writer
		}
	}

	return manager, nil
}

// WriteObject writes an object to the output writer
func (m *WriterManager) WriteObject(obj types.Object) error {
	writer, closeFunc, err := m.getWriterForObj(obj)
	if err != nil {
		return err
	}
	defer closeFunc() // nolint:errcheck,gosec
	return m.ObjectWriter.WriteObject(writer, obj, m.objectWriterOpts)
}

// Close closes the output writer
func (m *WriterManager) Close() error {
	if err := m.ObjectWriter.Close(m.W, m.objectWriterOpts); err != nil {
		return err
	}
	if m.closeFn != nil {
		defer m.closeFn() // nolint:errcheck,gosec
	}
	return nil
}

func (m *WriterManager) initializeOutputFile(filePath string) (io.Writer, func() error, error) {
	f, err := m.FileOps.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create file %v", err)
	}
	return f, f.Close, nil

}

func (m *WriterManager) getWriterForObj(obj types.Object) (io.Writer, func() error, error) {
	if !m.FilePerObj {
		noOpFunc := func() error { return nil }
		return m.W, noOpFunc, nil
	}
	baseFile := types.DisplayName(obj.Type().Kind, obj.Identifier())

	// Ensure that we don't open the same file multiple times by tracking
	// opened files, and using a incrementing suffix to avoid dupes.
	suffix := 1
	uniqueFile := baseFile
	for m.opened[uniqueFile] {
		suffix++
		uniqueFile = baseFile + fmt.Sprintf("_%v", suffix)
	}
	m.opened[uniqueFile] = true

	filePath := filepath.Join(m.Dir, uniqueFile+extensionsMap[m.Format])
	return m.initializeOutputFile(filePath)
}
