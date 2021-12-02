// Code generated by go-swagger; DO NOT EDIT.

package forward_nsg

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

// NewForwardNsgCreateParams creates a new ForwardNsgCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewForwardNsgCreateParams() *ForwardNsgCreateParams {
	return &ForwardNsgCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewForwardNsgCreateParamsWithTimeout creates a new ForwardNsgCreateParams object
// with the ability to set a timeout on a request.
func NewForwardNsgCreateParamsWithTimeout(timeout time.Duration) *ForwardNsgCreateParams {
	return &ForwardNsgCreateParams{
		timeout: timeout,
	}
}

// NewForwardNsgCreateParamsWithContext creates a new ForwardNsgCreateParams object
// with the ability to set a context for a request.
func NewForwardNsgCreateParamsWithContext(ctx context.Context) *ForwardNsgCreateParams {
	return &ForwardNsgCreateParams{
		Context: ctx,
	}
}

// NewForwardNsgCreateParamsWithHTTPClient creates a new ForwardNsgCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewForwardNsgCreateParamsWithHTTPClient(client *http.Client) *ForwardNsgCreateParams {
	return &ForwardNsgCreateParams{
		HTTPClient: client,
	}
}

/* ForwardNsgCreateParams contains all the parameters to send to the API endpoint
   for the forward nsg create operation.

   Typically these are written to a http.Request.
*/
type ForwardNsgCreateParams struct {

	// Body.
	Body *models.ConfigForwardNSG

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the forward nsg create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForwardNsgCreateParams) WithDefaults() *ForwardNsgCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the forward nsg create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForwardNsgCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the forward nsg create params
func (o *ForwardNsgCreateParams) WithTimeout(timeout time.Duration) *ForwardNsgCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the forward nsg create params
func (o *ForwardNsgCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the forward nsg create params
func (o *ForwardNsgCreateParams) WithContext(ctx context.Context) *ForwardNsgCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the forward nsg create params
func (o *ForwardNsgCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the forward nsg create params
func (o *ForwardNsgCreateParams) WithHTTPClient(client *http.Client) *ForwardNsgCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the forward nsg create params
func (o *ForwardNsgCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the forward nsg create params
func (o *ForwardNsgCreateParams) WithBody(body *models.ConfigForwardNSG) *ForwardNsgCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the forward nsg create params
func (o *ForwardNsgCreateParams) SetBody(body *models.ConfigForwardNSG) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ForwardNsgCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}