// Code generated by go-swagger; DO NOT EDIT.

package subnet

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

// NewSubnetDeleteParams creates a new SubnetDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubnetDeleteParams() *SubnetDeleteParams {
	return &SubnetDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubnetDeleteParamsWithTimeout creates a new SubnetDeleteParams object
// with the ability to set a timeout on a request.
func NewSubnetDeleteParamsWithTimeout(timeout time.Duration) *SubnetDeleteParams {
	return &SubnetDeleteParams{
		timeout: timeout,
	}
}

// NewSubnetDeleteParamsWithContext creates a new SubnetDeleteParams object
// with the ability to set a context for a request.
func NewSubnetDeleteParamsWithContext(ctx context.Context) *SubnetDeleteParams {
	return &SubnetDeleteParams{
		Context: ctx,
	}
}

// NewSubnetDeleteParamsWithHTTPClient creates a new SubnetDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubnetDeleteParamsWithHTTPClient(client *http.Client) *SubnetDeleteParams {
	return &SubnetDeleteParams{
		HTTPClient: client,
	}
}

/* SubnetDeleteParams contains all the parameters to send to the API endpoint
   for the subnet delete operation.

   Typically these are written to a http.Request.
*/
type SubnetDeleteParams struct {

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subnet delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubnetDeleteParams) WithDefaults() *SubnetDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subnet delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubnetDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subnet delete params
func (o *SubnetDeleteParams) WithTimeout(timeout time.Duration) *SubnetDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subnet delete params
func (o *SubnetDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subnet delete params
func (o *SubnetDeleteParams) WithContext(ctx context.Context) *SubnetDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subnet delete params
func (o *SubnetDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subnet delete params
func (o *SubnetDeleteParams) WithHTTPClient(client *http.Client) *SubnetDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subnet delete params
func (o *SubnetDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the subnet delete params
func (o *SubnetDeleteParams) WithID(id string) *SubnetDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the subnet delete params
func (o *SubnetDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SubnetDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
