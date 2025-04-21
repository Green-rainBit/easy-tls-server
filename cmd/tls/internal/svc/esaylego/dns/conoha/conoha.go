package conoha

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/conoha"
)

func NewDNSProviderByValues(values map[string]string) (*conoha.DNSProvider, error) {
	if _, ok := values[conoha.EnvTenantID]; ok {
		return nil, fmt.Errorf("conoha: config is nil")
	}

	config := conoha.NewDefaultConfig()
	config.TenantID = values[conoha.EnvTenantID]
	config.Username = values[conoha.EnvAPIUsername]
	config.Password = values[conoha.EnvAPIPassword]

	return conoha.NewDNSProviderConfig(config)
}
