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

// NewDeleteNoopEntityParams creates a new DeleteNoopEntityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNoopEntityParams() *DeleteNoopEntityParams {
	return &DeleteNoopEntityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNoopEntityParamsWithTimeout creates a new DeleteNoopEntityParams object
// with the ability to set a timeout on a request.
func NewDeleteNoopEntityParamsWithTimeout(timeout time.Duration) *DeleteNoopEntityParams {
	return &DeleteNoopEntityParams{
		timeout: timeout,
	}
}

// NewDeleteNoopEntityParamsWithContext creates a new DeleteNoopEntityParams object
// with the ability to set a context for a request.
func NewDeleteNoopEntityParamsWithContext(ctx context.Context) *DeleteNoopEntityParams {
	return &DeleteNoopEntityParams{
		Context: ctx,
	}
}

// NewDeleteNoopEntityParamsWithHTTPClient creates a new DeleteNoopEntityParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNoopEntityParamsWithHTTPClient(client *http.Client) *DeleteNoopEntityParams {
	return &DeleteNoopEntityParams{
		HTTPClient: client,
	}
}

/*
DeleteNoopEntityParams contains all the parameters to send to the API endpoint

	for the delete noop entity operation.

	Typically these are written to a http.Request.
*/
type DeleteNoopEntityParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete noop entity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNoopEntityParams) WithDefaults() *DeleteNoopEntityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete noop entity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNoopEntityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete noop entity params
func (o *DeleteNoopEntityParams) WithTimeout(timeout time.Duration) *DeleteNoopEntityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete noop entity params
func (o *DeleteNoopEntityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete noop entity params
func (o *DeleteNoopEntityParams) WithContext(ctx context.Context) *DeleteNoopEntityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete noop entity params
func (o *DeleteNoopEntityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete noop entity params
func (o *DeleteNoopEntityParams) WithHTTPClient(client *http.Client) *DeleteNoopEntityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete noop entity params
func (o *DeleteNoopEntityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the delete noop entity params
func (o *DeleteNoopEntityParams) WithSlug(slug string) *DeleteNoopEntityParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete noop entity params
func (o *DeleteNoopEntityParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNoopEntityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
