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

// HardwareFilterUpdateReader is a Reader for the HardwareFilterUpdate structure.
type HardwareFilterUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareFilterUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewHardwareFilterUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHardwareFilterUpdateCreated creates a HardwareFilterUpdateCreated with default headers values
func NewHardwareFilterUpdateCreated() *HardwareFilterUpdateCreated {
	return &HardwareFilterUpdateCreated{}
}

/* HardwareFilterUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type HardwareFilterUpdateCreated struct {
	Payload *models.IpamsvcUpdateHardwareFilterResponse
}

func (o *HardwareFilterUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /dhcp/hardware_filter/{id}][%d] hardwareFilterUpdateCreated  %+v", 201, o.Payload)
}
func (o *HardwareFilterUpdateCreated) GetPayload() *models.IpamsvcUpdateHardwareFilterResponse {
	return o.Payload
}

func (o *HardwareFilterUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcUpdateHardwareFilterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
