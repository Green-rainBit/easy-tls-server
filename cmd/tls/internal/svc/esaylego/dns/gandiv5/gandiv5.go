package gandiv5

import (
	"github.com/go-acme/lego/v4/providers/dns/gandiv5"
)

func NewDNSProviderByValues(values map[string]string) (*gandiv5.DNSProvider, error) {
	// TODO(ldez): rewrite this when APIKey will be removed.
	config := gandiv5.NewDefaultConfig()
	config.APIKey = values[gandiv5.EnvAPIKey]
	config.PersonalAccessToken = values[gandiv5.EnvPersonalAccessToken]

	return gandiv5.NewDNSProviderConfig(config)
}
