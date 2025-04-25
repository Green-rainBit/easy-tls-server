package transip

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/transip"
)

func NewDNSProviderByValues(values map[string]string) (*transip.DNSProvider, error) {
	if values[transip.EnvAccountName] == "" || values[transip.EnvPrivateKeyPath] == "" {
		return nil, fmt.Errorf("transip: TRANSIP_ACCOUNT_NAME or TRANSIP_PRIVATE_KEY_PATH or STACKPATH_STACK_ID")
	}

	config := transip.NewDefaultConfig()
	config.AccountName = values[transip.EnvAccountName]
	config.PrivateKeyPath = values[transip.EnvPrivateKeyPath]

	return transip.NewDNSProviderConfig(config)
}
