package conoha

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/conoha"
)

func NewDNSProviderByValues(values map[string]string) (*conoha.DNSProvider, error) {
	if values[conoha.EnvTenantID] == "" || values[conoha.EnvAPIUsername] == "" || values[conoha.EnvAPIPassword] == "" {
		return nil, fmt.Errorf("conoha: CONOHA_TENANT_ID or CONOHA_API_USERNAME or CONOHA_API_PASSWORD is nil")
	}

	config := conoha.NewDefaultConfig()
	config.TenantID = values[conoha.EnvTenantID]
	config.Username = values[conoha.EnvAPIUsername]
	config.Password = values[conoha.EnvAPIPassword]

	return conoha.NewDNSProviderConfig(config)
}
