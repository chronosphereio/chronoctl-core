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

// NewDeleteGCPMetricsIntegrationParams creates a new DeleteGCPMetricsIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGCPMetricsIntegrationParams() *DeleteGCPMetricsIntegrationParams {
	return &DeleteGCPMetricsIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGCPMetricsIntegrationParamsWithTimeout creates a new DeleteGCPMetricsIntegrationParams object
// with the ability to set a timeout on a request.
func NewDeleteGCPMetricsIntegrationParamsWithTimeout(timeout time.Duration) *DeleteGCPMetricsIntegrationParams {
	return &DeleteGCPMetricsIntegrationParams{
		timeout: timeout,
	}
}

// NewDeleteGCPMetricsIntegrationParamsWithContext creates a new DeleteGCPMetricsIntegrationParams object
// with the ability to set a context for a request.
func NewDeleteGCPMetricsIntegrationParamsWithContext(ctx context.Context) *DeleteGCPMetricsIntegrationParams {
	return &DeleteGCPMetricsIntegrationParams{
		Context: ctx,
	}
}

// NewDeleteGCPMetricsIntegrationParamsWithHTTPClient creates a new DeleteGCPMetricsIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGCPMetricsIntegrationParamsWithHTTPClient(client *http.Client) *DeleteGCPMetricsIntegrationParams {
	return &DeleteGCPMetricsIntegrationParams{
		HTTPClient: client,
	}
}

/*
DeleteGCPMetricsIntegrationParams contains all the parameters to send to the API endpoint

	for the delete g c p metrics integration operation.

	Typically these are written to a http.Request.
*/
type DeleteGCPMetricsIntegrationParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete g c p metrics integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGCPMetricsIntegrationParams) WithDefaults() *DeleteGCPMetricsIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete g c p metrics integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGCPMetricsIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete g c p metrics integration params
func (o *DeleteGCPMetricsIntegrationParams) WithTimeout(timeout time.Duration) *DeleteGCPMetricsIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete g c p metrics integration params
func (o *DeleteGCPMetricsIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete g c p metrics integration params
func (o *DeleteGCPMetricsIntegrationParams) WithContext(ctx context.Context) *DeleteGCPMetricsIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete g c p metrics integration params
func (o *DeleteGCPMetricsIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete g c p metrics integration params
func (o *DeleteGCPMetricsIntegrationParams) WithHTTPClient(client *http.Client) *DeleteGCPMetricsIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete g c p metrics integration params
func (o *DeleteGCPMetricsIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the delete g c p metrics integration params
func (o *DeleteGCPMetricsIntegrationParams) WithSlug(slug string) *DeleteGCPMetricsIntegrationParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete g c p metrics integration params
func (o *DeleteGCPMetricsIntegrationParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGCPMetricsIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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