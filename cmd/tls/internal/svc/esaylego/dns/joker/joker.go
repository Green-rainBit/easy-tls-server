package joker

import (
	"fmt"

	"github.com/go-acme/lego/challenge"
	"github.com/go-acme/lego/v4/providers/dns/joker"
)

func NewDNSProviderByValues(values map[string]string) (challenge.ProviderTimeout, error) {

	if values[joker.EnvAPIKey] == "" {
		if values[joker.EnvUsername] == "" || values[joker.EnvPassword] == "" {
			//nolint:errorlint // false-positive
			return nil, fmt.Errorf("joker: JOKER_USERNAME or JOKER_PASSWORD is nil")
		}
	}

	config := joker.NewDefaultConfig()
	config.APIKey = values[joker.EnvAPIKey]
	config.Username = values[joker.EnvUsername]
	config.Password = values[joker.EnvPassword]

	return joker.NewDNSProviderConfig(config)
}
