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
)

// AtcfwListPoPRegionsResponse atcfw list po p regions response
//
// swagger:model atcfwListPoPRegionsResponse
type AtcfwListPoPRegionsResponse struct {

	// results
	Results []*AtcfwPoPRegion `json:"results"`

	// total result count
	TotalResultCount int32 `json:"total_result_count,omitempty"`
}

// Validate validates this atcfw list po p regions response
func (m *AtcfwListPoPRegionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwListPoPRegionsResponse) validateResults(formats strfmt.Registry) error {
	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this atcfw list po p regions response based on the context it is used
func (m *AtcfwListPoPRegionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwListPoPRegionsResponse) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Results); i++ {

		if m.Results[i] != nil {

			if swag.IsZero(m.Results[i]) { // not required
				return nil
			}

			if err := m.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AtcfwListPoPRegionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AtcfwListPoPRegionsResponse) UnmarshalBinary(b []byte) error {
	var res AtcfwListPoPRegionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
