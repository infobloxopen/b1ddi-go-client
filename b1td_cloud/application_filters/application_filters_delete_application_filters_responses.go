// Code generated by go-swagger; DO NOT EDIT.

package application_filters

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

// ApplicationFiltersDeleteApplicationFiltersReader is a Reader for the ApplicationFiltersDeleteApplicationFilters structure.
type ApplicationFiltersDeleteApplicationFiltersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationFiltersDeleteApplicationFiltersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewApplicationFiltersDeleteApplicationFiltersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewApplicationFiltersDeleteApplicationFiltersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewApplicationFiltersDeleteApplicationFiltersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /application_filters] application_filtersDeleteApplicationFilters", response, response.Code())
	}
}

// NewApplicationFiltersDeleteApplicationFiltersNoContent creates a ApplicationFiltersDeleteApplicationFiltersNoContent with default headers values
func NewApplicationFiltersDeleteApplicationFiltersNoContent() *ApplicationFiltersDeleteApplicationFiltersNoContent {
	return &ApplicationFiltersDeleteApplicationFiltersNoContent{}
}

/*
ApplicationFiltersDeleteApplicationFiltersNoContent describes a response with status code 204, with default header values.

No Content
*/
type ApplicationFiltersDeleteApplicationFiltersNoContent struct {
}

// IsSuccess returns true when this application filters delete application filters no content response has a 2xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application filters delete application filters no content response has a 3xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application filters delete application filters no content response has a 4xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this application filters delete application filters no content response has a 5xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this application filters delete application filters no content response a status code equal to that given
func (o *ApplicationFiltersDeleteApplicationFiltersNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the application filters delete application filters no content response
func (o *ApplicationFiltersDeleteApplicationFiltersNoContent) Code() int {
	return 204
}

func (o *ApplicationFiltersDeleteApplicationFiltersNoContent) Error() string {
	return fmt.Sprintf("[DELETE /application_filters][%d] applicationFiltersDeleteApplicationFiltersNoContent ", 204)
}

func (o *ApplicationFiltersDeleteApplicationFiltersNoContent) String() string {
	return fmt.Sprintf("[DELETE /application_filters][%d] applicationFiltersDeleteApplicationFiltersNoContent ", 204)
}

func (o *ApplicationFiltersDeleteApplicationFiltersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApplicationFiltersDeleteApplicationFiltersBadRequest creates a ApplicationFiltersDeleteApplicationFiltersBadRequest with default headers values
func NewApplicationFiltersDeleteApplicationFiltersBadRequest() *ApplicationFiltersDeleteApplicationFiltersBadRequest {
	return &ApplicationFiltersDeleteApplicationFiltersBadRequest{}
}

/*
	ApplicationFiltersDeleteApplicationFiltersBadRequest describes a response with status code 400, with default header values.

- 'ids' value must contain existing application filter identifiers
- application filters assigned to a security policy cannot be deleted
- application filters assigned to a bypass code cannot be deleted
*/
type ApplicationFiltersDeleteApplicationFiltersBadRequest struct {
	Payload *ApplicationFiltersDeleteApplicationFiltersBadRequestBody
}

// IsSuccess returns true when this application filters delete application filters bad request response has a 2xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this application filters delete application filters bad request response has a 3xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application filters delete application filters bad request response has a 4xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this application filters delete application filters bad request response has a 5xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this application filters delete application filters bad request response a status code equal to that given
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the application filters delete application filters bad request response
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequest) Code() int {
	return 400
}

func (o *ApplicationFiltersDeleteApplicationFiltersBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /application_filters][%d] applicationFiltersDeleteApplicationFiltersBadRequest  %+v", 400, o.Payload)
}

func (o *ApplicationFiltersDeleteApplicationFiltersBadRequest) String() string {
	return fmt.Sprintf("[DELETE /application_filters][%d] applicationFiltersDeleteApplicationFiltersBadRequest  %+v", 400, o.Payload)
}

func (o *ApplicationFiltersDeleteApplicationFiltersBadRequest) GetPayload() *ApplicationFiltersDeleteApplicationFiltersBadRequestBody {
	return o.Payload
}

