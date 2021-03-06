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
)

// IpamsvcInheritedDHCPOptionList InheritedDHCPOptionList
//
// The inheritance configuration for a field that contains list of _OptionItem_.
//
// swagger:model ipamsvcInheritedDHCPOptionList
type IpamsvcInheritedDHCPOptionList struct {

	// The inheritance setting.
	//
	// Valid values are:
	// * _inherit_: Use the inherited value.
	// * _block_: Don't use the inherited value.
	//
	// Defaults to _inherit_.
	Action string `json:"action,omitempty"`

	// The inherited DHCP option values.
	Value []*IpamsvcInheritedDHCPOption `json:"value"`
}

// Validate validates this ipamsvc inherited d h c p option list
func (m *IpamsvcInheritedDHCPOptionList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcInheritedDHCPOptionList) validateValue(formats strfmt.Registry) error {
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

// ContextValidate validate this ipamsvc inherited d h c p option list based on the context it is used
func (m *IpamsvcInheritedDHCPOptionList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcInheritedDHCPOptionList) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IpamsvcInheritedDHCPOptionList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcInheritedDHCPOptionList) UnmarshalBinary(b []byte) error {
	var res IpamsvcInheritedDHCPOptionList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
