// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InfraReplaceHostRequest ReplaceHostRequest Request Structure
//
// swagger:model infraReplaceHostRequest
type InfraReplaceHostRequest struct {

	// The resource identifier.
	From string `json:"from,omitempty"`

	// The resource identifier.
	To string `json:"to,omitempty"`
}

// Validate validates this infra replace host request
func (m *InfraReplaceHostRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this infra replace host request based on context it is used
func (m *InfraReplaceHostRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfraReplaceHostRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraReplaceHostRequest) UnmarshalBinary(b []byte) error {
	var res InfraReplaceHostRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}