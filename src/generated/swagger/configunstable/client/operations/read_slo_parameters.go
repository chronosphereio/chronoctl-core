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

// NewReadSLOParams creates a new ReadSLOParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadSLOParams() *ReadSLOParams {
	return &ReadSLOParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadSLOParamsWithTimeout creates a new ReadSLOParams object
// with the ability to set a timeout on a request.
func NewReadSLOParamsWithTimeout(timeout time.Duration) *ReadSLOParams {
	return &ReadSLOParams{
		timeout: timeout,
	}
}

// NewReadSLOParamsWithContext creates a new ReadSLOParams object
// with the ability to set a context for a request.
func NewReadSLOParamsWithContext(ctx context.Context) *ReadSLOParams {
	return &ReadSLOParams{
		Context: ctx,
	}
}

// NewReadSLOParamsWithHTTPClient creates a new ReadSLOParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadSLOParamsWithHTTPClient(client *http.Client) *ReadSLOParams {
	return &ReadSLOParams{
		HTTPClient: client,
	}
}

/*
ReadSLOParams contains all the parameters to send to the API endpoint

	for the read SLO operation.

	Typically these are written to a http.Request.
*/
type ReadSLOParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read SLO params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadSLOParams) WithDefaults() *ReadSLOParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read SLO params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadSLOParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read SLO params
func (o *ReadSLOParams) WithTimeout(timeout time.Duration) *ReadSLOParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read SLO params
func (o *ReadSLOParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read SLO params
func (o *ReadSLOParams) WithContext(ctx context.Context) *ReadSLOParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read SLO params
func (o *ReadSLOParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read SLO params
func (o *ReadSLOParams) WithHTTPClient(client *http.Client) *ReadSLOParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read SLO params
func (o *ReadSLOParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read SLO params
func (o *ReadSLOParams) WithSlug(slug string) *ReadSLOParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read SLO params
func (o *ReadSLOParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadSLOParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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