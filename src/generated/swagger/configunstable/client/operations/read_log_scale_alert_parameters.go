// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewReadLogScaleAlertParams creates a new ReadLogScaleAlertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadLogScaleAlertParams() *ReadLogScaleAlertParams {
	return &ReadLogScaleAlertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadLogScaleAlertParamsWithTimeout creates a new ReadLogScaleAlertParams object
// with the ability to set a timeout on a request.
func NewReadLogScaleAlertParamsWithTimeout(timeout time.Duration) *ReadLogScaleAlertParams {
	return &ReadLogScaleAlertParams{
		timeout: timeout,
	}
}

// NewReadLogScaleAlertParamsWithContext creates a new ReadLogScaleAlertParams object
// with the ability to set a context for a request.
func NewReadLogScaleAlertParamsWithContext(ctx context.Context) *ReadLogScaleAlertParams {
	return &ReadLogScaleAlertParams{
		Context: ctx,
	}
}

// NewReadLogScaleAlertParamsWithHTTPClient creates a new ReadLogScaleAlertParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadLogScaleAlertParamsWithHTTPClient(client *http.Client) *ReadLogScaleAlertParams {
	return &ReadLogScaleAlertParams{
		HTTPClient: client,
	}
}

/*
ReadLogScaleAlertParams contains all the parameters to send to the API endpoint

	for the read log scale alert operation.

	Typically these are written to a http.Request.
*/
type ReadLogScaleAlertParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read log scale alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadLogScaleAlertParams) WithDefaults() *ReadLogScaleAlertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read log scale alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadLogScaleAlertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read log scale alert params
func (o *ReadLogScaleAlertParams) WithTimeout(timeout time.Duration) *ReadLogScaleAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read log scale alert params
func (o *ReadLogScaleAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read log scale alert params
func (o *ReadLogScaleAlertParams) WithContext(ctx context.Context) *ReadLogScaleAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read log scale alert params
func (o *ReadLogScaleAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read log scale alert params
func (o *ReadLogScaleAlertParams) WithHTTPClient(client *http.Client) *ReadLogScaleAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read log scale alert params
func (o *ReadLogScaleAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read log scale alert params
func (o *ReadLogScaleAlertParams) WithSlug(slug string) *ReadLogScaleAlertParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read log scale alert params
func (o *ReadLogScaleAlertParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadLogScaleAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}