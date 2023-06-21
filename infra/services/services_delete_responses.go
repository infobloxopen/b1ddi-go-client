// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ServicesDeleteReader is a Reader for the ServicesDelete structure.
type ServicesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewServicesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /services/{id}] ServicesDelete", response, response.Code())
	}
}

// NewServicesDeleteNoContent creates a ServicesDeleteNoContent with default headers values
func NewServicesDeleteNoContent() *ServicesDeleteNoContent {
	return &ServicesDeleteNoContent{}
}

/*
ServicesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ServicesDeleteNoContent struct {
}

// IsSuccess returns true when this services delete no content response has a 2xx status code
func (o *ServicesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this services delete no content response has a 3xx status code
func (o *ServicesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this services delete no content response has a 4xx status code
func (o *ServicesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this services delete no content response has a 5xx status code
func (o *ServicesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this services delete no content response a status code equal to that given
func (o *ServicesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the services delete no content response
func (o *ServicesDeleteNoContent) Code() int {
	return 204
}

func (o *ServicesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/{id}][%d] servicesDeleteNoContent ", 204)
}

func (o *ServicesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/{id}][%d] servicesDeleteNoContent ", 204)
}

func (o *ServicesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}