package namecheap

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/namecheap"
)

func NewDNSProviderByValues(values map[string]string) (*namecheap.DNSProvider, error) {
	if values[namecheap.EnvAPIUser] == "" || values[namecheap.EnvAPIKey] == "" {
		return nil, fmt.Errorf("namecheap: NAMECHEAP_API_USER or NAMECHEAP_API_KEY is nil")
	}

	config := namecheap.NewDefaultConfig()
	config.APIUser = values[namecheap.EnvAPIUser]
	config.APIKey = values[namecheap.EnvAPIKey]

	return namecheap.NewDNSProviderConfig(config)
}
