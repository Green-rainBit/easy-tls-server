package glesys

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/glesys"
)

func NewDNSProviderByValues(values map[string]string) (*glesys.DNSProvider, error) {

	if values[glesys.EnvAPIUser] == "" || values[glesys.EnvAPIKey] == "" {
		return nil, fmt.Errorf("glesys: GLESYS_API_USER or GLESYS_API_KEY is nil")
	}

	config := glesys.NewDefaultConfig()
	config.APIUser = values[glesys.EnvAPIUser]
	config.APIKey = values[glesys.EnvAPIKey]

	return glesys.NewDNSProviderConfig(config)
}
