package strutil

import (
	"errors"
	"testing"

	errors2 "github.com/bmj2728/utils/pkg/internal/errors"
)

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
			if builderErr != nil && !errors.Is(builderErr, errors2.ErrInvalidEmail) {
				t.Errorf("IsEmail(%q) = %v; want %v", tt.input, builderErr.Error(), errors2.ErrInvalidEmail)
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
		{"ValidUrlWithPathAndQuery",
			"https://www.google.com/search?q=hello",
			true},
		{"ValidUrlWithPathAndQueryAndFragment",
			"https://www.google.com/search?q=hello#fragment",
			true},
		{"ValidUrlWithPathAndQueryAndFragmentAndParams",
			"https://www.google.com/search?q=hello#fragment&param1=value1&param2=value2",
			true},
		{"ValidUrlWithPathAndQueryAndFragmentAndParamsAndTrailingSlash",
			"https://www.google.com/search?q=hello#fragment&param1=value1&param2=value2/",
			true},
		{"ValidURLWithPort",
			"https://www.google.com:443",
			true},
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
			if builderErr != nil && !errors.Is(builderErr, errors2.ErrInvalidURL) {
				t.Errorf("IsURL(%q) = %v; want %v", tt.input, builderErr.Error(), errors2.ErrInvalidURL)
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
			if (tt.min < 0 || tt.max < 0) && !errors.Is(builderErr, errors2.ErrInvalidLengthRange) {
				t.Errorf("IsValidLength(%q, %d, %d) = %v; want %v",
					tt.input, tt.min, tt.max, builderErr.Error(), errors2.ErrInvalidLengthRange)
			}

			if tt.min > tt.max && !errors.Is(builderErr, errors2.ErrInvalidLengthRange) {
				t.Errorf("IsValidLength(%q, %d, %d) = %v; want %v",
					tt.input, tt.min, tt.max, builderErr.Error(), errors2.ErrInvalidLengthRange)
			}

			if tt.min > 0 &&
				tt.max > 0 &&
				tt.min <= tt.max &&
				!isLengthInRange(tt.input, tt.min, tt.max) &&
				!errors.Is(builderErr, errors2.ErrInvalidLength) {
				t.Errorf("IsValidLength(%q, %d, %d) = %v; want %v",
					tt.input, tt.min, tt.max, builderErr.Error(), errors2.ErrInvalidLength)
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
			if isEmpty(tt.input) && !errors.Is(builderErr, errors2.ErrInvalidEmpty) {
				t.Errorf("IsEmpty(%q) = %v; want %v", tt.input, builderErr.Error(), errors2.ErrInvalidEmpty)
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
			if isEmptyNormalized(tt.input) && !errors.Is(builderErr, errors2.ErrInvalidEmptyAfterNormalization) {
				t.Errorf("IsEmptyNormalized(%q) = %v; want %v",
					tt.input, builderErr.Error(), errors2.ErrInvalidEmptyAfterNormalization)
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
			if builderResult != nil && !errors.Is(builderResult, errors2.ErrInvalidNotAlphaNumeric) {
				t.Errorf("IsAlphaNumericString(%q) = %v; want %v", tt.input, builderResult, errors2.ErrInvalidNotAlphaNumeric)
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
			if builderResult != nil && !errors.Is(builderResult, errors2.ErrInvalidNotAlpha) {
				t.Errorf("IsAlphaString(%q) = %v; want %v", tt.input, builderResult.Error(), errors2.ErrInvalidNotAlpha)
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
		{"IsDomainValidTooLong",
			"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz.com",
			false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := isValidDomain(tt.input)
			result := IsDomain(tt.input)
			builderErr := New(tt.input).RequireDomain().Error()
			ok := builderErr == nil
			if result != tt.expected || helperResult != tt.expected || ok != tt.expected {
				t.Errorf("IsDomain - expected %t - got %t", tt.expected, result)
			}
		})
	}
}

func TestIsNormalizedUnicode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		format   NormalizationFormat
		expected bool
	}{
		{"NFC_AlreadyComposed", "café", NFC, true},
		{"NFD_AlreadyDecomposed", "cafe\u0301", NFD, true},
		{"NFKC_FractionHalf", "½", NFKC, false},
		{"NFKC_WithAccents", "ﬁlé", NFKC, false},
		{"NFKD_FractionHalf", "½", NFKD, false},
		{"NFKD_SuperscriptTwo", "x²", NFKD, false},
		{"NFKD_Ligature", "ﬁle", NFKD, false},
		{"NFKD_CircledOne", "①", NFKD, false},
		{"NFD_Empty", "", NFD, true},
		{"NFC_Empty", "", NFC, true},
		{"NFKC_Empty", "", NFKC, true},
		{"NFKD_Empty", "", NFKD, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := isNormalizedUnicode(tt.input, tt.format)
			result := IsNormalizedUnicode(tt.input, tt.format)
			if result != tt.expected || helperResult != tt.expected {
				t.Errorf("IsNormalizedUnicode(%q, %v) = %v; want %v", tt.input, tt.format, result, tt.expected)
			}
			builderErr := New(tt.input).RequireNormalizedUnicode(tt.format)
			if builderErr.Error() != nil && tt.expected == true {
				t.Errorf("IsNormalizedUnicode(%q, %v) = %v; want %v",
					tt.input, tt.format, builderErr.Error(), tt.expected)
			}
			if tt.expected == false && (builderErr.Error() == nil ||
				!errors.Is(builderErr.Error(), errors2.ErrNotNormalizedUnicode)) {
				t.Errorf("IsNormalizedUnicode(%q, %v) = %v; want %v",
					tt.input, tt.format, builderErr.Error(), tt.expected)
			}

		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		substr   string
		expected bool
	}{
		{"Contains", "hello world", "hello", true},
		{"ContainsUpper", "HELLO WORLD", "hello", false},
		{"ContainsInvalid", "hello world!", "John", false},
		{"ContainsEmpty", "", "hello", false},
		{"ContainsEmptySubstr", "hello", "", false},
		{"ContainsEmptyBoth", "", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := contains(tt.input, tt.substr)
			result := Contains(tt.input, tt.substr)
			builderResult := New(tt.input).RequireContains(tt.substr).Error() == nil
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected {
				t.Errorf("Contains(%q, %q) = %t/%t/%t; want %t",
					tt.input, tt.substr, helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		substr   string
		expected bool
	}{
		{"Contains", "hello world", "HELLO", true},
		{"ContainsUpper", "HELLO WORLD", "hello", true},
		{"ContainsInvalid", "hello world!", "John", false},
		{"ContainsEmpty", "", "hello", false},
		{"ContainsEmptySubstr", "hello", "", false},
		{"ContainsEmptyBoth", "", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := containsIgnoreCase(tt.input, tt.substr)
			result := ContainsIgnoreCase(tt.input, tt.substr)
			builderResult := New(tt.input).RequireContainsIgnoreCase(tt.substr).Error() == nil
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected {
				t.Errorf("ContainsIgnoreCase(%q, %q) = %t/%t/%t; want %t",
					tt.input, tt.substr, helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		substrs  []string
		expected bool
	}{
		{"Contains", "hello world", []string{"hello", "world"}, true},
		{"ContainsUpper", "HELLO WORLD", []string{"hello", "world"}, false},
		{"ContainsInvalid", "hello world!", []string{"hi", "John"}, false},
		{"ContainsOneInvalid", "hello world!", []string{"hello", "John"}, true},
		{"ContainsEmpty", "", []string{"hello", "world"}, false},
		{"ContainsEmptySubstr", "hello", []string{}, false},
		{"ContainsEmptyBoth", "", []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := containsAny(tt.input, tt.substrs)
			result := ContainsAny(tt.input, tt.substrs)
			builderResult := New(tt.input).RequireContainsAny(tt.substrs).Error() == nil
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected {
				t.Errorf("ContainsAny(%q, %q) = %t/%t/%t; want %t",
					tt.input, tt.substrs, helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestContainsAnyIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		substrs  []string
		expected bool
	}{
		{"Contains", "hello world", []string{"HELLO", "WORLD"}, true},
		{"ContainsUpper", "HELLO WORLD", []string{"hello", "world"}, true},
		{"ContainsInvalid", "hello world!", []string{"hi", "John"}, false},
		{"ContainsOneInvalid", "hello world!", []string{"hello", "John"}, true},
		{"ContainsEmpty", "", []string{"hello", "world"}, false},
		{"ContainsEmptySubstr", "hello", []string{}, false},
		{"ContainsEmptyBoth", "", []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := containsAnyIgnoreCase(tt.input, tt.substrs)
			result := ContainsAnyIgnoreCase(tt.input, tt.substrs)
			builderResult := New(tt.input).RequireContainsAnyIgnoreCase(tt.substrs).Error() == nil
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected {
				t.Errorf("ContainsAnyIgnoreCase(%q, %q) = %t/%t/%t; want %t",
					tt.input, tt.substrs, helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		substrs  []string
		expected bool
	}{
		{"Contains", "hello world", []string{"hello", "world"}, true},
		{"ContainsUpper", "HELLO WORLD", []string{"hello", "world"}, false},
		{"ContainsInvalid", "hello world!", []string{"hi", "John"}, false},
		{"ContainsOneInvalid", "hello world!", []string{"hello", "John"}, false},
		{"ContainsEmpty", "", []string{"hello", "world"}, false},
		{"ContainsEmptySubstr", "hello", []string{}, false},
		{"ContainsEmptyBoth", "", []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := containsAll(tt.input, tt.substrs)
			result := ContainsAll(tt.input, tt.substrs)
			builderResult := New(tt.input).RequireContainsAll(tt.substrs).Error() == nil
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected {
				t.Errorf("ContainsAll(%q, %q) = %t/%t/%t; want %t",
					tt.input, tt.substrs, helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestContainsAllIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		substrs  []string
		expected bool
	}{
		{"Contains", "hello world", []string{"hello", "WORLD"}, true},
		{"ContainsUpper", "HELLO WORLD", []string{"hello", "world"}, true},
		{"ContainsInvalid", "hello world!", []string{"hi", "John"}, false},
		{"ContainsOneInvalid", "hello world!", []string{"hello", "John"}, false},
		{"ContainsEmpty", "", []string{"hello", "world"}, false},
		{"ContainsEmptySubstr", "hello", []string{}, false},
		{"ContainsEmptyBoth", "", []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := containsAllIgnoreCase(tt.input, tt.substrs)
			result := ContainsAllIgnoreCase(tt.input, tt.substrs)
			builderResult := New(tt.input).RequireContainsAllIgnoreCase(tt.substrs).Error() == nil
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected {
				t.Errorf("ContainsAllIgnoreCase(%q, %q) = %t/%t/%t; want %t",
					tt.input, tt.substrs, helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		prefix   string
		expected bool
	}{
		{"HasPrefix", "Prehistoric", "Pre", true},
		{"HasPrefixUpper", "APP_ENV_VAR", "APP", true},
		{"HasPrefixInvalid", "hello world!", "hi", false},
		{"HasPrefixEmpty", "", "hello", false},
		{"HasPrefixEmptyPrefix", "hello", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := hasPrefix(tt.input, tt.prefix)
			result := HasPrefix(tt.input, tt.prefix)
			builderResult := New(tt.input).RequireHasPrefix(tt.prefix).Error() == nil
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected {
				t.Errorf("HasPrefix(%q, %q) = %t/%t/%t; want %t",
					tt.input, tt.prefix, helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestHasSuffix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		suffix   string
		expected bool
	}{
		{"HasPrefix", "Running", "ing", true},
		{"HasPrefixUpper", "APP_ENV_VAR_DEV", "DEV", true},
		{"HasPrefixInvalid", "hello world!", "hi", false},
		{"HasPrefixEmpty", "", "hello", false},
		{"HasPrefixEmptyPrefix", "hello", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := hasSuffix(tt.input, tt.suffix)
			result := HasSuffix(tt.input, tt.suffix)
			builderResult := New(tt.input).RequireHasSuffix(tt.suffix).Error() == nil
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected {
				t.Errorf("HasPrefix(%q, %q) = %t/%t/%t; want %t",
					tt.input, tt.suffix, helperResult, result, builderResult, tt.expected)
			}
		})
	}
}
