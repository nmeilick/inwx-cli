package output

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	"github.com/nmeilick/inwx-cli/pkg/inwx"
)

type TableFormatter struct {
	useColors bool
}

func NewTableFormatter() *TableFormatter {
	return &TableFormatter{
		useColors: IsColorSupported(),
	}
}

func (f *TableFormatter) SetColors(enabled bool) {
	f.useColors = enabled
}

func (f *TableFormatter) FormatDNSRecords(records []inwx.DNSRecord) string {
	if len(records) == 0 {
		return "No DNS records found"
	}

	// Calculate dynamic column widths
	widths := f.calculateDNSRecordWidths(records)

	var output strings.Builder

	// Header
	header := fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s %-*s %s",
		widths[0], "ID",
		widths[1], "DOMAIN",
		widths[2], "NAME",
		widths[3], "TYPE",
		widths[4], "TTL",
		widths[5], "PRIO",
		"CONTENT")

	if f.useColors {
		output.WriteString(color.New(color.Bold, color.FgCyan).Sprint(header))
	} else {
		output.WriteString(header)
	}
	output.WriteString("\n")

	// Separator
	totalWidth := widths[0] + widths[1] + widths[2] + widths[3] + widths[4] + widths[5] + 12 // spaces between columns
	separator := strings.Repeat("-", totalWidth)
	if f.useColors {
		output.WriteString(color.New(color.FgBlue).Sprint(separator))
	} else {
		output.WriteString(separator)
	}
	output.WriteString("\n")

	// Records
	for _, record := range records {
		name := record.Name
		if name == "" {
			name = "@"
		}

		prio := ""
		if record.Prio > 0 {
			prio = fmt.Sprintf("%d", record.Prio)
		}

		line := fmt.Sprintf("%-*d %-*s %-*s %-*s %-*d %-*s %s",
			widths[0], record.ID,
			widths[1], record.Domain,
			widths[2], name,
			widths[3], record.Type,
			widths[4], record.TTL,
			widths[5], prio,
			record.Content)

		if f.useColors {
			// Color code by record type
			switch record.Type {
			case "A":
				line = color.New(color.FgGreen).Sprint(line)
			case "AAAA":
				line = color.New(color.FgGreen).Sprint(line)
			case "CNAME":
				line = color.New(color.FgYellow).Sprint(line)
			case "MX":
				line = color.New(color.FgMagenta).Sprint(line)
			case "TXT":
				line = color.New(color.FgCyan).Sprint(line)
			case "NS":
				line = color.New(color.FgBlue).Sprint(line)
			default:
				line = color.New(color.FgWhite).Sprint(line)
			}
		}

		output.WriteString(line)
		output.WriteString("\n")
	}

	return output.String()
}

func (f *TableFormatter) calculateDNSRecordWidths(records []inwx.DNSRecord) []int {
	// Minimum widths for headers
	widths := []int{5, 6, 8, 8, 8, 6} // ID, DOMAIN, NAME, TYPE, TTL, PRIO

	for _, record := range records {
		name := record.Name
		if name == "" {
			name = "@"
		}

		prio := ""
		if record.Prio > 0 {
			prio = fmt.Sprintf("%d", record.Prio)
		}

		// Update widths if current record has longer values
		if len(fmt.Sprintf("%d", record.ID)) > widths[0] {
			widths[0] = len(fmt.Sprintf("%d", record.ID))
		}
		if len(record.Domain) > widths[1] {
			widths[1] = len(record.Domain)
		}
		if len(name) > widths[2] {
			widths[2] = len(name)
		}
		if len(record.Type) > widths[3] {
			widths[3] = len(record.Type)
		}
		if len(fmt.Sprintf("%d", record.TTL)) > widths[4] {
			widths[4] = len(fmt.Sprintf("%d", record.TTL))
		}
		if len(prio) > widths[5] {
			widths[5] = len(prio)
		}
	}

	return widths
}

