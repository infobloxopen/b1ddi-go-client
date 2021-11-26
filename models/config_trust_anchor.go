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

// ConfigTrustAnchor TrustAnchor
//
// DNSSEC trust anchor.
//
// swagger:model configTrustAnchor
type ConfigTrustAnchor struct {

	// Key algorithm.
	// Algorithm values are as per standards.
	// The mapping is as follows:
	//  * _RSAMD5_ = 1,
	//  * _DH_ = 2,
	//  * _DSA_ = 3,
	//  * _RSASHA1_ = 5,
	//  * _DSANSEC3SHA1_ = 6,
	//  * _RSASHA1NSEC3SHA1_ = 7,
	//  * _RSASHA256_ = 8,
	//  * _RSASHA512_ = 10,
	//  * _ECDSAP256SHA256_ = 13,
	//  * _ECDSAP384SHA384_ = 14.
	// Below algorithms are deprecated and not supported anymore
	//  * _RSAMD5_ = 1,
	//  * _DSA_ = 3,
	//  * _DSANSEC3SHA1_ = 6,
	// Required: true
	Algorithm *int64 `json:"algorithm"`

	// Zone FQDN in punycode.
	// Read Only: true
	ProtocolZone string `json:"protocol_zone,omitempty"`

	// DNSSEC key data. Non-empty, valid base64 string.
	// Required: true
	PublicKey *string `json:"public_key"`

	// Optional. Secure Entry Point flag.
	//
	// Defaults to _true_.
	Sep bool `json:"sep,omitempty"`

	// Zone FQDN.
	// Required: true
	Zone *string `json:"zone"`
}

// Validate validates this config trust anchor
func (m *ConfigTrustAnchor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigTrustAnchor) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *ConfigTrustAnchor) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("public_key", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

func (m *ConfigTrustAnchor) validateZone(formats strfmt.Registry) error {

	if err := validate.Required("zone", "body", m.Zone); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this config trust anchor based on the context it is used
func (m *ConfigTrustAnchor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtocolZone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigTrustAnchor) contextValidateProtocolZone(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol_zone", "body", string(m.ProtocolZone)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigTrustAnchor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigTrustAnchor) UnmarshalBinary(b []byte) error {
	var res ConfigTrustAnchor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
