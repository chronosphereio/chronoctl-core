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

// Package output provides functionality to write output in different formats.
package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"

	"github.com/chronosphereio/chronoctl-core/src/types"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const (
	// JSONFormat is used to indicate that the output of a command should be in JSON format
	JSONFormat = "json"
	// JSONLFormat is used to indicate that the output of a command should be in JSON format using
	// new line delimited JSON.
	JSONLFormat = "jsonl"
	// YAMLFormat is used to indicate that the output of a command should be in YAML format
	YAMLFormat = "yaml"
	// TableFormat is used to indicate that the output of a command should be printed as a table
	TableFormat = "table"
)

// Flags is a struct that contains the configuration for the output flags
type Flags struct {
	// Format is which Format to use when writing output.
	Format string
	// Writers maps a Format to a specific writer, allowing commands to register
	// their own output Format, e.g. a table view.
	Writers map[string]ObjectWriter
	// OutputDirectory is the directory where the output is written to.
	OutputDirectory string
	// CreateFilePerObject indicates whether to create a file per object.
	CreateFilePerObject bool
	// noFilePerObjectArg indicates whether the command supports the --create-file-per-object flag.
	noFilePerObjectArg bool
	// noOutputDirArg indicates that we should not include the output dir fllag.
	noOutputDirArg bool
}

// OutputOption is a function that configures output options for writers
type OutputOption func(*Flags) //nolint:revive // ignore repetitive name warning

// WithWriter registers a writer with the given Format name.
func WithWriter(formatName string, writer ObjectWriter) OutputOption {
	return func(o *Flags) {
		o.Writers[formatName] = writer
	}
}

// WithoutCreateFilePerObject is an option that configuration a writer to not
// output each object to a separate file.
func WithoutCreateFilePerObject() OutputOption {
	return func(o *Flags) {
		o.noFilePerObjectArg = true
	}
}

// WithoutOutputDirectory disables the out-dir flag
func WithoutOutputDirectory() OutputOption {
	return func(o *Flags) {
		o.noOutputDirArg = true
	}
}

// NewFlags returns a new Flags object with the given options applied.
func NewFlags(opts ...OutputOption) *Flags {
	o := &Flags{
		Format: YAMLFormat,
		Writers: map[string]ObjectWriter{
			JSONFormat:  NewJSONObjectWriter(true /* useArrayJSON */),
			JSONLFormat: NewJSONObjectWriter(false /* useArrayJSON */),
			YAMLFormat:  NewYAMLObjectWriter(),
		},
	}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

// WithOutputOptions allows late application of options.
func (f *Flags) WithOutputOptions(opts ...OutputOption) {
	for _, opt := range opts {
		opt(f)
	}
}

// AddFlags adds the output flags to the given command.fr
func (f *Flags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Format, "output", "o", f.Format, "Output format. Must be one of ["+f.validOutputFormatsStr()+"].")
	if !f.noOutputDirArg {
		cmd.Flags().StringVar(&f.OutputDirectory, "output-dir", "", "Path of the directory where output file is stored.")
	}
	if !f.noFilePerObjectArg {
		cmd.Flags().BoolVar(&f.CreateFilePerObject, "create-file-per-object", false, "Split the output into multiple files. Must be used with output-dir flag. File name is in the format {object_type}_{slug_name}.yaml/json")
	}
}

func (f *Flags) validateFilePerObjectOption() error {
	if f.CreateFilePerObject && f.OutputDirectory == "" {
		return errors.New("output directory must be provided when using --create-file-per-object")
	}
	return nil
}

func (f *Flags) validOutputFormats() []string {
	formats := []string{}
	for format := range f.Writers {
		formats = append(formats, format)
	}

	// sort order for deterministic test assertion
	sort.Strings(formats)
	return formats
}

func (f *Flags) validOutputFormatsStr() string {
	return strings.Join(f.validOutputFormats(), "|")
}

