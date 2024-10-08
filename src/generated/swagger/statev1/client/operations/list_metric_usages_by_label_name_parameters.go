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

// NewListMetricUsagesByLabelNameParams creates a new ListMetricUsagesByLabelNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListMetricUsagesByLabelNameParams() *ListMetricUsagesByLabelNameParams {
	return &ListMetricUsagesByLabelNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListMetricUsagesByLabelNameParamsWithTimeout creates a new ListMetricUsagesByLabelNameParams object
// with the ability to set a timeout on a request.
func NewListMetricUsagesByLabelNameParamsWithTimeout(timeout time.Duration) *ListMetricUsagesByLabelNameParams {
	return &ListMetricUsagesByLabelNameParams{
		timeout: timeout,
	}
}

// NewListMetricUsagesByLabelNameParamsWithContext creates a new ListMetricUsagesByLabelNameParams object
// with the ability to set a context for a request.
func NewListMetricUsagesByLabelNameParamsWithContext(ctx context.Context) *ListMetricUsagesByLabelNameParams {
	return &ListMetricUsagesByLabelNameParams{
		Context: ctx,
	}
}

// NewListMetricUsagesByLabelNameParamsWithHTTPClient creates a new ListMetricUsagesByLabelNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewListMetricUsagesByLabelNameParamsWithHTTPClient(client *http.Client) *ListMetricUsagesByLabelNameParams {
	return &ListMetricUsagesByLabelNameParams{
		HTTPClient: client,
	}
}

/*
ListMetricUsagesByLabelNameParams contains all the parameters to send to the API endpoint

	for the list metric usages by label name operation.

	Typically these are written to a http.Request.
*/
type ListMetricUsagesByLabelNameParams struct {

	/* IncludeCountsByType.

	   If `true`, responses include the fields `reference_counts_by_type` and `query_execution_counts_by_type`. If `false`, these fields are returned empty.
	*/
	IncludeCountsByType *bool

	/* LabelNameGlob.

	   Glob match string for filtering results by label name.
	*/
	LabelNameGlob *string

	/* LookbackSecs.

	   Controls the time range over which query executions are included in usages. Defaults to `2592000` (30 days).

	   Format: int32
	*/
	LookbackSecs *int32

	// OrderAscending.
	OrderAscending *bool

	// OrderBy.
	OrderBy *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list metric usages by label name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMetricUsagesByLabelNameParams) WithDefaults() *ListMetricUsagesByLabelNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list metric usages by label name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMetricUsagesByLabelNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) WithTimeout(timeout time.Duration) *ListMetricUsagesByLabelNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) WithContext(ctx context.Context) *ListMetricUsagesByLabelNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) WithHTTPClient(client *http.Client) *ListMetricUsagesByLabelNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeCountsByType adds the includeCountsByType to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) WithIncludeCountsByType(includeCountsByType *bool) *ListMetricUsagesByLabelNameParams {
	o.SetIncludeCountsByType(includeCountsByType)
	return o
}

// SetIncludeCountsByType adds the includeCountsByType to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) SetIncludeCountsByType(includeCountsByType *bool) {
	o.IncludeCountsByType = includeCountsByType
}

// WithLabelNameGlob adds the labelNameGlob to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) WithLabelNameGlob(labelNameGlob *string) *ListMetricUsagesByLabelNameParams {
	o.SetLabelNameGlob(labelNameGlob)
	return o
}

// SetLabelNameGlob adds the labelNameGlob to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) SetLabelNameGlob(labelNameGlob *string) {
	o.LabelNameGlob = labelNameGlob
}

// WithLookbackSecs adds the lookbackSecs to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) WithLookbackSecs(lookbackSecs *int32) *ListMetricUsagesByLabelNameParams {
	o.SetLookbackSecs(lookbackSecs)
	return o
}

// SetLookbackSecs adds the lookbackSecs to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) SetLookbackSecs(lookbackSecs *int32) {
	o.LookbackSecs = lookbackSecs
}

// WithOrderAscending adds the orderAscending to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) WithOrderAscending(orderAscending *bool) *ListMetricUsagesByLabelNameParams {
	o.SetOrderAscending(orderAscending)
	return o
}

// SetOrderAscending adds the orderAscending to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) SetOrderAscending(orderAscending *bool) {
	o.OrderAscending = orderAscending
}

// WithOrderBy adds the orderBy to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) WithOrderBy(orderBy *string) *ListMetricUsagesByLabelNameParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithPageMaxSize adds the pageMaxSize to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) WithPageMaxSize(pageMaxSize *int64) *ListMetricUsagesByLabelNameParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) WithPageToken(pageToken *string) *ListMetricUsagesByLabelNameParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list metric usages by label name params
func (o *ListMetricUsagesByLabelNameParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WriteToRequest writes these params to a swagger request
func (o *ListMetricUsagesByLabelNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeCountsByType != nil {

		// query param include_counts_by_type
		var qrIncludeCountsByType bool

		if o.IncludeCountsByType != nil {
			qrIncludeCountsByType = *o.IncludeCountsByType
		}
		qIncludeCountsByType := swag.FormatBool(qrIncludeCountsByType)
		if qIncludeCountsByType != "" {

			if err := r.SetQueryParam("include_counts_by_type", qIncludeCountsByType); err != nil {
				return err
			}
		}
	}

	if o.LabelNameGlob != nil {

		// query param label_name_glob
		var qrLabelNameGlob string

		if o.LabelNameGlob != nil {
			qrLabelNameGlob = *o.LabelNameGlob
		}
		qLabelNameGlob := qrLabelNameGlob
		if qLabelNameGlob != "" {

			if err := r.SetQueryParam("label_name_glob", qLabelNameGlob); err != nil {
				return err
			}
		}
	}

	if o.LookbackSecs != nil {

		// query param lookback_secs
		var qrLookbackSecs int32

		if o.LookbackSecs != nil {
			qrLookbackSecs = *o.LookbackSecs
		}
		qLookbackSecs := swag.FormatInt32(qrLookbackSecs)
		if qLookbackSecs != "" {

			if err := r.SetQueryParam("lookback_secs", qLookbackSecs); err != nil {
				return err
			}
		}
	}

	if o.OrderAscending != nil {

		// query param order.ascending
		var qrOrderAscending bool

		if o.OrderAscending != nil {
			qrOrderAscending = *o.OrderAscending
		}
		qOrderAscending := swag.FormatBool(qrOrderAscending)
		if qOrderAscending != "" {

			if err := r.SetQueryParam("order.ascending", qOrderAscending); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// query param order.by
		var qrOrderBy string

		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {

			if err := r.SetQueryParam("order.by", qOrderBy); err != nil {
				return err
			}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
