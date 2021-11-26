// Code generated by go-swagger; DO NOT EDIT.

package delegation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// DelegationListReader is a Reader for the DelegationList structure.
type DelegationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DelegationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDelegationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDelegationListOK creates a DelegationListOK with default headers values
func NewDelegationListOK() *DelegationListOK {
	return &DelegationListOK{}
}

/* DelegationListOK describes a response with status code 200, with default header values.

GET operation response
*/
type DelegationListOK struct {
	Payload *models.ConfigListDelegationResponse
}

func (o *DelegationListOK) Error() string {
	return fmt.Sprintf("[GET /dns/delegation][%d] delegationListOK  %+v", 200, o.Payload)
}
func (o *DelegationListOK) GetPayload() *models.ConfigListDelegationResponse {
	return o.Payload
}

func (o *DelegationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigListDelegationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
