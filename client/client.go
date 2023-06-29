// Package client provides useful primitives for working with BloxOne DDI APIs
package client

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/infobloxopen/b1ddi-go-client/dns_config"
	"github.com/infobloxopen/b1ddi-go-client/dns_data"
	"github.com/infobloxopen/b1ddi-go-client/infra"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
)

// Client is an aggregation of different BloxOne DDI API clients.
type Client struct {
	IPAddressManagementAPI      *ipamsvc.IPAddressManagementAPI
	DNSConfigurationAPI         *dns_config.DNSConfigurationAPI
	DNSDataAPI                  *dns_data.DNSDataAPI
	InfrastructureManagementAPI *infra.InfrastructureManagementAPI
}

// NewClient creates a new BloxOne DDI API Client.
func NewClient(host string, api_key string, formats strfmt.Registry) *Client {

	// Create Transports via DefaultBasePaths and passed in host
	infratransport := httptransport.New(
		host, infra.DefaultBasePath, nil,
	)
	// ipamsvc, dns_config, and dns_data share the same BasePath
	dditransport := httptransport.New(
		host, ipamsvc.DefaultBasePath, nil,
	)

	// Create default auth header for all API requests
	tokenAuth := B1DDIAPIKey(api_key)
	infratransport.DefaultAuthentication = tokenAuth
	dditransport.DefaultAuthentication = tokenAuth

	return &Client{
		IPAddressManagementAPI:      ipamsvc.New(dditransport, formats),
		DNSConfigurationAPI:         dns_config.New(dditransport, formats),
		DNSDataAPI:                  dns_data.New(dditransport, formats),
		InfrastructureManagementAPI: infra.New(infratransport, formats),
	}
}

// B1DDIAPIKey provides a header for the BloxOne DDI API authentication.
//
// See https://docs.infoblox.com/display/BloxOneDDI/BloxOne+DDI+API+Guide learn how to get the API key.
func B1DDIAPIKey(apiKey string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("Authorization", "Token "+apiKey)
	})
}
