// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigHostInheritance HostInheritance
//
// Inheritance configuration specifies how and which fields _Host_ object inherits from _Global_ or _Server_ parent.
//
// swagger:model configHostInheritance
type ConfigHostInheritance struct {

	// Optional. Field config for _kerberos_keys_ field from _Host_ object.
	KerberosKeys *ConfigInheritedKerberosKeys `json:"kerberos_keys,omitempty"`
}

// Validate validates this config host inheritance
func (m *ConfigHostInheritance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKerberosKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigHostInheritance) validateKerberosKeys(formats strfmt.Registry) error {
	if swag.IsZero(m.KerberosKeys) { // not required
		return nil
	}

	if m.KerberosKeys != nil {
		if err := m.KerberosKeys.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kerberos_keys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kerberos_keys")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config host inheritance based on the context it is used
func (m *ConfigHostInheritance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKerberosKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigHostInheritance) contextValidateKerberosKeys(ctx context.Context, formats strfmt.Registry) error {

	if m.KerberosKeys != nil {
		if err := m.KerberosKeys.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kerberos_keys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kerberos_keys")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigHostInheritance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigHostInheritance) UnmarshalBinary(b []byte) error {
	var res ConfigHostInheritance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}