package utils

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

var (
	domainRegex = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?)*$`)
	emailRegex  = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func ValidateDomain(domain string) error {
	if domain == "" {
		return fmt.Errorf("domain cannot be empty")
	}

	if len(domain) > 253 {
		return fmt.Errorf("domain too long")
	}

	if !domainRegex.MatchString(domain) {
		return fmt.Errorf("invalid domain format")
	}

	return nil
}

func ValidateRecordType(recordType string) error {
	validTypes := []string{
		"A", "AAAA", "CNAME", "MX", "TXT", "NS", "PTR", "SRV", "CAA",
		"ALIAS", "AFSDB", "CERT", "HINFO", "HTTPS", "IPSECKEY", "LOC",
		"NAPTR", "OPENPGPKEY", "RP", "SMIMEA", "SOA", "SSHFP", "SVCB",
		"TLSA", "URI", "URL",
	}

	recordType = strings.ToUpper(recordType)
	for _, validType := range validTypes {
		if recordType == validType {
			return nil
		}
	}

	return fmt.Errorf("invalid record type: %s", recordType)
}

func ValidateRecordContent(recordType, content string) error {
	if content == "" {
		return fmt.Errorf("record content cannot be empty")
	}

	switch strings.ToUpper(recordType) {
	case "A":
		return validateIPv4(content)
	case "AAAA":
		return validateIPv6(content)
	case "CNAME", "NS", "PTR":
		return ValidateDomain(content)
	case "MX":
		return ValidateDomain(content)
	case "TXT":
		return validateTXT(content)
	case "SRV":
		return validateSRV(content)
	}

	return nil
}

func validateIPv4(ip string) error {
	// Check for IPv6 format (contains ':')
	if strings.Contains(ip, ":") {
		return fmt.Errorf("invalid IPv4 address (appears to be IPv6): %s", ip)
	}

	parsed := net.ParseIP(ip)
	if parsed == nil {
		return fmt.Errorf("invalid IP address: %s", ip)
	}
	if parsed.To4() == nil {
		return fmt.Errorf("invalid IPv4 address: %s", ip)
	}
	return nil
}

func validateIPv6(ip string) error {
	parsed := net.ParseIP(ip)
	if parsed == nil || parsed.To4() != nil {
		return fmt.Errorf("invalid IPv6 address: %s", ip)
	}
	return nil
}

func validateTXT(content string) error {
	// DNS TXT records can contain multiple 255-character strings
	// Most DNS servers support TXT records up to 4KB in practice
	// Allow up to 4096 characters to support common use cases (DKIM, SPF, etc.)
	if len(content) > 4096 {
		return fmt.Errorf("TXT record too long (max 4096 characters)")
	}
	return nil
}

func validateSRV(content string) error {
	parts := strings.Fields(content)
	if len(parts) != 4 {
		return fmt.Errorf("SRV record must have format: priority weight port target")
	}

	// Validate priority
	if _, err := strconv.Atoi(parts[0]); err != nil {
		return fmt.Errorf("invalid SRV priority: %s", parts[0])
	}

	// Validate weight
	if _, err := strconv.Atoi(parts[1]); err != nil {
		return fmt.Errorf("invalid SRV weight: %s", parts[1])
	}

	// Validate port
	port, err := strconv.Atoi(parts[2])
	if err != nil || port < 0 || port > 65535 {
		return fmt.Errorf("invalid SRV port: %s", parts[2])
	}

	// Validate target (should be a domain name)
	if err := ValidateDomain(parts[3]); err != nil {
		return fmt.Errorf("invalid SRV target: %w", err)
	}

	return nil
}

func ValidateTTL(ttl int) error {
	if ttl < 1 || ttl > 2147483647 {
		return fmt.Errorf("TTL must be between 1 and 2147483647")
	}
	return nil
}

func ValidateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func ValidateHostname(hostname string) error {
	if hostname == "" {
		return fmt.Errorf("hostname cannot be empty")
	}

	if len(hostname) > 253 {
		return fmt.Errorf("hostname too long")
	}

	// Allow wildcards in hostnames
	hostname = strings.TrimPrefix(hostname, "*.")

	return ValidateDomain(hostname)
}

func ValidatePort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("port must be between 1 and 65535")
	}
	return nil
}

func ValidateWeight(weight int) error {
	if weight < 0 || weight > 65535 {
		return fmt.Errorf("weight must be between 0 and 65535")
	}
	return nil
}

func ValidatePriority(priority int) error {
	if priority < 0 || priority > 65535 {
		return fmt.Errorf("priority must be between 0 and 65535")
	}
	return nil
}
