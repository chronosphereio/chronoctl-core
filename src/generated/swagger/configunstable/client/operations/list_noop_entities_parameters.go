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

// NewListNoopEntitiesParams creates a new ListNoopEntitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListNoopEntitiesParams() *ListNoopEntitiesParams {
	return &ListNoopEntitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListNoopEntitiesParamsWithTimeout creates a new ListNoopEntitiesParams object
// with the ability to set a timeout on a request.
func NewListNoopEntitiesParamsWithTimeout(timeout time.Duration) *ListNoopEntitiesParams {
	return &ListNoopEntitiesParams{
		timeout: timeout,
	}
}

// NewListNoopEntitiesParamsWithContext creates a new ListNoopEntitiesParams object
// with the ability to set a context for a request.
func NewListNoopEntitiesParamsWithContext(ctx context.Context) *ListNoopEntitiesParams {
	return &ListNoopEntitiesParams{
		Context: ctx,
	}
}

// NewListNoopEntitiesParamsWithHTTPClient creates a new ListNoopEntitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListNoopEntitiesParamsWithHTTPClient(client *http.Client) *ListNoopEntitiesParams {
	return &ListNoopEntitiesParams{
		HTTPClient: client,
	}
}

/*
ListNoopEntitiesParams contains all the parameters to send to the API endpoint

	for the list noop entities operation.

	Typically these are written to a http.Request.
*/
type ListNoopEntitiesParams struct {

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

	   Filters results by slug, where any NoopEntity with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list noop entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNoopEntitiesParams) WithDefaults() *ListNoopEntitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list noop entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNoopEntitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list noop entities params
func (o *ListNoopEntitiesParams) WithTimeout(timeout time.Duration) *ListNoopEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list noop entities params
func (o *ListNoopEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list noop entities params
func (o *ListNoopEntitiesParams) WithContext(ctx context.Context) *ListNoopEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list noop entities params
func (o *ListNoopEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list noop entities params
func (o *ListNoopEntitiesParams) WithHTTPClient(client *http.Client) *ListNoopEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list noop entities params
func (o *ListNoopEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageMaxSize adds the pageMaxSize to the list noop entities params
func (o *ListNoopEntitiesParams) WithPageMaxSize(pageMaxSize *int64) *ListNoopEntitiesParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list noop entities params
func (o *ListNoopEntitiesParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list noop entities params
func (o *ListNoopEntitiesParams) WithPageToken(pageToken *string) *ListNoopEntitiesParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list noop entities params
func (o *ListNoopEntitiesParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithSlugs adds the slugs to the list noop entities params
func (o *ListNoopEntitiesParams) WithSlugs(slugs []string) *ListNoopEntitiesParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list noop entities params
func (o *ListNoopEntitiesParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WriteToRequest writes these params to a swagger request
func (o *ListNoopEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

// bindParamListNoopEntities binds the parameter slugs
func (o *ListNoopEntitiesParams) bindParamSlugs(formats strfmt.Registry) []string {
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
