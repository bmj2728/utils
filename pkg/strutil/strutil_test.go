package strutil

import (
	"github.com/google/uuid"
	"strings"
	"testing"
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
		t.Errorf("Expected NewUUID() to generate unique values, but got identical UUIDs: %s and %s", uuid5.String(), uuid6.String())
	}

	if uuid7.String() == uuid8.String() {
		t.Errorf("Expected NewUUIDV7 to generate unique values, but got identical UUIDs: %s and %s", uuid7.String(), uuid8.String())
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
		{"RandomStringNegLen", RandomString, -1},
		{"RandomHex", RandomHex, 16},
		{"RandomHexZeroLen", RandomHex, 0},
		{"RandomHexNegLen", RandomHex, -1},
		{"RandomUrlSafe", RandomUrlSafe, 8},
		{"RandomUrlSafeZeroLen", RandomUrlSafe, 0},
		{"RandomUrlSafeNegLen", RandomUrlSafe, -1},
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
	}{
		{"NewRandom", NewRandom, 10},
		{"NewRandomZero", NewRandom, 0},
		{"NewRandomNeg", NewRandom, -1},
		{"NewRandomHex", NewRandomHex, 6},
		{"NewRandomHexZero", NewRandomHex, 0},
		{"NewRandomHexNeg", NewRandomHex, -1},
		{"NewRandomURLSafe", NewRandomURLSafe, 8},
		{"NewRandomURLSafeZero", NewRandomURLSafe, 0},
		{"NewRandomURLSafeNeg", NewRandomURLSafe, -1},
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
			helperResult := isValidURL(tt.input)
			result := IsURL(tt.input)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("IsURL(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			builderErr := New(tt.input).RequireURL().Error()
			if isValidURL(tt.input) && builderErr != nil {
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

func TestIsAlphaString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"AlphaStringValid", "hello", true},
		{"AlphaStringValidUpper", "HELLO", true},
		{"AlphaStringInvalid", "hello world", false},
		{"AlphaStringInvalidSpecial", "hello world!", false},
		{"AlphaStringInvalidNumber", "rogue1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := isAlphaString(tt.input)
			result := IsAlphaString(tt.input)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("IsAlphaString(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			builderResult := New(tt.input).RequireAlpha().Error()
			if isAlphaString(tt.input) && builderResult != nil {
				t.Errorf("IsAlphaString(%q) = %v; want %v", tt.input, result, tt.expected)
			}
			if builderResult != nil && builderResult.Error() != ErrInvalidNotAlpha {
				t.Errorf("IsAlphaString(%q) = %v; want %v", tt.input, builderResult.Error(), ErrInvalidNotAlpha)
			}
		})
	}
}

func TestKeepAlphaNumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		ws       bool
		expected string
	}{
		{"KeepAlphaNumericRemoveSpace", "hello world 123", false, "helloworld123"},
		{"KeepAlphaNumericPreserveSpace", "hello world 123", true, "hello world 123"},
		{"KeepAlphaNumericNoSpaceSpecialsValid", "hello! world! 123!@#$%^&*()", false, "helloworld123"},
		{"KeepAlphaNumericPreserveSpaceSpecialsValid", "hello! world! 123!@#$%^&*()", true, "hello world 123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := alphaNumeric(tt.input, tt.ws)
			result := KeepAlphaNumeric(tt.input, tt.ws)
			builderResult := New(tt.input).KeepAlphaNumeric(tt.ws).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("KeepAlphaNumeric(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestKeepAlpha(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		ws       bool
		expected string
	}{
		{"KeepAlphaRemoveSpace", "hello world", false, "helloworld"},
		{"KeepAlphaPreserveSpace", "hello world", true, "hello world"},
		{"KeepAlphaRemoveSpaceDigits", "hello world 123", false, "helloworld"},
		{"KeepAlphaPreserveSpaceDigits", "hello world 123", true, "hello world "},
		{"KeepAlphaNoSpaceSpecialsValid", "hello! world! 123!@#$%^&*()", false, "helloworld"},
		{"KeepAlphaPreserveSpaceSpecialsValid", "hello! world! 123!@#$%^&*()", true, "hello world "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := alpha(tt.input, tt.ws)
			result := KeepAlpha(tt.input, tt.ws)
			builderResult := New(tt.input).KeepAlpha(tt.ws).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("KeepAlpha(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestLoremWord(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremWord1"},
		{"LoremWord2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremWord()
			result := LoremWord()
			builderResult := NewLoremWord().String()
			if helperResult == "" || builderResult == "" || result == "" {
				t.Errorf("LoremWord - %q, %q, %q", result, helperResult, builderResult)
			}

		})
	}
}

func TestLoremWords(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{"LoremWords1", 1},
		{"LoremWords30", 30},
		{"LoremWords100", 100},
		{"LoremWordsZero", 0},
		{"LoremWordsNegative", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremWords(tt.count)
			result := LoremWords(tt.count)
			builderResult := NewLoremWords(tt.count).String()

			if tt.count < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremWords - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.count > 0 && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremWords - %q, %q, %q", result, helperResult, builderResult)
			}

			helperResultCount := len(strings.Split(helperResult, " ")) - 1
			helperArr := strings.Split(helperResult, " ")
			resultCount := len(strings.Split(result, " ")) - 1
			resultArr := strings.Split(result, " ")
			builderResultCount := len(strings.Split(builderResult, " ")) - 1
			builderArr := strings.Split(builderResult, " ")
			if tt.count > 0 && (helperResultCount != tt.count || resultCount != tt.count || builderResultCount != tt.count) {
				t.Errorf("LoremWords(%d, %d, %d) - expected %d - words: %q / %q / %q", resultCount, helperResultCount, builderResultCount, tt.count, resultArr, helperArr, builderArr)
			}
		})
	}
}

