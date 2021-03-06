// Code generated by go-swagger; DO NOT EDIT.

package ipam_host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// New creates a new ipam host API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ipam host API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	IpamHostCreate(params *IpamHostCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamHostCreateCreated, error)

	IpamHostDelete(params *IpamHostDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamHostDeleteNoContent, error)

	IpamHostList(params *IpamHostListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamHostListOK, error)

	IpamHostRead(params *IpamHostReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamHostReadOK, error)

	IpamHostUpdate(params *IpamHostUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamHostUpdateCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  IpamHostCreate creates the IP a m host

  Use this method to create an __IpamHost__ object.
The __IpamHost__ object (_ipam/host_) represents any network connected equipment that is assigned one or more IP Addresses.
*/
func (a *Client) IpamHostCreate(params *IpamHostCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamHostCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamHostCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_hostCreate",
		Method:             "POST",
		PathPattern:        "/ipam/host",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IpamHostCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamHostCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_hostCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamHostDelete moves the IP a m host to the recycle bin

  Use this method to move an __IpamHost__ object to the recycle bin.
The __IpamHost__ object (_ipam/host_) represents any network connected equipment that is assigned one or more IP Addresses.
*/
func (a *Client) IpamHostDelete(params *IpamHostDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamHostDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamHostDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_hostDelete",
		Method:             "DELETE",
		PathPattern:        "/ipam/host/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IpamHostDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamHostDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_hostDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamHostList retrieves the IP a m hosts

  Use this method to retrieve __IpamHost__ objects.
The __IpamHost__ object (_ipam/host_) represents any network connected equipment that is assigned one or more IP Addresses.
*/
func (a *Client) IpamHostList(params *IpamHostListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamHostListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamHostListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_hostList",
		Method:             "GET",
		PathPattern:        "/ipam/host",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IpamHostListReader{formats: a.formats},
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
	success, ok := result.(*IpamHostListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_hostList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamHostRead retrieves the IP a m host

  Use this method to retrieve an __IpamHost__ object.
The __IpamHost__ object (_ipam/host_) represents any network connected equipment that is assigned one or more IP Addresses.
*/
func (a *Client) IpamHostRead(params *IpamHostReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamHostReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamHostReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_hostRead",
		Method:             "GET",
		PathPattern:        "/ipam/host/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IpamHostReadReader{formats: a.formats},
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
	success, ok := result.(*IpamHostReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_hostRead: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamHostUpdate updates the IP a m host

  Use this method to update an __IpamHost__ object.
The __IpamHost__ object (_ipam/host_) represents any network connected equipment that is assigned one or more IP Addresses.
*/
func (a *Client) IpamHostUpdate(params *IpamHostUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamHostUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamHostUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_hostUpdate",
		Method:             "PATCH",
		PathPattern:        "/ipam/host/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IpamHostUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamHostUpdateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_hostUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
