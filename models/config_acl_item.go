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

// ConfigACLItem ACLItem
//
// Element in an ACL.
//
// Error if both _acl_ and _address_ are given.
//
// swagger:model configACLItem
type ConfigACLItem struct {

	// Access permission for _element_.
	//
	// Allowed values:
	//  * _allow_,
	//  * _deny_.
	// Required: true
	Access *string `json:"access"`

	// The resource identifier.
	ACL string `json:"acl,omitempty"`

	// Optional. Data for _ip_ _element_.
	//
	// Must be empty if _element_ is not _ip_.
	Address string `json:"address,omitempty"`

	// Type of element.
	//
	// Allowed values:
	//  * _any_,
	//  * _ip_,
	//  * _acl_,
	//  * _tsig_key_.
	// Required: true
	Element *string `json:"element"`

	// Optional. TSIG key.
	//
	// Must be empty if _element_ is not _tsig_key_.
	TsigKey *ConfigTSIGKey `json:"tsig_key,omitempty"`
}

// Validate validates this config ACL item
func (m *ConfigACLItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTsigKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigACLItem) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", m.Access); err != nil {
		return err
	}

	return nil
}

func (m *ConfigACLItem) validateElement(formats strfmt.Registry) error {

	if err := validate.Required("element", "body", m.Element); err != nil {
		return err
	}

	return nil
}

func (m *ConfigACLItem) validateTsigKey(formats strfmt.Registry) error {
	if swag.IsZero(m.TsigKey) { // not required
		return nil
	}

	if m.TsigKey != nil {
		if err := m.TsigKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tsig_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tsig_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config ACL item based on the context it is used
func (m *ConfigACLItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTsigKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigACLItem) contextValidateTsigKey(ctx context.Context, formats strfmt.Registry) error {

	if m.TsigKey != nil {

		if swag.IsZero(m.TsigKey) { // not required
			return nil
		}

		if err := m.TsigKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tsig_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tsig_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigACLItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigACLItem) UnmarshalBinary(b []byte) error {
	var res ConfigACLItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
