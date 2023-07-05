// Code generated by go-swagger; DO NOT EDIT.

package named_lists

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

// NewNamedListsReadNamedListParams creates a new NamedListsReadNamedListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNamedListsReadNamedListParams() *NamedListsReadNamedListParams {
	return &NamedListsReadNamedListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNamedListsReadNamedListParamsWithTimeout creates a new NamedListsReadNamedListParams object
// with the ability to set a timeout on a request.
func NewNamedListsReadNamedListParamsWithTimeout(timeout time.Duration) *NamedListsReadNamedListParams {
	return &NamedListsReadNamedListParams{
		timeout: timeout,
	}
}

// NewNamedListsReadNamedListParamsWithContext creates a new NamedListsReadNamedListParams object
// with the ability to set a context for a request.
func NewNamedListsReadNamedListParamsWithContext(ctx context.Context) *NamedListsReadNamedListParams {
	return &NamedListsReadNamedListParams{
		Context: ctx,
	}
}

// NewNamedListsReadNamedListParamsWithHTTPClient creates a new NamedListsReadNamedListParams object
// with the ability to set a custom HTTPClient for a request.
func NewNamedListsReadNamedListParamsWithHTTPClient(client *http.Client) *NamedListsReadNamedListParams {
	return &NamedListsReadNamedListParams{
		HTTPClient: client,
	}
}

/*
NamedListsReadNamedListParams contains all the parameters to send to the API endpoint

	for the named lists read named list operation.

	Typically these are written to a http.Request.
*/
type NamedListsReadNamedListParams struct {

	/* Fields.



	A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.

	Specify this parameter as a comma-separated list of JSON tag names.


	*/
	Fields *string

	/* Limit.



	The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.


	*/
	Limit *int64

	/* Offset.



	The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.


	*/
	Offset *int64

	/* PageToken.



	The service-defined string used to identify a page of resources. A null value indicates the first page.


	*/
	PageToken *string

	/* ID.

	   The Named List identifier.

	   Format: int32
	*/
	ID int32

	/* Name.

	     The name of the named list.
	Can be used in pair with 'type' (both fields are mandatory) to request
	the object by their name. This aproach available only if the field 'id'
	is empty (==0).
	*/
	Name *string

	/* Type.

	   The type of the named list. See 'NamedList' for more details.
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the named lists read named list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamedListsReadNamedListParams) WithDefaults() *NamedListsReadNamedListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the named lists read named list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamedListsReadNamedListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the named lists read named list params
func (o *NamedListsReadNamedListParams) WithTimeout(timeout time.Duration) *NamedListsReadNamedListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the named lists read named list params
func (o *NamedListsReadNamedListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the named lists read named list params
func (o *NamedListsReadNamedListParams) WithContext(ctx context.Context) *NamedListsReadNamedListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the named lists read named list params
func (o *NamedListsReadNamedListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the named lists read named list params
func (o *NamedListsReadNamedListParams) WithHTTPClient(client *http.Client) *NamedListsReadNamedListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the named lists read named list params
func (o *NamedListsReadNamedListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the named lists read named list params
func (o *NamedListsReadNamedListParams) WithFields(fields *string) *NamedListsReadNamedListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the named lists read named list params
func (o *NamedListsReadNamedListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the named lists read named list params
func (o *NamedListsReadNamedListParams) WithLimit(limit *int64) *NamedListsReadNamedListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the named lists read named list params
func (o *NamedListsReadNamedListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the named lists read named list params
func (o *NamedListsReadNamedListParams) WithOffset(offset *int64) *NamedListsReadNamedListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the named lists read named list params
func (o *NamedListsReadNamedListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPageToken adds the pageToken to the named lists read named list params
func (o *NamedListsReadNamedListParams) WithPageToken(pageToken *string) *NamedListsReadNamedListParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the named lists read named list params
func (o *NamedListsReadNamedListParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithID adds the id to the named lists read named list params
func (o *NamedListsReadNamedListParams) WithID(id int32) *NamedListsReadNamedListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the named lists read named list params
func (o *NamedListsReadNamedListParams) SetID(id int32) {
	o.ID = id
}

// WithName adds the name to the named lists read named list params
func (o *NamedListsReadNamedListParams) WithName(name *string) *NamedListsReadNamedListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the named lists read named list params
func (o *NamedListsReadNamedListParams) SetName(name *string) {
	o.Name = name
}

// WithType adds the typeVar to the named lists read named list params
func (o *NamedListsReadNamedListParams) WithType(typeVar *string) *NamedListsReadNamedListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the named lists read named list params
func (o *NamedListsReadNamedListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *NamedListsReadNamedListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param _fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("_fields", qFields); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param _limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("_limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param _offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("_offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.PageToken != nil {

		// query param _page_token
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("_page_token", qPageToken); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
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

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}