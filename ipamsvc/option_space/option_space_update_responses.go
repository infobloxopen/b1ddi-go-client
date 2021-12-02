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

// OptionSpaceUpdateReader is a Reader for the OptionSpaceUpdate structure.
type OptionSpaceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OptionSpaceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewOptionSpaceUpdateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewOptionSpaceUpdateCreated creates a OptionSpaceUpdateCreated with default headers values
func NewOptionSpaceUpdateCreated() *OptionSpaceUpdateCreated {
	return &OptionSpaceUpdateCreated{}
}

/* OptionSpaceUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type OptionSpaceUpdateCreated struct {
	Payload *models.IpamsvcUpdateOptionSpaceResponse
}

func (o *OptionSpaceUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /dhcp/option_space/{id}][%d] optionSpaceUpdateCreated  %+v", 201, o.Payload)
}
func (o *OptionSpaceUpdateCreated) GetPayload() *models.IpamsvcUpdateOptionSpaceResponse {
	return o.Payload
}

func (o *OptionSpaceUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcUpdateOptionSpaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
