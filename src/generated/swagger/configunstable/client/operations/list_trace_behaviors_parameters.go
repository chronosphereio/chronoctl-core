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
	"github.com/go-openapi/swag"
)

// NewListTraceBehaviorsParams creates a new ListTraceBehaviorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTraceBehaviorsParams() *ListTraceBehaviorsParams {
	return &ListTraceBehaviorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTraceBehaviorsParamsWithTimeout creates a new ListTraceBehaviorsParams object
// with the ability to set a timeout on a request.
func NewListTraceBehaviorsParamsWithTimeout(timeout time.Duration) *ListTraceBehaviorsParams {
	return &ListTraceBehaviorsParams{
		timeout: timeout,
	}
}

// NewListTraceBehaviorsParamsWithContext creates a new ListTraceBehaviorsParams object
// with the ability to set a context for a request.
func NewListTraceBehaviorsParamsWithContext(ctx context.Context) *ListTraceBehaviorsParams {
	return &ListTraceBehaviorsParams{
		Context: ctx,
	}
}

// NewListTraceBehaviorsParamsWithHTTPClient creates a new ListTraceBehaviorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTraceBehaviorsParamsWithHTTPClient(client *http.Client) *ListTraceBehaviorsParams {
	return &ListTraceBehaviorsParams{
		HTTPClient: client,
	}
}

/*
ListTraceBehaviorsParams contains all the parameters to send to the API endpoint

	for the list trace behaviors operation.

	Typically these are written to a http.Request.
*/
type ListTraceBehaviorsParams struct {

	/* Names.

	   Filters results by name, where any TraceBehavior with a matching name in the given list (and matches all other filters) is returned.
	*/
	Names []string

	/* PageMaxSize.

	     Page size preference (i.e. how many items are returned in the next
	page). If zero, the server will use a default. Regardless of what size
	is given, clients must never assume how many items will be returned.

	     Format: int64
	*/
	PageMaxSize *int64

	/* PageToken.

	     Opaque page token identifying which page to request. An empty token
	identifies the first page.
	*/
	PageToken *string

	/* Slugs.

	   Filters results by slug, where any TraceBehavior with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list trace behaviors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTraceBehaviorsParams) WithDefaults() *ListTraceBehaviorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list trace behaviors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTraceBehaviorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list trace behaviors params
func (o *ListTraceBehaviorsParams) WithTimeout(timeout time.Duration) *ListTraceBehaviorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list trace behaviors params
func (o *ListTraceBehaviorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list trace behaviors params
func (o *ListTraceBehaviorsParams) WithContext(ctx context.Context) *ListTraceBehaviorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list trace behaviors params
func (o *ListTraceBehaviorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list trace behaviors params
func (o *ListTraceBehaviorsParams) WithHTTPClient(client *http.Client) *ListTraceBehaviorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list trace behaviors params
func (o *ListTraceBehaviorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNames adds the names to the list trace behaviors params
func (o *ListTraceBehaviorsParams) WithNames(names []string) *ListTraceBehaviorsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the list trace behaviors params
func (o *ListTraceBehaviorsParams) SetNames(names []string) {
	o.Names = names
}

// WithPageMaxSize adds the pageMaxSize to the list trace behaviors params
func (o *ListTraceBehaviorsParams) WithPageMaxSize(pageMaxSize *int64) *ListTraceBehaviorsParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list trace behaviors params
func (o *ListTraceBehaviorsParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list trace behaviors params
func (o *ListTraceBehaviorsParams) WithPageToken(pageToken *string) *ListTraceBehaviorsParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list trace behaviors params
func (o *ListTraceBehaviorsParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithSlugs adds the slugs to the list trace behaviors params
func (o *ListTraceBehaviorsParams) WithSlugs(slugs []string) *ListTraceBehaviorsParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list trace behaviors params
func (o *ListTraceBehaviorsParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WriteToRequest writes these params to a swagger request
func (o *ListTraceBehaviorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Names != nil {

		// binding items for names
		joinedNames := o.bindParamNames(reg)

		// query array param names
		if err := r.SetQueryParam("names", joinedNames...); err != nil {
			return err
		}
	}

	if o.PageMaxSize != nil {

		// query param page.max_size
		var qrPageMaxSize int64

		if o.PageMaxSize != nil {
			qrPageMaxSize = *o.PageMaxSize
		}
		qPageMaxSize := swag.FormatInt64(qrPageMaxSize)
		if qPageMaxSize != "" {

			if err := r.SetQueryParam("page.max_size", qPageMaxSize); err != nil {
				return err
			}
		}
	}

	if o.PageToken != nil {

		// query param page.token
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("page.token", qPageToken); err != nil {
				return err
			}
		}
	}

	if o.Slugs != nil {

		// binding items for slugs
		joinedSlugs := o.bindParamSlugs(reg)

		// query array param slugs
		if err := r.SetQueryParam("slugs", joinedSlugs...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListTraceBehaviors binds the parameter names
func (o *ListTraceBehaviorsParams) bindParamNames(formats strfmt.Registry) []string {
	namesIR := o.Names

	var namesIC []string
	for _, namesIIR := range namesIR { // explode []string

		namesIIV := namesIIR // string as string
		namesIC = append(namesIC, namesIIV)
	}

	// items.CollectionFormat: "multi"
	namesIS := swag.JoinByFormat(namesIC, "multi")

	return namesIS
}

// bindParamListTraceBehaviors binds the parameter slugs
func (o *ListTraceBehaviorsParams) bindParamSlugs(formats strfmt.Registry) []string {
	slugsIR := o.Slugs

	var slugsIC []string
	for _, slugsIIR := range slugsIR { // explode []string

		slugsIIV := slugsIIR // string as string
		slugsIC = append(slugsIC, slugsIIV)
	}

	// items.CollectionFormat: "multi"
	slugsIS := swag.JoinByFormat(slugsIC, "multi")

	return slugsIS
}
