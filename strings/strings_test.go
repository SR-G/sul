package strings

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"regexp"
	stdstrings "strings"
	"testing"

	sultest "github.com/SR-G/sul/tests"
)

func TestSlug(t *testing.T) {
	sultest.Assert(t, "this-is-a-fu-s3nt3nce-withweird@characters", Slugify("This is A Fu!! S3nT3nce ?! ... With_weird_@characters"))
}

func TestCamelCase(t *testing.T) {
	sultest.Assert(t, "Full Lower Case", CamelCase("full lower case"))
	sultest.Assert(t, "Full Upper Case", CamelCase("FULL UPPER CASE"))
	sultest.Assert(t, "Mixed Case", CamelCase("mIXED cAsE"))
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
			sultest.Assert(t, tt.expected, len(res))
		})
	}
}

func TestContains(t *testing.T) {
	sultest.Assert(t, true, ContainsI("ABCdefGHI", "ABC"))
	sultest.Assert(t, true, ContainsI("ABCdefGHI", "abc"))
	sultest.Assert(t, true, Contains("ABCdefGHI", "abc", true))
	sultest.Assert(t, false, Contains("ABCdefGHI", "abc", false))
}

func TestUUID(t *testing.T) {
	sultest.Assert(t, true, IsValidUUID(UUID()))
}

func TestUUIDWithPotentialErrors(t *testing.T) {
	got, err := UUIDWithPotentialErrors()
	if err != nil {
		t.Fatalf("UUIDWithPotentialErrors() error = %v", err)
	}
	if !IsValidUUID(got) {
		t.Errorf("UUIDWithPotentialErrors() returned invalid UUID %q", got)
	}
	if IsValidUUID("not-a-uuid") {
		t.Error("IsValidUUID() accepted an invalid UUID")
	}
}

func TestCamelCaseSeparators(t *testing.T) {
	if got, want := CamelCase("hello_world 42"), "Hello World 42"; got != want {
		t.Errorf("CamelCase() = %q, want %q", got, want)
	}
}

func TestRandomString(t *testing.T) {
	for _, size := range []int{-1, 0} {
		if got := RandomString(size); got != "" {
			t.Errorf("RandomString(%d) = %q, want empty string", size, got)
		}
	}
	got := RandomString(128)
	if len(got) != 128 {
		t.Fatalf("RandomString(128) length = %d, want 128", len(got))
	}
	for _, r := range got {
		if !stdstrings.ContainsRune(string(LETTERS), r) {
			t.Errorf("RandomString() returned unexpected character %q", r)
			break
		}
	}
}

func TestAsSha256(t *testing.T) {
	hash := func(data []byte) string {
		sum := sha256.Sum256(data)
		return hex.EncodeToString(sum[:])
	}
	if got, want := AsSha256("hello"), hash([]byte("hello")); got != want {
		t.Errorf("AsSha256(string) = %q, want %q", got, want)
	}
	if got, want := AsSha256([]byte("hello")), hash([]byte("hello")); got != want {
		t.Errorf("AsSha256([]byte) = %q, want %q", got, want)
	}
	if got, want := AsSha256(map[string]int{"answer": 42}), hash([]byte(`{"answer":42}`)); got != want {
		t.Errorf("AsSha256(JSON value) = %q, want %q", got, want)
	}
	unmarshalable := make(chan int)
	if got, want := AsSha256(unmarshalable), hash([]byte(fmt.Sprintf("%v", unmarshalable))); got != want {
		t.Errorf("AsSha256(fallback value) = %q, want %q", got, want)
	}
}

func TestReallySplit(t *testing.T) {
	if got := ReallySplit("", ","); len(got) != 0 {
		t.Errorf("ReallySplit(empty) = %v, want empty slice", got)
	}
	if got, want := ReallySplit("a,b", ","), []string{"a", "b"}; fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("ReallySplit() = %v, want %v", got, want)
	}
}

func TestRegexpReplaceAll(t *testing.T) {
	if got := RegexpReplaceAll("a1 b2", `\d`, "_"); got != "a_ b_" {
		t.Errorf("RegexpReplaceAll() = %q, want %q", got, "a_ b_")
	}
	if got := RegexpReplaceAll("unchanged", "[", "_"); got != "unchanged" {
		t.Errorf("RegexpReplaceAll(invalid pattern) = %q, want unchanged", got)
	}
	if got := RegexpReplaceAllPattern("a1", regexp.MustCompile(`(\d)`), "[$1]"); got != "a[1]" {
		t.Errorf("RegexpReplaceAllPattern() = %q, want %q", got, "a[1]")
	}
	if got := RegexpReplaceAllPattern("unchanged", nil, "_"); got != "unchanged" {
		t.Errorf("RegexpReplaceAllPattern(nil) = %q, want unchanged", got)
	}
}

func TestReplacePlaceholders(t *testing.T) {
	got := ReplacePlaceholders("{name}: {value}; {name}", map[string]string{"name": "Ada", "value": "{name}"})
	if want := "Ada: {name}; Ada"; got != want {
		t.Errorf("ReplacePlaceholders() = %q, want %q", got, want)
	}
}

func TestFileNameHelpers(t *testing.T) {
	if got := FileNameExtensionClean(" report.pdf?download=1"); got != " report.pdf" {
		t.Errorf("FileNameExtensionClean(query) = %q, want %q", got, " report.pdf")
	}
	if got := FileNameExtensionClean(" report.pdf "); got != "report.pdf" {
		t.Errorf("FileNameExtensionClean() = %q, want %q", got, "report.pdf")
	}
	if got := FileNameWithoutExtension(filepath.Join("dir", "report.tar.gz")); got != filepath.Join("dir", "report.tar") {
		t.Errorf("FileNameWithoutExtension() = %q, want %q", got, filepath.Join("dir", "report.tar"))
	}
}
