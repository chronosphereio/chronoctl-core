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

// NewReadOpenTelemetryMetricsIngestionParams creates a new ReadOpenTelemetryMetricsIngestionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadOpenTelemetryMetricsIngestionParams() *ReadOpenTelemetryMetricsIngestionParams {
	return &ReadOpenTelemetryMetricsIngestionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadOpenTelemetryMetricsIngestionParamsWithTimeout creates a new ReadOpenTelemetryMetricsIngestionParams object
// with the ability to set a timeout on a request.
func NewReadOpenTelemetryMetricsIngestionParamsWithTimeout(timeout time.Duration) *ReadOpenTelemetryMetricsIngestionParams {
	return &ReadOpenTelemetryMetricsIngestionParams{
		timeout: timeout,
	}
}

// NewReadOpenTelemetryMetricsIngestionParamsWithContext creates a new ReadOpenTelemetryMetricsIngestionParams object
// with the ability to set a context for a request.
func NewReadOpenTelemetryMetricsIngestionParamsWithContext(ctx context.Context) *ReadOpenTelemetryMetricsIngestionParams {
	return &ReadOpenTelemetryMetricsIngestionParams{
		Context: ctx,
	}
}

// NewReadOpenTelemetryMetricsIngestionParamsWithHTTPClient creates a new ReadOpenTelemetryMetricsIngestionParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadOpenTelemetryMetricsIngestionParamsWithHTTPClient(client *http.Client) *ReadOpenTelemetryMetricsIngestionParams {
	return &ReadOpenTelemetryMetricsIngestionParams{
		HTTPClient: client,
	}
}

/*
ReadOpenTelemetryMetricsIngestionParams contains all the parameters to send to the API endpoint

	for the read open telemetry metrics ingestion operation.

	Typically these are written to a http.Request.
*/
type ReadOpenTelemetryMetricsIngestionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read open telemetry metrics ingestion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadOpenTelemetryMetricsIngestionParams) WithDefaults() *ReadOpenTelemetryMetricsIngestionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read open telemetry metrics ingestion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadOpenTelemetryMetricsIngestionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read open telemetry metrics ingestion params
func (o *ReadOpenTelemetryMetricsIngestionParams) WithTimeout(timeout time.Duration) *ReadOpenTelemetryMetricsIngestionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read open telemetry metrics ingestion params
func (o *ReadOpenTelemetryMetricsIngestionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read open telemetry metrics ingestion params
func (o *ReadOpenTelemetryMetricsIngestionParams) WithContext(ctx context.Context) *ReadOpenTelemetryMetricsIngestionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read open telemetry metrics ingestion params
func (o *ReadOpenTelemetryMetricsIngestionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read open telemetry metrics ingestion params
func (o *ReadOpenTelemetryMetricsIngestionParams) WithHTTPClient(client *http.Client) *ReadOpenTelemetryMetricsIngestionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read open telemetry metrics ingestion params
func (o *ReadOpenTelemetryMetricsIngestionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReadOpenTelemetryMetricsIngestionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
