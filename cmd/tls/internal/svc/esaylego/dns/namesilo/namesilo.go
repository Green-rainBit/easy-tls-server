package namesilo

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/namesilo"
)

func NewDNSProviderByValues(values map[string]string) (*namesilo.DNSProvider, error) {
	if values[namesilo.EnvAPIKey] == "" {
		return nil, fmt.Errorf("namesilo: NAMESILO_API_KEY is nil")
	}

	config := namesilo.NewDefaultConfig()
	config.APIKey = values[namesilo.EnvAPIKey]

	return namesilo.NewDNSProviderConfig(config)
}
