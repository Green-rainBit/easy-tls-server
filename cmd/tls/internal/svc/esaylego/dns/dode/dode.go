package dode

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/dode"
)

func NewDNSProviderByValues(values map[string]string) (*dode.DNSProvider, error) {

	if values[dode.EnvToken] == "" {
		return nil, fmt.Errorf("do.de: DODE_TOKEN is nil")
	}

	config := dode.NewDefaultConfig()
	config.Token = values[dode.EnvToken]

	return dode.NewDNSProviderConfig(config)
}
