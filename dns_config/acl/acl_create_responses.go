// Code generated by go-swagger; DO NOT EDIT.

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ACLCreateReader is a Reader for the ACLCreate structure.
type ACLCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ACLCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewACLCreateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewACLCreateCreated creates a ACLCreateCreated with default headers values
func NewACLCreateCreated() *ACLCreateCreated {
	return &ACLCreateCreated{}
}

/* ACLCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type ACLCreateCreated struct {
	Payload *models.ConfigCreateACLResponse
}

func (o *ACLCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dns/acl][%d] aclCreateCreated  %+v", 201, o.Payload)
}
func (o *ACLCreateCreated) GetPayload() *models.ConfigCreateACLResponse {
	return o.Payload
}

func (o *ACLCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCreateACLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
