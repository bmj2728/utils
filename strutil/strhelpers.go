package strutil

import (
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
