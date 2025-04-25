package easydns

import (
	"fmt"
	"net/url"

	"github.com/go-acme/lego/v4/providers/dns/easydns"
)

const DefaultBaseURL = "https://rest.easydns.net"

func NewDNSProviderByValues(values map[string]string) (*easydns.DNSProvider, error) {
	config := easydns.NewDefaultConfig()
	defaultBaseURL := DefaultBaseURL
	if values[easydns.EnvEndpoint] != "" {
		defaultBaseURL = values[easydns.EnvEndpoint]
	}
	endpoint, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, fmt.Errorf("easydns: %w", err)
	}
	config.Endpoint = endpoint

	if values[easydns.EnvToken] == "" || values[easydns.EnvKey] == "" {
		return nil, fmt.Errorf("easydns: EASYDNS_TOKEN or EASYDNS_KEY is nil")
	}

	config.Token = values[easydns.EnvToken]
	config.Key = values[easydns.EnvKey]

	return easydns.NewDNSProviderConfig(config)
}
