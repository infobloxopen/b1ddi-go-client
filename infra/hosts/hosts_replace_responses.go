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

// HostsReplaceReader is a Reader for the HostsReplace structure.
type HostsReplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsReplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewHostsReplaceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /hosts/{from.resource_id}/replace/{to.resource_id}] HostsReplace", response, response.Code())
	}
}

// NewHostsReplaceCreated creates a HostsReplaceCreated with default headers values
func NewHostsReplaceCreated() *HostsReplaceCreated {
	return &HostsReplaceCreated{}
}

/*
HostsReplaceCreated describes a response with status code 201, with default header values.

POST operation response
*/
type HostsReplaceCreated struct {
	Payload models.InfraReplaceHostResponse
}

// IsSuccess returns true when this hosts replace created response has a 2xx status code
func (o *HostsReplaceCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hosts replace created response has a 3xx status code
func (o *HostsReplaceCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hosts replace created response has a 4xx status code
func (o *HostsReplaceCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this hosts replace created response has a 5xx status code
func (o *HostsReplaceCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this hosts replace created response a status code equal to that given
func (o *HostsReplaceCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the hosts replace created response
func (o *HostsReplaceCreated) Code() int {
	return 201
}

func (o *HostsReplaceCreated) Error() string {
	return fmt.Sprintf("[POST /hosts/{from.resource_id}/replace/{to.resource_id}][%d] hostsReplaceCreated  %+v", 201, o.Payload)
}

func (o *HostsReplaceCreated) String() string {
	return fmt.Sprintf("[POST /hosts/{from.resource_id}/replace/{to.resource_id}][%d] hostsReplaceCreated  %+v", 201, o.Payload)
}

func (o *HostsReplaceCreated) GetPayload() models.InfraReplaceHostResponse {
	return o.Payload
}

func (o *HostsReplaceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
