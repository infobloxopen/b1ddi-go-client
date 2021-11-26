// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigInheritedKerberosKeys InheritedKerberosKeys
//
// Inheritance configuration for a field of type list of _kerberos_key_.
//
// swagger:model configInheritedKerberosKeys
type ConfigInheritedKerberosKeys struct {

	// Optional. Inheritance setting for a field.
	// Defaults to _inherit_.
	Action string `json:"action,omitempty"`

	// Human-readable display name for the object referred to by _source_.
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// The resource identifier.
	Source string `json:"source,omitempty"`

	// Inherited value.
	// Read Only: true
	Value []*ConfigKerberosKey `json:"value"`
}

// Validate validates this config inherited kerberos keys
func (m *ConfigInheritedKerberosKeys) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigInheritedKerberosKeys) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	for i := 0; i < len(m.Value); i++ {
		if swag.IsZero(m.Value[i]) { // not required
			continue
		}

		if m.Value[i] != nil {
			if err := m.Value[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this config inherited kerberos keys based on the context it is used
func (m *ConfigInheritedKerberosKeys) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplayName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigInheritedKerberosKeys) contextValidateDisplayName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display_name", "body", string(m.DisplayName)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigInheritedKerberosKeys) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", []*ConfigKerberosKey(m.Value)); err != nil {
		return err
	}

	for i := 0; i < len(m.Value); i++ {

		if m.Value[i] != nil {
			if err := m.Value[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigInheritedKerberosKeys) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigInheritedKerberosKeys) UnmarshalBinary(b []byte) error {
	var res ConfigInheritedKerberosKeys
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
