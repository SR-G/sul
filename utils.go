package sul

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"
)

func humanizeValue(i int64, typeOfValue string) string {
	if i == 0 {
		return ""
	} else if i == 1 {
		return fmt.Sprintf("%d "+typeOfValue, i)
	} else {
		if !strings.HasSuffix(typeOfValue, "s") {
			return fmt.Sprintf("%d "+typeOfValue+"s", i)
		} else {
			return fmt.Sprintf("%d "+typeOfValue, i)
		}
	}
}

// humanizeDuration humanizes time.Duration output to a meaningful value,
// golang's default “time.Duration“ output is badly formatted and unreadable.
func HumanizeDuration(duration time.Duration) string {
	var d string
	var h string
	var m string
	var s string
	var ms string
	if duration.Milliseconds() == 0 {
		return "0 ms"
	} else if duration.Seconds() < 1 {
		ms = humanizeValue(int64(duration.Milliseconds()), "ms")
	} else if duration.Seconds() < 60.0 {
		remainingMilliSeconds := math.Mod(float64(duration.Milliseconds()), 1000)
		s = humanizeValue(int64(duration.Seconds()), "second")
		ms = humanizeValue(int64(remainingMilliSeconds), "ms")
	} else if duration.Minutes() < 60.0 {
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		remainingMilliSeconds := math.Mod(float64(duration.Milliseconds()), 1000)
		m = humanizeValue(int64(duration.Minutes()), "minute")
		s = humanizeValue(int64(remainingSeconds), "second")
		ms = humanizeValue(int64(remainingMilliSeconds), "ms")
	} else if duration.Hours() < 24.0 {
		remainingMinutes := math.Mod(duration.Minutes(), 60)
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		remainingMilliSeconds := math.Mod(float64(duration.Milliseconds()), 1000)
		h = humanizeValue(int64(duration.Hours()), "hour")
		m = humanizeValue(int64(remainingMinutes), "minute")
		s = humanizeValue(int64(remainingSeconds), "second")
		ms = humanizeValue(int64(remainingMilliSeconds), "ms")
	} else {
		remainingHours := math.Mod(duration.Hours(), 24)
		remainingMinutes := math.Mod(duration.Minutes(), 60)
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		remainingMilliSeconds := math.Mod(float64(duration.Milliseconds()), 1000)
		d = humanizeValue(int64(duration.Hours()/24), "day")
		h = humanizeValue(int64(remainingHours), "hour")
		m = humanizeValue(int64(remainingMinutes), "minute")
		s = humanizeValue(int64(remainingSeconds), "second")
		ms = humanizeValue(int64(remainingMilliSeconds), "ms")
	}
	var result string
	if d != "" {
		result += d
	}
	if h != "" {
		result += " " + h
	}
	if m != "" {
		result += " " + m
	}
	if s != "" {
		result += " " + s
	}
	if ms != "" {
		result += " " + ms
	}
	return strings.TrimSpace(result) // strings.TrimSpace(d + " " + h + " " + m + " " + s)
}

func Unmarshal(raw json.RawMessage, destination interface{}) error {
	err := json.Unmarshal(raw, &destination)
	if err != nil {
		return err
	}
	return nil
}
