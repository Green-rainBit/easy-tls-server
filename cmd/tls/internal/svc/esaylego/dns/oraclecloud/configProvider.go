package oraclecloud

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/oracle/oci-go-sdk/v65/common"
)

type configProvider struct {
	values               map[string]string
	privateKeyPassphrase string
}

func newConfigProvider(values map[string]string) *configProvider {
	return &configProvider{
		values:               values,
		privateKeyPassphrase: values[EnvPrivKeyPass],
	}
}

func (p *configProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	privateKey, err := p.getPrivateKey(envPrivKey)
	if err != nil {
		return nil, fmt.Errorf("oraclecloud: %w", err)
	}

	return common.PrivateKeyFromBytesWithPassword(privateKey, []byte(p.privateKeyPassphrase))
}

func (p *configProvider) KeyID() (string, error) {
	tenancy, err := p.TenancyOCID()
	if err != nil {
		return "", err
	}

	user, err := p.UserOCID()
	if err != nil {
		return "", err
	}

	fingerprint, err := p.KeyFingerprint()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", tenancy, user, fingerprint), nil
}

func (p *configProvider) TenancyOCID() (value string, err error) {
	return p.values[EnvTenancyOCID], nil
}

func (p *configProvider) UserOCID() (string, error) {
	return p.values[EnvUserOCID], nil
}

func (p *configProvider) KeyFingerprint() (string, error) {
	return p.values[EnvPubKeyFingerprint], nil
}

func (p *configProvider) Region() (string, error) {
	return p.values[EnvRegion], nil
}

func (p *configProvider) AuthType() (common.AuthConfig, error) {
	// Inspired by https://github.com/oracle/oci-go-sdk/blob/e7635c292e60d0a9dcdd3a1e7de180d7c99b1eee/common/configuration.go#L231-L234
	return common.AuthConfig{AuthType: common.UnknownAuthenticationType}, errors.New("unsupported, keep the interface")
}

func (p *configProvider) getPrivateKey(envVar string) ([]byte, error) {
	envVarValue := p.values[envVar]
	if envVarValue != "" {
		bytes, err := base64.StdEncoding.DecodeString(envVarValue)
		if err != nil {
			return nil, fmt.Errorf("failed to read base64 value %s (defined by env var %s): %w", envVarValue, envVar, err)
		}
		return bytes, nil
	}

	fileVar := envVar + "_FILE"
	fileVarValue := p.values[fileVar]
	if fileVarValue == "" {
		return nil, fmt.Errorf("no value provided for: %s or %s", envVar, fileVar)
	}

	fileContents, ok := p.values[fileVarValue]
	if !ok {
		return nil, fmt.Errorf("failed to read the file %s (defined by env var %s)", fileVarValue, fileVar)
	}

	return []byte(fileContents), nil
}
