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

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// InternalDomainListsInternalDomainsItemsPartialUpdateReader is a Reader for the InternalDomainListsInternalDomainsItemsPartialUpdate structure.
type InternalDomainListsInternalDomainsItemsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInternalDomainListsInternalDomainsItemsPartialUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalDomainListsInternalDomainsItemsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInternalDomainListsInternalDomainsItemsPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /internal_domain_lists/{id}/items] internal_domain_listsInternalDomainsItemsPartialUpdate", response, response.Code())
	}
}

// NewInternalDomainListsInternalDomainsItemsPartialUpdateCreated creates a InternalDomainListsInternalDomainsItemsPartialUpdateCreated with default headers values
func NewInternalDomainListsInternalDomainsItemsPartialUpdateCreated() *InternalDomainListsInternalDomainsItemsPartialUpdateCreated {
	return &InternalDomainListsInternalDomainsItemsPartialUpdateCreated{}
}

/*
InternalDomainListsInternalDomainsItemsPartialUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type InternalDomainListsInternalDomainsItemsPartialUpdateCreated struct {
	Payload models.AtcfwInternalDomainsItemsPartialResponse
}

// IsSuccess returns true when this internal domain lists internal domains items partial update created response has a 2xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal domain lists internal domains items partial update created response has a 3xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists internal domains items partial update created response has a 4xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal domain lists internal domains items partial update created response has a 5xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this internal domain lists internal domains items partial update created response a status code equal to that given
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the internal domain lists internal domains items partial update created response
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateCreated) Code() int {
	return 201
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /internal_domain_lists/{id}/items][%d] internalDomainListsInternalDomainsItemsPartialUpdateCreated  %+v", 201, o.Payload)
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateCreated) String() string {
	return fmt.Sprintf("[PATCH /internal_domain_lists/{id}/items][%d] internalDomainListsInternalDomainsItemsPartialUpdateCreated  %+v", 201, o.Payload)
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateCreated) GetPayload() models.AtcfwInternalDomainsItemsPartialResponse {
	return o.Payload
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalDomainListsInternalDomainsItemsPartialUpdateBadRequest creates a InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest with default headers values
func NewInternalDomainListsInternalDomainsItemsPartialUpdateBadRequest() *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest {
	return &InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest{}
}

/*
	InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest describes a response with status code 400, with default header values.

- 'id' value must be greater than or equal to zero
- Invalid domain or IP address or network.
*/
type InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest struct {
	Payload *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody
}

