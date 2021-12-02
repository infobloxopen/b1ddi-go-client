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

// IpamsvcHostAssociationsResponse HostAssociationsResponse
//
// The response format to retrieve __HAGroup__ and __Subnet__ objects associated with the DHCP __Host__ object. The host in question is also included in the output, for the convenience reasons.
//
// swagger:model ipamsvcHostAssociationsResponse
type IpamsvcHostAssociationsResponse struct {

	// The list of HA groups.
	HaGroups []*IpamsvcHAGroup `json:"ha_groups"`

	// The host for which the associated objects, subnets and HA groups, are returned.
	Host *IpamsvcHost `json:"host,omitempty"`

	// The list of subnets.
	Subnets []*IpamsvcSubnet `json:"subnets"`
}

// Validate validates this ipamsvc host associations response
func (m *IpamsvcHostAssociationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHaGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcHostAssociationsResponse) validateHaGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.HaGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.HaGroups); i++ {
		if swag.IsZero(m.HaGroups[i]) { // not required
			continue
		}

		if m.HaGroups[i] != nil {
			if err := m.HaGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ha_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ha_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcHostAssociationsResponse) validateHost(formats strfmt.Registry) error {
	if swag.IsZero(m.Host) { // not required
		return nil
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcHostAssociationsResponse) validateSubnets(formats strfmt.Registry) error {
	if swag.IsZero(m.Subnets) { // not required
		return nil
	}

	for i := 0; i < len(m.Subnets); i++ {
		if swag.IsZero(m.Subnets[i]) { // not required
			continue
		}

		if m.Subnets[i] != nil {
			if err := m.Subnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ipamsvc host associations response based on the context it is used
func (m *IpamsvcHostAssociationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHaGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcHostAssociationsResponse) contextValidateHaGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HaGroups); i++ {

		if m.HaGroups[i] != nil {
			if err := m.HaGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ha_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ha_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcHostAssociationsResponse) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if m.Host != nil {
		if err := m.Host.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcHostAssociationsResponse) contextValidateSubnets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Subnets); i++ {

		if m.Subnets[i] != nil {
			if err := m.Subnets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcHostAssociationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcHostAssociationsResponse) UnmarshalBinary(b []byte) error {
	var res IpamsvcHostAssociationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}