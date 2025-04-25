package ns1

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/ns1"
)

func NewDNSProviderByValues(values map[string]string) (*ns1.DNSProvider, error) {
	if values[ns1.EnvAPIKey] == "" {
		return nil, fmt.Errorf("ns1: NS1_API_KEY is nil")
	}

	config := ns1.NewDefaultConfig()
	config.APIKey = values[ns1.EnvAPIKey]

	return ns1.NewDNSProviderConfig(config)
}
