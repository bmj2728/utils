package strutil

import (
	"github.com/google/uuid"
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

	e1 := uuid.Validate(uuid1)
	e2 := uuid.Validate(uuid2)
	e3 := uuid.Validate(uuid3)
	e4 := uuid.Validate(uuid4)
	if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
		t.Errorf("Generated UUIDs are not valid UUIDs")
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
			helperResult := isValidUUID(tt.input)
			result := IsUUID(tt.input)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("IsUUID(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			builderErr := New(tt.input).RequireUUID().Error()
			if isValidUUID(tt.input) && builderErr != nil {
				t.Errorf("IsUUID(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			if builderErr != nil && builderErr.Error() != ErrInvalidUUID {
				t.Errorf("IsUUID(%q) = %v; want %v", tt.input, builderErr.Error(), ErrInvalidUUID)
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
			builderResult := New(tt.input).CleanWhitespace().String()
			builderError := New(tt.input).CleanWhitespace().Error()
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected || builderError != nil {
				t.Errorf("cleanWhitespace(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestNormalizeWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"NoWhitespace", "hello world", "hello world"},
		{"ExtraWhitespace", "hello     world", "hello world"},
		{"Tabs", "hello\t\tworld", "hello world"},
		{"Newlines", "hello\n\nworld", "hello world"},
		{"CarriageReturns", "hello\r\rworld", "hello world"},
		{"MixedWhitespace", "hello  \t\tworld\n\n\r\r", "hello world"},
		{"Empty", "", ""},
		{"LeadingWhitespace", "   hello world", "hello world"},
		{"TrailingWhitespace", "hello world   ", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := normalizeWhitespace(tt.input)
			result := NormalizeWhitespace(tt.input)
			builderResult := New(tt.input).NormalizeWhitespace().String()
			builderError := New(tt.input).NormalizeWhitespace().Error()
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected || builderError != nil {
				t.Errorf("normalizeWhitespace(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCollapseWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"NoWhitespace", "hello world", "hello world"},
		{"ExtraWhitespace", "hello     world", "hello world"},
		{"Tabs", "hello\t\tworld", "hello world"},
		{"Newlines", "hello\n\nworld", "hello world"},
		{"CarriageReturns", "hello\r\rworld", "hello world"},
		{"MixedWhitespace", "hello  \t\tworld\n\n\r\r", "hello world "},
		{"Empty", "", ""},
		{"LeadingWhitespace", "   hello world", " hello world"},
		{"LeadingAndTrailingWhitespace", "   hello world   ", " hello world "},
		{"TrailingWhitespace", "hello world   ", "hello world "},
		{"LeadingTrailingWhitespaceMixed", " hello   world\n\r\t   ", " hello world "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := collapseWhitespace(tt.input)
			result := CollapseWhitespace(tt.input)
			builderResult := New(tt.input).CollapseWhitespace().String()
			builderError := New(tt.input).CollapseWhitespace().Error()
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected || builderError != nil {
				t.Errorf("collapseWhitespace(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"ValidEmail", "somebody@gmail.com", true},
		{"ValidEmailUppercase", "SOMEBODY@GMAIL.COM", true},
		{"ValidEmailNoDomain", "somebody@", false},
		{"ValidEmailNoUsername", "@gmail.com", false},
		{"ValidEmailNoTld", "somebody@gmail", true},
		{"ValidEmailMultiLevelDomain", "somebody@place.uk.co", true},
		{"ValidEmailWithSpecial", "john.doe@gmail.com", true},
		{"ValidEmailWithSpecials", "john.doe!#$%&'*+-/=?^_`{|}~@gmail.com", true},
		{"ValidEmailIllegalSpecial", "john.doe @gmail.com", false},
		{"ValidIllegalLeadingPeriod", ".somebody@gmail.com", false},
		{"ValidEmailIllegalDoublePeriod", "john..doe@gmail.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := isValidEmail(tt.input)
			result := IsEmail(tt.input)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("IsEmail(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			builderErr := New(tt.input).RequireEmail().Error()
			if isValidEmail(tt.input) && builderErr != nil {
				t.Errorf("IsEmail(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			if builderErr != nil && builderErr.Error() != ErrInvalidEmail {
				t.Errorf("IsEmail(%q) = %v; want %v", tt.input, builderErr.Error(), ErrInvalidEmail)
			}
		})
	}
}

func TestIsUrl(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"ValidUrl", "https://www.google.com", true},
		{"ValidUrlUppercase", "HTTPS://WWW.GOOGLE.COM", true},
		{"ValidUrlNoScheme", "www.google.com", false},
		{"ValidUrlNoHost", "https://", false},
		{"ValidUrlNoSchemeOrHost", ":", false},
		{"ValidUrlWithPath", "https://www.google.com/search", true},
		{"ValidUrlWithPathAndQuery", "https://www.google.com/search?q=hello", true},
		{"ValidUrlWithPathAndQueryAndFragment", "https://www.google.com/search?q=hello#fragment", true},
		{"ValidUrlWithPathAndQueryAndFragmentAndParams", "https://www.google.com/search?q=hello#fragment&param1=value1&param2=value2", true},
		{"ValidUrlWithPathAndQueryAndFragmentAndParamsAndTrailingSlash", "https://www.google.com/search?q=hello#fragment&param1=value1&param2=value2/", true},
		{"ValidURLWithPort", "https://www.google.com:443", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := isValidUrl(tt.input)
			result := IsURL(tt.input)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("IsURL(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			builderErr := New(tt.input).RequireURL().Error()
			if isValidUrl(tt.input) && builderErr != nil {
				t.Errorf("IsURL(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			if builderErr != nil && builderErr.Error() != ErrInvalidURL {
				t.Errorf("IsURL(%q) = %v; want %v", tt.input, builderErr.Error(), ErrInvalidURL)
			}
		})
	}
}

func TestIsValidLength(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		min      int
		max      int
		expected bool
	}{
		{"ValidLength", "hello world", 0, 20, true},
		{"ValidLengthMin", "hello world", 5, 20, true},
		{"ValidLengthMax", "hello world", 0, 20, true},
		{"ValidLengthMinMax", "hello world", 5, 20, true},
		{"InvalidLengthMin", "hello world", 20, 30, false},
		{"InvalidLengthMax", "hello world", 0, 5, false},
		{"InvalidLengthMinMax", "hello world", 20, 5, false},
		{"InvalidLengthNegativeMin", "hello world", -1, 20, false},
		{"InvalidLengthNegativeMax", "hello world", 0, -1, false},
		{"InvalidLengthNegativeMinMax", "hello world", -1, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := isLengthInRange(tt.input, tt.min, tt.max)
			result := IsValidLength(tt.input, tt.min, tt.max)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("IsValidLength(%q, %d, %d) = %v; want %v", tt.input, tt.min, tt.max, result, tt.expected)
			}

			builderErr := New(tt.input).RequireLength(tt.min, tt.max).Error()
			if isLengthInRange(tt.input, tt.min, tt.max) && builderErr != nil {
				t.Errorf("IsValidLength(%q, %d, %d) = %v; want %v", tt.input, tt.min, tt.max, result, tt.expected)
			}
			if (tt.min < 0 || tt.max < 0) && builderErr.Error() != ErrInvalidLengthRange {
				t.Errorf("IsValidLength(%q, %d, %d) = %v; want %v", tt.input, tt.min, tt.max, builderErr.Error(), ErrInvalidLengthRange)
			}

			if tt.min > tt.max && builderErr.Error() != ErrInvalidLengthRange {
				t.Errorf("IsValidLength(%q, %d, %d) = %v; want %v", tt.input, tt.min, tt.max, builderErr.Error(), ErrInvalidLengthRange)
			}

			if tt.min > 0 && tt.max > 0 && tt.min <= tt.max && !isLengthInRange(tt.input, tt.min, tt.max) && builderErr.Error() != ErrInvalidLength {
				t.Errorf("IsValidLength(%q, %d, %d) = %v; want %v", tt.input, tt.min, tt.max, builderErr.Error(), ErrInvalidLength)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"EmptyString", "", true},
		{"NonEmptyString", "hello world", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := isEmpty(tt.input)
			result := IsEmpty(tt.input)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("IsEmpty(%q) = %v; want %v", tt.input, result, tt.expected)
			}

			builderErr := New(tt.input).RequireNotEmpty().Error()
			if !isEmpty(tt.input) && builderErr != nil {
				t.Errorf("IsEmpty(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			if isEmpty(tt.input) && builderErr.Error() != ErrInvalidEmpty {
				t.Errorf("IsEmpty(%q) = %v; want %v", tt.input, builderErr.Error(), ErrInvalidEmpty)
			}
		})
	}

}

func TestIsEmptyNormalized(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"EmptyString", "", true},
		{"NonEmptyString", "hello world", false},
		{"EmptyStringWithWhitespace", "   ", true},
		{"NonEmptyStringWithWhitespace", "   hello world   ", false},
		{"EmptyStringWithTabs", "\t\t", true},
		{"NonEmptyStringWithTabs", "\t\thello world\t\t", false},
		{"EmptyStringWithNewlines", "\n\n", true},
		{"NonEmptyStringWithNewlines", "\n\nhello world\n\n", false},
		{"EmptyStringWithCarriageReturns", "\r\r", true},
		{"NonEmptyStringWithCarriageReturns", "\r\rhello world\r\r", false},
		{"EmptyStringWithMixedWhitespace", "\t\t\n\n\r\r", true},
		{"NonEmptyStringWithMixedWhitespace", "\t\t\n\n\r\rhello world\t\t\n\n\r\r", false},
		{"EmptyStringWithMixedWhitespaceAndTabs", "\t\t\n\n\r\r   ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := isEmptyNormalized(tt.input)
			result := IsEmptyNormalized(tt.input)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("IsEmptyNormalized(%q) = %v; want %v", tt.input, result, tt.expected)
			}

			builderErr := New(tt.input).RequireNotEmptyNormalized().Error()

			if !isEmptyNormalized(tt.input) && builderErr != nil {
				t.Errorf("IsEmptyNormalized(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			if isEmptyNormalized(tt.input) && builderErr.Error() != ErrInvalidEmptyAfterNormalization {
				t.Errorf("IsEmptyNormalized(%q) = %v; want %v", tt.input, builderErr.Error(), ErrInvalidEmptyAfterNormalization)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"ToUpperNormal", "hello world", "HELLO WORLD"},
		{"ToUpperEmptyString", "", ""},
		{"ToUpperSingleChar", "a", "A"},
		{"ToUpperSingleCharUpper", "A", "A"},
		{"ToUpperWithSpecials", "hello world!@#$%^&*()_+", "HELLO WORLD!@#$%^&*()_+"},
		{"ToUpperWithWhitespace", "hello world", "HELLO WORLD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := toUpper(tt.input)
			result := ToUpper(tt.input)
			builderResult := New(tt.input).ToUpper().String()
			builderError := New(tt.input).ToUpper().Error()
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected || builderError != nil {
				t.Errorf("ToUpper(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"ToLowerNormal", "Hello World", "hello world"},
		{"ToLowerEmptyString", "", ""},
		{"ToLowerSingleChar", "A", "a"},
		{"ToLowerSingleCharLower", "a", "a"},
		{"ToLowerWithSpecials", "HELLO WORLD!@#$%^&*()_+", "hello world!@#$%^&*()_+"},
		{"ToLowerWithWhitespace", " HELLO WORLD ", " hello world "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := toLower(tt.input)
			result := ToLower(tt.input)
			builderResult := New(tt.input).ToLower().String()
			builderError := New(tt.input).ToLower().Error()
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected || builderError != nil {
				t.Errorf("ToLower(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsAlphaNumericString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"AlphaNumericStringValid", "hello123", true},
		{"AlphaNumericStringValidUpper", "HELLO123", true},
		{"AlphaNumericStringInvalid", "hello world", false},
		{"AlphaNumericStringInvalidSpecial", "hello world!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := isAlphaNumericString(tt.input)
			result := IsAlphaNumericString(tt.input)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("IsAlphaNumericString(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			builderResult := New(tt.input).RequireAlphaNumeric().Error()
			if isAlphaNumericString(tt.input) && builderResult != nil {
				t.Errorf("IsAlphaNumericString(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			if builderResult != nil && builderResult.Error() != ErrInvalidNotAlphaNumeric {
				t.Errorf("IsAlphaNumericString(%q) = %v; want %v", tt.input, builderResult.Error(), ErrInvalidNotAlphaNumeric)
			}
		})
	}
}
