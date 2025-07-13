package strutil

import (
	"github.com/google/uuid"
	"net/mail"
	"net/url"
	"strings"
	"unicode"
)

// cleanWhitespace removes all whitespace characters (spaces, tabs, newlines, etc.) from the input string and returns the result.
func cleanWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, c := range s {
		if !unicode.IsSpace(c) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// normalizeWhitespace collapses whitespace runs to single spaces and trims
func normalizeWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// collapseWhitespace reduces all consecutive whitespace characters in the input string to a single space preserving leading and trailing whitespace.
func collapseWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	prevWasSpace := false

	for _, r := range s {
		if unicode.IsSpace(r) {
			if !prevWasSpace {
				b.WriteRune(' ')
			}
			prevWasSpace = true
		} else {
			b.WriteRune(r)
			prevWasSpace = false
		}
	}
	return b.String()
}

// isValidEmail checks if the provided string s is a valid email address format and returns true if valid, false otherwise.
func isValidEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

// isValidUrl checks if the provided string is a valid URL with a defined scheme and host. Returns true if valid, false otherwise.
func isValidUrl(s string) bool {
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

// isValidUUID checks if the provided string is a valid UUID. Returns true if valid, otherwise false.
func isValidUUID(s string) bool {
	err := uuid.Validate(s)
	if err != nil {
		return false
	}
	return true
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
	return len(normalizeWhitespace(s)) == 0
}

// toUpper converts the input string to uppercase and returns the result.
func toUpper(s string) string {
	return strings.ToUpper(s)
}

// toLower converts all characters in the input string to their lowercase equivalents and returns the resulting string.
func toLower(s string) string {
	return strings.ToLower(s)
}

// isAlphaNumericString checks whether a given string consists only of alphanumeric characters (letters and digits).
func isAlphaNumericString(s string) bool {
	for _, c := range s {
		if !isAlphaNumericRune(c) {
			return false
		}
	}
	return true
}

// isAlphaNumericRune determines if the given rune is an alphanumeric character (letter or digit).
func isAlphaNumericRune(r rune) bool {
	if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
		return false
	}
	return true
}
