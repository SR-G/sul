package github.com/SR-G/sul

import "testing"

func AssertString(t *testing.T, wanted string, actual string) {
	if wanted != actual {
		t.Errorf("got %s, wanted %s", actual, wanted)
	}
}

func AssertInt(t *testing.T, wanted int, actual int) {
	if wanted != actual {
		t.Errorf("got %d, wanted %d", actual, wanted)
	}
}

func AssertBool(t *testing.T, wanted bool, actual bool) {
	if wanted != actual {
		t.Errorf("got %t, wanted %t", actual, wanted)
	}
}
