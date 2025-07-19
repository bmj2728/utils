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
