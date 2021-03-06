// Code generated by go-swagger; DO NOT EDIT.

package forward_nsg

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

// ForwardNsgCreateReader is a Reader for the ForwardNsgCreate structure.
type ForwardNsgCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForwardNsgCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewForwardNsgCreateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewForwardNsgCreateCreated creates a ForwardNsgCreateCreated with default headers values
func NewForwardNsgCreateCreated() *ForwardNsgCreateCreated {
	return &ForwardNsgCreateCreated{}
}

/* ForwardNsgCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type ForwardNsgCreateCreated struct {
	Payload *models.ConfigCreateForwardNSGResponse
}

func (o *ForwardNsgCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dns/forward_nsg][%d] forwardNsgCreateCreated  %+v", 201, o.Payload)
}
func (o *ForwardNsgCreateCreated) GetPayload() *models.ConfigCreateForwardNSGResponse {
	return o.Payload
}

func (o *ForwardNsgCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCreateForwardNSGResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
