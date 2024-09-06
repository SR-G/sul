package sul

import (
	"fmt"
	"sort"
)

func HasKey(m any, key string) bool {
	if m != nil {
		convertedMap := m.(map[string]any)
		if convertedMap != nil {
			for k := range convertedMap {
				if k == key {
					return true
				}
			}
		}
	}
	return false
}

func FirstKey(m map[string]interface{}) string {
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

func UniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	sort.Strings(us)
	return us
}
