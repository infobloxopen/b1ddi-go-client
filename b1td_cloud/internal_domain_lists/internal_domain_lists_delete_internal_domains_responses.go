// Code generated by go-swagger; DO NOT EDIT.

package internal_domain_lists

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
)

// InternalDomainListsDeleteInternalDomainsReader is a Reader for the InternalDomainListsDeleteInternalDomains structure.
type InternalDomainListsDeleteInternalDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalDomainListsDeleteInternalDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewInternalDomainListsDeleteInternalDomainsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalDomainListsDeleteInternalDomainsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInternalDomainListsDeleteInternalDomainsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalDomainListsDeleteInternalDomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /internal_domain_lists] internal_domain_listsDeleteInternalDomains", response, response.Code())
	}
}

// NewInternalDomainListsDeleteInternalDomainsNoContent creates a InternalDomainListsDeleteInternalDomainsNoContent with default headers values
func NewInternalDomainListsDeleteInternalDomainsNoContent() *InternalDomainListsDeleteInternalDomainsNoContent {
	return &InternalDomainListsDeleteInternalDomainsNoContent{}
}

/*
InternalDomainListsDeleteInternalDomainsNoContent describes a response with status code 204, with default header values.

No Content
*/
type InternalDomainListsDeleteInternalDomainsNoContent struct {
}

