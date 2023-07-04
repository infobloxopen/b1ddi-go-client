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

// IpamsvcInheritedASMConfig InheritedASMConfig
//
// The inheritance configuration for the __ASMConfig__ object.
//
// swagger:model ipamsvcInheritedASMConfig
type IpamsvcInheritedASMConfig struct {

	// The block of ASM fields: _enable_, _enable_notification_, _reenable_date_.
	AsmEnableBlock *IpamsvcInheritedAsmEnableBlock `json:"asm_enable_block,omitempty"`

	// The block of ASM fields: _growth_factor_, _growth_type_.
	AsmGrowthBlock *IpamsvcInheritedAsmGrowthBlock `json:"asm_growth_block,omitempty"`

	// ASM shows the number of addresses forecast to be used _forecast_period_ days in the future, if it is greater than _asm_threshold_percent_ * _dhcp_total_ (see _dhcp_utilization_) then the subnet is flagged.
	AsmThreshold *InheritanceInheritedUInt32 `json:"asm_threshold,omitempty"`

	// The forecast period in days.
	ForecastPeriod *InheritanceInheritedUInt32 `json:"forecast_period,omitempty"`

	// The minimum amount of history needed before ASM can run on this subnet.
	History *InheritanceInheritedUInt32 `json:"history,omitempty"`

	// The minimum size of range needed for ASM to run on this subnet.
	MinTotal *InheritanceInheritedUInt32 `json:"min_total,omitempty"`

	// The minimum percentage of addresses that must be available outside of the DHCP ranges and fixed addresses when making a suggested change.
	MinUnused *InheritanceInheritedUInt32 `json:"min_unused,omitempty"`
}

// Validate validates this ipamsvc inherited a s m config
func (m *IpamsvcInheritedASMConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsmEnableBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAsmGrowthBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAsmThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForecastPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinUnused(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcInheritedASMConfig) validateAsmEnableBlock(formats strfmt.Registry) error {
	if swag.IsZero(m.AsmEnableBlock) { // not required
		return nil
	}

	if m.AsmEnableBlock != nil {
		if err := m.AsmEnableBlock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asm_enable_block")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asm_enable_block")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) validateAsmGrowthBlock(formats strfmt.Registry) error {
	if swag.IsZero(m.AsmGrowthBlock) { // not required
		return nil
	}

	if m.AsmGrowthBlock != nil {
		if err := m.AsmGrowthBlock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asm_growth_block")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asm_growth_block")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) validateAsmThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.AsmThreshold) { // not required
		return nil
	}

	if m.AsmThreshold != nil {
		if err := m.AsmThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asm_threshold")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asm_threshold")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) validateForecastPeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.ForecastPeriod) { // not required
		return nil
	}

	if m.ForecastPeriod != nil {
		if err := m.ForecastPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forecast_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("forecast_period")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) validateHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.History) { // not required
		return nil
	}

	if m.History != nil {
		if err := m.History.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("history")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("history")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) validateMinTotal(formats strfmt.Registry) error {
	if swag.IsZero(m.MinTotal) { // not required
		return nil
	}

	if m.MinTotal != nil {
		if err := m.MinTotal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("min_total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("min_total")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) validateMinUnused(formats strfmt.Registry) error {
	if swag.IsZero(m.MinUnused) { // not required
		return nil
	}

	if m.MinUnused != nil {
		if err := m.MinUnused.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("min_unused")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("min_unused")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ipamsvc inherited a s m config based on the context it is used
func (m *IpamsvcInheritedASMConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAsmEnableBlock(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAsmGrowthBlock(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAsmThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateForecastPeriod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinUnused(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcInheritedASMConfig) contextValidateAsmEnableBlock(ctx context.Context, formats strfmt.Registry) error {

	if m.AsmEnableBlock != nil {

		if swag.IsZero(m.AsmEnableBlock) { // not required
			return nil
		}

		if err := m.AsmEnableBlock.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asm_enable_block")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asm_enable_block")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) contextValidateAsmGrowthBlock(ctx context.Context, formats strfmt.Registry) error {

	if m.AsmGrowthBlock != nil {

		if swag.IsZero(m.AsmGrowthBlock) { // not required
			return nil
		}

		if err := m.AsmGrowthBlock.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asm_growth_block")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asm_growth_block")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) contextValidateAsmThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.AsmThreshold != nil {

		if swag.IsZero(m.AsmThreshold) { // not required
			return nil
		}

		if err := m.AsmThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asm_threshold")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asm_threshold")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) contextValidateForecastPeriod(ctx context.Context, formats strfmt.Registry) error {

	if m.ForecastPeriod != nil {

		if swag.IsZero(m.ForecastPeriod) { // not required
			return nil
		}

		if err := m.ForecastPeriod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forecast_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("forecast_period")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) contextValidateHistory(ctx context.Context, formats strfmt.Registry) error {

	if m.History != nil {

		if swag.IsZero(m.History) { // not required
			return nil
		}

		if err := m.History.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("history")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("history")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) contextValidateMinTotal(ctx context.Context, formats strfmt.Registry) error {

	if m.MinTotal != nil {

		if swag.IsZero(m.MinTotal) { // not required
			return nil
		}

		if err := m.MinTotal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("min_total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("min_total")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcInheritedASMConfig) contextValidateMinUnused(ctx context.Context, formats strfmt.Registry) error {

	if m.MinUnused != nil {

		if swag.IsZero(m.MinUnused) { // not required
			return nil
		}

		if err := m.MinUnused.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("min_unused")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("min_unused")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcInheritedASMConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcInheritedASMConfig) UnmarshalBinary(b []byte) error {
	var res IpamsvcInheritedASMConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
