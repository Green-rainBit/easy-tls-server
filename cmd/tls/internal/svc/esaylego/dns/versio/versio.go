package versio

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/versio"
)

func NewDNSProviderByValues(values map[string]string) (*versio.DNSProvider, error) {
	if values[versio.EnvUsername] == "" || values[versio.EnvPassword] == "" {
		return nil, fmt.Errorf("versio: VERSIO_USERNAME or VERSIO_PASSWORD is nil")
	}

	config := versio.NewDefaultConfig()
	config.Username = values[versio.EnvUsername]
	config.Password = values[versio.EnvPassword]

	return versio.NewDNSProviderConfig(config)
}
