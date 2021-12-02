// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// HostListReader is a Reader for the HostList structure.
type HostListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewHostListOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewHostListOK creates a HostListOK with default headers values
func NewHostListOK() *HostListOK {
	return &HostListOK{}
}

/* HostListOK describes a response with status code 200, with default header values.

GET operation response
*/
type HostListOK struct {
	Payload *models.ConfigListHostResponse
}

func (o *HostListOK) Error() string {
	return fmt.Sprintf("[GET /dns/host][%d] hostListOK  %+v", 200, o.Payload)
}
func (o *HostListOK) GetPayload() *models.ConfigListHostResponse {
	return o.Payload
}

func (o *HostListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigListHostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
