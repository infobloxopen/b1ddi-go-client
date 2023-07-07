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

// AtcfwSecurityPolicyRule The Security Policy Rule object.
//
// The Security Policy Rule object represents a rule and action that you define to balance access and constraints so you can mitigate malicious attacks and provide security for your networks.
//
// swagger:model atcfwSecurityPolicyRule
type AtcfwSecurityPolicyRule struct {

	// The action for the policy rule that can be either "action_allow" or "action_log" or "action_redirect" or "action_block" or "action_allow_with_local_resolution".
	// "action_allow_with_local_resolution" only supported for application filter rule with enabled onprem_resolve flag on the Security policy.
	// Example: action_block
	Action string `json:"action,omitempty"`

	// The data source for the policy rule, that can be either a name of the predefined feed for "named_feed", custom list name for "custom_list" type, category filter name for "category_filter" type and application filter name for "application_filter" type.
	// Example: custom_list_a
	Data string `json:"data,omitempty"`

	// The Custom List object identifier with which the policy rule is associated. 0 value means no custom list is associated with this policy rule.
	// Example: 12345
	// Read Only: true
	ListID int32 `json:"list_id,omitempty"`

	// The identifier of the Security Policy object with which the policy rule is associated.
	// Example: 12345
	// Read Only: true
	PolicyID int32 `json:"policy_id,omitempty"`

	// The name of the security policy with which the policy rule is associated.
	// Example: security_policy_a
	PolicyName string `json:"policy_name,omitempty"`

	// The name of the redirect address for redirect actions that can be either IPv4 address or a domain name.
	// Example: redirect_a
	RedirectName string `json:"redirect_name,omitempty"`

	// The policy rule type that can be either "named_feed" or "custom_list" or "category_filter" or "application_filter".
	// Example: custom_list
	Type string `json:"type,omitempty"`
}

// Validate validates this atcfw security policy rule
func (m *AtcfwSecurityPolicyRule) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this atcfw security policy rule based on the context it is used
func (m *AtcfwSecurityPolicyRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateListID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwSecurityPolicyRule) contextValidateListID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "list_id", "body", int32(m.ListID)); err != nil {
		return err
	}

	return nil
}

func (m *AtcfwSecurityPolicyRule) contextValidatePolicyID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "policy_id", "body", int32(m.PolicyID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AtcfwSecurityPolicyRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AtcfwSecurityPolicyRule) UnmarshalBinary(b []byte) error {
	var res AtcfwSecurityPolicyRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
