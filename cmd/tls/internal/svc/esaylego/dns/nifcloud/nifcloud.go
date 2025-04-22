package nifcloud

import (
	"fmt"

	"github.com/go-acme/lego/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/nifcloud"
)

func NewDNSProviderByValues(values map[string]string) (*nifcloud.DNSProvider, error) {
	if values[nifcloud.EnvAccessKeyID] == "" || values[nifcloud.EnvSecretAccessKey] == "" {
		return nil, fmt.Errorf("nifcloud: NIFCLOUD_ACCESS_KEY_ID or NIFCLOUD_SECRET_ACCESS_KEY is nil")
	}

	config := nifcloud.NewDefaultConfig()
	config.BaseURL = env.GetOrFile(nifcloud.EnvDNSEndpoint)
	config.AccessKey = values[nifcloud.EnvAccessKeyID]
	config.SecretKey = values[nifcloud.EnvSecretAccessKey]

	return nifcloud.NewDNSProviderConfig(config)
}
