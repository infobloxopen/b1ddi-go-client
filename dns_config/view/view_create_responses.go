// Code generated by go-swagger; DO NOT EDIT.

package view

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ViewCreateReader is a Reader for the ViewCreate structure.
type ViewCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ViewCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewViewCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewViewCreateCreated creates a ViewCreateCreated with default headers values
func NewViewCreateCreated() *ViewCreateCreated {
	return &ViewCreateCreated{}
}

/* ViewCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type ViewCreateCreated struct {
	Payload *models.ConfigCreateViewResponse
}

func (o *ViewCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dns/view][%d] viewCreateCreated  %+v", 201, o.Payload)
}
func (o *ViewCreateCreated) GetPayload() *models.ConfigCreateViewResponse {
	return o.Payload
}

func (o *ViewCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCreateViewResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
