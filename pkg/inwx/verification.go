package inwx

import (
	"context"
	"fmt"
	"net"
	"sort"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	// DNSQueryTimeout is the timeout for DNS queries
	DNSQueryTimeout = 5 * time.Second
)

// VerificationResult contains the results of live DNS verification
type VerificationResult struct {
	Domain  string
	Records []RecordVerification
	Summary VerificationSummary
}

// NameserverResult contains the result from querying a specific nameserver
type NameserverResult struct {
	Server   string
	Type     string // "authoritative", "public", "custom"
	Response []string
	Status   string // "match", "mismatch", "missing", "error"
	Latency  time.Duration
	Error    string
}

// RecordVerification contains verification for a single record
type RecordVerification struct {
	Hostname    string
	Type        string
	Expected    []string
	Nameservers []NameserverResult
	Status      string // "match", "partial", "mismatch", "missing"
}

// VerificationSummary provides overall statistics
type VerificationSummary struct {
	Total    int
	Match    int
	Partial  int
	Mismatch int
	Missing  int
}

// VerifyDomainRecords verifies all records for a domain with better grouping
func (s *DNSService) VerifyDomainRecords(ctx context.Context, domain, name, recordType string) (*VerificationResult, error) {
	log.Debug().
		Str("domain", domain).
		Str("name", name).
		Str("type", recordType).
		Msg("Verifying DNS records")

	result := &VerificationResult{
		Domain: domain,
	}

	// 1. Get records from INWX API
	filters := []RecordFilter{WithDomainFilter(domain)}
	if name != "" {
		filters = append(filters, WithRecordName(name))
	}
	if recordType != "" {
		filters = append(filters, WithRecordType(recordType))
	}

	records, err := s.ListRecords(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("failed to list records: %w", err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("no matching records found in INWX")
	}

	// 2. Group records by hostname and type
	type recordKey struct {
		hostname string
		rtype    string
	}
	recordGroups := make(map[recordKey][]string)

	for _, record := range records {
		// Skip SOA records - not verified
		if record.Type == "SOA" {
			continue
		}

		hostname := s.getFullName(record)
		key := recordKey{hostname: hostname, rtype: record.Type}

		// Deduplicate content
		content := record.Content
		found := false
		for _, existing := range recordGroups[key] {
			if existing == content {
				found = true
				break
			}
		}
		if !found {
			recordGroups[key] = append(recordGroups[key], content)
		}
	}

	// 3. Verify each group
	for key, expectedValues := range recordGroups {
		verification := s.verifyRecordGroup(ctx, domain, key.hostname, key.rtype, expectedValues)
		result.Records = append(result.Records, verification)

		// Update summary
		result.Summary.Total++
		switch verification.Status {
		case "match":
			result.Summary.Match++
		case "partial":
			result.Summary.Partial++
		case "mismatch":
			result.Summary.Mismatch++
		case "missing":
			result.Summary.Missing++
		}
	}

	// Sort records for consistent output
	sort.Slice(result.Records, func(i, j int) bool {
		if result.Records[i].Hostname != result.Records[j].Hostname {
			return result.Records[i].Hostname < result.Records[j].Hostname
		}
		return result.Records[i].Type < result.Records[j].Type
	})

	return result, nil
}

// verifyRecordGroup verifies a single hostname+type combination
func (s *DNSService) verifyRecordGroup(ctx context.Context, domain, hostname, recordType string, expected []string) RecordVerification {
	verification := RecordVerification{
		Hostname: hostname,
		Type:     recordType,
		Expected: expected,
	}

	// Build FQDN for query
	fqdn := hostname
	if !strings.HasSuffix(fqdn, ".") {
		fqdn = fqdn + "."
	}

	// Query authoritative nameservers
	nsRecords, err := net.LookupNS(domain)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to lookup authoritative nameservers")
	} else {
		for _, ns := range nsRecords {
			nsResult := s.queryNameserver(ctx, ns.Host, fqdn, recordType, expected)
			nsResult.Type = "authoritative"
			verification.Nameservers = append(verification.Nameservers, nsResult)
		}
	}

	// Query public resolvers
	publicResolvers := []struct {
		ip   string
		name string
	}{
		{"8.8.8.8", "Google"},
		{"1.1.1.1", "Cloudflare"},
		{"9.9.9.9", "Quad9"},
	}

	for _, resolver := range publicResolvers {
		nsResult := s.queryNameserver(ctx, resolver.ip, fqdn, recordType, expected)
		nsResult.Type = "public"
		nsResult.Server = fmt.Sprintf("%s (%s)", resolver.ip, resolver.name)
		verification.Nameservers = append(verification.Nameservers, nsResult)
	}

	// Determine status
	matchCount := 0
	totalCount := len(verification.Nameservers)

	for _, ns := range verification.Nameservers {
		if ns.Status == "match" {
			matchCount++
		}
	}

	if matchCount == totalCount {
		verification.Status = "match"
	} else if matchCount > 0 {
		verification.Status = "partial"
	} else if matchCount == 0 {
		// Check if any returned data (mismatch vs missing)
		hasData := false
		for _, ns := range verification.Nameservers {
			if len(ns.Response) > 0 {
				hasData = true
				break
			}
		}
		if hasData {
			verification.Status = "mismatch"
		} else {
			verification.Status = "missing"
		}
	}

	return verification
}

