package version

import (
	"fmt"
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

	return buildInfo
}

// String returns a human-readable version string
func (bi *BuildInfo) String() string {
	if bi.CommitHash != "unknown" {
		return fmt.Sprintf("%s (commit: %s, built: %s, %s)",
			bi.Version, bi.CommitHash, bi.BuildDate, bi.Platform)
	}
	return fmt.Sprintf("%s (%s)", bi.Version, bi.Platform)
}

// IsDevelopment returns true if this is a development build
func (bi *BuildInfo) IsDevelopment() bool {
	return bi.Version == "dev" || bi.CommitHash == "unknown"
}

// GetShortCommit returns the first 7 characters of the commit hash
func (bi *BuildInfo) GetShortCommit() string {
	if len(bi.CommitHash) >= 7 {
		return bi.CommitHash[:7]
	}
	return bi.CommitHash
}