func TestLoremSentence(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremSentence1"},
		{"LoremSentence2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremSentence()
			result := LoremSentence()
			builderResult := NewLoremSentence().String()
			if helperResult == "" || builderResult == "" || result == "" {
				t.Errorf("LoremSentence - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if helperResultCount != 8 || builderResultCount != 8 || resultCount != 8 {
				t.Errorf("LoremSentence - expected 8 - words: %d / %d / %d", resultCount, helperResultCount, builderResultCount)
			}

			if helperResult[0] < 'A' || helperResult[0] > 'Z' || result[0] < 'A' || result[0] > 'Z' || builderResult[0] < 'A' || builderResult[0] > 'Z' {
				t.Errorf("LoremSentence - first character not uppercase: %q / %q / %q", helperResult[0], result[0], builderResult[0])
			}

			if helperResult[len(helperResult)-1] != '.' || result[len(result)-1] != '.' || builderResult[len(builderResult)-1] != '.' {
				t.Errorf("LoremSentence - last character not period: %q / %q / %q", helperResult[len(helperResult)-1], result[len(result)-1], builderResult[len(builderResult)-1])
			}
		})
	}
}

func TestLoremSentenceCustom(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"LoremSentenceCustom1", 1},
		{"LoremSentenceCustom11", 11},
		{"LoremSentenceCustom25", 25},
		{"LoremSentenceCustomZero", 0},
		{"LoremSentenceCustomNeg", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremSentenceCustom(tt.length)
			result := LoremSentenceCustom(tt.length)
			builderResult := NewLoremSentenceCustom(tt.length).String()
			if tt.length > 0 && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremSentenceCustom - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.length < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremSentenceCustom - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if tt.length > 0 && (helperResultCount != tt.length || builderResultCount != tt.length || resultCount != tt.length) {
				t.Errorf("LoremSentenceCustom - expected %d - words: %d / %d / %d", tt.length, resultCount, helperResultCount, builderResultCount)
			}

			if tt.length > 0 && (helperResult[0] < 'A' || helperResult[0] > 'Z' || result[0] < 'A' || result[0] > 'Z' || builderResult[0] < 'A' || builderResult[0] > 'Z') {
				t.Errorf("LoremSentenceCustom - first character not uppercase: %q / %q / %q", helperResult[0], result[0], builderResult[0])
			}

			if tt.length > 0 && (helperResult[len(helperResult)-1] != '.' || result[len(result)-1] != '.' || builderResult[len(builderResult)-1] != '.') {
				t.Errorf("LoremSentenceCustom - last character not period: %q / %q / %q", helperResult[len(helperResult)-1], result[len(result)-1], builderResult[len(builderResult)-1])
			}
		})
	}
}

func TestLoremSentences(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		expected int
	}{
		{"LoremSentences1", 1, 8},
		{"LoremSentences3", 3, 24},
		{"LoremSentences10", 10, 80},
		{"LoremSentencesZero", 0, 0},
		{"LoremSentencesNegative", -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremSentences(tt.count)
			result := LoremSentences(tt.count)
			builderResult := NewLoremSentences(tt.count).String()
			if tt.count < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremSentencesNotNil - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.count > 0 && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremSentencesIsNil - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if tt.count > 0 && (helperResultCount != tt.expected || resultCount != tt.expected || builderResultCount != tt.expected) {
				t.Errorf("LoremSentencesLen1 - expected %d - sentences: %d / %d / %d", tt.expected, helperResultCount, resultCount, builderResultCount)
			}
			if tt.count < 1 && (helperResultCount != tt.expected+1 || resultCount != tt.expected+1 || builderResultCount != tt.expected+1) {
				t.Errorf("LoremSentenceslen2 - expected %d - sentences: %d / %d / %d", tt.expected, helperResultCount, resultCount, builderResultCount)
			}
		})
	}
}

