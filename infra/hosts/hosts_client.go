// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new hosts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hosts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	HostsCreate(params *HostsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsCreateCreated, error)

	HostsDelete(params *HostsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsDeleteNoContent, error)

	HostsDisconnect(params *HostsDisconnectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsDisconnectOK, error)

	HostsList(params *HostsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsListOK, error)

	HostsRead(params *HostsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsReadOK, error)

	HostsReplace(params *HostsReplaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsReplaceCreated, error)

	HostsUpdate(params *HostsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsUpdateCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	HostsCreate creates a host resource

	Validation:

- "display_name" is required and should be unique.
*/
func (a *Client) HostsCreate(params *HostsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HostsCreate",
		Method:             "POST",
		PathPattern:        "/hosts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HostsCreateReader{formats: a.formats},
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
	success, ok := result.(*HostsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HostsCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	HostsDelete deletes a host resource

	Validation:

- "id" is required.
*/
func (a *Client) HostsDelete(params *HostsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HostsDelete",
		Method:             "DELETE",
		PathPattern:        "/hosts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HostsDeleteReader{formats: a.formats},
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
	success, ok := result.(*HostsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HostsDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HostsDisconnect disconnects a host by resource ID

The user can disconnect the host from the cloud (for example, if in case a host is compromised).
*/
func (a *Client) HostsDisconnect(params *HostsDisconnectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsDisconnectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsDisconnectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HostsDisconnect",
		Method:             "POST",
		PathPattern:        "/hosts/{id}/disconnect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HostsDisconnectReader{formats: a.formats},
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
	success, ok := result.(*HostsDisconnectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HostsDisconnect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HostsList lists all the host resources for an account
*/
func (a *Client) HostsList(params *HostsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HostsList",
		Method:             "GET",
		PathPattern:        "/hosts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HostsListReader{formats: a.formats},
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
	success, ok := result.(*HostsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HostsList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	HostsRead gets a host resource

	Validation:

- "id" is required.
*/
func (a *Client) HostsRead(params *HostsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HostsRead",
		Method:             "GET",
		PathPattern:        "/hosts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HostsReadReader{formats: a.formats},
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
	success, ok := result.(*HostsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HostsRead: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HostsReplace migrates a host s configuration from one to another
*/
func (a *Client) HostsReplace(params *HostsReplaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsReplaceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsReplaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HostsReplace",
		Method:             "POST",
		PathPattern:        "/hosts/{from.resource_id}/replace/{to.resource_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HostsReplaceReader{formats: a.formats},
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
	success, ok := result.(*HostsReplaceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HostsReplace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	HostsUpdate updates a host resource

	Validation:

- "id" is required.
- "display_name" is required and should be unique.
- "pool_id" is required.
*/
func (a *Client) HostsUpdate(params *HostsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HostsUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HostsUpdate",
		Method:             "PUT",
		PathPattern:        "/hosts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HostsUpdateReader{formats: a.formats},
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
	success, ok := result.(*HostsUpdateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HostsUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
