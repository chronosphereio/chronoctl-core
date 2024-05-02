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

// NewCreateLogAllocationConfigParams creates a new CreateLogAllocationConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateLogAllocationConfigParams() *CreateLogAllocationConfigParams {
	return &CreateLogAllocationConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLogAllocationConfigParamsWithTimeout creates a new CreateLogAllocationConfigParams object
// with the ability to set a timeout on a request.
func NewCreateLogAllocationConfigParamsWithTimeout(timeout time.Duration) *CreateLogAllocationConfigParams {
	return &CreateLogAllocationConfigParams{
		timeout: timeout,
	}
}

// NewCreateLogAllocationConfigParamsWithContext creates a new CreateLogAllocationConfigParams object
// with the ability to set a context for a request.
func NewCreateLogAllocationConfigParamsWithContext(ctx context.Context) *CreateLogAllocationConfigParams {
	return &CreateLogAllocationConfigParams{
		Context: ctx,
	}
}

// NewCreateLogAllocationConfigParamsWithHTTPClient creates a new CreateLogAllocationConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateLogAllocationConfigParamsWithHTTPClient(client *http.Client) *CreateLogAllocationConfigParams {
	return &CreateLogAllocationConfigParams{
		HTTPClient: client,
	}
}

/*
CreateLogAllocationConfigParams contains all the parameters to send to the API endpoint

	for the create log allocation config operation.

	Typically these are written to a http.Request.
*/
type CreateLogAllocationConfigParams struct {

	// Body.
	Body *models.ConfigunstableCreateLogAllocationConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create log allocation config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLogAllocationConfigParams) WithDefaults() *CreateLogAllocationConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create log allocation config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLogAllocationConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create log allocation config params
func (o *CreateLogAllocationConfigParams) WithTimeout(timeout time.Duration) *CreateLogAllocationConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create log allocation config params
func (o *CreateLogAllocationConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create log allocation config params
func (o *CreateLogAllocationConfigParams) WithContext(ctx context.Context) *CreateLogAllocationConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create log allocation config params
func (o *CreateLogAllocationConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create log allocation config params
func (o *CreateLogAllocationConfigParams) WithHTTPClient(client *http.Client) *CreateLogAllocationConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create log allocation config params
func (o *CreateLogAllocationConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create log allocation config params
func (o *CreateLogAllocationConfigParams) WithBody(body *models.ConfigunstableCreateLogAllocationConfigRequest) *CreateLogAllocationConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create log allocation config params
func (o *CreateLogAllocationConfigParams) SetBody(body *models.ConfigunstableCreateLogAllocationConfigRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLogAllocationConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
