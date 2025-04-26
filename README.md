# 介绍
    easy-tls-server是基于gozero与lego的tls生成证书服务，只需要简单的配置、无需中间件，就可以稳定为在线服务提供Let's Encryptls证书。
    一个服务管理多个域名证书。

# 配置说明

```yml
Name: tls-api
Host: 0.0.0.0
Port: 8888
Timeout: 100000

LegoConf:
  Email:    		    // Let's Encrypt邮箱
  PrivateKeyStre:		// Let's Encrypt密钥
  Env:				// production or develop 默认为develop
  DomainMap:		// 域名配置
    www.demo.top:		// 域名
      DNS_CHALLENGE:		// dns 服务商
      CLOUDFLARE_EMAIL:		// 运营商密钥
      CLOUDFLARE_DNS_API_TOKEN:
    www.demo.cn: 
      DNS_CHALLENGE: 
      ALICLOUD_ACCESS_KEY: 
      ALICLOUD_SECRET_KEY: 
      ALICLOUD_REGION_ID: 
```




# 目录结构描述
    ├── README.md           	// 介绍文档
    └── cmd
        └── tls     			 //   tls服务包含各版本的include、src文件夹
            ├── tls.go 			 //  main.go
            ├── ect 			 //  存放配置文件
            |   └── tls-api.yaml  //  配置文件
            ├── config 			 //  静态配置文件对应的结构体声明目录
            └── internal          //  单个服务内部文件，其可见范围仅限当前服务
                └── handler		  //  服务路由管理
                ├── logic		  //  具体逻辑
                ├── svc		      //  依赖注入
               	|	└── esaylego		  //  lego封装
               	├		└── dns		  		//  dns客户端封装
                └── types
         		   └── types.go   //结构体存放目录

# 使用说明

需要预获取Let's Encrypt邮箱与密钥，填写到tls-api.yaml文件内，再填写证书对应的域名的配置信息。

```sh
cd cmd
go mod tidy
go run tls.go
```


## 支持的DNS运营商

