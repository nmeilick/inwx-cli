package utils

import (
	"context"
	"fmt"
	"strings"

	"github.com/nmeilick/inwx-cli/pkg/inwx"
)

// InferDomainAndName attempts to infer the domain and record name from a qualified hostname
func InferDomainAndName(client *inwx.Client, ctx context.Context, hostname string) (domain, name string, err error) {
	if hostname == "" {
		return "", "", fmt.Errorf("hostname cannot be empty")
	}

	// Get list of domains owned by the account
	domainService := client.Domain()
	domains, err := domainService.List(ctx)
	if err != nil {
		return "", "", fmt.Errorf("failed to list domains: %w", err)
	}

	// Find the longest domain that is a suffix of the hostname
	var bestMatch string
	for _, d := range domains {
		if strings.HasSuffix(hostname, d.Name) {
			if len(d.Name) > len(bestMatch) {
				bestMatch = d.Name
			}
		}
	}

	if bestMatch == "" {
		return "", "", fmt.Errorf("no matching domain found for hostname %s", hostname)
	}

	// Calculate the record name
	if hostname == bestMatch {
		return bestMatch, "@", nil
	}

	name = strings.TrimSuffix(hostname, "."+bestMatch)
	return bestMatch, name, nil
}
