package cloudns

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/cloudns"
)

func NewDNSProviderByValues(values map[string]string) (*cloudns.DNSProvider, error) {
	var subAuthID string
	authID := values[cloudns.EnvAuthID]
	if authID == "" {
		subAuthID = values[cloudns.EnvSubAuthID]
	}

	if authID == "" && subAuthID == "" {
		return nil, fmt.Errorf("ClouDNS: some credentials information are missing: %s or %s", cloudns.EnvAuthID, cloudns.EnvSubAuthID)
	}

	if _, ok := values[cloudns.EnvAuthPassword]; !ok {
		return nil, fmt.Errorf("ClouDNS: CLOUDNS_AUTH_PASSWORD is nil")
	}

	config := cloudns.NewDefaultConfig()
	config.AuthID = authID
	config.SubAuthID = subAuthID
	config.AuthPassword = values[cloudns.EnvAuthPassword]

	return cloudns.NewDNSProviderConfig(config)
}
