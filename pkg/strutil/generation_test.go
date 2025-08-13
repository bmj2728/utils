package strutil

import (
	"errors"
	"strings"
	"testing"

	errors2 "github.com/bmj2728/utils/pkg/internal/errors"

	"github.com/google/uuid"
)

func TestNewUUID(t *testing.T) {
	uuid1 := GenerateUUID()
	uuid2 := GenerateUUID()
	uuid3 := GenerateUUIDV7()
	uuid4 := GenerateUUIDV7()
	uuid5 := NewUUID()
	uuid6 := NewUUID()
	uuid7 := NewUUIDV7()
	uuid8 := NewUUIDV7()

	uuidLength := 36

	// Check UUID format (very basic validation)
	if len(uuid1) != uuidLength ||
		len(uuid2) != uuidLength ||
		len(uuid3) != uuidLength ||
		len(uuid4) != uuidLength ||
		len(uuid5.String()) != uuidLength ||
		len(uuid6.String()) != uuidLength ||
		len(uuid7.String()) != uuidLength ||
		len(uuid8.String()) != uuidLength {
		t.Errorf("Generated UUIDs are not of valid length")
	}

	e1 := uuid.Validate(uuid1)
	e2 := uuid.Validate(uuid2)
	e3 := uuid.Validate(uuid3)
	e4 := uuid.Validate(uuid4)
	e5 := uuid.Validate(uuid5.String())
	e6 := uuid.Validate(uuid6.String())
	e7 := uuid.Validate(uuid7.String())
	e8 := uuid.Validate(uuid8.String())
	if e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil || e6 != nil || e7 != nil || e8 != nil {
		t.Errorf("Generated UUIDs are not valid UUIDs")
	}

	// Ensure different UUIDs are generated each time
	if uuid1 == uuid2 {
		t.Errorf("Expected GenerateUUID to generate unique values, but got identical UUIDs: %s and %s", uuid1, uuid2)
	}

	if uuid3 == uuid4 {
		t.Errorf("Expected GenerateUUIDV7 to generate unique values, but got identical UUIDs: %s and %s", uuid3, uuid4)
	}

	if uuid5.String() == uuid6.String() {
		t.Errorf("Expected NewUUID() to generate unique values, but got identical UUIDs: %s and %s",
			uuid5.String(),
			uuid6.String(),
		)
	}

	if uuid7.String() == uuid8.String() {
		t.Errorf("Expected NewUUIDV7 to generate unique values, but got identical UUIDs: %s and %s",
			uuid7.String(),
			uuid8.String(),
		)
	}
}

func TestRandomStringFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func(int) string
		length   int
		charSet  CharacterSet
	}{
		{"RandomAlphaString", RandomAlphaNumericString, 10, AlphaNumericChars},
		{"RandomStringZeroLen", RandomAlphaNumericString, 0, AlphaNumericChars},
		{"RandomStringNegLen", RandomAlphaNumericString, -1, AlphaNumericChars},
		{"RandomAlphaString", RandomAlphaString, 10, Alpha},
		{"RandomStringZeroLen", RandomAlphaString, 0, Alpha},
		{"RandomStringNegLen", RandomAlphaString, -1, Alpha},
		{"RandomHex", RandomHex, 16, HexChars},
		{"RandomHexZeroLen", RandomHex, 0, HexChars},
		{"RandomHexNegLen", RandomHex, -1, HexChars},
		{"RandomUrlSafe", RandomUrlSafe, 8, URLSafe},
		{"RandomUrlSafeZeroLen", RandomUrlSafe, 0, URLSafe},
		{"RandomUrlSafeNegLen", RandomUrlSafe, -1, URLSafe},
		{"RandomAlphaString", randomAlphaNumericString, 10, AlphaNumericChars},
		{"RandomStringZeroLen", randomAlphaNumericString, 0, AlphaNumericChars},
		{"RandomStringNegLen", randomAlphaNumericString, -1, AlphaNumericChars},
		{"RandomAlphaString", randomAlphaString, 10, Alpha},
		{"RandomStringZeroLen", randomAlphaString, 0, Alpha},
		{"RandomStringNegLen", randomAlphaString, -1, Alpha},
		{"RandomHex", randomHex, 16, HexChars},
		{"RandomHexZeroLen", randomHex, 0, HexChars},
		{"RandomHexNegLen", randomHex, -1, HexChars},
		{"RandomUrlSafe", randomURLSafe, 8, URLSafe},
		{"RandomUrlSafeZeroLen", randomURLSafe, 0, URLSafe},
		{"RandomUrlSafeNegLen", randomURLSafe, -1, URLSafe},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(tt.length)
			if tt.length > 0 && len(result) != tt.length {
				t.Errorf("Expected length %d, got %d for %s", tt.length, len(result), tt.name)
			}
			if tt.length == 0 && result != "" {
				t.Errorf("Expected empty string for length 0, but got %s", result)
			}
			if tt.length < 0 && result != "" {
				t.Errorf("Expected empty string for negative length, but got %s", result)
			}
			for _, c := range result {
				if !strings.Contains(string(tt.charSet), string(c)) {
					t.Errorf("Expected %s to be in custom char set for %s", string(c), tt.name)
				}
			}
		})
	}
}

