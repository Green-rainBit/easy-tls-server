package dnsimple

import (
	"github.com/go-acme/lego/v4/providers/dns/dnsimple"
)

func NewDNSProviderByValues(values map[string]string) (*dnsimple.DNSProvider, error) {
	config := dnsimple.NewDefaultConfig()
	config.AccessToken = values[dnsimple.EnvOAuthToken]
	config.BaseURL = values[dnsimple.EnvBaseURL]

	return dnsimple.NewDNSProviderConfig(config)
}
