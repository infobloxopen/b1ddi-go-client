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

// AtcfwNetworkListReadResponse The Network List read response.
//
// swagger:model atcfwNetworkListReadResponse
type AtcfwNetworkListReadResponse struct {

	// The Network List object.
	Results *AtcfwNetworkList `json:"results,omitempty"`
}

// Validate validates this atcfw network list read response
func (m *AtcfwNetworkListReadResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwNetworkListReadResponse) validateResults(formats strfmt.Registry) error {
	if swag.IsZero(m.Results) { // not required
		return nil
	}

	if m.Results != nil {
		if err := m.Results.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this atcfw network list read response based on the context it is used
func (m *AtcfwNetworkListReadResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwNetworkListReadResponse) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	if m.Results != nil {

		if swag.IsZero(m.Results) { // not required
			return nil
		}

		if err := m.Results.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AtcfwNetworkListReadResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AtcfwNetworkListReadResponse) UnmarshalBinary(b []byte) error {
	var res AtcfwNetworkListReadResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}