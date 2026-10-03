package net

import (
	"regexp"
	"strings"
)

func ExtractFromHTMLSingleValueInTag(html string, balise string) string {
	re, err := regexp.Compile(`(?is)<` + regexp.QuoteMeta(balise) + `[^>]*>(.*?)</` + regexp.QuoteMeta(balise) + `>`)
	if err != nil {
		return ""
	}
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return ""
}

func ExtractFromHTMLTitle(html string) string {
	return ExtractFromHTMLSingleValueInTag(html, "title")
}
