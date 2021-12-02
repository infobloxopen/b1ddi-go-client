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

// IpamsvcOptionFilterRule OptionFilterRule
//
// An __OptionFilterRule__ object (_dhcp/option_filter_rule_) represents a filter rule to match a DHCP client.
//
// swagger:model ipamsvcOptionFilterRule
type IpamsvcOptionFilterRule struct {

	// Indicates how to compare the _option_value_ to the client option.
	//
	// Success by comparison:
	//  * _equals_: value and client option are the same,
	//  * _not_equals_: value and client option are not the same,
	//  * _exists_: client option exists,
	//  * _not_exists_: client option does not exist,
	//  * _text_substring_: value is the specified substring of the option,
	//  * _not_text_substring_: value is not the specified substring of the option.
	//  * _hex_substring_: value is the specified hexadecimal substring of the option,
	//  * _not_hex_substring_: value is not the specified hexadecimal substring of the option.
	// Required: true
	Compare *string `json:"compare"`

	// The resource identifier.
	// Required: true
	OptionCode *string `json:"option_code"`

	// The value to match against.
	OptionValue string `json:"option_value,omitempty"`

	// The offset where the substring match starts. This is used only if comparing the _option_value_ using any of the substring modes.
	SubstringOffset int64 `json:"substring_offset,omitempty"`
}

// Validate validates this ipamsvc option filter rule
func (m *IpamsvcOptionFilterRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompare(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcOptionFilterRule) validateCompare(formats strfmt.Registry) error {

	if err := validate.Required("compare", "body", m.Compare); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcOptionFilterRule) validateOptionCode(formats strfmt.Registry) error {

	if err := validate.Required("option_code", "body", m.OptionCode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ipamsvc option filter rule based on context it is used
func (m *IpamsvcOptionFilterRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcOptionFilterRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcOptionFilterRule) UnmarshalBinary(b []byte) error {
	var res IpamsvcOptionFilterRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}