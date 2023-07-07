// Code generated by go-swagger; DO NOT EDIT.

package category_filters

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

// NewCategoryFiltersCreateCategoryFilterParams creates a new CategoryFiltersCreateCategoryFilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCategoryFiltersCreateCategoryFilterParams() *CategoryFiltersCreateCategoryFilterParams {
	return &CategoryFiltersCreateCategoryFilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCategoryFiltersCreateCategoryFilterParamsWithTimeout creates a new CategoryFiltersCreateCategoryFilterParams object
// with the ability to set a timeout on a request.
func NewCategoryFiltersCreateCategoryFilterParamsWithTimeout(timeout time.Duration) *CategoryFiltersCreateCategoryFilterParams {
	return &CategoryFiltersCreateCategoryFilterParams{
		timeout: timeout,
	}
}

// NewCategoryFiltersCreateCategoryFilterParamsWithContext creates a new CategoryFiltersCreateCategoryFilterParams object
// with the ability to set a context for a request.
func NewCategoryFiltersCreateCategoryFilterParamsWithContext(ctx context.Context) *CategoryFiltersCreateCategoryFilterParams {
	return &CategoryFiltersCreateCategoryFilterParams{
		Context: ctx,
	}
}

// NewCategoryFiltersCreateCategoryFilterParamsWithHTTPClient creates a new CategoryFiltersCreateCategoryFilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCategoryFiltersCreateCategoryFilterParamsWithHTTPClient(client *http.Client) *CategoryFiltersCreateCategoryFilterParams {
	return &CategoryFiltersCreateCategoryFilterParams{
		HTTPClient: client,
	}
}

/*
CategoryFiltersCreateCategoryFilterParams contains all the parameters to send to the API endpoint

	for the category filters create category filter operation.

	Typically these are written to a http.Request.
*/
type CategoryFiltersCreateCategoryFilterParams struct {

	/* Body.

	   The Category Filter object.
	*/
	Body *models.AtcfwCategoryFilter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the category filters create category filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CategoryFiltersCreateCategoryFilterParams) WithDefaults() *CategoryFiltersCreateCategoryFilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the category filters create category filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CategoryFiltersCreateCategoryFilterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the category filters create category filter params
func (o *CategoryFiltersCreateCategoryFilterParams) WithTimeout(timeout time.Duration) *CategoryFiltersCreateCategoryFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the category filters create category filter params
func (o *CategoryFiltersCreateCategoryFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the category filters create category filter params
func (o *CategoryFiltersCreateCategoryFilterParams) WithContext(ctx context.Context) *CategoryFiltersCreateCategoryFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the category filters create category filter params
func (o *CategoryFiltersCreateCategoryFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the category filters create category filter params
func (o *CategoryFiltersCreateCategoryFilterParams) WithHTTPClient(client *http.Client) *CategoryFiltersCreateCategoryFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the category filters create category filter params
func (o *CategoryFiltersCreateCategoryFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the category filters create category filter params
func (o *CategoryFiltersCreateCategoryFilterParams) WithBody(body *models.AtcfwCategoryFilter) *CategoryFiltersCreateCategoryFilterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the category filters create category filter params
func (o *CategoryFiltersCreateCategoryFilterParams) SetBody(body *models.AtcfwCategoryFilter) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CategoryFiltersCreateCategoryFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
