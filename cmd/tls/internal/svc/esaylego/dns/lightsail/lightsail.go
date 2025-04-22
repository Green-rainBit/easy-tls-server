package lightsail

import (
	"github.com/go-acme/lego/v4/providers/dns/lightsail"
)

func NewDNSProviderByValues(values map[string]string) (*lightsail.DNSProvider, error) {
	config := lightsail.NewDefaultConfig()

	config.DNSZone = values[lightsail.EnvDNSZone]
	config.Region = "us-east-1"
	if values[lightsail.EnvRegion] != "" {
		config.Region = values[lightsail.EnvRegion]
	}

	return lightsail.NewDNSProviderConfig(config)
}
