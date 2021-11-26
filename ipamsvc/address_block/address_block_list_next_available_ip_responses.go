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

// AddressBlockListNextAvailableIPReader is a Reader for the AddressBlockListNextAvailableIP structure.
type AddressBlockListNextAvailableIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressBlockListNextAvailableIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddressBlockListNextAvailableIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddressBlockListNextAvailableIPOK creates a AddressBlockListNextAvailableIPOK with default headers values
func NewAddressBlockListNextAvailableIPOK() *AddressBlockListNextAvailableIPOK {
	return &AddressBlockListNextAvailableIPOK{}
}

/* AddressBlockListNextAvailableIPOK describes a response with status code 200, with default header values.

GET operation response
*/
type AddressBlockListNextAvailableIPOK struct {
	Payload *models.IpamsvcNextAvailableIPResponse
}

func (o *AddressBlockListNextAvailableIPOK) Error() string {
	return fmt.Sprintf("[GET /ipam/address_block/{id}/nextavailableip][%d] addressBlockListNextAvailableIpOK  %+v", 200, o.Payload)
}
func (o *AddressBlockListNextAvailableIPOK) GetPayload() *models.IpamsvcNextAvailableIPResponse {
	return o.Payload
}

func (o *AddressBlockListNextAvailableIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcNextAvailableIPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
