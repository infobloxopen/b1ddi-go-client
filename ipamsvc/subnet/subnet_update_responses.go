// Code generated by go-swagger; DO NOT EDIT.

package subnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// SubnetUpdateReader is a Reader for the SubnetUpdate structure.
type SubnetUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubnetUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSubnetUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubnetUpdateCreated creates a SubnetUpdateCreated with default headers values
func NewSubnetUpdateCreated() *SubnetUpdateCreated {
	return &SubnetUpdateCreated{}
}

/* SubnetUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type SubnetUpdateCreated struct {
	Payload *models.IpamsvcUpdateSubnetResponse
}

func (o *SubnetUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /ipam/subnet/{id}][%d] subnetUpdateCreated  %+v", 201, o.Payload)
}
func (o *SubnetUpdateCreated) GetPayload() *models.IpamsvcUpdateSubnetResponse {
	return o.Payload
}

func (o *SubnetUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcUpdateSubnetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
