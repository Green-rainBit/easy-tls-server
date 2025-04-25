package gandi

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/gandi"
)

func NewDNSProviderByValues(values map[string]string) (*gandi.DNSProvider, error) {

	if _, ok := values[gandi.EnvAPIKey]; !ok {
		return nil, fmt.Errorf("gandi: GANDI_API_KEY is nil")
	}

	config := gandi.NewDefaultConfig()
	config.APIKey = values[gandi.EnvAPIKey]

	return gandi.NewDNSProviderConfig(config)
}
