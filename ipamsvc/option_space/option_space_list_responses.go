// Code generated by go-swagger; DO NOT EDIT.

package option_space

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

// OptionSpaceListReader is a Reader for the OptionSpaceList structure.
type OptionSpaceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OptionSpaceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewOptionSpaceListOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewOptionSpaceListOK creates a OptionSpaceListOK with default headers values
func NewOptionSpaceListOK() *OptionSpaceListOK {
	return &OptionSpaceListOK{}
}

/* OptionSpaceListOK describes a response with status code 200, with default header values.

GET operation response
*/
type OptionSpaceListOK struct {
	Payload *models.IpamsvcListOptionSpaceResponse
}

func (o *OptionSpaceListOK) Error() string {
	return fmt.Sprintf("[GET /dhcp/option_space][%d] optionSpaceListOK  %+v", 200, o.Payload)
}
func (o *OptionSpaceListOK) GetPayload() *models.IpamsvcListOptionSpaceResponse {
	return o.Payload
}

func (o *OptionSpaceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcListOptionSpaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
