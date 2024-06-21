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

// NewUpdateDerivedLabelParams creates a new UpdateDerivedLabelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDerivedLabelParams() *UpdateDerivedLabelParams {
	return &UpdateDerivedLabelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDerivedLabelParamsWithTimeout creates a new UpdateDerivedLabelParams object
// with the ability to set a timeout on a request.
func NewUpdateDerivedLabelParamsWithTimeout(timeout time.Duration) *UpdateDerivedLabelParams {
	return &UpdateDerivedLabelParams{
		timeout: timeout,
	}
}

// NewUpdateDerivedLabelParamsWithContext creates a new UpdateDerivedLabelParams object
// with the ability to set a context for a request.
func NewUpdateDerivedLabelParamsWithContext(ctx context.Context) *UpdateDerivedLabelParams {
	return &UpdateDerivedLabelParams{
		Context: ctx,
	}
}

// NewUpdateDerivedLabelParamsWithHTTPClient creates a new UpdateDerivedLabelParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDerivedLabelParamsWithHTTPClient(client *http.Client) *UpdateDerivedLabelParams {
	return &UpdateDerivedLabelParams{
		HTTPClient: client,
	}
}

/*
UpdateDerivedLabelParams contains all the parameters to send to the API endpoint

	for the update derived label operation.

	Typically these are written to a http.Request.
*/
type UpdateDerivedLabelParams struct {

	// Body.
	Body *models.ConfigV1UpdateDerivedLabelBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update derived label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDerivedLabelParams) WithDefaults() *UpdateDerivedLabelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update derived label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDerivedLabelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update derived label params
func (o *UpdateDerivedLabelParams) WithTimeout(timeout time.Duration) *UpdateDerivedLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update derived label params
func (o *UpdateDerivedLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update derived label params
func (o *UpdateDerivedLabelParams) WithContext(ctx context.Context) *UpdateDerivedLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update derived label params
func (o *UpdateDerivedLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update derived label params
func (o *UpdateDerivedLabelParams) WithHTTPClient(client *http.Client) *UpdateDerivedLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update derived label params
func (o *UpdateDerivedLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update derived label params
func (o *UpdateDerivedLabelParams) WithBody(body *models.ConfigV1UpdateDerivedLabelBody) *UpdateDerivedLabelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update derived label params
func (o *UpdateDerivedLabelParams) SetBody(body *models.ConfigV1UpdateDerivedLabelBody) {
	o.Body = body
}

// WithSlug adds the slug to the update derived label params
func (o *UpdateDerivedLabelParams) WithSlug(slug string) *UpdateDerivedLabelParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update derived label params
func (o *UpdateDerivedLabelParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDerivedLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
