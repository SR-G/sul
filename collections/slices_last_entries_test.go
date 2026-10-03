package collections

import (
	"testing"

	sultest "github.com/SR-G/sul/tests"
)

func TestIsStringFoundInLastEntriesOfSlice(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	var tests = []struct {
		s        string
		max      int
		expected bool
	}{
		{"d", 10, true}, {"d", 100, true},
		{"j", 10, true}, {"j", 1, true}, {"j", 2, true}, {"j", 3, true}, {"j", 4, true},
		{"i", 2, true}, {"i", 1, false},
		{"f", 1, false}, {"f", 2, false}, {"f", 3, false}, {"f", 4, false}, {"f", 5, true},
		{"a", 1, false}, {"a", 10, true},
	}
	for _, tt := range tests {
		sultest.Assert(t, tt.expected, IsStringFoundInLastEntriesOfSlice(slice, tt.s, tt.max))
	}
}
