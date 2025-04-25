package digitalocean

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/digitalocean"
)

func NewDNSProviderByValues(values map[string]string) (*digitalocean.DNSProvider, error) {
	_, ok := values[digitalocean.EnvAuthToken]
	if !ok {
		return nil, fmt.Errorf("digitalocean: DO_AUTH_TOKEN is nil")
	}

	config := digitalocean.NewDefaultConfig()
	config.AuthToken = values[digitalocean.EnvAuthToken]

	return digitalocean.NewDNSProviderConfig(config)
}
