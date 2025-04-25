package bluecat

import (
	"fmt"

	"github.com/go-acme/lego/providers/dns/bluecat"
)

func NewDNSProviderByValues(values map[string]string) (*bluecat.DNSProvider, error) {
	if values["BLUECAT_SERVER_URL"] == "" || values["BLUECAT_USER_NAME"] == "" || values["BLUECAT_PASSWORD"] == "" || values["BLUECAT_CONFIG_NAME"] == "" || values["BLUECAT_DNS_VIEW"] == "" {
		return nil, fmt.Errorf("bluecat: BLUECAT_SERVER_URL or BLUECAT_USER_NAME or BLUECAT_PASSWORD or BLUECAT_CONFIG_NAME or BLUECAT_DNS_VIEW is nil")
	}

	config := bluecat.NewDefaultConfig()
	config.BaseURL = values["BLUECAT_SERVER_URL"]
	config.UserName = values["BLUECAT_USER_NAME"]
	config.Password = values["BLUECAT_PASSWORD"]
	config.ConfigName = values["BLUECAT_CONFIG_NAME"]
	config.DNSView = values["BLUECAT_DNS_VIEW"]

	return bluecat.NewDNSProviderConfig(config)
}
