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

// NewRecordCreateParams creates a new RecordCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRecordCreateParams() *RecordCreateParams {
	return &RecordCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRecordCreateParamsWithTimeout creates a new RecordCreateParams object
// with the ability to set a timeout on a request.
func NewRecordCreateParamsWithTimeout(timeout time.Duration) *RecordCreateParams {
	return &RecordCreateParams{
		timeout: timeout,
	}
}

// NewRecordCreateParamsWithContext creates a new RecordCreateParams object
// with the ability to set a context for a request.
func NewRecordCreateParamsWithContext(ctx context.Context) *RecordCreateParams {
	return &RecordCreateParams{
		Context: ctx,
	}
}

// NewRecordCreateParamsWithHTTPClient creates a new RecordCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewRecordCreateParamsWithHTTPClient(client *http.Client) *RecordCreateParams {
	return &RecordCreateParams{
		HTTPClient: client,
	}
}

/* RecordCreateParams contains all the parameters to send to the API endpoint
   for the record create operation.

   Typically these are written to a http.Request.
*/
type RecordCreateParams struct {

	// Body.
	Body *models.DataRecord

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the record create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecordCreateParams) WithDefaults() *RecordCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the record create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecordCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the record create params
func (o *RecordCreateParams) WithTimeout(timeout time.Duration) *RecordCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the record create params
func (o *RecordCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the record create params
func (o *RecordCreateParams) WithContext(ctx context.Context) *RecordCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the record create params
func (o *RecordCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the record create params
func (o *RecordCreateParams) WithHTTPClient(client *http.Client) *RecordCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the record create params
func (o *RecordCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the record create params
func (o *RecordCreateParams) WithBody(body *models.DataRecord) *RecordCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the record create params
func (o *RecordCreateParams) SetBody(body *models.DataRecord) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RecordCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
