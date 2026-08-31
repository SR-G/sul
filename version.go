package sul

import (
	"bytes"
	"fmt"
	"runtime/debug"
	"strconv"
)

// BuildVersion and BuildVersionLabel are meant to be injected at compile time via
// -ldflags "-X github.com/<module>/sul.BuildVersion=1.2.3 -X github.com/<module>/sul.BuildVersionLabel=RELEASE"
// Commit, compilation timestamp and commit status are not injected: they are read from runtime build info instead.
var BuildVersion = "dev"
var BuildVersionLabel = ""
var BuildVersionName = ""

type Version struct {
	ApplicationName           string // Name of the program
	Version                   string // Version, e.g. 1.0.0 (injected from Makefile)
	VersionLabel              string // Label of the version, like SNAPSHOT or RELEASE (injected from Makefile)
	VersionName               string // Name of the version, like Buster, ...
	Commit                    string // GiT commit hash (retrieved at RunTime)
	CompilationTimestamp      string // When program has been compiled (retrieved at RunTime)
	HasEverythingBeenCommited bool   // Has everything been committed before building (retrieved at RunTime)
}

// NewVersion assembles a Version from the ldflags-injected BuildVersion/BuildVersionLabel
// and the commit/timestamp/modified-status read from the module's runtime build info.
func NewVersion(applicationName string) *Version {
	v := &Version{}
	v.ApplicationName = applicationName

	// injected from Makefile
	v.Version = BuildVersion
	v.VersionLabel = BuildVersionLabel
	v.VersionName = BuildVersionName

	// retrieved from binary at runtime, since go 1.18+
	v.Commit = ExtractCommitFromRuntime()
	v.CompilationTimestamp = ExtractBuildTimeFromRuntime()
	v.HasEverythingBeenCommited = ExtractBuildCommitedFromRuntime()

	return v
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
	if !v.HasEverythingBeenCommited {
		buf.WriteString(" (some changes uncommited)")
	}
	return buf.String()
}

/*
	path	bambu-lab-observer
	mod	bambu-lab-observer	(devel)
	dep	github.com/eclipse/paho.mqtt.golang	v1.4.2	h1:66wOzfUHSSI1zamx7jR6yMEI5EuHnT1G6rNA5PM12m4=
	dep	github.com/gorilla/websocket	v1.4.2	h1:+/TMaTYc4QFitKJxsQ7Yye35DkWvkdLcvGKqM+x0Ufc=
	dep	github.com/integrii/flaggy	v1.5.2	h1:bWV20MQEngo4hWhno3i5Z9ISPxLPKj9NOGNwTWb/8IQ=
	dep	github.com/juju/loggo	v1.0.0	h1:Y6ZMQOGR9Aj3BGkiWx7HBbIx6zNwNkxhVNOHU2i1bl0=
	dep	github.com/labstack/gommon	v0.4.0	h1:y7cvthEAEbU0yHOf4axH8ZG2NH8knB9iNSoTO8dyIk8=
	dep	github.com/mattn/go-colorable	v0.1.11	h1:nQ+aFkoE2TMGc0b68U2OKSexC+eq46+XwZzWXHRmPYs=
	dep	github.com/mattn/go-isatty	v0.0.14	h1:yVuAays6BHfxijgZPzw+3Zlu5yQgKGP2/hcQbHb7S9Y=
	dep	golang.org/x/net	v0.0.0-20200425230154-ff2c4b7c35a0	h1:Jcxah/M+oLZ/R4/z5RzfPzGbPXnVDPkEDtf2JnuxN+U=
	dep	golang.org/x/sync	v0.0.0-20210220032951-036812b2e83c	h1:5KslGYwFpkhGh+Q16bwMP3cOontH8FOep7tGV86Y7SQ=
	dep	golang.org/x/sys	v0.0.0-20211103235746-7861aae1554b	h1:1VkfZQv42XQlA/jchYumAnv1UPo6RgF9rJFkTgZIxO4=
	build	-compiler=gc
	build	-ldflags="-s -w"
	build	-tags=netgo
	build	CGO_ENABLED=1
	build	CGO_CFLAGS=
	build	CGO_CPPFLAGS=
	build	CGO_CXXFLAGS=
	build	CGO_LDFLAGS=
	build	GOARCH=amd64
	build	GOOS=linux
	build	GOAMD64=v1
	build	vcs=git
	build	vcs.revision=9c9fec529615cfb9ac57fe1c8791ceb695de2861
	build	vcs.time=2023-01-30T19:24:55Z
	build	vcs.modified=true
*/

func extractInformationFromRuntime(key string) string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == key {
				return setting.Value
			}
		}
	}
	return ""
}

func ExtractCommitFromRuntime() string {
	return extractInformationFromRuntime("vcs.revision")
}

func ExtractBuildTimeFromRuntime() string {
	return extractInformationFromRuntime("vcs.time")
}

func ExtractBuildCommitedFromRuntime() bool {
	s, _ := strconv.ParseBool(extractInformationFromRuntime("vcs.modified"))
	return !s
}
