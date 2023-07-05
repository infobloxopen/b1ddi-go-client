// Code generated by go-swagger; DO NOT EDIT.

package named_lists

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

// NamedListsMultiListUpdateReader is a Reader for the NamedListsMultiListUpdate structure.
type NamedListsMultiListUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NamedListsMultiListUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNamedListsMultiListUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNamedListsMultiListUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNamedListsMultiListUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNamedListsMultiListUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /named_lists] named_listsMultiListUpdate", response, response.Code())
	}
}

// NewNamedListsMultiListUpdateCreated creates a NamedListsMultiListUpdateCreated with default headers values
func NewNamedListsMultiListUpdateCreated() *NamedListsMultiListUpdateCreated {
	return &NamedListsMultiListUpdateCreated{}
}

/*
NamedListsMultiListUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type NamedListsMultiListUpdateCreated struct {
	Payload models.AtcfwMultiListUpdateResponse
}

// IsSuccess returns true when this named lists multi list update created response has a 2xx status code
func (o *NamedListsMultiListUpdateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this named lists multi list update created response has a 3xx status code
func (o *NamedListsMultiListUpdateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named lists multi list update created response has a 4xx status code
func (o *NamedListsMultiListUpdateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this named lists multi list update created response has a 5xx status code
func (o *NamedListsMultiListUpdateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this named lists multi list update created response a status code equal to that given
func (o *NamedListsMultiListUpdateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the named lists multi list update created response
func (o *NamedListsMultiListUpdateCreated) Code() int {
	return 201
}

func (o *NamedListsMultiListUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /named_lists][%d] namedListsMultiListUpdateCreated  %+v", 201, o.Payload)
}

func (o *NamedListsMultiListUpdateCreated) String() string {
	return fmt.Sprintf("[PATCH /named_lists][%d] namedListsMultiListUpdateCreated  %+v", 201, o.Payload)
}

func (o *NamedListsMultiListUpdateCreated) GetPayload() models.AtcfwMultiListUpdateResponse {
	return o.Payload
}

func (o *NamedListsMultiListUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamedListsMultiListUpdateBadRequest creates a NamedListsMultiListUpdateBadRequest with default headers values
func NewNamedListsMultiListUpdateBadRequest() *NamedListsMultiListUpdateBadRequest {
	return &NamedListsMultiListUpdateBadRequest{}
}

/*
	NamedListsMultiListUpdateBadRequest describes a response with status code 400, with default header values.

- 'ids' values must be greater than or equal to zero
- 'inserted_items_described' value must contain either valid domain names or IPv4 or network addresses.
*/
type NamedListsMultiListUpdateBadRequest struct {
	Payload *NamedListsMultiListUpdateBadRequestBody
}

