// Code generated by go-swagger; DO NOT EDIT.

package forward_zone

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

// NewForwardZoneUpdateParams creates a new ForwardZoneUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewForwardZoneUpdateParams() *ForwardZoneUpdateParams {
	return &ForwardZoneUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewForwardZoneUpdateParamsWithTimeout creates a new ForwardZoneUpdateParams object
// with the ability to set a timeout on a request.
func NewForwardZoneUpdateParamsWithTimeout(timeout time.Duration) *ForwardZoneUpdateParams {
	return &ForwardZoneUpdateParams{
		timeout: timeout,
	}
}

// NewForwardZoneUpdateParamsWithContext creates a new ForwardZoneUpdateParams object
// with the ability to set a context for a request.
func NewForwardZoneUpdateParamsWithContext(ctx context.Context) *ForwardZoneUpdateParams {
	return &ForwardZoneUpdateParams{
		Context: ctx,
	}
}

// NewForwardZoneUpdateParamsWithHTTPClient creates a new ForwardZoneUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewForwardZoneUpdateParamsWithHTTPClient(client *http.Client) *ForwardZoneUpdateParams {
	return &ForwardZoneUpdateParams{
		HTTPClient: client,
	}
}

/* ForwardZoneUpdateParams contains all the parameters to send to the API endpoint
   for the forward zone update operation.

   Typically these are written to a http.Request.
*/
type ForwardZoneUpdateParams struct {

	// Body.
	Body *models.ConfigForwardZone

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the forward zone update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForwardZoneUpdateParams) WithDefaults() *ForwardZoneUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the forward zone update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForwardZoneUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the forward zone update params
func (o *ForwardZoneUpdateParams) WithTimeout(timeout time.Duration) *ForwardZoneUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the forward zone update params
func (o *ForwardZoneUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the forward zone update params
func (o *ForwardZoneUpdateParams) WithContext(ctx context.Context) *ForwardZoneUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the forward zone update params
func (o *ForwardZoneUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the forward zone update params
func (o *ForwardZoneUpdateParams) WithHTTPClient(client *http.Client) *ForwardZoneUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the forward zone update params
func (o *ForwardZoneUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the forward zone update params
func (o *ForwardZoneUpdateParams) WithBody(body *models.ConfigForwardZone) *ForwardZoneUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the forward zone update params
func (o *ForwardZoneUpdateParams) SetBody(body *models.ConfigForwardZone) {
	o.Body = body
}

// WithID adds the id to the forward zone update params
func (o *ForwardZoneUpdateParams) WithID(id string) *ForwardZoneUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the forward zone update params
func (o *ForwardZoneUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ForwardZoneUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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