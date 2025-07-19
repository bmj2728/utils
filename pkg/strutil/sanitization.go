package strutil

// SanitizeHTML removes potentially unsafe or harmful content from the input HTML string.
func SanitizeHTML(s string) string {
	return sanitizeHTML(s)
}

// SanitizeHTMLCustom sanitizes an HTML string, allowing only the tags specified in the allowedTags slice.
// Returns a sanitized string with disallowed HTML elements removed or escaped.
func SanitizeHTMLCustom(s string, allowedTags []string) string {
	return sanitizeHTMLCustom(s, allowedTags)
}

func SanitizeFilename(s string) string {
	panic("Implement me!")
}

func SanitizePath(s string) string {
	panic("Implement me!")
}

func SanitizeEmail(s string) string {
	panic("Implement me!")
}

func SanitizeUsername(s string) string {
	panic("Implement me!")
}

func SanitizeSearchQuery(s string) string {
	panic("Implement me!")
}

func SanitizePhone(s string) string {
	panic("Implement me!")
}

func SanitizeCardNumber(s string) string {
	panic("Implement me!")
}

func SanitizePostalCode(s string) string {
	panic("Implement me!")
}

func SanitizeSSN(s string) string {
	panic("Implement me!")
}

func SanitizeEnvValue(s string) string {
	panic("Implement me!")
}

func SanitizeCommandLineArg(s string) string {
	panic("Implement me!")
}

func SanitizeLogMessage(s string) string {
	panic("Implement me!")
}

func SanitizeLogMessageWithLevel(s string, level string) string {
	panic("Implement me!")
}

func SanitizeShellArg(s string) string {
	panic("Implement me!")
}

// CleanWhitespace removes all whitespace characters (spaces, tabs, newlines, etc.)
// from the input string and returns the result.
func CleanWhitespace(s string) string {
	return cleanWhitespace(s)
}

// CleanWhitespaceWithIgnore removes whitespace from the input string while
// ignoring whitespace characters specified in the charset string.
// Whitespace Chars: '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP)
func CleanWhitespaceWithIgnore(s string, charset string) string {
	return cleanWhitespaceWithIgnore(s, charset)
}

// KeepAlpha removes all non-alphabetic characters from the input string, optionally keeping whitespace if ws is true.
func KeepAlpha(s string, ws bool) string {
	return alphaRemove(s, ws)
}

// KeepAlphaNumeric removes all non-alphanumeric characters from the input string,
// optionally preserving whitespace if ws is true.
func KeepAlphaNumeric(s string, ws bool) string {
	return alphaNumericRemove(s, ws)
}

// EscapeHTML escapes special HTML characters in a string, replacing them with their corresponding HTML entity codes.
func EscapeHTML(s string) string {
	return escapeHTML(s)
}

// StripHTML removes all HTML tags from the input string and sanitizes it to ensure safe content.
func StripHTML(s string) string {
	return stripHTML(s)
}

func RemoveNonPrintable(s string) string {
	panic("Implement me!")
}

func StripAnsi(s string) string {
	panic("Implement me!")
}

func MaskCardNumber(s string) string {
	panic("Implement me!")
}

func MaskSSN(s string) string {
	panic("Implement me!")
}

func EscapeSQL(s string) string {
	panic("Implement me!")
}
