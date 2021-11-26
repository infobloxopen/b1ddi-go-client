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

// Inheritance2AssignedHost AssignedHost
//
// _ddi/assigned_host_ represents a BloxOne DDI host assigned to an object.
//
// swagger:model inheritance2AssignedHost
type Inheritance2AssignedHost struct {

	// The human-readable display name for the host referred to by _ophid_.
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// The resource identifier.
	Host string `json:"host,omitempty"`

	// The on-prem host ID.
	// Read Only: true
	Ophid string `json:"ophid,omitempty"`
}

// Validate validates this inheritance2 assigned host
func (m *Inheritance2AssignedHost) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this inheritance2 assigned host based on the context it is used
func (m *Inheritance2AssignedHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplayName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOphid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Inheritance2AssignedHost) contextValidateDisplayName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display_name", "body", string(m.DisplayName)); err != nil {
		return err
	}

	return nil
}

func (m *Inheritance2AssignedHost) contextValidateOphid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ophid", "body", string(m.Ophid)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Inheritance2AssignedHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Inheritance2AssignedHost) UnmarshalBinary(b []byte) error {
	var res Inheritance2AssignedHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
