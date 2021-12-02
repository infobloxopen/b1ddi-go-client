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

// ConfigAuthZoneInheritance config auth zone inheritance
//
// swagger:model configAuthZoneInheritance
type ConfigAuthZoneInheritance struct {

	// Optional. Field config for _gss_tsig_enabled_ field from _AuthZone_ object.
	GssTsigEnabled *Inheritance2InheritedBool `json:"gss_tsig_enabled,omitempty"`

	// Field config for _notify_ field from _AuthZone_ object.
	Notify *Inheritance2InheritedBool `json:"notify,omitempty"`

	// Optional. Field config for _query_acl_ field from _AuthZone_ object.
	QueryACL *ConfigInheritedACLItems `json:"query_acl,omitempty"`

	// Optional. Field config for _transfer_acl_ field from _AuthZone_ object.
	TransferACL *ConfigInheritedACLItems `json:"transfer_acl,omitempty"`

	// Optional. Field config for _update_acl_ field from _AuthZone_ object.
	UpdateACL *ConfigInheritedACLItems `json:"update_acl,omitempty"`

	// Optional. Field config for _use_forwarders_for_subzones_ field from _AuthZone_ object.
	UseForwardersForSubzones *Inheritance2InheritedBool `json:"use_forwarders_for_subzones,omitempty"`

	// Optional. Field config for _zone_authority_ field from _AuthZone_ object.
	ZoneAuthority *ConfigInheritedZoneAuthority `json:"zone_authority,omitempty"`
}

// Validate validates this config auth zone inheritance
func (m *ConfigAuthZoneInheritance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGssTsigEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotify(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransferACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseForwardersForSubzones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneAuthority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigAuthZoneInheritance) validateGssTsigEnabled(formats strfmt.Registry) error {
	if swag.IsZero(m.GssTsigEnabled) { // not required
		return nil
	}

	if m.GssTsigEnabled != nil {
		if err := m.GssTsigEnabled.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gss_tsig_enabled")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gss_tsig_enabled")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) validateNotify(formats strfmt.Registry) error {
	if swag.IsZero(m.Notify) { // not required
		return nil
	}

	if m.Notify != nil {
		if err := m.Notify.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notify")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notify")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) validateQueryACL(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryACL) { // not required
		return nil
	}

	if m.QueryACL != nil {
		if err := m.QueryACL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query_acl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query_acl")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) validateTransferACL(formats strfmt.Registry) error {
	if swag.IsZero(m.TransferACL) { // not required
		return nil
	}

	if m.TransferACL != nil {
		if err := m.TransferACL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transfer_acl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transfer_acl")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) validateUpdateACL(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateACL) { // not required
		return nil
	}

	if m.UpdateACL != nil {
		if err := m.UpdateACL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update_acl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("update_acl")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) validateUseForwardersForSubzones(formats strfmt.Registry) error {
	if swag.IsZero(m.UseForwardersForSubzones) { // not required
		return nil
	}

	if m.UseForwardersForSubzones != nil {
		if err := m.UseForwardersForSubzones.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("use_forwarders_for_subzones")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("use_forwarders_for_subzones")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) validateZoneAuthority(formats strfmt.Registry) error {
	if swag.IsZero(m.ZoneAuthority) { // not required
		return nil
	}

	if m.ZoneAuthority != nil {
		if err := m.ZoneAuthority.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone_authority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zone_authority")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config auth zone inheritance based on the context it is used
func (m *ConfigAuthZoneInheritance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGssTsigEnabled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotify(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueryACL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransferACL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateACL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUseForwardersForSubzones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZoneAuthority(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigAuthZoneInheritance) contextValidateGssTsigEnabled(ctx context.Context, formats strfmt.Registry) error {

	if m.GssTsigEnabled != nil {
		if err := m.GssTsigEnabled.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gss_tsig_enabled")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gss_tsig_enabled")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) contextValidateNotify(ctx context.Context, formats strfmt.Registry) error {

	if m.Notify != nil {
		if err := m.Notify.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notify")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notify")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) contextValidateQueryACL(ctx context.Context, formats strfmt.Registry) error {

	if m.QueryACL != nil {
		if err := m.QueryACL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query_acl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query_acl")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) contextValidateTransferACL(ctx context.Context, formats strfmt.Registry) error {

	if m.TransferACL != nil {
		if err := m.TransferACL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transfer_acl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transfer_acl")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) contextValidateUpdateACL(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdateACL != nil {
		if err := m.UpdateACL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update_acl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("update_acl")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) contextValidateUseForwardersForSubzones(ctx context.Context, formats strfmt.Registry) error {

	if m.UseForwardersForSubzones != nil {
		if err := m.UseForwardersForSubzones.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("use_forwarders_for_subzones")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("use_forwarders_for_subzones")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAuthZoneInheritance) contextValidateZoneAuthority(ctx context.Context, formats strfmt.Registry) error {

	if m.ZoneAuthority != nil {
		if err := m.ZoneAuthority.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone_authority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zone_authority")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigAuthZoneInheritance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigAuthZoneInheritance) UnmarshalBinary(b []byte) error {
	var res ConfigAuthZoneInheritance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}