package sakuracloud

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/sakuracloud"
)

func NewDNSProviderByValues(values map[string]string) (*sakuracloud.DNSProvider, error) {
	if values[sakuracloud.EnvAccessToken] == "" || values[sakuracloud.EnvAccessTokenSecret] == "" {
		return nil, fmt.Errorf("sakuracloud: SAKURACLOUD_ACCESS_TOKEN or SAKURACLOUD_ACCESS_TOKEN_SECRET is nil")
	}

	config := sakuracloud.NewDefaultConfig()
	config.Token = values[sakuracloud.EnvAccessToken]
	config.Secret = values[sakuracloud.EnvAccessTokenSecret]

	return sakuracloud.NewDNSProviderConfig(config)
}
