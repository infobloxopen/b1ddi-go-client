// Code generated by go-swagger; DO NOT EDIT.

package network_lists

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

// NetworkListsDeleteNetworkListsReader is a Reader for the NetworkListsDeleteNetworkLists structure.
type NetworkListsDeleteNetworkListsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkListsDeleteNetworkListsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewNetworkListsDeleteNetworkListsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNetworkListsDeleteNetworkListsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNetworkListsDeleteNetworkListsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNetworkListsDeleteNetworkListsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /network_lists] network_listsDeleteNetworkLists", response, response.Code())
	}
}

// NewNetworkListsDeleteNetworkListsNoContent creates a NetworkListsDeleteNetworkListsNoContent with default headers values
func NewNetworkListsDeleteNetworkListsNoContent() *NetworkListsDeleteNetworkListsNoContent {
	return &NetworkListsDeleteNetworkListsNoContent{}
}

/*
NetworkListsDeleteNetworkListsNoContent describes a response with status code 204, with default header values.

No Content
*/
type NetworkListsDeleteNetworkListsNoContent struct {
}

// IsSuccess returns true when this network lists delete network lists no content response has a 2xx status code
func (o *NetworkListsDeleteNetworkListsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network lists delete network lists no content response has a 3xx status code
func (o *NetworkListsDeleteNetworkListsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network lists delete network lists no content response has a 4xx status code
func (o *NetworkListsDeleteNetworkListsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this network lists delete network lists no content response has a 5xx status code
func (o *NetworkListsDeleteNetworkListsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this network lists delete network lists no content response a status code equal to that given
func (o *NetworkListsDeleteNetworkListsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the network lists delete network lists no content response
func (o *NetworkListsDeleteNetworkListsNoContent) Code() int {
	return 204
}

func (o *NetworkListsDeleteNetworkListsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /network_lists][%d] networkListsDeleteNetworkListsNoContent ", 204)
}

func (o *NetworkListsDeleteNetworkListsNoContent) String() string {
	return fmt.Sprintf("[DELETE /network_lists][%d] networkListsDeleteNetworkListsNoContent ", 204)
}

func (o *NetworkListsDeleteNetworkListsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkListsDeleteNetworkListsBadRequest creates a NetworkListsDeleteNetworkListsBadRequest with default headers values
func NewNetworkListsDeleteNetworkListsBadRequest() *NetworkListsDeleteNetworkListsBadRequest {
	return &NetworkListsDeleteNetworkListsBadRequest{}
}

/*
	NetworkListsDeleteNetworkListsBadRequest describes a response with status code 400, with default header values.

- 'ids' value must be non-empty
- 'ids' value must contain unique elements
- 'ids' value must contain values that are greater than or equal to zero
- network list that is assigned to a security policy cannot be deleted
*/
type NetworkListsDeleteNetworkListsBadRequest struct {
	Payload *NetworkListsDeleteNetworkListsBadRequestBody
}

// IsSuccess returns true when this network lists delete network lists bad request response has a 2xx status code
func (o *NetworkListsDeleteNetworkListsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network lists delete network lists bad request response has a 3xx status code
func (o *NetworkListsDeleteNetworkListsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network lists delete network lists bad request response has a 4xx status code
func (o *NetworkListsDeleteNetworkListsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this network lists delete network lists bad request response has a 5xx status code
func (o *NetworkListsDeleteNetworkListsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this network lists delete network lists bad request response a status code equal to that given
func (o *NetworkListsDeleteNetworkListsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the network lists delete network lists bad request response
func (o *NetworkListsDeleteNetworkListsBadRequest) Code() int {
	return 400
}

func (o *NetworkListsDeleteNetworkListsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /network_lists][%d] networkListsDeleteNetworkListsBadRequest  %+v", 400, o.Payload)
}

func (o *NetworkListsDeleteNetworkListsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /network_lists][%d] networkListsDeleteNetworkListsBadRequest  %+v", 400, o.Payload)
}

func (o *NetworkListsDeleteNetworkListsBadRequest) GetPayload() *NetworkListsDeleteNetworkListsBadRequestBody {
	return o.Payload
}

func (o *NetworkListsDeleteNetworkListsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkListsDeleteNetworkListsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkListsDeleteNetworkListsNotFound creates a NetworkListsDeleteNetworkListsNotFound with default headers values
func NewNetworkListsDeleteNetworkListsNotFound() *NetworkListsDeleteNetworkListsNotFound {
	return &NetworkListsDeleteNetworkListsNotFound{}
}

/*
	NetworkListsDeleteNetworkListsNotFound describes a response with status code 404, with default header values.

- 'ids' value must contain existing network list identifiers
*/
type NetworkListsDeleteNetworkListsNotFound struct {
	Payload *NetworkListsDeleteNetworkListsNotFoundBody
}

// IsSuccess returns true when this network lists delete network lists not found response has a 2xx status code
func (o *NetworkListsDeleteNetworkListsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network lists delete network lists not found response has a 3xx status code
func (o *NetworkListsDeleteNetworkListsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network lists delete network lists not found response has a 4xx status code
func (o *NetworkListsDeleteNetworkListsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this network lists delete network lists not found response has a 5xx status code
func (o *NetworkListsDeleteNetworkListsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this network lists delete network lists not found response a status code equal to that given
func (o *NetworkListsDeleteNetworkListsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the network lists delete network lists not found response
func (o *NetworkListsDeleteNetworkListsNotFound) Code() int {
	return 404
}

func (o *NetworkListsDeleteNetworkListsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /network_lists][%d] networkListsDeleteNetworkListsNotFound  %+v", 404, o.Payload)
}

func (o *NetworkListsDeleteNetworkListsNotFound) String() string {
	return fmt.Sprintf("[DELETE /network_lists][%d] networkListsDeleteNetworkListsNotFound  %+v", 404, o.Payload)
}

func (o *NetworkListsDeleteNetworkListsNotFound) GetPayload() *NetworkListsDeleteNetworkListsNotFoundBody {
	return o.Payload
}

func (o *NetworkListsDeleteNetworkListsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkListsDeleteNetworkListsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkListsDeleteNetworkListsInternalServerError creates a NetworkListsDeleteNetworkListsInternalServerError with default headers values
func NewNetworkListsDeleteNetworkListsInternalServerError() *NetworkListsDeleteNetworkListsInternalServerError {
	return &NetworkListsDeleteNetworkListsInternalServerError{}
}

/*
	NetworkListsDeleteNetworkListsInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type NetworkListsDeleteNetworkListsInternalServerError struct {
	Payload *NetworkListsDeleteNetworkListsInternalServerErrorBody
}

// IsSuccess returns true when this network lists delete network lists internal server error response has a 2xx status code
func (o *NetworkListsDeleteNetworkListsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network lists delete network lists internal server error response has a 3xx status code
func (o *NetworkListsDeleteNetworkListsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network lists delete network lists internal server error response has a 4xx status code
func (o *NetworkListsDeleteNetworkListsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this network lists delete network lists internal server error response has a 5xx status code
func (o *NetworkListsDeleteNetworkListsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this network lists delete network lists internal server error response a status code equal to that given
func (o *NetworkListsDeleteNetworkListsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the network lists delete network lists internal server error response
func (o *NetworkListsDeleteNetworkListsInternalServerError) Code() int {
	return 500
}

func (o *NetworkListsDeleteNetworkListsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /network_lists][%d] networkListsDeleteNetworkListsInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkListsDeleteNetworkListsInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /network_lists][%d] networkListsDeleteNetworkListsInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkListsDeleteNetworkListsInternalServerError) GetPayload() *NetworkListsDeleteNetworkListsInternalServerErrorBody {
	return o.Payload
}

func (o *NetworkListsDeleteNetworkListsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkListsDeleteNetworkListsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NetworkListsDeleteNetworkListsBadRequestBody network lists delete network lists bad request body
swagger:model NetworkListsDeleteNetworkListsBadRequestBody
*/
type NetworkListsDeleteNetworkListsBadRequestBody struct {

	// error
	Error *NetworkListsDeleteNetworkListsBadRequestBodyError `json:"error,omitempty"`
}

// Validate validates this network lists delete network lists bad request body
func (o *NetworkListsDeleteNetworkListsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkListsDeleteNetworkListsBadRequestBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkListsDeleteNetworkListsBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkListsDeleteNetworkListsBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network lists delete network lists bad request body based on the context it is used
func (o *NetworkListsDeleteNetworkListsBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkListsDeleteNetworkListsBadRequestBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkListsDeleteNetworkListsBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkListsDeleteNetworkListsBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res NetworkListsDeleteNetworkListsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetworkListsDeleteNetworkListsBadRequestBodyError network lists delete network lists bad request body error
swagger:model NetworkListsDeleteNetworkListsBadRequestBodyError
*/
type NetworkListsDeleteNetworkListsBadRequestBodyError struct {

	// code
	// Example: INVALID_ARGUMENT
	Code string `json:"code,omitempty"`

	// message
	// Example: Network List ids can't be empty
	Message string `json:"message,omitempty"`

	// status
	// Example: 400
	Status string `json:"status,omitempty"`
}

// Validate validates this network lists delete network lists bad request body error
func (o *NetworkListsDeleteNetworkListsBadRequestBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network lists delete network lists bad request body error based on context it is used
func (o *NetworkListsDeleteNetworkListsBadRequestBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsBadRequestBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsBadRequestBodyError) UnmarshalBinary(b []byte) error {
	var res NetworkListsDeleteNetworkListsBadRequestBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetworkListsDeleteNetworkListsInternalServerErrorBody network lists delete network lists internal server error body
swagger:model NetworkListsDeleteNetworkListsInternalServerErrorBody
*/
type NetworkListsDeleteNetworkListsInternalServerErrorBody struct {

	// error
	Error *NetworkListsDeleteNetworkListsInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this network lists delete network lists internal server error body
func (o *NetworkListsDeleteNetworkListsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkListsDeleteNetworkListsInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkListsDeleteNetworkListsInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkListsDeleteNetworkListsInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network lists delete network lists internal server error body based on the context it is used
func (o *NetworkListsDeleteNetworkListsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkListsDeleteNetworkListsInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkListsDeleteNetworkListsInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkListsDeleteNetworkListsInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res NetworkListsDeleteNetworkListsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetworkListsDeleteNetworkListsInternalServerErrorBodyError network lists delete network lists internal server error body error
swagger:model NetworkListsDeleteNetworkListsInternalServerErrorBodyError
*/
type NetworkListsDeleteNetworkListsInternalServerErrorBodyError struct {

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

// Validate validates this network lists delete network lists internal server error body error
func (o *NetworkListsDeleteNetworkListsInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network lists delete network lists internal server error body error based on context it is used
func (o *NetworkListsDeleteNetworkListsInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res NetworkListsDeleteNetworkListsInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetworkListsDeleteNetworkListsNotFoundBody network lists delete network lists not found body
swagger:model NetworkListsDeleteNetworkListsNotFoundBody
*/
type NetworkListsDeleteNetworkListsNotFoundBody struct {

	// error
	Error *NetworkListsDeleteNetworkListsNotFoundBodyError `json:"error,omitempty"`
}

// Validate validates this network lists delete network lists not found body
func (o *NetworkListsDeleteNetworkListsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkListsDeleteNetworkListsNotFoundBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkListsDeleteNetworkListsNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkListsDeleteNetworkListsNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network lists delete network lists not found body based on the context it is used
func (o *NetworkListsDeleteNetworkListsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkListsDeleteNetworkListsNotFoundBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkListsDeleteNetworkListsNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkListsDeleteNetworkListsNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res NetworkListsDeleteNetworkListsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetworkListsDeleteNetworkListsNotFoundBodyError network lists delete network lists not found body error
swagger:model NetworkListsDeleteNetworkListsNotFoundBodyError
*/
type NetworkListsDeleteNetworkListsNotFoundBodyError struct {

	// code
	// Example: NOT_FOUND
	Code string `json:"code,omitempty"`

	// message
	// Example: Network List does not exist
	Message string `json:"message,omitempty"`

	// status
	// Example: 404
	Status string `json:"status,omitempty"`
}

// Validate validates this network lists delete network lists not found body error
func (o *NetworkListsDeleteNetworkListsNotFoundBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network lists delete network lists not found body error based on context it is used
func (o *NetworkListsDeleteNetworkListsNotFoundBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsNotFoundBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkListsDeleteNetworkListsNotFoundBodyError) UnmarshalBinary(b []byte) error {
	var res NetworkListsDeleteNetworkListsNotFoundBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}