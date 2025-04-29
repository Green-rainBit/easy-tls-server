package esaylego

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"log"
	"testing"

	"github.com/go-acme/lego/v4/certcrypto"
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

func TestRegister(t *testing.T) {

	ecPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
		return
	}
	privateKey, err := x509.MarshalECPrivateKey(ecPrivateKey)
	if err != nil {
		log.Fatal(err)
		return
	}
	t.Log("priKey:", base64.StdEncoding.EncodeToString(privateKey))
	myUser := myUser{
		Email:        "",
		Key:          ecPrivateKey,
		Registration: &registration.Resource{},
	}
	config := lego.NewConfig(&myUser)
	// This CA URL is configured for a local dev instance of Boulder running in Docker in a VM.
	config.CADirURL = lego.LEDirectoryProduction
	config.Certificate.KeyType = certcrypto.EC256
	client, err := lego.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	regOptions := registration.RegisterOptions{TermsOfServiceAgreed: true}
	reg, err := client.Registration.Register(regOptions)
	if err != nil {
		log.Fatal(err)
	}
	myUser.Registration = reg
}
