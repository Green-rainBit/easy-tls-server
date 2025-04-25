package ovh

import (
	"github.com/go-acme/lego/v4/providers/dns/ovh"
)

func NewDNSProviderByValues(values map[string]string) (*ovh.DNSProvider, error) {
	config := ovh.NewDefaultConfig()

	// https://github.com/ovh/go-ovh/blob/6817886d12a8c5650794b28da635af9fcdfd1162/ovh/configuration.go#L105
	config.APIEndpoint = "ovh-eu"
	if values[ovh.EnvEndpoint] != "" {
		config.APIEndpoint = values[ovh.EnvEndpoint]
	}

	config.ApplicationKey = values[ovh.EnvApplicationKey]
	config.ApplicationSecret = values[ovh.EnvApplicationSecret]
	config.ConsumerKey = values[ovh.EnvConsumerKey]

	config.AccessToken = values[ovh.EnvAccessToken]

	clientID := values[ovh.EnvClientID]
	clientSecret := values[ovh.EnvClientSecret]

	if clientID != "" || clientSecret != "" {
		config.OAuth2Config = &ovh.OAuth2Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}
	}

	return ovh.NewDNSProviderConfig(config)
}