func (o *ApplicationFiltersDeleteApplicationFiltersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ApplicationFiltersDeleteApplicationFiltersBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationFiltersDeleteApplicationFiltersInternalServerError creates a ApplicationFiltersDeleteApplicationFiltersInternalServerError with default headers values
func NewApplicationFiltersDeleteApplicationFiltersInternalServerError() *ApplicationFiltersDeleteApplicationFiltersInternalServerError {
	return &ApplicationFiltersDeleteApplicationFiltersInternalServerError{}
}

/*
	ApplicationFiltersDeleteApplicationFiltersInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type ApplicationFiltersDeleteApplicationFiltersInternalServerError struct {
	Payload *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody
}

// IsSuccess returns true when this application filters delete application filters internal server error response has a 2xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this application filters delete application filters internal server error response has a 3xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application filters delete application filters internal server error response has a 4xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this application filters delete application filters internal server error response has a 5xx status code
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this application filters delete application filters internal server error response a status code equal to that given
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the application filters delete application filters internal server error response
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerError) Code() int {
	return 500
}

func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /application_filters][%d] applicationFiltersDeleteApplicationFiltersInternalServerError  %+v", 500, o.Payload)
}

func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /application_filters][%d] applicationFiltersDeleteApplicationFiltersInternalServerError  %+v", 500, o.Payload)
}

func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerError) GetPayload() *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody {
	return o.Payload
}

func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ApplicationFiltersDeleteApplicationFiltersBadRequestBody application filters delete application filters bad request body
swagger:model ApplicationFiltersDeleteApplicationFiltersBadRequestBody
*/
type ApplicationFiltersDeleteApplicationFiltersBadRequestBody struct {

	// error
	Error *ApplicationFiltersDeleteApplicationFiltersBadRequestBodyError `json:"error,omitempty"`
}

// Validate validates this application filters delete application filters bad request body
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ApplicationFiltersDeleteApplicationFiltersBadRequestBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationFiltersDeleteApplicationFiltersBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationFiltersDeleteApplicationFiltersBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application filters delete application filters bad request body based on the context it is used
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ApplicationFiltersDeleteApplicationFiltersBadRequestBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationFiltersDeleteApplicationFiltersBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationFiltersDeleteApplicationFiltersBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ApplicationFiltersDeleteApplicationFiltersBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ApplicationFiltersDeleteApplicationFiltersBadRequestBodyError application filters delete application filters bad request body error
swagger:model ApplicationFiltersDeleteApplicationFiltersBadRequestBodyError
*/
type ApplicationFiltersDeleteApplicationFiltersBadRequestBodyError struct {

	// code
	// Example: INVALID_ARGUMENT
	Code string `json:"code,omitempty"`

	// message
	// Example: Application Filter ids can't be empty
	Message string `json:"message,omitempty"`

	// status
	// Example: 400
	Status string `json:"status,omitempty"`
}

// Validate validates this application filters delete application filters bad request body error
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequestBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this application filters delete application filters bad request body error based on context it is used
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequestBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequestBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ApplicationFiltersDeleteApplicationFiltersBadRequestBodyError) UnmarshalBinary(b []byte) error {
	var res ApplicationFiltersDeleteApplicationFiltersBadRequestBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody application filters delete application filters internal server error body
swagger:model ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody
*/
type ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody struct {

	// error
	Error *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this application filters delete application filters internal server error body
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationFiltersDeleteApplicationFiltersInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationFiltersDeleteApplicationFiltersInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application filters delete application filters internal server error body based on the context it is used
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationFiltersDeleteApplicationFiltersInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationFiltersDeleteApplicationFiltersInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBodyError application filters delete application filters internal server error body error
swagger:model ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBodyError
*/
type ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBodyError struct {

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

// Validate validates this application filters delete application filters internal server error body error
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this application filters delete application filters internal server error body error based on context it is used
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res ApplicationFiltersDeleteApplicationFiltersInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}