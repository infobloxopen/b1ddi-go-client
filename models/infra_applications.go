// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InfraApplications infra applications
//
// swagger:model infraApplications
type InfraApplications struct {

	// applications
	Applications []string `json:"applications"`
}

// Validate validates this infra applications
func (m *InfraApplications) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this infra applications based on context it is used
func (m *InfraApplications) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfraApplications) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraApplications) UnmarshalBinary(b []byte) error {
	var res InfraApplications
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
