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

// AtcfwThreatFeed The Threat Feed object.
//
// BloxOne Cloud provides predefined threat intelligence feeds based on your subscription. The Plus subscription offers a few more feeds than the Standard subscription. The Advanced subscription offers a few more feeds than the Plus subscription. A threat feed subscription for RPZ updates offers protection against malicious hostnames.
//
// swagger:model atcfwThreatFeed
type AtcfwThreatFeed struct {

	// The Confidence Level of the Feed.
	// Example: LOW
	// Read Only: true
	ConfidenceLevel string `json:"confidence_level,omitempty"`

	// The brief description for the thread feed.
	// Example: Threat Feed Description
	// Read Only: true
	Description string `json:"description,omitempty"`

	// The TSIG key of the threat feed.
	// Example: etiqrisk
	// Read Only: true
	Key string `json:"key,omitempty"`

	// The name of the thread feed.
	// Example: ETIQRisk
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The source of the threat feed.
	// Example: Infoblox
	// Read Only: true
	Source *ThreatFeedSource `json:"source,omitempty"`

	// The Threat Level of the Feed.
	// Example: LOW
	// Read Only: true
	ThreatLevel string `json:"threat_level,omitempty"`
}

// Validate validates this atcfw threat feed
func (m *AtcfwThreatFeed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwThreatFeed) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this atcfw threat feed based on the context it is used
func (m *AtcfwThreatFeed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfidenceLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThreatLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwThreatFeed) contextValidateConfidenceLevel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "confidence_level", "body", string(m.ConfidenceLevel)); err != nil {
		return err
	}

	return nil
}

func (m *AtcfwThreatFeed) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *AtcfwThreatFeed) contextValidateKey(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "key", "body", string(m.Key)); err != nil {
		return err
	}

	return nil
}

func (m *AtcfwThreatFeed) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *AtcfwThreatFeed) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if swag.IsZero(m.Source) { // not required
			return nil
		}

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *AtcfwThreatFeed) contextValidateThreatLevel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "threat_level", "body", string(m.ThreatLevel)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AtcfwThreatFeed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AtcfwThreatFeed) UnmarshalBinary(b []byte) error {
	var res AtcfwThreatFeed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}