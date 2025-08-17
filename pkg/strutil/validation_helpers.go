package strutil

import (
	"net/mail"
	"net/url"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/text/unicode/norm"
)

// isEmail checks if the provided string s is a valid email address format
// and returns true if valid, false otherwise.
func isEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

// isURL checks if the provided string is a valid URL with a defined scheme and host.
// Returns true if valid, false otherwise.
func isURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}
	u, err := url.Parse(s)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

// isDomain validates whether a given string conforms to a valid domain name format based on specified rules.
func isDomain(domain string) bool {
	if domain == "" {
		return false
	}
	return DomainRegex.MatchString(strings.TrimSpace(domain))
}

// isUUID checks if the provided string is a valid UUID. Returns true if valid, otherwise false.
func isUUID(s string) bool {
	err := uuid.Validate(s)
	return err == nil
}

// isLengthInRange checks if the length of the string s is within the inclusive range specified by min and max values.
func isLengthInRange(s string, min, max int) bool {
	if min < 0 || max < 0 {
		return false
	}
	if min > max {
		return false
	}
	return len(s) >= min && len(s) <= max
}

// isEmpty checks if the provided string is empty and returns true if it is, otherwise false.
func isEmpty(s string) bool {
	return len(s) == 0
}

// isEmptyNormalized checks if the input string is empty after normalizing and trimming whitespace.
func isEmptyNormalized(s string) bool {
	return len(normalizeWhitespace(s, ' ')) == 0
}

// isAlphaNumeric checks whether a given string consists only of alphanumeric characters (letters and digits).
func isAlphaNumeric(s string) bool {
	if isEmpty(s) {
		return false
	}
	for _, c := range s {
		if !IsAlphaNumericRune(c) {
			return false
		}
	}
	return true
}

// isNumeric checks if the input string contains only numeric characters,
// optionally enforcing strict digit-only validation.
func isNumeric(s string, strict bool) bool {
	if isEmpty(s) {
		return false
	}
	for _, c := range s {
		if strict {
			if !unicode.IsDigit(c) {
				return false
			}
		} else {
			if !unicode.IsNumber(c) {
				return false
			}
		}
	}
	return true
}

// isAlpha checks if the given string contains only alphabetic characters.
func isAlpha(s string) bool {
	if isEmpty(s) {
		return false
	}
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

// isNormalizedUnicode checks if a string is normalized according to the specified Unicode normalization format.
func isNormalizedUnicode(s string, format NormalizationFormat) bool {
	return norm.Form(format).IsNormalString(s)
}

// contains determines if the substring `substr` exists within the string `s` and returns true if found,
// otherwise false.
func contains(s string, substr string) bool {
	if isEmpty(s) || isEmpty(substr) {
		return false
	}
	return strings.Contains(s, substr)
}

// containsIgnoreCase checks if substr is present in s, ignoring case sensitivity. Returns true if found,
// false otherwise.
func containsIgnoreCase(s string, substr string) bool {
	if isEmpty(s) || isEmpty(substr) {
		return false
	}
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// containsAny checks if any string in the slice substrs is present in the string s and returns true if found.
func containsAny(s string, substrs []string) bool {
	if isEmpty(s) || len(substrs) == 0 {
		return false
	}
	for _, substr := range substrs {
		if contains(s, substr) {
			return true
		}
	}
	return false
}

// containsAnyIgnoreCase checks if the input string contains any of the given substrings, ignoring case sensitivity.
func containsAnyIgnoreCase(s string, substrs []string) bool {
	if isEmpty(s) || len(substrs) == 0 {
		return false
	}
	for _, substr := range substrs {
		if containsIgnoreCase(s, substr) {
			return true
		}
	}
	return false
}

// containsAll returns true if all strings in substrs are present in s; returns false if s is empty or substrs is empty.
func containsAll(s string, substrs []string) bool {
	if isEmpty(s) || len(substrs) == 0 {
		return false
	}
	for _, substr := range substrs {
		if !contains(s, substr) {
			return false
		}
	}
	return true
}

// containsAllIgnoreCase checks if all substrings in the slice are present in the given string,
// ignoring case sensitivity.
func containsAllIgnoreCase(s string, substrs []string) bool {
	if isEmpty(s) || len(substrs) == 0 {
		return false
	}
	for _, substr := range substrs {
		if !containsIgnoreCase(s, substr) {
			return false
		}
	}
	return true
}

// hasPrefix checks if the string 's' starts with the specified 'prefix'. Returns true if it does, otherwise false.
func hasPrefix(s string, prefix string) bool {
	if isEmpty(s) || isEmpty(prefix) {
		return false
	}
	return strings.HasPrefix(s, prefix)
}

// hasSuffix checks if the string `s` ends with the given `suffix` and returns true if it does, otherwise false.
func hasSuffix(s string, suffix string) bool {
	if isEmpty(s) || isEmpty(suffix) {
		return false
	}
	return strings.HasSuffix(s, suffix)
}
