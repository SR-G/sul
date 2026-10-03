package collections

import (
	"reflect"
	"testing"
)

func TestHasKey(t *testing.T) {
	tests := []struct {
		name string
		m    any
		key  string
		want bool
	}{
		{name: "present", m: map[string]any{"name": "Ada"}, key: "name", want: true},
		{name: "missing", m: map[string]any{"name": "Ada"}, key: "age"},
		{name: "nil map", m: map[string]any(nil), key: "name"},
		{name: "unsupported map", m: map[string]string{"name": "Ada"}, key: "name"},
		{name: "nil value", m: nil, key: "name"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasKey(tt.m, tt.key); got != tt.want {
				t.Errorf("HasKey(%v, %q) = %t, want %t", tt.m, tt.key, got, tt.want)
			}
		})
	}
}

func TestFirstKey(t *testing.T) {
	if got := FirstKey(map[string]any{"only": 1}); got != "only" {
		t.Errorf("FirstKey(single-entry map) = %q, want %q", got, "only")
	}
	if got := FirstKey(map[string]any{}); got != "" {
		t.Errorf("FirstKey(empty map) = %q, want empty string", got)
	}
}

func TestExtractKey(t *testing.T) {
	tests := []struct {
		name    string
		m       any
		key     string
		want    string
		wantErr bool
	}{
		{name: "found", m: map[string]any{"count": 12}, key: "count", want: "12"},
		{name: "missing", m: map[string]any{"count": 12}, key: "other"},
		{name: "nil map", m: map[string]any(nil), key: "count"},
		{name: "nil input", key: "count"},
		{name: "unsupported type", m: "not a map", key: "count", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractKey(tt.m, tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ExtractKey() error = %v, wantErr %t", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ExtractKey() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestUniqueNonEmptyElementsOf(t *testing.T) {
	input := []string{"pear", "", "apple", "pear", "banana", "apple"}
	want := []string{"apple", "banana", "pear"}
	if got := UniqueNonEmptyElementsOf(input); !reflect.DeepEqual(got, want) {
		t.Errorf("UniqueNonEmptyElementsOf() = %v, want %v", got, want)
	}
	if got := UniqueNonEmptyElementsOf(nil); len(got) != 0 {
		t.Errorf("UniqueNonEmptyElementsOf(nil) = %v, want empty result", got)
	}
}
