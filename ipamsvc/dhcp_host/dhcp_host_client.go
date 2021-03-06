// Code generated by go-swagger; DO NOT EDIT.

package dhcp_host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// New creates a new dhcp host API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dhcp host API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DhcpHostList(params *DhcpHostListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DhcpHostListOK, error)

	DhcpHostListAssociations(params *DhcpHostListAssociationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DhcpHostListAssociationsOK, error)

	DhcpHostRead(params *DhcpHostReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DhcpHostReadOK, error)

	DhcpHostUpdate(params *DhcpHostUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DhcpHostUpdateCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DhcpHostList retrieves d h c p hosts

  Use this method to retrieve DHCP __Host__ objects.
A DHCP __Host__ object associates a __DHCPConfigProfile__ object with an on-prem host.
*/
func (a *Client) DhcpHostList(params *DhcpHostListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DhcpHostListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDhcpHostListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dhcp_hostList",
		Method:             "GET",
		PathPattern:        "/dhcp/host",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DhcpHostListReader{formats: a.formats},
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
	success, ok := result.(*DhcpHostListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dhcp_hostList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DhcpHostListAssociations retrieves d h c p host associations

  Use this method to retrieve __HostAssociation__ objects.
*/
func (a *Client) DhcpHostListAssociations(params *DhcpHostListAssociationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DhcpHostListAssociationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDhcpHostListAssociationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dhcp_hostListAssociations",
		Method:             "GET",
		PathPattern:        "/dhcp/host/{id}/associations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DhcpHostListAssociationsReader{formats: a.formats},
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
	success, ok := result.(*DhcpHostListAssociationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dhcp_hostListAssociations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DhcpHostRead retrieves the d h c p host

  Use this method to retrieve a DHCP Host object.
A DHCP __Host__ object associates a __DHCPConfigProfile__ object with an on-prem host.
*/
func (a *Client) DhcpHostRead(params *DhcpHostReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DhcpHostReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDhcpHostReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dhcp_hostRead",
		Method:             "GET",
		PathPattern:        "/dhcp/host/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DhcpHostReadReader{formats: a.formats},
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
	success, ok := result.(*DhcpHostReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dhcp_hostRead: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DhcpHostUpdate updates the d h c p hosts

  Use this method to update a DHCP __Host__ object.
A DHCP __Host__ object associates a __DHCPConfigProfile__ object with an on-prem host.
*/
func (a *Client) DhcpHostUpdate(params *DhcpHostUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DhcpHostUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDhcpHostUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dhcp_hostUpdate",
		Method:             "PATCH",
		PathPattern:        "/dhcp/host/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DhcpHostUpdateReader{formats: a.formats},
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
	success, ok := result.(*DhcpHostUpdateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dhcp_hostUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
