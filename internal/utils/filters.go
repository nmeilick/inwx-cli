package utils

import (
	"path/filepath"
	"strings"
)

func MatchWildcard(pattern, text string) bool {
	matched, _ := filepath.Match(pattern, text)
	return matched
}

func FilterByPattern(items []string, pattern string) []string {
	if pattern == "" {
		return items
	}

	var filtered []string
	for _, item := range items {
		if MatchWildcard(pattern, item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func FilterByType(records interface{}, types []string) interface{} {
	// This would be implemented based on the specific record type
	// For now, returning the input as-is
	return records
}

func ContainsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func ContainsStringIgnoreCase(slice []string, item string) bool {
	itemLower := strings.ToLower(item)
	for _, s := range slice {
		if strings.ToLower(s) == itemLower {
			return true
		}
	}
	return false
}