func TestRandomStringWithCharSet(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		charSet CharacterSet
	}{
		{"RandomAlpha", 10, Alpha},
		{"RandomAlphaZeroLen", 0, Alpha},
		{"RandomAlphaNegLen", -1, Alpha},
		{"RandomAlphaNumeric", 10, AlphaNumericChars},
		{"RandomAlphaNumericZeroLen", 0, AlphaNumericChars},
		{"RandomAlphaNumericNegLen", -1, AlphaNumericChars},
		{"RandomHex", 10, HexChars},
		{"RandomHexZeroLen", 0, HexChars},
		{"RandomHexNegLen", -1, HexChars},
		{"RandomUrlSafe", 8, URLSafe},
		{"RandomUrlSafeZeroLen", 0, URLSafe},
		{"RandomUrlSafeNegLen", -1, URLSafe},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := randomFromCharset(tt.length, tt.charSet)
			result2 := RandomString(tt.length, tt.charSet)
			if tt.length > 0 && (len(result) != tt.length || len(result2) != tt.length) {
				t.Errorf("Expected length %d, got %d for %s", tt.length, len(result), tt.name)
			}
			if tt.length == 0 && (result != "" || result2 != "") {
				t.Errorf("Expected empty string for length 0, but got %s", result)
			}
			if tt.length < 0 && (result != "" || result2 != "") {
				t.Errorf("Expected empty string for negative length, but got %s", result)
			}
			for _, c := range result {
				if !strings.Contains(string(tt.charSet), string(c)) {
					t.Errorf("Expected %s to be in custom char set for %s", string(c), tt.name)
				}
			}
			for _, c := range result2 {
				if !strings.Contains(string(tt.charSet), string(c)) {
					t.Errorf("Expected %s to be in custom char set for %s", string(c), tt.name)
				}
			}
		})
	}
}

func TestRandomStringFromCustomCharset(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		charSet string
	}{
		{"RandomFromCust1", 10, "abc"},
		{"RandomFromCust2", 10, "pn;YUOIVBOIYU.H5NVAE48657wtsjgf4725723@!##!%$@^WUFNGKXLM3UIOR"},
		{"RandomFromCust3", 0, "abc"},
		{"RandomFromCust4", -1, "abc"},
		{"RandomFromCust5",
			10,
			"zxcvbnm,./asdfghjkl;'qwertyuiop[]1234567890-=!@#$%^&*()_+{}:><|~"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := randomFromCustomCharset(tt.length, tt.charSet)
			result2 := RandomStringFromCustomCharset(tt.length, tt.charSet)
			if tt.length > 0 && (len(result) != tt.length || len(result2) != tt.length) {
				t.Errorf("Expected length %d, got %d for %s", tt.length, len(result), tt.name)
			}
			if tt.length == 0 && (result != "" || result2 != "") {
				t.Errorf("Expected empty string for length 0, but got %s", result)
			}
			if tt.length < 0 && (result != "" || result2 != "") {
				t.Errorf("Expected empty string for negative length, but got %s", result)
			}
			for _, c := range result {
				if !strings.Contains(tt.charSet, string(c)) {
					t.Errorf("Expected %s to be in custom char set for %s", string(c), tt.name)
				}
			}
			for _, c := range result2 {
				if !strings.Contains(tt.charSet, string(c)) {
					t.Errorf("Expected %s to be in custom char set for %s", string(c), tt.name)
				}
			}
		})
	}
}

