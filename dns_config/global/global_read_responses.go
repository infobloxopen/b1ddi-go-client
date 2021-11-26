// Code generated by go-swagger; DO NOT EDIT.

package global

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// GlobalReadReader is a Reader for the GlobalRead structure.
type GlobalReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GlobalReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGlobalReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGlobalReadOK creates a GlobalReadOK with default headers values
func NewGlobalReadOK() *GlobalReadOK {
	return &GlobalReadOK{}
}

/* GlobalReadOK describes a response with status code 200, with default header values.

GET operation response
*/
type GlobalReadOK struct {
	Payload *models.ConfigReadGlobalResponse
}

func (o *GlobalReadOK) Error() string {
	return fmt.Sprintf("[GET /dns/global][%d] globalReadOK  %+v", 200, o.Payload)
}
func (o *GlobalReadOK) GetPayload() *models.ConfigReadGlobalResponse {
	return o.Payload
}

func (o *GlobalReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigReadGlobalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
