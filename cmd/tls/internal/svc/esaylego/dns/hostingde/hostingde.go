package hostingde

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/hostingde"
)

func NewDNSProviderByValues(values map[string]string) (*hostingde.DNSProvider, error) {
	if values[hostingde.EnvAPIKey] == "" {
		return nil, fmt.Errorf("hostingde: HOSTINGDE_API_KEY is nil")
	}

	config := hostingde.NewDefaultConfig()
	config.APIKey = values[hostingde.EnvAPIKey]

	return hostingde.NewDNSProviderConfig(config)
}
