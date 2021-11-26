// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IpamsvcInheritedDHCPConfig InheritedDHCPConfig
//
// The inheritance configuration for a field of type _DHCPConfig_.
//
// swagger:model ipamsvcInheritedDHCPConfig
type IpamsvcInheritedDHCPConfig struct {

	// The inheritance configuration for _allow_unknown_ field from _DHCPConfig_ object.
	AllowUnknown *InheritanceInheritedBool `json:"allow_unknown,omitempty"`

	// The inheritance configuration for filters field from _DHCPConfig_ object.
	Filters *InheritedDHCPConfigFilterList `json:"filters,omitempty"`

	// The inheritance configuration for _ignore_list_ field from _DHCPConfig_ object.
	IgnoreList *InheritedDHCPConfigIgnoreItemList `json:"ignore_list,omitempty"`

	// The inheritance configuration for _lease_time_ field from _DHCPConfig_ object.
	LeaseTime *InheritanceInheritedUInt32 `json:"lease_time,omitempty"`
}

// Validate validates this ipamsvc inherited d h c p config
func (m *IpamsvcInheritedDHCPConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowUnknown(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgnoreList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeaseTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcInheritedDHCPConfig) validateAllowUnknown(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowUnknown) { // not required
		return nil
	}

	if m.AllowUnknown != nil {
		if err := m.AllowUnknown.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allow_unknown")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allow_unknown")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedDHCPConfig) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	if m.Filters != nil {
		if err := m.Filters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedDHCPConfig) validateIgnoreList(formats strfmt.Registry) error {
	if swag.IsZero(m.IgnoreList) { // not required
		return nil
	}

	if m.IgnoreList != nil {
		if err := m.IgnoreList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ignore_list")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ignore_list")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedDHCPConfig) validateLeaseTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LeaseTime) { // not required
		return nil
	}

	if m.LeaseTime != nil {
		if err := m.LeaseTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lease_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lease_time")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ipamsvc inherited d h c p config based on the context it is used
func (m *IpamsvcInheritedDHCPConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowUnknown(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIgnoreList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLeaseTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcInheritedDHCPConfig) contextValidateAllowUnknown(ctx context.Context, formats strfmt.Registry) error {

	if m.AllowUnknown != nil {
		if err := m.AllowUnknown.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allow_unknown")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allow_unknown")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedDHCPConfig) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	if m.Filters != nil {
		if err := m.Filters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedDHCPConfig) contextValidateIgnoreList(ctx context.Context, formats strfmt.Registry) error {

	if m.IgnoreList != nil {
		if err := m.IgnoreList.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ignore_list")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ignore_list")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedDHCPConfig) contextValidateLeaseTime(ctx context.Context, formats strfmt.Registry) error {

	if m.LeaseTime != nil {
		if err := m.LeaseTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lease_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lease_time")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcInheritedDHCPConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcInheritedDHCPConfig) UnmarshalBinary(b []byte) error {
	var res IpamsvcInheritedDHCPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
