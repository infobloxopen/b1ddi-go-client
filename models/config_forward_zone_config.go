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

// ConfigForwardZoneConfig config forward zone config
//
// swagger:model configForwardZoneConfig
type ConfigForwardZoneConfig struct {

	// Optional. External DNS servers to forward to. Order is not significant.
	ExternalForwarders []*ConfigForwarder `json:"external_forwarders"`

	// The resource identifier.
	Hosts []string `json:"hosts"`

	// The resource identifier.
	InternalForwarders []string `json:"internal_forwarders"`

	// The resource identifier.
	Nsgs []string `json:"nsgs"`
}

// Validate validates this config forward zone config
func (m *ConfigForwardZoneConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalForwarders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigForwardZoneConfig) validateExternalForwarders(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalForwarders) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalForwarders); i++ {
		if swag.IsZero(m.ExternalForwarders[i]) { // not required
			continue
		}

		if m.ExternalForwarders[i] != nil {
			if err := m.ExternalForwarders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_forwarders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_forwarders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this config forward zone config based on the context it is used
func (m *ConfigForwardZoneConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalForwarders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigForwardZoneConfig) contextValidateExternalForwarders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExternalForwarders); i++ {

		if m.ExternalForwarders[i] != nil {
			if err := m.ExternalForwarders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_forwarders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_forwarders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigForwardZoneConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigForwardZoneConfig) UnmarshalBinary(b []byte) error {
	var res ConfigForwardZoneConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
