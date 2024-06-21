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

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
)

// NewUpdateGcpMetricsIntegrationParams creates a new UpdateGcpMetricsIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGcpMetricsIntegrationParams() *UpdateGcpMetricsIntegrationParams {
	return &UpdateGcpMetricsIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGcpMetricsIntegrationParamsWithTimeout creates a new UpdateGcpMetricsIntegrationParams object
// with the ability to set a timeout on a request.
func NewUpdateGcpMetricsIntegrationParamsWithTimeout(timeout time.Duration) *UpdateGcpMetricsIntegrationParams {
	return &UpdateGcpMetricsIntegrationParams{
		timeout: timeout,
	}
}

// NewUpdateGcpMetricsIntegrationParamsWithContext creates a new UpdateGcpMetricsIntegrationParams object
// with the ability to set a context for a request.
func NewUpdateGcpMetricsIntegrationParamsWithContext(ctx context.Context) *UpdateGcpMetricsIntegrationParams {
	return &UpdateGcpMetricsIntegrationParams{
		Context: ctx,
	}
}

// NewUpdateGcpMetricsIntegrationParamsWithHTTPClient creates a new UpdateGcpMetricsIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGcpMetricsIntegrationParamsWithHTTPClient(client *http.Client) *UpdateGcpMetricsIntegrationParams {
	return &UpdateGcpMetricsIntegrationParams{
		HTTPClient: client,
	}
}

/*
UpdateGcpMetricsIntegrationParams contains all the parameters to send to the API endpoint

	for the update gcp metrics integration operation.

	Typically these are written to a http.Request.
*/
type UpdateGcpMetricsIntegrationParams struct {

	// Body.
	Body *models.ConfigV1UpdateGcpMetricsIntegrationBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update gcp metrics integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGcpMetricsIntegrationParams) WithDefaults() *UpdateGcpMetricsIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update gcp metrics integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGcpMetricsIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update gcp metrics integration params
func (o *UpdateGcpMetricsIntegrationParams) WithTimeout(timeout time.Duration) *UpdateGcpMetricsIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update gcp metrics integration params
func (o *UpdateGcpMetricsIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update gcp metrics integration params
func (o *UpdateGcpMetricsIntegrationParams) WithContext(ctx context.Context) *UpdateGcpMetricsIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update gcp metrics integration params
func (o *UpdateGcpMetricsIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update gcp metrics integration params
func (o *UpdateGcpMetricsIntegrationParams) WithHTTPClient(client *http.Client) *UpdateGcpMetricsIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update gcp metrics integration params
func (o *UpdateGcpMetricsIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update gcp metrics integration params
func (o *UpdateGcpMetricsIntegrationParams) WithBody(body *models.ConfigV1UpdateGcpMetricsIntegrationBody) *UpdateGcpMetricsIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update gcp metrics integration params
func (o *UpdateGcpMetricsIntegrationParams) SetBody(body *models.ConfigV1UpdateGcpMetricsIntegrationBody) {
	o.Body = body
}

// WithSlug adds the slug to the update gcp metrics integration params
func (o *UpdateGcpMetricsIntegrationParams) WithSlug(slug string) *UpdateGcpMetricsIntegrationParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update gcp metrics integration params
func (o *UpdateGcpMetricsIntegrationParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGcpMetricsIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
