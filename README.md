**Read this in other languages: [English](README.md), [中文](README_zh.md).**

# Introduction
    easy-tls-server is a TLS certificate generation service based on gozero and lego. With simple configuration and no middleware required, it can stably provide Let's Encrypt certificates for multiple domain services online.
    You can achieve non-stop TLS certificate replacement for multiple services using easy-tls-server.

# Configuration Description

```yml
Name: tls-api
Host: 0.0.0.0
Port: 8888
Timeout: 100000

LegoConf:
  Email:    		    // Let's Encrypt email
  PrivateKeyStre:		// Let's Encrypt private key
  Env:				// production or develop (default is develop)
  DomainMap:		// Domain configuration
    www.demo.top:		// Domain
      DNS_CHALLENGE:		// DNS provider
      CLOUDFLARE_EMAIL:		// Provider credentials
      CLOUDFLARE_DNS_API_TOKEN:
    www.demo.cn: 
      DNS_CHALLENGE: 
      ALICLOUD_ACCESS_KEY: 
      ALICLOUD_SECRET_KEY: 
      ALICLOUD_REGION_ID: 
```

# Directory Structure Description
    ├── README.md           	// Documentation file
    └── cmd
        └── tls     			 //   tls service includes various versions of include, src folders
            ├── tls.go 			 //  main.go
            ├── ect 			 //  Configuration file storage
            |   └── tls-api.yaml  //  Configuration file
            ├── config 			 //  Static configuration file corresponding to the structure declaration directory
            └── internal          //  Single service internal files, visible only within the current service
                └── handler		  //  Service route management
                ├── logic		  //  Specific logic
                ├── svc		      //  Dependency injection
               	|	└── esaylego		  //  Lego encapsulation
               	|		└── dns		  		//  DNS client encapsulation
                └── types
         		   └── types.go   // Structure storage directory

# Usage Instructions

You need to pre-obtain the Let's Encrypt email and private key, fill them into the tls-api.yaml file, and then configure the domain information corresponding to the certificate.

```sh
cd cmd
go mod tidy
go run tls.go

GET http://127.0.0.1:8888/get/tls?domains=www.demo.top

* Response
{
    "domain": "",
    "certUrl": "",
    "certStableUrl": "",
    "privateKey": "",
    "certificate": "",
    "issuerCertificate": "",
    "cSR": ""
}
```

## Supported DNS Providers

