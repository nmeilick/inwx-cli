package output

import (
	"encoding/json"

	"github.com/nmeilick/inwx-cli/pkg/inwx"
)

type JSONFormatter struct{}

func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

func (f *JSONFormatter) FormatDNSRecords(records []inwx.DNSRecord) string {
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (f *JSONFormatter) FormatDomains(domains []inwx.Domain) string {
	data, err := json.MarshalIndent(domains, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (f *JSONFormatter) FormatAccountInfo(info *inwx.AccountInfo) string {
	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (f *JSONFormatter) FormatBackupEntries(entries []*inwx.BackupEntry) string {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}
