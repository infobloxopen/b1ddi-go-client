// Code generated by go-swagger; DO NOT EDIT.

package asm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// AsmCreateReader is a Reader for the AsmCreate structure.
type AsmCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AsmCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewAsmCreateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewAsmCreateCreated creates a AsmCreateCreated with default headers values
func NewAsmCreateCreated() *AsmCreateCreated {
	return &AsmCreateCreated{}
}

/* AsmCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type AsmCreateCreated struct {
	Payload *models.IpamsvcCreateASMResponse
}

func (o *AsmCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/asm][%d] asmCreateCreated  %+v", 201, o.Payload)
}
func (o *AsmCreateCreated) GetPayload() *models.IpamsvcCreateASMResponse {
	return o.Payload
}

func (o *AsmCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcCreateASMResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
