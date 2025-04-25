package selectel

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/selectel"
)

func NewDNSProviderByValues(values map[string]string) (*selectel.DNSProvider, error) {
	if values[selectel.EnvAPIToken] == "" {
		return nil, fmt.Errorf("selectel: SELECTEL_API_TOKEN is nil")
	}

	config := selectel.NewDefaultConfig()
	config.Token = values[selectel.EnvAPIToken]

	return selectel.NewDNSProviderConfig(config)
}
