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

// GlobalUpdate2Reader is a Reader for the GlobalUpdate2 structure.
type GlobalUpdate2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GlobalUpdate2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, runtime.NewAPIError("response status code indicates client error", response, response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, runtime.NewAPIError("response status code indicates server error", response, response.Code())
	}

	result := NewGlobalUpdate2Created()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewGlobalUpdate2Created creates a GlobalUpdate2Created with default headers values
func NewGlobalUpdate2Created() *GlobalUpdate2Created {
	return &GlobalUpdate2Created{}
}

/* GlobalUpdate2Created describes a response with status code 201, with default header values.

PATCH operation response
*/
type GlobalUpdate2Created struct {
	Payload *models.IpamsvcUpdateGlobalResponse
}

func (o *GlobalUpdate2Created) Error() string {
	return fmt.Sprintf("[PATCH /dhcp/global/{id}][%d] globalUpdate2Created  %+v", 201, o.Payload)
}
func (o *GlobalUpdate2Created) GetPayload() *models.IpamsvcUpdateGlobalResponse {
	return o.Payload
}

func (o *GlobalUpdate2Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcUpdateGlobalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}