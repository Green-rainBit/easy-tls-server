package route53

import "github.com/go-acme/lego/v4/providers/dns/route53"

func NewDNSProviderByValues(values map[string]string) (*route53.DNSProvider, error) {
	return route53.NewDNSProviderConfig(route53.NewDefaultConfig())
}
