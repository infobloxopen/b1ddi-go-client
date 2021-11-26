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

// ForwardNsgUpdateReader is a Reader for the ForwardNsgUpdate structure.
type ForwardNsgUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForwardNsgUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewForwardNsgUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewForwardNsgUpdateCreated creates a ForwardNsgUpdateCreated with default headers values
func NewForwardNsgUpdateCreated() *ForwardNsgUpdateCreated {
	return &ForwardNsgUpdateCreated{}
}

/* ForwardNsgUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type ForwardNsgUpdateCreated struct {
	Payload *models.ConfigUpdateForwardNSGResponse
}

func (o *ForwardNsgUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /dns/forward_nsg/{id}][%d] forwardNsgUpdateCreated  %+v", 201, o.Payload)
}
func (o *ForwardNsgUpdateCreated) GetPayload() *models.ConfigUpdateForwardNSGResponse {
	return o.Payload
}

func (o *ForwardNsgUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigUpdateForwardNSGResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
