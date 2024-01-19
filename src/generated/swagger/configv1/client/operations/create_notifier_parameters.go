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

// NewCreateNotifierParams creates a new CreateNotifierParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNotifierParams() *CreateNotifierParams {
	return &CreateNotifierParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNotifierParamsWithTimeout creates a new CreateNotifierParams object
// with the ability to set a timeout on a request.
func NewCreateNotifierParamsWithTimeout(timeout time.Duration) *CreateNotifierParams {
	return &CreateNotifierParams{
		timeout: timeout,
	}
}

// NewCreateNotifierParamsWithContext creates a new CreateNotifierParams object
// with the ability to set a context for a request.
func NewCreateNotifierParamsWithContext(ctx context.Context) *CreateNotifierParams {
	return &CreateNotifierParams{
		Context: ctx,
	}
}

// NewCreateNotifierParamsWithHTTPClient creates a new CreateNotifierParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNotifierParamsWithHTTPClient(client *http.Client) *CreateNotifierParams {
	return &CreateNotifierParams{
		HTTPClient: client,
	}
}

/*
CreateNotifierParams contains all the parameters to send to the API endpoint

	for the create notifier operation.

	Typically these are written to a http.Request.
*/
type CreateNotifierParams struct {

	// Body.
	Body *models.Configv1CreateNotifierRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create notifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNotifierParams) WithDefaults() *CreateNotifierParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create notifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNotifierParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create notifier params
func (o *CreateNotifierParams) WithTimeout(timeout time.Duration) *CreateNotifierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create notifier params
func (o *CreateNotifierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create notifier params
func (o *CreateNotifierParams) WithContext(ctx context.Context) *CreateNotifierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create notifier params
func (o *CreateNotifierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create notifier params
func (o *CreateNotifierParams) WithHTTPClient(client *http.Client) *CreateNotifierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create notifier params
func (o *CreateNotifierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create notifier params
func (o *CreateNotifierParams) WithBody(body *models.Configv1CreateNotifierRequest) *CreateNotifierParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create notifier params
func (o *CreateNotifierParams) SetBody(body *models.Configv1CreateNotifierRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNotifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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