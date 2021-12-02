// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ServerCreateReader is a Reader for the ServerCreate structure.
type ServerCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServerCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewServerCreateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewServerCreateCreated creates a ServerCreateCreated with default headers values
func NewServerCreateCreated() *ServerCreateCreated {
	return &ServerCreateCreated{}
}

/* ServerCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type ServerCreateCreated struct {
	Payload *models.ConfigCreateServerResponse
}

func (o *ServerCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dns/server][%d] serverCreateCreated  %+v", 201, o.Payload)
}
func (o *ServerCreateCreated) GetPayload() *models.ConfigCreateServerResponse {
	return o.Payload
}

func (o *ServerCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCreateServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
