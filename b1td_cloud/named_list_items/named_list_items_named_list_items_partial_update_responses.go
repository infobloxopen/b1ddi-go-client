// Code generated by go-swagger; DO NOT EDIT.

package named_list_items

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

// NamedListItemsNamedListItemsPartialUpdateReader is a Reader for the NamedListItemsNamedListItemsPartialUpdate structure.
type NamedListItemsNamedListItemsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NamedListItemsNamedListItemsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNamedListItemsNamedListItemsPartialUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNamedListItemsNamedListItemsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNamedListItemsNamedListItemsPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNamedListItemsNamedListItemsPartialUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /named_lists/{id}/items] named_list_itemsNamedListItemsPartialUpdate", response, response.Code())
	}
}

// NewNamedListItemsNamedListItemsPartialUpdateCreated creates a NamedListItemsNamedListItemsPartialUpdateCreated with default headers values
func NewNamedListItemsNamedListItemsPartialUpdateCreated() *NamedListItemsNamedListItemsPartialUpdateCreated {
	return &NamedListItemsNamedListItemsPartialUpdateCreated{}
}

/*
NamedListItemsNamedListItemsPartialUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type NamedListItemsNamedListItemsPartialUpdateCreated struct {
	Payload models.AtcfwNamedListItemsPartialUpdateResponse
}

// IsSuccess returns true when this named list items named list items partial update created response has a 2xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this named list items named list items partial update created response has a 3xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named list items named list items partial update created response has a 4xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this named list items named list items partial update created response has a 5xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this named list items named list items partial update created response a status code equal to that given
func (o *NamedListItemsNamedListItemsPartialUpdateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the named list items named list items partial update created response
func (o *NamedListItemsNamedListItemsPartialUpdateCreated) Code() int {
	return 201
}

func (o *NamedListItemsNamedListItemsPartialUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /named_lists/{id}/items][%d] namedListItemsNamedListItemsPartialUpdateCreated  %+v", 201, o.Payload)
}

func (o *NamedListItemsNamedListItemsPartialUpdateCreated) String() string {
	return fmt.Sprintf("[PATCH /named_lists/{id}/items][%d] namedListItemsNamedListItemsPartialUpdateCreated  %+v", 201, o.Payload)
}

func (o *NamedListItemsNamedListItemsPartialUpdateCreated) GetPayload() models.AtcfwNamedListItemsPartialUpdateResponse {
	return o.Payload
}

func (o *NamedListItemsNamedListItemsPartialUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamedListItemsNamedListItemsPartialUpdateBadRequest creates a NamedListItemsNamedListItemsPartialUpdateBadRequest with default headers values
func NewNamedListItemsNamedListItemsPartialUpdateBadRequest() *NamedListItemsNamedListItemsPartialUpdateBadRequest {
	return &NamedListItemsNamedListItemsPartialUpdateBadRequest{}
}

/*
	NamedListItemsNamedListItemsPartialUpdateBadRequest describes a response with status code 400, with default header values.

- 'id' value must be greater than or equal to zero
- 'items' value must contain either valid domain names or IPv4 or network addresses.
*/
type NamedListItemsNamedListItemsPartialUpdateBadRequest struct {
	Payload *NamedListItemsNamedListItemsPartialUpdateBadRequestBody
}

// IsSuccess returns true when this named list items named list items partial update bad request response has a 2xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this named list items named list items partial update bad request response has a 3xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named list items named list items partial update bad request response has a 4xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this named list items named list items partial update bad request response has a 5xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this named list items named list items partial update bad request response a status code equal to that given
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the named list items named list items partial update bad request response
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequest) Code() int {
	return 400
}

func (o *NamedListItemsNamedListItemsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /named_lists/{id}/items][%d] namedListItemsNamedListItemsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *NamedListItemsNamedListItemsPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /named_lists/{id}/items][%d] namedListItemsNamedListItemsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *NamedListItemsNamedListItemsPartialUpdateBadRequest) GetPayload() *NamedListItemsNamedListItemsPartialUpdateBadRequestBody {
	return o.Payload
}

func (o *NamedListItemsNamedListItemsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NamedListItemsNamedListItemsPartialUpdateBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamedListItemsNamedListItemsPartialUpdateNotFound creates a NamedListItemsNamedListItemsPartialUpdateNotFound with default headers values
func NewNamedListItemsNamedListItemsPartialUpdateNotFound() *NamedListItemsNamedListItemsPartialUpdateNotFound {
	return &NamedListItemsNamedListItemsPartialUpdateNotFound{}
}

/*
	NamedListItemsNamedListItemsPartialUpdateNotFound describes a response with status code 404, with default header values.

- 'id' value must contain existing named list identifier
*/
type NamedListItemsNamedListItemsPartialUpdateNotFound struct {
	Payload *NamedListItemsNamedListItemsPartialUpdateNotFoundBody
}

// IsSuccess returns true when this named list items named list items partial update not found response has a 2xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this named list items named list items partial update not found response has a 3xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named list items named list items partial update not found response has a 4xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this named list items named list items partial update not found response has a 5xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this named list items named list items partial update not found response a status code equal to that given
func (o *NamedListItemsNamedListItemsPartialUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the named list items named list items partial update not found response
func (o *NamedListItemsNamedListItemsPartialUpdateNotFound) Code() int {
	return 404
}

func (o *NamedListItemsNamedListItemsPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /named_lists/{id}/items][%d] namedListItemsNamedListItemsPartialUpdateNotFound  %+v", 404, o.Payload)
}

func (o *NamedListItemsNamedListItemsPartialUpdateNotFound) String() string {
	return fmt.Sprintf("[PATCH /named_lists/{id}/items][%d] namedListItemsNamedListItemsPartialUpdateNotFound  %+v", 404, o.Payload)
}

func (o *NamedListItemsNamedListItemsPartialUpdateNotFound) GetPayload() *NamedListItemsNamedListItemsPartialUpdateNotFoundBody {
	return o.Payload
}

func (o *NamedListItemsNamedListItemsPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NamedListItemsNamedListItemsPartialUpdateNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamedListItemsNamedListItemsPartialUpdateInternalServerError creates a NamedListItemsNamedListItemsPartialUpdateInternalServerError with default headers values
func NewNamedListItemsNamedListItemsPartialUpdateInternalServerError() *NamedListItemsNamedListItemsPartialUpdateInternalServerError {
	return &NamedListItemsNamedListItemsPartialUpdateInternalServerError{}
}

/*
	NamedListItemsNamedListItemsPartialUpdateInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type NamedListItemsNamedListItemsPartialUpdateInternalServerError struct {
	Payload *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody
}

// IsSuccess returns true when this named list items named list items partial update internal server error response has a 2xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this named list items named list items partial update internal server error response has a 3xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named list items named list items partial update internal server error response has a 4xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this named list items named list items partial update internal server error response has a 5xx status code
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this named list items named list items partial update internal server error response a status code equal to that given
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the named list items named list items partial update internal server error response
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerError) Code() int {
	return 500
}

func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /named_lists/{id}/items][%d] namedListItemsNamedListItemsPartialUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /named_lists/{id}/items][%d] namedListItemsNamedListItemsPartialUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerError) GetPayload() *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody {
	return o.Payload
}

func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NamedListItemsNamedListItemsPartialUpdateBadRequestBody named list items named list items partial update bad request body
swagger:model NamedListItemsNamedListItemsPartialUpdateBadRequestBody
*/
type NamedListItemsNamedListItemsPartialUpdateBadRequestBody struct {

	// error
	Error *NamedListItemsNamedListItemsPartialUpdateBadRequestBodyError `json:"error,omitempty"`
}

// Validate validates this named list items named list items partial update bad request body
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsNamedListItemsPartialUpdateBadRequestBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsNamedListItemsPartialUpdateBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsNamedListItemsPartialUpdateBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this named list items named list items partial update bad request body based on the context it is used
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsNamedListItemsPartialUpdateBadRequestBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsNamedListItemsPartialUpdateBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsNamedListItemsPartialUpdateBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res NamedListItemsNamedListItemsPartialUpdateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListItemsNamedListItemsPartialUpdateBadRequestBodyError named list items named list items partial update bad request body error
swagger:model NamedListItemsNamedListItemsPartialUpdateBadRequestBodyError
*/
type NamedListItemsNamedListItemsPartialUpdateBadRequestBodyError struct {

	// code
	// Example: INVALID_ARGUMENT
	Code string `json:"code,omitempty"`

	// message
	// Example: Invalid domain or IP address or network '1.1.1.1.1'
	Message string `json:"message,omitempty"`

	// status
	// Example: 400
	Status string `json:"status,omitempty"`
}

// Validate validates this named list items named list items partial update bad request body error
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequestBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this named list items named list items partial update bad request body error based on context it is used
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequestBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequestBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateBadRequestBodyError) UnmarshalBinary(b []byte) error {
	var res NamedListItemsNamedListItemsPartialUpdateBadRequestBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody named list items named list items partial update internal server error body
swagger:model NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody
*/
type NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody struct {

	// error
	Error *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this named list items named list items partial update internal server error body
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsNamedListItemsPartialUpdateInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsNamedListItemsPartialUpdateInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this named list items named list items partial update internal server error body based on the context it is used
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsNamedListItemsPartialUpdateInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsNamedListItemsPartialUpdateInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBodyError named list items named list items partial update internal server error body error
swagger:model NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBodyError
*/
type NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBodyError struct {

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

// Validate validates this named list items named list items partial update internal server error body error
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this named list items named list items partial update internal server error body error based on context it is used
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res NamedListItemsNamedListItemsPartialUpdateInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListItemsNamedListItemsPartialUpdateNotFoundBody named list items named list items partial update not found body
swagger:model NamedListItemsNamedListItemsPartialUpdateNotFoundBody
*/
type NamedListItemsNamedListItemsPartialUpdateNotFoundBody struct {

	// error
	Error *NamedListItemsNamedListItemsPartialUpdateNotFoundBodyError `json:"error,omitempty"`
}

// Validate validates this named list items named list items partial update not found body
func (o *NamedListItemsNamedListItemsPartialUpdateNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsNamedListItemsPartialUpdateNotFoundBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsNamedListItemsPartialUpdateNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsNamedListItemsPartialUpdateNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this named list items named list items partial update not found body based on the context it is used
func (o *NamedListItemsNamedListItemsPartialUpdateNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsNamedListItemsPartialUpdateNotFoundBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsNamedListItemsPartialUpdateNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsNamedListItemsPartialUpdateNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res NamedListItemsNamedListItemsPartialUpdateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListItemsNamedListItemsPartialUpdateNotFoundBodyError named list items named list items partial update not found body error
swagger:model NamedListItemsNamedListItemsPartialUpdateNotFoundBodyError
*/
type NamedListItemsNamedListItemsPartialUpdateNotFoundBodyError struct {

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

// Validate validates this named list items named list items partial update not found body error
func (o *NamedListItemsNamedListItemsPartialUpdateNotFoundBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this named list items named list items partial update not found body error based on context it is used
func (o *NamedListItemsNamedListItemsPartialUpdateNotFoundBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateNotFoundBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsNamedListItemsPartialUpdateNotFoundBodyError) UnmarshalBinary(b []byte) error {
	var res NamedListItemsNamedListItemsPartialUpdateNotFoundBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}