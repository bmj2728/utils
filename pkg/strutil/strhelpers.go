package strutil

import (
	"strings"
	"unicode"

	"github.com/microcosm-cc/bluemonday"
	"github.com/mrz1836/go-sanitize"
)

// cleanWhitespace removes all whitespace characters (spaces, tabs, newlines, etc.)
// from the input string and returns the result.
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

// cleanWhitespaceWithIgnore removes all whitespace characters from the input string
// while ignoring whitespace characters in the given charset.
func cleanWhitespaceWithIgnore(s string, charset string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, c := range s {
		if !unicode.IsSpace(c) || (unicode.IsSpace(c) && strings.ContainsRune(charset, c)) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// normalizeWhitespace collapses whitespace runs to single spaces and trims
func normalizeWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// collapseWhitespace reduces all consecutive whitespace characters in the input string to a single space preserving
// leading and trailing whitespace.
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

// replaceWhitespace replaces all whitespace characters in the input string with the specified replacement string.
func replaceWhitespace(s string, replacement string) string {
	if isEmpty(s) {
		return s
	}

	for _, c := range s {
		if isWhiteSpaceRune(c) {
			s = strings.Replace(s, string(c), replacement, 1)
		}
	}
	return s
}

// replaceSpaces replaces all spaces in the input string with the specified replacement string.
func replaceSpaces(s string, replacement string) string {
	return strings.ReplaceAll(s, " ", replacement)
}

// alpha removes all non-alphabetic characters from the given string, optionally retaining whitespace if ws is true.
func alpha(s string, ws bool) string {
	return sanitize.Alpha(s, ws)
}

// alphaNumeric removes all non-alphanumeric characters from the input string,
// optionally preserving whitespace if ws is true.
func alphaNumeric(s string, ws bool) string {
	return sanitize.AlphaNumeric(s, ws)
}

// alphaReplace replaces all non-alphabetic characters in the input string with the specified replacement string.
func alphaReplace(s string, replacement string) string {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			s = strings.Replace(s, string(c), replacement, 1)
		}
	}
	return s
}

// alphaNumericReplace replaces all non-alphanumeric characters
// in the input string with the specified replacement string.
func alphaNumericReplace(s string, replacement string) string {
	for _, c := range s {
		if !isAlphaNumericRune(c) {
			s = strings.Replace(s, string(c), replacement, 1)
		}
	}
	return s
}

// stripHTML removes all HTML tags and sanitizes the input string to prevent potential security risks.
func stripHTML(s string) string {
	p := bluemonday.StrictPolicy()
	return p.Sanitize(s)
}

// escapeHTML escapes special HTML characters in a string to their corresponding HTML entity codes.
func escapeHTML(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		switch r {
		case '"':
			b.WriteString("&quot;")
		case '&':
			b.WriteString("&amp;")
		case '<':
			b.WriteString("&lt;")
		case '>':
			b.WriteString("&gt;")
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}
