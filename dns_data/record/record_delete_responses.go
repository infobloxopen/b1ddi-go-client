// Code generated by go-swagger; DO NOT EDIT.

package record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// RecordDeleteReader is a Reader for the RecordDelete structure.
type RecordDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecordDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewRecordDeleteNoContent()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewRecordDeleteNoContent creates a RecordDeleteNoContent with default headers values
func NewRecordDeleteNoContent() *RecordDeleteNoContent {
	return &RecordDeleteNoContent{}
}

/* RecordDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type RecordDeleteNoContent struct {
}

func (o *RecordDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dns/record/{id}][%d] recordDeleteNoContent ", 204)
}

func (o *RecordDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
