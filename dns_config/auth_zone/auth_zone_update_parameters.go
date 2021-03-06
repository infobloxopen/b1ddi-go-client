// Code generated by go-swagger; DO NOT EDIT.

package auth_zone

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

// NewAuthZoneUpdateParams creates a new AuthZoneUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthZoneUpdateParams() *AuthZoneUpdateParams {
	return &AuthZoneUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthZoneUpdateParamsWithTimeout creates a new AuthZoneUpdateParams object
// with the ability to set a timeout on a request.
func NewAuthZoneUpdateParamsWithTimeout(timeout time.Duration) *AuthZoneUpdateParams {
	return &AuthZoneUpdateParams{
		timeout: timeout,
	}
}

// NewAuthZoneUpdateParamsWithContext creates a new AuthZoneUpdateParams object
// with the ability to set a context for a request.
func NewAuthZoneUpdateParamsWithContext(ctx context.Context) *AuthZoneUpdateParams {
	return &AuthZoneUpdateParams{
		Context: ctx,
	}
}

// NewAuthZoneUpdateParamsWithHTTPClient creates a new AuthZoneUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthZoneUpdateParamsWithHTTPClient(client *http.Client) *AuthZoneUpdateParams {
	return &AuthZoneUpdateParams{
		HTTPClient: client,
	}
}

/* AuthZoneUpdateParams contains all the parameters to send to the API endpoint
   for the auth zone update operation.

   Typically these are written to a http.Request.
*/
type AuthZoneUpdateParams struct {

	// Body.
	Body *models.ConfigAuthZone

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auth zone update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthZoneUpdateParams) WithDefaults() *AuthZoneUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auth zone update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthZoneUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auth zone update params
func (o *AuthZoneUpdateParams) WithTimeout(timeout time.Duration) *AuthZoneUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth zone update params
func (o *AuthZoneUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth zone update params
func (o *AuthZoneUpdateParams) WithContext(ctx context.Context) *AuthZoneUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth zone update params
func (o *AuthZoneUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth zone update params
func (o *AuthZoneUpdateParams) WithHTTPClient(client *http.Client) *AuthZoneUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth zone update params
func (o *AuthZoneUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the auth zone update params
func (o *AuthZoneUpdateParams) WithBody(body *models.ConfigAuthZone) *AuthZoneUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the auth zone update params
func (o *AuthZoneUpdateParams) SetBody(body *models.ConfigAuthZone) {
	o.Body = body
}

// WithID adds the id to the auth zone update params
func (o *AuthZoneUpdateParams) WithID(id string) *AuthZoneUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the auth zone update params
func (o *AuthZoneUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AuthZoneUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
