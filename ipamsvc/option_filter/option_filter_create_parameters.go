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

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// NewOptionFilterCreateParams creates a new OptionFilterCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOptionFilterCreateParams() *OptionFilterCreateParams {
	return &OptionFilterCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOptionFilterCreateParamsWithTimeout creates a new OptionFilterCreateParams object
// with the ability to set a timeout on a request.
func NewOptionFilterCreateParamsWithTimeout(timeout time.Duration) *OptionFilterCreateParams {
	return &OptionFilterCreateParams{
		timeout: timeout,
	}
}

// NewOptionFilterCreateParamsWithContext creates a new OptionFilterCreateParams object
// with the ability to set a context for a request.
func NewOptionFilterCreateParamsWithContext(ctx context.Context) *OptionFilterCreateParams {
	return &OptionFilterCreateParams{
		Context: ctx,
	}
}

// NewOptionFilterCreateParamsWithHTTPClient creates a new OptionFilterCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOptionFilterCreateParamsWithHTTPClient(client *http.Client) *OptionFilterCreateParams {
	return &OptionFilterCreateParams{
		HTTPClient: client,
	}
}

/* OptionFilterCreateParams contains all the parameters to send to the API endpoint
   for the option filter create operation.

   Typically these are written to a http.Request.
*/
type OptionFilterCreateParams struct {

	// Body.
	Body *models.IpamsvcOptionFilter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the option filter create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OptionFilterCreateParams) WithDefaults() *OptionFilterCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the option filter create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OptionFilterCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the option filter create params
func (o *OptionFilterCreateParams) WithTimeout(timeout time.Duration) *OptionFilterCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the option filter create params
func (o *OptionFilterCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the option filter create params
func (o *OptionFilterCreateParams) WithContext(ctx context.Context) *OptionFilterCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the option filter create params
func (o *OptionFilterCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the option filter create params
func (o *OptionFilterCreateParams) WithHTTPClient(client *http.Client) *OptionFilterCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the option filter create params
func (o *OptionFilterCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the option filter create params
func (o *OptionFilterCreateParams) WithBody(body *models.IpamsvcOptionFilter) *OptionFilterCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the option filter create params
func (o *OptionFilterCreateParams) SetBody(body *models.IpamsvcOptionFilter) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *OptionFilterCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
