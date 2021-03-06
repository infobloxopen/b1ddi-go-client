// Code generated by go-swagger; DO NOT EDIT.

package option_filter

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

// NewOptionFilterDeleteParams creates a new OptionFilterDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOptionFilterDeleteParams() *OptionFilterDeleteParams {
	return &OptionFilterDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOptionFilterDeleteParamsWithTimeout creates a new OptionFilterDeleteParams object
// with the ability to set a timeout on a request.
func NewOptionFilterDeleteParamsWithTimeout(timeout time.Duration) *OptionFilterDeleteParams {
	return &OptionFilterDeleteParams{
		timeout: timeout,
	}
}

// NewOptionFilterDeleteParamsWithContext creates a new OptionFilterDeleteParams object
// with the ability to set a context for a request.
func NewOptionFilterDeleteParamsWithContext(ctx context.Context) *OptionFilterDeleteParams {
	return &OptionFilterDeleteParams{
		Context: ctx,
	}
}

// NewOptionFilterDeleteParamsWithHTTPClient creates a new OptionFilterDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewOptionFilterDeleteParamsWithHTTPClient(client *http.Client) *OptionFilterDeleteParams {
	return &OptionFilterDeleteParams{
		HTTPClient: client,
	}
}

/* OptionFilterDeleteParams contains all the parameters to send to the API endpoint
   for the option filter delete operation.

   Typically these are written to a http.Request.
*/
type OptionFilterDeleteParams struct {

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the option filter delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OptionFilterDeleteParams) WithDefaults() *OptionFilterDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the option filter delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OptionFilterDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the option filter delete params
func (o *OptionFilterDeleteParams) WithTimeout(timeout time.Duration) *OptionFilterDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the option filter delete params
func (o *OptionFilterDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the option filter delete params
func (o *OptionFilterDeleteParams) WithContext(ctx context.Context) *OptionFilterDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the option filter delete params
func (o *OptionFilterDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the option filter delete params
func (o *OptionFilterDeleteParams) WithHTTPClient(client *http.Client) *OptionFilterDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the option filter delete params
func (o *OptionFilterDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the option filter delete params
func (o *OptionFilterDeleteParams) WithID(id string) *OptionFilterDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the option filter delete params
func (o *OptionFilterDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OptionFilterDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
