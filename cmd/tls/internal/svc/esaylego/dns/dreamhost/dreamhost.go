package dreamhost

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/dreamhost"
)

func NewDNSProviderByValues(values map[string]string) (*dreamhost.DNSProvider, error) {
	if values[dreamhost.EnvAPIKey] == "" {
		return nil, fmt.Errorf("dreamhost: DREAMHOST_API_KEY is nil")
	}

	config := dreamhost.NewDefaultConfig()
	config.APIKey = values[dreamhost.EnvAPIKey]

	return dreamhost.NewDNSProviderConfig(config)
}