func TestLoremSentencesCustom(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		length   int
		expected int
	}{
		{"LoremSentencesCustom1", 1, 1, 1},
		{"LoremSentencesCustom11", 3, 7, 21},
		{"LoremSentencesCustom25", 5, 12, 60},
		{"LoremSentencesCustomZero", 0, 0, 0},
		{"LoremSentencesCustomNeg", -1, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremSentencesCustom(tt.count, tt.length)
			result := LoremSentencesCustom(tt.count, tt.length)
			builderResult := NewLoremSentencesCustom(tt.count, tt.length).String()
			if tt.count < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremSentencesCustomNotNil - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.count > 0 && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremSentencesCustomIsNil - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if tt.count > 0 && (helperResultCount != tt.expected || resultCount != tt.expected || builderResultCount != tt.expected) {
				t.Errorf("LoremSentencesCustomLen1 - expected %d - sentences: %d / %d / %d", tt.expected, helperResultCount, resultCount, builderResultCount)
			}
			if tt.count < 1 && (helperResultCount != tt.expected+1 || resultCount != tt.expected+1 || builderResultCount != tt.expected+1) {
				t.Errorf("LoremSentencesCustomLen2 - expected %d - sentences: %d / %d / %d", tt.expected, helperResultCount, resultCount, builderResultCount)
			}
		})
	}
}

func TestLoremSentencesVariable(t *testing.T) {
	tests := []struct {
		name  string
		count int
		min   int
		max   int
	}{
		{"LoremSentencesVariable1", 1, 1, 10},
		{"LoremSentencesVariable3", 3, 1, 10},
		{"LoremSentencesVariable10", 5, 6, 14},
		{"LoremSentencesVariableZero", 0, 0, 0},
		{"LoremSentencesVariableNeg", -1, 0, 0},
		{"LoremSentencesVariableInvalid", 3, 10, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremSentencesVariable(tt.count, tt.min, tt.max)
			result := LoremSentencesVariable(tt.count, tt.min, tt.max)
			builderResult := NewLoremSentencesVariable(tt.count, tt.min, tt.max).String()
			if tt.count < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremSentencesVariableNotNil - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.count > 0 && tt.min <= tt.max && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremSentencesVariableIsNil - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			minRange := tt.min * tt.count
			maxRange := tt.max * tt.count
			if tt.count > 0 && tt.min <= tt.max && (helperResultCount < minRange || helperResultCount > maxRange || resultCount < minRange || resultCount > maxRange || builderResultCount < minRange || builderResultCount > maxRange) {
				t.Errorf("LoremSentencesVariableLen1 - expected %d - %d - %d - %d - sentences: %d / %d / %d / %d", tt.count, tt.min, tt.max, minRange, helperResultCount, resultCount, builderResultCount, maxRange)
			}
			if tt.max < tt.min && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremSentencesVariableLen2 - expected %d - %d - %d - %d - sentences: %d / %d / %d / %d", tt.count, tt.min, tt.max, minRange, helperResultCount, resultCount, builderResultCount, maxRange)
			}
		})
	}
}

func TestLoremParagraph(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremParagraph1"},
		{"LoremParagraph2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremParagraph()
			result := LoremParagraph()
			builderResult := NewLoremParagraph().String()
			if helperResult == "" || builderResult == "" || result == "" {
				t.Errorf("LoremParagraph - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if helperResultCount != 45 || builderResultCount != 45 || resultCount != 45 {
				t.Errorf("LoremParagraph - expected 45 - words: %d / %d / %d", resultCount, helperResultCount, builderResultCount)
			}
		})
	}
}

func TestLoremParagraphs(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		expected int
	}{
		{"LoremParagraphs1", 1, 45},
		{"LoremParagraphs3", 3, 133},
		{"LoremParagraphs10", 10, 441},
		{"LoremParagraphs1000", 1000, 44001},
		{"LoremParagraphsZero", 0, 0},
		{"LoremParagraphsNegative", -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremParagraphs(tt.count)
			result := LoremParagraphs(tt.count)
			builderResult := NewLoremParagraphs(tt.count).String()
			if tt.count < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremParagraphsNotNil - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.count > 0 && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremParagraphsIsNil - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if tt.count > 0 && (helperResultCount != tt.expected || resultCount != tt.expected || builderResultCount != tt.expected) {
				t.Errorf("LoremParagraphsLen1 - expected %d - sentences: %d / %d / %d", tt.expected, helperResultCount, resultCount, builderResultCount)
			}
			if tt.count < 1 && (helperResultCount != tt.expected+1 || resultCount != tt.expected+1 || builderResultCount != tt.expected+1) {
				t.Errorf("LoremParagraphsLen2 - expected %d - sentences: %d / %d / %d", tt.expected, helperResultCount, resultCount, builderResultCount)
			}
		})
	}
}

