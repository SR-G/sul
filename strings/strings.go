package strings

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	mrand "math/rand/v2"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

var (
	LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func Slugify(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)
	// Replace spaces with hyphens
	s = strings.ReplaceAll(s, " ", "-")
	// Remove special characters
	s = strings.Map(func(r rune) rune {
		if r == '?' || r == '$' || r == '_' || r == '.' || r == '~' || r == '!' || r == '*' || r == '\'' || r == '(' || r == ')' || r == '[' || r == ']' || r == '"' {
			return -1
		}
		return r
	}, s)

	regexp, _ := regexp.Compile("-+")
	s = regexp.ReplaceAllString(s, "-")
	return s
}

func UUID() string {
	uuid, _ := uuid.NewRandom()
	return uuid.String()
}

func UUIDWithPotentialErrors() (string, error) {
	uuid, err := uuid.NewRandom()
	return uuid.String(), err

}

func IsValidUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

func CamelCase(s string) string {
	// Remove all characters that are not alphanumeric or spaces or underscores
	// s = regexp.MustCompile("[^a-zA-Z0-9_ ,.-]+").ReplaceAllString(s, "")

	// Replace all underscores with spaces
	s = strings.ReplaceAll(s, "_", " ")

	// Title case: first letter of each word upper, the rest lower
	var sb strings.Builder
	inWord := false
	for _, r := range s {
		if inWord {
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(unicode.ToUpper(r))
		}
		inWord = unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\''
	}

	return sb.String()
}

// RandomString returns a cryptographically secure random string of letters; empty if size <= 0.
// If the system entropy source fails, it degrades to math/rand rather than panicking.
func RandomString(size int) string {
	if size <= 0 || len(LETTERS) == 0 {
		return ""
	}
	// Reject bytes above the largest multiple of len(LETTERS) to avoid modulo bias.
	limit := 256 - 256%len(LETTERS)
	b := make([]rune, 0, size)
	buf := make([]byte, size)
	for len(b) < size {
		if _, err := rand.Read(buf); err != nil {
			for i := range buf {
				buf[i] = byte(mrand.UintN(256))
			}
		}
		for _, v := range buf {
			if int(v) < limit && len(b) < size {
				b = append(b, LETTERS[int(v)%len(LETTERS)])
			}
		}
	}
	return string(b)
}

// AsSha256 returns the hex SHA-256 of o: strings and []byte are hashed as-is, anything else via its JSON
// encoding (stable for maps and structs), falling back to fmt's %v when it can't be marshaled.
func AsSha256(o any) string {
	var data []byte
	switch v := o.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		if encoded, err := json.Marshal(o); err == nil {
			data = encoded
		} else {
			data = []byte(fmt.Sprintf("%v", o))
		}
	}
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum)
}

func ReallySplit(s, sep string) []string {
	if len(s) == 0 {
		return []string{}
	}
	return strings.Split(s, sep)
}

func SplitAny(s string, separators string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(separators, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func Contains(s string, searched string, insensitive bool) bool {
	if insensitive {
		return ContainsI(s, searched)
	} else {
		return strings.Contains(s, searched)
	}
}

func ContainsI(s string, searched string) bool {
	return strings.Contains(
		strings.ToLower(s),
		strings.ToLower(searched),
	)
}

// RegexpReplaceAll compiles r on each call; an invalid pattern leaves s unchanged.
// Prefer RegexpReplaceAllPattern with a precompiled pattern in hot paths.
func RegexpReplaceAll(s string, r string, repl string) string {
	re, err := regexp.Compile(r)
	if err != nil {
		return s
	}
	return RegexpReplaceAllPattern(s, re, repl)
}

// RegexpReplaceAllPattern is RegexpReplaceAll with an already compiled (hence valid) pattern.
func RegexpReplaceAllPattern(s string, re *regexp.Regexp, repl string) string {
	if re == nil {
		return s
	}
	return re.ReplaceAllString(s, repl)
}

// ReplacePlaceholders replaces every {key} in a single pass, so substituted values are never re-expanded.
func ReplacePlaceholders(template string, replacements map[string]string) string {
	oldnew := make([]string, 0, len(replacements)*2)
	for key, value := range replacements {
		oldnew = append(oldnew, "{"+key+"}", value)
	}
	return strings.NewReplacer(oldnew...).Replace(template)
}

func FileNameExtensionClean(s string) string {
	if strings.Contains(s, "?") {
		tokens := strings.Split(s, "?")
		if len(tokens) > 0 {
			return tokens[0]
		}
	}
	return strings.TrimSpace(s)
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
