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

// InternalDomainListsCreateInternalDomainsReader is a Reader for the InternalDomainListsCreateInternalDomains structure.
type InternalDomainListsCreateInternalDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalDomainListsCreateInternalDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInternalDomainListsCreateInternalDomainsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalDomainListsCreateInternalDomainsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewInternalDomainListsCreateInternalDomainsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalDomainListsCreateInternalDomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /internal_domain_lists] internal_domain_listsCreateInternalDomains", response, response.Code())
	}
}

// NewInternalDomainListsCreateInternalDomainsCreated creates a InternalDomainListsCreateInternalDomainsCreated with default headers values
func NewInternalDomainListsCreateInternalDomainsCreated() *InternalDomainListsCreateInternalDomainsCreated {
	return &InternalDomainListsCreateInternalDomainsCreated{}
}

/*
InternalDomainListsCreateInternalDomainsCreated describes a response with status code 201, with default header values.

POST operation response
*/
type InternalDomainListsCreateInternalDomainsCreated struct {
	Payload *models.AtcfwInternalDomainsCreateResponse
}

// IsSuccess returns true when this internal domain lists create internal domains created response has a 2xx status code
func (o *InternalDomainListsCreateInternalDomainsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal domain lists create internal domains created response has a 3xx status code
func (o *InternalDomainListsCreateInternalDomainsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists create internal domains created response has a 4xx status code
func (o *InternalDomainListsCreateInternalDomainsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal domain lists create internal domains created response has a 5xx status code
func (o *InternalDomainListsCreateInternalDomainsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this internal domain lists create internal domains created response a status code equal to that given
func (o *InternalDomainListsCreateInternalDomainsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the internal domain lists create internal domains created response
func (o *InternalDomainListsCreateInternalDomainsCreated) Code() int {
	return 201
}

func (o *InternalDomainListsCreateInternalDomainsCreated) Error() string {
	return fmt.Sprintf("[POST /internal_domain_lists][%d] internalDomainListsCreateInternalDomainsCreated  %+v", 201, o.Payload)
}

func (o *InternalDomainListsCreateInternalDomainsCreated) String() string {
	return fmt.Sprintf("[POST /internal_domain_lists][%d] internalDomainListsCreateInternalDomainsCreated  %+v", 201, o.Payload)
}

func (o *InternalDomainListsCreateInternalDomainsCreated) GetPayload() *models.AtcfwInternalDomainsCreateResponse {
	return o.Payload
}

func (o *InternalDomainListsCreateInternalDomainsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AtcfwInternalDomainsCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalDomainListsCreateInternalDomainsBadRequest creates a InternalDomainListsCreateInternalDomainsBadRequest with default headers values
func NewInternalDomainListsCreateInternalDomainsBadRequest() *InternalDomainListsCreateInternalDomainsBadRequest {
	return &InternalDomainListsCreateInternalDomainsBadRequest{}
}

/*
	InternalDomainListsCreateInternalDomainsBadRequest describes a response with status code 400, with default header values.

- 'name' length cannot exceed 256 characters limit
- 'description' length cannot exceed 256 characters limit
- 'internal_domains' must not contain existing Internal domain Lists
- 'Non-empty lists'
*/
type InternalDomainListsCreateInternalDomainsBadRequest struct {
	Payload *InternalDomainListsCreateInternalDomainsBadRequestBody
}

// IsSuccess returns true when this internal domain lists create internal domains bad request response has a 2xx status code
func (o *InternalDomainListsCreateInternalDomainsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal domain lists create internal domains bad request response has a 3xx status code
func (o *InternalDomainListsCreateInternalDomainsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists create internal domains bad request response has a 4xx status code
func (o *InternalDomainListsCreateInternalDomainsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal domain lists create internal domains bad request response has a 5xx status code
func (o *InternalDomainListsCreateInternalDomainsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this internal domain lists create internal domains bad request response a status code equal to that given
func (o *InternalDomainListsCreateInternalDomainsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the internal domain lists create internal domains bad request response
func (o *InternalDomainListsCreateInternalDomainsBadRequest) Code() int {
	return 400
}

func (o *InternalDomainListsCreateInternalDomainsBadRequest) Error() string {
	return fmt.Sprintf("[POST /internal_domain_lists][%d] internalDomainListsCreateInternalDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *InternalDomainListsCreateInternalDomainsBadRequest) String() string {
	return fmt.Sprintf("[POST /internal_domain_lists][%d] internalDomainListsCreateInternalDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *InternalDomainListsCreateInternalDomainsBadRequest) GetPayload() *InternalDomainListsCreateInternalDomainsBadRequestBody {
	return o.Payload
}

func (o *InternalDomainListsCreateInternalDomainsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InternalDomainListsCreateInternalDomainsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalDomainListsCreateInternalDomainsConflict creates a InternalDomainListsCreateInternalDomainsConflict with default headers values
func NewInternalDomainListsCreateInternalDomainsConflict() *InternalDomainListsCreateInternalDomainsConflict {
	return &InternalDomainListsCreateInternalDomainsConflict{}
}

/*
	InternalDomainListsCreateInternalDomainsConflict describes a response with status code 409, with default header values.

- 'name' value must be unique among internal domains lists belonging to the same account
*/
type InternalDomainListsCreateInternalDomainsConflict struct {
	Payload *InternalDomainListsCreateInternalDomainsConflictBody
}

// IsSuccess returns true when this internal domain lists create internal domains conflict response has a 2xx status code
func (o *InternalDomainListsCreateInternalDomainsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal domain lists create internal domains conflict response has a 3xx status code
func (o *InternalDomainListsCreateInternalDomainsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists create internal domains conflict response has a 4xx status code
func (o *InternalDomainListsCreateInternalDomainsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal domain lists create internal domains conflict response has a 5xx status code
func (o *InternalDomainListsCreateInternalDomainsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this internal domain lists create internal domains conflict response a status code equal to that given
func (o *InternalDomainListsCreateInternalDomainsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the internal domain lists create internal domains conflict response
func (o *InternalDomainListsCreateInternalDomainsConflict) Code() int {
	return 409
}

func (o *InternalDomainListsCreateInternalDomainsConflict) Error() string {
	return fmt.Sprintf("[POST /internal_domain_lists][%d] internalDomainListsCreateInternalDomainsConflict  %+v", 409, o.Payload)
}

func (o *InternalDomainListsCreateInternalDomainsConflict) String() string {
	return fmt.Sprintf("[POST /internal_domain_lists][%d] internalDomainListsCreateInternalDomainsConflict  %+v", 409, o.Payload)
}

func (o *InternalDomainListsCreateInternalDomainsConflict) GetPayload() *InternalDomainListsCreateInternalDomainsConflictBody {
	return o.Payload
}

func (o *InternalDomainListsCreateInternalDomainsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InternalDomainListsCreateInternalDomainsConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalDomainListsCreateInternalDomainsInternalServerError creates a InternalDomainListsCreateInternalDomainsInternalServerError with default headers values
func NewInternalDomainListsCreateInternalDomainsInternalServerError() *InternalDomainListsCreateInternalDomainsInternalServerError {
	return &InternalDomainListsCreateInternalDomainsInternalServerError{}
}

/*
	InternalDomainListsCreateInternalDomainsInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type InternalDomainListsCreateInternalDomainsInternalServerError struct {
	Payload *InternalDomainListsCreateInternalDomainsInternalServerErrorBody
}

// IsSuccess returns true when this internal domain lists create internal domains internal server error response has a 2xx status code
func (o *InternalDomainListsCreateInternalDomainsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal domain lists create internal domains internal server error response has a 3xx status code
func (o *InternalDomainListsCreateInternalDomainsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal domain lists create internal domains internal server error response has a 4xx status code
func (o *InternalDomainListsCreateInternalDomainsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal domain lists create internal domains internal server error response has a 5xx status code
func (o *InternalDomainListsCreateInternalDomainsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal domain lists create internal domains internal server error response a status code equal to that given
func (o *InternalDomainListsCreateInternalDomainsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal domain lists create internal domains internal server error response
func (o *InternalDomainListsCreateInternalDomainsInternalServerError) Code() int {
	return 500
}

func (o *InternalDomainListsCreateInternalDomainsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /internal_domain_lists][%d] internalDomainListsCreateInternalDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalDomainListsCreateInternalDomainsInternalServerError) String() string {
	return fmt.Sprintf("[POST /internal_domain_lists][%d] internalDomainListsCreateInternalDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalDomainListsCreateInternalDomainsInternalServerError) GetPayload() *InternalDomainListsCreateInternalDomainsInternalServerErrorBody {
	return o.Payload
}

func (o *InternalDomainListsCreateInternalDomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InternalDomainListsCreateInternalDomainsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
InternalDomainListsCreateInternalDomainsBadRequestBody internal domain lists create internal domains bad request body
swagger:model InternalDomainListsCreateInternalDomainsBadRequestBody
*/
type InternalDomainListsCreateInternalDomainsBadRequestBody struct {

	// error
	Error *InternalDomainListsCreateInternalDomainsBadRequestBodyError `json:"error,omitempty"`
}

// Validate validates this internal domain lists create internal domains bad request body
func (o *InternalDomainListsCreateInternalDomainsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsCreateInternalDomainsBadRequestBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsCreateInternalDomainsBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsCreateInternalDomainsBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this internal domain lists create internal domains bad request body based on the context it is used
func (o *InternalDomainListsCreateInternalDomainsBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsCreateInternalDomainsBadRequestBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsCreateInternalDomainsBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsCreateInternalDomainsBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsCreateInternalDomainsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsCreateInternalDomainsBadRequestBodyError internal domain lists create internal domains bad request body error
swagger:model InternalDomainListsCreateInternalDomainsBadRequestBodyError
*/
type InternalDomainListsCreateInternalDomainsBadRequestBodyError struct {

	// code
	// Example: INVALID_ARGUMENT
	Code string `json:"code,omitempty"`

	// message
	// Example: internal_domains' must not be empty
	Message string `json:"message,omitempty"`

	// status
	// Example: 400
	Status string `json:"status,omitempty"`
}

// Validate validates this internal domain lists create internal domains bad request body error
func (o *InternalDomainListsCreateInternalDomainsBadRequestBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this internal domain lists create internal domains bad request body error based on context it is used
func (o *InternalDomainListsCreateInternalDomainsBadRequestBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsBadRequestBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsBadRequestBodyError) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsCreateInternalDomainsBadRequestBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsCreateInternalDomainsConflictBody internal domain lists create internal domains conflict body
swagger:model InternalDomainListsCreateInternalDomainsConflictBody
*/
type InternalDomainListsCreateInternalDomainsConflictBody struct {

	// error
	Error *InternalDomainListsCreateInternalDomainsConflictBodyError `json:"error,omitempty"`
}

// Validate validates this internal domain lists create internal domains conflict body
func (o *InternalDomainListsCreateInternalDomainsConflictBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsCreateInternalDomainsConflictBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsCreateInternalDomainsConflict" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsCreateInternalDomainsConflict" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this internal domain lists create internal domains conflict body based on the context it is used
func (o *InternalDomainListsCreateInternalDomainsConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsCreateInternalDomainsConflictBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsCreateInternalDomainsConflict" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsCreateInternalDomainsConflict" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsConflictBody) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsCreateInternalDomainsConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsCreateInternalDomainsConflictBodyError internal domain lists create internal domains conflict body error
swagger:model InternalDomainListsCreateInternalDomainsConflictBodyError
*/
type InternalDomainListsCreateInternalDomainsConflictBodyError struct {

	// code
	// Example: ALREADY_EXISTS
	Code string `json:"code,omitempty"`

	// message
	// Example: Cannot use duplicate name \u003cname\u003e
	Message string `json:"message,omitempty"`

	// status
	// Example: 409
	Status string `json:"status,omitempty"`
}

// Validate validates this internal domain lists create internal domains conflict body error
func (o *InternalDomainListsCreateInternalDomainsConflictBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this internal domain lists create internal domains conflict body error based on context it is used
func (o *InternalDomainListsCreateInternalDomainsConflictBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsConflictBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsConflictBodyError) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsCreateInternalDomainsConflictBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsCreateInternalDomainsInternalServerErrorBody internal domain lists create internal domains internal server error body
swagger:model InternalDomainListsCreateInternalDomainsInternalServerErrorBody
*/
type InternalDomainListsCreateInternalDomainsInternalServerErrorBody struct {

	// error
	Error *InternalDomainListsCreateInternalDomainsInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this internal domain lists create internal domains internal server error body
func (o *InternalDomainListsCreateInternalDomainsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsCreateInternalDomainsInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsCreateInternalDomainsInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsCreateInternalDomainsInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this internal domain lists create internal domains internal server error body based on the context it is used
func (o *InternalDomainListsCreateInternalDomainsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InternalDomainListsCreateInternalDomainsInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internalDomainListsCreateInternalDomainsInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("internalDomainListsCreateInternalDomainsInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsCreateInternalDomainsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
InternalDomainListsCreateInternalDomainsInternalServerErrorBodyError internal domain lists create internal domains internal server error body error
swagger:model InternalDomainListsCreateInternalDomainsInternalServerErrorBodyError
*/
type InternalDomainListsCreateInternalDomainsInternalServerErrorBodyError struct {

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

// Validate validates this internal domain lists create internal domains internal server error body error
func (o *InternalDomainListsCreateInternalDomainsInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this internal domain lists create internal domains internal server error body error based on context it is used
func (o *InternalDomainListsCreateInternalDomainsInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InternalDomainListsCreateInternalDomainsInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res InternalDomainListsCreateInternalDomainsInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}