func TestBasicConstructor(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty", "", ""},
		{"Normal", "hello world", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := New(tt.input).String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q for %s", tt.expected, result, tt.name)
			}
		})
	}
}

func TestStringBuilderConstructors(t *testing.T) {
	tests := []struct {
		name     string
		function func(int) *StringBuilder
		length   int
		charSet  CharacterSet
	}{
		{"NewRandomAlpha", NewRandomAlpha, 10, Alpha},
		{"NewRandomAlphaZero", NewRandomAlpha, 0, Alpha},
		{"NewRandomAlphaNeg", NewRandomAlpha, -1, Alpha},
		{"NewRandomAlphaNumeric", NewRandomAlphaNumeric, 10, AlphaNumericChars},
		{"NewRandomAlphaNumericZero", NewRandomAlphaNumeric, 0, AlphaNumericChars},
		{"NewRandomAlphaNumericNeg", NewRandomAlphaNumeric, -1, AlphaNumericChars},
		{"NewRandomHex", NewRandomHex, 6, HexChars},
		{"NewRandomHexZero", NewRandomHex, 0, HexChars},
		{"NewRandomHexNeg", NewRandomHex, -1, HexChars},
		{"NewRandomURLSafe", NewRandomURLSafe, 8, URLSafe},
		{"NewRandomURLSafeZero", NewRandomURLSafe, 0, URLSafe},
		{"NewRandomURLSafeNeg", NewRandomURLSafe, -1, URLSafe},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(tt.length)
			if result == nil {
				t.Errorf("Expected non-nil result for %s", tt.name)
			}
			if result.String() == "" && tt.length > 0 {
				t.Errorf("Expected non-empty string for %s", tt.name)
			}
			if result.String() != "" && tt.length == 0 {
				t.Errorf("Expected empty string for %s", tt.name)
			}
			if result.String() != "" && tt.length < 0 {
				t.Errorf("Expected empty string for %s", tt.name)
			}
			for _, c := range result.String() {
				if !strings.Contains(string(tt.charSet), string(c)) {
					t.Errorf("Expected %s to be in custom char set for %s", string(c), tt.name)
				}
			}
		})
	}
}

func TestStringBuilderConstructorsWithCharSet(t *testing.T) {
	tests := []struct {
		name     string
		function func(int, CharacterSet) *StringBuilder
		length   int
		charSet  CharacterSet
	}{
		{"TestAlpha", NewRandom, 10, Alpha},
		{"TestAlphaZero", NewRandom, 0, Alpha},
		{"TestAlphaNeg", NewRandom, -1, Alpha},
		{"TestAlphaNumeric", NewRandom, 10, AlphaNumericChars},
		{"TestAlphaNumericZero", NewRandom, 0, AlphaNumericChars},
		{"TestAlphaNumericNeg", NewRandom, -1, AlphaNumericChars},
		{"TestHex", NewRandom, 10, HexChars},
		{"TestHexZero", NewRandom, 0, HexChars},
		{"TestHexNeg", NewRandom, -1, HexChars},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(tt.length, tt.charSet)
			if result == nil {
				t.Errorf("Expected non-nil result for %s", tt.name)
			}
			if result.String() == "" && tt.length > 0 {
				t.Errorf("Expected non-empty string for %s", tt.name)
			}
			if result.String() != "" && tt.length == 0 {
				t.Errorf("Expected empty string for %s", tt.name)
			}
			if result.String() != "" && tt.length < 0 {
				t.Errorf("Expected empty string for %s", tt.name)
			}
			for _, c := range result.String() {
				if !strings.Contains(string(tt.charSet), string(c)) {
					t.Errorf("Expected %s to be in custom char set for %s", string(c), tt.name)
				}
			}
		})
	}
}

