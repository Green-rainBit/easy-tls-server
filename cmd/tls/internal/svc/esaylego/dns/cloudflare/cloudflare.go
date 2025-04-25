package cloudflare

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/cloudflare"
)

const (
	envNamespace    = "CLOUDFLARE_"
	altEnvNamespace = "CF_"

	altEnvEmail = altEnvNamespace + "API_EMAIL"
)

func NewDNSProviderByValues(values map[string]string) (*cloudflare.DNSProvider, error) {
	if (values[cloudflare.EnvEmail] == "" && values[altEnvEmail] == "") || (values[cloudflare.EnvAPIKey] == "" && values[altEnvName(cloudflare.EnvAPIKey)] == "") {
		errT := fmt.Errorf("cloudflare: %v and %v or %v and %v is nil", cloudflare.EnvEmail, altEnvEmail, cloudflare.EnvAPIKey, cloudflare.EnvAPIKey)
		if (values[cloudflare.EnvDNSAPIToken] == "" && values[altEnvName(cloudflare.EnvDNSAPIToken)] == "") ||
			(values[cloudflare.EnvZoneAPIToken] == "" && values[altEnvName(cloudflare.EnvZoneAPIToken)] == "" && values[cloudflare.EnvDNSAPIToken] == "" && values[altEnvName(cloudflare.EnvDNSAPIToken)] == "") {
			err := fmt.Errorf("cloudflare: %v and %v or %v and %v and %v and %v is nil", cloudflare.EnvDNSAPIToken, altEnvName(cloudflare.EnvDNSAPIToken),
				cloudflare.EnvZoneAPIToken, altEnvName(cloudflare.EnvZoneAPIToken), cloudflare.EnvDNSAPIToken, altEnvName(cloudflare.EnvDNSAPIToken))
			return nil, fmt.Errorf("cloudflare: %v or %v", err, errT)
		}
	}

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

func altEnvName(v string) string {
	return strings.ReplaceAll(v, envNamespace, altEnvNamespace)
}
