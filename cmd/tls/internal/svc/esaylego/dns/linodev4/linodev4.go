package linodev4

import (
	"fmt"

	"github.com/go-acme/lego/providers/dns/linodev4"
)

func NewDNSProviderByValues(values map[string]string) (*linodev4.DNSProvider, error) {
	if values["LINODE_TOKEN"] == "" {
		return nil, fmt.Errorf("linodev4: LINODE_TOKEN is nil")
	}

	config := linodev4.NewDefaultConfig()
	config.Token = values["LINODE_TOKEN"]

	return linodev4.NewDNSProviderConfig(config)
}
