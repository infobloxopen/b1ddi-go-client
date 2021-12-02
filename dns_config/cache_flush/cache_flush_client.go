// Code generated by go-swagger; DO NOT EDIT.

package cache_flush

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cache flush API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cache flush API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CacheFlushCreate(params *CacheFlushCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CacheFlushCreateCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CacheFlushCreate creates the cache flush object

  Use this method to create a Cache Flush object.
The Cache Flush object is for removing entries from the DNS cache on a host. The host must be available and running DNS for this to succeed.
*/
func (a *Client) CacheFlushCreate(params *CacheFlushCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CacheFlushCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCacheFlushCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cache_flushCreate",
		Method:             "POST",
		PathPattern:        "/dns/cache_flush",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CacheFlushCreateReader{formats: a.formats},
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
	success, ok := result.(*CacheFlushCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cache_flushCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}