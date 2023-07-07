// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// HostsCreateReader is a Reader for the HostsCreate structure.
type HostsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewHostsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /hosts] HostsCreate", response, response.Code())
	}
}

// NewHostsCreateCreated creates a HostsCreateCreated with default headers values
func NewHostsCreateCreated() *HostsCreateCreated {
	return &HostsCreateCreated{}
}

/*
HostsCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type HostsCreateCreated struct {
	Payload *models.InfraCreateHostResponse
}

// IsSuccess returns true when this hosts create created response has a 2xx status code
func (o *HostsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hosts create created response has a 3xx status code
func (o *HostsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hosts create created response has a 4xx status code
func (o *HostsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this hosts create created response has a 5xx status code
func (o *HostsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this hosts create created response a status code equal to that given
func (o *HostsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the hosts create created response
func (o *HostsCreateCreated) Code() int {
	return 201
}

func (o *HostsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /hosts][%d] hostsCreateCreated  %+v", 201, o.Payload)
}

func (o *HostsCreateCreated) String() string {
	return fmt.Sprintf("[POST /hosts][%d] hostsCreateCreated  %+v", 201, o.Payload)
}

func (o *HostsCreateCreated) GetPayload() *models.InfraCreateHostResponse {
	return o.Payload
}

func (o *HostsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraCreateHostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
