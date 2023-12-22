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

// NewDeleteRecordingRuleParams creates a new DeleteRecordingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRecordingRuleParams() *DeleteRecordingRuleParams {
	return &DeleteRecordingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRecordingRuleParamsWithTimeout creates a new DeleteRecordingRuleParams object
// with the ability to set a timeout on a request.
func NewDeleteRecordingRuleParamsWithTimeout(timeout time.Duration) *DeleteRecordingRuleParams {
	return &DeleteRecordingRuleParams{
		timeout: timeout,
	}
}

// NewDeleteRecordingRuleParamsWithContext creates a new DeleteRecordingRuleParams object
// with the ability to set a context for a request.
func NewDeleteRecordingRuleParamsWithContext(ctx context.Context) *DeleteRecordingRuleParams {
	return &DeleteRecordingRuleParams{
		Context: ctx,
	}
}

// NewDeleteRecordingRuleParamsWithHTTPClient creates a new DeleteRecordingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRecordingRuleParamsWithHTTPClient(client *http.Client) *DeleteRecordingRuleParams {
	return &DeleteRecordingRuleParams{
		HTTPClient: client,
	}
}

/*
DeleteRecordingRuleParams contains all the parameters to send to the API endpoint

	for the delete recording rule operation.

	Typically these are written to a http.Request.
*/
type DeleteRecordingRuleParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete recording rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRecordingRuleParams) WithDefaults() *DeleteRecordingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete recording rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRecordingRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete recording rule params
func (o *DeleteRecordingRuleParams) WithTimeout(timeout time.Duration) *DeleteRecordingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete recording rule params
func (o *DeleteRecordingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete recording rule params
func (o *DeleteRecordingRuleParams) WithContext(ctx context.Context) *DeleteRecordingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete recording rule params
func (o *DeleteRecordingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete recording rule params
func (o *DeleteRecordingRuleParams) WithHTTPClient(client *http.Client) *DeleteRecordingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete recording rule params
func (o *DeleteRecordingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the delete recording rule params
func (o *DeleteRecordingRuleParams) WithSlug(slug string) *DeleteRecordingRuleParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete recording rule params
func (o *DeleteRecordingRuleParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRecordingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
