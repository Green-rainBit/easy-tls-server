package vscale

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/vscale"
)

func NewDNSProviderByValues(values map[string]string) (*vscale.DNSProvider, error) {
	if values[vscale.EnvAPIToken] == "" {
		return nil, fmt.Errorf("vscale: VSCALE_API_TOKEN is nil")
	}

	config := vscale.NewDefaultConfig()
	config.Token = values[vscale.EnvAPIToken]

	return vscale.NewDNSProviderConfig(config)
}
