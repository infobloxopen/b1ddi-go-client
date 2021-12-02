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

// NewViewUpdateParams creates a new ViewUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewViewUpdateParams() *ViewUpdateParams {
	return &ViewUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewViewUpdateParamsWithTimeout creates a new ViewUpdateParams object
// with the ability to set a timeout on a request.
func NewViewUpdateParamsWithTimeout(timeout time.Duration) *ViewUpdateParams {
	return &ViewUpdateParams{
		timeout: timeout,
	}
}

// NewViewUpdateParamsWithContext creates a new ViewUpdateParams object
// with the ability to set a context for a request.
func NewViewUpdateParamsWithContext(ctx context.Context) *ViewUpdateParams {
	return &ViewUpdateParams{
		Context: ctx,
	}
}

// NewViewUpdateParamsWithHTTPClient creates a new ViewUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewViewUpdateParamsWithHTTPClient(client *http.Client) *ViewUpdateParams {
	return &ViewUpdateParams{
		HTTPClient: client,
	}
}

/* ViewUpdateParams contains all the parameters to send to the API endpoint
   for the view update operation.

   Typically these are written to a http.Request.
*/
type ViewUpdateParams struct {

	// Body.
	Body *models.ConfigView

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the view update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ViewUpdateParams) WithDefaults() *ViewUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the view update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ViewUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the view update params
func (o *ViewUpdateParams) WithTimeout(timeout time.Duration) *ViewUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the view update params
func (o *ViewUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the view update params
func (o *ViewUpdateParams) WithContext(ctx context.Context) *ViewUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the view update params
func (o *ViewUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the view update params
func (o *ViewUpdateParams) WithHTTPClient(client *http.Client) *ViewUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the view update params
func (o *ViewUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the view update params
func (o *ViewUpdateParams) WithBody(body *models.ConfigView) *ViewUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the view update params
func (o *ViewUpdateParams) SetBody(body *models.ConfigView) {
	o.Body = body
}

// WithID adds the id to the view update params
func (o *ViewUpdateParams) WithID(id string) *ViewUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the view update params
func (o *ViewUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ViewUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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