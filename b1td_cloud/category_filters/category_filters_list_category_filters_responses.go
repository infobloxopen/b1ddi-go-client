// Code generated by go-swagger; DO NOT EDIT.

package category_filters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// CategoryFiltersListCategoryFiltersReader is a Reader for the CategoryFiltersListCategoryFilters structure.
type CategoryFiltersListCategoryFiltersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CategoryFiltersListCategoryFiltersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCategoryFiltersListCategoryFiltersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCategoryFiltersListCategoryFiltersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /category_filters] category_filtersListCategoryFilters", response, response.Code())
	}
}

// NewCategoryFiltersListCategoryFiltersOK creates a CategoryFiltersListCategoryFiltersOK with default headers values
func NewCategoryFiltersListCategoryFiltersOK() *CategoryFiltersListCategoryFiltersOK {
	return &CategoryFiltersListCategoryFiltersOK{}
}

/*
CategoryFiltersListCategoryFiltersOK describes a response with status code 200, with default header values.

GET operation response
*/
type CategoryFiltersListCategoryFiltersOK struct {
	Payload *models.AtcfwCategoryFilterMultiResponse
}

// IsSuccess returns true when this category filters list category filters o k response has a 2xx status code
func (o *CategoryFiltersListCategoryFiltersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this category filters list category filters o k response has a 3xx status code
func (o *CategoryFiltersListCategoryFiltersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this category filters list category filters o k response has a 4xx status code
func (o *CategoryFiltersListCategoryFiltersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this category filters list category filters o k response has a 5xx status code
func (o *CategoryFiltersListCategoryFiltersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this category filters list category filters o k response a status code equal to that given
func (o *CategoryFiltersListCategoryFiltersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the category filters list category filters o k response
func (o *CategoryFiltersListCategoryFiltersOK) Code() int {
	return 200
}

func (o *CategoryFiltersListCategoryFiltersOK) Error() string {
	return fmt.Sprintf("[GET /category_filters][%d] categoryFiltersListCategoryFiltersOK  %+v", 200, o.Payload)
}

func (o *CategoryFiltersListCategoryFiltersOK) String() string {
	return fmt.Sprintf("[GET /category_filters][%d] categoryFiltersListCategoryFiltersOK  %+v", 200, o.Payload)
}

func (o *CategoryFiltersListCategoryFiltersOK) GetPayload() *models.AtcfwCategoryFilterMultiResponse {
	return o.Payload
}

func (o *CategoryFiltersListCategoryFiltersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AtcfwCategoryFilterMultiResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCategoryFiltersListCategoryFiltersInternalServerError creates a CategoryFiltersListCategoryFiltersInternalServerError with default headers values
func NewCategoryFiltersListCategoryFiltersInternalServerError() *CategoryFiltersListCategoryFiltersInternalServerError {
	return &CategoryFiltersListCategoryFiltersInternalServerError{}
}

/*
	CategoryFiltersListCategoryFiltersInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type CategoryFiltersListCategoryFiltersInternalServerError struct {
	Payload *CategoryFiltersListCategoryFiltersInternalServerErrorBody
}

// IsSuccess returns true when this category filters list category filters internal server error response has a 2xx status code
func (o *CategoryFiltersListCategoryFiltersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this category filters list category filters internal server error response has a 3xx status code
func (o *CategoryFiltersListCategoryFiltersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this category filters list category filters internal server error response has a 4xx status code
func (o *CategoryFiltersListCategoryFiltersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this category filters list category filters internal server error response has a 5xx status code
func (o *CategoryFiltersListCategoryFiltersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this category filters list category filters internal server error response a status code equal to that given
func (o *CategoryFiltersListCategoryFiltersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the category filters list category filters internal server error response
func (o *CategoryFiltersListCategoryFiltersInternalServerError) Code() int {
	return 500
}

func (o *CategoryFiltersListCategoryFiltersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /category_filters][%d] categoryFiltersListCategoryFiltersInternalServerError  %+v", 500, o.Payload)
}

func (o *CategoryFiltersListCategoryFiltersInternalServerError) String() string {
	return fmt.Sprintf("[GET /category_filters][%d] categoryFiltersListCategoryFiltersInternalServerError  %+v", 500, o.Payload)
}

func (o *CategoryFiltersListCategoryFiltersInternalServerError) GetPayload() *CategoryFiltersListCategoryFiltersInternalServerErrorBody {
	return o.Payload
}

func (o *CategoryFiltersListCategoryFiltersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CategoryFiltersListCategoryFiltersInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CategoryFiltersListCategoryFiltersInternalServerErrorBody category filters list category filters internal server error body
swagger:model CategoryFiltersListCategoryFiltersInternalServerErrorBody
*/
type CategoryFiltersListCategoryFiltersInternalServerErrorBody struct {

	// error
	Error *CategoryFiltersListCategoryFiltersInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this category filters list category filters internal server error body
func (o *CategoryFiltersListCategoryFiltersInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CategoryFiltersListCategoryFiltersInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("categoryFiltersListCategoryFiltersInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("categoryFiltersListCategoryFiltersInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this category filters list category filters internal server error body based on the context it is used
func (o *CategoryFiltersListCategoryFiltersInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CategoryFiltersListCategoryFiltersInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("categoryFiltersListCategoryFiltersInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("categoryFiltersListCategoryFiltersInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CategoryFiltersListCategoryFiltersInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CategoryFiltersListCategoryFiltersInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res CategoryFiltersListCategoryFiltersInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CategoryFiltersListCategoryFiltersInternalServerErrorBodyError category filters list category filters internal server error body error
swagger:model CategoryFiltersListCategoryFiltersInternalServerErrorBodyError
*/
type CategoryFiltersListCategoryFiltersInternalServerErrorBodyError struct {

	// code
	// Example: INTERNAL
	Code string `json:"code,omitempty"`

	// message
	// Example: Internal Server Error
	Message string `json:"message,omitempty"`

	// status
	// Example: 500
	Status string `json:"status,omitempty"`
}

// Validate validates this category filters list category filters internal server error body error
func (o *CategoryFiltersListCategoryFiltersInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this category filters list category filters internal server error body error based on context it is used
func (o *CategoryFiltersListCategoryFiltersInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CategoryFiltersListCategoryFiltersInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CategoryFiltersListCategoryFiltersInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res CategoryFiltersListCategoryFiltersInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}