// Code generated by go-swagger; DO NOT EDIT.

package access_codes

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

// AccessCodesReadAccessCodeReader is a Reader for the AccessCodesReadAccessCode structure.
type AccessCodesReadAccessCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccessCodesReadAccessCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccessCodesReadAccessCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAccessCodesReadAccessCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccessCodesReadAccessCodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /access_codes/{access_key}] access_codesReadAccessCode", response, response.Code())
	}
}

// NewAccessCodesReadAccessCodeOK creates a AccessCodesReadAccessCodeOK with default headers values
func NewAccessCodesReadAccessCodeOK() *AccessCodesReadAccessCodeOK {
	return &AccessCodesReadAccessCodeOK{}
}

/*
AccessCodesReadAccessCodeOK describes a response with status code 200, with default header values.

GET operation response
*/
type AccessCodesReadAccessCodeOK struct {
	Payload *models.AtcfwAccessCodeReadResponse
}

// IsSuccess returns true when this access codes read access code o k response has a 2xx status code
func (o *AccessCodesReadAccessCodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this access codes read access code o k response has a 3xx status code
func (o *AccessCodesReadAccessCodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access codes read access code o k response has a 4xx status code
func (o *AccessCodesReadAccessCodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this access codes read access code o k response has a 5xx status code
func (o *AccessCodesReadAccessCodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this access codes read access code o k response a status code equal to that given
func (o *AccessCodesReadAccessCodeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the access codes read access code o k response
func (o *AccessCodesReadAccessCodeOK) Code() int {
	return 200
}

func (o *AccessCodesReadAccessCodeOK) Error() string {
	return fmt.Sprintf("[GET /access_codes/{access_key}][%d] accessCodesReadAccessCodeOK  %+v", 200, o.Payload)
}

func (o *AccessCodesReadAccessCodeOK) String() string {
	return fmt.Sprintf("[GET /access_codes/{access_key}][%d] accessCodesReadAccessCodeOK  %+v", 200, o.Payload)
}

func (o *AccessCodesReadAccessCodeOK) GetPayload() *models.AtcfwAccessCodeReadResponse {
	return o.Payload
}

func (o *AccessCodesReadAccessCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AtcfwAccessCodeReadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessCodesReadAccessCodeNotFound creates a AccessCodesReadAccessCodeNotFound with default headers values
func NewAccessCodesReadAccessCodeNotFound() *AccessCodesReadAccessCodeNotFound {
	return &AccessCodesReadAccessCodeNotFound{}
}

/*
	AccessCodesReadAccessCodeNotFound describes a response with status code 404, with default header values.

- 'access_codes' value must contain existing bypass code key
*/
type AccessCodesReadAccessCodeNotFound struct {
	Payload *AccessCodesReadAccessCodeNotFoundBody
}

// IsSuccess returns true when this access codes read access code not found response has a 2xx status code
func (o *AccessCodesReadAccessCodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access codes read access code not found response has a 3xx status code
func (o *AccessCodesReadAccessCodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access codes read access code not found response has a 4xx status code
func (o *AccessCodesReadAccessCodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this access codes read access code not found response has a 5xx status code
func (o *AccessCodesReadAccessCodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this access codes read access code not found response a status code equal to that given
func (o *AccessCodesReadAccessCodeNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the access codes read access code not found response
func (o *AccessCodesReadAccessCodeNotFound) Code() int {
	return 404
}

func (o *AccessCodesReadAccessCodeNotFound) Error() string {
	return fmt.Sprintf("[GET /access_codes/{access_key}][%d] accessCodesReadAccessCodeNotFound  %+v", 404, o.Payload)
}

func (o *AccessCodesReadAccessCodeNotFound) String() string {
	return fmt.Sprintf("[GET /access_codes/{access_key}][%d] accessCodesReadAccessCodeNotFound  %+v", 404, o.Payload)
}

func (o *AccessCodesReadAccessCodeNotFound) GetPayload() *AccessCodesReadAccessCodeNotFoundBody {
	return o.Payload
}

func (o *AccessCodesReadAccessCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AccessCodesReadAccessCodeNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessCodesReadAccessCodeInternalServerError creates a AccessCodesReadAccessCodeInternalServerError with default headers values
func NewAccessCodesReadAccessCodeInternalServerError() *AccessCodesReadAccessCodeInternalServerError {
	return &AccessCodesReadAccessCodeInternalServerError{}
}

/*
	AccessCodesReadAccessCodeInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type AccessCodesReadAccessCodeInternalServerError struct {
	Payload *AccessCodesReadAccessCodeInternalServerErrorBody
}

// IsSuccess returns true when this access codes read access code internal server error response has a 2xx status code
func (o *AccessCodesReadAccessCodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access codes read access code internal server error response has a 3xx status code
func (o *AccessCodesReadAccessCodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access codes read access code internal server error response has a 4xx status code
func (o *AccessCodesReadAccessCodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this access codes read access code internal server error response has a 5xx status code
func (o *AccessCodesReadAccessCodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this access codes read access code internal server error response a status code equal to that given
func (o *AccessCodesReadAccessCodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the access codes read access code internal server error response
func (o *AccessCodesReadAccessCodeInternalServerError) Code() int {
	return 500
}

func (o *AccessCodesReadAccessCodeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /access_codes/{access_key}][%d] accessCodesReadAccessCodeInternalServerError  %+v", 500, o.Payload)
}

func (o *AccessCodesReadAccessCodeInternalServerError) String() string {
	return fmt.Sprintf("[GET /access_codes/{access_key}][%d] accessCodesReadAccessCodeInternalServerError  %+v", 500, o.Payload)
}

func (o *AccessCodesReadAccessCodeInternalServerError) GetPayload() *AccessCodesReadAccessCodeInternalServerErrorBody {
	return o.Payload
}

func (o *AccessCodesReadAccessCodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AccessCodesReadAccessCodeInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AccessCodesReadAccessCodeInternalServerErrorBody access codes read access code internal server error body
swagger:model AccessCodesReadAccessCodeInternalServerErrorBody
*/
type AccessCodesReadAccessCodeInternalServerErrorBody struct {

	// error
	Error *AccessCodesReadAccessCodeInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this access codes read access code internal server error body
func (o *AccessCodesReadAccessCodeInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesReadAccessCodeInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesReadAccessCodeInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesReadAccessCodeInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this access codes read access code internal server error body based on the context it is used
func (o *AccessCodesReadAccessCodeInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesReadAccessCodeInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesReadAccessCodeInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesReadAccessCodeInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesReadAccessCodeInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesReadAccessCodeInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res AccessCodesReadAccessCodeInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessCodesReadAccessCodeInternalServerErrorBodyError access codes read access code internal server error body error
swagger:model AccessCodesReadAccessCodeInternalServerErrorBodyError
*/
type AccessCodesReadAccessCodeInternalServerErrorBodyError struct {

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

// Validate validates this access codes read access code internal server error body error
func (o *AccessCodesReadAccessCodeInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access codes read access code internal server error body error based on context it is used
func (o *AccessCodesReadAccessCodeInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesReadAccessCodeInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesReadAccessCodeInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res AccessCodesReadAccessCodeInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessCodesReadAccessCodeNotFoundBody access codes read access code not found body
swagger:model AccessCodesReadAccessCodeNotFoundBody
*/
type AccessCodesReadAccessCodeNotFoundBody struct {

	// error
	Error *AccessCodesReadAccessCodeNotFoundBodyError `json:"error,omitempty"`
}

// Validate validates this access codes read access code not found body
func (o *AccessCodesReadAccessCodeNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesReadAccessCodeNotFoundBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesReadAccessCodeNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesReadAccessCodeNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this access codes read access code not found body based on the context it is used
func (o *AccessCodesReadAccessCodeNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesReadAccessCodeNotFoundBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesReadAccessCodeNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesReadAccessCodeNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesReadAccessCodeNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesReadAccessCodeNotFoundBody) UnmarshalBinary(b []byte) error {
	var res AccessCodesReadAccessCodeNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessCodesReadAccessCodeNotFoundBodyError access codes read access code not found body error
swagger:model AccessCodesReadAccessCodeNotFoundBodyError
*/
type AccessCodesReadAccessCodeNotFoundBodyError struct {

	// code
	// Example: NOT_FOUND
	Code string `json:"code,omitempty"`

	// message
	// Example: Bypass Code does not exist
	Message string `json:"message,omitempty"`

	// status
	// Example: 404
	Status string `json:"status,omitempty"`
}

// Validate validates this access codes read access code not found body error
func (o *AccessCodesReadAccessCodeNotFoundBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access codes read access code not found body error based on context it is used
func (o *AccessCodesReadAccessCodeNotFoundBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesReadAccessCodeNotFoundBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesReadAccessCodeNotFoundBodyError) UnmarshalBinary(b []byte) error {
	var res AccessCodesReadAccessCodeNotFoundBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
