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

// NewDeleteTraceBehaviorParams creates a new DeleteTraceBehaviorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTraceBehaviorParams() *DeleteTraceBehaviorParams {
	return &DeleteTraceBehaviorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTraceBehaviorParamsWithTimeout creates a new DeleteTraceBehaviorParams object
// with the ability to set a timeout on a request.
func NewDeleteTraceBehaviorParamsWithTimeout(timeout time.Duration) *DeleteTraceBehaviorParams {
	return &DeleteTraceBehaviorParams{
		timeout: timeout,
	}
}

// NewDeleteTraceBehaviorParamsWithContext creates a new DeleteTraceBehaviorParams object
// with the ability to set a context for a request.
func NewDeleteTraceBehaviorParamsWithContext(ctx context.Context) *DeleteTraceBehaviorParams {
	return &DeleteTraceBehaviorParams{
		Context: ctx,
	}
}

// NewDeleteTraceBehaviorParamsWithHTTPClient creates a new DeleteTraceBehaviorParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTraceBehaviorParamsWithHTTPClient(client *http.Client) *DeleteTraceBehaviorParams {
	return &DeleteTraceBehaviorParams{
		HTTPClient: client,
	}
}

/*
DeleteTraceBehaviorParams contains all the parameters to send to the API endpoint

	for the delete trace behavior operation.

	Typically these are written to a http.Request.
*/
type DeleteTraceBehaviorParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete trace behavior params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTraceBehaviorParams) WithDefaults() *DeleteTraceBehaviorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete trace behavior params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTraceBehaviorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete trace behavior params
func (o *DeleteTraceBehaviorParams) WithTimeout(timeout time.Duration) *DeleteTraceBehaviorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete trace behavior params
func (o *DeleteTraceBehaviorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete trace behavior params
func (o *DeleteTraceBehaviorParams) WithContext(ctx context.Context) *DeleteTraceBehaviorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete trace behavior params
func (o *DeleteTraceBehaviorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete trace behavior params
func (o *DeleteTraceBehaviorParams) WithHTTPClient(client *http.Client) *DeleteTraceBehaviorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete trace behavior params
func (o *DeleteTraceBehaviorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the delete trace behavior params
func (o *DeleteTraceBehaviorParams) WithSlug(slug string) *DeleteTraceBehaviorParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete trace behavior params
func (o *DeleteTraceBehaviorParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTraceBehaviorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
