// Code generated by go-swagger; DO NOT EDIT.

package auth_nsg

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AuthNsgDeleteReader is a Reader for the AuthNsgDelete structure.
type AuthNsgDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthNsgDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAuthNsgDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthNsgDeleteNoContent creates a AuthNsgDeleteNoContent with default headers values
func NewAuthNsgDeleteNoContent() *AuthNsgDeleteNoContent {
	return &AuthNsgDeleteNoContent{}
}

/* AuthNsgDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type AuthNsgDeleteNoContent struct {
}

func (o *AuthNsgDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dns/auth_nsg/{id}][%d] authNsgDeleteNoContent ", 204)
}

func (o *AuthNsgDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
