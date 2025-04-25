package acmedns

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v4/providers/dns/acmedns"
)

func NewDNSProviderByValues(values map[string]string) (*acmedns.DNSProvider, error) {
	if values[acmedns.EnvAPIBase] == "" {
		return nil, fmt.Errorf("acme-dns EnvAPIBase is nil")
	}

	config := acmedns.NewDefaultConfig()
	config.APIBase = values[acmedns.EnvAPIBase]
	config.StoragePath = values[acmedns.EnvStoragePath]
	config.StorageBaseURL = values[acmedns.EnvStorageBaseURL]

	allowList := values[acmedns.EnvAllowList]
	if allowList != "" {
		config.AllowList = strings.Split(allowList, ",")
	}

	return acmedns.NewDNSProviderConfig(config)
}
