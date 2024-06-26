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

// NewDeleteOtelMetricsIngestionParams creates a new DeleteOtelMetricsIngestionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOtelMetricsIngestionParams() *DeleteOtelMetricsIngestionParams {
	return &DeleteOtelMetricsIngestionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOtelMetricsIngestionParamsWithTimeout creates a new DeleteOtelMetricsIngestionParams object
// with the ability to set a timeout on a request.
func NewDeleteOtelMetricsIngestionParamsWithTimeout(timeout time.Duration) *DeleteOtelMetricsIngestionParams {
	return &DeleteOtelMetricsIngestionParams{
		timeout: timeout,
	}
}

// NewDeleteOtelMetricsIngestionParamsWithContext creates a new DeleteOtelMetricsIngestionParams object
// with the ability to set a context for a request.
func NewDeleteOtelMetricsIngestionParamsWithContext(ctx context.Context) *DeleteOtelMetricsIngestionParams {
	return &DeleteOtelMetricsIngestionParams{
		Context: ctx,
	}
}

// NewDeleteOtelMetricsIngestionParamsWithHTTPClient creates a new DeleteOtelMetricsIngestionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOtelMetricsIngestionParamsWithHTTPClient(client *http.Client) *DeleteOtelMetricsIngestionParams {
	return &DeleteOtelMetricsIngestionParams{
		HTTPClient: client,
	}
}

/*
DeleteOtelMetricsIngestionParams contains all the parameters to send to the API endpoint

	for the delete otel metrics ingestion operation.

	Typically these are written to a http.Request.
*/
type DeleteOtelMetricsIngestionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete otel metrics ingestion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOtelMetricsIngestionParams) WithDefaults() *DeleteOtelMetricsIngestionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete otel metrics ingestion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOtelMetricsIngestionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete otel metrics ingestion params
func (o *DeleteOtelMetricsIngestionParams) WithTimeout(timeout time.Duration) *DeleteOtelMetricsIngestionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete otel metrics ingestion params
func (o *DeleteOtelMetricsIngestionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete otel metrics ingestion params
func (o *DeleteOtelMetricsIngestionParams) WithContext(ctx context.Context) *DeleteOtelMetricsIngestionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete otel metrics ingestion params
func (o *DeleteOtelMetricsIngestionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete otel metrics ingestion params
func (o *DeleteOtelMetricsIngestionParams) WithHTTPClient(client *http.Client) *DeleteOtelMetricsIngestionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete otel metrics ingestion params
func (o *DeleteOtelMetricsIngestionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOtelMetricsIngestionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
