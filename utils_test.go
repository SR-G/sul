package sul

import (
	"testing"
	"time"
)

func init() {
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

func TestSlug(t *testing.T) {
	AssertString(t, "this-is-a-fu-s3nt3nce-withweird@characters", Slugify("This is A Fu!! S3nT3nce ?! ... With_weird_@characters"))
}

func TestCamelCase(t *testing.T) {
	AssertString(t, "Full Lower Case", CamelCase("full lower case"))
	AssertString(t, "Full Upper Case", CamelCase("FULL UPPER CASE"))
	AssertString(t, "Mixed Case", CamelCase("mIXED cAsE"))
}

func TestSplitAny(t *testing.T) {
	var tests = []struct {
		name       string
		input      string
		separators string
		expected   int
	}{
		{"Test 1", "A,B,C,D", ", ", 4},
		{"Test 2", "A B C D", ", ", 4},
		{"Test 3", "A;B;C;D;", ", ", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SplitAny(tt.input, tt.separators)
			size := len(res)
			AssertInt(t, tt.expected, size)
		})
	}
}

func TestContains(t *testing.T) {
	AssertBool(t, true, ContainsI("ABCdefGHI", "ABC"))
	AssertBool(t, true, ContainsI("ABCdefGHI", "abc"))
	AssertBool(t, true, Contains("ABCdefGHI", "abc", true))
	AssertBool(t, false, Contains("ABCdefGHI", "abc", false))
}

func TestSlice(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	AssertBool(t, true, IsStringFoundInLastEntriesOfSlice(slice, "d", 10))
	AssertBool(t, true, IsStringFoundInLastEntriesOfSlice(slice, "d", 100))
	AssertBool(t, true, IsStringFoundInLastEntriesOfSlice(slice, "j", 10))
	AssertBool(t, true, IsStringFoundInLastEntriesOfSlice(slice, "j", 1))
	AssertBool(t, true, IsStringFoundInLastEntriesOfSlice(slice, "j", 2))
	AssertBool(t, true, IsStringFoundInLastEntriesOfSlice(slice, "j", 3))
	AssertBool(t, true, IsStringFoundInLastEntriesOfSlice(slice, "j", 4))
	AssertBool(t, true, IsStringFoundInLastEntriesOfSlice(slice, "i", 2))
	AssertBool(t, false, IsStringFoundInLastEntriesOfSlice(slice, "i", 1))
	AssertBool(t, false, IsStringFoundInLastEntriesOfSlice(slice, "f", 1))
	AssertBool(t, false, IsStringFoundInLastEntriesOfSlice(slice, "f", 2))
	AssertBool(t, false, IsStringFoundInLastEntriesOfSlice(slice, "f", 3))
	AssertBool(t, false, IsStringFoundInLastEntriesOfSlice(slice, "f", 4))
	AssertBool(t, true, IsStringFoundInLastEntriesOfSlice(slice, "f", 5))
	AssertBool(t, false, IsStringFoundInLastEntriesOfSlice(slice, "a", 1))
	AssertBool(t, true, IsStringFoundInLastEntriesOfSlice(slice, "a", 10))
}

func TestUUID(t *testing.T) {
	AssertBool(t, true, IsValidUUID(UUID()))
}
