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

// NewUpdateDropRuleParams creates a new UpdateDropRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDropRuleParams() *UpdateDropRuleParams {
	return &UpdateDropRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDropRuleParamsWithTimeout creates a new UpdateDropRuleParams object
// with the ability to set a timeout on a request.
func NewUpdateDropRuleParamsWithTimeout(timeout time.Duration) *UpdateDropRuleParams {
	return &UpdateDropRuleParams{
		timeout: timeout,
	}
}

// NewUpdateDropRuleParamsWithContext creates a new UpdateDropRuleParams object
// with the ability to set a context for a request.
func NewUpdateDropRuleParamsWithContext(ctx context.Context) *UpdateDropRuleParams {
	return &UpdateDropRuleParams{
		Context: ctx,
	}
}

// NewUpdateDropRuleParamsWithHTTPClient creates a new UpdateDropRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDropRuleParamsWithHTTPClient(client *http.Client) *UpdateDropRuleParams {
	return &UpdateDropRuleParams{
		HTTPClient: client,
	}
}

/*
UpdateDropRuleParams contains all the parameters to send to the API endpoint

	for the update drop rule operation.

	Typically these are written to a http.Request.
*/
type UpdateDropRuleParams struct {

	// Body.
	Body UpdateDropRuleBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update drop rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDropRuleParams) WithDefaults() *UpdateDropRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update drop rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDropRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update drop rule params
func (o *UpdateDropRuleParams) WithTimeout(timeout time.Duration) *UpdateDropRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update drop rule params
func (o *UpdateDropRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update drop rule params
func (o *UpdateDropRuleParams) WithContext(ctx context.Context) *UpdateDropRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update drop rule params
func (o *UpdateDropRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update drop rule params
func (o *UpdateDropRuleParams) WithHTTPClient(client *http.Client) *UpdateDropRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update drop rule params
func (o *UpdateDropRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update drop rule params
func (o *UpdateDropRuleParams) WithBody(body UpdateDropRuleBody) *UpdateDropRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update drop rule params
func (o *UpdateDropRuleParams) SetBody(body UpdateDropRuleBody) {
	o.Body = body
}

// WithSlug adds the slug to the update drop rule params
func (o *UpdateDropRuleParams) WithSlug(slug string) *UpdateDropRuleParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update drop rule params
func (o *UpdateDropRuleParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDropRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
