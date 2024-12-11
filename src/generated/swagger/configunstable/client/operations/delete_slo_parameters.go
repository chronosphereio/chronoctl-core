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

// NewDeleteSLOParams creates a new DeleteSLOParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSLOParams() *DeleteSLOParams {
	return &DeleteSLOParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSLOParamsWithTimeout creates a new DeleteSLOParams object
// with the ability to set a timeout on a request.
func NewDeleteSLOParamsWithTimeout(timeout time.Duration) *DeleteSLOParams {
	return &DeleteSLOParams{
		timeout: timeout,
	}
}

// NewDeleteSLOParamsWithContext creates a new DeleteSLOParams object
// with the ability to set a context for a request.
func NewDeleteSLOParamsWithContext(ctx context.Context) *DeleteSLOParams {
	return &DeleteSLOParams{
		Context: ctx,
	}
}

// NewDeleteSLOParamsWithHTTPClient creates a new DeleteSLOParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSLOParamsWithHTTPClient(client *http.Client) *DeleteSLOParams {
	return &DeleteSLOParams{
		HTTPClient: client,
	}
}

/*
DeleteSLOParams contains all the parameters to send to the API endpoint

	for the delete SLO operation.

	Typically these are written to a http.Request.
*/
type DeleteSLOParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete SLO params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSLOParams) WithDefaults() *DeleteSLOParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete SLO params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSLOParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete SLO params
func (o *DeleteSLOParams) WithTimeout(timeout time.Duration) *DeleteSLOParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete SLO params
func (o *DeleteSLOParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete SLO params
func (o *DeleteSLOParams) WithContext(ctx context.Context) *DeleteSLOParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete SLO params
func (o *DeleteSLOParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete SLO params
func (o *DeleteSLOParams) WithHTTPClient(client *http.Client) *DeleteSLOParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete SLO params
func (o *DeleteSLOParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the delete SLO params
func (o *DeleteSLOParams) WithSlug(slug string) *DeleteSLOParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete SLO params
func (o *DeleteSLOParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSLOParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
