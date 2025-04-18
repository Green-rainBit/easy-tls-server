package alidns

import (
	"github.com/go-acme/lego/v4/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/alidns"
)

func NewDNSProviderByValues(values map[string]string) (*alidns.DNSProvider, error) {

	config := alidns.NewDefaultConfig()
	config.APIKey = values["ALICLOUD_ACCESS_KEY"]
	config.SecretKey = values["ALICLOUD_SECRET_KEY"]
	config.RegionID = env.GetOrFile("ALICLOUD_REGION_ID")

	return alidns.NewDNSProviderConfig(config)
}
