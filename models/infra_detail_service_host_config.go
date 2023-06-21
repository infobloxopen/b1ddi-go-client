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

// InfraDetailServiceHostConfig DetailServiceHostConfig
//
// swagger:model infraDetailServiceHostConfig
type InfraDetailServiceHostConfig struct {

	// The current version of the Service deployed on the Host.
	CurrentVersion string `json:"current_version,omitempty"`

	// Service status information.
	Status *InfraShortServiceStatus `json:"status,omitempty"`

	// upgraded at
	// Format: date-time
	UpgradedAt strfmt.DateTime `json:"upgraded_at,omitempty"`
}

// Validate validates this infra detail service host config
func (m *InfraDetailServiceHostConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfraDetailServiceHostConfig) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *InfraDetailServiceHostConfig) validateUpgradedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpgradedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("upgraded_at", "body", "date-time", m.UpgradedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this infra detail service host config based on the context it is used
func (m *InfraDetailServiceHostConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfraDetailServiceHostConfig) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfraDetailServiceHostConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraDetailServiceHostConfig) UnmarshalBinary(b []byte) error {
	var res InfraDetailServiceHostConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}