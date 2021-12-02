// Code generated by go-swagger; DO NOT EDIT.

package view

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new view API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for view API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ViewBulkCopy(params *ViewBulkCopyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewBulkCopyCreated, error)

	ViewCreate(params *ViewCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewCreateCreated, error)

	ViewDelete(params *ViewDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewDeleteNoContent, error)

	ViewList(params *ViewListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewListOK, error)

	ViewRead(params *ViewReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewReadOK, error)

	ViewUpdate(params *ViewUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewUpdateCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ViewBulkCopy copies the specified auth zone and forward zone objects in the view

  Use this method to bulk copy __AuthZone__ and __ForwardZone__ objects from one __View__ object to another __View__ object.
The __AuthZone__ object represents an authoritative zone.
The __ForwardZone__ object represents a forwarding zone.
*/
func (a *Client) ViewBulkCopy(params *ViewBulkCopyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewBulkCopyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewViewBulkCopyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "viewBulkCopy",
		Method:             "POST",
		PathPattern:        "/dns/view/bulk_copy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ViewBulkCopyReader{formats: a.formats},
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
	success, ok := result.(*ViewBulkCopyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for viewBulkCopy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ViewCreate creates the view object

  Use this method to create a View object.
Named collection of DNS View settings.
*/
func (a *Client) ViewCreate(params *ViewCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewViewCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "viewCreate",
		Method:             "POST",
		PathPattern:        "/dns/view",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ViewCreateReader{formats: a.formats},
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
	success, ok := result.(*ViewCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for viewCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ViewDelete moves the view object to recyclebin

  Use this method to move a View object to Recyclebin.
Named collection of DNS View settings.
*/
func (a *Client) ViewDelete(params *ViewDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewViewDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "viewDelete",
		Method:             "DELETE",
		PathPattern:        "/dns/view/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ViewDeleteReader{formats: a.formats},
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
	success, ok := result.(*ViewDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for viewDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ViewList lists view objects

  Use this method to list View objects.
Named collection of DNS View settings.
*/
func (a *Client) ViewList(params *ViewListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewViewListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "viewList",
		Method:             "GET",
		PathPattern:        "/dns/view",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ViewListReader{formats: a.formats},
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
	success, ok := result.(*ViewListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for viewList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ViewRead reads the view object

  Use this method to read a View object.
Named collection of DNS View settings.
*/
func (a *Client) ViewRead(params *ViewReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewViewReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "viewRead",
		Method:             "GET",
		PathPattern:        "/dns/view/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ViewReadReader{formats: a.formats},
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
	success, ok := result.(*ViewReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for viewRead: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ViewUpdate updates the view object

  Use this method to update a View object.
Named collection of DNS View settings.
*/
func (a *Client) ViewUpdate(params *ViewUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ViewUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewViewUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "viewUpdate",
		Method:             "PATCH",
		PathPattern:        "/dns/view/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ViewUpdateReader{formats: a.formats},
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
	success, ok := result.(*ViewUpdateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for viewUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}