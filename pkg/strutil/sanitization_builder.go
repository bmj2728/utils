package strutil

// CleanWhitespace removes all whitespace characters from the StringBuilder's value
// and returns the updated StringBuilder.
func (sb *StringBuilder) CleanWhitespace() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(cleanWhitespace(sb.value))
	return sb
}

// CleanWhitespaceWithIgnore removes whitespace from the StringBuilder's value
// except those whitespace chars specified in the provided charset.
func (sb *StringBuilder) CleanWhitespaceWithIgnore(charset string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(cleanWhitespaceWithIgnore(sb.value, charset))
	return sb
}

// KeepAlpha removes all non-alphabetic characters from the StringBuilder's value,
// optionally retaining whitespace if ws is true.
func (sb *StringBuilder) KeepAlpha(ws bool) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(alphaRemove(sb.value, ws))
	return sb
}

// KeepAlphaNumeric removes all non-alphanumeric characters from the StringBuilder's value,
// optionally preserving whitespace.
func (sb *StringBuilder) KeepAlphaNumeric(ws bool) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(alphaNumericRemove(sb.value, ws))
	return sb
}

// StripHTML removes all HTML tags from the StringBuilder's value and returns the updated StringBuilder.
func (sb *StringBuilder) StripHTML() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(stripHTML(sb.value))
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

// SanitizeHTMLCustom sanitizes the StringBuilder's value, allowing only specified HTML tags,
// and returns the updated StringBuilder.
func (sb *StringBuilder) SanitizeHTMLCustom(allowedTags []string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(sanitizeHTMLCustom(sb.value, allowedTags))
	return sb
}

func (sb *StringBuilder) RemoveNonPrintable() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) StripAnsi() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) SanitizeFilename() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) SanitizePath() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) SanitizeEmail() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) SanitizeUsername() *StringBuilder {
	panic("Implement me!")
}
