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

// DataSOASerialIncrementRequest data s o a serial increment request
//
// swagger:model dataSOASerialIncrementRequest
type DataSOASerialIncrementRequest struct {

	// fields
	Fields *ProtobufFieldMask `json:"fields,omitempty"`

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Default increment SOA serial number by 1,000.
	SerialIncrement int64 `json:"serial_increment,omitempty"`
}

// Validate validates this data s o a serial increment request
func (m *DataSOASerialIncrementRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSOASerialIncrementRequest) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	if m.Fields != nil {
		if err := m.Fields.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fields")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this data s o a serial increment request based on the context it is used
func (m *DataSOASerialIncrementRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSOASerialIncrementRequest) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	if m.Fields != nil {
		if err := m.Fields.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fields")
			}
			return err
		}
	}

	return nil
}

func (m *DataSOASerialIncrementRequest) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataSOASerialIncrementRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSOASerialIncrementRequest) UnmarshalBinary(b []byte) error {
	var res DataSOASerialIncrementRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
