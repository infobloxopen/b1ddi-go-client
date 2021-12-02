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

// ConfigECSBlock ECSBlock
//
// Block for fields: _ecs_enabled_, _ecs_forwarding_, _ecs_prefix_v4_, _ecs_prefix_v6_, _ecs_zones_.
//
// swagger:model configECSBlock
type ConfigECSBlock struct {

	// Optional. Field config for _ecs_enabled_ field.
	EcsEnabled bool `json:"ecs_enabled,omitempty"`

	// Optional. Field config for _ecs_forwarding_ field.
	EcsForwarding bool `json:"ecs_forwarding,omitempty"`

	// Optional. Field config for _ecs_prefix_v4_ field.
	EcsPrefixV4 int64 `json:"ecs_prefix_v4,omitempty"`

	// Optional. Field config for _ecs_prefix_v6_ field.
	EcsPrefixV6 int64 `json:"ecs_prefix_v6,omitempty"`

	// Optional. Field config for _ecs_zones_ field.
	EcsZones []*ConfigECSZone `json:"ecs_zones"`
}

// Validate validates this config e c s block
func (m *ConfigECSBlock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEcsZones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigECSBlock) validateEcsZones(formats strfmt.Registry) error {
	if swag.IsZero(m.EcsZones) { // not required
		return nil
	}

	for i := 0; i < len(m.EcsZones); i++ {
		if swag.IsZero(m.EcsZones[i]) { // not required
			continue
		}

		if m.EcsZones[i] != nil {
			if err := m.EcsZones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ecs_zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ecs_zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this config e c s block based on the context it is used
func (m *ConfigECSBlock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEcsZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigECSBlock) contextValidateEcsZones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EcsZones); i++ {

		if m.EcsZones[i] != nil {
			if err := m.EcsZones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ecs_zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ecs_zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigECSBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigECSBlock) UnmarshalBinary(b []byte) error {
	var res ConfigECSBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}