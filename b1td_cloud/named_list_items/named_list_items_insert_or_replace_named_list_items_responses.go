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

// NamedListItemsInsertOrReplaceNamedListItemsReader is a Reader for the NamedListItemsInsertOrReplaceNamedListItems structure.
type NamedListItemsInsertOrReplaceNamedListItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NamedListItemsInsertOrReplaceNamedListItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNamedListItemsInsertOrReplaceNamedListItemsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNamedListItemsInsertOrReplaceNamedListItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNamedListItemsInsertOrReplaceNamedListItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNamedListItemsInsertOrReplaceNamedListItemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /named_lists/{id}/items] named_list_itemsInsertOrReplaceNamedListItems", response, response.Code())
	}
}

// NewNamedListItemsInsertOrReplaceNamedListItemsCreated creates a NamedListItemsInsertOrReplaceNamedListItemsCreated with default headers values
func NewNamedListItemsInsertOrReplaceNamedListItemsCreated() *NamedListItemsInsertOrReplaceNamedListItemsCreated {
	return &NamedListItemsInsertOrReplaceNamedListItemsCreated{}
}

/*
NamedListItemsInsertOrReplaceNamedListItemsCreated describes a response with status code 201, with default header values.

POST operation response
*/
type NamedListItemsInsertOrReplaceNamedListItemsCreated struct {
	Payload *models.AtcfwNamedListItemsInsertOrUpdateResponse
}

