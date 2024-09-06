package sul

import (
	"os"
	"testing"
	"time"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("")

func init() {
	logger.SetLogLevel(loggo.INFO)
	loggo.ReplaceDefaultWriter(loggo.NewSimpleWriter(os.Stderr, LocalLoggoFormatter))
}

func TestDatesHumanRendered(t *testing.T) {

	AssertString(t, "0 ms", HumanizeDuration(0*time.Second))
	AssertString(t, "5 ms", HumanizeDuration(5*time.Millisecond))
	AssertString(t, "1 second", HumanizeDuration(1*time.Second))
	AssertString(t, "2 seconds", HumanizeDuration(2*time.Second))
	AssertString(t, "2 seconds 10 ms", HumanizeDuration(2*time.Second+10*time.Millisecond))
	AssertString(t, "12 seconds", HumanizeDuration(12*time.Second))
	AssertString(t, "1 minute", HumanizeDuration(60*time.Second))
	AssertString(t, "1 minute 12 seconds", HumanizeDuration(72*time.Second))
	AssertString(t, "2 minutes", HumanizeDuration(120*time.Second))
	AssertString(t, "1 hour", HumanizeDuration(3600*time.Second))
	AssertString(t, "3 hours", HumanizeDuration(3*3600*time.Second))
	AssertString(t, "3 hours 1 second", HumanizeDuration((3*3600+1)*time.Second))

}
