// Code generated by go-swagger; DO NOT EDIT.

package address_block

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
	"github.com/go-openapi/swag"
)

// NewAddressBlockListNextAvailableABParams creates a new AddressBlockListNextAvailableABParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressBlockListNextAvailableABParams() *AddressBlockListNextAvailableABParams {
	return &AddressBlockListNextAvailableABParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressBlockListNextAvailableABParamsWithTimeout creates a new AddressBlockListNextAvailableABParams object
// with the ability to set a timeout on a request.
func NewAddressBlockListNextAvailableABParamsWithTimeout(timeout time.Duration) *AddressBlockListNextAvailableABParams {
	return &AddressBlockListNextAvailableABParams{
		timeout: timeout,
	}
}

// NewAddressBlockListNextAvailableABParamsWithContext creates a new AddressBlockListNextAvailableABParams object
// with the ability to set a context for a request.
func NewAddressBlockListNextAvailableABParamsWithContext(ctx context.Context) *AddressBlockListNextAvailableABParams {
	return &AddressBlockListNextAvailableABParams{
		Context: ctx,
	}
}

// NewAddressBlockListNextAvailableABParamsWithHTTPClient creates a new AddressBlockListNextAvailableABParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressBlockListNextAvailableABParamsWithHTTPClient(client *http.Client) *AddressBlockListNextAvailableABParams {
	return &AddressBlockListNextAvailableABParams{
		HTTPClient: client,
	}
}

/* AddressBlockListNextAvailableABParams contains all the parameters to send to the API endpoint
   for the address block list next available a b operation.

   Typically these are written to a http.Request.
*/
type AddressBlockListNextAvailableABParams struct {

	/* Cidr.

	   The cidr value of address blocks to be created.

	   Format: int32
	*/
	Cidr *int32

	/* Comment.

	   Comment of next available address blocks.
	*/
	Comment *string

	/* Count.

	   Number of address blocks to generate. Default 1 if not set.

	   Format: int32
	*/
	Count *int32

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	/* Name.

	   Name of next available address blocks.
	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address block list next available a b params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressBlockListNextAvailableABParams) WithDefaults() *AddressBlockListNextAvailableABParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address block list next available a b params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressBlockListNextAvailableABParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) WithTimeout(timeout time.Duration) *AddressBlockListNextAvailableABParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) WithContext(ctx context.Context) *AddressBlockListNextAvailableABParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) WithHTTPClient(client *http.Client) *AddressBlockListNextAvailableABParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCidr adds the cidr to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) WithCidr(cidr *int32) *AddressBlockListNextAvailableABParams {
	o.SetCidr(cidr)
	return o
}

// SetCidr adds the cidr to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) SetCidr(cidr *int32) {
	o.Cidr = cidr
}

// WithComment adds the comment to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) WithComment(comment *string) *AddressBlockListNextAvailableABParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithCount adds the count to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) WithCount(count *int32) *AddressBlockListNextAvailableABParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) SetCount(count *int32) {
	o.Count = count
}

// WithID adds the id to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) WithID(id string) *AddressBlockListNextAvailableABParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) WithName(name *string) *AddressBlockListNextAvailableABParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the address block list next available a b params
func (o *AddressBlockListNextAvailableABParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *AddressBlockListNextAvailableABParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cidr != nil {

		// query param cidr
		var qrCidr int32

		if o.Cidr != nil {
			qrCidr = *o.Cidr
		}
		qCidr := swag.FormatInt32(qrCidr)
		if qCidr != "" {

			if err := r.SetQueryParam("cidr", qCidr); err != nil {
				return err
			}
		}
	}

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.Count != nil {

		// query param count
		var qrCount int32

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
