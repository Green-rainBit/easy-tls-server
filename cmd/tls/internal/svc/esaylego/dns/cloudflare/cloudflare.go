package cloudflare

import (
	"github.com/go-acme/lego/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/cloudflare"
)

func NewDNSProviderByValues(values map[string]string) (*cloudflare.DNSProvider, error) {

	config := cloudflare.NewDefaultConfig()
	config.AuthEmail = values[cloudflare.EnvEmail]
	config.AuthKey = values[cloudflare.EnvAPIKey]
	config.AuthToken = values[cloudflare.EnvDNSAPIToken]
	config.ZoneToken = values[cloudflare.EnvZoneAPIToken]
	config.BaseURL = env.GetOrFile(cloudflare.EnvBaseURL)
	if values[cloudflare.EnvBaseURL] != "" {
		config.BaseURL = values[cloudflare.EnvBaseURL]
	}
	return cloudflare.NewDNSProviderConfig(config)
}
