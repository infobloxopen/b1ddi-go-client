// Code generated by go-swagger; DO NOT EDIT.

package security_policies

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

// NewSecurityPoliciesListSecurityPoliciesParams creates a new SecurityPoliciesListSecurityPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityPoliciesListSecurityPoliciesParams() *SecurityPoliciesListSecurityPoliciesParams {
	return &SecurityPoliciesListSecurityPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityPoliciesListSecurityPoliciesParamsWithTimeout creates a new SecurityPoliciesListSecurityPoliciesParams object
// with the ability to set a timeout on a request.
func NewSecurityPoliciesListSecurityPoliciesParamsWithTimeout(timeout time.Duration) *SecurityPoliciesListSecurityPoliciesParams {
	return &SecurityPoliciesListSecurityPoliciesParams{
		timeout: timeout,
	}
}

// NewSecurityPoliciesListSecurityPoliciesParamsWithContext creates a new SecurityPoliciesListSecurityPoliciesParams object
// with the ability to set a context for a request.
func NewSecurityPoliciesListSecurityPoliciesParamsWithContext(ctx context.Context) *SecurityPoliciesListSecurityPoliciesParams {
	return &SecurityPoliciesListSecurityPoliciesParams{
		Context: ctx,
	}
}

// NewSecurityPoliciesListSecurityPoliciesParamsWithHTTPClient creates a new SecurityPoliciesListSecurityPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityPoliciesListSecurityPoliciesParamsWithHTTPClient(client *http.Client) *SecurityPoliciesListSecurityPoliciesParams {
	return &SecurityPoliciesListSecurityPoliciesParams{
		HTTPClient: client,
	}
}

/*
SecurityPoliciesListSecurityPoliciesParams contains all the parameters to send to the API endpoint

	for the security policies list security policies operation.

	Typically these are written to a http.Request.
*/
type SecurityPoliciesListSecurityPoliciesParams struct {

	/* Fields.



	A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.

	Specify this parameter as a comma-separated list of JSON tag names.


	*/
	Fields *string

	/* Filter.

	     A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.

	Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.

	You can filter by following fields:

	| Name               | type   | Supported Op                |
	| ------------------ | ------ | --------------------------- |
	| id                 | int32  | !=, ==, >, <, <=, >=        |
	| name               | string | !=, ==, ~, !~, >, <, <=, >= |
	| description        | string | !=, ==, ~, !~, >, <, <=, >= |
	| is_default         | bool   | !=, ==                      |

	In addition grouping operators are supported:

	| Op  | Description          |
	| --- | -------------------- |
	| and | Logical AND          |
	| or  | Logical OR           |
	| not | Logical NOT          |
	| ()  | Groupping Operators  |

	Example:
	```
	?_filter="((name=='sec_policy_a')or(name~'policy_b'))and(is_default!='true')"
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

	/* Tfilter.

	   Filtering by tags.
	*/
	Tfilter *string

	/* TorderBy.

	   Sorting by tags.
	*/
	TorderBy *string

	// IncludeAccessCodes.
	//
	// Format: boolean
	IncludeAccessCodes *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security policies list security policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityPoliciesListSecurityPoliciesParams) WithDefaults() *SecurityPoliciesListSecurityPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security policies list security policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityPoliciesListSecurityPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) WithTimeout(timeout time.Duration) *SecurityPoliciesListSecurityPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) WithContext(ctx context.Context) *SecurityPoliciesListSecurityPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) WithHTTPClient(client *http.Client) *SecurityPoliciesListSecurityPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) WithFields(fields *string) *SecurityPoliciesListSecurityPoliciesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) WithFilter(filter *string) *SecurityPoliciesListSecurityPoliciesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) WithLimit(limit *int64) *SecurityPoliciesListSecurityPoliciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) WithOffset(offset *int64) *SecurityPoliciesListSecurityPoliciesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPageToken adds the pageToken to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) WithPageToken(pageToken *string) *SecurityPoliciesListSecurityPoliciesParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithTfilter adds the tfilter to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) WithTfilter(tfilter *string) *SecurityPoliciesListSecurityPoliciesParams {
	o.SetTfilter(tfilter)
	return o
}

// SetTfilter adds the tfilter to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) SetTfilter(tfilter *string) {
	o.Tfilter = tfilter
}

// WithTorderBy adds the torderBy to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) WithTorderBy(torderBy *string) *SecurityPoliciesListSecurityPoliciesParams {
	o.SetTorderBy(torderBy)
	return o
}

// SetTorderBy adds the torderBy to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) SetTorderBy(torderBy *string) {
	o.TorderBy = torderBy
}

// WithIncludeAccessCodes adds the includeAccessCodes to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) WithIncludeAccessCodes(includeAccessCodes *bool) *SecurityPoliciesListSecurityPoliciesParams {
	o.SetIncludeAccessCodes(includeAccessCodes)
	return o
}

// SetIncludeAccessCodes adds the includeAccessCodes to the security policies list security policies params
func (o *SecurityPoliciesListSecurityPoliciesParams) SetIncludeAccessCodes(includeAccessCodes *bool) {
	o.IncludeAccessCodes = includeAccessCodes
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityPoliciesListSecurityPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Tfilter != nil {

		// query param _tfilter
		var qrTfilter string

		if o.Tfilter != nil {
			qrTfilter = *o.Tfilter
		}
		qTfilter := qrTfilter
		if qTfilter != "" {

			if err := r.SetQueryParam("_tfilter", qTfilter); err != nil {
				return err
			}
		}
	}

	if o.TorderBy != nil {

		// query param _torder_by
		var qrTorderBy string

		if o.TorderBy != nil {
			qrTorderBy = *o.TorderBy
		}
		qTorderBy := qrTorderBy
		if qTorderBy != "" {

			if err := r.SetQueryParam("_torder_by", qTorderBy); err != nil {
				return err
			}
		}
	}

	if o.IncludeAccessCodes != nil {

		// query param include_access_codes
		var qrIncludeAccessCodes bool

		if o.IncludeAccessCodes != nil {
			qrIncludeAccessCodes = *o.IncludeAccessCodes
		}
		qIncludeAccessCodes := swag.FormatBool(qrIncludeAccessCodes)
		if qIncludeAccessCodes != "" {

			if err := r.SetQueryParam("include_access_codes", qIncludeAccessCodes); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
