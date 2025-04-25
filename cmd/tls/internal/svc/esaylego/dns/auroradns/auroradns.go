package auroradns

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/auroradns"
)

func NewDNSProviderByValues(values map[string]string) (*auroradns.DNSProvider, error) {

	if values[auroradns.EnvAPIKey] == "" || values[auroradns.EnvSecret] == "" {
		return nil, fmt.Errorf("aurora EnvAPIKey or AURORA_SECRET is nil")
	}

	config := auroradns.NewDefaultConfig()
	config.BaseURL = values[auroradns.EnvEndpoint]
	config.APIKey = values[auroradns.EnvAPIKey]
	config.Secret = values[auroradns.EnvSecret]

	return auroradns.NewDNSProviderConfig(config)
}
