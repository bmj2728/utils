package strutil

import (
	"html"
	"strings"
	"unicode"

	"github.com/acarl005/stripansi"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mrz1836/go-sanitize"
)

var (
	// StrictPolicy is a sanitizer instance that removes all HTML tags
	// and only allows plain text content for strict enforcement.
	StrictPolicy = bluemonday.StrictPolicy()
	// StrictPolicyWithSpaces is a sanitizer that removes all HTML tags but
	//retains spaces when stripping tags for better readability.
	StrictPolicyWithSpaces = bluemonday.StrictPolicy().AddSpaceWhenStrippingTag(true)
	// UGCPolicy defines a bluemonday policy specifically designed for user-generated content sanitization.
	UGCPolicy = bluemonday.UGCPolicy()
)

// sanitizeHTML sanitizes an input HTML string by removing potentially unsafe or harmful content.
func sanitizeHTML(s string) string {
	return UGCPolicy.Sanitize(s)
}

// removeHTML removes all HTML tags and sanitizes the input string to prevent potential security risks.
func removeHTML(s string, preserveSpaces bool) string {
	stripped := ""
	if preserveSpaces {
		stripped = StrictPolicyWithSpaces.Sanitize(s)
	} else {
		stripped = StrictPolicy.Sanitize(s)
	}
	return strings.TrimSpace(stripped)
}

// escapeHTML escapes special HTML characters in a string to their corresponding HTML entity codes.
func escapeHTML(s string) string {
	return html.EscapeString(s)
}

// removeWhitespace removes all whitespace characters (spaces, tabs, newlines, etc.)
// from the input string and returns the result.
func removeWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, c := range s {
		if !unicode.IsSpace(c) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// removeWhitespaceWithIgnore removes all whitespace characters from the input string
// while ignoring whitespace characters in the given charset.
func removeWhitespaceWithIgnore(s string, charset string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, c := range s {
		if !unicode.IsSpace(c) || (unicode.IsSpace(c) && strings.ContainsRune(charset, c)) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// removeNonAlpha removes all non-alphabetic characters from the given string,
// optionally retaining whitespace if ws is true.
func removeNonAlpha(s string, ws bool) string {
	return sanitize.Alpha(s, ws)
}

// removeNonAlphaNumeric removes all non-alphanumeric characters from the input string,
// optionally preserving whitespace if ws is true.
func removeNonAlphaNumeric(s string, ws bool) string {
	return sanitize.AlphaNumeric(s, ws)
}

// removeNonPrintable replaces non-printable characters in a string with '_' and returns the modified string.
func removeNonPrintable(s string) string {
	if isEmpty(s) {
		return s
	}
	var b strings.Builder
	b.Grow(len(s))
	for _, c := range s {
		if unicode.IsPrint(c) {
			b.WriteRune(c)
		} else {
			b.WriteRune(' ')
		}
	}
	return b.String()
}

// removeANSIEscapeCodes removes ANSI escape codes from a given string, returning the clean text.
func removeANSIEscapeCodes(s string) string {
	return stripansi.Strip(s)
}
