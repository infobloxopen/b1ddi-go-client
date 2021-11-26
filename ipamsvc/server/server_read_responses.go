// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ServerReadReader is a Reader for the ServerRead structure.
type ServerReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServerReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServerReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServerReadOK creates a ServerReadOK with default headers values
func NewServerReadOK() *ServerReadOK {
	return &ServerReadOK{}
}

/* ServerReadOK describes a response with status code 200, with default header values.

GET operation response
*/
type ServerReadOK struct {
	Payload *models.IpamsvcReadServerResponse
}

func (o *ServerReadOK) Error() string {
	return fmt.Sprintf("[GET /dhcp/server/{id}][%d] serverReadOK  %+v", 200, o.Payload)
}
func (o *ServerReadOK) GetPayload() *models.IpamsvcReadServerResponse {
	return o.Payload
}

func (o *ServerReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcReadServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
