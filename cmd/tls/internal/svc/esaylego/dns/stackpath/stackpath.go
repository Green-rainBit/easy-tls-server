package stackpath

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/stackpath"
)

func NewDNSProviderByValues(values map[string]string) (*stackpath.DNSProvider, error) {
	if values[stackpath.EnvClientID] == "" || values[stackpath.EnvClientSecret] == "" || values[stackpath.EnvStackID] == "" {
		return nil, fmt.Errorf("stackpath: STACKPATH_CLIENT_ID or STACKPATH_CLIENT_SECRET or STACKPATH_STACK_ID")
	}

	config := stackpath.NewDefaultConfig()
	config.ClientID = values[stackpath.EnvClientID]
	config.ClientSecret = values[stackpath.EnvClientSecret]
	config.StackID = values[stackpath.EnvStackID]

	return stackpath.NewDNSProviderConfig(config)
}
