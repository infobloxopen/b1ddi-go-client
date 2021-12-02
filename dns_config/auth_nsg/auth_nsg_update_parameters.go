// Code generated by go-swagger; DO NOT EDIT.

package auth_nsg

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

// NewAuthNsgUpdateParams creates a new AuthNsgUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthNsgUpdateParams() *AuthNsgUpdateParams {
	return &AuthNsgUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthNsgUpdateParamsWithTimeout creates a new AuthNsgUpdateParams object
// with the ability to set a timeout on a request.
func NewAuthNsgUpdateParamsWithTimeout(timeout time.Duration) *AuthNsgUpdateParams {
	return &AuthNsgUpdateParams{
		timeout: timeout,
	}
}

// NewAuthNsgUpdateParamsWithContext creates a new AuthNsgUpdateParams object
// with the ability to set a context for a request.
func NewAuthNsgUpdateParamsWithContext(ctx context.Context) *AuthNsgUpdateParams {
	return &AuthNsgUpdateParams{
		Context: ctx,
	}
}

// NewAuthNsgUpdateParamsWithHTTPClient creates a new AuthNsgUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthNsgUpdateParamsWithHTTPClient(client *http.Client) *AuthNsgUpdateParams {
	return &AuthNsgUpdateParams{
		HTTPClient: client,
	}
}

/* AuthNsgUpdateParams contains all the parameters to send to the API endpoint
   for the auth nsg update operation.

   Typically these are written to a http.Request.
*/
type AuthNsgUpdateParams struct {

	// Body.
	Body *models.ConfigAuthNSG

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auth nsg update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthNsgUpdateParams) WithDefaults() *AuthNsgUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auth nsg update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthNsgUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auth nsg update params
func (o *AuthNsgUpdateParams) WithTimeout(timeout time.Duration) *AuthNsgUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth nsg update params
func (o *AuthNsgUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth nsg update params
func (o *AuthNsgUpdateParams) WithContext(ctx context.Context) *AuthNsgUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth nsg update params
func (o *AuthNsgUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth nsg update params
func (o *AuthNsgUpdateParams) WithHTTPClient(client *http.Client) *AuthNsgUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth nsg update params
func (o *AuthNsgUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the auth nsg update params
func (o *AuthNsgUpdateParams) WithBody(body *models.ConfigAuthNSG) *AuthNsgUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the auth nsg update params
func (o *AuthNsgUpdateParams) SetBody(body *models.ConfigAuthNSG) {
	o.Body = body
}

// WithID adds the id to the auth nsg update params
func (o *AuthNsgUpdateParams) WithID(id string) *AuthNsgUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the auth nsg update params
func (o *AuthNsgUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AuthNsgUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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