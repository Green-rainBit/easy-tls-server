package gcloud

import (
	"context"

	"cloud.google.com/go/compute/metadata"
	"github.com/go-acme/lego/v4/providers/dns/gcloud"
)

func NewDNSProviderByValues(values map[string]string) (*gcloud.DNSProvider, error) {
	// Use a service account file if specified via environment variable.
	if saKey := values[gcloud.EnvServiceAccount]; saKey != "" {
		return gcloud.NewDNSProviderServiceAccountKey([]byte(saKey))
	}

	// Use default credentials.
	project := autodetectProjectID(context.Background())

	if values[gcloud.EnvProject] == "" {
		project = values[gcloud.EnvProject]
	}
	return gcloud.NewDNSProviderCredentials(project)
}

func autodetectProjectID(ctx context.Context) string {
	if pid, err := metadata.ProjectIDWithContext(ctx); err == nil {
		return pid
	}

	return ""
}
