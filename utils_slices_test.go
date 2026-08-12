package sul

import "testing"

func TestInputs(t *testing.T) {
	// TODO
	StoreEntriesIntoNewSlice("a", "b", "c")
	StoreEntriesIntoNewSlice([]string{"a", "b", "c"}...)
	StoreEntriesIntoExistingSlice([]string{}, "a", "b", "c")
	StoreEntriesIntoExistingSlice([]string{}, []string{"a", "b", "c"}...)
}
