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

// ConfigDelegation Delegation
//
// DNS zone delegation.
//
// swagger:model configDelegation
type ConfigDelegation struct {

	// Optional. Comment for zone delegation.
	Comment string `json:"comment,omitempty"`

	// Required. DNS zone delegation servers. Order is not significant.
	DelegationServers []*ConfigDelegationServer `json:"delegation_servers"`

	// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating resource records.
	Disabled bool `json:"disabled,omitempty"`

	// Delegation FQDN.
	// The FQDN supplied at creation will be converted to canonical form.
	//
	// Read-only after creation.
	// Required: true
	Fqdn *string `json:"fqdn"`

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The resource identifier.
	Parent string `json:"parent,omitempty"`

	// Delegation FQDN in punycode.
	// Read Only: true
	ProtocolFqdn string `json:"protocol_fqdn,omitempty"`

	// Tagging specifics.
	Tags interface{} `json:"tags,omitempty"`

	// The resource identifier.
	View string `json:"view,omitempty"`
}

// Validate validates this config delegation
func (m *ConfigDelegation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelegationServers(formats); err != nil {
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

func (m *ConfigDelegation) validateDelegationServers(formats strfmt.Registry) error {
	if swag.IsZero(m.DelegationServers) { // not required
		return nil
	}

	for i := 0; i < len(m.DelegationServers); i++ {
		if swag.IsZero(m.DelegationServers[i]) { // not required
			continue
		}

		if m.DelegationServers[i] != nil {
			if err := m.DelegationServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegation_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delegation_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigDelegation) validateFqdn(formats strfmt.Registry) error {

	if err := validate.Required("fqdn", "body", m.Fqdn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this config delegation based on the context it is used
func (m *ConfigDelegation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDelegationServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocolFqdn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigDelegation) contextValidateDelegationServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DelegationServers); i++ {

		if m.DelegationServers[i] != nil {
			if err := m.DelegationServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegation_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delegation_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigDelegation) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigDelegation) contextValidateProtocolFqdn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol_fqdn", "body", string(m.ProtocolFqdn)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigDelegation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigDelegation) UnmarshalBinary(b []byte) error {
	var res ConfigDelegation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}