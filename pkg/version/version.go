// Package version provides functionality for managing and retrieving version information of the application.
// Data is populated during the build process or as outlined during dev
//
// Default Version: v0.0.0-dev.unknown
// is returned in the event Version is invalid
//
// Default CommitHash, BuildDate, GoVersion: unknown
// is the default value applied in the absence of a built binary
package version

import (
	"fmt"
	"regexp"
	"runtime"
	"runtime/debug"
)

// Build information - set by ldflags during compilation
var (
	Version    = "dev"     // e.g. "v0.1.0" or "v0.0.0-dev.abc1234"
	CommitHash = "unknown" // e.g. "abc1234"
	BuildDate  = "unknown" // e.g. "2025-07-15T10:30:00Z"
	GoVersion  = "unknown" // e.g. "go1.24.1"
)

// BuildInfo contains version and build information
type BuildInfo struct {
	Version    string `json:"version"`
	CommitHash string `json:"commit_hash"`
	BuildDate  string `json:"build_date"`
	GoVersion  string `json:"go_version"`
	Platform   string `json:"platform"`
	ModuleInfo string `json:"module_info,omitempty"`
}

// GetVersion returns just the version string
func GetVersion() string {
	if !IsValidSemVer(Version) {
		return "v0.0.0-dev.unknown"
	}
	return Version
}

// GetBuildInfo returns complete build information
func GetBuildInfo() *BuildInfo {
	platform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

	buildInfo := &BuildInfo{
		Version:    Version,
		CommitHash: CommitHash,
		BuildDate:  BuildDate,
		GoVersion:  GoVersion,
		Platform:   platform,
	}

	// Try to get module info from runtime
	if info, ok := debug.ReadBuildInfo(); ok {
		buildInfo.ModuleInfo = info.Main.Version
	}
	if buildInfo.Version == "dev" {
		buildInfo.ModuleInfo = "v0.0.0-dev.unknown"
	}

	return buildInfo
}

// String returns a human-readable version string
func (bi *BuildInfo) String() string {
	if bi.CommitHash != "unknown" {
		return fmt.Sprintf("%s (commit: %s, built: %s, %s)",
			GetVersion(), bi.CommitHash, bi.BuildDate, bi.Platform)
	}
	return fmt.Sprintf("%s (%s)", GetVersion(), bi.Platform)
}

// IsDevelopment returns true if this is a development build
func (bi *BuildInfo) IsDevelopment() bool {
	return bi.Version == "dev" || bi.CommitHash == "unknown"
}

// GetShortCommit returns the first 7 characters of the commit hash
func (bi *BuildInfo) GetShortCommit() string {
	if bi.CommitHash == "unknown" || bi.CommitHash == "" {
		return "unknown"
	}
	if len(bi.CommitHash) >= 7 {
		return bi.CommitHash[:7]
	}
	return bi.CommitHash
}

// semVerPattern is the regex pattern for semantic version validation
const semVerPattern = `^v\d+\.\d+\.\d+(-[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?$`

// IsValidSemVer checks if the given version string follows semantic versioning format
func IsValidSemVer(version string) bool {
	match, err := regexp.MatchString(semVerPattern, version)
	return match && err == nil
}
