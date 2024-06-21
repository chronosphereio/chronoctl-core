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

// NewUpdateLogScaleAlertParams creates a new UpdateLogScaleAlertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLogScaleAlertParams() *UpdateLogScaleAlertParams {
	return &UpdateLogScaleAlertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLogScaleAlertParamsWithTimeout creates a new UpdateLogScaleAlertParams object
// with the ability to set a timeout on a request.
func NewUpdateLogScaleAlertParamsWithTimeout(timeout time.Duration) *UpdateLogScaleAlertParams {
	return &UpdateLogScaleAlertParams{
		timeout: timeout,
	}
}

// NewUpdateLogScaleAlertParamsWithContext creates a new UpdateLogScaleAlertParams object
// with the ability to set a context for a request.
func NewUpdateLogScaleAlertParamsWithContext(ctx context.Context) *UpdateLogScaleAlertParams {
	return &UpdateLogScaleAlertParams{
		Context: ctx,
	}
}

// NewUpdateLogScaleAlertParamsWithHTTPClient creates a new UpdateLogScaleAlertParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLogScaleAlertParamsWithHTTPClient(client *http.Client) *UpdateLogScaleAlertParams {
	return &UpdateLogScaleAlertParams{
		HTTPClient: client,
	}
}

/*
UpdateLogScaleAlertParams contains all the parameters to send to the API endpoint

	for the update log scale alert operation.

	Typically these are written to a http.Request.
*/
type UpdateLogScaleAlertParams struct {

	// Body.
	Body *models.ConfigUnstableUpdateLogScaleAlertBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update log scale alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLogScaleAlertParams) WithDefaults() *UpdateLogScaleAlertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update log scale alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLogScaleAlertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update log scale alert params
func (o *UpdateLogScaleAlertParams) WithTimeout(timeout time.Duration) *UpdateLogScaleAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update log scale alert params
func (o *UpdateLogScaleAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update log scale alert params
func (o *UpdateLogScaleAlertParams) WithContext(ctx context.Context) *UpdateLogScaleAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update log scale alert params
func (o *UpdateLogScaleAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update log scale alert params
func (o *UpdateLogScaleAlertParams) WithHTTPClient(client *http.Client) *UpdateLogScaleAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update log scale alert params
func (o *UpdateLogScaleAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update log scale alert params
func (o *UpdateLogScaleAlertParams) WithBody(body *models.ConfigUnstableUpdateLogScaleAlertBody) *UpdateLogScaleAlertParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update log scale alert params
func (o *UpdateLogScaleAlertParams) SetBody(body *models.ConfigUnstableUpdateLogScaleAlertBody) {
	o.Body = body
}

// WithSlug adds the slug to the update log scale alert params
func (o *UpdateLogScaleAlertParams) WithSlug(slug string) *UpdateLogScaleAlertParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update log scale alert params
func (o *UpdateLogScaleAlertParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLogScaleAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
