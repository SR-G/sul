package sul

import (
	"bytes"
	"fmt"
)

/* Example of use case - note : These 4 variables are injected through Makefile at compile time

var BuildVersion = "1.0.0"
var BuildVersionLabel = "RELEASE"
var BuildCommit = ""
var BuildCompilationTimestamp = ""

var Version = sul.Version{ApplicationName: PROGRAM_NAME, Version: BuildVersion, VersionLabel: BuildVersionLabel, Commit: BuildCommit, CompilationTimestamp: BuildCompilationTimestamp}
*/

type Version struct {
	ApplicationName      string // Name of the program
	Version              string // Version, e.g. 1.0.0
	VersionLabel         string // Label of the version, like SNAPSHOT or RELEASE
	VersionName          string // Name of the version, like Buster, ...
	Commit               string // GiT commit hash
	CompilationTimestamp string // When program has been compiled
}

// Return just the base version "Version-VersionLabel" (1.0.0-SNAPSHOT, etc.)
func (v *Version) GetVersion() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s", v.Version))
	if v.VersionLabel != "" {
		buf.WriteString("-")
		buf.WriteString(v.VersionLabel)
	}
	return buf.String()
}

// Return full version, with all informations
func (v *Version) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s %s", v.ApplicationName, v.Version))
	if v.VersionLabel != "" {
		buf.WriteString("-")
		buf.WriteString(v.VersionLabel)
	}
	if v.VersionName != "" {
		buf.WriteString(" \"")
		buf.WriteString(v.VersionName)
		buf.WriteString("\"")
	}
	if v.CompilationTimestamp != "" {
		buf.WriteString(" (")
		buf.WriteString(v.CompilationTimestamp)
		buf.WriteString(")")
	}
	if v.Commit != "" {
		buf.WriteString("\nGit commit hash: ")
		buf.WriteString(v.Commit)
	}
	return buf.String()
}
