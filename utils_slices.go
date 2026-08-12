package sul

import (
	"slices"
	"strings"

	"github.com/samber/lo"
)

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

func IsSliceContainingValueI(s []string, searched string) bool {
	for _, value := range s {
		if strings.EqualFold(searched, value) {
			return true
		}
	}
	return false
}

func IsSliceContainingValue(s []string, searched string) bool {
	for _, value := range s {
		if searched == value {
			return true
		}
	}
	return false
}

func StoreEntriesIntoNewSlice(inputs ...string) []string {
	results := make([]string, 0)
	for _, input := range inputs {
		s := strings.TrimSpace(input)
		if s != "" {
			results = append(results, s)
		}
	}
	results = lo.Uniq(results)
	slices.Sort(results)
	return results
}

func StoreEntriesIntoExistingSlice(current []string, inputs ...string) []string {
	results := make([]string, 0)
	results = append(results, current...)
	for _, input := range inputs {
		s := strings.TrimSpace(input)
		if s != "" {
			results = append(results, s)
		}
	}
	results = lo.Uniq(results)
	slices.Sort(results)
	return results
}

func MergeSlices(inputs ...[]string) []string {
	results := make([]string, 0)
	for _, input := range inputs {
		results = append(results, input...)
	}
	return results
}
