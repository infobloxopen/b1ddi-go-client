// Code generated by go-swagger; DO NOT EDIT.

package auth_nsg

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// AuthNsgCreateReader is a Reader for the AuthNsgCreate structure.
type AuthNsgCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthNsgCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewAuthNsgCreateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewAuthNsgCreateCreated creates a AuthNsgCreateCreated with default headers values
func NewAuthNsgCreateCreated() *AuthNsgCreateCreated {
	return &AuthNsgCreateCreated{}
}

/* AuthNsgCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type AuthNsgCreateCreated struct {
	Payload *models.ConfigCreateAuthNSGResponse
}

func (o *AuthNsgCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dns/auth_nsg][%d] authNsgCreateCreated  %+v", 201, o.Payload)
}
func (o *AuthNsgCreateCreated) GetPayload() *models.ConfigCreateAuthNSGResponse {
	return o.Payload
}

func (o *AuthNsgCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCreateAuthNSGResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
