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
		{
			name:     "empty version",
			version:  Version{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertString(t, tt.expected, tt.version.GetVersion())
		})
	}
}

func TestVersionString(t *testing.T) {
	tests := []struct {
		name     string
		version  Version
		expected string
	}{
		{
			name: "full information, everything committed",
			version: Version{
				ApplicationName:           "sul",
				Version:                   "1.2.3",
				VersionLabel:              "SNAPSHOT",
				VersionName:               "Buster",
				Commit:                    "abc123",
				CompilationTimestamp:      "2026-07-12",
				HasEverythingBeenCommited: true,
			},
			expected: "sul 1.2.3-SNAPSHOT \"Buster\" (2026-07-12)\nGit commit hash: abc123",
		},
		{
			name: "full information, uncommitted changes",
			version: Version{
				ApplicationName:           "sul",
				Version:                   "1.2.3",
				VersionLabel:              "SNAPSHOT",
				VersionName:               "Buster",
				Commit:                    "abc123",
				CompilationTimestamp:      "2026-07-12",
				HasEverythingBeenCommited: false,
			},
			expected: "sul 1.2.3-SNAPSHOT \"Buster\" (2026-07-12)\nGit commit hash: abc123 (some changes uncommited)",
		},
		{
			name: "minimal information",
			version: Version{
				ApplicationName:           "sul",
				Version:                   "1.2.3",
				HasEverythingBeenCommited: true,
			},
			expected: "sul 1.2.3",
		},
		{
			name: "no commit hash",
			version: Version{
				ApplicationName:           "sul",
				Version:                   "1.2.3",
				CompilationTimestamp:      "2026-07-12",
				HasEverythingBeenCommited: true,
			},
			expected: "sul 1.2.3 (2026-07-12)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertString(t, tt.expected, tt.version.String())
		})
	}
}

func TestNewVersion(t *testing.T) {
	originalVersion, originalLabel, originalName := BuildVersion, BuildVersionLabel, BuildVersionName
	defer func() {
		BuildVersion, BuildVersionLabel, BuildVersionName = originalVersion, originalLabel, originalName
	}()

	BuildVersion = "9.9.9"
	BuildVersionLabel = "TEST"
	BuildVersionName = "Codename"

	v := NewVersion("my-app")

	AssertString(t, "my-app", v.ApplicationName)
	AssertString(t, "9.9.9", v.Version)
	AssertString(t, "TEST", v.VersionLabel)
	AssertString(t, "Codename", v.VersionName)
	AssertString(t, ExtractCommitFromRuntime(), v.Commit)
	AssertString(t, ExtractBuildTimeFromRuntime(), v.CompilationTimestamp)
	AssertBool(t, ExtractBuildCommitedFromRuntime(), v.HasEverythingBeenCommited)
}
