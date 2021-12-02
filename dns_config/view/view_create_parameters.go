// Code generated by go-swagger; DO NOT EDIT.

package view

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

// NewViewCreateParams creates a new ViewCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewViewCreateParams() *ViewCreateParams {
	return &ViewCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewViewCreateParamsWithTimeout creates a new ViewCreateParams object
// with the ability to set a timeout on a request.
func NewViewCreateParamsWithTimeout(timeout time.Duration) *ViewCreateParams {
	return &ViewCreateParams{
		timeout: timeout,
	}
}

// NewViewCreateParamsWithContext creates a new ViewCreateParams object
// with the ability to set a context for a request.
func NewViewCreateParamsWithContext(ctx context.Context) *ViewCreateParams {
	return &ViewCreateParams{
		Context: ctx,
	}
}

// NewViewCreateParamsWithHTTPClient creates a new ViewCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewViewCreateParamsWithHTTPClient(client *http.Client) *ViewCreateParams {
	return &ViewCreateParams{
		HTTPClient: client,
	}
}

/* ViewCreateParams contains all the parameters to send to the API endpoint
   for the view create operation.

   Typically these are written to a http.Request.
*/
type ViewCreateParams struct {

	// Body.
	Body *models.ConfigView

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the view create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ViewCreateParams) WithDefaults() *ViewCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the view create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ViewCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the view create params
func (o *ViewCreateParams) WithTimeout(timeout time.Duration) *ViewCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the view create params
func (o *ViewCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the view create params
func (o *ViewCreateParams) WithContext(ctx context.Context) *ViewCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the view create params
func (o *ViewCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the view create params
func (o *ViewCreateParams) WithHTTPClient(client *http.Client) *ViewCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the view create params
func (o *ViewCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the view create params
func (o *ViewCreateParams) WithBody(body *models.ConfigView) *ViewCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the view create params
func (o *ViewCreateParams) SetBody(body *models.ConfigView) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ViewCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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