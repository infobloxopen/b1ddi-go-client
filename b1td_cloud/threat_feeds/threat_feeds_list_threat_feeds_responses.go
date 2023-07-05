// Code generated by go-swagger; DO NOT EDIT.

package threat_feeds

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

// ThreatFeedsListThreatFeedsReader is a Reader for the ThreatFeedsListThreatFeeds structure.
type ThreatFeedsListThreatFeedsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThreatFeedsListThreatFeedsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThreatFeedsListThreatFeedsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewThreatFeedsListThreatFeedsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /threat_feeds] threat_feedsListThreatFeeds", response, response.Code())
	}
}

// NewThreatFeedsListThreatFeedsOK creates a ThreatFeedsListThreatFeedsOK with default headers values
func NewThreatFeedsListThreatFeedsOK() *ThreatFeedsListThreatFeedsOK {
	return &ThreatFeedsListThreatFeedsOK{}
}

/*
ThreatFeedsListThreatFeedsOK describes a response with status code 200, with default header values.

GET operation response
*/
type ThreatFeedsListThreatFeedsOK struct {
	Payload *models.AtcfwThreatFeedMultiResponse
}

// IsSuccess returns true when this threat feeds list threat feeds o k response has a 2xx status code
func (o *ThreatFeedsListThreatFeedsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this threat feeds list threat feeds o k response has a 3xx status code
func (o *ThreatFeedsListThreatFeedsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this threat feeds list threat feeds o k response has a 4xx status code
func (o *ThreatFeedsListThreatFeedsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this threat feeds list threat feeds o k response has a 5xx status code
func (o *ThreatFeedsListThreatFeedsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this threat feeds list threat feeds o k response a status code equal to that given
func (o *ThreatFeedsListThreatFeedsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the threat feeds list threat feeds o k response
func (o *ThreatFeedsListThreatFeedsOK) Code() int {
	return 200
}

func (o *ThreatFeedsListThreatFeedsOK) Error() string {
	return fmt.Sprintf("[GET /threat_feeds][%d] threatFeedsListThreatFeedsOK  %+v", 200, o.Payload)
}

func (o *ThreatFeedsListThreatFeedsOK) String() string {
	return fmt.Sprintf("[GET /threat_feeds][%d] threatFeedsListThreatFeedsOK  %+v", 200, o.Payload)
}

func (o *ThreatFeedsListThreatFeedsOK) GetPayload() *models.AtcfwThreatFeedMultiResponse {
	return o.Payload
}

func (o *ThreatFeedsListThreatFeedsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AtcfwThreatFeedMultiResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThreatFeedsListThreatFeedsInternalServerError creates a ThreatFeedsListThreatFeedsInternalServerError with default headers values
func NewThreatFeedsListThreatFeedsInternalServerError() *ThreatFeedsListThreatFeedsInternalServerError {
	return &ThreatFeedsListThreatFeedsInternalServerError{}
}

/*
	ThreatFeedsListThreatFeedsInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type ThreatFeedsListThreatFeedsInternalServerError struct {
	Payload *ThreatFeedsListThreatFeedsInternalServerErrorBody
}

// IsSuccess returns true when this threat feeds list threat feeds internal server error response has a 2xx status code
func (o *ThreatFeedsListThreatFeedsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this threat feeds list threat feeds internal server error response has a 3xx status code
func (o *ThreatFeedsListThreatFeedsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this threat feeds list threat feeds internal server error response has a 4xx status code
func (o *ThreatFeedsListThreatFeedsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this threat feeds list threat feeds internal server error response has a 5xx status code
func (o *ThreatFeedsListThreatFeedsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this threat feeds list threat feeds internal server error response a status code equal to that given
func (o *ThreatFeedsListThreatFeedsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the threat feeds list threat feeds internal server error response
func (o *ThreatFeedsListThreatFeedsInternalServerError) Code() int {
	return 500
}

func (o *ThreatFeedsListThreatFeedsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /threat_feeds][%d] threatFeedsListThreatFeedsInternalServerError  %+v", 500, o.Payload)
}

func (o *ThreatFeedsListThreatFeedsInternalServerError) String() string {
	return fmt.Sprintf("[GET /threat_feeds][%d] threatFeedsListThreatFeedsInternalServerError  %+v", 500, o.Payload)
}

func (o *ThreatFeedsListThreatFeedsInternalServerError) GetPayload() *ThreatFeedsListThreatFeedsInternalServerErrorBody {
	return o.Payload
}

func (o *ThreatFeedsListThreatFeedsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ThreatFeedsListThreatFeedsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ThreatFeedsListThreatFeedsInternalServerErrorBody threat feeds list threat feeds internal server error body
swagger:model ThreatFeedsListThreatFeedsInternalServerErrorBody
*/
type ThreatFeedsListThreatFeedsInternalServerErrorBody struct {

	// error
	Error *ThreatFeedsListThreatFeedsInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this threat feeds list threat feeds internal server error body
func (o *ThreatFeedsListThreatFeedsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ThreatFeedsListThreatFeedsInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("threatFeedsListThreatFeedsInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("threatFeedsListThreatFeedsInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this threat feeds list threat feeds internal server error body based on the context it is used
func (o *ThreatFeedsListThreatFeedsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ThreatFeedsListThreatFeedsInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("threatFeedsListThreatFeedsInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("threatFeedsListThreatFeedsInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ThreatFeedsListThreatFeedsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ThreatFeedsListThreatFeedsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ThreatFeedsListThreatFeedsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ThreatFeedsListThreatFeedsInternalServerErrorBodyError threat feeds list threat feeds internal server error body error
swagger:model ThreatFeedsListThreatFeedsInternalServerErrorBodyError
*/
type ThreatFeedsListThreatFeedsInternalServerErrorBodyError struct {

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

// Validate validates this threat feeds list threat feeds internal server error body error
func (o *ThreatFeedsListThreatFeedsInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this threat feeds list threat feeds internal server error body error based on context it is used
func (o *ThreatFeedsListThreatFeedsInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ThreatFeedsListThreatFeedsInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ThreatFeedsListThreatFeedsInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res ThreatFeedsListThreatFeedsInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}