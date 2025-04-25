package godaddy

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/godaddy"
)

func NewDNSProviderByValues(values map[string]string) (*godaddy.DNSProvider, error) {
	if values[godaddy.EnvAPIKey] == "" || values[godaddy.EnvAPISecret] == "" {
		return nil, fmt.Errorf("godaddy: GODADDY_API_KEY or GODADDY_API_SECRET is nil")
	}

	config := godaddy.NewDefaultConfig()
	config.APIKey = values[godaddy.EnvAPIKey]
	config.APISecret = values[godaddy.EnvAPISecret]

	return godaddy.NewDNSProviderConfig(config)
}
