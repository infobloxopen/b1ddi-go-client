// Code generated by go-swagger; DO NOT EDIT.

package app_approvals

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

// NewAppApprovalsListAppApprovalsParams creates a new AppApprovalsListAppApprovalsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppApprovalsListAppApprovalsParams() *AppApprovalsListAppApprovalsParams {
	return &AppApprovalsListAppApprovalsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppApprovalsListAppApprovalsParamsWithTimeout creates a new AppApprovalsListAppApprovalsParams object
// with the ability to set a timeout on a request.
func NewAppApprovalsListAppApprovalsParamsWithTimeout(timeout time.Duration) *AppApprovalsListAppApprovalsParams {
	return &AppApprovalsListAppApprovalsParams{
		timeout: timeout,
	}
}

// NewAppApprovalsListAppApprovalsParamsWithContext creates a new AppApprovalsListAppApprovalsParams object
// with the ability to set a context for a request.
func NewAppApprovalsListAppApprovalsParamsWithContext(ctx context.Context) *AppApprovalsListAppApprovalsParams {
	return &AppApprovalsListAppApprovalsParams{
		Context: ctx,
	}
}

// NewAppApprovalsListAppApprovalsParamsWithHTTPClient creates a new AppApprovalsListAppApprovalsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppApprovalsListAppApprovalsParamsWithHTTPClient(client *http.Client) *AppApprovalsListAppApprovalsParams {
	return &AppApprovalsListAppApprovalsParams{
		HTTPClient: client,
	}
}

/*
AppApprovalsListAppApprovalsParams contains all the parameters to send to the API endpoint

	for the app approvals list app approvals operation.

	Typically these are written to a http.Request.
*/
type AppApprovalsListAppApprovalsParams struct {

	/* Filter.



	A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.

	Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'. The following operators are commonly used in filter expressions:

	|  Op   |  Description               |
	|  --   |  -----------               |
	|  ==   |  Equal                     |
	|  !=   |  Not Equal                 |
	|  >    |  Greater Than              |
	|   >=  |  Greater Than or Equal To  |
	|  <    |  Less Than                 |
	|  <=   |  Less Than or Equal To     |
	|  and  |  Logical AND               |
	|  ~    |  Matches Regex             |
	|  !~   |  Does Not Match Regex      |
	|  or   |  Logical OR                |
	|  not  |  Logical NOT               |
	|  ()   |  Groupping Operators       |


	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the app approvals list app approvals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppApprovalsListAppApprovalsParams) WithDefaults() *AppApprovalsListAppApprovalsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the app approvals list app approvals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppApprovalsListAppApprovalsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the app approvals list app approvals params
func (o *AppApprovalsListAppApprovalsParams) WithTimeout(timeout time.Duration) *AppApprovalsListAppApprovalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the app approvals list app approvals params
func (o *AppApprovalsListAppApprovalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the app approvals list app approvals params
func (o *AppApprovalsListAppApprovalsParams) WithContext(ctx context.Context) *AppApprovalsListAppApprovalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the app approvals list app approvals params
func (o *AppApprovalsListAppApprovalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the app approvals list app approvals params
func (o *AppApprovalsListAppApprovalsParams) WithHTTPClient(client *http.Client) *AppApprovalsListAppApprovalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the app approvals list app approvals params
func (o *AppApprovalsListAppApprovalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the app approvals list app approvals params
func (o *AppApprovalsListAppApprovalsParams) WithFilter(filter *string) *AppApprovalsListAppApprovalsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the app approvals list app approvals params
func (o *AppApprovalsListAppApprovalsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *AppApprovalsListAppApprovalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}