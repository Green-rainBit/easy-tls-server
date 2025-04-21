package dnspod

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/dnspod"
)

func NewDNSProviderByValues(values map[string]string) (*dnspod.DNSProvider, error) {
	if values[dnspod.EnvAPIKey] == "" {
		return nil, fmt.Errorf("dnspod: DNSPOD_API_KEY is nil")
	}

	config := dnspod.NewDefaultConfig()
	config.LoginToken = values[dnspod.EnvAPIKey]

	return dnspod.NewDNSProviderConfig(config)
}
