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

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// NewSubnetUpdateParams creates a new SubnetUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubnetUpdateParams() *SubnetUpdateParams {
	return &SubnetUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubnetUpdateParamsWithTimeout creates a new SubnetUpdateParams object
// with the ability to set a timeout on a request.
func NewSubnetUpdateParamsWithTimeout(timeout time.Duration) *SubnetUpdateParams {
	return &SubnetUpdateParams{
		timeout: timeout,
	}
}

// NewSubnetUpdateParamsWithContext creates a new SubnetUpdateParams object
// with the ability to set a context for a request.
func NewSubnetUpdateParamsWithContext(ctx context.Context) *SubnetUpdateParams {
	return &SubnetUpdateParams{
		Context: ctx,
	}
}

// NewSubnetUpdateParamsWithHTTPClient creates a new SubnetUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubnetUpdateParamsWithHTTPClient(client *http.Client) *SubnetUpdateParams {
	return &SubnetUpdateParams{
		HTTPClient: client,
	}
}

/* SubnetUpdateParams contains all the parameters to send to the API endpoint
   for the subnet update operation.

   Typically these are written to a http.Request.
*/
type SubnetUpdateParams struct {

	// Body.
	Body *models.IpamsvcSubnet

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subnet update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubnetUpdateParams) WithDefaults() *SubnetUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subnet update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubnetUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subnet update params
func (o *SubnetUpdateParams) WithTimeout(timeout time.Duration) *SubnetUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subnet update params
func (o *SubnetUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subnet update params
func (o *SubnetUpdateParams) WithContext(ctx context.Context) *SubnetUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subnet update params
func (o *SubnetUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subnet update params
func (o *SubnetUpdateParams) WithHTTPClient(client *http.Client) *SubnetUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subnet update params
func (o *SubnetUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subnet update params
func (o *SubnetUpdateParams) WithBody(body *models.IpamsvcSubnet) *SubnetUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subnet update params
func (o *SubnetUpdateParams) SetBody(body *models.IpamsvcSubnet) {
	o.Body = body
}

// WithID adds the id to the subnet update params
func (o *SubnetUpdateParams) WithID(id string) *SubnetUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the subnet update params
func (o *SubnetUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SubnetUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
