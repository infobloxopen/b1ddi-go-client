// Code generated by go-swagger; DO NOT EDIT.

package address_block

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

// NewAddressBlockReadParams creates a new AddressBlockReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressBlockReadParams() *AddressBlockReadParams {
	return &AddressBlockReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressBlockReadParamsWithTimeout creates a new AddressBlockReadParams object
// with the ability to set a timeout on a request.
func NewAddressBlockReadParamsWithTimeout(timeout time.Duration) *AddressBlockReadParams {
	return &AddressBlockReadParams{
		timeout: timeout,
	}
}

// NewAddressBlockReadParamsWithContext creates a new AddressBlockReadParams object
// with the ability to set a context for a request.
func NewAddressBlockReadParamsWithContext(ctx context.Context) *AddressBlockReadParams {
	return &AddressBlockReadParams{
		Context: ctx,
	}
}

// NewAddressBlockReadParamsWithHTTPClient creates a new AddressBlockReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressBlockReadParamsWithHTTPClient(client *http.Client) *AddressBlockReadParams {
	return &AddressBlockReadParams{
		HTTPClient: client,
	}
}

/* AddressBlockReadParams contains all the parameters to send to the API endpoint
   for the address block read operation.

   Typically these are written to a http.Request.
*/
type AddressBlockReadParams struct {

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

// WithDefaults hydrates default values in the address block read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressBlockReadParams) WithDefaults() *AddressBlockReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address block read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressBlockReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address block read params
func (o *AddressBlockReadParams) WithTimeout(timeout time.Duration) *AddressBlockReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address block read params
func (o *AddressBlockReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address block read params
func (o *AddressBlockReadParams) WithContext(ctx context.Context) *AddressBlockReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address block read params
func (o *AddressBlockReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address block read params
func (o *AddressBlockReadParams) WithHTTPClient(client *http.Client) *AddressBlockReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address block read params
func (o *AddressBlockReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the address block read params
func (o *AddressBlockReadParams) WithFields(fields *string) *AddressBlockReadParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the address block read params
func (o *AddressBlockReadParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the address block read params
func (o *AddressBlockReadParams) WithID(id string) *AddressBlockReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the address block read params
func (o *AddressBlockReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddressBlockReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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