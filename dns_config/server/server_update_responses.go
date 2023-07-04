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
	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// ServerUpdateReader is a Reader for the ServerUpdate structure.
type ServerUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServerUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewServerUpdateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewServerUpdateCreated creates a ServerUpdateCreated with default headers values
func NewServerUpdateCreated() *ServerUpdateCreated {
	return &ServerUpdateCreated{}
}

/*
ServerUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type ServerUpdateCreated struct {
	Payload *models.ConfigUpdateServerResponse
}

func (o *ServerUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /dns/server/{id}][%d] serverUpdateCreated  %+v", 201, o.Payload)
}
func (o *ServerUpdateCreated) GetPayload() *models.ConfigUpdateServerResponse {
	return o.Payload
}

func (o *ServerUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigUpdateServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
