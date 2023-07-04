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

// ConfigECSZone ECSZone
//
// EDNS Client Subnet zone.
//
// swagger:model configECSZone
type ConfigECSZone struct {

	// Access control for zone.
	//
	// Allowed values:
	// * _allow_,
	// * _deny_.
	// Required: true
	Access *string `json:"access"`

	// Zone FQDN.
	// Required: true
	Fqdn *string `json:"fqdn"`

	// Zone FQDN in punycode.
	// Read Only: true
	ProtocolFqdn string `json:"protocol_fqdn,omitempty"`
}

// Validate validates this config e c s zone
func (m *ConfigECSZone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFqdn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigECSZone) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", m.Access); err != nil {
		return err
	}

	return nil
}

func (m *ConfigECSZone) validateFqdn(formats strfmt.Registry) error {

	if err := validate.Required("fqdn", "body", m.Fqdn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this config e c s zone based on the context it is used
func (m *ConfigECSZone) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtocolFqdn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigECSZone) contextValidateProtocolFqdn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol_fqdn", "body", string(m.ProtocolFqdn)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigECSZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigECSZone) UnmarshalBinary(b []byte) error {
	var res ConfigECSZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
