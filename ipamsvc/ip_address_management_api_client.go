// Code generated by go-swagger; DO NOT EDIT.

package ipamsvc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address_block"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/asm"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/dhcp_host"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/dns_usage"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/filter"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/fixed_address"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/global"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ha_group"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/hardware_filter"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ip_space"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ipam_host"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/option_code"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/option_filter"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/option_group"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/option_space"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/range_operations"
	serverops "github.com/infobloxopen/b1ddi-go-client/ipamsvc/server"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/subnet"
)

// Default IP address management API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/ddi/v1/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new IP address management API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *IPAddressManagementAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new IP address management API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *IPAddressManagementAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new IP address management API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *IPAddressManagementAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(IPAddressManagementAPI)
	cli.Transport = transport
	cli.Address = address.New(transport, formats)
	cli.AddressBlock = address_block.New(transport, formats)
	cli.Asm = asm.New(transport, formats)
	cli.DhcpHost = dhcp_host.New(transport, formats)
	cli.DNSUsage = dns_usage.New(transport, formats)
	cli.Filter = filter.New(transport, formats)
	cli.FixedAddress = fixed_address.New(transport, formats)
	cli.Global = global.New(transport, formats)
	cli.HaGroup = ha_group.New(transport, formats)
	cli.HardwareFilter = hardware_filter.New(transport, formats)
	cli.IPSpace = ip_space.New(transport, formats)
	cli.IpamHost = ipam_host.New(transport, formats)
	cli.OptionCode = option_code.New(transport, formats)
	cli.OptionFilter = option_filter.New(transport, formats)
	cli.OptionGroup = option_group.New(transport, formats)
	cli.OptionSpace = option_space.New(transport, formats)
	cli.RangeOperations = range_operations.New(transport, formats)
	cli.Server = serverops.New(transport, formats)
	cli.Subnet = subnet.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// IPAddressManagementAPI is a client for IP address management API
type IPAddressManagementAPI struct {
	Address address.ClientService

	AddressBlock address_block.ClientService

	Asm asm.ClientService

	DhcpHost dhcp_host.ClientService

	DNSUsage dns_usage.ClientService

	Filter filter.ClientService

	FixedAddress fixed_address.ClientService

	Global global.ClientService

	HaGroup ha_group.ClientService

	HardwareFilter hardware_filter.ClientService

	IPSpace ip_space.ClientService

	IpamHost ipam_host.ClientService

	OptionCode option_code.ClientService

	OptionFilter option_filter.ClientService

	OptionGroup option_group.ClientService

	OptionSpace option_space.ClientService

	RangeOperations range_operations.ClientService

	Server serverops.ClientService

	Subnet subnet.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *IPAddressManagementAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Address.SetTransport(transport)
	c.AddressBlock.SetTransport(transport)
	c.Asm.SetTransport(transport)
	c.DhcpHost.SetTransport(transport)
	c.DNSUsage.SetTransport(transport)
	c.Filter.SetTransport(transport)
	c.FixedAddress.SetTransport(transport)
	c.Global.SetTransport(transport)
	c.HaGroup.SetTransport(transport)
	c.HardwareFilter.SetTransport(transport)
	c.IPSpace.SetTransport(transport)
	c.IpamHost.SetTransport(transport)
	c.OptionCode.SetTransport(transport)
	c.OptionFilter.SetTransport(transport)
	c.OptionGroup.SetTransport(transport)
	c.OptionSpace.SetTransport(transport)
	c.RangeOperations.SetTransport(transport)
	c.Server.SetTransport(transport)
	c.Subnet.SetTransport(transport)
}
