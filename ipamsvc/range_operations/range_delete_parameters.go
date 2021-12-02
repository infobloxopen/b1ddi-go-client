// Code generated by go-swagger; DO NOT EDIT.

package range_operations

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

// NewRangeDeleteParams creates a new RangeDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRangeDeleteParams() *RangeDeleteParams {
	return &RangeDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRangeDeleteParamsWithTimeout creates a new RangeDeleteParams object
// with the ability to set a timeout on a request.
func NewRangeDeleteParamsWithTimeout(timeout time.Duration) *RangeDeleteParams {
	return &RangeDeleteParams{
		timeout: timeout,
	}
}

// NewRangeDeleteParamsWithContext creates a new RangeDeleteParams object
// with the ability to set a context for a request.
func NewRangeDeleteParamsWithContext(ctx context.Context) *RangeDeleteParams {
	return &RangeDeleteParams{
		Context: ctx,
	}
}

// NewRangeDeleteParamsWithHTTPClient creates a new RangeDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewRangeDeleteParamsWithHTTPClient(client *http.Client) *RangeDeleteParams {
	return &RangeDeleteParams{
		HTTPClient: client,
	}
}

/* RangeDeleteParams contains all the parameters to send to the API endpoint
   for the range delete operation.

   Typically these are written to a http.Request.
*/
type RangeDeleteParams struct {

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the range delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RangeDeleteParams) WithDefaults() *RangeDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the range delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RangeDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the range delete params
func (o *RangeDeleteParams) WithTimeout(timeout time.Duration) *RangeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the range delete params
func (o *RangeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the range delete params
func (o *RangeDeleteParams) WithContext(ctx context.Context) *RangeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the range delete params
func (o *RangeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the range delete params
func (o *RangeDeleteParams) WithHTTPClient(client *http.Client) *RangeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the range delete params
func (o *RangeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the range delete params
func (o *RangeDeleteParams) WithID(id string) *RangeDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the range delete params
func (o *RangeDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RangeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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