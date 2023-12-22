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

// NewReadServiceAccountParams creates a new ReadServiceAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadServiceAccountParams() *ReadServiceAccountParams {
	return &ReadServiceAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadServiceAccountParamsWithTimeout creates a new ReadServiceAccountParams object
// with the ability to set a timeout on a request.
func NewReadServiceAccountParamsWithTimeout(timeout time.Duration) *ReadServiceAccountParams {
	return &ReadServiceAccountParams{
		timeout: timeout,
	}
}

// NewReadServiceAccountParamsWithContext creates a new ReadServiceAccountParams object
// with the ability to set a context for a request.
func NewReadServiceAccountParamsWithContext(ctx context.Context) *ReadServiceAccountParams {
	return &ReadServiceAccountParams{
		Context: ctx,
	}
}

// NewReadServiceAccountParamsWithHTTPClient creates a new ReadServiceAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadServiceAccountParamsWithHTTPClient(client *http.Client) *ReadServiceAccountParams {
	return &ReadServiceAccountParams{
		HTTPClient: client,
	}
}

/*
ReadServiceAccountParams contains all the parameters to send to the API endpoint

	for the read service account operation.

	Typically these are written to a http.Request.
*/
type ReadServiceAccountParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadServiceAccountParams) WithDefaults() *ReadServiceAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadServiceAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read service account params
func (o *ReadServiceAccountParams) WithTimeout(timeout time.Duration) *ReadServiceAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read service account params
func (o *ReadServiceAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read service account params
func (o *ReadServiceAccountParams) WithContext(ctx context.Context) *ReadServiceAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read service account params
func (o *ReadServiceAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read service account params
func (o *ReadServiceAccountParams) WithHTTPClient(client *http.Client) *ReadServiceAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read service account params
func (o *ReadServiceAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read service account params
func (o *ReadServiceAccountParams) WithSlug(slug string) *ReadServiceAccountParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read service account params
func (o *ReadServiceAccountParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadServiceAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
