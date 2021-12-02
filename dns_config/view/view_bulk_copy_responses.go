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

// ViewBulkCopyReader is a Reader for the ViewBulkCopy structure.
type ViewBulkCopyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ViewBulkCopyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewViewBulkCopyCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewViewBulkCopyCreated creates a ViewBulkCopyCreated with default headers values
func NewViewBulkCopyCreated() *ViewBulkCopyCreated {
	return &ViewBulkCopyCreated{}
}

/* ViewBulkCopyCreated describes a response with status code 201, with default header values.

POST operation response
*/
type ViewBulkCopyCreated struct {
	Payload *models.ConfigBulkCopyViewResponse
}

func (o *ViewBulkCopyCreated) Error() string {
	return fmt.Sprintf("[POST /dns/view/bulk_copy][%d] viewBulkCopyCreated  %+v", 201, o.Payload)
}
func (o *ViewBulkCopyCreated) GetPayload() *models.ConfigBulkCopyViewResponse {
	return o.Payload
}

func (o *ViewBulkCopyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigBulkCopyViewResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
