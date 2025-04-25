package vultr

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/vultr"
)

func NewDNSProviderByValues(values map[string]string) (*vultr.DNSProvider, error) {
	if values[vultr.EnvAPIKey] == "" {
		return nil, fmt.Errorf("vultr: VULTR_API_KEY is nil")
	}

	config := vultr.NewDefaultConfig()
	config.APIKey = values[vultr.EnvAPIKey]

	return vultr.NewDNSProviderConfig(config)
}
