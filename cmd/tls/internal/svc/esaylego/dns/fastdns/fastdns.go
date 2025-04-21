package fastdns

import (
	"fmt"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/go-acme/lego/providers/dns/fastdns"
)

func NewDNSProviderByValues(values map[string]string) (*fastdns.DNSProvider, error) {
	if values["AKAMAI_HOST"] == "" || values["AKAMAI_CLIENT_TOKEN"] == "" || values["AKAMAI_CLIENT_SECRET"] == "" || values["AKAMAI_ACCESS_TOKEN"] == "" {
		return nil, fmt.Errorf("fastdns: AKAMAI_HOST ar AKAMAI_CLIENT_TOKEN ar AKAMAI_CLIENT_SECRET or AKAMAI_ACCESS_TOKEN is nil")
	}

	config := fastdns.NewDefaultConfig()
	config.Config = edgegrid.Config{
		Host:         values["AKAMAI_HOST"],
		ClientToken:  values["AKAMAI_CLIENT_TOKEN"],
		ClientSecret: values["AKAMAI_CLIENT_SECRET"],
		AccessToken:  values["AKAMAI_ACCESS_TOKEN"],
		MaxBody:      131072,
	}

	return fastdns.NewDNSProviderConfig(config)
}
