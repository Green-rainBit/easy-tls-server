package vegadns

import (
	"fmt"

	"github.com/go-acme/lego/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/vegadns"
)

func NewDNSProviderByValues(values map[string]string) (*vegadns.DNSProvider, error) {
	if values[vegadns.EnvURL] == "" {
		return nil, fmt.Errorf("vegadns: VEGADNS_URL is nil")
	}

	config := vegadns.NewDefaultConfig()
	config.BaseURL = values[vegadns.EnvURL]
	config.APIKey = env.GetOrFile(vegadns.EnvKey)
	config.APISecret = env.GetOrFile(vegadns.EnvSecret)

	return vegadns.NewDNSProviderConfig(config)
}
