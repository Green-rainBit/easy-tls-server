package dns

import (
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/acmedns"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/alidns"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/auroradns"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/azure"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/bindman"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/bluecat"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/cloudflare"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/cloudns"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/cloudxns"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/conoha"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/designate"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/digitalocean"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/dnsimple"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/dnsmadeeasy"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/dnspod"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/dode"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/dreamhost"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/duckdns"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/dyn"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/easydns"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/exec"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/exoscale"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/fastdns"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/gandi"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/gandiv5"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/gcloud"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/glesys"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/godaddy"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/hostingde"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/httpreq"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/iij"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/inwx"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/joker"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/lightsail"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/linode"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/linodev4"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/mydnsjp"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/namecheap"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/namedotcom"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/namesilo"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/netcup"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/nifcloud"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/ns1"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns/oraclecloud"
	"fmt"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/otc"
	"github.com/go-acme/lego/v4/providers/dns/ovh"
	"github.com/go-acme/lego/v4/providers/dns/pdns"
	"github.com/go-acme/lego/v4/providers/dns/rackspace"
	"github.com/go-acme/lego/v4/providers/dns/rfc2136"
	"github.com/go-acme/lego/v4/providers/dns/route53"
	"github.com/go-acme/lego/v4/providers/dns/sakuracloud"
	"github.com/go-acme/lego/v4/providers/dns/selectel"
	"github.com/go-acme/lego/v4/providers/dns/stackpath"
	"github.com/go-acme/lego/v4/providers/dns/transip"
	"github.com/go-acme/lego/v4/providers/dns/vegadns"
	"github.com/go-acme/lego/v4/providers/dns/versio"
	"github.com/go-acme/lego/v4/providers/dns/vscale"
	"github.com/go-acme/lego/v4/providers/dns/vultr"
	"github.com/go-acme/lego/v4/providers/dns/zoneee"
)

func NewDNSChallengeProviderByName(name string, values map[string]string) (challenge.Provider, error) {
	switch name {
	case "acme-dns":
		return acmedns.NewDNSProviderByValues(values)
	case "alidns":
		return alidns.NewDNSProviderByValues(values)
	case "azure":
		return azure.NewDNSProviderByValues(values)
	case "auroradns":
		return auroradns.NewDNSProviderByValues(values)
	case "bindman":
		return bindman.NewDNSProviderByValues(values)
	case "bluecat":
		return bluecat.NewDNSProviderByValues(values)
	case "cloudflare":
		return cloudflare.NewDNSProviderByValues(values)
	case "cloudns":
		return cloudns.NewDNSProviderByValues(values)
	case "cloudxns":
		return cloudxns.NewDNSProviderByValues(values)
	case "conoha":
		return conoha.NewDNSProviderByValues(values)
	case "designate":
		return designate.NewDNSProviderByValues(values)
	case "digitalocean":
		return digitalocean.NewDNSProviderByValues(values)
	case "dnsimple":
		return dnsimple.NewDNSProviderByValues(values)
	case "dnsmadeeasy":
		return dnsmadeeasy.NewDNSProviderByValues(values)
	case "dnspod":
		return dnspod.NewDNSProviderByValues(values)
	case "dode":
		return dode.NewDNSProviderByValues(values)
	case "dreamhost":
		return dreamhost.NewDNSProviderByValues(values)
	case "duckdns":
		return duckdns.NewDNSProviderByValues(values)
	case "dyn":
		return dyn.NewDNSProviderByValues(values)
	case "fastdns":
		return fastdns.NewDNSProviderByValues(values)
	case "easydns":
		return easydns.NewDNSProviderByValues(values)
	case "exec":
		return exec.NewDNSProviderByValues(values)
	case "exoscale":
		return exoscale.NewDNSProviderByValues(values)
	case "gandi":
		return gandi.NewDNSProviderByValues(values)
	case "gandiv5":
		return gandiv5.NewDNSProviderByValues(values)
	case "glesys":
		return glesys.NewDNSProviderByValues(values)
	case "gcloud":
		return gcloud.NewDNSProviderByValues(values)
	case "godaddy":
		return godaddy.NewDNSProviderByValues(values)
	case "hostingde":
		return hostingde.NewDNSProviderByValues(values)
	case "httpreq":
		return httpreq.NewDNSProviderByValues(values)
	case "iij":
		return iij.NewDNSProviderByValues(values)
	case "inwx":
		return inwx.NewDNSProviderByValues(values)
	case "joker":
		return joker.NewDNSProviderByValues(values)
	case "lightsail":
		return lightsail.NewDNSProviderByValues(values)
	case "linode":
		return linode.NewDNSProviderByValues(values)
	case "linodev4":
		return linodev4.NewDNSProviderByValues(values)
	case "manual":
		return dns01.NewDNSProviderManual()
	case "mydnsjp":
		return mydnsjp.NewDNSProviderByValues(values)
	case "namecheap":
		return namecheap.NewDNSProviderByValues(values)
	case "namedotcom":
		return namedotcom.NewDNSProviderByValues(values)
	case "namesilo":
		return namesilo.NewDNSProviderByValues(values)
	case "netcup":
		return netcup.NewDNSProviderByValues(values)
	case "nifcloud":
		return nifcloud.NewDNSProviderByValues(values)
	case "ns1":
		return ns1.NewDNSProviderByValues(values)
	case "oraclecloud":
		return oraclecloud.NewDNSProviderByValues(values)
	case "otc":
		return otc.NewDNSProvider()
	case "ovh":
		return ovh.NewDNSProvider()
	case "pdns":
		return pdns.NewDNSProvider()
	case "rackspace":
		return rackspace.NewDNSProvider()
	case "route53":
		return route53.NewDNSProvider()
	case "rfc2136":
		return rfc2136.NewDNSProvider()
	case "sakuracloud":
		return sakuracloud.NewDNSProvider()
	case "stackpath":
		return stackpath.NewDNSProvider()
	case "selectel":
		return selectel.NewDNSProvider()
	case "transip":
		return transip.NewDNSProvider()
	case "vegadns":
		return vegadns.NewDNSProvider()
	case "versio":
		return versio.NewDNSProvider()
	case "vultr":
		return vultr.NewDNSProvider()
	case "vscale":
		return vscale.NewDNSProvider()
	case "zoneee":
		return zoneee.NewDNSProvider()
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}