// IsSuccess returns true when this named lists multi list update bad request response has a 2xx status code
func (o *NamedListsMultiListUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this named lists multi list update bad request response has a 3xx status code
func (o *NamedListsMultiListUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named lists multi list update bad request response has a 4xx status code
func (o *NamedListsMultiListUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this named lists multi list update bad request response has a 5xx status code
func (o *NamedListsMultiListUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this named lists multi list update bad request response a status code equal to that given
func (o *NamedListsMultiListUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the named lists multi list update bad request response
func (o *NamedListsMultiListUpdateBadRequest) Code() int {
	return 400
}

func (o *NamedListsMultiListUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /named_lists][%d] namedListsMultiListUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *NamedListsMultiListUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /named_lists][%d] namedListsMultiListUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *NamedListsMultiListUpdateBadRequest) GetPayload() *NamedListsMultiListUpdateBadRequestBody {
	return o.Payload
}

func (o *NamedListsMultiListUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NamedListsMultiListUpdateBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamedListsMultiListUpdateNotFound creates a NamedListsMultiListUpdateNotFound with default headers values
func NewNamedListsMultiListUpdateNotFound() *NamedListsMultiListUpdateNotFound {
	return &NamedListsMultiListUpdateNotFound{}
}

/*
	NamedListsMultiListUpdateNotFound describes a response with status code 404, with default header values.

- 'ids' values must contain existing named list identifier
*/
type NamedListsMultiListUpdateNotFound struct {
	Payload *NamedListsMultiListUpdateNotFoundBody
}

// IsSuccess returns true when this named lists multi list update not found response has a 2xx status code
func (o *NamedListsMultiListUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this named lists multi list update not found response has a 3xx status code
func (o *NamedListsMultiListUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named lists multi list update not found response has a 4xx status code
func (o *NamedListsMultiListUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this named lists multi list update not found response has a 5xx status code
func (o *NamedListsMultiListUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this named lists multi list update not found response a status code equal to that given
func (o *NamedListsMultiListUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the named lists multi list update not found response
func (o *NamedListsMultiListUpdateNotFound) Code() int {
	return 404
}

func (o *NamedListsMultiListUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /named_lists][%d] namedListsMultiListUpdateNotFound  %+v", 404, o.Payload)
}

func (o *NamedListsMultiListUpdateNotFound) String() string {
	return fmt.Sprintf("[PATCH /named_lists][%d] namedListsMultiListUpdateNotFound  %+v", 404, o.Payload)
}

func (o *NamedListsMultiListUpdateNotFound) GetPayload() *NamedListsMultiListUpdateNotFoundBody {
	return o.Payload
}

func (o *NamedListsMultiListUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NamedListsMultiListUpdateNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamedListsMultiListUpdateInternalServerError creates a NamedListsMultiListUpdateInternalServerError with default headers values
func NewNamedListsMultiListUpdateInternalServerError() *NamedListsMultiListUpdateInternalServerError {
	return &NamedListsMultiListUpdateInternalServerError{}
}

/*
	NamedListsMultiListUpdateInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type NamedListsMultiListUpdateInternalServerError struct {
	Payload *NamedListsMultiListUpdateInternalServerErrorBody
}

// IsSuccess returns true when this named lists multi list update internal server error response has a 2xx status code
func (o *NamedListsMultiListUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this named lists multi list update internal server error response has a 3xx status code
func (o *NamedListsMultiListUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named lists multi list update internal server error response has a 4xx status code
func (o *NamedListsMultiListUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this named lists multi list update internal server error response has a 5xx status code
func (o *NamedListsMultiListUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this named lists multi list update internal server error response a status code equal to that given
func (o *NamedListsMultiListUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the named lists multi list update internal server error response
func (o *NamedListsMultiListUpdateInternalServerError) Code() int {
	return 500
}

func (o *NamedListsMultiListUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /named_lists][%d] namedListsMultiListUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *NamedListsMultiListUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /named_lists][%d] namedListsMultiListUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *NamedListsMultiListUpdateInternalServerError) GetPayload() *NamedListsMultiListUpdateInternalServerErrorBody {
	return o.Payload
}

func (o *NamedListsMultiListUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NamedListsMultiListUpdateInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NamedListsMultiListUpdateBadRequestBody named lists multi list update bad request body
swagger:model NamedListsMultiListUpdateBadRequestBody
*/
type NamedListsMultiListUpdateBadRequestBody struct {

	// error
	Error *NamedListsMultiListUpdateBadRequestBodyError `json:"error,omitempty"`
}

// Validate validates this named lists multi list update bad request body
func (o *NamedListsMultiListUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListsMultiListUpdateBadRequestBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListsMultiListUpdateBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListsMultiListUpdateBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this named lists multi list update bad request body based on the context it is used
func (o *NamedListsMultiListUpdateBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListsMultiListUpdateBadRequestBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListsMultiListUpdateBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListsMultiListUpdateBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NamedListsMultiListUpdateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListsMultiListUpdateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res NamedListsMultiListUpdateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListsMultiListUpdateBadRequestBodyError named lists multi list update bad request body error
swagger:model NamedListsMultiListUpdateBadRequestBodyError
*/
type NamedListsMultiListUpdateBadRequestBodyError struct {

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

// Validate validates this named lists multi list update bad request body error
func (o *NamedListsMultiListUpdateBadRequestBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this named lists multi list update bad request body error based on context it is used
func (o *NamedListsMultiListUpdateBadRequestBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NamedListsMultiListUpdateBadRequestBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListsMultiListUpdateBadRequestBodyError) UnmarshalBinary(b []byte) error {
	var res NamedListsMultiListUpdateBadRequestBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListsMultiListUpdateInternalServerErrorBody named lists multi list update internal server error body
swagger:model NamedListsMultiListUpdateInternalServerErrorBody
*/
type NamedListsMultiListUpdateInternalServerErrorBody struct {

	// error
	Error *NamedListsMultiListUpdateInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this named lists multi list update internal server error body
func (o *NamedListsMultiListUpdateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListsMultiListUpdateInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListsMultiListUpdateInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListsMultiListUpdateInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this named lists multi list update internal server error body based on the context it is used
func (o *NamedListsMultiListUpdateInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListsMultiListUpdateInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListsMultiListUpdateInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListsMultiListUpdateInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NamedListsMultiListUpdateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListsMultiListUpdateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res NamedListsMultiListUpdateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListsMultiListUpdateInternalServerErrorBodyError named lists multi list update internal server error body error
swagger:model NamedListsMultiListUpdateInternalServerErrorBodyError
*/
type NamedListsMultiListUpdateInternalServerErrorBodyError struct {

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

// Validate validates this named lists multi list update internal server error body error
func (o *NamedListsMultiListUpdateInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this named lists multi list update internal server error body error based on context it is used
func (o *NamedListsMultiListUpdateInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NamedListsMultiListUpdateInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListsMultiListUpdateInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res NamedListsMultiListUpdateInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListsMultiListUpdateNotFoundBody named lists multi list update not found body
swagger:model NamedListsMultiListUpdateNotFoundBody
*/
type NamedListsMultiListUpdateNotFoundBody struct {

	// error
	Error *NamedListsMultiListUpdateNotFoundBodyError `json:"error,omitempty"`
}

// Validate validates this named lists multi list update not found body
func (o *NamedListsMultiListUpdateNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListsMultiListUpdateNotFoundBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListsMultiListUpdateNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListsMultiListUpdateNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this named lists multi list update not found body based on the context it is used
func (o *NamedListsMultiListUpdateNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListsMultiListUpdateNotFoundBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListsMultiListUpdateNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListsMultiListUpdateNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NamedListsMultiListUpdateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListsMultiListUpdateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res NamedListsMultiListUpdateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListsMultiListUpdateNotFoundBodyError named lists multi list update not found body error
swagger:model NamedListsMultiListUpdateNotFoundBodyError
*/
type NamedListsMultiListUpdateNotFoundBodyError struct {

	// code
	// Example: NOT_FOUND
	Code string `json:"code,omitempty"`

	// message
	// Example: Named List does not exist
	Message string `json:"message,omitempty"`

	// status
	// Example: 404
	Status string `json:"status,omitempty"`
}

// Validate validates this named lists multi list update not found body error
func (o *NamedListsMultiListUpdateNotFoundBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this named lists multi list update not found body error based on context it is used
func (o *NamedListsMultiListUpdateNotFoundBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NamedListsMultiListUpdateNotFoundBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListsMultiListUpdateNotFoundBodyError) UnmarshalBinary(b []byte) error {
	var res NamedListsMultiListUpdateNotFoundBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}