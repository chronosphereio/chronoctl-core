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
	"errors"
	"io"
	"os"
	"testing"

	"github.com/chronosphereio/chronoctl-core/src/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type nopWriteCloser struct {
	io.Writer
}

func (nopWriteCloser) Close() error {
	return nil
}

type testFileOps struct {
	fileNames       []string
	shouldMkDirFail bool
}

// MkDirAll is a mock implementation of os.MkDirAll that can simulate failure
func (o *testFileOps) MkDirAll(path string, perm os.FileMode) error {
	if o.shouldMkDirFail {
		return errors.New("test fail to make dir")
	}
	return nil
}

// OpenFile is a mock implementation of FileOperation.OpenFile
func (o *testFileOps) OpenFile(name string, flag int, perm os.FileMode) (io.WriteCloser, error) {
	o.fileNames = append(o.fileNames, name)
	return nopWriteCloser{io.Discard}, nil
}

var (
	testYAMLObjWriter = NewYAMLObjectWriter()
	testJSONObjWriter = NewJSONObjectWriter(false /* useArray */)
	writersMap        = map[string]ObjectWriter{
		JSONFormat: testJSONObjWriter,
		YAMLFormat: testYAMLObjWriter,
	}
)

func TestInitWriterManager(t *testing.T) {
	tests := []struct {
		name              string
		flags             *Flags
		mockFileOps       *testFileOps
		format            string
		wantWriterManager WriterManager
		wantFilePath      string
		wantErr           string
	}{
		{
			name: "with empty directory",
			flags: &Flags{
				Format:              JSONFormat,
				OutputDirectory:     "",
				CreateFilePerObject: false,
				Writers:             writersMap,
			},
			mockFileOps:  &testFileOps{},
			format:       JSONFormat,
			wantFilePath: "",
			wantWriterManager: WriterManager{
				FilePerObj:   false,
				Dir:          "",
				ObjectWriter: testJSONObjWriter,
				Format:       JSONFormat,
				closeFn:      nil,
				opened:       map[string]bool{},
				objectWriterOpts: ObjectWriterOpts{
					multiObjects: true,
				},
			},
		},
		{
			name: "with directory but no file per object",
			flags: &Flags{
				Format:              YAMLFormat,
				OutputDirectory:     "/path/to/dir",
				CreateFilePerObject: false,
				Writers:             writersMap,
			},
			mockFileOps:  &testFileOps{},
			format:       YAMLFormat,
			wantFilePath: "/path/to/dir/output.yaml",
			wantWriterManager: WriterManager{
				FilePerObj:   false,
				Dir:          "/path/to/dir",
				ObjectWriter: testYAMLObjWriter,
				Format:       YAMLFormat,
				opened:       map[string]bool{},
				objectWriterOpts: ObjectWriterOpts{
					multiObjects: true,
				},
			},
		},
		{
			name: "with directory and file per object is true",
			flags: &Flags{
				Format:              YAMLFormat,
				OutputDirectory:     "/path/to/dir",
				CreateFilePerObject: true,
				Writers:             writersMap,
			},
			mockFileOps:  &testFileOps{},
			format:       YAMLFormat,
			wantFilePath: "",
			wantWriterManager: WriterManager{
				FilePerObj:   true,
				Dir:          "/path/to/dir",
				ObjectWriter: testYAMLObjWriter,
				Format:       YAMLFormat,
				opened:       map[string]bool{},
				objectWriterOpts: ObjectWriterOpts{
					multiObjects: false,
				},
			},
		},
		{
			name: "failed making directory",
			flags: &Flags{
				OutputDirectory: "/path/to/dir",
			},
			mockFileOps: &testFileOps{shouldMkDirFail: true},
			wantErr:     "unable to make directory: test fail to make dir",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			writerManager, err := NewWriterManager(tt.flags, os.Stdout, tt.mockFileOps)
			if tt.wantErr != "" {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			// setting file and writer to nil here since we want to compare file paths instead of the actual file
			writerManager.closeFn = nil
			writerManager.W = nil
			if tt.wantFilePath != "" {
				assert.Equal(t, []string{tt.wantFilePath}, tt.mockFileOps.fileNames)
			} else {
				assert.Empty(t, tt.mockFileOps.fileNames)
			}

			writerManager.FileOps = nil
			assert.Equal(t, tt.wantWriterManager, *writerManager)
		})
	}
}

func TestWriteObject(t *testing.T) {
	tests := []struct {
		name          string
		objs          []types.Object
		flags         *Flags
		outputDir     string
		filePerObj    bool
		wantFilePaths []string
	}{
		{
			name: "file per object is true, single file",
			objs: []types.Object{testGoodObject{
				Foo: "foo",
				Bar: 42,
			}},
			outputDir:     "/path/to/dir",
			filePerObj:    true,
			wantFilePaths: []string{"/path/to/dir/type_kind_slug_foo.yaml"},
		},
		{
			name: "file per object is true, multiple files with clash",
			objs: []types.Object{
				testGoodObject{
					Foo: "foo",
					Bar: 42,
				},
				testGoodObject{
					Foo: "foo", // clashes with above
					Bar: 42,
				},
				testGoodObject{
					Foo: "bar", // clashes with above
					Bar: 42,
				},
			},
			outputDir:  "/path/to/dir",
			filePerObj: true,
			wantFilePaths: []string{
				"/path/to/dir/type_kind_slug_foo.yaml",
				"/path/to/dir/type_kind_slug_foo_2.yaml",
				"/path/to/dir/type_kind_slug_bar.yaml",
			},
		},
		{
			name: "file per object is false",
			objs: []types.Object{testGoodObject{
				Foo: "foo",
				Bar: 42,
			}},
			outputDir:     "",
			filePerObj:    false,
			wantFilePaths: nil, // no open expected
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flags := NewFlags()
			flags.OutputDirectory = tt.outputDir
			flags.CreateFilePerObject = tt.filePerObj

			fileOps := &testFileOps{}
			writerManager, err := NewWriterManager(flags, io.Discard, fileOps)
			require.NoError(t, err, "NewWriterManager failed")

			for _, obj := range tt.objs {
				require.NoError(t, writerManager.WriteObject(obj), "WriteObject failed")
			}
			assert.Equal(t, tt.wantFilePaths, fileOps.fileNames, "files mismatch")
		})
	}
}

func (testGoodObject) Type() types.TypeMeta {
	return types.TypeMeta{Kind: "type-kind"}
}

func (testGoodObject) Description() string {
	return "testObjectDescription"
}

func (o testGoodObject) Identifier() string {
	return "slug-" + o.Foo
}
