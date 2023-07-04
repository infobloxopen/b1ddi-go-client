// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IpamsvcLeaseSubnet ---------------------------  leases_command  ---------------------------
//
// swagger:model ipamsvcLeaseSubnet
type IpamsvcLeaseSubnet struct {

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`
}

// Validate validates this ipamsvc lease subnet
func (m *IpamsvcLeaseSubnet) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this ipamsvc lease subnet based on the context it is used
func (m *IpamsvcLeaseSubnet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcLeaseSubnet) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcLeaseSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcLeaseSubnet) UnmarshalBinary(b []byte) error {
	var res IpamsvcLeaseSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
