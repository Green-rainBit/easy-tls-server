package pdns

import (
	"fmt"
	"net/url"

	"github.com/go-acme/lego/v4/providers/dns/pdns"
)

func NewDNSProviderByValues(values map[string]string) (*pdns.DNSProvider, error) {
	if values[pdns.EnvAPIKey] == "" || values[pdns.EnvAPIURL] == "" {
		return nil, fmt.Errorf("pdns: PDNS_API_KEY or PDNS_API_URL is nil")
	}

	hostURL, err := url.Parse(values[pdns.EnvAPIURL])
	if err != nil {
		return nil, fmt.Errorf("pdns: %w", err)
	}

	config := pdns.NewDefaultConfig()
	config.Host = hostURL
	config.APIKey = values[pdns.EnvAPIKey]

	return pdns.NewDNSProviderConfig(config)
}
