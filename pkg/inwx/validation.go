package inwx

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

// ValidationIssue represents a DNS configuration issue
type ValidationIssue struct {
	Severity   string     // "error", "warning", "info"
	Type       string     // "orphaned_cname", "missing_target", "dangling_ref", etc.
	RecordID   int        // 0 if not specific to a record
	Record     *DNSRecord // nil if not specific to a record
	Message    string
	Suggestion string
}

// ValidationResult contains validation results for a domain
type ValidationResult struct {
	Domain  string
	Issues  []ValidationIssue
	Summary ValidationSummary
}

// ValidationSummary provides counts of issues by severity
type ValidationSummary struct {
	Total    int
	Errors   int
	Warnings int
	Info     int
}

// ValidateDomain performs logical validation checks on DNS records
func (s *DNSService) ValidateDomain(ctx context.Context, domain string) (*ValidationResult, error) {
	log.Debug().Str("domain", domain).Msg("Validating domain")

	result := &ValidationResult{Domain: domain}

	// Get all records for the domain
	records, err := s.ListRecords(ctx, WithDomainFilter(domain))
	if err != nil {
		return nil, fmt.Errorf("failed to list records: %w", err)
	}

	log.Debug().Int("count", len(records)).Msg("Retrieved records for validation")

	// Build lookup maps for efficient checking
	recordsByName := make(map[string][]DNSRecord)
	aRecords := make(map[string]bool) // Has A or AAAA record
	recordsByType := make(map[string][]DNSRecord)

	for _, record := range records {
		// Full name (e.g., "www.example.com" or "example.com" for @)
		fullName := s.getFullName(record)
		recordsByName[fullName] = append(recordsByName[fullName], record)
		recordsByType[record.Type] = append(recordsByType[record.Type], record)

		if record.Type == "A" || record.Type == "AAAA" {
			aRecords[fullName] = true
		}
	}

	// Run validation checks
	s.checkOrphanedCNAMEs(domain, records, aRecords, &result.Issues)
	s.checkMXTargets(domain, records, aRecords, &result.Issues)
	s.checkNSTargets(domain, records, aRecords, &result.Issues)
	s.checkCNAMEConflicts(recordsByName, &result.Issues)
	s.checkSRVTargets(domain, records, aRecords, &result.Issues)
	s.checkCommonBestPractices(domain, recordsByName, recordsByType, &result.Issues)
	s.checkTTLValues(records, &result.Issues)

	// Calculate summary
	for _, issue := range result.Issues {
		result.Summary.Total++
		switch issue.Severity {
		case "error":
			result.Summary.Errors++
		case "warning":
			result.Summary.Warnings++
		case "info":
			result.Summary.Info++
		}
	}

	log.Debug().
		Int("total", result.Summary.Total).
		Int("errors", result.Summary.Errors).
		Int("warnings", result.Summary.Warnings).
		Int("info", result.Summary.Info).
		Msg("Validation complete")

	return result, nil
}

// getFullName returns the fully qualified name for a record
func (s *DNSService) getFullName(record DNSRecord) string {
	if record.Name == "" || record.Name == "@" {
		return record.Domain
	}
	return record.Name + "." + record.Domain
}

// checkOrphanedCNAMEs detects CNAME records pointing to non-existent targets
func (s *DNSService) checkOrphanedCNAMEs(domain string, records []DNSRecord, aRecords map[string]bool, issues *[]ValidationIssue) {
	for _, record := range records {
		if record.Type != "CNAME" {
			continue
		}

		target := strings.TrimSuffix(record.Content, ".")

		// Only check targets within the same domain
		if strings.HasSuffix(target, "."+domain) || target == domain {
			if !aRecords[target] {
				rec := record // Create a copy for the pointer
				*issues = append(*issues, ValidationIssue{
					Severity:   "error",
					Type:       "orphaned_cname",
					RecordID:   record.ID,
					Record:     &rec,
					Message:    fmt.Sprintf("CNAME %s points to non-existent target %s", s.getFullName(record), target),
					Suggestion: fmt.Sprintf("Create an A/AAAA record for %s or update the CNAME target", target),
				})
			}
		}
	}
}

