package sul

import "testing"

func TestVersionGetVersion(t *testing.T) {
	tests := []struct {
		name     string
		version  Version
		expected string
	}{
		{
			name: "version without label",
			version: Version{
				Version: "1.2.3",
			},
			expected: "1.2.3",
		},
		{
			name: "version with label",
			version: Version{
				Version:      "1.2.3",
				VersionLabel: "SNAPSHOT",
			},
			expected: "1.2.3-SNAPSHOT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertString(t, tt.expected, tt.version.GetVersion())
		})
	}
}

func TestVersionString(t *testing.T) {
	version := Version{
		ApplicationName:      "sul",
		Version:              "1.2.3",
		VersionLabel:         "SNAPSHOT",
		VersionName:          "Buster",
		Commit:               "abc123",
		CompilationTimestamp: "2026-07-12",
	}

	AssertString(t, "sul 1.2.3-SNAPSHOT \"Buster\" (2026-07-12)\nGit commit hash: abc123", version.String())
}
