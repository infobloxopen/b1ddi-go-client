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

// IpamsvcHostAddress Addresses
//
// The __HostAddress__ object represents addresses associated with a Host object.
//
// swagger:model ipamsvcHostAddress
type IpamsvcHostAddress struct {

	// Field usage depends on the operation:
	//  * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.
	//  * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:
	//     * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.
	//     * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.
	//     * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_.
	// Required: true
	Address *string `json:"address,omitempty"`

	// The resource identifier.
	// Required: true
	Ref *string `json:"ref,omitempty"`

	// The resource identifier.
	// Required: true
	Space *string `json:"space,omitempty"`
}

// Validate validates this ipamsvc host address
func (m *IpamsvcHostAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcHostAddress) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcHostAddress) validateRef(formats strfmt.Registry) error {

	if err := validate.Required("ref", "body", m.Ref); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcHostAddress) validateSpace(formats strfmt.Registry) error {

	if err := validate.Required("space", "body", m.Space); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ipamsvc host address based on context it is used
func (m *IpamsvcHostAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcHostAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcHostAddress) UnmarshalBinary(b []byte) error {
	var res IpamsvcHostAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
