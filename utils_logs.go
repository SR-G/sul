package github.com/SR-G/sul

import (
	"fmt"
	"front-matter-updater/constants"
	"path/filepath"
	"time"

	"github.com/juju/loggo"
)

func LocalLoggoFormatter(entry loggo.Entry) string {
	ts := entry.Timestamp.In(time.Local).Format(constants.TIME_FORMAT_LOGS)
	// Just get the basename from the filename
	filename := filepath.Base(entry.Filename)
	return fmt.Sprintf("%s %7s %s %s:%d %s", ts, entry.Level, entry.Module, filename, entry.Line, entry.Message)
}
