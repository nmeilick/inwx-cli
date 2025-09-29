package inwx

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func ExportZonefile(records []DNSRecord, domain string) ([]byte, error) {
	var buffer bytes.Buffer

	// Write zone header
	buffer.WriteString(fmt.Sprintf("; Zone file for %s\n", domain))
	buffer.WriteString(fmt.Sprintf("$ORIGIN %s.\n", domain))
	buffer.WriteString("$TTL 3600\n\n")

	// Write SOA record if present
	for _, record := range records {
		if record.Type == "SOA" {
			buffer.WriteString(formatZoneRecord(record))
			buffer.WriteString("\n")
			break
		}
	}

	// Write other records
	for _, record := range records {
		if record.Type != "SOA" {
			buffer.WriteString(formatZoneRecord(record))
			buffer.WriteString("\n")
		}
	}

	return buffer.Bytes(), nil
}

func ImportZonefile(data []byte, domain string) ([]DNSRecord, error) {
	var records []DNSRecord
	scanner := bufio.NewScanner(bytes.NewReader(data))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip comments and empty lines
		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "$") {
			continue
		}

		record, err := parseZoneRecord(line, domain)
		if err != nil {
			continue // Skip invalid records
		}

		if record != nil {
			records = append(records, *record)
		}
	}

	return records, scanner.Err()
}

func formatZoneRecord(record DNSRecord) string {
	name := record.Name
	if name == "" || name == "@" {
		name = "@"
	}

	parts := []string{
		fmt.Sprintf("%-30s", name),
		fmt.Sprintf("%-8s", strconv.Itoa(record.TTL)),
		"IN",
		fmt.Sprintf("%-8s", record.Type),
	}

	if record.Type == "MX" && record.Prio > 0 {
		parts = append(parts, fmt.Sprintf("%-4s", strconv.Itoa(record.Prio)))
	}

	parts = append(parts, record.Content)

	return strings.Join(parts, " ")
}

func parseZoneRecord(line, domain string) (*DNSRecord, error) {
	fields := strings.Fields(line)
	if len(fields) < 4 {
		return nil, fmt.Errorf("invalid record format")
	}

	record := &DNSRecord{
		Domain: domain,
	}

	// Parse name
	name := fields[0]
	if name == "@" {
		record.Name = ""
	} else {
		record.Name = name
	}

	// Parse TTL
	ttl, err := strconv.Atoi(fields[1])
	if err != nil {
		ttl = 3600 // Default TTL
	}
	record.TTL = ttl

	// Skip "IN" class (fields[2])

	// Parse type
	record.Type = fields[3]

	// Parse priority and content
	contentStart := 4
	if record.Type == "MX" && len(fields) > 4 {
		if prio, err := strconv.Atoi(fields[4]); err == nil {
			record.Prio = prio
			contentStart = 5
		}
	}

	// Parse content (remaining fields)
	if len(fields) > contentStart {
		record.Content = strings.Join(fields[contentStart:], " ")
	}

	return record, nil
}
