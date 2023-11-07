// Code generated by go-swagger; DO NOT EDIT.

package address_block

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"github.com/go-openapi/swag"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewAddressBlockCreateNextAvailableSubnetParams creates a new AddressBlockCreateNextAvailableSubnetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressBlockCreateNextAvailableSubnetParams() *AddressBlockCreateNextAvailableSubnetParams {
	return &AddressBlockCreateNextAvailableSubnetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressBlockCreateNextAvailableSubnetParamsWithTimeout creates a new AddressBlockCreateNextAvailableSubnetParams object
// with the ability to set a timeout on a request.
func NewAddressBlockCreateNextAvailableSubnetParamsWithTimeout(timeout time.Duration) *AddressBlockCreateNextAvailableSubnetParams {
	return &AddressBlockCreateNextAvailableSubnetParams{
		timeout: timeout,
	}
}

// NewAddressBlockCreateNextAvailableSubnetParamsWithContext creates a new AddressBlockCreateNextAvailableSubnetParams object
// with the ability to set a context for a request.
func NewAddressBlockCreateNextAvailableSubnetParamsWithContext(ctx context.Context) *AddressBlockCreateNextAvailableSubnetParams {
	return &AddressBlockCreateNextAvailableSubnetParams{
		Context: ctx,
	}
}

// NewAddressBlockCreateNextAvailableSubnetParamsWithHTTPClient creates a new AddressBlockCreateNextAvailableSubnetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressBlockCreateNextAvailableSubnetParamsWithHTTPClient(client *http.Client) *AddressBlockCreateNextAvailableSubnetParams {
	return &AddressBlockCreateNextAvailableSubnetParams{
		HTTPClient: client,
	}
}

/* AddressBlockCreateNextAvailableSubnetParams contains all the parameters to send to the API endpoint
   for the address block create next available subnet operation.

   Typically these are written to a http.Request.
*/
type AddressBlockCreateNextAvailableSubnetParams struct {

	/* Cidr.
	   The cidr value of subnets to be created.
	   Format: int32
	*/
	Cidr *int32

	/* Comment.
	   Comment of next available subnets.
	*/
	Comment *string

	/* Count.
	   Number of subnets to generate. Default 1 if not set.
	   Format: int32
	*/
	Count *int32

	/* DhcpHost.
	   Reference of OnPrem Host associated with the next available subnets to be created.
	*/
	DhcpHost *string

	/* ID.
	   An application specific resource identity of a resource
	*/
	ID string

	/* Name.
	   Name of next available subnets.
	*/
	Name *string


	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address block create next available subnet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressBlockCreateNextAvailableSubnetParams) WithDefaults() *AddressBlockCreateNextAvailableSubnetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address block create next available subnet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressBlockCreateNextAvailableSubnetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) WithTimeout(timeout time.Duration) *AddressBlockCreateNextAvailableSubnetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) WithContext(ctx context.Context) *AddressBlockCreateNextAvailableSubnetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) WithHTTPClient(client *http.Client) *AddressBlockCreateNextAvailableSubnetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCidr adds the cidr to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) WithCidr(cidr *int32) *AddressBlockCreateNextAvailableSubnetParams {
	o.SetCidr(cidr)
	return o
}

// SetCidr adds the cidr to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) SetCidr(cidr *int32) {
	o.Cidr = cidr
}

// WithComment adds the comment to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) WithComment(comment *string) *AddressBlockCreateNextAvailableSubnetParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithCount adds the count to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) WithCount(count *int32) *AddressBlockCreateNextAvailableSubnetParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) SetCount(count *int32) {
	o.Count = count
}

// WithDhcpHost adds the dhcpHost to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) WithDhcpHost(dhcpHost *string) *AddressBlockCreateNextAvailableSubnetParams {
	o.SetDhcpHost(dhcpHost)
	return o
}

// SetDhcpHost adds the dhcpHost to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) SetDhcpHost(dhcpHost *string) {
	o.DhcpHost = dhcpHost
}

// WithID adds the id to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) WithID(id string) *AddressBlockCreateNextAvailableSubnetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) WithName(name *string) *AddressBlockCreateNextAvailableSubnetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the address block create next available subnet params
func (o *AddressBlockCreateNextAvailableSubnetParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *AddressBlockCreateNextAvailableSubnetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DhcpHost != nil {

		// query param dhcp_host
		var qrDhcpHost string

		if o.DhcpHost != nil {
			qrDhcpHost = *o.DhcpHost
		}
		qDhcpHost := qrDhcpHost
		if qDhcpHost != "" {

			if err := r.SetQueryParam("dhcp_host", qDhcpHost); err != nil {
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