func TestStringBuilderConstructorCustCharSet(t *testing.T) {
	tests := []struct {
		name    string
		charset string
		length  int
	}{
		{"TestCustomChars1", "abc", 10},
		{"TestCustomChars2", "abc", 0},
		{"TestCustomChars3", "abc", -1},
		{"TestCustomChars4", "golang", 10},
		{"TestCustomChars5", "a", 100},
		{"TestCustomChars6", "abracadabra", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewRandomFromCustomCharSet(tt.length, tt.charset)
			if result == nil {
				t.Errorf("Expected non-nil result for %s", tt.name)
			}
			if result.String() == "" && tt.length > 0 {
				t.Errorf("Expected non-empty string for %s", tt.name)
			}
			if result.String() != "" && tt.length == 0 {
				t.Errorf("Expected empty string for %s", tt.name)
			}
			for _, c := range result.String() {
				if !strings.Contains(tt.charset, string(c)) {
					t.Errorf("Expected %s to be in custom char set for %s", string(c), tt.name)
				}
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
			helperResult := isValidUUID(tt.input)
			result := IsUUID(tt.input)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("IsUUID(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			builderErr := New(tt.input).RequireUUID().Error()
			if isValidUUID(tt.input) && builderErr != nil {
				t.Errorf("IsUUID(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			if builderErr != nil && !errors.Is(builderErr, errors2.ErrInvalidUUID) {
				t.Errorf("IsUUID(%q) = %v; want %v", tt.input, builderErr.Error(), errors2.ErrInvalidUUID)
			}
		})
	}
}

func TestStringBuilderPrint(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "hello world", "hello world"},
		{"Test 2", "hello world",
			"error"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Test 1" {
				if New(tt.input).formatOutput() != tt.expected {
					t.Errorf("Expected %q, got %q for %s", tt.expected, New(tt.input).formatOutput(), tt.name)
				}
				New(tt.input).Print()
			}
			if tt.name == "Test 2" {
				if New(tt.input).
					WithComparisonManager().
					setError(errors.New("error"), true).
					formatOutput() != tt.expected {
					t.Errorf("Expected %q, got %q for %s", tt.expected,
						New(tt.input).WithComparisonManager().HammingDistance("ERROR").formatOutput(), tt.name)
				}
				New(tt.input).WithComparisonManager().HammingDistance("ERROR").Print()
			}
		})
	}
}

func TestStringBuilderResult(t *testing.T) {
	tests := []struct {
		name     string
		input    *StringBuilder
		expected string
		err      bool
		compMan  bool
	}{
		{"BuilderResult1", New("Hello World"), "Hello World", false, false},
		{"BuilderResult2",
			New("Hello World").
				WithComparisonManager(),
			"Hello World",
			false,
			true},
		{"BuilderResult3",
			New("Hello World").
				WithComparisonManager().HammingDistance("oops"),
			"Hello World",
			true,
			true},
		{"BuilderResult4",
			New("Hello World").
				WithComparisonManager().LevenshteinDistance("Hello World!"),
			"Hello World",
			false,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex, err := tt.input.Result()
			if ex != tt.expected {
				t.Errorf("Expected %q, got %q for %s", tt.expected, ex, tt.name)
			}
			if tt.err == true && err == nil {
				t.Errorf("Expected error for %s", tt.name)
			}
		})
	}
}

func TestStringBuilderBuild(t *testing.T) {
	tests := []struct {
		name     string
		input    *StringBuilder
		expected string
		err      bool
	}{
		{"BuilderBuild1", New("Hello World"), "Hello World", false},
		{"BuilderBuild2",
			New("Hello World").
				WithComparisonManager(),
			"Hello World",
			false},
		{"BuilderBuild3",
			New("Hello World").
				WithComparisonManager().HammingDistance("oops"),
			"",
			true},
		{"BuilderBuild4",
			New("Hello World").
				WithComparisonManager().HammingDistance("Hello World!"),
			"",
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex, err := tt.input.Build()
			if ex != tt.expected {
				t.Errorf("Expected %q, got %q for %s", tt.expected, ex, tt.name)
			}
			if tt.err == true && err == nil {
				t.Errorf("Expected error for %s", tt.name)
			}
		})
	}
}