// IsSuccess returns true when this internal domain lists internal domains items partial update bad request response has a 2xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal domain lists internal domains items partial update bad request response has a 3xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists internal domains items partial update bad request response has a 4xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal domain lists internal domains items partial update bad request response has a 5xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this internal domain lists internal domains items partial update bad request response a status code equal to that given
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the internal domain lists internal domains items partial update bad request response
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest) Code() int {
	return 400
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /internal_domain_lists/{id}/items][%d] internalDomainListsInternalDomainsItemsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /internal_domain_lists/{id}/items][%d] internalDomainListsInternalDomainsItemsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest) GetPayload() *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody {
	return o.Payload
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalDomainListsInternalDomainsItemsPartialUpdateNotFound creates a InternalDomainListsInternalDomainsItemsPartialUpdateNotFound with default headers values
func NewInternalDomainListsInternalDomainsItemsPartialUpdateNotFound() *InternalDomainListsInternalDomainsItemsPartialUpdateNotFound {
	return &InternalDomainListsInternalDomainsItemsPartialUpdateNotFound{}
}

/*
	InternalDomainListsInternalDomainsItemsPartialUpdateNotFound describes a response with status code 404, with default header values.

- 'id' value must contain existing named list identifier
*/
type InternalDomainListsInternalDomainsItemsPartialUpdateNotFound struct {
	Payload *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody
}

// IsSuccess returns true when this internal domain lists internal domains items partial update not found response has a 2xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal domain lists internal domains items partial update not found response has a 3xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists internal domains items partial update not found response has a 4xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal domain lists internal domains items partial update not found response has a 5xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this internal domain lists internal domains items partial update not found response a status code equal to that given
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the internal domain lists internal domains items partial update not found response
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFound) Code() int {
	return 404
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /internal_domain_lists/{id}/items][%d] internalDomainListsInternalDomainsItemsPartialUpdateNotFound  %+v", 404, o.Payload)
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFound) String() string {
	return fmt.Sprintf("[PATCH /internal_domain_lists/{id}/items][%d] internalDomainListsInternalDomainsItemsPartialUpdateNotFound  %+v", 404, o.Payload)
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFound) GetPayload() *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody {
	return o.Payload
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError creates a InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError with default headers values
func NewInternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError() *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError {
	return &InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError{}
}

/*
	InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError struct {
	Payload *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody
}

// IsSuccess returns true when this internal domain lists internal domains items partial update internal server error response has a 2xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal domain lists internal domains items partial update internal server error response has a 3xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists internal domains items partial update internal server error response has a 4xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal domain lists internal domains items partial update internal server error response has a 5xx status code
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal domain lists internal domains items partial update internal server error response a status code equal to that given
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal domain lists internal domains items partial update internal server error response
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError) Code() int {
	return 500
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /internal_domain_lists/{id}/items][%d] internalDomainListsInternalDomainsItemsPartialUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /internal_domain_lists/{id}/items][%d] internalDomainListsInternalDomainsItemsPartialUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError) GetPayload() *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody {
	return o.Payload
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody internal domain lists internal domains items partial update bad request body
swagger:model InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody
*/
type InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody struct {

	// error
	Error *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBodyError `json:"error,omitempty"`
}

// Validate validates this internal domain lists internal domains items partial update bad request body
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this internal domain lists internal domains items partial update bad request body based on the context it is used
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBodyError internal domain lists internal domains items partial update bad request body error
swagger:model InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBodyError
*/
type InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBodyError struct {

	// code
	// Example: INVALID_ARGUMENT
	Code string `json:"code,omitempty"`

	// message
	// Example: Invalid domain or IPv4 address or network '1.1.1'
	Message string `json:"message,omitempty"`

	// status
	// Example: 400
	Status string `json:"status,omitempty"`
}

// Validate validates this internal domain lists internal domains items partial update bad request body error
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this internal domain lists internal domains items partial update bad request body error based on context it is used
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBodyError) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsInternalDomainsItemsPartialUpdateBadRequestBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody internal domain lists internal domains items partial update internal server error body
swagger:model InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody
*/
type InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody struct {

	// error
	Error *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this internal domain lists internal domains items partial update internal server error body
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this internal domain lists internal domains items partial update internal server error body based on the context it is used
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBodyError internal domain lists internal domains items partial update internal server error body error
swagger:model InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBodyError
*/
type InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBodyError struct {

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

// Validate validates this internal domain lists internal domains items partial update internal server error body error
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this internal domain lists internal domains items partial update internal server error body error based on context it is used
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsInternalDomainsItemsPartialUpdateInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody internal domain lists internal domains items partial update not found body
swagger:model InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody
*/
type InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody struct {

	// error
	Error *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBodyError `json:"error,omitempty"`
}

// Validate validates this internal domain lists internal domains items partial update not found body
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this internal domain lists internal domains items partial update not found body based on the context it is used
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsInternalDomainsItemsPartialUpdateNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBodyError internal domain lists internal domains items partial update not found body error
swagger:model InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBodyError
*/
type InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBodyError struct {

	// code
	// Example: NOT_FOUND
	Code string `json:"code,omitempty"`

	// message
	// Example: List does not exist
	Message string `json:"message,omitempty"`

	// status
	// Example: 404
	Status string `json:"status,omitempty"`
}

// Validate validates this internal domain lists internal domains items partial update not found body error
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this internal domain lists internal domains items partial update not found body error based on context it is used
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBodyError) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsInternalDomainsItemsPartialUpdateNotFoundBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}