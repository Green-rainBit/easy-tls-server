package exec

import (
	"fmt"

	"github.com/go-acme/lego/v4/providers/dns/exec"
)

func NewDNSProviderByValues(values map[string]string) (*exec.DNSProvider, error) {
	if values[exec.EnvPath] == "" {
		return nil, fmt.Errorf("exec: EXEC_PATH is nil")
	}

	config := exec.NewDefaultConfig()
	config.Program = values[exec.EnvPath]
	config.Mode = values[exec.EnvMode]

	return exec.NewDNSProviderConfig(config)
}
