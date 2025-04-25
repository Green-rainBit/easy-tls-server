package cloudxns

import (
	"fmt"

	"github.com/go-acme/lego/providers/dns/cloudxns"
	v4 "github.com/go-acme/lego/v4/providers/dns/cloudxns"
)

func NewDNSProviderByValues(values map[string]string) (*cloudxns.DNSProvider, error) {
	if values[v4.EnvAPIKey] == "" || values[v4.EnvSecretKey] == "" {
		return nil, fmt.Errorf("CloudXNS: CLOUDXNS_API_KEY or CLOUDXNS_SECRET_KEY is nil ")
	}

	config := cloudxns.NewDefaultConfig()
	config.APIKey = values[v4.EnvAPIKey]
	config.SecretKey = values[v4.EnvSecretKey]

	return cloudxns.NewDNSProviderConfig(config)
}
