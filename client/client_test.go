package client

import (
	"fmt"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	// transport is created inside the client now, passing in the host as the first parameter
	cli := NewClient("", "", strfmt.Default)

	assert.NotNil(t, cli)
	assert.NotNil(t, cli.IPAddressManagementAPI)
	assert.NotNil(t, cli.DNSConfigurationAPI)
	assert.NotNil(t, cli.DNSDataAPI)
	assert.NotNil(t, cli.InfrastructureManagementAPI)
}

func TestB1DDIAPIKey(t *testing.T) {
	key := "KxPlKBne1NtqkcgaUgiU29cME9Y0I13JtZ8QPyDLTalv0yXLriJZgXF4lDXYV31Ky"

	authFunc := B1DDIAPIKey(key)

	req := &runtime.TestClientRequest{}

	err := authFunc.AuthenticateRequest(req, strfmt.Default)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(
		t,
		fmt.Sprintf("Token %s", key),
		req.Headers["Authorization"][0],
	)
}
