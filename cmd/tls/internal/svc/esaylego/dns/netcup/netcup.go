package netcup

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/netcup"
)

func NewDNSProviderByValues(values map[string]string) (*netcup.DNSProvider, error) {
	if values[netcup.EnvCustomerNumber] == "" || values[netcup.EnvAPIKey] == "" || values[netcup.EnvAPIPassword] == "" {
		return nil, fmt.Errorf("netcup: EnvAPIPassword or NETCUP_API_KEY or NETCUP_API_PASSWORD")
	}

	config := netcup.NewDefaultConfig()
	config.Customer = values[netcup.EnvCustomerNumber]
	config.Key = values[netcup.EnvAPIKey]
	config.Password = values[netcup.EnvAPIPassword]

	return netcup.NewDNSProviderConfig(config)
}
