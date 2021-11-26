// Code generated by go-swagger; DO NOT EDIT.

package dns_usage

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

// NewDNSUsageReadParams creates a new DNSUsageReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDNSUsageReadParams() *DNSUsageReadParams {
	return &DNSUsageReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDNSUsageReadParamsWithTimeout creates a new DNSUsageReadParams object
// with the ability to set a timeout on a request.
func NewDNSUsageReadParamsWithTimeout(timeout time.Duration) *DNSUsageReadParams {
	return &DNSUsageReadParams{
		timeout: timeout,
	}
}

// NewDNSUsageReadParamsWithContext creates a new DNSUsageReadParams object
// with the ability to set a context for a request.
func NewDNSUsageReadParamsWithContext(ctx context.Context) *DNSUsageReadParams {
	return &DNSUsageReadParams{
		Context: ctx,
	}
}

// NewDNSUsageReadParamsWithHTTPClient creates a new DNSUsageReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDNSUsageReadParamsWithHTTPClient(client *http.Client) *DNSUsageReadParams {
	return &DNSUsageReadParams{
		HTTPClient: client,
	}
}

/* DNSUsageReadParams contains all the parameters to send to the API endpoint
   for the dns usage read operation.

   Typically these are written to a http.Request.
*/
type DNSUsageReadParams struct {

	/* Fields.



	A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.

	Specify this parameter as a comma-separated list of JSON tag names.


	*/
	Fields *string

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dns usage read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DNSUsageReadParams) WithDefaults() *DNSUsageReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dns usage read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DNSUsageReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dns usage read params
func (o *DNSUsageReadParams) WithTimeout(timeout time.Duration) *DNSUsageReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dns usage read params
func (o *DNSUsageReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dns usage read params
func (o *DNSUsageReadParams) WithContext(ctx context.Context) *DNSUsageReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dns usage read params
func (o *DNSUsageReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dns usage read params
func (o *DNSUsageReadParams) WithHTTPClient(client *http.Client) *DNSUsageReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dns usage read params
func (o *DNSUsageReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the dns usage read params
func (o *DNSUsageReadParams) WithFields(fields *string) *DNSUsageReadParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the dns usage read params
func (o *DNSUsageReadParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the dns usage read params
func (o *DNSUsageReadParams) WithID(id string) *DNSUsageReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dns usage read params
func (o *DNSUsageReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DNSUsageReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
