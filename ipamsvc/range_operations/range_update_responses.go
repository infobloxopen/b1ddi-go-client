// Code generated by go-swagger; DO NOT EDIT.

package range_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// RangeUpdateReader is a Reader for the RangeUpdate structure.
type RangeUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RangeUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewRangeUpdateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewRangeUpdateCreated creates a RangeUpdateCreated with default headers values
func NewRangeUpdateCreated() *RangeUpdateCreated {
	return &RangeUpdateCreated{}
}

/* RangeUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type RangeUpdateCreated struct {
	Payload *models.IpamsvcUpdateRangeResponse
}

func (o *RangeUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /ipam/range/{id}][%d] rangeUpdateCreated  %+v", 201, o.Payload)
}
func (o *RangeUpdateCreated) GetPayload() *models.IpamsvcUpdateRangeResponse {
	return o.Payload
}

func (o *RangeUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcUpdateRangeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
