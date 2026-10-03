// Package tests provides assertion helpers for tests; kept separate so that
// the main package does not link the testing package into production binaries.
package tests

import "testing"

// Assert fails the test when actual differs from wanted.
func Assert[T comparable](t testing.TB, wanted T, actual T) {
	t.Helper()
	if wanted != actual {
		t.Errorf("got %v, wanted %v", actual, wanted)
	}
}
