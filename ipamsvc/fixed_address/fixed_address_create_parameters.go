// Code generated by go-swagger; DO NOT EDIT.

package fixed_address

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

// NewFixedAddressCreateParams creates a new FixedAddressCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFixedAddressCreateParams() *FixedAddressCreateParams {
	return &FixedAddressCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFixedAddressCreateParamsWithTimeout creates a new FixedAddressCreateParams object
// with the ability to set a timeout on a request.
func NewFixedAddressCreateParamsWithTimeout(timeout time.Duration) *FixedAddressCreateParams {
	return &FixedAddressCreateParams{
		timeout: timeout,
	}
}

// NewFixedAddressCreateParamsWithContext creates a new FixedAddressCreateParams object
// with the ability to set a context for a request.
func NewFixedAddressCreateParamsWithContext(ctx context.Context) *FixedAddressCreateParams {
	return &FixedAddressCreateParams{
		Context: ctx,
	}
}

// NewFixedAddressCreateParamsWithHTTPClient creates a new FixedAddressCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewFixedAddressCreateParamsWithHTTPClient(client *http.Client) *FixedAddressCreateParams {
	return &FixedAddressCreateParams{
		HTTPClient: client,
	}
}

/* FixedAddressCreateParams contains all the parameters to send to the API endpoint
   for the fixed address create operation.

   Typically these are written to a http.Request.
*/
type FixedAddressCreateParams struct {

	// Body.
	Body *models.IpamsvcFixedAddress

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fixed address create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FixedAddressCreateParams) WithDefaults() *FixedAddressCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fixed address create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FixedAddressCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fixed address create params
func (o *FixedAddressCreateParams) WithTimeout(timeout time.Duration) *FixedAddressCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fixed address create params
func (o *FixedAddressCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fixed address create params
func (o *FixedAddressCreateParams) WithContext(ctx context.Context) *FixedAddressCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fixed address create params
func (o *FixedAddressCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fixed address create params
func (o *FixedAddressCreateParams) WithHTTPClient(client *http.Client) *FixedAddressCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fixed address create params
func (o *FixedAddressCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the fixed address create params
func (o *FixedAddressCreateParams) WithBody(body *models.IpamsvcFixedAddress) *FixedAddressCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the fixed address create params
func (o *FixedAddressCreateParams) SetBody(body *models.IpamsvcFixedAddress) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FixedAddressCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
