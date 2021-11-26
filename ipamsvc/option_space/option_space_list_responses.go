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
)

// OptionSpaceListReader is a Reader for the OptionSpaceList structure.
type OptionSpaceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OptionSpaceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOptionSpaceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
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
