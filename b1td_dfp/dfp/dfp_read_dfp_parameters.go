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

// NewDfpReadDfpParams creates a new DfpReadDfpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDfpReadDfpParams() *DfpReadDfpParams {
	return &DfpReadDfpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDfpReadDfpParamsWithTimeout creates a new DfpReadDfpParams object
// with the ability to set a timeout on a request.
func NewDfpReadDfpParamsWithTimeout(timeout time.Duration) *DfpReadDfpParams {
	return &DfpReadDfpParams{
		timeout: timeout,
	}
}

// NewDfpReadDfpParamsWithContext creates a new DfpReadDfpParams object
// with the ability to set a context for a request.
func NewDfpReadDfpParamsWithContext(ctx context.Context) *DfpReadDfpParams {
	return &DfpReadDfpParams{
		Context: ctx,
	}
}

// NewDfpReadDfpParamsWithHTTPClient creates a new DfpReadDfpParams object
// with the ability to set a custom HTTPClient for a request.
func NewDfpReadDfpParamsWithHTTPClient(client *http.Client) *DfpReadDfpParams {
	return &DfpReadDfpParams{
		HTTPClient: client,
	}
}

/*
DfpReadDfpParams contains all the parameters to send to the API endpoint

	for the dfp read dfp operation.

	Typically these are written to a http.Request.
*/
type DfpReadDfpParams struct {

	/* Fields.



	A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.

	Specify this parameter as a comma-separated list of JSON tag names.


	*/
	Fields *string

	/* ID.

	   The DNS Forwarding Proxy object identifier.

	   Format: int32
	*/
	ID int32

	/* Name.

	     The name of the DNS Forwarding Proxy.
	Used only if the 'id' field is empty.
	*/
	Name *string

	/* ServiceID.

	   The On-Prem Application Service identifier. For internal Use only.
	*/
	ServiceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dfp read dfp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DfpReadDfpParams) WithDefaults() *DfpReadDfpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dfp read dfp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DfpReadDfpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dfp read dfp params
func (o *DfpReadDfpParams) WithTimeout(timeout time.Duration) *DfpReadDfpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dfp read dfp params
func (o *DfpReadDfpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dfp read dfp params
func (o *DfpReadDfpParams) WithContext(ctx context.Context) *DfpReadDfpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dfp read dfp params
func (o *DfpReadDfpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dfp read dfp params
func (o *DfpReadDfpParams) WithHTTPClient(client *http.Client) *DfpReadDfpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dfp read dfp params
func (o *DfpReadDfpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the dfp read dfp params
func (o *DfpReadDfpParams) WithFields(fields *string) *DfpReadDfpParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the dfp read dfp params
func (o *DfpReadDfpParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the dfp read dfp params
func (o *DfpReadDfpParams) WithID(id int32) *DfpReadDfpParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dfp read dfp params
func (o *DfpReadDfpParams) SetID(id int32) {
	o.ID = id
}

// WithName adds the name to the dfp read dfp params
func (o *DfpReadDfpParams) WithName(name *string) *DfpReadDfpParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dfp read dfp params
func (o *DfpReadDfpParams) SetName(name *string) {
	o.Name = name
}

// WithServiceID adds the serviceID to the dfp read dfp params
func (o *DfpReadDfpParams) WithServiceID(serviceID *string) *DfpReadDfpParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the dfp read dfp params
func (o *DfpReadDfpParams) SetServiceID(serviceID *string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DfpReadDfpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.ServiceID != nil {

		// query param service_id
		var qrServiceID string

		if o.ServiceID != nil {
			qrServiceID = *o.ServiceID
		}
		qServiceID := qrServiceID
		if qServiceID != "" {

			if err := r.SetQueryParam("service_id", qServiceID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
