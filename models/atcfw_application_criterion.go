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

// AtcfwApplicationCriterion atcfw application criterion
//
// swagger:model atcfwApplicationCriterion
type AtcfwApplicationCriterion struct {

	// category
	Category string `json:"category,omitempty"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Name for the application. Since the name of application is unique it may
	// be used as alternate key for the application. The 'name' is used for
	// import-export workflow and should be resolved to the 'id' before continue
	// processing Create/Update operations.
	Name string `json:"name,omitempty"`

	// subcategory
	Subcategory string `json:"subcategory,omitempty"`
}

// Validate validates this atcfw application criterion
func (m *AtcfwApplicationCriterion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this atcfw application criterion based on the context it is used
func (m *AtcfwApplicationCriterion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwApplicationCriterion) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AtcfwApplicationCriterion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AtcfwApplicationCriterion) UnmarshalBinary(b []byte) error {
	var res AtcfwApplicationCriterion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
