// Code generated by go-swagger; DO NOT EDIT.

package ip_space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IPSpaceDeleteReader is a Reader for the IPSpaceDelete structure.
type IPSpaceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPSpaceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewIPSpaceDeleteNoContent()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewIPSpaceDeleteNoContent creates a IPSpaceDeleteNoContent with default headers values
func NewIPSpaceDeleteNoContent() *IPSpaceDeleteNoContent {
	return &IPSpaceDeleteNoContent{}
}

/* IPSpaceDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type IPSpaceDeleteNoContent struct {
}

func (o *IPSpaceDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip_space/{id}][%d] ipSpaceDeleteNoContent ", 204)
}

func (o *IPSpaceDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
