// Code generated by go-swagger; DO NOT EDIT.

package hardware_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HardwareFilterDeleteReader is a Reader for the HardwareFilterDelete structure.
type HardwareFilterDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareFilterDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewHardwareFilterDeleteNoContent()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewHardwareFilterDeleteNoContent creates a HardwareFilterDeleteNoContent with default headers values
func NewHardwareFilterDeleteNoContent() *HardwareFilterDeleteNoContent {
	return &HardwareFilterDeleteNoContent{}
}

/* HardwareFilterDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type HardwareFilterDeleteNoContent struct {
}

func (o *HardwareFilterDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dhcp/hardware_filter/{id}][%d] hardwareFilterDeleteNoContent ", 204)
}

func (o *HardwareFilterDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
