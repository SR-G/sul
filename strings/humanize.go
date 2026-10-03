package strings

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func humanizeValue(i int64, typeOfValue string) string {
	switch {
	case i == 0:
		return ""
	case i == 1 || strings.HasSuffix(typeOfValue, "s"):
		return strconv.FormatInt(i, 10) + " " + typeOfValue
	default:
		return strconv.FormatInt(i, 10) + " " + typeOfValue + "s"
	}
}

func HumanizeFileSize(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

// HumanizeDuration renders a duration as e.g. "1 hour 2 minutes 3 seconds 4 ms" (millisecond precision),
// as time.Duration's default output is hard to read.
func HumanizeDuration(duration time.Duration) string {
	if duration < 0 {
		if duration == math.MinInt64 {
			duration = math.MaxInt64
		} else {
			duration = -duration
		}
		if result := HumanizeDuration(duration); result != "0 ms" {
			return "-" + result
		}
		return "0 ms"
	}

	remaining := duration.Milliseconds()
	if remaining == 0 {
		return "0 ms"
	}

	units := []struct {
		name string
		ms   int64
	}{
		{"day", 24 * 60 * 60 * 1000},
		{"hour", 60 * 60 * 1000},
		{"minute", 60 * 1000},
		{"second", 1000},
		{"ms", 1},
	}
	parts := make([]string, 0, len(units))
	for _, u := range units {
		if part := humanizeValue(remaining/u.ms, u.name); part != "" {
			parts = append(parts, part)
		}
		remaining %= u.ms
	}
	return strings.Join(parts, " ")
}
