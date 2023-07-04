// Code generated by go-swagger; DO NOT EDIT.

package asm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// New creates a new asm API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for asm API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AsmCreate(params *AsmCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AsmCreateCreated, error)

	AsmList(params *AsmListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AsmListOK, error)

	AsmRead(params *AsmReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AsmReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	AsmCreate updates subnet and ranges for automated scope management

	Use this method to update the subnet and range for Automated Scope Management.

The __ASM__ object generates and returns the suggestions from the ASM suggestion engine and allows for updating the subnet and range.
This method attempts to expand the scope by expanding a range or adding a new range and, if necessary, expanding the subnet.
*/
func (a *Client) AsmCreate(params *AsmCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AsmCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAsmCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "asmCreate",
		Method:             "POST",
		PathPattern:        "/ipam/asm",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AsmCreateReader{formats: a.formats},
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
	success, ok := result.(*AsmCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for asmCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AsmList retrieves suggested updates for automated scope management

	Use this method to retrieve __ASM__ objects for Automated Scope Management.

The __ASM__ object returns the suggested updates for the subnet from the ASM suggestion engine and allows for updating the subnet and range information.
*/
func (a *Client) AsmList(params *AsmListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AsmListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAsmListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "asmList",
		Method:             "GET",
		PathPattern:        "/ipam/asm",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AsmListReader{formats: a.formats},
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
	success, ok := result.(*AsmListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for asmList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AsmRead retrieves the suggested update for automated scope management

	Use this method to retrieve an __ASM__ object for Automated Scope Management.

The __ASM__ object returns the suggested updates for the subnet from the ASM suggestion engine and allows for updating the subnet and range information.
*/
func (a *Client) AsmRead(params *AsmReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AsmReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAsmReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "asmRead",
		Method:             "GET",
		PathPattern:        "/ipam/asm/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AsmReadReader{formats: a.formats},
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
	success, ok := result.(*AsmReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for asmRead: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
