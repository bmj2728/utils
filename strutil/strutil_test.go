package strutil

import (
	"testing"
)

func TestNewUUID(t *testing.T) {
	uuid1 := NewUUID()
	uuid2 := NewUUID()
	uuid3 := NewUUIDV7()
	uuid4 := NewUUIDV7()

	// Check UUID format (very basic validation)
	if len(uuid1) != 36 || len(uuid2) != 36 || len(uuid3) != 36 || len(uuid4) != 36 {
		t.Errorf("Generated UUIDs are not of valid length")
	}

	// Ensure different UUIDs are generated each time
	if uuid1 == uuid2 {
		t.Errorf("Expected NewUUID to generate unique values, but got identical UUIDs: %s and %s", uuid1, uuid2)
	}

	if uuid3 == uuid4 {
		t.Errorf("Expected NewUUIDV7 to generate unique values, but got identical UUIDs: %s and %s", uuid3, uuid4)
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
func TestIsUUID(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"ValidUUID", "550e8400-e29b-41d4-a716-446655440000", true},
		{"ValidUUIDUppercase", "550E8400-E29B-41D4-A716-446655440000", true},
		{"ValidUUIDURN", "urn:uuid:550E8400-E29B-41D4-A716-446655440000", true},
		{"ValidUUIDMS", "{550E8400-E29B-41D4-A716-446655440000}", true},
		{"ValidUUIDNoDash", "550E8400E29B41D4A716446655440000", true},
		{"InvalidUUID_TooShort", "550e8400-e29b-41d4-a716-44665544", false},
		{"InvalidUUID_ExtraChars", "550e8400-e29b-41d4-a716-446655440000-extra", false},
		{"InvalidUUID_Malformed", "550e8400-e29b-41d4ZXa716-446655440000", false},
		{"CompletelyInvalid", "not-a-uuid", false},
		{"EmptyString", "", false},
		{"NilUUID", "00000000-0000-0000-0000-000000000000", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsUUID(tt.input)
			if result != tt.expected {
				t.Errorf("IsUUID(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCleanWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"NoWhitespace", "hello world", "helloworld"},
		{"Whitespace", "hello world ", "helloworld"},
		{"Tabs", "hello\tworld", "helloworld"},
		{"Newlines", "hello\nworld", "helloworld"},
		{"CarriageReturns", "hello\rworld", "helloworld"},
		{"MixedWhitespace", "hello\tworld\n\r", "helloworld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := cleanWhitespace(tt.input)
			result := CleanWhitespace(tt.input)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("cleanWhitespace(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
