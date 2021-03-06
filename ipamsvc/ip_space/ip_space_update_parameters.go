// Code generated by go-swagger; DO NOT EDIT.

package ip_space

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

// NewIPSpaceUpdateParams creates a new IPSpaceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIPSpaceUpdateParams() *IPSpaceUpdateParams {
	return &IPSpaceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIPSpaceUpdateParamsWithTimeout creates a new IPSpaceUpdateParams object
// with the ability to set a timeout on a request.
func NewIPSpaceUpdateParamsWithTimeout(timeout time.Duration) *IPSpaceUpdateParams {
	return &IPSpaceUpdateParams{
		timeout: timeout,
	}
}

// NewIPSpaceUpdateParamsWithContext creates a new IPSpaceUpdateParams object
// with the ability to set a context for a request.
func NewIPSpaceUpdateParamsWithContext(ctx context.Context) *IPSpaceUpdateParams {
	return &IPSpaceUpdateParams{
		Context: ctx,
	}
}

// NewIPSpaceUpdateParamsWithHTTPClient creates a new IPSpaceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIPSpaceUpdateParamsWithHTTPClient(client *http.Client) *IPSpaceUpdateParams {
	return &IPSpaceUpdateParams{
		HTTPClient: client,
	}
}

/* IPSpaceUpdateParams contains all the parameters to send to the API endpoint
   for the ip space update operation.

   Typically these are written to a http.Request.
*/
type IPSpaceUpdateParams struct {

	// Body.
	Body *models.IpamsvcIPSpace

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ip space update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPSpaceUpdateParams) WithDefaults() *IPSpaceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ip space update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPSpaceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ip space update params
func (o *IPSpaceUpdateParams) WithTimeout(timeout time.Duration) *IPSpaceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ip space update params
func (o *IPSpaceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ip space update params
func (o *IPSpaceUpdateParams) WithContext(ctx context.Context) *IPSpaceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ip space update params
func (o *IPSpaceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ip space update params
func (o *IPSpaceUpdateParams) WithHTTPClient(client *http.Client) *IPSpaceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ip space update params
func (o *IPSpaceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ip space update params
func (o *IPSpaceUpdateParams) WithBody(body *models.IpamsvcIPSpace) *IPSpaceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ip space update params
func (o *IPSpaceUpdateParams) SetBody(body *models.IpamsvcIPSpace) {
	o.Body = body
}

// WithID adds the id to the ip space update params
func (o *IPSpaceUpdateParams) WithID(id string) *IPSpaceUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ip space update params
func (o *IPSpaceUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPSpaceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
