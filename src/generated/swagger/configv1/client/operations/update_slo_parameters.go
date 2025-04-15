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

// NewUpdateSLOParams creates a new UpdateSLOParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSLOParams() *UpdateSLOParams {
	return &UpdateSLOParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSLOParamsWithTimeout creates a new UpdateSLOParams object
// with the ability to set a timeout on a request.
func NewUpdateSLOParamsWithTimeout(timeout time.Duration) *UpdateSLOParams {
	return &UpdateSLOParams{
		timeout: timeout,
	}
}

// NewUpdateSLOParamsWithContext creates a new UpdateSLOParams object
// with the ability to set a context for a request.
func NewUpdateSLOParamsWithContext(ctx context.Context) *UpdateSLOParams {
	return &UpdateSLOParams{
		Context: ctx,
	}
}

// NewUpdateSLOParamsWithHTTPClient creates a new UpdateSLOParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSLOParamsWithHTTPClient(client *http.Client) *UpdateSLOParams {
	return &UpdateSLOParams{
		HTTPClient: client,
	}
}

/*
UpdateSLOParams contains all the parameters to send to the API endpoint

	for the update SLO operation.

	Typically these are written to a http.Request.
*/
type UpdateSLOParams struct {

	// Body.
	Body *models.ConfigV1UpdateSLOBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update SLO params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSLOParams) WithDefaults() *UpdateSLOParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update SLO params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSLOParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update SLO params
func (o *UpdateSLOParams) WithTimeout(timeout time.Duration) *UpdateSLOParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update SLO params
func (o *UpdateSLOParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update SLO params
func (o *UpdateSLOParams) WithContext(ctx context.Context) *UpdateSLOParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update SLO params
func (o *UpdateSLOParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update SLO params
func (o *UpdateSLOParams) WithHTTPClient(client *http.Client) *UpdateSLOParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update SLO params
func (o *UpdateSLOParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update SLO params
func (o *UpdateSLOParams) WithBody(body *models.ConfigV1UpdateSLOBody) *UpdateSLOParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update SLO params
func (o *UpdateSLOParams) SetBody(body *models.ConfigV1UpdateSLOBody) {
	o.Body = body
}

// WithSlug adds the slug to the update SLO params
func (o *UpdateSLOParams) WithSlug(slug string) *UpdateSLOParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update SLO params
func (o *UpdateSLOParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSLOParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
