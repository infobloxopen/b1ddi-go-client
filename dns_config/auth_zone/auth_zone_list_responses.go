// Code generated by go-swagger; DO NOT EDIT.

package auth_zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// AuthZoneListReader is a Reader for the AuthZoneList structure.
type AuthZoneListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthZoneListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthZoneListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthZoneListOK creates a AuthZoneListOK with default headers values
func NewAuthZoneListOK() *AuthZoneListOK {
	return &AuthZoneListOK{}
}

/* AuthZoneListOK describes a response with status code 200, with default header values.

GET operation response
*/
type AuthZoneListOK struct {
	Payload *models.ConfigListAuthZoneResponse
}

func (o *AuthZoneListOK) Error() string {
	return fmt.Sprintf("[GET /dns/auth_zone][%d] authZoneListOK  %+v", 200, o.Payload)
}
func (o *AuthZoneListOK) GetPayload() *models.ConfigListAuthZoneResponse {
	return o.Payload
}

func (o *AuthZoneListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigListAuthZoneResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
