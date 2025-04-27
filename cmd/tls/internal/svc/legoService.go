package svc

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"

	"esay-tls-server/cmd/tls/internal/svc/esaylego"
	"esay-tls-server/cmd/tls/internal/svc/esaylego/dns"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/resolver"

	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

type myUser struct {
	Email        string
	Registration *registration.Resource
	Key          crypto.PrivateKey
}

func (u *myUser) GetEmail() string {
	return u.Email
}
func (u myUser) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *myUser) GetPrivateKey() crypto.PrivateKey {
	return u.Key
}

func NewLegoSerivce(email, privateKeyStrey, env string) *LegoSerivce {
	bytes, err := base64.StdEncoding.DecodeString(privateKeyStrey)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := x509.ParseECPrivateKey(bytes)
	if err != nil {
		log.Fatal(err)
	}
	legoMyUser := myUser{
		Email:        email,
		Key:          privateKey,
		Registration: &registration.Resource{},
	}

	config := NewLegoConfigByUser(&legoMyUser)
	if env != "production" {
		config.CADirURL = lego.LEDirectoryStaging
	} else {
		config.CADirURL = lego.LEDirectoryProduction
	}
	tlsClient, err := esaylego.NewEsayClient(config)
	if err != nil {
		log.Fatal(err)
	}
	reg, err := tlsClient.Registration.ResolveAccountByKey()
	if err != nil {
		log.Fatal(err)
	}
	legoMyUser.Registration = reg
	return &LegoSerivce{
		client: tlsClient,
	}
}

func NewLegoConfigByUser(legoMyUser *myUser) *lego.Config {
	config := lego.NewConfig(legoMyUser)
	// This CA URL is configured for a local dev instance of Boulder running in Docker in a VM.
	// config.CADirURL = lego.LEDirectoryProduction
	config.CADirURL = lego.LEDirectoryStaging
	config.Certificate.KeyType = certcrypto.EC256

	return config
}

type LegoSerivce struct {
	client *esaylego.EsayClient
}

func (t *LegoSerivce) GetTlsByDomain(domains []string, name string, values map[string]string) (certificate.Resource, error) {
	provider, err := dns.NewDNSChallengeProviderByName(name, values)
	if err != nil {
		return certificate.Resource{}, err
	}
	solversManager := resolver.NewSolversManager(t.client.GetCore())
	solversManager.SetDNS01Provider(provider)
	prober := resolver.NewProber(solversManager)
	certifier := certificate.NewCertifier(t.client.GetCore(), prober, t.client.GetCertifierOptionse())
	if len(domains) == 0 {
		return certificate.Resource{}, fmt.Errorf("domains not is empty")
	}
	request := certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	}
	certificates, err := certifier.Obtain(request)
	if err != nil {
		return certificate.Resource{}, err
	}
	if certificates != nil {
		return *certificates, nil
	}
	return certificate.Resource{}, fmt.Errorf("certificates not is empty")
}
