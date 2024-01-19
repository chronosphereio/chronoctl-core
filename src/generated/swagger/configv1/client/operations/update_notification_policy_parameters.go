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

// NewUpdateNotificationPolicyParams creates a new UpdateNotificationPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNotificationPolicyParams() *UpdateNotificationPolicyParams {
	return &UpdateNotificationPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNotificationPolicyParamsWithTimeout creates a new UpdateNotificationPolicyParams object
// with the ability to set a timeout on a request.
func NewUpdateNotificationPolicyParamsWithTimeout(timeout time.Duration) *UpdateNotificationPolicyParams {
	return &UpdateNotificationPolicyParams{
		timeout: timeout,
	}
}

// NewUpdateNotificationPolicyParamsWithContext creates a new UpdateNotificationPolicyParams object
// with the ability to set a context for a request.
func NewUpdateNotificationPolicyParamsWithContext(ctx context.Context) *UpdateNotificationPolicyParams {
	return &UpdateNotificationPolicyParams{
		Context: ctx,
	}
}

// NewUpdateNotificationPolicyParamsWithHTTPClient creates a new UpdateNotificationPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNotificationPolicyParamsWithHTTPClient(client *http.Client) *UpdateNotificationPolicyParams {
	return &UpdateNotificationPolicyParams{
		HTTPClient: client,
	}
}

/*
UpdateNotificationPolicyParams contains all the parameters to send to the API endpoint

	for the update notification policy operation.

	Typically these are written to a http.Request.
*/
type UpdateNotificationPolicyParams struct {

	// Body.
	Body UpdateNotificationPolicyBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update notification policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNotificationPolicyParams) WithDefaults() *UpdateNotificationPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update notification policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNotificationPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update notification policy params
func (o *UpdateNotificationPolicyParams) WithTimeout(timeout time.Duration) *UpdateNotificationPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update notification policy params
func (o *UpdateNotificationPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update notification policy params
func (o *UpdateNotificationPolicyParams) WithContext(ctx context.Context) *UpdateNotificationPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update notification policy params
func (o *UpdateNotificationPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update notification policy params
func (o *UpdateNotificationPolicyParams) WithHTTPClient(client *http.Client) *UpdateNotificationPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update notification policy params
func (o *UpdateNotificationPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update notification policy params
func (o *UpdateNotificationPolicyParams) WithBody(body UpdateNotificationPolicyBody) *UpdateNotificationPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update notification policy params
func (o *UpdateNotificationPolicyParams) SetBody(body UpdateNotificationPolicyBody) {
	o.Body = body
}

// WithSlug adds the slug to the update notification policy params
func (o *UpdateNotificationPolicyParams) WithSlug(slug string) *UpdateNotificationPolicyParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update notification policy params
func (o *UpdateNotificationPolicyParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNotificationPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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