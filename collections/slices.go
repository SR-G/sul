package collections

import (
	"slices"
	"strings"
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

// Alias for slices.Contains, just allows to have the case sensitivie/insensitive searches with the same kind of namings
func IsSliceContainingValue(s []string, searched string) bool {
	return slices.Contains(s, searched)
}

func StoreEntriesIntoNewSlice(inputs ...string) []string {
	results := make([]string, 0)
	for _, input := range inputs {
		s := strings.TrimSpace(input)
		if s != "" {
			results = append(results, s)
		}
	}
	slices.Sort(results)
	return slices.Compact(results)
}

func StoreEntriesIntoExistingSlice(current []string, inputs ...string) []string {
	results := StoreEntriesIntoNewSlice(current...)
	results = append(results, StoreEntriesIntoNewSlice(inputs...)...)
	slices.Sort(results)
	return slices.Compact(results)
}

// Merge several strings slices
// Deprecated: MergeSlices` is `slices.Concat` (Go 1.22+).
func MergeSlices(inputs ...[]string) []string {
	results := make([]string, 0)
	for _, input := range inputs {
		results = append(results, input...)
	}
	return results
}