// Validate validates the output flags.
func (f *Flags) Validate() error {
	var validFormat bool
	for _, ft := range f.validOutputFormats() {
		if ft == f.Format {
			validFormat = true
			break
		}
	}
	if !validFormat {
		return fmt.Errorf("invalid output format '%s', must be one of %s", f.Format, f.validOutputFormatsStr())
	}

	return f.validateFilePerObjectOption()
}

// NewWriterManager returns a writer manager
func (f Flags) NewWriterManager(w io.Writer) (*WriterManager, error) {
	return NewWriterManager(&f, w, &FileOps{})
}

// WriteObject marshals and writes an object.
func (f Flags) WriteObject(object interface{}, w io.Writer) error {
	writer := f.Writers[f.Format]
	opts := ObjectWriterOpts{
		multiObjects: true,
	}
	return writer.WriteObject(w, object, opts)
}

// WriteAfterClose prints out final messages after closing our object writer stream.
// Use for this can be printing an additional command to be run.
func (f Flags) WriteAfterClose(s string) {
	writer := f.Writers[f.Format]
	writer.WriteAfterClose(s)
}

// Close closes the object writer stream.
func (f Flags) Close(w io.Writer) error {
	writer := f.Writers[f.Format]
	return writer.Close(w, ObjectWriterOpts{multiObjects: true})
}

// ObjectWriterOpts are options for writing objects.
type ObjectWriterOpts struct {
	multiObjects bool
}

// ObjectWriter marshals an object and streams it to output.
type ObjectWriter interface {
	WriteObject(io.Writer, interface{}, ObjectWriterOpts) error
	WriteAfterClose(string)
	Close(io.Writer, ObjectWriterOpts) error
}

// YAMLObjectWriter is an ObjectWriter that marshals and streams yaml.
type YAMLObjectWriter struct {
	after       []string
	writtenOnce bool
}

// NewYAMLObjectWriter returns a new YAMLObjectWriter, which
// marshals and streams yaml.
func NewYAMLObjectWriter() *YAMLObjectWriter {
	return &YAMLObjectWriter{}
}

// WriteObject marshals and writes a YAML object to the output stream.
func (y *YAMLObjectWriter) WriteObject(w io.Writer, obj interface{}, opts ObjectWriterOpts) error {
	if opts.multiObjects && y.writtenOnce {
		if _, err := w.Write([]byte("---\n")); err != nil {
			return fmt.Errorf("unable to output converted file: %v", err)
		}
	}
	b, err := types.EncodeYAML(obj)
	if err != nil {
		return fmt.Errorf("unable to marshal object to YAML: %v", err)
	}
	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("unable to write object: %v", err)
	}
	y.writtenOnce = true
	return nil
}

// WriteAfterClose sets a final message to print after closing our object writer stream.
func (y *YAMLObjectWriter) WriteAfterClose(s string) {
	y.after = append(y.after, s)
}

// Close closes the object writer stream and prints any final messages
func (y *YAMLObjectWriter) Close(w io.Writer, _ ObjectWriterOpts) error {
	for _, a := range y.after {
		if _, err := w.Write([]byte(a)); err != nil {
			return err
		}
	}
	return nil
}

// JSONObjectWriter marshalls and writes JSON to an output stream. Not safe for concurrent use.
type JSONObjectWriter struct {
	useArray    bool
	after       []string
	writtenOnce bool
}

// NewJSONObjectWriter returns a new writer for JSON output
func NewJSONObjectWriter(useArray bool) *JSONObjectWriter {
	return &JSONObjectWriter{
		useArray: useArray,
	}
}

// WriteObject marshals and writes a JSON object to the output stream.
func (j *JSONObjectWriter) WriteObject(w io.Writer, obj interface{}, opts ObjectWriterOpts) error {
	var (
		delim        = "\n"
		indentPrefix = ""
	)
	if j.useArray {
		delim = ",\n  "
		indentPrefix = "  "

	}
	// For valid JSON, we need to start this with an open bracket to begin an array.
	if opts.multiObjects && j.useArray && !j.writtenOnce {
		if _, err := w.Write([]byte("[\n  ")); err != nil {
			return err
		}
	}
	// Once we've written a json object, we append `\n` to continue streaming
	// valid JSON. If using an array, we print `,\n`
	if opts.multiObjects && j.writtenOnce {
		if _, err := w.Write([]byte(delim)); err != nil {
			return fmt.Errorf("unable to output converted file: %v", err)
		}
	}

	b, err := json.MarshalIndent(obj, indentPrefix, "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal object to JSON: %v", err)
	}
	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("unable to write object: %v", err)
	}
	j.writtenOnce = true
	return nil
}

