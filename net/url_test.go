package net

import (
	"testing"

	sultest "github.com/SR-G/sul/tests"
)

func TestExtractDomainFromURL(t *testing.T) {
	var tests = []struct {
		url      string
		expected string
	}{
		{"http://www.domain.tld/", "www.domain.tld"},
		{"https://www.domain.tld/", "www.domain.tld"},
		{"https://www.domain.tld", "www.domain.tld"},
		{"https://www.domain.tld/path/subpath/", "www.domain.tld"},
		{"www.domain.tld/path/subpath/", "www.domain.tld"},
		{"domain.tld/path/subpath/", "domain.tld"},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			actual, err := ExtractDomainFromURL(tt.url)
			sultest.Assert(t, tt.expected, actual)
			sultest.Assert(t, true, err == nil)
		})
	}
}