Detailed documentation is available [here](https://go-acme.github.io/lego/dns).

<!-- START DNS PROVIDERS LIST -->
<table><tr>
  <td><a href="https://go-acme.github.io/lego/dns/active24/">Active24</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/edgedns/">Akamai EdgeDNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/alidns/">Alibaba Cloud DNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/allinkl/">all-inkl</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/lightsail/">Amazon Lightsail</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/route53/">Amazon Route 53</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/arvancloud/">ArvanCloud</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/auroradns/">Aurora DNS</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/autodns/">Autodns</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/axelname/">Axelname</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/azure/">Azure (deprecated)</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/azuredns/">Azure DNS</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/baiducloud/">Baidu Cloud</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/bindman/">Bindman</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/bluecat/">Bluecat</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/bookmyname/">BookMyName</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/brandit/">Brandit (deprecated)</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/bunny/">Bunny</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/checkdomain/">Checkdomain</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/civo/">Civo</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/cloudru/">Cloud.ru</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/clouddns/">CloudDNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/cloudflare/">Cloudflare</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/cloudns/">ClouDNS</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/cloudxns/">CloudXNS (Deprecated)</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/conoha/">ConoHa</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/constellix/">Constellix</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/corenetworks/">Core-Networks</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/cpanel/">CPanel/WHM</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/derak/">Derak Cloud</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/desec/">deSEC.io</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/designate/">Designate DNSaaS for Openstack</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/digitalocean/">Digital Ocean</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/directadmin/">DirectAdmin</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/dnsmadeeasy/">DNS Made Easy</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/dnshomede/">dnsHome.de</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/dnsimple/">DNSimple</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/dnspod/">DNSPod (deprecated)</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/dode/">Domain Offensive (do.de)</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/domeneshop/">Domeneshop</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/dreamhost/">DreamHost</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/duckdns/">Duck DNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/dyn/">Dyn</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/dynu/">Dynu</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/easydns/">EasyDNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/efficientip/">Efficient IP</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/epik/">Epik</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/exoscale/">Exoscale</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/exec/">External program</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/f5xc/">F5 XC</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/freemyip/">freemyip.com</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/gcore/">G-Core</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/gandi/">Gandi</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/gandiv5/">Gandi Live DNS (v5)</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/glesys/">Glesys</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/godaddy/">Go Daddy</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/gcloud/">Google Cloud</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/googledomains/">Google Domains</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/hetzner/">Hetzner</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/hostingde/">Hosting.de</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/hosttech/">Hosttech</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/httpreq/">HTTP request</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/httpnet/">http.net</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/huaweicloud/">Huawei Cloud</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/hurricane/">Hurricane Electric DNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/hyperone/">HyperOne</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/ibmcloud/">IBM Cloud (SoftLayer)</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/iijdpf/">IIJ DNS Platform Service</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/infoblox/">Infoblox</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/infomaniak/">Infomaniak</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/iij/">Internet Initiative Japan</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/internetbs/">Internet.bs</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/inwx/">INWX</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/ionos/">Ionos</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/ipv64/">IPv64</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/iwantmyname/">iwantmyname</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/joker/">Joker</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/acme-dns/">Joohoi&#39;s ACME-DNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/liara/">Liara</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/limacity/">Lima-City</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/linode/">Linode (v4)</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/liquidweb/">Liquid Web</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/loopia/">Loopia</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/luadns/">LuaDNS</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/mailinabox/">Mail-in-a-Box</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/manageengine/">ManageEngine CloudDNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/manual/">Manual</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/metaname/">Metaname</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/metaregistrar/">Metaregistrar</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/mijnhost/">mijn.host</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/mittwald/">Mittwald</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/myaddr/">myaddr.{tools,dev,io}</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/mydnsjp/">MyDNS.jp</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/mythicbeasts/">MythicBeasts</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/namedotcom/">Name.com</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/namecheap/">Namecheap</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/namesilo/">Namesilo</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/nearlyfreespeech/">NearlyFreeSpeech.NET</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/netcup/">Netcup</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/netlify/">Netlify</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/nicmanager/">Nicmanager</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/nifcloud/">NIFCloud</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/njalla/">Njalla</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/nodion/">Nodion</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/ns1/">NS1</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/otc/">Open Telekom Cloud</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/oraclecloud/">Oracle Cloud</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/ovh/">OVH</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/plesk/">plesk.com</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/porkbun/">Porkbun</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/pdns/">PowerDNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/rackspace/">Rackspace</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/rainyun/">Rain Yun/雨云</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/rcodezero/">RcodeZero</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/regru/">reg.ru</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/regfish/">Regfish</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/rfc2136/">RFC2136</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/rimuhosting/">RimuHosting</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/sakuracloud/">Sakura Cloud</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/scaleway/">Scaleway</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/selectel/">Selectel</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/selectelv2/">Selectel v2</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/selfhostde/">SelfHost.(de|eu)</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/servercow/">Servercow</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/shellrent/">Shellrent</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/simply/">Simply.com</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/sonic/">Sonic</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/spaceship/">Spaceship</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/stackpath/">Stackpath</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/technitium/">Technitium</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/tencentcloud/">Tencent Cloud DNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/timewebcloud/">Timeweb Cloud</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/transip/">TransIP</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/safedns/">UKFast SafeDNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/ultradns/">Ultradns</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/variomedia/">Variomedia</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/vegadns/">VegaDNS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/vercel/">Vercel</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/versio/">Versio.[nl|eu|uk]</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/vinyldns/">VinylDNS</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/vkcloud/">VK Cloud</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/volcengine/">Volcano Engine/火山引擎</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/vscale/">Vscale</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/vultr/">Vultr</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/webnames/">Webnames</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/websupport/">Websupport</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/wedos/">WEDOS</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/westcn/">West.cn/西部数码</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/yandex360/">Yandex 360</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/yandexcloud/">Yandex Cloud</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/yandex/">Yandex PDD</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/zoneee/">Zone.ee</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/zonomi/">Zonomi</a></td>
  <td></td>
  <td></td>
  <td></td>
</tr></table>

If your DNS provider is not supported, please open an [issue](https://github.com/go-acme/lego/issues/new?assignees=&labels=enhancement%2C+new-provider&template=new_dns_provider.md).