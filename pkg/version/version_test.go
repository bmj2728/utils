package version

import "testing"

var (
	TestBuildInfoValid1 = BuildInfo{
		Version:    "v0.0.0-dev.abc123",
		CommitHash: "abc123",
		BuildDate:  "2025-07-15T10:30:00Z",
		GoVersion:  "go1.24.1",
		Platform:   "linux/amd64",
	}

	TestBuildInfoValid2 = BuildInfo{
		Version:    "v0.0.0-dev.abc1234",
		CommitHash: "abc1234kjafgadkgahslkfjsaladkghsdlgkjsdgjh",
		BuildDate:  "2025-07-15T10:30:00Z",
		GoVersion:  "go1.24.1",
		Platform:   "linux/amd64",
	}
	TestBuildInfoInvalid1 = BuildInfo{
		Version:    "v&helloW0Rld",
		CommitHash: "abc1234",
		BuildDate:  "2025-07-15T10:30:00Z",
		GoVersion:  "go1.24.1",
		Platform:   "linux/amd64",
	}
	TestBuildInfoInvalidEmpty = BuildInfo{
		Version:    "",
		CommitHash: "",
		BuildDate:  "",
		GoVersion:  "",
		Platform:   "",
	}
	TestBuildInfoDefault = BuildInfo{
		Version:    "dev",
		CommitHash: "unknown",
		BuildDate:  "unknown",
		GoVersion:  "unknown",
	}
)

func TestGetVersion(t *testing.T) {
	tests := []struct {
		name      string
		buildInfo BuildInfo
		expect    string
	}{
		{"Test GetVersionValid1", TestBuildInfoValid1, "v0.0.0-dev.abc123"},
		{"Test GetVersionValid2", TestBuildInfoValid2, "v0.0.0-dev.abc1234"},
		{"Test GetVersionInvalid1", TestBuildInfoInvalid1, "v0.0.0-dev.unknown"},
		{"Test GetVersionEmpty", TestBuildInfoInvalidEmpty, "v0.0.0-dev.unknown"},
		{"Test GetVersionDefault", TestBuildInfoDefault, "v0.0.0-dev.unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Version = tt.buildInfo.Version
			if got := GetVersion(); got != tt.expect {
				t.Errorf("GetVersion() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name      string
		buildInfo BuildInfo
		expect    string
	}{
		{"Test StringValid1",
			TestBuildInfoValid1,
			"v0.0.0-dev.abc123 (commit: abc123, built: 2025-07-15T10:30:00Z, linux/amd64)"},
		{"Test StringValid2",
			TestBuildInfoValid2,
			"v0.0.0-dev.abc1234 (commit: abc1234kjafgadkgahslkfjsaladkghsdlgkjsdgjh, built: 2025-07-15T10:30:00Z, linux/amd64)"},
		{"Test StringInvalid1",
			TestBuildInfoInvalid1,
			"v0.0.0-dev.unknown (commit: abc1234, built: 2025-07-15T10:30:00Z, linux/amd64)"},
		{"Test StringEmpty",
			TestBuildInfoInvalidEmpty,
			"v0.0.0-dev.unknown (commit: , built: , )"},
		{"Test StringDefault",
			TestBuildInfoDefault,
			"v0.0.0-dev.unknown ()"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bi := tt.buildInfo
			Version = bi.Version
			if got := bi.String(); got != tt.expect {
				t.Errorf("String() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestIsDev(t *testing.T) {
	tests := []struct {
		name      string
		buildInfo BuildInfo
		expect    bool
	}{
		{"Test StringValid1",
			TestBuildInfoValid1,
			false},
		{"Test StringValid2",
			TestBuildInfoValid2,
			false},
		{"Test StringInvalid1",
			TestBuildInfoInvalid1,
			false},
		{"Test StringEmpty",
			TestBuildInfoInvalidEmpty,
			false},
		{"Test StringDefault",
			TestBuildInfoDefault,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bi := tt.buildInfo
			Version = bi.Version
			if got := bi.IsDevelopment(); got != tt.expect {
				t.Errorf("String() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestShortCommit(t *testing.T) {
	tests := []struct {
		name      string
		buildInfo BuildInfo
		expect    string
	}{
		{"Test StringValid1",
			TestBuildInfoValid1,
			"abc123"},
		{"Test StringValid2",
			TestBuildInfoValid2,
			"abc1234"},
		{"Test StringInvalid1",
			TestBuildInfoInvalid1,
			"abc1234"},
		{"Test StringEmpty",
			TestBuildInfoInvalidEmpty,
			"unknown"},
		{"Test StringDefault",
			TestBuildInfoDefault,
			"unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bi := tt.buildInfo
			Version = bi.Version
			if got := bi.GetShortCommit(); got != tt.expect {
				t.Errorf("String() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestGetBuildInfo(t *testing.T) {
	tests := []struct {
		name      string
		buildInfo BuildInfo
	}{
		{"Test StringValid1", TestBuildInfoValid1},
		{"Test StringValid2", TestBuildInfoValid2},
		{"Test StringInvalid1", TestBuildInfoInvalid1},
		{"Test StringEmpty", TestBuildInfoInvalidEmpty},
		{"Test StringDefault", TestBuildInfoDefault},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bi := tt.buildInfo
			Version = bi.Version
			CommitHash = bi.CommitHash
			BuildDate = bi.BuildDate
			GoVersion = bi.GoVersion
			value := GetBuildInfo()
			if value.Version != bi.Version {
				t.Errorf("Version = %v, want %v", value.Version, bi.Version)
			}
			if value.CommitHash != bi.CommitHash {
				t.Errorf("CommitHash = %v, want %v", value.CommitHash, bi.CommitHash)
			}
			if value.BuildDate != bi.BuildDate {
				t.Errorf("BuildDate = %v, want %v", value.BuildDate, bi.BuildDate)
			}
			if value.GoVersion != bi.GoVersion {
				t.Errorf("GoVersion = %v, want %v", value.GoVersion, bi.GoVersion)
			}
			if bi.Platform != "" && bi.Platform != "unknown" && value.Platform != bi.Platform {
				t.Errorf("Platform = %v, want %v", value.Platform, bi.Platform)
			}

		})
	}
}

func TestIsValidSemVer(t *testing.T) {
	tests := []struct {
		name    string
		version string
		expect  bool
	}{
		{"Test Valid1", "v0.0.0-dev.abc123", true},
		{"Test Valid2", "v0.0.0-dev.abc1234", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSemVer(tt.version); got != tt.expect {
				t.Errorf("IsValidSemVer() = %v, want %v", got, tt.expect)
			}
		})
	}
}
