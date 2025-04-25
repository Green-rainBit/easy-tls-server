package mydnsjp

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/mydnsjp"
)

func NewDNSProviderByValues(values map[string]string) (*mydnsjp.DNSProvider, error) {
	if values[mydnsjp.EnvMasterID] == "" || values[mydnsjp.EnvPassword] == "" {
		return nil, fmt.Errorf("mydnsjp: MYDNSJP_MASTER_ID or MYDNSJP_PASSWORD is nil")
	}
	config := mydnsjp.NewDefaultConfig()
	config.MasterID = values[mydnsjp.EnvMasterID]
	config.Password = values[mydnsjp.EnvPassword]

	return mydnsjp.NewDNSProviderConfig(config)
}
