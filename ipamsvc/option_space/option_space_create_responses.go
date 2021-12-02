// Code generated by go-swagger; DO NOT EDIT.

package option_space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// OptionSpaceCreateReader is a Reader for the OptionSpaceCreate structure.
type OptionSpaceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OptionSpaceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewOptionSpaceCreateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewOptionSpaceCreateCreated creates a OptionSpaceCreateCreated with default headers values
func NewOptionSpaceCreateCreated() *OptionSpaceCreateCreated {
	return &OptionSpaceCreateCreated{}
}

/* OptionSpaceCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type OptionSpaceCreateCreated struct {
	Payload *models.IpamsvcCreateOptionSpaceResponse
}

func (o *OptionSpaceCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dhcp/option_space][%d] optionSpaceCreateCreated  %+v", 201, o.Payload)
}
func (o *OptionSpaceCreateCreated) GetPayload() *models.IpamsvcCreateOptionSpaceResponse {
	return o.Payload
}

func (o *OptionSpaceCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcCreateOptionSpaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}