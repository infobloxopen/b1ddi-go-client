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
)

// ForwardNsgListReader is a Reader for the ForwardNsgList structure.
type ForwardNsgListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForwardNsgListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewForwardNsgListOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewForwardNsgListOK creates a ForwardNsgListOK with default headers values
func NewForwardNsgListOK() *ForwardNsgListOK {
	return &ForwardNsgListOK{}
}

/* ForwardNsgListOK describes a response with status code 200, with default header values.

GET operation response
*/
type ForwardNsgListOK struct {
	Payload *models.ConfigListForwardNSGResponse
}

func (o *ForwardNsgListOK) Error() string {
	return fmt.Sprintf("[GET /dns/forward_nsg][%d] forwardNsgListOK  %+v", 200, o.Payload)
}
func (o *ForwardNsgListOK) GetPayload() *models.ConfigListForwardNSGResponse {
	return o.Payload
}

func (o *ForwardNsgListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigListForwardNSGResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}