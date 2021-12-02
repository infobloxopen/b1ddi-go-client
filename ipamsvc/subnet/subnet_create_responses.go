// Code generated by go-swagger; DO NOT EDIT.

package subnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// SubnetCreateReader is a Reader for the SubnetCreate structure.
type SubnetCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubnetCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewSubnetCreateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewSubnetCreateCreated creates a SubnetCreateCreated with default headers values
func NewSubnetCreateCreated() *SubnetCreateCreated {
	return &SubnetCreateCreated{}
}

/* SubnetCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type SubnetCreateCreated struct {
	Payload *models.IpamsvcCreateSubnetResponse
}

func (o *SubnetCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/subnet][%d] subnetCreateCreated  %+v", 201, o.Payload)
}
func (o *SubnetCreateCreated) GetPayload() *models.IpamsvcCreateSubnetResponse {
	return o.Payload
}

func (o *SubnetCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcCreateSubnetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}