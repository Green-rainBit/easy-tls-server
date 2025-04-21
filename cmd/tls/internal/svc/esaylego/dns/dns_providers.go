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
	"fmt"

	"github.com/go-acme/lego/providers/dns/linodev4"
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/gcloud"
	"github.com/go-acme/lego/v4/providers/dns/glesys"
	"github.com/go-acme/lego/v4/providers/dns/godaddy"
	"github.com/go-acme/lego/v4/providers/dns/hostingde"
	"github.com/go-acme/lego/v4/providers/dns/httpreq"
	"github.com/go-acme/lego/v4/providers/dns/iij"
	"github.com/go-acme/lego/v4/providers/dns/inwx"
	"github.com/go-acme/lego/v4/providers/dns/joker"
	"github.com/go-acme/lego/v4/providers/dns/lightsail"
	"github.com/go-acme/lego/v4/providers/dns/linode"
	"github.com/go-acme/lego/v4/providers/dns/mydnsjp"
	"github.com/go-acme/lego/v4/providers/dns/namecheap"
	"github.com/go-acme/lego/v4/providers/dns/namedotcom"
	"github.com/go-acme/lego/v4/providers/dns/namesilo"
	"github.com/go-acme/lego/v4/providers/dns/netcup"
	"github.com/go-acme/lego/v4/providers/dns/nifcloud"
	"github.com/go-acme/lego/v4/providers/dns/ns1"
	"github.com/go-acme/lego/v4/providers/dns/oraclecloud"
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
		return glesys.NewDNSProvider()
	case "gcloud":
		return gcloud.NewDNSProvider()
	case "godaddy":
		return godaddy.NewDNSProvider()
	case "hostingde":
		return hostingde.NewDNSProvider()
	case "httpreq":
		return httpreq.NewDNSProvider()
	case "iij":
		return iij.NewDNSProvider()
	case "inwx":
		return inwx.NewDNSProvider()
	case "joker":
		return joker.NewDNSProvider()
	case "lightsail":
		return lightsail.NewDNSProvider()
	case "linode":
		return linode.NewDNSProvider()
	case "linodev4":
		return linodev4.NewDNSProvider()
	case "manual":
		return dns01.NewDNSProviderManual()
	case "mydnsjp":
		return mydnsjp.NewDNSProvider()
	case "namecheap":
		return namecheap.NewDNSProvider()
	case "namedotcom":
		return namedotcom.NewDNSProvider()
	case "namesilo":
		return namesilo.NewDNSProvider()
	case "netcup":
		return netcup.NewDNSProvider()
	case "nifcloud":
		return nifcloud.NewDNSProvider()
	case "ns1":
		return ns1.NewDNSProvider()
	case "oraclecloud":
		return oraclecloud.NewDNSProvider()
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
