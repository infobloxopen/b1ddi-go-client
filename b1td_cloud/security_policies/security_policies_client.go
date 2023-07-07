// Code generated by go-swagger; DO NOT EDIT.

package security_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new security policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for security policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SecurityPoliciesCreateSecurityPolicy(params *SecurityPoliciesCreateSecurityPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesCreateSecurityPolicyCreated, error)

	SecurityPoliciesDeleteSecurityPolicy(params *SecurityPoliciesDeleteSecurityPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesDeleteSecurityPolicyNoContent, error)

	SecurityPoliciesDeleteSingleSecurityPolicy(params *SecurityPoliciesDeleteSingleSecurityPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesDeleteSingleSecurityPolicyNoContent, error)

	SecurityPoliciesListSecurityPolicies(params *SecurityPoliciesListSecurityPoliciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesListSecurityPoliciesOK, error)

	SecurityPoliciesReadSecurityPolicy(params *SecurityPoliciesReadSecurityPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesReadSecurityPolicyOK, error)

	SecurityPoliciesUpdateSecurityPolicy(params *SecurityPoliciesUpdateSecurityPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesUpdateSecurityPolicyCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	SecurityPoliciesCreateSecurityPolicy creates security policy

	Use this method to create a Security Policy object. If no rule list is specified, the newly created Security Policy object will create these rules as a copy of default Security Policy rules ("Default Global Policy"). If rule list is provided it must contain at least the complete list of policy rules, including the rules based on all feeds that the account is entitled to. If no network list is specified, networking scope for this policy is empty, thus no action can be performed by this policy. Note that you are not allowed to add DNS Forwarding Proxies and Roaming Device Groups that are already assigned to a Security Policy different from "Default Global Policy", to assign them to this Security Policy object you should remove them from other Security Policy first.

A security policy defines a set of rules and actions that you define to balance access and constraints so you can mitigate malicious attacks and provide security for your networks. When you create a new security policy, you first define a network scope to which you add networks, DNS forwarding proxies, and BloxOne Endpoint groups. BloxOne Cloud applies the security policy to all the entities that you include in the network scope. You can also include DNS forwarding proxies to which you want to apply the security policy.  After you define the network scope, you can add custom lists and category filters to the security policy. You can also specify actions for the added lists and filters, and to determine the precedence order for the entities. Depending on your subscription level, each security policy also comes with a set of predefined threat intelligence feeds and Threat Insight rules that are inherited from the default global policy. You cannot delete the inherited feeds and rules, but you can change their precedence order. For each policy you can define policy-level action (Default Action) gets applied when none of the policy rules apply/match. If an user really needs access to some blocked domain (web page) via a browser - there is a possibility to assign special bypass code(s) (Bypass Code) to any policy.

Required:
- name
*/
func (a *Client) SecurityPoliciesCreateSecurityPolicy(params *SecurityPoliciesCreateSecurityPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesCreateSecurityPolicyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecurityPoliciesCreateSecurityPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "security_policiesCreateSecurityPolicy",
		Method:             "POST",
		PathPattern:        "/security_policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecurityPoliciesCreateSecurityPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecurityPoliciesCreateSecurityPolicyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for security_policiesCreateSecurityPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SecurityPoliciesDeleteSecurityPolicy deletes security policies

	Use this method to delete Security Policy objects. Deletion of multiple lists is an all-or-nothing operation (if any of the specified lists can not be deleted then none of the specified lists will be deleted).

A security policy defines a set of rules and actions that you define to balance access and constraints so you can mitigate malicious attacks and provide security for your networks. When you create a new security policy, you first define a network scope to which you add networks, DNS forwarding proxies, and BloxOne Endpoint groups. BloxOne Cloud applies the security policy to all the entities that you include in the network scope. You can also include DNS forwarding proxies to which you want to apply the security policy.  After you define the network scope, you can add custom lists and category filters to the security policy. You can also specify actions for the added lists and filters, and to determine the precedence order for the entities. Depending on your subscription level, each security policy also comes with a set of predefined threat intelligence feeds and Threat Insight rules that are inherited from the default global policy. You cannot delete the inherited feeds and rules, but you can change their precedence order. For each policy you can define policy-level action (Default Action) gets applied when none of the policy rules apply/match. If an user really needs access to some blocked domain (web page) via a browser - there is a possibility to assign special bypass code(s) (Bypass Code) to any policy.

Required:
- ids
*/
func (a *Client) SecurityPoliciesDeleteSecurityPolicy(params *SecurityPoliciesDeleteSecurityPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesDeleteSecurityPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecurityPoliciesDeleteSecurityPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "security_policiesDeleteSecurityPolicy",
		Method:             "DELETE",
		PathPattern:        "/security_policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecurityPoliciesDeleteSecurityPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecurityPoliciesDeleteSecurityPolicyNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for security_policiesDeleteSecurityPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SecurityPoliciesDeleteSingleSecurityPolicy deletes security policy

	Use this method to delete Security Policy object by given Security Policy object id.

A security policy defines a set of rules and actions that you define to balance access and constraints so you can mitigate malicious attacks and provide security for your networks. When you create a new security policy, you first define a network scope to which you add networks, DNS forwarding proxies, and BloxOne Endpoint groups. BloxOne Cloud applies the security policy to all the entities that you include in the network scope. You can also include DNS forwarding proxies to which you want to apply the security policy.  After you define the network scope, you can add custom lists and category filters to the security policy. You can also specify actions for the added lists and filters, and to determine the precedence order for the entities. Depending on your subscription level, each security policy also comes with a set of predefined threat intelligence feeds and Threat Insight rules that are inherited from the default global policy. You cannot delete the inherited feeds and rules, but you can change their precedence order. For each policy you can define policy-level action (Default Action) gets applied when none of the policy rules apply/match. If an user really needs access to some blocked domain (web page) via a browser - there is a possibility to assign special bypass code(s) (Bypass Code) to any policy.
*/
func (a *Client) SecurityPoliciesDeleteSingleSecurityPolicy(params *SecurityPoliciesDeleteSingleSecurityPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesDeleteSingleSecurityPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecurityPoliciesDeleteSingleSecurityPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "security_policiesDeleteSingleSecurityPolicy",
		Method:             "DELETE",
		PathPattern:        "/security_policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecurityPoliciesDeleteSingleSecurityPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecurityPoliciesDeleteSingleSecurityPolicyNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for security_policiesDeleteSingleSecurityPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SecurityPoliciesListSecurityPolicies lists security policies

	Use this method to retrieve information on all Security Policy objects for the account.

A security policy defines a set of rules and actions that you define to balance access and constraints so you can mitigate malicious attacks and provide security for your networks. When you create a new security policy, you first define a network scope to which you add networks, DNS forwarding proxies, and BloxOne Endpoint groups. BloxOne Cloud applies the security policy to all the entities that you include in the network scope. You can also include DNS forwarding proxies to which you want to apply the security policy.  After you define the network scope, you can add custom lists and category filters to the security policy. You can also specify actions for the added lists and filters, and to determine the precedence order for the entities. Depending on your subscription level, each security policy also comes with a set of predefined threat intelligence feeds and Threat Insight rules that are inherited from the default global policy. You cannot delete the inherited feeds and rules, but you can change their precedence order. For each policy you can define policy-level action (Default Action) gets applied when none of the policy rules apply/match. If an user really needs access to some blocked domain (web page) via a browser - there is a possibility to assign special bypass code(s) (Bypass Code) to any policy.
*/
func (a *Client) SecurityPoliciesListSecurityPolicies(params *SecurityPoliciesListSecurityPoliciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesListSecurityPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecurityPoliciesListSecurityPoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "security_policiesListSecurityPolicies",
		Method:             "GET",
		PathPattern:        "/security_policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecurityPoliciesListSecurityPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecurityPoliciesListSecurityPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for security_policiesListSecurityPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SecurityPoliciesReadSecurityPolicy reads security policy

	Use this method to retrieve information on the specified Security Policy object.

A security policy defines a set of rules and actions that you define to balance access and constraints so you can mitigate malicious attacks and provide security for your networks. When you create a new security policy, you first define a network scope to which you add networks, DNS forwarding proxies, and BloxOne Endpoint groups. BloxOne Cloud applies the security policy to all the entities that you include in the network scope. You can also include DNS forwarding proxies to which you want to apply the security policy.  After you define the network scope, you can add custom lists and category filters to the security policy. You can also specify actions for the added lists and filters, and to determine the precedence order for the entities. Depending on your subscription level, each security policy also comes with a set of predefined threat intelligence feeds and Threat Insight rules that are inherited from the default global policy. You cannot delete the inherited feeds and rules, but you can change their precedence order. For each policy you can define policy-level action (Default Action) gets applied when none of the policy rules apply/match. If an user really needs access to some blocked domain (web page) via a browser - there is a possibility to assign special bypass code(s) (Bypass Code) to any policy.
*/
func (a *Client) SecurityPoliciesReadSecurityPolicy(params *SecurityPoliciesReadSecurityPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesReadSecurityPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecurityPoliciesReadSecurityPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "security_policiesReadSecurityPolicy",
		Method:             "GET",
		PathPattern:        "/security_policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecurityPoliciesReadSecurityPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecurityPoliciesReadSecurityPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for security_policiesReadSecurityPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SecurityPoliciesUpdateSecurityPolicy updates security policy

	Use this method to update a specified Network List object. The policy data supplied with the update request must have the complete list of policy rules, including the rules based on all feeds that the account is entitled to. If no network list is specified, networking scope for this policy is empty, thus no action can be performed by this policy. Note that you are not allowed to add DNS Forwarding Proxies and Roaming Device Groups that are already assigned to a Security Policy different from "Default Global Policy", to assign them to this Security Policy object you should remove them from other Security Policy first.

A security policy defines a set of rules and actions that you define to balance access and constraints so you can mitigate malicious attacks and provide security for your networks. When you create a new security policy, you first define a network scope to which you add networks, DNS forwarding proxies, and BloxOne Endpoint groups. BloxOne Cloud applies the security policy to all the entities that you include in the network scope. You can also include DNS forwarding proxies to which you want to apply the security policy.  After you define the network scope, you can add custom lists and category filters to the security policy. You can also specify actions for the added lists and filters, and to determine the precedence order for the entities. Depending on your subscription level, each security policy also comes with a set of predefined threat intelligence feeds and Threat Insight rules that are inherited from the default global policy. You cannot delete the inherited feeds and rules, but you can change their precedence order. For each policy you can define policy-level action (Default Action) gets applied when none of the policy rules apply/match. If an user really needs access to some blocked domain (web page) via a browser - there is a possibility to assign special bypass code(s) (Bypass Code) to any policy.

Required:
- name
- rules
- dfps
- network_lists
- roaming_device_groups
*/
func (a *Client) SecurityPoliciesUpdateSecurityPolicy(params *SecurityPoliciesUpdateSecurityPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SecurityPoliciesUpdateSecurityPolicyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecurityPoliciesUpdateSecurityPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "security_policiesUpdateSecurityPolicy",
		Method:             "PUT",
		PathPattern:        "/security_policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecurityPoliciesUpdateSecurityPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecurityPoliciesUpdateSecurityPolicyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for security_policiesUpdateSecurityPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
