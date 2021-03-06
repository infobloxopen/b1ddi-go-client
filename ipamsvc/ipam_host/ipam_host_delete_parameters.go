// Code generated by go-swagger; DO NOT EDIT.

package ipam_host

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

// NewIpamHostDeleteParams creates a new IpamHostDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamHostDeleteParams() *IpamHostDeleteParams {
	return &IpamHostDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamHostDeleteParamsWithTimeout creates a new IpamHostDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamHostDeleteParamsWithTimeout(timeout time.Duration) *IpamHostDeleteParams {
	return &IpamHostDeleteParams{
		timeout: timeout,
	}
}

// NewIpamHostDeleteParamsWithContext creates a new IpamHostDeleteParams object
// with the ability to set a context for a request.
func NewIpamHostDeleteParamsWithContext(ctx context.Context) *IpamHostDeleteParams {
	return &IpamHostDeleteParams{
		Context: ctx,
	}
}

// NewIpamHostDeleteParamsWithHTTPClient creates a new IpamHostDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamHostDeleteParamsWithHTTPClient(client *http.Client) *IpamHostDeleteParams {
	return &IpamHostDeleteParams{
		HTTPClient: client,
	}
}

/* IpamHostDeleteParams contains all the parameters to send to the API endpoint
   for the ipam host delete operation.

   Typically these are written to a http.Request.
*/
type IpamHostDeleteParams struct {

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam host delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamHostDeleteParams) WithDefaults() *IpamHostDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam host delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamHostDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam host delete params
func (o *IpamHostDeleteParams) WithTimeout(timeout time.Duration) *IpamHostDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam host delete params
func (o *IpamHostDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam host delete params
func (o *IpamHostDeleteParams) WithContext(ctx context.Context) *IpamHostDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam host delete params
func (o *IpamHostDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam host delete params
func (o *IpamHostDeleteParams) WithHTTPClient(client *http.Client) *IpamHostDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam host delete params
func (o *IpamHostDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam host delete params
func (o *IpamHostDeleteParams) WithID(id string) *IpamHostDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam host delete params
func (o *IpamHostDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamHostDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
