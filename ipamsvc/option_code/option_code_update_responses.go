// Code generated by go-swagger; DO NOT EDIT.

package option_code

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// OptionCodeUpdateReader is a Reader for the OptionCodeUpdate structure.
type OptionCodeUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OptionCodeUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewOptionCodeUpdateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewOptionCodeUpdateCreated creates a OptionCodeUpdateCreated with default headers values
func NewOptionCodeUpdateCreated() *OptionCodeUpdateCreated {
	return &OptionCodeUpdateCreated{}
}

/* OptionCodeUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type OptionCodeUpdateCreated struct {
	Payload *models.IpamsvcUpdateOptionCodeResponse
}

func (o *OptionCodeUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /dhcp/option_code/{id}][%d] optionCodeUpdateCreated  %+v", 201, o.Payload)
}
func (o *OptionCodeUpdateCreated) GetPayload() *models.IpamsvcUpdateOptionCodeResponse {
	return o.Payload
}

func (o *OptionCodeUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcUpdateOptionCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}