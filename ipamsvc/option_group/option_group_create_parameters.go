// Code generated by go-swagger; DO NOT EDIT.

package option_group

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

// NewOptionGroupCreateParams creates a new OptionGroupCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOptionGroupCreateParams() *OptionGroupCreateParams {
	return &OptionGroupCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOptionGroupCreateParamsWithTimeout creates a new OptionGroupCreateParams object
// with the ability to set a timeout on a request.
func NewOptionGroupCreateParamsWithTimeout(timeout time.Duration) *OptionGroupCreateParams {
	return &OptionGroupCreateParams{
		timeout: timeout,
	}
}

// NewOptionGroupCreateParamsWithContext creates a new OptionGroupCreateParams object
// with the ability to set a context for a request.
func NewOptionGroupCreateParamsWithContext(ctx context.Context) *OptionGroupCreateParams {
	return &OptionGroupCreateParams{
		Context: ctx,
	}
}

// NewOptionGroupCreateParamsWithHTTPClient creates a new OptionGroupCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOptionGroupCreateParamsWithHTTPClient(client *http.Client) *OptionGroupCreateParams {
	return &OptionGroupCreateParams{
		HTTPClient: client,
	}
}

/* OptionGroupCreateParams contains all the parameters to send to the API endpoint
   for the option group create operation.

   Typically these are written to a http.Request.
*/
type OptionGroupCreateParams struct {

	// Body.
	Body *models.IpamsvcOptionGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the option group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OptionGroupCreateParams) WithDefaults() *OptionGroupCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the option group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OptionGroupCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the option group create params
func (o *OptionGroupCreateParams) WithTimeout(timeout time.Duration) *OptionGroupCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the option group create params
func (o *OptionGroupCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the option group create params
func (o *OptionGroupCreateParams) WithContext(ctx context.Context) *OptionGroupCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the option group create params
func (o *OptionGroupCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the option group create params
func (o *OptionGroupCreateParams) WithHTTPClient(client *http.Client) *OptionGroupCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the option group create params
func (o *OptionGroupCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the option group create params
func (o *OptionGroupCreateParams) WithBody(body *models.IpamsvcOptionGroup) *OptionGroupCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the option group create params
func (o *OptionGroupCreateParams) SetBody(body *models.IpamsvcOptionGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *OptionGroupCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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