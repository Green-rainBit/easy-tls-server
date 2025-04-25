package rackspace

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/rackspace"
)

func NewDNSProviderByValues(values map[string]string) (*rackspace.DNSProvider, error) {
	if values[rackspace.EnvUser] == "" || values[rackspace.EnvAPIKey] == "" {
		return nil, fmt.Errorf("rackspace: RACKSPACE_USER or RACKSPACE_API_KEY is nil")
	}

	config := rackspace.NewDefaultConfig()
	config.APIUser = values[rackspace.EnvUser]
	config.APIKey = values[rackspace.EnvAPIKey]

	return rackspace.NewDNSProviderConfig(config)
}
