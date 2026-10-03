package tests

import "testing"

// recorder captures failures instead of failing the real test.
type recorder struct {
	testing.TB
	failures []string
	helpers  int
}

func (r *recorder) Helper() { r.helpers++ }

func (r *recorder) Errorf(format string, args ...any) {
	r.failures = append(r.failures, format)
}

func TestAssertEqualValuesDoNotFail(t *testing.T) {
	r := &recorder{TB: t}
	Assert(r, "a", "a")
	Assert(r, 1, 1)
	Assert(r, true, true)
	Assert(r, 1.5, 1.5)
	if len(r.failures) != 0 {
		t.Errorf("expected no failure, got %d", len(r.failures))
	}
}

func TestAssertDifferentValuesFail(t *testing.T) {
	r := &recorder{TB: t}
	Assert(r, "a", "b")
	Assert(r, 1, 2)
	Assert(r, true, false)
	if len(r.failures) != 3 {
		t.Errorf("expected 3 failures, got %d", len(r.failures))
	}
}

func TestAssertMarksItselfAsHelper(t *testing.T) {
	r := &recorder{TB: t}
	Assert(r, 1, 1)
	if r.helpers != 1 {
		t.Errorf("expected Helper() to be called once, got %d", r.helpers)
	}
}

func TestAssertWorksWithStructsAndPointers(t *testing.T) {
	type point struct{ X, Y int }
	r := &recorder{TB: t}
	Assert(r, point{1, 2}, point{1, 2})
	if len(r.failures) != 0 {
		t.Errorf("equal structs should not fail")
	}
	Assert(r, point{1, 2}, point{2, 1})
	if len(r.failures) != 1 {
		t.Errorf("different structs should fail once, got %d", len(r.failures))
	}

	p := &point{1, 2}
	Assert(r, p, p)
	Assert(r, p, &point{1, 2}) // distinct pointers are not equal
	if len(r.failures) != 2 {
		t.Errorf("expected 2 failures in total, got %d", len(r.failures))
	}
}
