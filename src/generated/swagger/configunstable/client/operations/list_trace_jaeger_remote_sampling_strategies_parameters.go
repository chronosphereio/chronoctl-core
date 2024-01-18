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

// NewListTraceJaegerRemoteSamplingStrategiesParams creates a new ListTraceJaegerRemoteSamplingStrategiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTraceJaegerRemoteSamplingStrategiesParams() *ListTraceJaegerRemoteSamplingStrategiesParams {
	return &ListTraceJaegerRemoteSamplingStrategiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTraceJaegerRemoteSamplingStrategiesParamsWithTimeout creates a new ListTraceJaegerRemoteSamplingStrategiesParams object
// with the ability to set a timeout on a request.
func NewListTraceJaegerRemoteSamplingStrategiesParamsWithTimeout(timeout time.Duration) *ListTraceJaegerRemoteSamplingStrategiesParams {
	return &ListTraceJaegerRemoteSamplingStrategiesParams{
		timeout: timeout,
	}
}

// NewListTraceJaegerRemoteSamplingStrategiesParamsWithContext creates a new ListTraceJaegerRemoteSamplingStrategiesParams object
// with the ability to set a context for a request.
func NewListTraceJaegerRemoteSamplingStrategiesParamsWithContext(ctx context.Context) *ListTraceJaegerRemoteSamplingStrategiesParams {
	return &ListTraceJaegerRemoteSamplingStrategiesParams{
		Context: ctx,
	}
}

// NewListTraceJaegerRemoteSamplingStrategiesParamsWithHTTPClient creates a new ListTraceJaegerRemoteSamplingStrategiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTraceJaegerRemoteSamplingStrategiesParamsWithHTTPClient(client *http.Client) *ListTraceJaegerRemoteSamplingStrategiesParams {
	return &ListTraceJaegerRemoteSamplingStrategiesParams{
		HTTPClient: client,
	}
}

/*
ListTraceJaegerRemoteSamplingStrategiesParams contains all the parameters to send to the API endpoint

	for the list trace jaeger remote sampling strategies operation.

	Typically these are written to a http.Request.
*/
type ListTraceJaegerRemoteSamplingStrategiesParams struct {

	// NameOrServiceContains.
	NameOrServiceContains *string

	/* Names.

	   Filters results by name, where any TraceJaegerRemoteSamplingStrategy with a matching name in the given list (and matches all other filters) is returned.
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

	/* ServiceNames.

	   Filters results by service_name, where any TraceJaegerRemoteSamplingStrategy with a matching service_name in the given list (and matches all other filters) is returned.
	*/
	ServiceNames []string

	/* Slugs.

	   Filters results by slug, where any TraceJaegerRemoteSamplingStrategy with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list trace jaeger remote sampling strategies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) WithDefaults() *ListTraceJaegerRemoteSamplingStrategiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list trace jaeger remote sampling strategies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) WithTimeout(timeout time.Duration) *ListTraceJaegerRemoteSamplingStrategiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) WithContext(ctx context.Context) *ListTraceJaegerRemoteSamplingStrategiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) WithHTTPClient(client *http.Client) *ListTraceJaegerRemoteSamplingStrategiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNameOrServiceContains adds the nameOrServiceContains to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) WithNameOrServiceContains(nameOrServiceContains *string) *ListTraceJaegerRemoteSamplingStrategiesParams {
	o.SetNameOrServiceContains(nameOrServiceContains)
	return o
}

// SetNameOrServiceContains adds the nameOrServiceContains to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) SetNameOrServiceContains(nameOrServiceContains *string) {
	o.NameOrServiceContains = nameOrServiceContains
}

// WithNames adds the names to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) WithNames(names []string) *ListTraceJaegerRemoteSamplingStrategiesParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) SetNames(names []string) {
	o.Names = names
}

// WithPageMaxSize adds the pageMaxSize to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) WithPageMaxSize(pageMaxSize *int64) *ListTraceJaegerRemoteSamplingStrategiesParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) WithPageToken(pageToken *string) *ListTraceJaegerRemoteSamplingStrategiesParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithServiceNames adds the serviceNames to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) WithServiceNames(serviceNames []string) *ListTraceJaegerRemoteSamplingStrategiesParams {
	o.SetServiceNames(serviceNames)
	return o
}

// SetServiceNames adds the serviceNames to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) SetServiceNames(serviceNames []string) {
	o.ServiceNames = serviceNames
}

// WithSlugs adds the slugs to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) WithSlugs(slugs []string) *ListTraceJaegerRemoteSamplingStrategiesParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list trace jaeger remote sampling strategies params
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WriteToRequest writes these params to a swagger request
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NameOrServiceContains != nil {

		// query param name_or_service_contains
		var qrNameOrServiceContains string

		if o.NameOrServiceContains != nil {
			qrNameOrServiceContains = *o.NameOrServiceContains
		}
		qNameOrServiceContains := qrNameOrServiceContains
		if qNameOrServiceContains != "" {

			if err := r.SetQueryParam("name_or_service_contains", qNameOrServiceContains); err != nil {
				return err
			}
		}
	}

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

	if o.ServiceNames != nil {

		// binding items for service_names
		joinedServiceNames := o.bindParamServiceNames(reg)

		// query array param service_names
		if err := r.SetQueryParam("service_names", joinedServiceNames...); err != nil {
			return err
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

// bindParamListTraceJaegerRemoteSamplingStrategies binds the parameter names
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) bindParamNames(formats strfmt.Registry) []string {
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

// bindParamListTraceJaegerRemoteSamplingStrategies binds the parameter service_names
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) bindParamServiceNames(formats strfmt.Registry) []string {
	serviceNamesIR := o.ServiceNames

	var serviceNamesIC []string
	for _, serviceNamesIIR := range serviceNamesIR { // explode []string

		serviceNamesIIV := serviceNamesIIR // string as string
		serviceNamesIC = append(serviceNamesIC, serviceNamesIIV)
	}

	// items.CollectionFormat: "multi"
	serviceNamesIS := swag.JoinByFormat(serviceNamesIC, "multi")

	return serviceNamesIS
}

// bindParamListTraceJaegerRemoteSamplingStrategies binds the parameter slugs
func (o *ListTraceJaegerRemoteSamplingStrategiesParams) bindParamSlugs(formats strfmt.Registry) []string {
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
