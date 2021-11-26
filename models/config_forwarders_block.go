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

// ConfigForwardersBlock ForwardersBlock
//
// Block for fields: _forwarders_, _forwarders_only_.
//
// swagger:model configForwardersBlock
type ConfigForwardersBlock struct {

	// Optional. Field config for _forwarders_ field from.
	Forwarders []*ConfigForwarder `json:"forwarders"`

	// Optional. Field config for _forwarders_only_ field.
	ForwardersOnly bool `json:"forwarders_only,omitempty"`
}

// Validate validates this config forwarders block
func (m *ConfigForwardersBlock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateForwarders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigForwardersBlock) validateForwarders(formats strfmt.Registry) error {
	if swag.IsZero(m.Forwarders) { // not required
		return nil
	}

	for i := 0; i < len(m.Forwarders); i++ {
		if swag.IsZero(m.Forwarders[i]) { // not required
			continue
		}

		if m.Forwarders[i] != nil {
			if err := m.Forwarders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("forwarders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("forwarders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this config forwarders block based on the context it is used
func (m *ConfigForwardersBlock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateForwarders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigForwardersBlock) contextValidateForwarders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Forwarders); i++ {

		if m.Forwarders[i] != nil {
			if err := m.Forwarders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("forwarders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("forwarders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigForwardersBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigForwardersBlock) UnmarshalBinary(b []byte) error {
	var res ConfigForwardersBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
