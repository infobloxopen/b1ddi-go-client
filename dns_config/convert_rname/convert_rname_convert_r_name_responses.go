// Code generated by go-swagger; DO NOT EDIT.

package convert_rname

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

// ConvertRnameConvertRNameReader is a Reader for the ConvertRnameConvertRName structure.
type ConvertRnameConvertRNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConvertRnameConvertRNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewConvertRnameConvertRNameOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewConvertRnameConvertRNameOK creates a ConvertRnameConvertRNameOK with default headers values
func NewConvertRnameConvertRNameOK() *ConvertRnameConvertRNameOK {
	return &ConvertRnameConvertRNameOK{}
}

/* ConvertRnameConvertRNameOK describes a response with status code 200, with default header values.

GET operation response
*/
type ConvertRnameConvertRNameOK struct {
	Payload *models.ConfigConvertRNameResponse
}

func (o *ConvertRnameConvertRNameOK) Error() string {
	return fmt.Sprintf("[GET /dns/convert_rname/{email_address}][%d] convertRnameConvertRNameOK  %+v", 200, o.Payload)
}
func (o *ConvertRnameConvertRNameOK) GetPayload() *models.ConfigConvertRNameResponse {
	return o.Payload
}

func (o *ConvertRnameConvertRNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigConvertRNameResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
