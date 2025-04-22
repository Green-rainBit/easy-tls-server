package iij

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/iij"
)

func NewDNSProviderByValues(values map[string]string) (*iij.DNSProvider, error) {
	if values[iij.EnvAPIAccessKey] == "" || values[iij.EnvAPISecretKey] == "" || values[iij.EnvDoServiceCode] == "" {
		return nil, fmt.Errorf("iij: IIJ_API_ACCESS_KEY or IIJ_API_SECRET_KEY or IIJ_DO_SERVICE_CODE is nil")
	}

	config := iij.NewDefaultConfig()
	config.AccessKey = values[iij.EnvAPIAccessKey]
	config.SecretKey = values[iij.EnvAPISecretKey]
	config.DoServiceCode = values[iij.EnvDoServiceCode]

	return iij.NewDNSProviderConfig(config)
}
