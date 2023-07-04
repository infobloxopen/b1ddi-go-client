//go:build e2e
// +build e2e

// Code generated by go-swagger; DO NOT EDIT.

package e2e_tests

import (
	"os"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	b1cli "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config"
)

var cli dns_config.DNSConfigurationAPI

func TestMain(m *testing.M) {
	host := os.Getenv("B1DDI_HOST")
	if host == "" {
		panic("B1DDI_HOST value is empty")
	}
	token := os.Getenv("B1DDI_API_KEY")
	if token == "" {
		panic("B1DDI_API_KEY token value is empty")
	}

	// Create new go-swagger runtime client
	transport := httptransport.New(
		host, "api/ddi/v1", nil,
	)

	// Create default auth header for all API requests
	tokenAuth := b1cli.B1DDIAPIKey(token)
	transport.DefaultAuthentication = tokenAuth

	// Create the BloxOne DNS configuration API client
	cli = *dns_config.New(transport, strfmt.Default)

	os.Exit(m.Run())
}
