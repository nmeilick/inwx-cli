package output

import (
	"bytes"
	"encoding/csv"
	"strconv"

	"github.com/nmeilick/inwx-cli/pkg/inwx"
)

type CSVFormatter struct{}

func NewCSVFormatter() *CSVFormatter {
	return &CSVFormatter{}
}

func (f *CSVFormatter) FormatDNSRecords(records []inwx.DNSRecord) string {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	// Write header
	header := []string{"ID", "Domain", "Name", "Type", "Content", "TTL", "Priority"}
	_ = writer.Write(header)

	// Write records
	for _, record := range records {
		name := record.Name
		if name == "" {
			name = "@"
		}

		row := []string{
			strconv.Itoa(record.ID),
			record.Domain,
			name,
			record.Type,
			record.Content,
			strconv.Itoa(record.TTL),
			strconv.Itoa(record.Prio),
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buffer.String()
}

func (f *CSVFormatter) FormatDomains(domains []inwx.Domain) string {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	// Write header
	header := []string{"Domain", "Status"}
	_ = writer.Write(header)

	// Write domains
	for _, domain := range domains {
		row := []string{domain.Name, domain.Status}
		writer.Write(row)
	}

	writer.Flush()
	return buffer.String()
}

func (f *CSVFormatter) FormatAccountInfo(info *inwx.AccountInfo) string {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	// Write header
	header := []string{"Field", "Value"}
	writer.Write(header)

	// Write account info
	rows := [][]string{
		{"AccountID", strconv.Itoa(info.AccountID)},
		{"CustomerID", strconv.Itoa(info.CustomerID)},
		{"Username", info.Username},
		{"Email", info.Email},
	}

	for _, row := range rows {
		_ = writer.Write(row)
	}

	writer.Flush()
	return buffer.String()
}

func (f *CSVFormatter) FormatBackupEntries(entries []*inwx.BackupEntry) string {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	// Write header
	header := []string{"ID", "Timestamp", "Operation", "Domain", "Name", "Type", "Content", "TTL", "Priority"}
	writer.Write(header)

	// Write entries
	for _, entry := range entries {
		name := entry.Record.Name
		if name == "" {
			name = "@"
		}

		row := []string{
			entry.ID,
			entry.Timestamp.Format("2006-01-02 15:04:05"),
			string(entry.Operation),
			entry.Record.Domain,
			name,
			entry.Record.Type,
			entry.Record.Content,
			strconv.Itoa(entry.Record.TTL),
			strconv.Itoa(entry.Record.Prio),
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buffer.String()
}