// checkMXTargets validates that MX records point to hosts with A/AAAA records
func (s *DNSService) checkMXTargets(domain string, records []DNSRecord, aRecords map[string]bool, issues *[]ValidationIssue) {
	for _, record := range records {
		if record.Type != "MX" {
			continue
		}

		// MX content format: "10 mail.example.com." or just "mail.example.com"
		parts := strings.Fields(record.Content)
		var mxTarget string

		if len(parts) >= 2 {
			mxTarget = strings.TrimSuffix(parts[1], ".")
		} else if len(parts) == 1 {
			mxTarget = strings.TrimSuffix(parts[0], ".")
		} else {
			rec := record
			*issues = append(*issues, ValidationIssue{
				Severity:   "error",
				Type:       "invalid_mx_format",
				RecordID:   record.ID,
				Record:     &rec,
				Message:    fmt.Sprintf("MX record has invalid format: %s", record.Content),
				Suggestion: "MX records should be in format 'priority hostname' (e.g., '10 mail.example.com')",
			})
			continue
		}

		// Only check targets within the same domain
		if strings.HasSuffix(mxTarget, "."+domain) || mxTarget == domain {
			if !aRecords[mxTarget] {
				rec := record
				*issues = append(*issues, ValidationIssue{
					Severity:   "error",
					Type:       "missing_mx_target",
					RecordID:   record.ID,
					Record:     &rec,
					Message:    fmt.Sprintf("MX record points to %s which has no A/AAAA record", mxTarget),
					Suggestion: fmt.Sprintf("Create an A or AAAA record for %s", mxTarget),
				})
			}
		}
	}
}

// checkNSTargets validates that NS records point to hosts with A/AAAA records
func (s *DNSService) checkNSTargets(domain string, records []DNSRecord, aRecords map[string]bool, issues *[]ValidationIssue) {
	for _, record := range records {
		if record.Type != "NS" {
			continue
		}

		nsTarget := strings.TrimSuffix(record.Content, ".")

		// Only check targets within the same domain (delegated subdomains)
		if strings.HasSuffix(nsTarget, "."+domain) || nsTarget == domain {
			if !aRecords[nsTarget] {
				rec := record
				*issues = append(*issues, ValidationIssue{
					Severity:   "warning",
					Type:       "missing_ns_target",
					RecordID:   record.ID,
					Record:     &rec,
					Message:    fmt.Sprintf("NS record points to %s which has no A/AAAA record", nsTarget),
					Suggestion: fmt.Sprintf("Create an A or AAAA record for %s (glue record)", nsTarget),
				})
			}
		}
	}
}

// checkCNAMEConflicts detects CNAME records coexisting with other record types
func (s *DNSService) checkCNAMEConflicts(recordsByName map[string][]DNSRecord, issues *[]ValidationIssue) {
	for name, recs := range recordsByName {
		if len(recs) <= 1 {
			continue
		}

		var cnameRecord *DNSRecord
		hasOtherTypes := false

		for _, rec := range recs {
			if rec.Type == "CNAME" {
				r := rec
				cnameRecord = &r
			} else {
				hasOtherTypes = true
			}
		}

		if cnameRecord != nil && hasOtherTypes {
			*issues = append(*issues, ValidationIssue{
				Severity:   "error",
				Type:       "cname_conflict",
				RecordID:   cnameRecord.ID,
				Record:     cnameRecord,
				Message:    fmt.Sprintf("CNAME record for %s conflicts with other record types (RFC 1034 violation)", name),
				Suggestion: "Remove either the CNAME or the other records for this hostname",
			})
		}

		// Check for multiple CNAMEs for the same name
		cnameCount := 0
		for _, rec := range recs {
			if rec.Type == "CNAME" {
				cnameCount++
			}
		}

		if cnameCount > 1 {
			for _, rec := range recs {
				if rec.Type == "CNAME" {
					r := rec
					*issues = append(*issues, ValidationIssue{
						Severity:   "error",
						Type:       "duplicate_cname",
						RecordID:   rec.ID,
						Record:     &r,
						Message:    fmt.Sprintf("Multiple CNAME records found for %s", name),
						Suggestion: "Only one CNAME record is allowed per hostname",
					})
				}
			}
		}
	}
}

