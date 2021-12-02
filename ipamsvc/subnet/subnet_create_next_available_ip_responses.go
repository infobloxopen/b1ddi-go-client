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

// SubnetCreateNextAvailableIPReader is a Reader for the SubnetCreateNextAvailableIP structure.
type SubnetCreateNextAvailableIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubnetCreateNextAvailableIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewSubnetCreateNextAvailableIPCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewSubnetCreateNextAvailableIPCreated creates a SubnetCreateNextAvailableIPCreated with default headers values
func NewSubnetCreateNextAvailableIPCreated() *SubnetCreateNextAvailableIPCreated {
	return &SubnetCreateNextAvailableIPCreated{}
}

/* SubnetCreateNextAvailableIPCreated describes a response with status code 201, with default header values.

POST operation response
*/
type SubnetCreateNextAvailableIPCreated struct {
	Payload *models.IpamsvcCreateNextAvailableIPResponse
}

func (o *SubnetCreateNextAvailableIPCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/subnet/{id}/nextavailableip][%d] subnetCreateNextAvailableIpCreated  %+v", 201, o.Payload)
}
func (o *SubnetCreateNextAvailableIPCreated) GetPayload() *models.IpamsvcCreateNextAvailableIPResponse {
	return o.Payload
}

func (o *SubnetCreateNextAvailableIPCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcCreateNextAvailableIPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}