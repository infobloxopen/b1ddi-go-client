// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IpamsvcCopyIPSpaceResponse ipamsvc copy IP space response
//
// swagger:model ipamsvcCopyIPSpaceResponse
type IpamsvcCopyIPSpaceResponse struct {

	// The resource identifier.
	OperationID string `json:"operation_id,omitempty"`
}

// Validate validates this ipamsvc copy IP space response
func (m *IpamsvcCopyIPSpaceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ipamsvc copy IP space response based on context it is used
func (m *IpamsvcCopyIPSpaceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcCopyIPSpaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcCopyIPSpaceResponse) UnmarshalBinary(b []byte) error {
	var res IpamsvcCopyIPSpaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