// checkSRVTargets validates SRV record targets
func (s *DNSService) checkSRVTargets(domain string, records []DNSRecord, aRecords map[string]bool, issues *[]ValidationIssue) {
	for _, record := range records {
		if record.Type != "SRV" {
			continue
		}

		// SRV format: "priority weight port target"
		parts := strings.Fields(record.Content)
		if len(parts) != 4 {
			rec := record
			*issues = append(*issues, ValidationIssue{
				Severity:   "error",
				Type:       "invalid_srv_format",
				RecordID:   record.ID,
				Record:     &rec,
				Message:    fmt.Sprintf("SRV record has invalid format: %s", record.Content),
				Suggestion: "SRV records should be in format 'priority weight port target'",
			})
			continue
		}

		srvTarget := strings.TrimSuffix(parts[3], ".")

		// Only check targets within the same domain
		if strings.HasSuffix(srvTarget, "."+domain) || srvTarget == domain {
			if !aRecords[srvTarget] {
				rec := record
				*issues = append(*issues, ValidationIssue{
					Severity:   "warning",
					Type:       "missing_srv_target",
					RecordID:   record.ID,
					Record:     &rec,
					Message:    fmt.Sprintf("SRV record points to %s which has no A/AAAA record", srvTarget),
					Suggestion: fmt.Sprintf("Create an A or AAAA record for %s", srvTarget),
				})
			}
		}
	}
}

// checkCommonBestPractices checks for common DNS configuration best practices
func (s *DNSService) checkCommonBestPractices(domain string, recordsByName map[string][]DNSRecord, recordsByType map[string][]DNSRecord, issues *[]ValidationIssue) {
	// Check for root domain record
	rootRecords, hasRoot := recordsByName[domain]
	hasRootA := false
	if hasRoot {
		for _, rec := range rootRecords {
			if rec.Type == "A" || rec.Type == "AAAA" {
				hasRootA = true
				break
			}
		}
	}

	if !hasRootA {
		*issues = append(*issues, ValidationIssue{
			Severity:   "warning",
			Type:       "missing_root",
			Message:    fmt.Sprintf("No A/AAAA record found for root domain %s", domain),
			Suggestion: "Consider adding an A or AAAA record for the root domain",
		})
	}

	// Check for www subdomain
	wwwRecords, hasWWW := recordsByName["www."+domain]
	if !hasWWW || len(wwwRecords) == 0 {
		*issues = append(*issues, ValidationIssue{
			Severity:   "info",
			Type:       "missing_www",
			Message:    "No www subdomain found",
			Suggestion: "Consider adding a www CNAME or A record pointing to your root domain",
		})
	}

	// Check for naked domain with MX but no A/AAAA
	hasMX := len(recordsByType["MX"]) > 0
	if hasMX && !hasRootA {
		*issues = append(*issues, ValidationIssue{
			Severity:   "warning",
			Type:       "mx_without_a",
			Message:    "Domain has MX records but no A/AAAA record for root domain",
			Suggestion: "Consider adding an A or AAAA record for the root domain (required by some email systems)",
		})
	}
}

// checkTTLValues warns about unusual TTL values
func (s *DNSService) checkTTLValues(records []DNSRecord, issues *[]ValidationIssue) {
	for _, record := range records {
		if record.TTL < 60 {
			rec := record
			*issues = append(*issues, ValidationIssue{
				Severity:   "warning",
				Type:       "low_ttl",
				RecordID:   record.ID,
				Record:     &rec,
				Message:    fmt.Sprintf("Record has very low TTL (%d seconds)", record.TTL),
				Suggestion: "Low TTL values (< 60s) can cause performance issues and increased DNS queries",
			})
		}

		if record.TTL > 604800 { // 7 days
			rec := record
			*issues = append(*issues, ValidationIssue{
				Severity:   "info",
				Type:       "high_ttl",
				RecordID:   record.ID,
				Record:     &rec,
				Message:    fmt.Sprintf("Record has very high TTL (%d seconds / %.1f days)", record.TTL, float64(record.TTL)/86400),
				Suggestion: "High TTL values can delay propagation of future changes",
			})
		}
	}
}
