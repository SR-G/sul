package misc

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type testStringer struct{}

func (testStringer) String() string { return "stringer" }

func TestIsExtensionAllowed(t *testing.T) {
	if !IsExtensionAllowed([]string{".png", ".jpg"}, ".PNG") {
		t.Error("IsExtensionAllowed() should match extensions without regard to case")
	}
	if IsExtensionAllowed([]string{".png", ".jpg"}, ".gif") {
		t.Error("IsExtensionAllowed() matched an extension that is not allowed")
	}
}

func TestUnmarshal(t *testing.T) {
	var got map[string]int
	if err := Unmarshal(json.RawMessage(`{"answer":42}`), &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got["answer"] != 42 {
		t.Errorf("Unmarshal() result = %v, want answer=42", got)
	}
	if err := Unmarshal(json.RawMessage(`{`), &got); err == nil {
		t.Error("Unmarshal() expected an error for malformed JSON")
	}
}

func TestGetEnv(t *testing.T) {
	t.Setenv("SUL_TEST_GET_ENV", "  value  ")
	if got := GetEnv("SUL_TEST_GET_ENV"); got != "value" {
		t.Errorf("GetEnv() = %q, want %q", got, "value")
	}
}

func TestGetEnvWithDefaultString(t *testing.T) {
	t.Setenv("SUL_TEST_DEFAULT_STRING", "  configured  ")
	if got := GetEnvWithDefaultString("SUL_TEST_DEFAULT_STRING", "fallback"); got != "configured" {
		t.Errorf("GetEnvWithDefaultString() = %q, want %q", got, "configured")
	}
	t.Setenv("SUL_TEST_DEFAULT_STRING", "")
	if got := GetEnvWithDefaultString("SUL_TEST_DEFAULT_STRING", "  fallback  "); got != "fallback" {
		t.Errorf("GetEnvWithDefaultString() default = %q, want %q", got, "fallback")
	}
}

func TestGetEnvWithDefaultInt(t *testing.T) {
	t.Setenv("SUL_TEST_DEFAULT_INT", " 23 ")
	if got := GetEnvWithDefaultInt("SUL_TEST_DEFAULT_INT", 7); got != 23 {
		t.Errorf("GetEnvWithDefaultInt() = %d, want 23", got)
	}
	t.Setenv("SUL_TEST_DEFAULT_INT", "invalid")
	if got := GetEnvWithDefaultInt("SUL_TEST_DEFAULT_INT", 7); got != 7 {
		t.Errorf("GetEnvWithDefaultInt() invalid value = %d, want default 7", got)
	}
	t.Setenv("SUL_TEST_DEFAULT_INT", " ")
	if got := GetEnvWithDefaultInt("SUL_TEST_DEFAULT_INT", 7); got != 7 {
		t.Errorf("GetEnvWithDefaultInt() empty value = %d, want default 7", got)
	}
}

func TestGetEnvWithDefaultBool(t *testing.T) {
	t.Setenv("SUL_TEST_DEFAULT_BOOL", " true ")
	if got := GetEnvWithDefaultBool("SUL_TEST_DEFAULT_BOOL", false); !got {
		t.Error("GetEnvWithDefaultBool() = false, want true")
	}
	t.Setenv("SUL_TEST_DEFAULT_BOOL", "invalid")
	if got := GetEnvWithDefaultBool("SUL_TEST_DEFAULT_BOOL", true); !got {
		t.Error("GetEnvWithDefaultBool() invalid value should use true default")
	}
	t.Setenv("SUL_TEST_DEFAULT_BOOL", " ")
	if got := GetEnvWithDefaultBool("SUL_TEST_DEFAULT_BOOL", true); !got {
		t.Error("GetEnvWithDefaultBool() empty value should use true default")
	}
}

func TestImplementsInterface(t *testing.T) {
	if !ImplementsInterface[fmt.Stringer](testStringer{}) {
		t.Error("ImplementsInterface() did not recognize fmt.Stringer")
	}
	if ImplementsInterface[fmt.Stringer](42) {
		t.Error("ImplementsInterface() recognized int as fmt.Stringer")
	}
}

func TestPrettyPrintJson(t *testing.T) {
	got, err := PrettyPrintJson(map[string]int{"answer": 42})
	if err != nil {
		t.Fatalf("PrettyPrintJson() error = %v", err)
	}
	want := "{\n  \"answer\": 42\n}"
	if got != want {
		t.Errorf("PrettyPrintJson() = %q, want %q", got, want)
	}
	if got, err := PrettyPrintJson(make(chan int)); err == nil || got != "" {
		t.Errorf("PrettyPrintJson(channel) = (%q, %v), want empty output and error", got, err)
	}
}

func TestGetHostname(t *testing.T) {
	want, err := os.Hostname()
	if err != nil {
		t.Skipf("cannot read hostname: %v", err)
	}
	if got := GetHostname(); got != want {
		t.Errorf("GetHostname() = %q, want %q", got, want)
	}
	if got := GetHostnameWithCustomDefault("fallback"); got != want {
		t.Errorf("GetHostnameWithCustomDefault() = %q, want %q", got, want)
	}
}
