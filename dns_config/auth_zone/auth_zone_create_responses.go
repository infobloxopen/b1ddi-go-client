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

// AuthZoneCreateReader is a Reader for the AuthZoneCreate structure.
type AuthZoneCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthZoneCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewAuthZoneCreateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewAuthZoneCreateCreated creates a AuthZoneCreateCreated with default headers values
func NewAuthZoneCreateCreated() *AuthZoneCreateCreated {
	return &AuthZoneCreateCreated{}
}

/* AuthZoneCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type AuthZoneCreateCreated struct {
	Payload *models.ConfigCreateAuthZoneResponse
}

func (o *AuthZoneCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dns/auth_zone][%d] authZoneCreateCreated  %+v", 201, o.Payload)
}
func (o *AuthZoneCreateCreated) GetPayload() *models.ConfigCreateAuthZoneResponse {
	return o.Payload
}

func (o *AuthZoneCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCreateAuthZoneResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
