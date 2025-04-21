package exoscale

import (
	"fmt"

	egoscale "github.com/exoscale/egoscale/v3"
	"github.com/go-acme/lego/v4/providers/dns/exoscale"
)

func NewDNSProviderByValues(values map[string]string) (*exoscale.DNSProvider, error) {
	if values[exoscale.EnvAPIKey] == "" || values[exoscale.EnvAPISecret] == "" {
		return nil, fmt.Errorf("exoscale: EXOSCALE_API_KEY or EXOSCALE_API_SECRET is nil")
	}

	config := exoscale.NewDefaultConfig()
	config.APIKey = values[exoscale.EnvAPIKey]
	config.APISecret = values[exoscale.EnvAPISecret]
	config.Endpoint = string(egoscale.CHGva2)
	if values[exoscale.EnvEndpoint] != "" {
		config.Endpoint = values[exoscale.EnvEndpoint]
	}

	return exoscale.NewDNSProviderConfig(config)
}
