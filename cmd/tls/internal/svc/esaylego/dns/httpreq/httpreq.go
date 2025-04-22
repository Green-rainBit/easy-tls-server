package httpreq

import (
	"fmt"
	"net/url"

	"github.com/go-acme/lego/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/httpreq"
)

func NewDNSProviderByValues(values map[string]string) (*httpreq.DNSProvider, error) {
	if values[httpreq.EnvEndpoint] == "" {
		return nil, fmt.Errorf("httpreq: HTTPREQ_ENDPOINT is nil")
	}

	endpoint, err := url.Parse(values[httpreq.EnvEndpoint])
	if err != nil {
		return nil, fmt.Errorf("httpreq: %w", err)
	}

	config := httpreq.NewDefaultConfig()
	config.Mode = env.GetOrFile(httpreq.EnvMode)
	config.Username = env.GetOrFile(httpreq.EnvUsername)
	config.Password = env.GetOrFile(httpreq.EnvPassword)
	config.Endpoint = endpoint
	return httpreq.NewDNSProviderConfig(config)
}
