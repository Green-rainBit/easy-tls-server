package duckdns

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/duckdns"
)

func NewDNSProviderByValues(values map[string]string) (*duckdns.DNSProvider, error) {
	if values[duckdns.EnvToken] == "" {
		return nil, fmt.Errorf("duckdns: DUCKDNS_TOKEN is nil")
	}

	config := duckdns.NewDefaultConfig()
	config.Token = values[duckdns.EnvToken]

	return duckdns.NewDNSProviderConfig(config)
}
