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

// NewUpdateObjectDiscoveryRuleParams creates a new UpdateObjectDiscoveryRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateObjectDiscoveryRuleParams() *UpdateObjectDiscoveryRuleParams {
	return &UpdateObjectDiscoveryRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateObjectDiscoveryRuleParamsWithTimeout creates a new UpdateObjectDiscoveryRuleParams object
// with the ability to set a timeout on a request.
func NewUpdateObjectDiscoveryRuleParamsWithTimeout(timeout time.Duration) *UpdateObjectDiscoveryRuleParams {
	return &UpdateObjectDiscoveryRuleParams{
		timeout: timeout,
	}
}

// NewUpdateObjectDiscoveryRuleParamsWithContext creates a new UpdateObjectDiscoveryRuleParams object
// with the ability to set a context for a request.
func NewUpdateObjectDiscoveryRuleParamsWithContext(ctx context.Context) *UpdateObjectDiscoveryRuleParams {
	return &UpdateObjectDiscoveryRuleParams{
		Context: ctx,
	}
}

// NewUpdateObjectDiscoveryRuleParamsWithHTTPClient creates a new UpdateObjectDiscoveryRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateObjectDiscoveryRuleParamsWithHTTPClient(client *http.Client) *UpdateObjectDiscoveryRuleParams {
	return &UpdateObjectDiscoveryRuleParams{
		HTTPClient: client,
	}
}

/*
UpdateObjectDiscoveryRuleParams contains all the parameters to send to the API endpoint

	for the update object discovery rule operation.

	Typically these are written to a http.Request.
*/
type UpdateObjectDiscoveryRuleParams struct {

	// Body.
	Body *models.ConfigUnstableUpdateObjectDiscoveryRuleBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update object discovery rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateObjectDiscoveryRuleParams) WithDefaults() *UpdateObjectDiscoveryRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update object discovery rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateObjectDiscoveryRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update object discovery rule params
func (o *UpdateObjectDiscoveryRuleParams) WithTimeout(timeout time.Duration) *UpdateObjectDiscoveryRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update object discovery rule params
func (o *UpdateObjectDiscoveryRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update object discovery rule params
func (o *UpdateObjectDiscoveryRuleParams) WithContext(ctx context.Context) *UpdateObjectDiscoveryRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update object discovery rule params
func (o *UpdateObjectDiscoveryRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update object discovery rule params
func (o *UpdateObjectDiscoveryRuleParams) WithHTTPClient(client *http.Client) *UpdateObjectDiscoveryRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update object discovery rule params
func (o *UpdateObjectDiscoveryRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update object discovery rule params
func (o *UpdateObjectDiscoveryRuleParams) WithBody(body *models.ConfigUnstableUpdateObjectDiscoveryRuleBody) *UpdateObjectDiscoveryRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update object discovery rule params
func (o *UpdateObjectDiscoveryRuleParams) SetBody(body *models.ConfigUnstableUpdateObjectDiscoveryRuleBody) {
	o.Body = body
}

// WithSlug adds the slug to the update object discovery rule params
func (o *UpdateObjectDiscoveryRuleParams) WithSlug(slug string) *UpdateObjectDiscoveryRuleParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update object discovery rule params
func (o *UpdateObjectDiscoveryRuleParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateObjectDiscoveryRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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