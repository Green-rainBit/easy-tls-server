package dyn

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/dyn"
)

func NewDNSProviderByValues(values map[string]string) (*dyn.DNSProvider, error) {
	if values[dyn.EnvCustomerName] == "" || values[dyn.EnvUserName] == "" || values[dyn.EnvPassword] == "" {
		return nil, fmt.Errorf("dyn: DYN_CUSTOMER_NAME and DYN_USER_NAME and DYN_PASSWORD is nil")
	}

	config := dyn.NewDefaultConfig()
	config.CustomerName = values[dyn.EnvCustomerName]
	config.UserName = values[dyn.EnvUserName]
	config.Password = values[dyn.EnvPassword]

	return dyn.NewDNSProviderConfig(config)
}
