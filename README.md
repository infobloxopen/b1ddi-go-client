# Overview

This library enables the management of BloxOne DDI resources.

The following BloxOne DDI APIs are supported:

- [IP Address Management (Ipamsvc API)](https://csp.infoblox.com/apidoc/?url=https://csp.infoblox.com/apidoc/docs/Ipamsvc)
- [DNS Configuration (DNSConfig API)](https://csp.infoblox.com/apidoc/?url=https://csp.infoblox.com/apidoc/docs/DnsConfig)
- [DNS Data (DnsData API)](https://csp.infoblox.com/apidoc/?url=https://csp.infoblox.com/apidoc/docs/DnsData)
- [Bloxone Cloud Infrastructure Management](https://csp.infoblox.com/apidoc?url=https:/csp.infoblox.com/apidoc/docs/Infrastructure)

# Installation

To install `b1ddi-go-client` use `go get` command:

```bash
go get github.com/infobloxopen/b1ddi-go-client
```

# Usage Guide

## Examples

The following program will print the subnet mask of each subnet in the BloxOne DDI environment:
```go
package main

import (
	"fmt"
	"os"

	"github.com/go-openapi/strfmt"
	b1cli "github.com/infobloxopen/b1ddi-go-client/client"
)

func main() {
	// Create the BloxOne API client
	//client := b1cli.NewClient(os.Getenv("B1DDI_HOST"), os.Getenv("B1DDI_API_KEY"), strfmt.Default)
	client := b1cli.NewClient("csp.infoblox.com", "b91338436bdf58bdf7c148cf35afe8cd3044f988fe1370133bf94f1a4906940f", strfmt.Default)

	// List all subnets using IPAM API client
	subnetList, err := client.IPAddressManagementAPI.Subnet.SubnetList(nil, nil)
	if err != nil {
		panic(err)
	}

	// Print subnet mask for each subnet
	for _, subnet := range subnetList.Payload.Results {
		fmt.Printf("%s/%d\n", *subnet.Address, subnet.Cidr)
	}

	// Get all hosts using Infrastructure Management client
	hostList, err := client.InfrastructureManagementAPI.Hosts.HostsList(nil, nil)
	if err != nil {
		panic(err)
	}

	// Print subnet mask for each subnet
	for _, host := range hostList.Payload.Results {
		fmt.Printf("%s\n", *host.DisplayName)
	}

	// Get all internal domain lists using b1td_cloud client
	domainList, err := client.B1TDCloudAPI.InternalDomainLists.InternalDomainListsListInternalDomains(nil, nil)
	if err != nil {
		panic(err)
	}

	for _, list := range domainList.GetPayload().Results {
		fmt.Printf("%s\n", list.Name)
	}

}
```
