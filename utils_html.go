package sul

import (
	"regexp"
	"strings"
)

func ExtractFromHTMLSingleValueInBalise(html string, balise string) string {
	re := regexp.MustCompile(`(?i)<` + balise + `[^>]*>(.*?)</` + balise + `>`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return ""
}

func ExtractFromHTMLTitle(html string) string {
	return ExtractFromHTMLSingleValueInBalise(html, "title")
}