func TestIsDomain(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"IsDomainValidBasic", "example.com", true},
		{"IsDomainValidMultiLevel", "example.com.au", true},
		{"IsDomainValidSubDomain", "www.example.com", true},
		{"IsDomainValidSubDomainMultiLevel", "www.example.com.au", true},
		{"IsDomainValidEmpty", "", false},
		{"IsDomainValidInvalidNoTLD", "example", false},
		{"IsDomainValidInvalidTrailingPeriod", "example.com.", false},
		{"IsDomainValidInvalidLeadingPeriod", ".example.com", false},
		{"IsDomainValidTooLong", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsDomain(tt.input)
			if result != tt.expected {
				t.Errorf("IsDomain - expected %t - got %t", tt.expected, result)
			}
		})
	}
}

func TestLoremDomain(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremDomain1"},
		{"LoremDomain2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremDomain()
			result := LoremDomain()
			builderResult := NewLoremDomain().String()
			if helperResult == "" || builderResult == "" || result == "" {
				t.Errorf("LoremDomain - %q, %q, %q", result, helperResult, builderResult)
			}
			if !isValidDomain(result) || !isValidDomain(builderResult) || !isValidDomain(helperResult) {
				t.Errorf("LoremDomain - invalid domain: %q / %q / %q", result, builderResult, helperResult)
			}
		})
	}
}

func TestLoremURL(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremURL1"},
		{"LoremURL2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremURL()
			result := LoremURL()
			builderResult := NewLoremURL().String()
			if result == "" || builderResult == "" || helperResult == "" {
				t.Errorf("LoremURL - %q, %q", result, builderResult)
			}
			if !isValidURL(result) || !isValidURL(builderResult) || !isValidURL(helperResult) {
				t.Errorf("LoremURL - invalid url: %q / %q", result, builderResult)
			}

		})
	}
}

func TestLoremEmail(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremEmail1"},
		{"LoremEmail2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremEmail()
			result := LoremEmail()
			builderResult := NewLoremEmail().String()
			if result == "" || builderResult == "" || helperResult == "" {
				t.Errorf("LoremEmail - %q, %q", result, builderResult)
			}
			if !isValidEmail(result) || !isValidEmail(builderResult) || !isValidEmail(helperResult) {
				t.Errorf("LoremEmail - invalid email: %q / %q", result, builderResult)
			}
		})
	}
}

func TestReplaceWhitespace(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		replacement string
		expected    string
	}{
		{"ReplaceWhiteSpaceDash", "a b\nc\rd\fe\vf\tg", "-", "a-b-c-d-e-f-g"},
		{"ReplaceWhiteSpaceSpace", "a b\nc\rd\fe\vf\tg", " ", "a b c d e f g"},
		{"ReplaceWhiteSpaceSlash", "a b\nc\rd\fe\vf\tg", "/", "a/b/c/d/e/f/g"},
		{"ReplaceWhiteSpaceText", "a b\nc\rd\fe\vf\tg", ", and ", "a, and b, and c, and d, and e, and f, and g"},
		{"ReplaceWhiteSpaceEmpty", "", " ", ""},
		{"ReplaceWhiteSpaceEmptyReplacement", "", " ", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := replaceWhitespace(tt.input, tt.replacement)
			result := ReplaceWhitespace(tt.input, tt.replacement)
			builderResult := New(tt.input).ReplaceWhitespace(tt.replacement).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ReplaceWhitespace - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestReplaceSpaces(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		replacement string
		expected    string
	}{
		{"ReplaceSpacesDash", "a b\nc\rd\fe\vf\tg", "-", "a-b\nc\rd\fe\vf\tg"},
		{"ReplaceSpacesSpace", "a b\nc\rd\fe\vf\tg", " ", "a b\nc\rd\fe\vf\tg"},
		{"ReplaceSpacesSlash", "a b\nc\rd\fe\vf\tg", "/", "a/b\nc\rd\fe\vf\tg"},
		{"ReplaceSpacesText", "a b\nc\rd\fe\vf\tg", ", and ", "a, and b\nc\rd\fe\vf\tg"},
		{"ReplaceSpacesEmpty", "", " ", ""},
		{"ReplaceSpacesEmptyReplacement", "", " ", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := replaceSpaces(tt.input, tt.replacement)
			result := ReplaceSpaces(tt.input, tt.replacement)
			builderResult := New(tt.input).ReplaceSpaces(tt.replacement).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ReplaceSpaces - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}