func (f *TableFormatter) FormatBackupEntries(entries []*inwx.BackupEntry) string {
	if len(entries) == 0 {
		return "No backup entries found"
	}

	// Calculate dynamic column widths
	widths := f.calculateBackupWidths(entries)

	var output strings.Builder

	// Header
	header := fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s %-*s %s",
		widths[0], "ID",
		widths[1], "TIMESTAMP",
		widths[2], "OPERATION",
		widths[3], "DOMAIN",
		widths[4], "NAME",
		widths[5], "TYPE",
		"CONTENT")

	if f.useColors {
		output.WriteString(color.New(color.Bold, color.FgCyan).Sprint(header))
	} else {
		output.WriteString(header)
	}
	output.WriteString("\n")

	// Separator
	totalWidth := widths[0] + widths[1] + widths[2] + widths[3] + widths[4] + widths[5] + 12 // spaces
	separator := strings.Repeat("-", totalWidth)
	if f.useColors {
		output.WriteString(color.New(color.FgBlue).Sprint(separator))
	} else {
		output.WriteString(separator)
	}
	output.WriteString("\n")

	// Entries
	for _, entry := range entries {
		name := entry.Record.Name
		if name == "" {
			name = "@"
		}

		line := fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s %-*s %s",
			widths[0], entry.ID[:8], // Show first 8 chars of ID
			widths[1], entry.Timestamp.Format("2006-01-02 15:04"),
			widths[2], string(entry.Operation),
			widths[3], entry.Record.Domain,
			widths[4], name,
			widths[5], entry.Record.Type,
			entry.Record.Content)

		if f.useColors {
			switch entry.Operation {
			case "create":
				line = color.New(color.FgGreen).Sprint(line)
			case "update":
				line = color.New(color.FgYellow).Sprint(line)
			case "delete":
				line = color.New(color.FgRed).Sprint(line)
			default:
				line = color.New(color.FgWhite).Sprint(line)
			}
		}

		output.WriteString(line)
		output.WriteString("\n")
	}

	return output.String()
}

func (f *TableFormatter) calculateBackupWidths(entries []*inwx.BackupEntry) []int {
	// Minimum widths for headers
	widths := []int{8, 16, 9, 6, 4, 4} // ID, TIMESTAMP, OPERATION, DOMAIN, NAME, TYPE

	for _, entry := range entries {
		name := entry.Record.Name
		if name == "" {
			name = "@"
		}

		// Update widths if current entry has longer values
		if len(entry.Record.Domain) > widths[3] {
			widths[3] = len(entry.Record.Domain)
		}
		if len(name) > widths[4] {
			widths[4] = len(name)
		}
		if len(entry.Record.Type) > widths[5] {
			widths[5] = len(entry.Record.Type)
		}
	}

	return widths
}

func (f *TableFormatter) FormatDomains(domains []inwx.Domain) string {
	if len(domains) == 0 {
		return "No domains found"
	}

	// Calculate dynamic column widths
	widths := f.calculateDomainWidths(domains)

	var output strings.Builder

	// Header
	header := fmt.Sprintf("%-*s %s", widths[0], "DOMAIN", "STATUS")
	if f.useColors {
		output.WriteString(color.New(color.Bold, color.FgCyan).Sprint(header))
	} else {
		output.WriteString(header)
	}
	output.WriteString("\n")

	// Separator
	totalWidth := widths[0] + widths[1] + 1 // space between columns
	separator := strings.Repeat("-", totalWidth)
	if f.useColors {
		output.WriteString(color.New(color.FgBlue).Sprint(separator))
	} else {
		output.WriteString(separator)
	}
	output.WriteString("\n")

	// Domains
	for _, domain := range domains {
		line := fmt.Sprintf("%-*s %s", widths[0], domain.Name, domain.Status)

		if f.useColors {
			switch domain.Status {
			case "OK":
				line = color.New(color.FgGreen).Sprint(line)
			case "PENDING":
				line = color.New(color.FgYellow).Sprint(line)
			case "EXPIRED":
				line = color.New(color.FgRed).Sprint(line)
			default:
				line = color.New(color.FgWhite).Sprint(line)
			}
		}

		output.WriteString(line)
		output.WriteString("\n")
	}

	return output.String()
}

func (f *TableFormatter) calculateDomainWidths(domains []inwx.Domain) []int {
	// Minimum widths for headers
	widths := []int{6, 6} // DOMAIN, STATUS

	for _, domain := range domains {
		if len(domain.Name) > widths[0] {
			widths[0] = len(domain.Name)
		}
		if len(domain.Status) > widths[1] {
			widths[1] = len(domain.Status)
		}
	}

	return widths
}

func (f *TableFormatter) FormatAccountInfo(info *inwx.AccountInfo) string {
	var output strings.Builder

	if f.useColors {
		output.WriteString(color.New(color.Bold, color.FgCyan).Sprint("Account Information"))
	} else {
		output.WriteString("Account Information")
	}
	output.WriteString("\n")

	separator := strings.Repeat("-", 30)
	if f.useColors {
		output.WriteString(color.New(color.FgBlue).Sprint(separator))
	} else {
		output.WriteString(separator)
	}
	output.WriteString("\n")

	lines := []string{
		fmt.Sprintf("Account ID:  %d", info.AccountID),
		fmt.Sprintf("Customer ID: %d", info.CustomerID),
		fmt.Sprintf("Username:    %s", info.Username),
		fmt.Sprintf("Email:       %s", info.Email),
	}

	for _, line := range lines {
		if f.useColors {
			output.WriteString(color.New(color.FgWhite).Sprint(line))
		} else {
			output.WriteString(line)
		}
		output.WriteString("\n")
	}

	return output.String()
}
