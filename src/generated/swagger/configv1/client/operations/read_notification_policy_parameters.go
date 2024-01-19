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

// NewReadNotificationPolicyParams creates a new ReadNotificationPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadNotificationPolicyParams() *ReadNotificationPolicyParams {
	return &ReadNotificationPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadNotificationPolicyParamsWithTimeout creates a new ReadNotificationPolicyParams object
// with the ability to set a timeout on a request.
func NewReadNotificationPolicyParamsWithTimeout(timeout time.Duration) *ReadNotificationPolicyParams {
	return &ReadNotificationPolicyParams{
		timeout: timeout,
	}
}

// NewReadNotificationPolicyParamsWithContext creates a new ReadNotificationPolicyParams object
// with the ability to set a context for a request.
func NewReadNotificationPolicyParamsWithContext(ctx context.Context) *ReadNotificationPolicyParams {
	return &ReadNotificationPolicyParams{
		Context: ctx,
	}
}

// NewReadNotificationPolicyParamsWithHTTPClient creates a new ReadNotificationPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadNotificationPolicyParamsWithHTTPClient(client *http.Client) *ReadNotificationPolicyParams {
	return &ReadNotificationPolicyParams{
		HTTPClient: client,
	}
}

/*
ReadNotificationPolicyParams contains all the parameters to send to the API endpoint

	for the read notification policy operation.

	Typically these are written to a http.Request.
*/
type ReadNotificationPolicyParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read notification policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadNotificationPolicyParams) WithDefaults() *ReadNotificationPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read notification policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadNotificationPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read notification policy params
func (o *ReadNotificationPolicyParams) WithTimeout(timeout time.Duration) *ReadNotificationPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read notification policy params
func (o *ReadNotificationPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read notification policy params
func (o *ReadNotificationPolicyParams) WithContext(ctx context.Context) *ReadNotificationPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read notification policy params
func (o *ReadNotificationPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read notification policy params
func (o *ReadNotificationPolicyParams) WithHTTPClient(client *http.Client) *ReadNotificationPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read notification policy params
func (o *ReadNotificationPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read notification policy params
func (o *ReadNotificationPolicyParams) WithSlug(slug string) *ReadNotificationPolicyParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read notification policy params
func (o *ReadNotificationPolicyParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadNotificationPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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