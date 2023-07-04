// Code generated by go-swagger; DO NOT EDIT.

package address

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

// NewAddressDeleteParams creates a new AddressDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressDeleteParams() *AddressDeleteParams {
	return &AddressDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressDeleteParamsWithTimeout creates a new AddressDeleteParams object
// with the ability to set a timeout on a request.
func NewAddressDeleteParamsWithTimeout(timeout time.Duration) *AddressDeleteParams {
	return &AddressDeleteParams{
		timeout: timeout,
	}
}

// NewAddressDeleteParamsWithContext creates a new AddressDeleteParams object
// with the ability to set a context for a request.
func NewAddressDeleteParamsWithContext(ctx context.Context) *AddressDeleteParams {
	return &AddressDeleteParams{
		Context: ctx,
	}
}

// NewAddressDeleteParamsWithHTTPClient creates a new AddressDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressDeleteParamsWithHTTPClient(client *http.Client) *AddressDeleteParams {
	return &AddressDeleteParams{
		HTTPClient: client,
	}
}

/*
AddressDeleteParams contains all the parameters to send to the API endpoint

	for the address delete operation.

	Typically these are written to a http.Request.
*/
type AddressDeleteParams struct {

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressDeleteParams) WithDefaults() *AddressDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address delete params
func (o *AddressDeleteParams) WithTimeout(timeout time.Duration) *AddressDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address delete params
func (o *AddressDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address delete params
func (o *AddressDeleteParams) WithContext(ctx context.Context) *AddressDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address delete params
func (o *AddressDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address delete params
func (o *AddressDeleteParams) WithHTTPClient(client *http.Client) *AddressDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address delete params
func (o *AddressDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the address delete params
func (o *AddressDeleteParams) WithID(id string) *AddressDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the address delete params
func (o *AddressDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddressDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
