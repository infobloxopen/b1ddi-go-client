// Code generated by go-swagger; DO NOT EDIT.

package hardware_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// HardwareFilterListReader is a Reader for the HardwareFilterList structure.
type HardwareFilterListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareFilterListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewHardwareFilterListOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewHardwareFilterListOK creates a HardwareFilterListOK with default headers values
func NewHardwareFilterListOK() *HardwareFilterListOK {
	return &HardwareFilterListOK{}
}

/* HardwareFilterListOK describes a response with status code 200, with default header values.

GET operation response
*/
type HardwareFilterListOK struct {
	Payload *models.IpamsvcListHardwareFilterResponse
}

func (o *HardwareFilterListOK) Error() string {
	return fmt.Sprintf("[GET /dhcp/hardware_filter][%d] hardwareFilterListOK  %+v", 200, o.Payload)
}
func (o *HardwareFilterListOK) GetPayload() *models.IpamsvcListHardwareFilterResponse {
	return o.Payload
}

func (o *HardwareFilterListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcListHardwareFilterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}