package linode

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/linode"
)

func NewDNSProviderByValues(values map[string]string) (*linode.DNSProvider, error) {
	if values[linode.EnvToken] == "" {
		return nil, fmt.Errorf("linode: LINODE_TOKEN is nil")
	}

	config := linode.NewDefaultConfig()
	config.Token = values[linode.EnvToken]

	return linode.NewDNSProviderConfig(config)
}
