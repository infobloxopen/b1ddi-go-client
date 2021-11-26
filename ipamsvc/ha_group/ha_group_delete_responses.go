// Code generated by go-swagger; DO NOT EDIT.

package ha_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HaGroupDeleteReader is a Reader for the HaGroupDelete structure.
type HaGroupDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HaGroupDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewHaGroupDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHaGroupDeleteNoContent creates a HaGroupDeleteNoContent with default headers values
func NewHaGroupDeleteNoContent() *HaGroupDeleteNoContent {
	return &HaGroupDeleteNoContent{}
}

/* HaGroupDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type HaGroupDeleteNoContent struct {
}

func (o *HaGroupDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dhcp/ha_group/{id}][%d] haGroupDeleteNoContent ", 204)
}

func (o *HaGroupDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
