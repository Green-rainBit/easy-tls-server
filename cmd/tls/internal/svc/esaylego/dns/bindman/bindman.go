package bindman

import (
	"fmt"

	"github.com/go-acme/lego/providers/dns/bindman"
)

func NewDNSProviderByValues(values map[string]string) (*bindman.DNSProvider, error) {
	if _, ok := values["BINDMAN_MANAGER_ADDRESS"]; !ok {
		return nil, fmt.Errorf("bindman: BINDMAN_MANAGER_ADDRESS is nil ")
	}

	config := bindman.NewDefaultConfig()
	config.BaseURL = values["BINDMAN_MANAGER_ADDRESS"]

	return bindman.NewDNSProviderConfig(config)
}