// IsSuccess returns true when this internal domain lists delete internal domains no content response has a 2xx status code
func (o *InternalDomainListsDeleteInternalDomainsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal domain lists delete internal domains no content response has a 3xx status code
func (o *InternalDomainListsDeleteInternalDomainsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists delete internal domains no content response has a 4xx status code
func (o *InternalDomainListsDeleteInternalDomainsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal domain lists delete internal domains no content response has a 5xx status code
func (o *InternalDomainListsDeleteInternalDomainsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this internal domain lists delete internal domains no content response a status code equal to that given
func (o *InternalDomainListsDeleteInternalDomainsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the internal domain lists delete internal domains no content response
func (o *InternalDomainListsDeleteInternalDomainsNoContent) Code() int {
	return 204
}

func (o *InternalDomainListsDeleteInternalDomainsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /internal_domain_lists][%d] internalDomainListsDeleteInternalDomainsNoContent ", 204)
}

func (o *InternalDomainListsDeleteInternalDomainsNoContent) String() string {
	return fmt.Sprintf("[DELETE /internal_domain_lists][%d] internalDomainListsDeleteInternalDomainsNoContent ", 204)
}

func (o *InternalDomainListsDeleteInternalDomainsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInternalDomainListsDeleteInternalDomainsBadRequest creates a InternalDomainListsDeleteInternalDomainsBadRequest with default headers values
func NewInternalDomainListsDeleteInternalDomainsBadRequest() *InternalDomainListsDeleteInternalDomainsBadRequest {
	return &InternalDomainListsDeleteInternalDomainsBadRequest{}
}

/*
	InternalDomainListsDeleteInternalDomainsBadRequest describes a response with status code 400, with default header values.

- 'ids' value must be non-empty
- 'ids' value must contain unique elements
- 'ids' value must contain values that are greater than or equal to zero
- internal domain list assigned to a endpoint group cannot be deleted
*/
type InternalDomainListsDeleteInternalDomainsBadRequest struct {
	Payload *InternalDomainListsDeleteInternalDomainsBadRequestBody
}

// IsSuccess returns true when this internal domain lists delete internal domains bad request response has a 2xx status code
func (o *InternalDomainListsDeleteInternalDomainsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal domain lists delete internal domains bad request response has a 3xx status code
func (o *InternalDomainListsDeleteInternalDomainsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists delete internal domains bad request response has a 4xx status code
func (o *InternalDomainListsDeleteInternalDomainsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal domain lists delete internal domains bad request response has a 5xx status code
func (o *InternalDomainListsDeleteInternalDomainsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this internal domain lists delete internal domains bad request response a status code equal to that given
func (o *InternalDomainListsDeleteInternalDomainsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the internal domain lists delete internal domains bad request response
func (o *InternalDomainListsDeleteInternalDomainsBadRequest) Code() int {
	return 400
}

func (o *InternalDomainListsDeleteInternalDomainsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /internal_domain_lists][%d] internalDomainListsDeleteInternalDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *InternalDomainListsDeleteInternalDomainsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /internal_domain_lists][%d] internalDomainListsDeleteInternalDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *InternalDomainListsDeleteInternalDomainsBadRequest) GetPayload() *InternalDomainListsDeleteInternalDomainsBadRequestBody {
	return o.Payload
}

func (o *InternalDomainListsDeleteInternalDomainsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InternalDomainListsDeleteInternalDomainsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalDomainListsDeleteInternalDomainsNotFound creates a InternalDomainListsDeleteInternalDomainsNotFound with default headers values
func NewInternalDomainListsDeleteInternalDomainsNotFound() *InternalDomainListsDeleteInternalDomainsNotFound {
	return &InternalDomainListsDeleteInternalDomainsNotFound{}
}

/*
	InternalDomainListsDeleteInternalDomainsNotFound describes a response with status code 404, with default header values.

- 'ids' value must contain existing internal domains identifiers
*/
type InternalDomainListsDeleteInternalDomainsNotFound struct {
	Payload *InternalDomainListsDeleteInternalDomainsNotFoundBody
}

// IsSuccess returns true when this internal domain lists delete internal domains not found response has a 2xx status code
func (o *InternalDomainListsDeleteInternalDomainsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal domain lists delete internal domains not found response has a 3xx status code
func (o *InternalDomainListsDeleteInternalDomainsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists delete internal domains not found response has a 4xx status code
func (o *InternalDomainListsDeleteInternalDomainsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal domain lists delete internal domains not found response has a 5xx status code
func (o *InternalDomainListsDeleteInternalDomainsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this internal domain lists delete internal domains not found response a status code equal to that given
func (o *InternalDomainListsDeleteInternalDomainsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the internal domain lists delete internal domains not found response
func (o *InternalDomainListsDeleteInternalDomainsNotFound) Code() int {
	return 404
}

func (o *InternalDomainListsDeleteInternalDomainsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /internal_domain_lists][%d] internalDomainListsDeleteInternalDomainsNotFound  %+v", 404, o.Payload)
}

func (o *InternalDomainListsDeleteInternalDomainsNotFound) String() string {
	return fmt.Sprintf("[DELETE /internal_domain_lists][%d] internalDomainListsDeleteInternalDomainsNotFound  %+v", 404, o.Payload)
}

func (o *InternalDomainListsDeleteInternalDomainsNotFound) GetPayload() *InternalDomainListsDeleteInternalDomainsNotFoundBody {
	return o.Payload
}

func (o *InternalDomainListsDeleteInternalDomainsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InternalDomainListsDeleteInternalDomainsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalDomainListsDeleteInternalDomainsInternalServerError creates a InternalDomainListsDeleteInternalDomainsInternalServerError with default headers values
func NewInternalDomainListsDeleteInternalDomainsInternalServerError() *InternalDomainListsDeleteInternalDomainsInternalServerError {
	return &InternalDomainListsDeleteInternalDomainsInternalServerError{}
}

/*
	InternalDomainListsDeleteInternalDomainsInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type InternalDomainListsDeleteInternalDomainsInternalServerError struct {
	Payload *InternalDomainListsDeleteInternalDomainsInternalServerErrorBody
}

// IsSuccess returns true when this internal domain lists delete internal domains internal server error response has a 2xx status code
func (o *InternalDomainListsDeleteInternalDomainsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal domain lists delete internal domains internal server error response has a 3xx status code
func (o *InternalDomainListsDeleteInternalDomainsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists delete internal domains internal server error response has a 4xx status code
func (o *InternalDomainListsDeleteInternalDomainsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal domain lists delete internal domains internal server error response has a 5xx status code
func (o *InternalDomainListsDeleteInternalDomainsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal domain lists delete internal domains internal server error response a status code equal to that given
func (o *InternalDomainListsDeleteInternalDomainsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal domain lists delete internal domains internal server error response
func (o *InternalDomainListsDeleteInternalDomainsInternalServerError) Code() int {
	return 500
}

func (o *InternalDomainListsDeleteInternalDomainsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /internal_domain_lists][%d] internalDomainListsDeleteInternalDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalDomainListsDeleteInternalDomainsInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /internal_domain_lists][%d] internalDomainListsDeleteInternalDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalDomainListsDeleteInternalDomainsInternalServerError) GetPayload() *InternalDomainListsDeleteInternalDomainsInternalServerErrorBody {
	return o.Payload
}

func (o *InternalDomainListsDeleteInternalDomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InternalDomainListsDeleteInternalDomainsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
InternalDomainListsDeleteInternalDomainsBadRequestBody internal domain lists delete internal domains bad request body
swagger:model InternalDomainListsDeleteInternalDomainsBadRequestBody
*/
type InternalDomainListsDeleteInternalDomainsBadRequestBody struct {

	// error
	Error *InternalDomainListsDeleteInternalDomainsBadRequestBodyError `json:"error,omitempty"`
}

// Validate validates this internal domain lists delete internal domains bad request body
func (o *InternalDomainListsDeleteInternalDomainsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsDeleteInternalDomainsBadRequestBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsDeleteInternalDomainsBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsDeleteInternalDomainsBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this internal domain lists delete internal domains bad request body based on the context it is used
func (o *InternalDomainListsDeleteInternalDomainsBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsDeleteInternalDomainsBadRequestBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsDeleteInternalDomainsBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsDeleteInternalDomainsBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsDeleteInternalDomainsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsDeleteInternalDomainsBadRequestBodyError internal domain lists delete internal domains bad request body error
swagger:model InternalDomainListsDeleteInternalDomainsBadRequestBodyError
*/
type InternalDomainListsDeleteInternalDomainsBadRequestBodyError struct {

	// code
	// Example: FAILED_PRECONDITION
	Code string `json:"code,omitempty"`

	// message
	// Example: Internal Domains List ids can't be empty
	Message string `json:"message,omitempty"`

	// status
	// Example: 400
	Status string `json:"status,omitempty"`
}

// Validate validates this internal domain lists delete internal domains bad request body error
func (o *InternalDomainListsDeleteInternalDomainsBadRequestBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this internal domain lists delete internal domains bad request body error based on context it is used
func (o *InternalDomainListsDeleteInternalDomainsBadRequestBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsBadRequestBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsBadRequestBodyError) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsDeleteInternalDomainsBadRequestBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsDeleteInternalDomainsInternalServerErrorBody internal domain lists delete internal domains internal server error body
swagger:model InternalDomainListsDeleteInternalDomainsInternalServerErrorBody
*/
type InternalDomainListsDeleteInternalDomainsInternalServerErrorBody struct {

	// error
	Error *InternalDomainListsDeleteInternalDomainsInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this internal domain lists delete internal domains internal server error body
func (o *InternalDomainListsDeleteInternalDomainsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsDeleteInternalDomainsInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsDeleteInternalDomainsInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsDeleteInternalDomainsInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this internal domain lists delete internal domains internal server error body based on the context it is used
func (o *InternalDomainListsDeleteInternalDomainsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsDeleteInternalDomainsInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsDeleteInternalDomainsInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsDeleteInternalDomainsInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsDeleteInternalDomainsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsDeleteInternalDomainsInternalServerErrorBodyError internal domain lists delete internal domains internal server error body error
swagger:model InternalDomainListsDeleteInternalDomainsInternalServerErrorBodyError
*/
type InternalDomainListsDeleteInternalDomainsInternalServerErrorBodyError struct {

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

// Validate validates this internal domain lists delete internal domains internal server error body error
func (o *InternalDomainListsDeleteInternalDomainsInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this internal domain lists delete internal domains internal server error body error based on context it is used
func (o *InternalDomainListsDeleteInternalDomainsInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsDeleteInternalDomainsInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsDeleteInternalDomainsNotFoundBody internal domain lists delete internal domains not found body
swagger:model InternalDomainListsDeleteInternalDomainsNotFoundBody
*/
type InternalDomainListsDeleteInternalDomainsNotFoundBody struct {

	// error
	Error *InternalDomainListsDeleteInternalDomainsNotFoundBodyError `json:"error,omitempty"`
}

// Validate validates this internal domain lists delete internal domains not found body
func (o *InternalDomainListsDeleteInternalDomainsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsDeleteInternalDomainsNotFoundBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsDeleteInternalDomainsNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsDeleteInternalDomainsNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this internal domain lists delete internal domains not found body based on the context it is used
func (o *InternalDomainListsDeleteInternalDomainsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsDeleteInternalDomainsNotFoundBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsDeleteInternalDomainsNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsDeleteInternalDomainsNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsDeleteInternalDomainsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsDeleteInternalDomainsNotFoundBodyError internal domain lists delete internal domains not found body error
swagger:model InternalDomainListsDeleteInternalDomainsNotFoundBodyError
*/
type InternalDomainListsDeleteInternalDomainsNotFoundBodyError struct {

	// code
	// Example: NOT_FOUND
	Code string `json:"code,omitempty"`

	// message
	// Example: Internal Domain List does not exist
	Message string `json:"message,omitempty"`

	// status
	// Example: 404
	Status string `json:"status,omitempty"`
}

// Validate validates this internal domain lists delete internal domains not found body error
func (o *InternalDomainListsDeleteInternalDomainsNotFoundBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this internal domain lists delete internal domains not found body error based on context it is used
func (o *InternalDomainListsDeleteInternalDomainsNotFoundBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsNotFoundBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsDeleteInternalDomainsNotFoundBodyError) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsDeleteInternalDomainsNotFoundBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}