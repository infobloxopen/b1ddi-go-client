// Code generated by go-swagger; DO NOT EDIT.

package fixed_address

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

// FixedAddressListReader is a Reader for the FixedAddressList structure.
type FixedAddressListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FixedAddressListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewFixedAddressListOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewFixedAddressListOK creates a FixedAddressListOK with default headers values
func NewFixedAddressListOK() *FixedAddressListOK {
	return &FixedAddressListOK{}
}

/* FixedAddressListOK describes a response with status code 200, with default header values.

GET operation response
*/
type FixedAddressListOK struct {
	Payload *models.IpamsvcListFixedAddressResponse
}

func (o *FixedAddressListOK) Error() string {
	return fmt.Sprintf("[GET /dhcp/fixed_address][%d] fixedAddressListOK  %+v", 200, o.Payload)
}
func (o *FixedAddressListOK) GetPayload() *models.IpamsvcListFixedAddressResponse {
	return o.Payload
}

func (o *FixedAddressListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcListFixedAddressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
