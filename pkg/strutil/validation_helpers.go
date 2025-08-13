package strutil

import (
	"net/mail"
	"net/url"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/text/unicode/norm"
)

// isValidEmail checks if the provided string s is a valid email address format
// and returns true if valid, false otherwise.
func isValidEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

// isValidURL checks if the provided string is a valid URL with a defined scheme and host.
// Returns true if valid, false otherwise.
func isValidURL(s string) bool {
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

// isValidDomain validates whether a given string conforms to a valid domain name format based on specified rules.
func isValidDomain(domain string) bool {
	if domain == "" {
		return false
	}
	return DomainRegex.MatchString(strings.TrimSpace(domain))
}

// isValidUUID checks if the provided string is a valid UUID. Returns true if valid, otherwise false.
func isValidUUID(s string) bool {
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

// isAlphaNumericString checks whether a given string consists only of alphanumeric characters (letters and digits).
func isAlphaNumericString(s string) bool {
	for _, c := range s {
		if !IsAlphaNumericRune(c) {
			return false
		}
	}
	return true
}

// isAlphaString checks if the given string contains only alphabetic characters.
func isAlphaString(s string) bool {
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
