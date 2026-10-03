package collections

import (
	"fmt"
	"slices"
)

// HasKey reports whether m is a map[string]any containing key; any other type yields false.
func HasKey(m any, key string) bool {
	converted, ok := m.(map[string]any)
	if !ok {
		return false
	}
	_, found := converted[key]
	return found
}

func FirstKey(m map[string]any) string {
	for k := range m {
		return k
	}
	return ""
}

func ExtractKey(m any, key string) (string, error) {
	if m != nil {
		if _, ok := m.(map[string]any); ok {
			convertedMap := m.(map[string]any)
			for k, v := range convertedMap {
				if key == k {
					return fmt.Sprintf("%v", v), nil
				}
			}
		} else {
			return "", fmt.Errorf("unknown key [%s] %s", key, m)
		}
	}
	return "", nil
}

// UniqueNonEmptyElementsOf returns the sorted, de-duplicated, non-empty elements of s.
func UniqueNonEmptyElementsOf(s []string) []string {
	us := make([]string, 0, len(s))
	for _, elem := range s {
		if elem != "" {
			us = append(us, elem)
		}
	}
	slices.Sort(us)
	return slices.Compact(us)
}
