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

// AddressBlockReadReader is a Reader for the AddressBlockRead structure.
type AddressBlockReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressBlockReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewAddressBlockReadOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewAddressBlockReadOK creates a AddressBlockReadOK with default headers values
func NewAddressBlockReadOK() *AddressBlockReadOK {
	return &AddressBlockReadOK{}
}

/* AddressBlockReadOK describes a response with status code 200, with default header values.

GET operation response
*/
type AddressBlockReadOK struct {
	Payload *models.IpamsvcReadAddressBlockResponse
}

func (o *AddressBlockReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/address_block/{id}][%d] addressBlockReadOK  %+v", 200, o.Payload)
}
func (o *AddressBlockReadOK) GetPayload() *models.IpamsvcReadAddressBlockResponse {
	return o.Payload
}

func (o *AddressBlockReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcReadAddressBlockResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
