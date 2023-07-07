// Code generated by go-swagger; DO NOT EDIT.

package dfp

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

// NewDfpListDfpParams creates a new DfpListDfpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDfpListDfpParams() *DfpListDfpParams {
	return &DfpListDfpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDfpListDfpParamsWithTimeout creates a new DfpListDfpParams object
// with the ability to set a timeout on a request.
func NewDfpListDfpParamsWithTimeout(timeout time.Duration) *DfpListDfpParams {
	return &DfpListDfpParams{
		timeout: timeout,
	}
}

// NewDfpListDfpParamsWithContext creates a new DfpListDfpParams object
// with the ability to set a context for a request.
func NewDfpListDfpParamsWithContext(ctx context.Context) *DfpListDfpParams {
	return &DfpListDfpParams{
		Context: ctx,
	}
}

// NewDfpListDfpParamsWithHTTPClient creates a new DfpListDfpParams object
// with the ability to set a custom HTTPClient for a request.
func NewDfpListDfpParamsWithHTTPClient(client *http.Client) *DfpListDfpParams {
	return &DfpListDfpParams{
		HTTPClient: client,
	}
}

/*
DfpListDfpParams contains all the parameters to send to the API endpoint

	for the dfp list dfp operation.

	Typically these are written to a http.Request.
*/
type DfpListDfpParams struct {

	/* Fields.



	A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.

	Specify this parameter as a comma-separated list of JSON tag names.


	*/
	Fields *string

	/* Filter.

	     A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.

	Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.

	You can filter by following fields:

	| Name                    | type   | Supported Op                |
	| ----------------------- | ------ | --------------------------- |
	| name                    | string | !=, ==, ~, !~, >, <, <=, >= |
	| site_id                 | string | !=, ==, ~, !~, >, <, <=, >= |
	| ophid                   | string | !=, ==, ~, !~, >, <, <=, >= |
	| policy_id               | int32  | !=, ==, >, <, <=, >=        |
	| default_security_policy | bool   | !=, ==                      |

	In addition groupping operators are supported:

	| Op  | Description          |
	| --- | -------------------- |
	| and | Logical AND          |
	| or  | Logical OR           |
	| not | Logical NOT          |
	| ()  | Groupping Operators  |

	Example:
	```
	?_filter="((name=='dfp1')or(ophid~'oph'))and(default_security_policy!='true')"
	```

	*/
	Filter *string

	/* Limit.



	The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.


	*/
	Limit *int64

	/* Offset.



	The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.


	*/
	Offset *int64

	/* PageToken.



	The service-defined string used to identify a page of resources. A null value indicates the first page.


	*/
	PageToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dfp list dfp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DfpListDfpParams) WithDefaults() *DfpListDfpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dfp list dfp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DfpListDfpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dfp list dfp params
func (o *DfpListDfpParams) WithTimeout(timeout time.Duration) *DfpListDfpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dfp list dfp params
func (o *DfpListDfpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dfp list dfp params
func (o *DfpListDfpParams) WithContext(ctx context.Context) *DfpListDfpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dfp list dfp params
func (o *DfpListDfpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dfp list dfp params
func (o *DfpListDfpParams) WithHTTPClient(client *http.Client) *DfpListDfpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dfp list dfp params
func (o *DfpListDfpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the dfp list dfp params
func (o *DfpListDfpParams) WithFields(fields *string) *DfpListDfpParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the dfp list dfp params
func (o *DfpListDfpParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the dfp list dfp params
func (o *DfpListDfpParams) WithFilter(filter *string) *DfpListDfpParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the dfp list dfp params
func (o *DfpListDfpParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the dfp list dfp params
func (o *DfpListDfpParams) WithLimit(limit *int64) *DfpListDfpParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dfp list dfp params
func (o *DfpListDfpParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the dfp list dfp params
func (o *DfpListDfpParams) WithOffset(offset *int64) *DfpListDfpParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dfp list dfp params
func (o *DfpListDfpParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPageToken adds the pageToken to the dfp list dfp params
func (o *DfpListDfpParams) WithPageToken(pageToken *string) *DfpListDfpParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the dfp list dfp params
func (o *DfpListDfpParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WriteToRequest writes these params to a swagger request
func (o *DfpListDfpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param _fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("_fields", qFields); err != nil {
				return err
			}
		}
	}

	if o.Filter != nil {

		// query param _filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("_filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param _limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("_limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param _offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("_offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.PageToken != nil {

		// query param _page_token
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("_page_token", qPageToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
