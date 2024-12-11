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

// NewCreateSLOParams creates a new CreateSLOParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSLOParams() *CreateSLOParams {
	return &CreateSLOParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSLOParamsWithTimeout creates a new CreateSLOParams object
// with the ability to set a timeout on a request.
func NewCreateSLOParamsWithTimeout(timeout time.Duration) *CreateSLOParams {
	return &CreateSLOParams{
		timeout: timeout,
	}
}

// NewCreateSLOParamsWithContext creates a new CreateSLOParams object
// with the ability to set a context for a request.
func NewCreateSLOParamsWithContext(ctx context.Context) *CreateSLOParams {
	return &CreateSLOParams{
		Context: ctx,
	}
}

// NewCreateSLOParamsWithHTTPClient creates a new CreateSLOParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSLOParamsWithHTTPClient(client *http.Client) *CreateSLOParams {
	return &CreateSLOParams{
		HTTPClient: client,
	}
}

/*
CreateSLOParams contains all the parameters to send to the API endpoint

	for the create SLO operation.

	Typically these are written to a http.Request.
*/
type CreateSLOParams struct {

	// Body.
	Body *models.ConfigunstableCreateSLORequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create SLO params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSLOParams) WithDefaults() *CreateSLOParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create SLO params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSLOParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create SLO params
func (o *CreateSLOParams) WithTimeout(timeout time.Duration) *CreateSLOParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create SLO params
func (o *CreateSLOParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create SLO params
func (o *CreateSLOParams) WithContext(ctx context.Context) *CreateSLOParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create SLO params
func (o *CreateSLOParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create SLO params
func (o *CreateSLOParams) WithHTTPClient(client *http.Client) *CreateSLOParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create SLO params
func (o *CreateSLOParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create SLO params
func (o *CreateSLOParams) WithBody(body *models.ConfigunstableCreateSLORequest) *CreateSLOParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create SLO params
func (o *CreateSLOParams) SetBody(body *models.ConfigunstableCreateSLORequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSLOParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
