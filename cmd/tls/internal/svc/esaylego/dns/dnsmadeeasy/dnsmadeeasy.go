package dnsmadeeasy

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/dnsmadeeasy"
)

func NewDNSProviderByValues(values map[string]string) (*dnsmadeeasy.DNSProvider, error) {
	if values[dnsmadeeasy.EnvAPIKey] == "" || values[dnsmadeeasy.EnvAPISecret] == "" {
		return nil, fmt.Errorf("dnsmadeeasy: DNSMADEEASY_API_KEY or DNSMADEEASY_API_SECRET is nil ")
	}

	config := dnsmadeeasy.NewDefaultConfig()

	if sandbox, ok := values[dnsmadeeasy.EnvSandbox]; !ok {
		config.Sandbox = sandbox == "true"
	}
	config.APIKey = values[dnsmadeeasy.EnvAPIKey]
	config.APISecret = values[dnsmadeeasy.EnvAPISecret]

	return dnsmadeeasy.NewDNSProviderConfig(config)
}
