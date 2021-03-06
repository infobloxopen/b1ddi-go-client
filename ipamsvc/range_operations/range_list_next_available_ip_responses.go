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
	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// RangeListNextAvailableIPReader is a Reader for the RangeListNextAvailableIP structure.
type RangeListNextAvailableIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RangeListNextAvailableIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewRangeListNextAvailableIPOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewRangeListNextAvailableIPOK creates a RangeListNextAvailableIPOK with default headers values
func NewRangeListNextAvailableIPOK() *RangeListNextAvailableIPOK {
	return &RangeListNextAvailableIPOK{}
}

/* RangeListNextAvailableIPOK describes a response with status code 200, with default header values.

GET operation response
*/
type RangeListNextAvailableIPOK struct {
	Payload *models.IpamsvcNextAvailableIPResponse
}

func (o *RangeListNextAvailableIPOK) Error() string {
	return fmt.Sprintf("[GET /ipam/range/{id}/nextavailableip][%d] rangeListNextAvailableIpOK  %+v", 200, o.Payload)
}
func (o *RangeListNextAvailableIPOK) GetPayload() *models.IpamsvcNextAvailableIPResponse {
	return o.Payload
}

func (o *RangeListNextAvailableIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcNextAvailableIPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
