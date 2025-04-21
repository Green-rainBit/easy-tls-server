package cloudxns

import (
	"fmt"

	"github.com/go-acme/lego/providers/dns/cloudxns"
	v4 "github.com/go-acme/lego/v4/providers/dns/cloudxns"
)

func NewDNSProviderByValues(values map[string]string) (*cloudxns.DNSProvider, error) {
	if _, ok := values[v4.EnvAPIKey]; !ok {
		return nil, fmt.Errorf("CloudXNS: CLOUDXNS_API_KEY is nil ")
	}

	config := cloudxns.NewDefaultConfig()
	config.APIKey = values[v4.EnvAPIKey]
	config.SecretKey = values[v4.EnvSecretKey]

	return cloudxns.NewDNSProviderConfig(config)
}
