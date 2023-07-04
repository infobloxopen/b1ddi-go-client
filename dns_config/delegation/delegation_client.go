// Code generated by go-swagger; DO NOT EDIT.

package delegation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// New creates a new delegation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for delegation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DelegationCreate(params *DelegationCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelegationCreateCreated, error)

	DelegationDelete(params *DelegationDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelegationDeleteNoContent, error)

	DelegationList(params *DelegationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelegationListOK, error)

	DelegationRead(params *DelegationReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelegationReadOK, error)

	DelegationUpdate(params *DelegationUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelegationUpdateCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	DelegationCreate creates the delegation object

	Use this method to create a Delegation object.

This object (_dns/delegation_) represents a zone delegation.
*/
func (a *Client) DelegationCreate(params *DelegationCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelegationCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDelegationCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delegationCreate",
		Method:             "POST",
		PathPattern:        "/dns/delegation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DelegationCreateReader{formats: a.formats},
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
	success, ok := result.(*DelegationCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delegationCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	DelegationDelete moves the delegation object to recyclebin

	Use this method to move a Delegation object to Recyclebin.

This object (_dns/delegation_) represents a zone delegation.
*/
func (a *Client) DelegationDelete(params *DelegationDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelegationDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDelegationDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delegationDelete",
		Method:             "DELETE",
		PathPattern:        "/dns/delegation/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DelegationDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	params.ID = b1cliruntime.TrimIDPrefix(op.PathPattern, params.ID)
	op.Params = params

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DelegationDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delegationDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	DelegationList lists delegation objects

	Use this method to list Delegation objects.

This object (_dns/delegation_) represents a zone delegation.
*/
func (a *Client) DelegationList(params *DelegationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelegationListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDelegationListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delegationList",
		Method:             "GET",
		PathPattern:        "/dns/delegation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DelegationListReader{formats: a.formats},
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
	success, ok := result.(*DelegationListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delegationList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	DelegationRead reads the delegation object

	Use this method to read a Delegation object.

This object (_dns/delegation)_ represents a zone delegation.
*/
func (a *Client) DelegationRead(params *DelegationReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelegationReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDelegationReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delegationRead",
		Method:             "GET",
		PathPattern:        "/dns/delegation/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DelegationReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	params.ID = b1cliruntime.TrimIDPrefix(op.PathPattern, params.ID)
	op.Params = params

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DelegationReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delegationRead: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	DelegationUpdate updates the delegation object

	Use this method to update a Delegation object.

This object (_dns/delegation_) represents a zone delegation.
*/
func (a *Client) DelegationUpdate(params *DelegationUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelegationUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDelegationUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delegationUpdate",
		Method:             "PATCH",
		PathPattern:        "/dns/delegation/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DelegationUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	params.ID = b1cliruntime.TrimIDPrefix(op.PathPattern, params.ID)
	op.Params = params

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DelegationUpdateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delegationUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
