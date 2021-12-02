// Code generated by go-swagger; DO NOT EDIT.

package address_block

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// AddressBlockCreateReader is a Reader for the AddressBlockCreate structure.
type AddressBlockCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressBlockCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewAddressBlockCreateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewAddressBlockCreateCreated creates a AddressBlockCreateCreated with default headers values
func NewAddressBlockCreateCreated() *AddressBlockCreateCreated {
	return &AddressBlockCreateCreated{}
}

/* AddressBlockCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type AddressBlockCreateCreated struct {
	Payload *models.IpamsvcCreateAddressBlockResponse
}

func (o *AddressBlockCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/address_block][%d] addressBlockCreateCreated  %+v", 201, o.Payload)
}
func (o *AddressBlockCreateCreated) GetPayload() *models.IpamsvcCreateAddressBlockResponse {
	return o.Payload
}

func (o *AddressBlockCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcCreateAddressBlockResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
