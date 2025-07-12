package strutil

// SanitizeHTML removes potentially unsafe or harmful content from the input HTML string.
func SanitizeHTML(s string) string {
	return sanitizeHTML(s)
}

// RemoveWhitespace removes all whitespace characters (spaces, tabs, newlines, etc.)
// from the input string and returns the result.
func RemoveWhitespace(s string) string {
	return removeWhitespace(s)
}

// RemoveWhitespaceWithIgnore removes whitespace from the input string while
// ignoring whitespace characters specified in the charset string.
// Whitespace Chars: '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP)
func RemoveWhitespaceWithIgnore(s string, charset string) string {
	return removeWhitespaceWithIgnore(s, charset)
}

// RemoveNonAlpha removes all non-alphabetic characters from the input string,
// optionally keeping whitespace if ws is true.
func RemoveNonAlpha(s string, ws bool) string {
	return removeNonAlpha(s, ws)
}

// RemoveNonAlphaNumeric removes all non-alphanumeric characters from the input string,
// optionally preserving whitespace if ws is true.
func RemoveNonAlphaNumeric(s string, ws bool) string {
	return removeNonAlphaNumeric(s, ws)
}

// EscapeHTML escapes special HTML characters in a string, replacing them with their corresponding HTML entity codes.
func EscapeHTML(s string) string {
	return escapeHTML(s)
}

// RemoveHTML removes all HTML tags from the input string and sanitizes it to ensure safe content.
func RemoveHTML(s string, preserveSpace bool) string {
	return removeHTML(s, preserveSpace)
}

// RemoveNonPrintable removes non-printable characters from a string and substitutes them with '_'.
func RemoveNonPrintable(s string) string {
	return removeNonPrintable(s)
}

// RemoveANSIEscapeCodes removes ANSI escape codes from the input string, returning a
// cleaned version without formatting sequences.
func RemoveANSIEscapeCodes(s string) string {
	return removeANSIEscapeCodes(s)
}
