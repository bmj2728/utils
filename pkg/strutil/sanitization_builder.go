package strutil

// RemoveWhitespace removes all whitespace characters from the StringBuilder's value
// and returns the updated StringBuilder.
func (sb *StringBuilder) RemoveWhitespace() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(removeWhitespace(sb.value))
	return sb
}

// RemoveWhitespaceWithIgnore removes whitespace from the StringBuilder's value
// except those whitespace chars specified in the provided charset.
func (sb *StringBuilder) RemoveWhitespaceWithIgnore(charset string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(removeWhitespaceWithIgnore(sb.value, charset))
	return sb
}

// RemoveNonAlpha removes all non-alphabetic characters from the StringBuilder's value,
// optionally retaining whitespace if ws is true.
func (sb *StringBuilder) RemoveNonAlpha(ws bool) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(removeNonAlpha(sb.value, ws))
	return sb
}

// RemoveNonAlphaNumeric removes all non-alphanumeric characters from the StringBuilder's value,
// optionally preserving whitespace.
func (sb *StringBuilder) RemoveNonAlphaNumeric(ws bool) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(removeNonAlphaNumeric(sb.value, ws))
	return sb
}

// RemoveHTML removes all HTML tags from the StringBuilder's value and returns the updated StringBuilder.
func (sb *StringBuilder) RemoveHTML(preserveSpace bool) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(removeHTML(sb.value, preserveSpace))
	return sb
}

// EscapeHTML escapes special HTML characters in the StringBuilder's value and returns the updated StringBuilder.
func (sb *StringBuilder) EscapeHTML() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(escapeHTML(sb.value))
	return sb
}

// SanitizeHTML sanitizes the StringBuilder's value by removing potentially unsafe or harmful HTML content.
func (sb *StringBuilder) SanitizeHTML() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(sanitizeHTML(sb.value))
	return sb
}

// RemoveNonPrintable removes non-printable characters from the StringBuilder's value and replaces them with '_'.
// Returns the modified StringBuilder instance.
func (sb *StringBuilder) RemoveNonPrintable() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(removeNonPrintable(sb.value))
	return sb
}

// RemoveANSIEscapeCodes removes ANSI escape codes from the string stored in the StringBuilder and updates its value.
func (sb *StringBuilder) RemoveANSIEscapeCodes() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(removeANSIEscapeCodes(sb.value))
	return sb
}
