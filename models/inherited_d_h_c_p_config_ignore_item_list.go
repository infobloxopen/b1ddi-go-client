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

// InheritedDHCPConfigIgnoreItemList IgnoreItemList
//
// The inheritance configuration for a field that contains a list of _IgnoreItem_ objects
//
// swagger:model InheritedDHCPConfigIgnoreItemList
type InheritedDHCPConfigIgnoreItemList struct {

	// The inheritance setting.
	//
	// Valid values are:
	// * _inherit_: Use the inherited value.
	// * _override_: Use the value set in the object.
	//
	// Defaults to _inherit_.
	Action string `json:"action,omitempty"`

	// The human-readable display name for the object referred to by _source_.
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// The resource identifier.
	Source string `json:"source,omitempty"`

	// The inherited value.
	// Read Only: true
	Value []*IpamsvcIgnoreItem `json:"value"`
}

// Validate validates this inherited d h c p config ignore item list
func (m *InheritedDHCPConfigIgnoreItemList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InheritedDHCPConfigIgnoreItemList) validateValue(formats strfmt.Registry) error {
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

// ContextValidate validate this inherited d h c p config ignore item list based on the context it is used
func (m *InheritedDHCPConfigIgnoreItemList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *InheritedDHCPConfigIgnoreItemList) contextValidateDisplayName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display_name", "body", string(m.DisplayName)); err != nil {
		return err
	}

	return nil
}

func (m *InheritedDHCPConfigIgnoreItemList) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", []*IpamsvcIgnoreItem(m.Value)); err != nil {
		return err
	}

	for i := 0; i < len(m.Value); i++ {

		if m.Value[i] != nil {

			if swag.IsZero(m.Value[i]) { // not required
				return nil
			}

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
func (m *InheritedDHCPConfigIgnoreItemList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InheritedDHCPConfigIgnoreItemList) UnmarshalBinary(b []byte) error {
	var res InheritedDHCPConfigIgnoreItemList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
