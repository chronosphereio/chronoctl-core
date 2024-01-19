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

// NewListDropRulesParams creates a new ListDropRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDropRulesParams() *ListDropRulesParams {
	return &ListDropRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDropRulesParamsWithTimeout creates a new ListDropRulesParams object
// with the ability to set a timeout on a request.
func NewListDropRulesParamsWithTimeout(timeout time.Duration) *ListDropRulesParams {
	return &ListDropRulesParams{
		timeout: timeout,
	}
}

// NewListDropRulesParamsWithContext creates a new ListDropRulesParams object
// with the ability to set a context for a request.
func NewListDropRulesParamsWithContext(ctx context.Context) *ListDropRulesParams {
	return &ListDropRulesParams{
		Context: ctx,
	}
}

// NewListDropRulesParamsWithHTTPClient creates a new ListDropRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDropRulesParamsWithHTTPClient(client *http.Client) *ListDropRulesParams {
	return &ListDropRulesParams{
		HTTPClient: client,
	}
}

/*
ListDropRulesParams contains all the parameters to send to the API endpoint

	for the list drop rules operation.

	Typically these are written to a http.Request.
*/
type ListDropRulesParams struct {

	/* Names.

	   Filters results by name, where any DropRule with a matching name in the given list (and matches all other filters) is returned.
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

	   Filters results by slug, where any DropRule with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list drop rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDropRulesParams) WithDefaults() *ListDropRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list drop rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDropRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list drop rules params
func (o *ListDropRulesParams) WithTimeout(timeout time.Duration) *ListDropRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list drop rules params
func (o *ListDropRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list drop rules params
func (o *ListDropRulesParams) WithContext(ctx context.Context) *ListDropRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list drop rules params
func (o *ListDropRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list drop rules params
func (o *ListDropRulesParams) WithHTTPClient(client *http.Client) *ListDropRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list drop rules params
func (o *ListDropRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNames adds the names to the list drop rules params
func (o *ListDropRulesParams) WithNames(names []string) *ListDropRulesParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the list drop rules params
func (o *ListDropRulesParams) SetNames(names []string) {
	o.Names = names
}

// WithPageMaxSize adds the pageMaxSize to the list drop rules params
func (o *ListDropRulesParams) WithPageMaxSize(pageMaxSize *int64) *ListDropRulesParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list drop rules params
func (o *ListDropRulesParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list drop rules params
func (o *ListDropRulesParams) WithPageToken(pageToken *string) *ListDropRulesParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list drop rules params
func (o *ListDropRulesParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithSlugs adds the slugs to the list drop rules params
func (o *ListDropRulesParams) WithSlugs(slugs []string) *ListDropRulesParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list drop rules params
func (o *ListDropRulesParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WriteToRequest writes these params to a swagger request
func (o *ListDropRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamListDropRules binds the parameter names
func (o *ListDropRulesParams) bindParamNames(formats strfmt.Registry) []string {
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

// bindParamListDropRules binds the parameter slugs
func (o *ListDropRulesParams) bindParamSlugs(formats strfmt.Registry) []string {
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