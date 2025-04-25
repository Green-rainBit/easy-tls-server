package inwx

import (
	"fmt"

	"github.com/go-acme/lego/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/inwx"
)

func NewDNSProviderByValues(values map[string]string) (*inwx.DNSProvider, error) {
	values, err := env.Get(inwx.EnvUsername, inwx.EnvPassword)
	if err != nil {
		return nil, fmt.Errorf("inwx: %w", err)
	}

	if values[inwx.EnvUsername] == "" || values[inwx.EnvPassword] == "" {
		return nil, fmt.Errorf("inwx: INWX_USERNAME or INWX_PASSWORD is nil")
	}
	config := inwx.NewDefaultConfig()
	config.Username = values[inwx.EnvUsername]
	config.Password = values[inwx.EnvPassword]
	config.SharedSecret = env.GetOrFile(inwx.EnvSharedSecret)

	return inwx.NewDNSProviderConfig(config)
}
