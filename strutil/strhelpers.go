package strutil

import (
	"net/mail"
	"strings"
)

// cleanWhitespace removes all whitespace characters (spaces, tabs, newlines, etc.) from the input string and returns the result.
func cleanWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, c := range s {
		if c != ' ' && c != '\t' && c != '\n' && c != '\r' {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// isEmail checks if the provided string s is a valid email address format and returns true if valid, false otherwise.
func isEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}
