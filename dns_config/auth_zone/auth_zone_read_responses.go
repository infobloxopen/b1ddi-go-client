// Code generated by go-swagger; DO NOT EDIT.

package auth_zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// AuthZoneReadReader is a Reader for the AuthZoneRead structure.
type AuthZoneReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthZoneReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewAuthZoneReadOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewAuthZoneReadOK creates a AuthZoneReadOK with default headers values
func NewAuthZoneReadOK() *AuthZoneReadOK {
	return &AuthZoneReadOK{}
}

/* AuthZoneReadOK describes a response with status code 200, with default header values.

GET operation response
*/
type AuthZoneReadOK struct {
	Payload *models.ConfigReadAuthZoneResponse
}

func (o *AuthZoneReadOK) Error() string {
	return fmt.Sprintf("[GET /dns/auth_zone/{id}][%d] authZoneReadOK  %+v", 200, o.Payload)
}
func (o *AuthZoneReadOK) GetPayload() *models.ConfigReadAuthZoneResponse {
	return o.Payload
}

func (o *AuthZoneReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigReadAuthZoneResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
