package namedotcom

import (
	"fmt"

	"github.com/go-acme/lego/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/namedotcom"
)

func NewDNSProviderByValues(values map[string]string) (*namedotcom.DNSProvider, error) {

	if values[namedotcom.EnvUsername] == "" || values[namedotcom.EnvAPIToken] == "" {
		return nil, fmt.Errorf("namedotcom: NAMECOM_USERNAME or NAMECOM_API_TOKEN is nil")
	}

	config := namedotcom.NewDefaultConfig()
	config.Username = values[namedotcom.EnvUsername]
	config.APIToken = values[namedotcom.EnvAPIToken]
	config.Server = env.GetOrFile(namedotcom.EnvServer)

	return namedotcom.NewDNSProviderConfig(config)
}
