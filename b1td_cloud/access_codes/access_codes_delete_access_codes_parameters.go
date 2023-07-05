// Code generated by go-swagger; DO NOT EDIT.

package access_codes

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

// NewAccessCodesDeleteAccessCodesParams creates a new AccessCodesDeleteAccessCodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccessCodesDeleteAccessCodesParams() *AccessCodesDeleteAccessCodesParams {
	return &AccessCodesDeleteAccessCodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccessCodesDeleteAccessCodesParamsWithTimeout creates a new AccessCodesDeleteAccessCodesParams object
// with the ability to set a timeout on a request.
func NewAccessCodesDeleteAccessCodesParamsWithTimeout(timeout time.Duration) *AccessCodesDeleteAccessCodesParams {
	return &AccessCodesDeleteAccessCodesParams{
		timeout: timeout,
	}
}

// NewAccessCodesDeleteAccessCodesParamsWithContext creates a new AccessCodesDeleteAccessCodesParams object
// with the ability to set a context for a request.
func NewAccessCodesDeleteAccessCodesParamsWithContext(ctx context.Context) *AccessCodesDeleteAccessCodesParams {
	return &AccessCodesDeleteAccessCodesParams{
		Context: ctx,
	}
}

// NewAccessCodesDeleteAccessCodesParamsWithHTTPClient creates a new AccessCodesDeleteAccessCodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccessCodesDeleteAccessCodesParamsWithHTTPClient(client *http.Client) *AccessCodesDeleteAccessCodesParams {
	return &AccessCodesDeleteAccessCodesParams{
		HTTPClient: client,
	}
}

/*
AccessCodesDeleteAccessCodesParams contains all the parameters to send to the API endpoint

	for the access codes delete access codes operation.

	Typically these are written to a http.Request.
*/
type AccessCodesDeleteAccessCodesParams struct {

	// Body.
	Body *models.AtcfwAccessCodeDeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the access codes delete access codes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessCodesDeleteAccessCodesParams) WithDefaults() *AccessCodesDeleteAccessCodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the access codes delete access codes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessCodesDeleteAccessCodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the access codes delete access codes params
func (o *AccessCodesDeleteAccessCodesParams) WithTimeout(timeout time.Duration) *AccessCodesDeleteAccessCodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the access codes delete access codes params
func (o *AccessCodesDeleteAccessCodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the access codes delete access codes params
func (o *AccessCodesDeleteAccessCodesParams) WithContext(ctx context.Context) *AccessCodesDeleteAccessCodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the access codes delete access codes params
func (o *AccessCodesDeleteAccessCodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the access codes delete access codes params
func (o *AccessCodesDeleteAccessCodesParams) WithHTTPClient(client *http.Client) *AccessCodesDeleteAccessCodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the access codes delete access codes params
func (o *AccessCodesDeleteAccessCodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the access codes delete access codes params
func (o *AccessCodesDeleteAccessCodesParams) WithBody(body *models.AtcfwAccessCodeDeleteRequest) *AccessCodesDeleteAccessCodesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the access codes delete access codes params
func (o *AccessCodesDeleteAccessCodesParams) SetBody(body *models.AtcfwAccessCodeDeleteRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AccessCodesDeleteAccessCodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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