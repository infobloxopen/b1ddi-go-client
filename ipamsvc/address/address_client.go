// Code generated by go-swagger; DO NOT EDIT.

package address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new address API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for address API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddressCreate(params *AddressCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddressCreateCreated, error)

	AddressDelete(params *AddressDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddressDeleteNoContent, error)

	AddressList(params *AddressListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddressListOK, error)

	AddressRead(params *AddressReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddressReadOK, error)

	AddressUpdate(params *AddressUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddressUpdateCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddressCreate creates the IP address

  Use this method to create an __Address__ object.
The __Address__ object represents any single IP address within a given IP space.
*/
func (a *Client) AddressCreate(params *AddressCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddressCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddressCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addressCreate",
		Method:             "POST",
		PathPattern:        "/ipam/address",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddressCreateReader{formats: a.formats},
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
	success, ok := result.(*AddressCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addressCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddressDelete moves the IP address to the recycle bin

  Use this method to move an __Address__ object to the recycle bin.
The __Address__ object represents any single IP address within a given IP space.
*/
func (a *Client) AddressDelete(params *AddressDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddressDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddressDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addressDelete",
		Method:             "DELETE",
		PathPattern:        "/ipam/address/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddressDeleteReader{formats: a.formats},
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
	success, ok := result.(*AddressDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addressDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddressList retrieves IP addresses

  Use this method to retrieve __Address__ objects.
The __Address__ object represents any single IP address within a given IP space.
*/
func (a *Client) AddressList(params *AddressListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddressListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddressListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addressList",
		Method:             "GET",
		PathPattern:        "/ipam/address",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddressListReader{formats: a.formats},
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
	success, ok := result.(*AddressListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addressList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddressRead retrieves the IP address

  Use this method to retrieve an __Address__ object.
The __Address__ object represents any single IP address within a given IP space.
*/
func (a *Client) AddressRead(params *AddressReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddressReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddressReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addressRead",
		Method:             "GET",
		PathPattern:        "/ipam/address/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddressReadReader{formats: a.formats},
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
	success, ok := result.(*AddressReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addressRead: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddressUpdate updates the IP address

  Use this method to update an __Address__ object.
The __Address__ object represents any single IP address within a given IP space.
*/
func (a *Client) AddressUpdate(params *AddressUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddressUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddressUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addressUpdate",
		Method:             "PATCH",
		PathPattern:        "/ipam/address/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddressUpdateReader{formats: a.formats},
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
	success, ok := result.(*AddressUpdateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addressUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}