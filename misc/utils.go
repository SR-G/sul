package misc

import (
	"encoding/json"

	"os"
	"strconv"
	"strings"

	"github.com/SR-G/sul/collections"
)

const (
	DEFAULT_UNKNOWN_HOSTNAME = "unknown"
)

func IsExtensionAllowed(allowedExtensions []string, fileExtension string) bool {
	return collections.IsSliceContainingValueI(allowedExtensions, fileExtension)
}

func Unmarshal(raw json.RawMessage, destination any) error {
	err := json.Unmarshal(raw, &destination)
	if err != nil {
		return err
	}
	return nil
}

func GetEnv(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}

func GetEnvWithDefaultString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return strings.TrimSpace(defaultValue)
	}
	return strings.TrimSpace(value)
}

func GetEnvWithDefaultInt(key string, defaultValue int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return defaultValue
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}

func GetEnvWithDefaultBool(key string, defaultValue bool) bool {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return defaultValue
	}
	result, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return result
}

func ImplementsInterface[T any](obj any) bool {
	_, ok := obj.(T)
	return ok
}

func PrettyPrintJson(data any) (string, error) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

func GetHostname() string {
	return GetHostnameWithCustomDefault(DEFAULT_UNKNOWN_HOSTNAME)
}

func GetHostnameWithCustomDefault(s string) string {
	hostname, err := os.Hostname()
	if err != nil {
		return s
	}
	return hostname
}
