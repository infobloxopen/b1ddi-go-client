// Code generated by go-swagger; DO NOT EDIT.

package option_space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OptionSpaceDeleteReader is a Reader for the OptionSpaceDelete structure.
type OptionSpaceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OptionSpaceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewOptionSpaceDeleteNoContent()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewOptionSpaceDeleteNoContent creates a OptionSpaceDeleteNoContent with default headers values
func NewOptionSpaceDeleteNoContent() *OptionSpaceDeleteNoContent {
	return &OptionSpaceDeleteNoContent{}
}

/* OptionSpaceDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type OptionSpaceDeleteNoContent struct {
}

func (o *OptionSpaceDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dhcp/option_space/{id}][%d] optionSpaceDeleteNoContent ", 204)
}

func (o *OptionSpaceDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
