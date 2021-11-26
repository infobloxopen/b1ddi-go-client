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

// IpamsvcDHCPConfig DHCPConfig
//
// A DHCP Config object (_dhcp/dhcp_config_) represents a shared DHCP configuration that controls how leases are issued.
//
// swagger:model ipamsvcDHCPConfig
type IpamsvcDHCPConfig struct {

	// Disable to allow leases only for known clients, those for which a fixed address is configured.
	AllowUnknown bool `json:"allow_unknown,omitempty"`

	// The resource identifier.
	Filters []string `json:"filters"`

	// The list of clients to ignore requests from.
	IgnoreList []*IpamsvcIgnoreItem `json:"ignore_list"`

	// The lease duration in seconds.
	LeaseTime int64 `json:"lease_time,omitempty"`
}

// Validate validates this ipamsvc d h c p config
func (m *IpamsvcDHCPConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIgnoreList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcDHCPConfig) validateIgnoreList(formats strfmt.Registry) error {
	if swag.IsZero(m.IgnoreList) { // not required
		return nil
	}

	for i := 0; i < len(m.IgnoreList); i++ {
		if swag.IsZero(m.IgnoreList[i]) { // not required
			continue
		}

		if m.IgnoreList[i] != nil {
			if err := m.IgnoreList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ignore_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ignore_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ipamsvc d h c p config based on the context it is used
func (m *IpamsvcDHCPConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIgnoreList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcDHCPConfig) contextValidateIgnoreList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IgnoreList); i++ {

		if m.IgnoreList[i] != nil {
			if err := m.IgnoreList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ignore_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ignore_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcDHCPConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcDHCPConfig) UnmarshalBinary(b []byte) error {
	var res IpamsvcDHCPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
