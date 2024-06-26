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

// NewUpdateServiceParams creates a new UpdateServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateServiceParams() *UpdateServiceParams {
	return &UpdateServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceParamsWithTimeout creates a new UpdateServiceParams object
// with the ability to set a timeout on a request.
func NewUpdateServiceParamsWithTimeout(timeout time.Duration) *UpdateServiceParams {
	return &UpdateServiceParams{
		timeout: timeout,
	}
}

// NewUpdateServiceParamsWithContext creates a new UpdateServiceParams object
// with the ability to set a context for a request.
func NewUpdateServiceParamsWithContext(ctx context.Context) *UpdateServiceParams {
	return &UpdateServiceParams{
		Context: ctx,
	}
}

// NewUpdateServiceParamsWithHTTPClient creates a new UpdateServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateServiceParamsWithHTTPClient(client *http.Client) *UpdateServiceParams {
	return &UpdateServiceParams{
		HTTPClient: client,
	}
}

/*
UpdateServiceParams contains all the parameters to send to the API endpoint

	for the update service operation.

	Typically these are written to a http.Request.
*/
type UpdateServiceParams struct {

	// Body.
	Body *models.ConfigUnstableUpdateServiceBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServiceParams) WithDefaults() *UpdateServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update service params
func (o *UpdateServiceParams) WithTimeout(timeout time.Duration) *UpdateServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service params
func (o *UpdateServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service params
func (o *UpdateServiceParams) WithContext(ctx context.Context) *UpdateServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service params
func (o *UpdateServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service params
func (o *UpdateServiceParams) WithHTTPClient(client *http.Client) *UpdateServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service params
func (o *UpdateServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update service params
func (o *UpdateServiceParams) WithBody(body *models.ConfigUnstableUpdateServiceBody) *UpdateServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update service params
func (o *UpdateServiceParams) SetBody(body *models.ConfigUnstableUpdateServiceBody) {
	o.Body = body
}

// WithSlug adds the slug to the update service params
func (o *UpdateServiceParams) WithSlug(slug string) *UpdateServiceParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update service params
func (o *UpdateServiceParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