// WriteAfterClose sets a final message to print after closing our object writer stream.
func (j *JSONObjectWriter) WriteAfterClose(s string) {
	j.after = append(j.after, s)
}

// Close closes the object writer stream and prints any final messages.
func (j *JSONObjectWriter) Close(w io.Writer, opts ObjectWriterOpts) error {
	if opts.multiObjects && j.writtenOnce {
		closing := "\n"
		if j.useArray {
			closing = "\n]\n"
		}
		if _, err := w.Write([]byte(closing)); err != nil {
			return err
		}
	}
	for _, a := range j.after {
		if _, err := w.Write([]byte(a)); err != nil {
			return err
		}
	}

	return nil
}

// RowExtractor takes in an object and extracts a map of header name -> value.
type RowExtractor func(interface{}) (map[string]string, error)

// TableWriter marshalls and writes objects in a table format to an output stream.
type TableWriter struct {
	once         sync.Once
	table        *tablewriter.Table
	headers      []string
	rowFn        RowExtractor
	after        []string
	autoWrapText bool
}

type options struct {
	autoWrapText bool
}

// TableOptions is a function that configures a TableWriter.
type TableOptions func(*options)

// SetAutoWrapText sets whether or not the table writer should wrap text.
func SetAutoWrapText(autoWrap bool) TableOptions {
	return func(o *options) { o.autoWrapText = autoWrap }
}

// NewTableWriter returns an object marshaller that renders a table. Headers are essentially
// the column keys for the table, and extractor takes in a single object and produces a map
// of header key -> header value. Not safe for concurrent use.
// The table writer can be customized by passing in options.
func NewTableWriter(headers []string, extractor RowExtractor, os ...TableOptions) *TableWriter {
	opts := options{autoWrapText: true}
	for _, optFn := range os {
		optFn(&opts)
	}
	return &TableWriter{
		headers:      headers,
		rowFn:        extractor,
		autoWrapText: opts.autoWrapText,
	}
}

// maybeCreateTableWriter initializes the actual underlying tablewriter object. The
// tablewriter library doesn't allow changing where it gets written after creation,
// so we punt until we have an io.Writer.
func (t *TableWriter) maybeCreateTableWriter(w io.Writer) {
	t.once.Do(func() {
		table := tablewriter.NewWriter(w)
		table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
		table.SetCenterSeparator("|")
		table.SetHeader(t.headers)
		table.SetAutoWrapText(t.autoWrapText)
		t.table = table
	})
}

// WriteObject writes an object in table format to the output stream.
func (t *TableWriter) WriteObject(w io.Writer, obj interface{}, _ ObjectWriterOpts) error {
	t.maybeCreateTableWriter(w)

	values, err := t.rowFn(obj)
	if err != nil {
		return err
	}
	row := make([]string, 0, len(t.headers))
	for _, key := range t.headers {
		// it's okay to not check ok, we want an empty string in the correct position.
		value := values[key]
		row = append(row, value)
	}
	t.table.Append(row)
	return nil
}

// WriteAfterClose sets a final message to print after closing our object writer stream.
func (t *TableWriter) WriteAfterClose(s string) {
	t.after = append(t.after, s)
}

// Close renders the table and prints any final messages.
func (t *TableWriter) Close(w io.Writer, _ ObjectWriterOpts) error {
	t.maybeCreateTableWriter(w)
	t.table.Render()

	for _, a := range t.after {
		if _, err := w.Write([]byte(a)); err != nil {
			return err
		}
	}
	return nil
}
