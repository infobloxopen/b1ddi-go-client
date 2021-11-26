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

// ConfigForwardZone ForwardZone
//
// Forward zone
//
// swagger:model configForwardZone
type ConfigForwardZone struct {

	// Optional. Comment for zone configuration.
	Comment string `json:"comment,omitempty"`

	// The timestamp when the object has been created.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
	Disabled bool `json:"disabled,omitempty"`

	// Optional. External DNS servers to forward to. Order is not significant.
	ExternalForwarders []*ConfigForwarder `json:"external_forwarders"`

	// Optional. _true_ to only forward.
	ForwardOnly bool `json:"forward_only,omitempty"`

	// Zone FQDN.
	// The FQDN supplied at creation will be converted to canonical form.
	//
	// Read-only after creation.
	// Required: true
	Fqdn *string `json:"fqdn"`

	// The resource identifier.
	Hosts []string `json:"hosts"`

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The resource identifier.
	InternalForwarders []string `json:"internal_forwarders"`

	// Reverse zone network address in the following format: "ip-address/cidr".
	// Defaults to empty.
	// Read Only: true
	MappedSubnet string `json:"mapped_subnet,omitempty"`

	// Read-only. Zone mapping type.
	// Allowed values:
	//  * _forward_,
	//  * _ipv4_reverse_.
	//  * _ipv6_reverse_.
	//
	// Defaults to _forward_.
	// Read Only: true
	Mapping string `json:"mapping,omitempty"`

	// The resource identifier.
	Nsgs []string `json:"nsgs"`

	// The resource identifier.
	Parent string `json:"parent,omitempty"`

	// Zone FQDN in punycode.
	// Read Only: true
	ProtocolFqdn string `json:"protocol_fqdn,omitempty"`

	// Tagging specifics.
	Tags interface{} `json:"tags,omitempty"`

	// The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// The resource identifier.
	View string `json:"view,omitempty"`
}

// Validate validates this config forward zone
func (m *ConfigForwardZone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalForwarders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFqdn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigForwardZone) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigForwardZone) validateExternalForwarders(formats strfmt.Registry) error {
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

func (m *ConfigForwardZone) validateFqdn(formats strfmt.Registry) error {

	if err := validate.Required("fqdn", "body", m.Fqdn); err != nil {
		return err
	}

	return nil
}

func (m *ConfigForwardZone) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this config forward zone based on the context it is used
func (m *ConfigForwardZone) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalForwarders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMappedSubnet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMapping(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocolFqdn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigForwardZone) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigForwardZone) contextValidateExternalForwarders(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConfigForwardZone) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigForwardZone) contextValidateMappedSubnet(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mapped_subnet", "body", string(m.MappedSubnet)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigForwardZone) contextValidateMapping(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mapping", "body", string(m.Mapping)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigForwardZone) contextValidateProtocolFqdn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol_fqdn", "body", string(m.ProtocolFqdn)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigForwardZone) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigForwardZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigForwardZone) UnmarshalBinary(b []byte) error {
	var res ConfigForwardZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
