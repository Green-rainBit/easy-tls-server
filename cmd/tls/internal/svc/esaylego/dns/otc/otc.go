package otc

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/otc"
)

func NewDNSProviderByValues(values map[string]string) (*otc.DNSProvider, error) {
	if values[otc.EnvDomainName] == "" || values[otc.EnvUserName] == "" || values[otc.EnvPassword] == "" || values[otc.EnvProjectName] == "" {
		return nil, fmt.Errorf("otc: OTC_DOMAIN_NAME or OTC_USER_NAME or OTC_PASSWORD or OTC_PROJECT_NAME is nil")
	}

	config := otc.NewDefaultConfig()
	config.DomainName = values[otc.EnvDomainName]
	config.UserName = values[otc.EnvUserName]
	config.Password = values[otc.EnvPassword]
	config.ProjectName = values[otc.EnvProjectName]

	return otc.NewDNSProviderConfig(config)
}
