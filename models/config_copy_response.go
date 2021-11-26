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

// ConfigCopyResponse config copy response
//
// swagger:model configCopyResponse
type ConfigCopyResponse struct {

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// An Unique Id to identify copy operation.
	JobID string `json:"job_id,omitempty"`
}

// Validate validates this config copy response
func (m *ConfigCopyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this config copy response based on the context it is used
func (m *ConfigCopyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigCopyResponse) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigCopyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigCopyResponse) UnmarshalBinary(b []byte) error {
	var res ConfigCopyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
