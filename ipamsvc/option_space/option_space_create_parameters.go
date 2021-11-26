// Code generated by go-swagger; DO NOT EDIT.

package option_space

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

// NewOptionSpaceCreateParams creates a new OptionSpaceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOptionSpaceCreateParams() *OptionSpaceCreateParams {
	return &OptionSpaceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOptionSpaceCreateParamsWithTimeout creates a new OptionSpaceCreateParams object
// with the ability to set a timeout on a request.
func NewOptionSpaceCreateParamsWithTimeout(timeout time.Duration) *OptionSpaceCreateParams {
	return &OptionSpaceCreateParams{
		timeout: timeout,
	}
}

// NewOptionSpaceCreateParamsWithContext creates a new OptionSpaceCreateParams object
// with the ability to set a context for a request.
func NewOptionSpaceCreateParamsWithContext(ctx context.Context) *OptionSpaceCreateParams {
	return &OptionSpaceCreateParams{
		Context: ctx,
	}
}

// NewOptionSpaceCreateParamsWithHTTPClient creates a new OptionSpaceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOptionSpaceCreateParamsWithHTTPClient(client *http.Client) *OptionSpaceCreateParams {
	return &OptionSpaceCreateParams{
		HTTPClient: client,
	}
}

/* OptionSpaceCreateParams contains all the parameters to send to the API endpoint
   for the option space create operation.

   Typically these are written to a http.Request.
*/
type OptionSpaceCreateParams struct {

	// Body.
	Body *models.IpamsvcOptionSpace

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the option space create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OptionSpaceCreateParams) WithDefaults() *OptionSpaceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the option space create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OptionSpaceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the option space create params
func (o *OptionSpaceCreateParams) WithTimeout(timeout time.Duration) *OptionSpaceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the option space create params
func (o *OptionSpaceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the option space create params
func (o *OptionSpaceCreateParams) WithContext(ctx context.Context) *OptionSpaceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the option space create params
func (o *OptionSpaceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the option space create params
func (o *OptionSpaceCreateParams) WithHTTPClient(client *http.Client) *OptionSpaceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the option space create params
func (o *OptionSpaceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the option space create params
func (o *OptionSpaceCreateParams) WithBody(body *models.IpamsvcOptionSpace) *OptionSpaceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the option space create params
func (o *OptionSpaceCreateParams) SetBody(body *models.IpamsvcOptionSpace) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *OptionSpaceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
