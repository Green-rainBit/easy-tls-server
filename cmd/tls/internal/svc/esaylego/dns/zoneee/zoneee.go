package zoneee

import (
	"fmt"
	"net/url"

	"github.com/go-acme/lego/v4/providers/dns/zoneee"
)

const DefaultEndpoint = "https://api.zone.eu/v2/"

func NewDNSProviderByValues(values map[string]string) (*zoneee.DNSProvider, error) {
	if values[zoneee.EnvAPIUser] == "" || values[zoneee.EnvAPIKey] == "" {
		return nil, fmt.Errorf("zoneee: ZONEEE_API_USER or ZONEEE_API_KEY is nil")
	}

	rawEndpoint := DefaultEndpoint
	if values[zoneee.EnvEndpoint] != "" {
		rawEndpoint = values[zoneee.EnvEndpoint]
	}
	endpoint, err := url.Parse(rawEndpoint)
	if err != nil {
		return nil, fmt.Errorf("zoneee: %w", err)
	}

	config := zoneee.NewDefaultConfig()
	config.Username = values[zoneee.EnvAPIUser]
	config.APIKey = values[zoneee.EnvAPIKey]
	config.Endpoint = endpoint

	return zoneee.NewDNSProviderConfig(config)
}
