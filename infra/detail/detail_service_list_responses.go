// Code generated by go-swagger; DO NOT EDIT.

package detail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// DetailServiceListReader is a Reader for the DetailServiceList structure.
type DetailServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /detail_services] DetailServiceList", response, response.Code())
	}
}

// NewDetailServiceListOK creates a DetailServiceListOK with default headers values
func NewDetailServiceListOK() *DetailServiceListOK {
	return &DetailServiceListOK{}
}

/*
DetailServiceListOK describes a response with status code 200, with default header values.

GET operation response
*/
type DetailServiceListOK struct {
	Payload *models.InfraListDetailServicesResponse
}

// IsSuccess returns true when this detail service list o k response has a 2xx status code
func (o *DetailServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this detail service list o k response has a 3xx status code
func (o *DetailServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detail service list o k response has a 4xx status code
func (o *DetailServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this detail service list o k response has a 5xx status code
func (o *DetailServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this detail service list o k response a status code equal to that given
func (o *DetailServiceListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the detail service list o k response
func (o *DetailServiceListOK) Code() int {
	return 200
}

func (o *DetailServiceListOK) Error() string {
	return fmt.Sprintf("[GET /detail_services][%d] detailServiceListOK  %+v", 200, o.Payload)
}

func (o *DetailServiceListOK) String() string {
	return fmt.Sprintf("[GET /detail_services][%d] detailServiceListOK  %+v", 200, o.Payload)
}

func (o *DetailServiceListOK) GetPayload() *models.InfraListDetailServicesResponse {
	return o.Payload
}

func (o *DetailServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraListDetailServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}