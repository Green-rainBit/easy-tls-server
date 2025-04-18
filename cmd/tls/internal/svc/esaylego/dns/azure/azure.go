package azure

import (
	"fmt"

	aazure "github.com/Azure/go-autorest/autorest/azure"
	"github.com/go-acme/lego/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/azure"
)

func NewDNSProviderByValues(values map[string]string) (*azure.DNSProvider, error) {

	config := azure.NewDefaultConfig()

	environmentName := values[azure.EnvEnvironment]
	if environmentName != "" {
		var environment aazure.Environment
		switch environmentName {
		case "china":
			environment = aazure.ChinaCloud
		case "german":
			environment = aazure.GermanCloud
		case "public":
			environment = aazure.PublicCloud
		case "usgovernment":
			environment = aazure.USGovernmentCloud
		default:
			return nil, fmt.Errorf("azure: unknown environment %s", environmentName)
		}

		config.ResourceManagerEndpoint = environment.ResourceManagerEndpoint
		config.ActiveDirectoryEndpoint = environment.ActiveDirectoryEndpoint
	}

	config.SubscriptionID = env.GetOrFile(azure.EnvSubscriptionID)
	config.ResourceGroup = env.GetOrFile(azure.EnvResourceGroup)
	config.ClientSecret = env.GetOrFile(azure.EnvClientSecret)
	config.ClientID = env.GetOrFile(azure.EnvClientID)
	config.TenantID = env.GetOrFile(azure.EnvTenantID)
	config.PrivateZone = env.GetOrDefaultBool(azure.EnvPrivateZone, false)

	return azure.NewDNSProviderConfig(config)
}
