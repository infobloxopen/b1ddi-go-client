// Code generated by go-swagger; DO NOT EDIT.

package detail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new detail API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for detail API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DetailHostList(params *DetailHostListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailHostListOK, error)

	DetailServiceList(params *DetailServiceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailServiceListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DetailHostList lists all the hosts along with its associated services applications
*/
func (a *Client) DetailHostList(params *DetailHostListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailHostListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailHostListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DetailHostList",
		Method:             "GET",
		PathPattern:        "/detail_hosts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DetailHostListReader{formats: a.formats},
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
	success, ok := result.(*DetailHostListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DetailHostList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DetailServiceList lists all the services applications along with its associated hosts
*/
func (a *Client) DetailServiceList(params *DetailServiceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DetailServiceList",
		Method:             "GET",
		PathPattern:        "/detail_services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DetailServiceListReader{formats: a.formats},
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
	success, ok := result.(*DetailServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DetailServiceList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}