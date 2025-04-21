package dnsmadeeasy

import (
	"fmt"

	"github.com/go-acme/lego/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/dnsmadeeasy"
)

func NewDNSProviderByValues(values map[string]string) (*dnsmadeeasy.DNSProvider, error) {
	env.Get(dnsmadeeasy.EnvAPIKey, dnsmadeeasy.EnvAPISecret)
	_, ok := values[dnsmadeeasy.EnvAPIKey]
	if !ok {
		return nil, fmt.Errorf("dnsmadeeasy: EnvAPIKey is nil ")
	}

	config := dnsmadeeasy.NewDefaultConfig()

	if sandbox, ok := values[dnsmadeeasy.EnvSandbox]; !ok {
		config.Sandbox = sandbox == "true"
	}
	config.APIKey = values[dnsmadeeasy.EnvAPIKey]
	config.APISecret = values[dnsmadeeasy.EnvAPISecret]

	return dnsmadeeasy.NewDNSProviderConfig(config)
}
