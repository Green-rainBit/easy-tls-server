package alidns

import (
	"fmt"

	"github.com/go-acme/lego/v4/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/alidns"
)

func NewDNSProviderByValues(values map[string]string) (*alidns.DNSProvider, error) {

	config := alidns.NewDefaultConfig()
	config.RegionID = values[alidns.EnvRegionID]

	if values[alidns.EnvRAMRole] != "" {
		config.RAMRole = values[alidns.EnvRAMRole]
		return alidns.NewDNSProviderConfig(config)
	}
	if values[alidns.EnvAccessKey] == "" || values[alidns.EnvSecretKey] == "" {
		return nil, fmt.Errorf("alicloud: ALICLOUD_ACCESS_KEY or ALICLOUD_SECRET_KEY is nil")
	}

	config.APIKey = values[alidns.EnvAccessKey]
	config.SecretKey = values[alidns.EnvSecretKey]
	config.SecurityToken = env.GetOrFile(alidns.EnvSecurityToken)

	return alidns.NewDNSProviderConfig(config)
}
