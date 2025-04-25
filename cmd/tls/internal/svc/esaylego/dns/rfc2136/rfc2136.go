package rfc2136

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/rfc2136"
)

func NewDNSProviderByValues(values map[string]string) (*rfc2136.DNSProvider, error) {
	if values[rfc2136.EnvNameserver] == "" {
		return nil, fmt.Errorf("rfc2136: RFC2136_NAMESERVER is nil")
	}

	config := rfc2136.NewDefaultConfig()
	config.Nameserver = values[rfc2136.EnvNameserver]

	config.TSIGFile = values[rfc2136.EnvTSIGFile]

	config.TSIGKey = values[rfc2136.EnvTSIGKey]
	config.TSIGSecret = values[rfc2136.EnvTSIGSecret]

	return rfc2136.NewDNSProviderConfig(config)
}
