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

// NewOptionFilterUpdateParams creates a new OptionFilterUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOptionFilterUpdateParams() *OptionFilterUpdateParams {
	return &OptionFilterUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOptionFilterUpdateParamsWithTimeout creates a new OptionFilterUpdateParams object
// with the ability to set a timeout on a request.
func NewOptionFilterUpdateParamsWithTimeout(timeout time.Duration) *OptionFilterUpdateParams {
	return &OptionFilterUpdateParams{
		timeout: timeout,
	}
}

// NewOptionFilterUpdateParamsWithContext creates a new OptionFilterUpdateParams object
// with the ability to set a context for a request.
func NewOptionFilterUpdateParamsWithContext(ctx context.Context) *OptionFilterUpdateParams {
	return &OptionFilterUpdateParams{
		Context: ctx,
	}
}

// NewOptionFilterUpdateParamsWithHTTPClient creates a new OptionFilterUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOptionFilterUpdateParamsWithHTTPClient(client *http.Client) *OptionFilterUpdateParams {
	return &OptionFilterUpdateParams{
		HTTPClient: client,
	}
}

/* OptionFilterUpdateParams contains all the parameters to send to the API endpoint
   for the option filter update operation.

   Typically these are written to a http.Request.
*/
type OptionFilterUpdateParams struct {

	// Body.
	Body *models.IpamsvcOptionFilter

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the option filter update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OptionFilterUpdateParams) WithDefaults() *OptionFilterUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the option filter update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OptionFilterUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the option filter update params
func (o *OptionFilterUpdateParams) WithTimeout(timeout time.Duration) *OptionFilterUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the option filter update params
func (o *OptionFilterUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the option filter update params
func (o *OptionFilterUpdateParams) WithContext(ctx context.Context) *OptionFilterUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the option filter update params
func (o *OptionFilterUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the option filter update params
func (o *OptionFilterUpdateParams) WithHTTPClient(client *http.Client) *OptionFilterUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the option filter update params
func (o *OptionFilterUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the option filter update params
func (o *OptionFilterUpdateParams) WithBody(body *models.IpamsvcOptionFilter) *OptionFilterUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the option filter update params
func (o *OptionFilterUpdateParams) SetBody(body *models.IpamsvcOptionFilter) {
	o.Body = body
}

// WithID adds the id to the option filter update params
func (o *OptionFilterUpdateParams) WithID(id string) *OptionFilterUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the option filter update params
func (o *OptionFilterUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OptionFilterUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
