package net

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	PROTOCOL_HTTP  = "http://"
	PROTOCOL_HTTPS = "https://"
)

func IsURLValid(rawURL string) error {
	if strings.TrimSpace(rawURL) == "" {
		return fmt.Errorf("URL is required")
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("invalid URL '%s': %w", rawURL, err)
	}
	if scheme := strings.ToLower(u.Scheme); scheme != "http" && scheme != "https" {
		return fmt.Errorf("invalid URL '%s': must start with '%s' or '%s'", rawURL, PROTOCOL_HTTP, PROTOCOL_HTTPS)
	}
	if u.Hostname() == "" {
		return fmt.Errorf("invalid URL '%s': host is required", rawURL)
	}
	return nil
}

func AddProtocolToURLIfNeeded(rawURL string) string {
	rawURL = strings.TrimSpace(rawURL)
	if !strings.HasPrefix(rawURL, PROTOCOL_HTTP) && !strings.HasPrefix(rawURL, PROTOCOL_HTTPS) {
		rawURL = PROTOCOL_HTTPS + rawURL
	}
	return rawURL
}

func ExtractDomainFromURL(s string) (string, error) {
	u, err := url.Parse(AddProtocolToURLIfNeeded(s))
	if err != nil {
		return "", err
	}
	return u.Hostname(), nil
}

func FetchURLContent(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
