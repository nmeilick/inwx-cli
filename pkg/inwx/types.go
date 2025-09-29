package inwx

import (
	"fmt"
	"strconv"
	"strings"
)

func formatRecord(record DNSRecord) string {
	name := record.Name
	if name == "" {
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
