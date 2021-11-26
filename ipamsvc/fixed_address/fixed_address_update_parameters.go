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

// NewFixedAddressUpdateParams creates a new FixedAddressUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFixedAddressUpdateParams() *FixedAddressUpdateParams {
	return &FixedAddressUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFixedAddressUpdateParamsWithTimeout creates a new FixedAddressUpdateParams object
// with the ability to set a timeout on a request.
func NewFixedAddressUpdateParamsWithTimeout(timeout time.Duration) *FixedAddressUpdateParams {
	return &FixedAddressUpdateParams{
		timeout: timeout,
	}
}

// NewFixedAddressUpdateParamsWithContext creates a new FixedAddressUpdateParams object
// with the ability to set a context for a request.
func NewFixedAddressUpdateParamsWithContext(ctx context.Context) *FixedAddressUpdateParams {
	return &FixedAddressUpdateParams{
		Context: ctx,
	}
}

// NewFixedAddressUpdateParamsWithHTTPClient creates a new FixedAddressUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewFixedAddressUpdateParamsWithHTTPClient(client *http.Client) *FixedAddressUpdateParams {
	return &FixedAddressUpdateParams{
		HTTPClient: client,
	}
}

/* FixedAddressUpdateParams contains all the parameters to send to the API endpoint
   for the fixed address update operation.

   Typically these are written to a http.Request.
*/
type FixedAddressUpdateParams struct {

	// Body.
	Body *models.IpamsvcFixedAddress

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fixed address update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FixedAddressUpdateParams) WithDefaults() *FixedAddressUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fixed address update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FixedAddressUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fixed address update params
func (o *FixedAddressUpdateParams) WithTimeout(timeout time.Duration) *FixedAddressUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fixed address update params
func (o *FixedAddressUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fixed address update params
func (o *FixedAddressUpdateParams) WithContext(ctx context.Context) *FixedAddressUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fixed address update params
func (o *FixedAddressUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fixed address update params
func (o *FixedAddressUpdateParams) WithHTTPClient(client *http.Client) *FixedAddressUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fixed address update params
func (o *FixedAddressUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the fixed address update params
func (o *FixedAddressUpdateParams) WithBody(body *models.IpamsvcFixedAddress) *FixedAddressUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the fixed address update params
func (o *FixedAddressUpdateParams) SetBody(body *models.IpamsvcFixedAddress) {
	o.Body = body
}

// WithID adds the id to the fixed address update params
func (o *FixedAddressUpdateParams) WithID(id string) *FixedAddressUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the fixed address update params
func (o *FixedAddressUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FixedAddressUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
