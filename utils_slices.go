package sul

import "strings"

func IsStringFoundInLastEntriesOfSlice(items []string, s string, max int) bool {
	stopIndex := len(items) - max
	if stopIndex < 0 {
		stopIndex = 0
	}
	for i := len(items) - 1; i >= stopIndex; i-- {
		if strings.Contains(s, items[i]) {
			return true
		}
	}
	return false
}
