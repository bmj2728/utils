package strutil

import "strings"

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
