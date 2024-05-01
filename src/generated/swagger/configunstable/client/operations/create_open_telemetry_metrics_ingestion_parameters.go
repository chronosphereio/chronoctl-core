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

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/models"
)

// NewCreateOpenTelemetryMetricsIngestionParams creates a new CreateOpenTelemetryMetricsIngestionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOpenTelemetryMetricsIngestionParams() *CreateOpenTelemetryMetricsIngestionParams {
	return &CreateOpenTelemetryMetricsIngestionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOpenTelemetryMetricsIngestionParamsWithTimeout creates a new CreateOpenTelemetryMetricsIngestionParams object
// with the ability to set a timeout on a request.
func NewCreateOpenTelemetryMetricsIngestionParamsWithTimeout(timeout time.Duration) *CreateOpenTelemetryMetricsIngestionParams {
	return &CreateOpenTelemetryMetricsIngestionParams{
		timeout: timeout,
	}
}

// NewCreateOpenTelemetryMetricsIngestionParamsWithContext creates a new CreateOpenTelemetryMetricsIngestionParams object
// with the ability to set a context for a request.
func NewCreateOpenTelemetryMetricsIngestionParamsWithContext(ctx context.Context) *CreateOpenTelemetryMetricsIngestionParams {
	return &CreateOpenTelemetryMetricsIngestionParams{
		Context: ctx,
	}
}

// NewCreateOpenTelemetryMetricsIngestionParamsWithHTTPClient creates a new CreateOpenTelemetryMetricsIngestionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOpenTelemetryMetricsIngestionParamsWithHTTPClient(client *http.Client) *CreateOpenTelemetryMetricsIngestionParams {
	return &CreateOpenTelemetryMetricsIngestionParams{
		HTTPClient: client,
	}
}

/*
CreateOpenTelemetryMetricsIngestionParams contains all the parameters to send to the API endpoint

	for the create open telemetry metrics ingestion operation.

	Typically these are written to a http.Request.
*/
type CreateOpenTelemetryMetricsIngestionParams struct {

	// Body.
	Body *models.ConfigunstableCreateOpenTelemetryMetricsIngestionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create open telemetry metrics ingestion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOpenTelemetryMetricsIngestionParams) WithDefaults() *CreateOpenTelemetryMetricsIngestionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create open telemetry metrics ingestion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOpenTelemetryMetricsIngestionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create open telemetry metrics ingestion params
func (o *CreateOpenTelemetryMetricsIngestionParams) WithTimeout(timeout time.Duration) *CreateOpenTelemetryMetricsIngestionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create open telemetry metrics ingestion params
func (o *CreateOpenTelemetryMetricsIngestionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create open telemetry metrics ingestion params
func (o *CreateOpenTelemetryMetricsIngestionParams) WithContext(ctx context.Context) *CreateOpenTelemetryMetricsIngestionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create open telemetry metrics ingestion params
func (o *CreateOpenTelemetryMetricsIngestionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create open telemetry metrics ingestion params
func (o *CreateOpenTelemetryMetricsIngestionParams) WithHTTPClient(client *http.Client) *CreateOpenTelemetryMetricsIngestionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create open telemetry metrics ingestion params
func (o *CreateOpenTelemetryMetricsIngestionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create open telemetry metrics ingestion params
func (o *CreateOpenTelemetryMetricsIngestionParams) WithBody(body *models.ConfigunstableCreateOpenTelemetryMetricsIngestionRequest) *CreateOpenTelemetryMetricsIngestionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create open telemetry metrics ingestion params
func (o *CreateOpenTelemetryMetricsIngestionParams) SetBody(body *models.ConfigunstableCreateOpenTelemetryMetricsIngestionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOpenTelemetryMetricsIngestionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
