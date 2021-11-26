// Code generated by go-swagger; DO NOT EDIT.

package option_code

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OptionCodeDeleteReader is a Reader for the OptionCodeDelete structure.
type OptionCodeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OptionCodeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOptionCodeDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOptionCodeDeleteNoContent creates a OptionCodeDeleteNoContent with default headers values
func NewOptionCodeDeleteNoContent() *OptionCodeDeleteNoContent {
	return &OptionCodeDeleteNoContent{}
}

/* OptionCodeDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type OptionCodeDeleteNoContent struct {
}

func (o *OptionCodeDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dhcp/option_code/{id}][%d] optionCodeDeleteNoContent ", 204)
}

func (o *OptionCodeDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
