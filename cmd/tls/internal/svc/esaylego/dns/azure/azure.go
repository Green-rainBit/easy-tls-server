package azure

import (
	"fmt"

	aazure "github.com/Azure/go-autorest/autorest/azure"
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

	config.SubscriptionID = values[azure.EnvSubscriptionID]
	config.ResourceGroup = values[azure.EnvResourceGroup]
	config.ClientSecret = values[azure.EnvClientSecret]
	config.ClientID = values[azure.EnvClientID]
	config.TenantID = values[azure.EnvTenantID]
	config.PrivateZone = values[azure.EnvPrivateZone] == "true"

	return azure.NewDNSProviderConfig(config)
}
