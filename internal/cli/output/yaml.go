package output

import (
	"gopkg.in/yaml.v3"

	"github.com/nmeilick/inwx-cli/pkg/inwx"
)

type YAMLFormatter struct{}

func NewYAMLFormatter() *YAMLFormatter {
	return &YAMLFormatter{}
}

func (f *YAMLFormatter) FormatDNSRecords(records []inwx.DNSRecord) string {
	data, err := yaml.Marshal(records)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (f *YAMLFormatter) FormatDomains(domains []inwx.Domain) string {
	data, err := yaml.Marshal(domains)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (f *YAMLFormatter) FormatAccountInfo(info *inwx.AccountInfo) string {
	data, err := yaml.Marshal(info)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (f *YAMLFormatter) FormatBackupEntries(entries []*inwx.BackupEntry) string {
	data, err := yaml.Marshal(entries)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
