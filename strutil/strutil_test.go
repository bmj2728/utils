package strutil

import (
	"testing"
)

func TestNewUUID(t *testing.T) {
	uuid1 := NewUUID()
	uuid2 := NewUUID()

	// Check UUID format (very basic validation)
	if len(uuid1) != 36 || len(uuid2) != 36 {
		t.Errorf("Generated UUIDs are not of valid length")
	}

	// Ensure different UUIDs are generated each time
	if uuid1 == uuid2 {
		t.Errorf("Expected NewUUID to generate unique values, but got identical UUIDs: %s and %s", uuid1, uuid2)
	}
}

func TestRandomStringFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func(int) string
		length   int
	}{
		{"RandomString", RandomString, 10},
		{"RandomStringZeroLen", RandomString, 0},
		{"RandomHex", RandomHex, 16},
		{"RandomHexZeroLen", RandomHex, 0},
		{"RandomUrlSafe", RandomUrlSafe, 8},
		{"RandomUrlSafeZeroLen", RandomUrlSafe, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(tt.length)
			if len(result) != tt.length {
				t.Errorf("Expected length %d, got %d for %s", tt.length, len(result), tt.name)
			}

			if tt.length == 0 && result != "" {
				t.Errorf("Expected empty string for length 0, but got %s", result)
			}
		})
	}
}