// IsSuccess returns true when this named list items insert or replace named list items created response has a 2xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this named list items insert or replace named list items created response has a 3xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named list items insert or replace named list items created response has a 4xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this named list items insert or replace named list items created response has a 5xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this named list items insert or replace named list items created response a status code equal to that given
func (o *NamedListItemsInsertOrReplaceNamedListItemsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the named list items insert or replace named list items created response
func (o *NamedListItemsInsertOrReplaceNamedListItemsCreated) Code() int {
	return 201
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsCreated) Error() string {
	return fmt.Sprintf("[POST /named_lists/{id}/items][%d] namedListItemsInsertOrReplaceNamedListItemsCreated  %+v", 201, o.Payload)
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsCreated) String() string {
	return fmt.Sprintf("[POST /named_lists/{id}/items][%d] namedListItemsInsertOrReplaceNamedListItemsCreated  %+v", 201, o.Payload)
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsCreated) GetPayload() *models.AtcfwNamedListItemsInsertOrUpdateResponse {
	return o.Payload
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AtcfwNamedListItemsInsertOrUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamedListItemsInsertOrReplaceNamedListItemsBadRequest creates a NamedListItemsInsertOrReplaceNamedListItemsBadRequest with default headers values
func NewNamedListItemsInsertOrReplaceNamedListItemsBadRequest() *NamedListItemsInsertOrReplaceNamedListItemsBadRequest {
	return &NamedListItemsInsertOrReplaceNamedListItemsBadRequest{}
}

/*
	NamedListItemsInsertOrReplaceNamedListItemsBadRequest describes a response with status code 400, with default header values.

- 'id' value must be greater than or equal to zero
- 'items' value must contain either valid domain names or IPv4 or IPv6 or network addresses.
*/
type NamedListItemsInsertOrReplaceNamedListItemsBadRequest struct {
	Payload *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody
}

// IsSuccess returns true when this named list items insert or replace named list items bad request response has a 2xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this named list items insert or replace named list items bad request response has a 3xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named list items insert or replace named list items bad request response has a 4xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this named list items insert or replace named list items bad request response has a 5xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this named list items insert or replace named list items bad request response a status code equal to that given
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the named list items insert or replace named list items bad request response
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequest) Code() int {
	return 400
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequest) Error() string {
	return fmt.Sprintf("[POST /named_lists/{id}/items][%d] namedListItemsInsertOrReplaceNamedListItemsBadRequest  %+v", 400, o.Payload)
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequest) String() string {
	return fmt.Sprintf("[POST /named_lists/{id}/items][%d] namedListItemsInsertOrReplaceNamedListItemsBadRequest  %+v", 400, o.Payload)
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequest) GetPayload() *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody {
	return o.Payload
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamedListItemsInsertOrReplaceNamedListItemsNotFound creates a NamedListItemsInsertOrReplaceNamedListItemsNotFound with default headers values
func NewNamedListItemsInsertOrReplaceNamedListItemsNotFound() *NamedListItemsInsertOrReplaceNamedListItemsNotFound {
	return &NamedListItemsInsertOrReplaceNamedListItemsNotFound{}
}

/*
	NamedListItemsInsertOrReplaceNamedListItemsNotFound describes a response with status code 404, with default header values.

- 'id' value must contain existing named list identifier
*/
type NamedListItemsInsertOrReplaceNamedListItemsNotFound struct {
	Payload *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody
}

// IsSuccess returns true when this named list items insert or replace named list items not found response has a 2xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this named list items insert or replace named list items not found response has a 3xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named list items insert or replace named list items not found response has a 4xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this named list items insert or replace named list items not found response has a 5xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this named list items insert or replace named list items not found response a status code equal to that given
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the named list items insert or replace named list items not found response
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFound) Code() int {
	return 404
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFound) Error() string {
	return fmt.Sprintf("[POST /named_lists/{id}/items][%d] namedListItemsInsertOrReplaceNamedListItemsNotFound  %+v", 404, o.Payload)
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFound) String() string {
	return fmt.Sprintf("[POST /named_lists/{id}/items][%d] namedListItemsInsertOrReplaceNamedListItemsNotFound  %+v", 404, o.Payload)
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFound) GetPayload() *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody {
	return o.Payload
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamedListItemsInsertOrReplaceNamedListItemsInternalServerError creates a NamedListItemsInsertOrReplaceNamedListItemsInternalServerError with default headers values
func NewNamedListItemsInsertOrReplaceNamedListItemsInternalServerError() *NamedListItemsInsertOrReplaceNamedListItemsInternalServerError {
	return &NamedListItemsInsertOrReplaceNamedListItemsInternalServerError{}
}

/*
	NamedListItemsInsertOrReplaceNamedListItemsInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type NamedListItemsInsertOrReplaceNamedListItemsInternalServerError struct {
	Payload *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody
}

// IsSuccess returns true when this named list items insert or replace named list items internal server error response has a 2xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this named list items insert or replace named list items internal server error response has a 3xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this named list items insert or replace named list items internal server error response has a 4xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this named list items insert or replace named list items internal server error response has a 5xx status code
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this named list items insert or replace named list items internal server error response a status code equal to that given
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the named list items insert or replace named list items internal server error response
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerError) Code() int {
	return 500
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /named_lists/{id}/items][%d] namedListItemsInsertOrReplaceNamedListItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerError) String() string {
	return fmt.Sprintf("[POST /named_lists/{id}/items][%d] namedListItemsInsertOrReplaceNamedListItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerError) GetPayload() *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody {
	return o.Payload
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody named list items insert or replace named list items bad request body
swagger:model NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody
*/
type NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody struct {

	// error
	Error *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBodyError `json:"error,omitempty"`
}

// Validate validates this named list items insert or replace named list items bad request body
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsInsertOrReplaceNamedListItemsBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsInsertOrReplaceNamedListItemsBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this named list items insert or replace named list items bad request body based on the context it is used
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsInsertOrReplaceNamedListItemsBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsInsertOrReplaceNamedListItemsBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res NamedListItemsInsertOrReplaceNamedListItemsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListItemsInsertOrReplaceNamedListItemsBadRequestBodyError named list items insert or replace named list items bad request body error
swagger:model NamedListItemsInsertOrReplaceNamedListItemsBadRequestBodyError
*/
type NamedListItemsInsertOrReplaceNamedListItemsBadRequestBodyError struct {

	// code
	// Example: INVALID_ARGUMENT
	Code string `json:"code,omitempty"`

	// message
	// Example: Invalid domain or IP address or network '1.1.1'
	Message string `json:"message,omitempty"`

	// status
	// Example: 400
	Status string `json:"status,omitempty"`
}

// Validate validates this named list items insert or replace named list items bad request body error
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this named list items insert or replace named list items bad request body error based on context it is used
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsBadRequestBodyError) UnmarshalBinary(b []byte) error {
	var res NamedListItemsInsertOrReplaceNamedListItemsBadRequestBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody named list items insert or replace named list items internal server error body
swagger:model NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody
*/
type NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody struct {

	// error
	Error *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this named list items insert or replace named list items internal server error body
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsInsertOrReplaceNamedListItemsInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsInsertOrReplaceNamedListItemsInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this named list items insert or replace named list items internal server error body based on the context it is used
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsInsertOrReplaceNamedListItemsInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsInsertOrReplaceNamedListItemsInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBodyError named list items insert or replace named list items internal server error body error
swagger:model NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBodyError
*/
type NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBodyError struct {

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

// Validate validates this named list items insert or replace named list items internal server error body error
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this named list items insert or replace named list items internal server error body error based on context it is used
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res NamedListItemsInsertOrReplaceNamedListItemsInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody named list items insert or replace named list items not found body
swagger:model NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody
*/
type NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody struct {

	// error
	Error *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBodyError `json:"error,omitempty"`
}

// Validate validates this named list items insert or replace named list items not found body
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsInsertOrReplaceNamedListItemsNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsInsertOrReplaceNamedListItemsNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this named list items insert or replace named list items not found body based on the context it is used
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namedListItemsInsertOrReplaceNamedListItemsNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namedListItemsInsertOrReplaceNamedListItemsNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res NamedListItemsInsertOrReplaceNamedListItemsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NamedListItemsInsertOrReplaceNamedListItemsNotFoundBodyError named list items insert or replace named list items not found body error
swagger:model NamedListItemsInsertOrReplaceNamedListItemsNotFoundBodyError
*/
type NamedListItemsInsertOrReplaceNamedListItemsNotFoundBodyError struct {

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

// Validate validates this named list items insert or replace named list items not found body error
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this named list items insert or replace named list items not found body error based on context it is used
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NamedListItemsInsertOrReplaceNamedListItemsNotFoundBodyError) UnmarshalBinary(b []byte) error {
	var res NamedListItemsInsertOrReplaceNamedListItemsNotFoundBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}