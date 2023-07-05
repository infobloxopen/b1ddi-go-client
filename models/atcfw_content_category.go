// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AtcfwContentCategory The Content Category object.
//
// The Content Category object represents a specific internet content and used to configure category filters. Based on your configuration, specific actions such as Allow or Block, will be taken on the detected content. BloxOne Cloud provides the following content categories from which you can build your category filters: Drugs, Risk/Fraud/Crime, Entertainment/Culture, Purchasing, Information/Communication, Business/Services, Information Technology, Lifestyle, Society/Education/Religion, Mature/Violent, Games/Gambling, Pornography/Nudity and Uncategorized. Each of these categories contains sub-categories that further define the respective content. When you configure your category filter, you can add as many categories and sub-categories as you need.
//
// swagger:model atcfwContentCategory
type AtcfwContentCategory struct {

	// The category code.
	// Example: 155
	CategoryCode int32 `json:"category_code,omitempty"`

	// The name of the category.
	// Example: Weapons
	CategoryName string `json:"category_name,omitempty"`

	// The functional group name of the category.
	// Example: Mature/Violent
	FunctionalGroup string `json:"functional_group,omitempty"`
}

// Validate validates this atcfw content category
func (m *AtcfwContentCategory) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this atcfw content category based on context it is used
func (m *AtcfwContentCategory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AtcfwContentCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AtcfwContentCategory) UnmarshalBinary(b []byte) error {
	var res AtcfwContentCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}