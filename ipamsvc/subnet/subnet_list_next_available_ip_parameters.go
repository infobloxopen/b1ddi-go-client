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
	"github.com/go-openapi/swag"
)

// NewSubnetListNextAvailableIPParams creates a new SubnetListNextAvailableIPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubnetListNextAvailableIPParams() *SubnetListNextAvailableIPParams {
	return &SubnetListNextAvailableIPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubnetListNextAvailableIPParamsWithTimeout creates a new SubnetListNextAvailableIPParams object
// with the ability to set a timeout on a request.
func NewSubnetListNextAvailableIPParamsWithTimeout(timeout time.Duration) *SubnetListNextAvailableIPParams {
	return &SubnetListNextAvailableIPParams{
		timeout: timeout,
	}
}

// NewSubnetListNextAvailableIPParamsWithContext creates a new SubnetListNextAvailableIPParams object
// with the ability to set a context for a request.
func NewSubnetListNextAvailableIPParamsWithContext(ctx context.Context) *SubnetListNextAvailableIPParams {
	return &SubnetListNextAvailableIPParams{
		Context: ctx,
	}
}

// NewSubnetListNextAvailableIPParamsWithHTTPClient creates a new SubnetListNextAvailableIPParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubnetListNextAvailableIPParamsWithHTTPClient(client *http.Client) *SubnetListNextAvailableIPParams {
	return &SubnetListNextAvailableIPParams{
		HTTPClient: client,
	}
}

/* SubnetListNextAvailableIPParams contains all the parameters to send to the API endpoint
   for the subnet list next available IP operation.

   Typically these are written to a http.Request.
*/
type SubnetListNextAvailableIPParams struct {

	/* Contiguous.

	     Indicates whether the IP addresses should belong to a contiguous block.

	Defaults to _false_.

	     Format: boolean
	*/
	Contiguous *bool

	/* Count.

	     The number of IP addresses requested.

	Defaults to 1.

	     Format: int32
	*/
	Count *int32

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subnet list next available IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubnetListNextAvailableIPParams) WithDefaults() *SubnetListNextAvailableIPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subnet list next available IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubnetListNextAvailableIPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) WithTimeout(timeout time.Duration) *SubnetListNextAvailableIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) WithContext(ctx context.Context) *SubnetListNextAvailableIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) WithHTTPClient(client *http.Client) *SubnetListNextAvailableIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContiguous adds the contiguous to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) WithContiguous(contiguous *bool) *SubnetListNextAvailableIPParams {
	o.SetContiguous(contiguous)
	return o
}

// SetContiguous adds the contiguous to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) SetContiguous(contiguous *bool) {
	o.Contiguous = contiguous
}

// WithCount adds the count to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) WithCount(count *int32) *SubnetListNextAvailableIPParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) SetCount(count *int32) {
	o.Count = count
}

// WithID adds the id to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) WithID(id string) *SubnetListNextAvailableIPParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the subnet list next available IP params
func (o *SubnetListNextAvailableIPParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SubnetListNextAvailableIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Contiguous != nil {

		// query param contiguous
		var qrContiguous bool

		if o.Contiguous != nil {
			qrContiguous = *o.Contiguous
		}
		qContiguous := swag.FormatBool(qrContiguous)
		if qContiguous != "" {

			if err := r.SetQueryParam("contiguous", qContiguous); err != nil {
				return err
			}
		}
	}

	if o.Count != nil {

		// query param count
		var qrCount int32

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
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
