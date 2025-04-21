package bluecat

import (
	"fmt"

	"github.com/go-acme/lego/providers/dns/bluecat"
)

func NewDNSProviderByValues(values map[string]string) (*bluecat.DNSProvider, error) {
	if _, ok := values["BLUECAT_SERVER_URL"]; !ok {
		return nil, fmt.Errorf("bluecat: BLUECAT_SERVER_URL nil")
	}

	config := bluecat.NewDefaultConfig()
	config.BaseURL = values["BLUECAT_SERVER_URL"]
	config.UserName = values["BLUECAT_USER_NAME"]
	config.Password = values["BLUECAT_PASSWORD"]
	config.ConfigName = values["BLUECAT_CONFIG_NAME"]
	config.DNSView = values["BLUECAT_DNS_VIEW"]

	return bluecat.NewDNSProviderConfig(config)
}
