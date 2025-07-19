package strutil

import (
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"unicode"

	"github.com/google/uuid"
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

	domainRegex := regexp.MustCompile(`^(` + LabelRegex + `\.)+` + TLDRegex + `$`)

	return domainRegex.MatchString(strings.TrimSpace(domain))
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
	return len(normalizeWhitespace(s)) == 0
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

// isAlphaString checks if the given string consists only of alphabetic characters.
// Returns true if all characters are letters.
func isAlphaString(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
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

// isWhiteSpaceRune checks if the given rune is classified as a whitespace character based on Unicode standards.
func isWhiteSpaceRune(r rune) bool {
	return unicode.IsSpace(r)
}
