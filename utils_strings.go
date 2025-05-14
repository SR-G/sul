package sul

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

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

	// Title case s
	s = cases.Title(language.AmericanEnglish).String(s)

	return s
}

func RandomString(size int) string {
	b := make([]rune, size)
	for i := range b {
		b[i] = LETTERS[rand.Intn(len(LETTERS))]
	}
	return string(b)
}

func AsSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func ReallySplit(s, sep string) []string {
	if len(s) == 0 {
		return []string{}
	}
	return strings.Split(s, sep)
}
