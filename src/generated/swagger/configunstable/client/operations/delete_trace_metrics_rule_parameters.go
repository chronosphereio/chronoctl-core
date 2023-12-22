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

// NewDeleteTraceMetricsRuleParams creates a new DeleteTraceMetricsRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTraceMetricsRuleParams() *DeleteTraceMetricsRuleParams {
	return &DeleteTraceMetricsRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTraceMetricsRuleParamsWithTimeout creates a new DeleteTraceMetricsRuleParams object
// with the ability to set a timeout on a request.
func NewDeleteTraceMetricsRuleParamsWithTimeout(timeout time.Duration) *DeleteTraceMetricsRuleParams {
	return &DeleteTraceMetricsRuleParams{
		timeout: timeout,
	}
}

// NewDeleteTraceMetricsRuleParamsWithContext creates a new DeleteTraceMetricsRuleParams object
// with the ability to set a context for a request.
func NewDeleteTraceMetricsRuleParamsWithContext(ctx context.Context) *DeleteTraceMetricsRuleParams {
	return &DeleteTraceMetricsRuleParams{
		Context: ctx,
	}
}

// NewDeleteTraceMetricsRuleParamsWithHTTPClient creates a new DeleteTraceMetricsRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTraceMetricsRuleParamsWithHTTPClient(client *http.Client) *DeleteTraceMetricsRuleParams {
	return &DeleteTraceMetricsRuleParams{
		HTTPClient: client,
	}
}

/*
DeleteTraceMetricsRuleParams contains all the parameters to send to the API endpoint

	for the delete trace metrics rule operation.

	Typically these are written to a http.Request.
*/
type DeleteTraceMetricsRuleParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete trace metrics rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTraceMetricsRuleParams) WithDefaults() *DeleteTraceMetricsRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete trace metrics rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTraceMetricsRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete trace metrics rule params
func (o *DeleteTraceMetricsRuleParams) WithTimeout(timeout time.Duration) *DeleteTraceMetricsRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete trace metrics rule params
func (o *DeleteTraceMetricsRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete trace metrics rule params
func (o *DeleteTraceMetricsRuleParams) WithContext(ctx context.Context) *DeleteTraceMetricsRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete trace metrics rule params
func (o *DeleteTraceMetricsRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete trace metrics rule params
func (o *DeleteTraceMetricsRuleParams) WithHTTPClient(client *http.Client) *DeleteTraceMetricsRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete trace metrics rule params
func (o *DeleteTraceMetricsRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the delete trace metrics rule params
func (o *DeleteTraceMetricsRuleParams) WithSlug(slug string) *DeleteTraceMetricsRuleParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete trace metrics rule params
func (o *DeleteTraceMetricsRuleParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTraceMetricsRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
