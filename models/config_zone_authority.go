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

// ConfigZoneAuthority ZoneAuthority
//
// Construct for fields: _refresh_, _retry_, _expire_, _default_ttl_, _negative_ttl_, _rname_, _protocol_rname_, _mname_, _protocol_mname_, _use_default_mname_.
//
// swagger:model configZoneAuthority
type ConfigZoneAuthority struct {

	// Optional. ZoneAuthority default ttl for resource records in zone (value in seconds).
	//
	// Defaults to 28800.
	DefaultTTL int64 `json:"default_ttl,omitempty"`

	// Optional. ZoneAuthority expire time in seconds.
	//
	// Defaults to 2419200.
	Expire int64 `json:"expire,omitempty"`

	// Optional. ZoneAuthority master name server (partially qualified domain name)
	//
	// Defaults to empty.
	Mname string `json:"mname,omitempty"`

	// Optional. ZoneAuthority negative caching (minimum) ttl in seconds.
	//
	// Defaults to 900.
	NegativeTTL int64 `json:"negative_ttl,omitempty"`

	// Optional. ZoneAuthority master name server in punycode.
	//
	// Defaults to empty.
	// Read Only: true
	ProtocolMname string `json:"protocol_mname,omitempty"`

	// Optional. A domain name which specifies the mailbox of the person responsible for this zone.
	//
	// Defaults to empty.
	// Read Only: true
	ProtocolRname string `json:"protocol_rname,omitempty"`

	// Optional. ZoneAuthority refresh.
	//
	// Defaults to 10800.
	Refresh int64 `json:"refresh,omitempty"`

	// Optional. ZoneAuthority retry.
	//
	// Defaults to 3600.
	Retry int64 `json:"retry,omitempty"`

	// Optional. ZoneAuthority rname.
	//
	// Defaults to empty.
	Rname string `json:"rname,omitempty"`

	// Optional. Use default value for master name server.
	//
	// Defaults to true.
	UseDefaultMname *bool `json:"use_default_mname,omitempty"`
}

// Validate validates this config zone authority
func (m *ConfigZoneAuthority) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this config zone authority based on the context it is used
func (m *ConfigZoneAuthority) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtocolMname(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocolRname(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigZoneAuthority) contextValidateProtocolMname(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol_mname", "body", string(m.ProtocolMname)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigZoneAuthority) contextValidateProtocolRname(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol_rname", "body", string(m.ProtocolRname)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigZoneAuthority) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigZoneAuthority) UnmarshalBinary(b []byte) error {
	var res ConfigZoneAuthority
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
