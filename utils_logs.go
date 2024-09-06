package sul

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/juju/loggo"
)

func LocalLoggoFormatter(entry loggo.Entry) string {
	ts := entry.Timestamp.In(time.Local).Format(TIME_FORMAT_LOGS)
	// Just get the basename from the filename
	filename := filepath.Base(entry.Filename)
	return fmt.Sprintf("%s %7s %s %s:%d %s", ts, entry.Level, entry.Module, filename, entry.Line, entry.Message)
}
