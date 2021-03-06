// Code generated by go-swagger; DO NOT EDIT.

package record

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

// NewRecordSOASerialIncrementParams creates a new RecordSOASerialIncrementParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRecordSOASerialIncrementParams() *RecordSOASerialIncrementParams {
	return &RecordSOASerialIncrementParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRecordSOASerialIncrementParamsWithTimeout creates a new RecordSOASerialIncrementParams object
// with the ability to set a timeout on a request.
func NewRecordSOASerialIncrementParamsWithTimeout(timeout time.Duration) *RecordSOASerialIncrementParams {
	return &RecordSOASerialIncrementParams{
		timeout: timeout,
	}
}

// NewRecordSOASerialIncrementParamsWithContext creates a new RecordSOASerialIncrementParams object
// with the ability to set a context for a request.
func NewRecordSOASerialIncrementParamsWithContext(ctx context.Context) *RecordSOASerialIncrementParams {
	return &RecordSOASerialIncrementParams{
		Context: ctx,
	}
}

// NewRecordSOASerialIncrementParamsWithHTTPClient creates a new RecordSOASerialIncrementParams object
// with the ability to set a custom HTTPClient for a request.
func NewRecordSOASerialIncrementParamsWithHTTPClient(client *http.Client) *RecordSOASerialIncrementParams {
	return &RecordSOASerialIncrementParams{
		HTTPClient: client,
	}
}

/* RecordSOASerialIncrementParams contains all the parameters to send to the API endpoint
   for the record s o a serial increment operation.

   Typically these are written to a http.Request.
*/
type RecordSOASerialIncrementParams struct {

	// Body.
	Body *models.DataSOASerialIncrementRequest

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the record s o a serial increment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecordSOASerialIncrementParams) WithDefaults() *RecordSOASerialIncrementParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the record s o a serial increment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecordSOASerialIncrementParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the record s o a serial increment params
func (o *RecordSOASerialIncrementParams) WithTimeout(timeout time.Duration) *RecordSOASerialIncrementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the record s o a serial increment params
func (o *RecordSOASerialIncrementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the record s o a serial increment params
func (o *RecordSOASerialIncrementParams) WithContext(ctx context.Context) *RecordSOASerialIncrementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the record s o a serial increment params
func (o *RecordSOASerialIncrementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the record s o a serial increment params
func (o *RecordSOASerialIncrementParams) WithHTTPClient(client *http.Client) *RecordSOASerialIncrementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the record s o a serial increment params
func (o *RecordSOASerialIncrementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the record s o a serial increment params
func (o *RecordSOASerialIncrementParams) WithBody(body *models.DataSOASerialIncrementRequest) *RecordSOASerialIncrementParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the record s o a serial increment params
func (o *RecordSOASerialIncrementParams) SetBody(body *models.DataSOASerialIncrementRequest) {
	o.Body = body
}

// WithID adds the id to the record s o a serial increment params
func (o *RecordSOASerialIncrementParams) WithID(id string) *RecordSOASerialIncrementParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the record s o a serial increment params
func (o *RecordSOASerialIncrementParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RecordSOASerialIncrementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