// recordsMatch checks if two sets of DNS records match (order-independent)
func (s *DNSService) recordsMatch(expected, actual []string) bool {
	if len(expected) != len(actual) {
		return false
	}

	// Normalize and sort both slices
	normalize := func(records []string) []string {
		normalized := make([]string, len(records))
		for i, r := range records {
			normalized[i] = strings.ToLower(strings.TrimSuffix(r, "."))
		}
		sort.Strings(normalized)
		return normalized
	}

	expectedNorm := normalize(expected)
	actualNorm := normalize(actual)

	for i := range expectedNorm {
		if expectedNorm[i] != actualNorm[i] {
			return false
		}
	}

	return true
}

// queryNameserver queries a specific DNS server for a record
func (s *DNSService) queryNameserver(ctx context.Context, server, fqdn, recordType string, expected []string) NameserverResult {
	result := NameserverResult{
		Server: server,
	}

	start := time.Now()
	defer func() {
		result.Latency = time.Since(start)
	}()

	// Use net package for DNS lookups
	var records []string
	var err error

	// Add timeout context
	lookupCtx, cancel := context.WithTimeout(ctx, DNSQueryTimeout)
	defer cancel()

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: DNSQueryTimeout,
			}
			if !strings.Contains(server, ":") {
				server = server + ":53"
			}
			return d.DialContext(ctx, "udp", server)
		},
	}

	switch strings.ToUpper(recordType) {
	case "A":
		ips, lookupErr := resolver.LookupIP(lookupCtx, "ip4", strings.TrimSuffix(fqdn, "."))
		err = lookupErr
		for _, ip := range ips {
			if ipv4 := ip.To4(); ipv4 != nil {
				records = append(records, ipv4.String())
			}
		}

	case "AAAA":
		ips, lookupErr := resolver.LookupIP(lookupCtx, "ip6", strings.TrimSuffix(fqdn, "."))
		err = lookupErr
		for _, ip := range ips {
			if ipv6 := ip.To16(); ipv6 != nil && ip.To4() == nil {
				records = append(records, ipv6.String())
			}
		}

	case "CNAME":
		cname, lookupErr := resolver.LookupCNAME(lookupCtx, strings.TrimSuffix(fqdn, "."))
		err = lookupErr
		if cname != "" && cname != fqdn {
			records = append(records, strings.TrimSuffix(cname, "."))
		}

	case "MX":
		mxRecords, lookupErr := resolver.LookupMX(lookupCtx, strings.TrimSuffix(fqdn, "."))
		err = lookupErr
		for _, mx := range mxRecords {
			// Just store hostname for now - we'll handle comparison separately
			records = append(records, strings.TrimSuffix(mx.Host, "."))
		}

	case "TXT":
		txtRecords, lookupErr := resolver.LookupTXT(lookupCtx, strings.TrimSuffix(fqdn, "."))
		err = lookupErr
		records = txtRecords

	case "NS":
		nsRecords, lookupErr := resolver.LookupNS(lookupCtx, strings.TrimSuffix(fqdn, "."))
		err = lookupErr
		for _, ns := range nsRecords {
			records = append(records, strings.TrimSuffix(ns.Host, "."))
		}

	default:
		result.Status = "error"
		result.Error = fmt.Sprintf("Unsupported record type: %s", recordType)
		return result
	}

	if err != nil {
		if dnsErr, ok := err.(*net.DNSError); ok {
			if dnsErr.IsNotFound {
				result.Status = "missing"
				result.Error = "Record not found"
			} else if dnsErr.IsTimeout {
				result.Status = "error"
				result.Error = "Query timeout"
			} else {
				result.Status = "error"
				result.Error = dnsErr.Error()
			}
		} else {
			result.Status = "error"
			result.Error = err.Error()
		}
		return result
	}

	result.Response = records

	// Compare with expected values (order-independent)
	if len(records) == 0 {
		result.Status = "missing"
	} else if s.recordsMatch(expected, records) {
		result.Status = "match"
	} else {
		result.Status = "mismatch"
	}

	return result
}
