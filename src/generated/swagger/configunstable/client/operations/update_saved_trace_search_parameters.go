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

// NewUpdateSavedTraceSearchParams creates a new UpdateSavedTraceSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSavedTraceSearchParams() *UpdateSavedTraceSearchParams {
	return &UpdateSavedTraceSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSavedTraceSearchParamsWithTimeout creates a new UpdateSavedTraceSearchParams object
// with the ability to set a timeout on a request.
func NewUpdateSavedTraceSearchParamsWithTimeout(timeout time.Duration) *UpdateSavedTraceSearchParams {
	return &UpdateSavedTraceSearchParams{
		timeout: timeout,
	}
}

// NewUpdateSavedTraceSearchParamsWithContext creates a new UpdateSavedTraceSearchParams object
// with the ability to set a context for a request.
func NewUpdateSavedTraceSearchParamsWithContext(ctx context.Context) *UpdateSavedTraceSearchParams {
	return &UpdateSavedTraceSearchParams{
		Context: ctx,
	}
}

// NewUpdateSavedTraceSearchParamsWithHTTPClient creates a new UpdateSavedTraceSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSavedTraceSearchParamsWithHTTPClient(client *http.Client) *UpdateSavedTraceSearchParams {
	return &UpdateSavedTraceSearchParams{
		HTTPClient: client,
	}
}

/*
UpdateSavedTraceSearchParams contains all the parameters to send to the API endpoint

	for the update saved trace search operation.

	Typically these are written to a http.Request.
*/
type UpdateSavedTraceSearchParams struct {

	// Body.
	Body *models.ConfigUnstableUpdateSavedTraceSearchBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update saved trace search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSavedTraceSearchParams) WithDefaults() *UpdateSavedTraceSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update saved trace search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSavedTraceSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update saved trace search params
func (o *UpdateSavedTraceSearchParams) WithTimeout(timeout time.Duration) *UpdateSavedTraceSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update saved trace search params
func (o *UpdateSavedTraceSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update saved trace search params
func (o *UpdateSavedTraceSearchParams) WithContext(ctx context.Context) *UpdateSavedTraceSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update saved trace search params
func (o *UpdateSavedTraceSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update saved trace search params
func (o *UpdateSavedTraceSearchParams) WithHTTPClient(client *http.Client) *UpdateSavedTraceSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update saved trace search params
func (o *UpdateSavedTraceSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update saved trace search params
func (o *UpdateSavedTraceSearchParams) WithBody(body *models.ConfigUnstableUpdateSavedTraceSearchBody) *UpdateSavedTraceSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update saved trace search params
func (o *UpdateSavedTraceSearchParams) SetBody(body *models.ConfigUnstableUpdateSavedTraceSearchBody) {
	o.Body = body
}

// WithSlug adds the slug to the update saved trace search params
func (o *UpdateSavedTraceSearchParams) WithSlug(slug string) *UpdateSavedTraceSearchParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update saved trace search params
func (o *UpdateSavedTraceSearchParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSavedTraceSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
