package strings

import (
	"testing"
	"time"

	sultest "github.com/SR-G/sul/tests"
)

func TestDatesHumanRendered(t *testing.T) {
	sultest.Assert(t, "0 ms", HumanizeDuration(0*time.Second))
	sultest.Assert(t, "5 ms", HumanizeDuration(5*time.Millisecond))
	sultest.Assert(t, "1 second", HumanizeDuration(1*time.Second))
	sultest.Assert(t, "2 seconds", HumanizeDuration(2*time.Second))
	sultest.Assert(t, "2 seconds 10 ms", HumanizeDuration(2*time.Second+10*time.Millisecond))
	sultest.Assert(t, "12 seconds", HumanizeDuration(12*time.Second))
	sultest.Assert(t, "1 minute", HumanizeDuration(60*time.Second))
	sultest.Assert(t, "1 minute 12 seconds", HumanizeDuration(72*time.Second))
	sultest.Assert(t, "2 minutes", HumanizeDuration(120*time.Second))
	sultest.Assert(t, "1 hour", HumanizeDuration(3600*time.Second))
	sultest.Assert(t, "3 hours", HumanizeDuration(3*3600*time.Second))
	sultest.Assert(t, "3 hours 1 second", HumanizeDuration((3*3600+1)*time.Second))
}
