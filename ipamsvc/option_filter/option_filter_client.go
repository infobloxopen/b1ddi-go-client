// Code generated by go-swagger; DO NOT EDIT.

package option_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new option filter API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for option filter API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	OptionFilterCreate(params *OptionFilterCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OptionFilterCreateCreated, error)

	OptionFilterDelete(params *OptionFilterDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OptionFilterDeleteNoContent, error)

	OptionFilterList(params *OptionFilterListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OptionFilterListOK, error)

	OptionFilterRead(params *OptionFilterReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OptionFilterReadOK, error)

	OptionFilterUpdate(params *OptionFilterUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OptionFilterUpdateCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  OptionFilterCreate creates the d h c p option filter

  Use this method to create an __OptionFilter__ object.
The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.
*/
func (a *Client) OptionFilterCreate(params *OptionFilterCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OptionFilterCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionFilterCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "option_filterCreate",
		Method:             "POST",
		PathPattern:        "/dhcp/option_filter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OptionFilterCreateReader{formats: a.formats},
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
	success, ok := result.(*OptionFilterCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for option_filterCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  OptionFilterDelete moves the d h c p option filter to the recycle bin

  Use this method to move an __OptionFilter__ object to the recycle bin.
The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.
*/
func (a *Client) OptionFilterDelete(params *OptionFilterDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OptionFilterDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionFilterDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "option_filterDelete",
		Method:             "DELETE",
		PathPattern:        "/dhcp/option_filter/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OptionFilterDeleteReader{formats: a.formats},
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
	success, ok := result.(*OptionFilterDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for option_filterDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  OptionFilterList retrieves d h c p option filters

  Use this method to retrieve __OptionFilter__ objects.
The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.
*/
func (a *Client) OptionFilterList(params *OptionFilterListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OptionFilterListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionFilterListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "option_filterList",
		Method:             "GET",
		PathPattern:        "/dhcp/option_filter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OptionFilterListReader{formats: a.formats},
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
	success, ok := result.(*OptionFilterListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for option_filterList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  OptionFilterRead retrieves the d h c p option filter

  Use this method to retrieve an __OptionFilter__ object.
The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.
*/
func (a *Client) OptionFilterRead(params *OptionFilterReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OptionFilterReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionFilterReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "option_filterRead",
		Method:             "GET",
		PathPattern:        "/dhcp/option_filter/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OptionFilterReadReader{formats: a.formats},
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
	success, ok := result.(*OptionFilterReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for option_filterRead: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  OptionFilterUpdate updates the d h c p option filter

  Use this method to update an __OptionFilter__ object.
The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.
*/
func (a *Client) OptionFilterUpdate(params *OptionFilterUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OptionFilterUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionFilterUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "option_filterUpdate",
		Method:             "PATCH",
		PathPattern:        "/dhcp/option_filter/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OptionFilterUpdateReader{formats: a.formats},
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
	success, ok := result.(*OptionFilterUpdateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for option_filterUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}