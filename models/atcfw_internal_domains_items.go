// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AtcfwInternalDomainsItems atcfw internal domains items
//
// swagger:model atcfwInternalDomainsItems
type AtcfwInternalDomainsItems struct {

	// The List of ItemStructs structure which contains the item and its description
	// Example: [{"description":"Item 1 Description","item":"example2.somedomain.com"},{"description":"Item 2 Description","item":"193.56.2.10/32"},{"description":"Item 3 Description","item":"193.56.2.9"}]
	DeletedItemsDescribed []*AtcfwItemStructs `json:"deleted_items_described"`

	// The Internal Domain List object identifier.
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The List of ItemStructs structure which contains the item and its description
	// Example: [{"description":"Item 1 Description","item":"example1.somedomain.com"},{"description":"Item 2 Description","item":"193.56.2.11/32"},{"description":"Item 3 Description","item":"193.56.2.11"}]
	InsertedItemsDescribed []*AtcfwItemStructs `json:"inserted_items_described"`
}

// Validate validates this atcfw internal domains items
func (m *AtcfwInternalDomainsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeletedItemsDescribed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInsertedItemsDescribed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwInternalDomainsItems) validateDeletedItemsDescribed(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletedItemsDescribed) { // not required
		return nil
	}

	for i := 0; i < len(m.DeletedItemsDescribed); i++ {
		if swag.IsZero(m.DeletedItemsDescribed[i]) { // not required
			continue
		}

		if m.DeletedItemsDescribed[i] != nil {
			if err := m.DeletedItemsDescribed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleted_items_described" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deleted_items_described" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AtcfwInternalDomainsItems) validateInsertedItemsDescribed(formats strfmt.Registry) error {
	if swag.IsZero(m.InsertedItemsDescribed) { // not required
		return nil
	}

	for i := 0; i < len(m.InsertedItemsDescribed); i++ {
		if swag.IsZero(m.InsertedItemsDescribed[i]) { // not required
			continue
		}

		if m.InsertedItemsDescribed[i] != nil {
			if err := m.InsertedItemsDescribed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inserted_items_described" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inserted_items_described" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this atcfw internal domains items based on the context it is used
func (m *AtcfwInternalDomainsItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeletedItemsDescribed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInsertedItemsDescribed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwInternalDomainsItems) contextValidateDeletedItemsDescribed(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DeletedItemsDescribed); i++ {

		if m.DeletedItemsDescribed[i] != nil {

			if swag.IsZero(m.DeletedItemsDescribed[i]) { // not required
				return nil
			}

			if err := m.DeletedItemsDescribed[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleted_items_described" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deleted_items_described" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AtcfwInternalDomainsItems) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *AtcfwInternalDomainsItems) contextValidateInsertedItemsDescribed(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InsertedItemsDescribed); i++ {

		if m.InsertedItemsDescribed[i] != nil {

			if swag.IsZero(m.InsertedItemsDescribed[i]) { // not required
				return nil
			}

			if err := m.InsertedItemsDescribed[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inserted_items_described" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inserted_items_described" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AtcfwInternalDomainsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AtcfwInternalDomainsItems) UnmarshalBinary(b []byte) error {
	var res AtcfwInternalDomainsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}