Detailed documentation is available [here](https://go-acme.github.io/lego/dns).

<!-- START DNS PROVIDERS LIST -->

<table><tr>
  <td>Active24</td>
  <td>Akamai EdgeDNS</td>
  <td>Alibaba Cloud DNS</td>
  <td>all-inkl</td>
</tr><tr>
  <td>Amazon Lightsail</td>
  <td>Amazon Route 53</td>
  <td>ArvanCloud</td>
  <td>Aurora DNS</td>
</tr><tr>
  <td>Autodns</td>
  <td>Axelname</td>
  <td>Azure (deprecated)</td>
  <td>Azure DNS</td>
</tr><tr>
  <td>Baidu Cloud</td>
  <td>Bindman</td>
  <td>Bluecat</td>
  <td>BookMyName</td>
</tr><tr>
  <td>Brandit (deprecated)</td>
  <td>Bunny</td>
  <td>Checkdomain</td>
  <td>Civo</td>
</tr><tr>
  <td>Cloud.ru</td>
  <td>CloudDNS</td>
  <td>Cloudflare</td>
  <td>ClouDNS</td>
</tr><tr>
  <td>CloudXNS (Deprecated)</td>
  <td>ConoHa</td>
  <td>Constellix</td>
  <td>Core-Networks</td>
</tr><tr>
  <td>CPanel/WHM</td>
  <td>Derak Cloud</td>
  <td>deSEC.io</td>
  <td>Designate DNSaaS for Openstack</td>
</tr><tr>
  <td>Digital Ocean</td>
  <td>DirectAdmin</td>
  <td>DNS Made Easy</td>
  <td>dnsHome.de</td>
</tr><tr>
  <td>DNSimple</td>
  <td>DNSPod (deprecated)</td>
  <td>Domain Offensive (do.de)</td>
  <td>Domeneshop</td>
</tr><tr>
  <td>DreamHost</td>
  <td>Duck DNS</td>
  <td>Dyn</td>
  <td>Dynu</td>
</tr><tr>
  <td>EasyDNS</td>
  <td>Efficient IP</td>
  <td>Epik</td>
  <td>Exoscale</td>
</tr><tr>
  <td>External program</td>
  <td>F5 XC</td>
  <td>freemyip.com</td>
  <td>G-Core</td>
</tr><tr>
  <td>Gandi</td>
  <td>Gandi Live DNS (v5)</td>
  <td>Glesys</td>
  <td>Go Daddy</td>
</tr><tr>
  <td>Google Cloud</td>
  <td>Google Domains</td>
  <td>Hetzner</td>
  <td>Hosting.de</td>
</tr><tr>
  <td>Hosttech</td>
  <td>HTTP request</td>
  <td>http.net</td>
  <td>Huawei Cloud</td>
</tr><tr>
  <td>Hurricane Electric DNS</td>
  <td>HyperOne</td>
  <td>IBM Cloud (SoftLayer)</td>
  <td>IIJ DNS Platform Service</td>
</tr><tr>
  <td>Infoblox</td>
  <td>Infomaniak</td>
  <td>Internet Initiative Japan</td>
  <td>Internet.bs</td>
</tr><tr>
  <td>INWX</td>
  <td>Ionos</td>
  <td>IPv64</td>
  <td>iwantmyname</td>
</tr><tr>
  <td>Joker</td>
  <td>Joohoi&#39;s ACME-DNS</td>
  <td>Liara</td>
  <td>Lima-City</td>
</tr><tr>
  <td>Linode (v4)</td>
  <td>Liquid Web</td>
  <td>Loopia</td>
  <td>LuaDNS</td>
</tr><tr>
  <td>Mail-in-a-Box</td>
  <td>ManageEngine CloudDNS</td>
  <td>Manual</td>
  <td>Metaname</td>
</tr><tr>
  <td>Metaregistrar</td>
  <td>mijn.host</td>
  <td>Mittwald</td>
  <td>myaddr.{tools,dev,io}</td>
</tr><tr>
  <td>MyDNS.jp</td>
  <td>MythicBeasts</td>
  <td>Name.com</td>
  <td>Namecheap</td>
</tr><tr>
  <td>Namesilo</td>
  <td>NearlyFreeSpeech.NET</td>
  <td>Netcup</td>
  <td>Netlify</td>
</tr><tr>
  <td>Nicmanager</td>
  <td>NIFCloud</td>
  <td>Njalla</td>
  <td>Nodion</td>
</tr><tr>
  <td>NS1</td>
  <td>Open Telekom Cloud</td>
  <td>Oracle Cloud</td>
  <td>OVH</td>
</tr><tr>
  <td>plesk.com</td>
  <td>Porkbun</td>
  <td>PowerDNS</td>
  <td>Rackspace</td>
</tr><tr>
  <td>Rain Yun/雨云</td>
  <td>RcodeZero</td>
  <td>reg.ru</td>
  <td>Regfish</td>
</tr><tr>
  <td>RFC2136</td>
  <td>RimuHosting</td>
  <td>Sakura Cloud</td>
  <td>Scaleway</td>
</tr><tr>
  <td>Selectel</td>
  <td>Selectel v2</td>
  <td>SelfHost.(de|eu)</td>
  <td>Servercow</td>
</tr><tr>
  <td>Shellrent</td>
  <td>Simply.com</td>
  <td>Sonic</td>
  <td>Spaceship</td>
</tr><tr>
  <td>Stackpath</td>
  <td>Technitium</td>
  <td>Tencent Cloud DNS</td>
  <td>Timeweb Cloud</td>
</tr><tr>
  <td>TransIP</td>
  <td>UKFast SafeDNS</td>
  <td>Ultradns</td>
  <td>Variomedia</td>
</tr><tr>
  <td>VegaDNS</td>
  <td>Vercel</td>
  <td>Versio.[nl|eu|uk]</td>
  <td>VinylDNS</td>
</tr><tr>
  <td>VK Cloud</td>
  <td>Volcano Engine/火山引擎</td>
  <td>Vscale</td>
  <td>Vultr</td>
</tr><tr>
  <td>Webnames</td>
  <td>Websupport</td>
  <td>WEDOS</td>
  <td>West.cn/西部数码</td>
</tr><tr>
  <td>Yandex 360</td>
  <td>Yandex Cloud</td>
  <td>Yandex PDD</td>
  <td>Zone.ee</td>
</tr><tr>
  <td>Zonomi</td>
  <td></td>
  <td></td>
  <td></td>
</tr></table>


