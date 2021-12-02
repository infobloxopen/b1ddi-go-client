// Code generated by go-swagger; DO NOT EDIT.

package convert_domain_name

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
)

// NewConvertDomainNameConvertParams creates a new ConvertDomainNameConvertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConvertDomainNameConvertParams() *ConvertDomainNameConvertParams {
	return &ConvertDomainNameConvertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConvertDomainNameConvertParamsWithTimeout creates a new ConvertDomainNameConvertParams object
// with the ability to set a timeout on a request.
func NewConvertDomainNameConvertParamsWithTimeout(timeout time.Duration) *ConvertDomainNameConvertParams {
	return &ConvertDomainNameConvertParams{
		timeout: timeout,
	}
}

// NewConvertDomainNameConvertParamsWithContext creates a new ConvertDomainNameConvertParams object
// with the ability to set a context for a request.
func NewConvertDomainNameConvertParamsWithContext(ctx context.Context) *ConvertDomainNameConvertParams {
	return &ConvertDomainNameConvertParams{
		Context: ctx,
	}
}

// NewConvertDomainNameConvertParamsWithHTTPClient creates a new ConvertDomainNameConvertParams object
// with the ability to set a custom HTTPClient for a request.
func NewConvertDomainNameConvertParamsWithHTTPClient(client *http.Client) *ConvertDomainNameConvertParams {
	return &ConvertDomainNameConvertParams{
		HTTPClient: client,
	}
}

/* ConvertDomainNameConvertParams contains all the parameters to send to the API endpoint
   for the convert domain name convert operation.

   Typically these are written to a http.Request.
*/
type ConvertDomainNameConvertParams struct {

	/* DomainName.

	   Input domain name in either of IDN or punycode representations.
	*/
	DomainName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the convert domain name convert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConvertDomainNameConvertParams) WithDefaults() *ConvertDomainNameConvertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the convert domain name convert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConvertDomainNameConvertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the convert domain name convert params
func (o *ConvertDomainNameConvertParams) WithTimeout(timeout time.Duration) *ConvertDomainNameConvertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the convert domain name convert params
func (o *ConvertDomainNameConvertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the convert domain name convert params
func (o *ConvertDomainNameConvertParams) WithContext(ctx context.Context) *ConvertDomainNameConvertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the convert domain name convert params
func (o *ConvertDomainNameConvertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the convert domain name convert params
func (o *ConvertDomainNameConvertParams) WithHTTPClient(client *http.Client) *ConvertDomainNameConvertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the convert domain name convert params
func (o *ConvertDomainNameConvertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainName adds the domainName to the convert domain name convert params
func (o *ConvertDomainNameConvertParams) WithDomainName(domainName string) *ConvertDomainNameConvertParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the convert domain name convert params
func (o *ConvertDomainNameConvertParams) SetDomainName(domainName string) {
	o.DomainName = domainName
}

// WriteToRequest writes these params to a swagger request
func (o *ConvertDomainNameConvertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domain_name
	if err := r.SetPathParam("domain_name", o.DomainName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}