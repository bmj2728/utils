package strutil

// Append adds the given string `s` to the current value with the specified
// separator `sep` and returns the updated StringBuilder.
func (sb *StringBuilder) Append(s string, sep string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(appendString(sb.value, s, sep))
	return sb
}

// Prepend adds the specified string and separator to the beginning of the StringBuilder's value.
func (sb *StringBuilder) Prepend(s string, ser string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(prependString(sb.value, s, ser))
	return sb
}

// Trim removes leading and trailing whitespace or specific characters from the StringBuilder's value.
func (sb *StringBuilder) Trim() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(trim(sb.value))
	return sb
}

// TrimLeft removes all leading whitespace characters from the string stored in the StringBuilder and updates its value.
func (sb *StringBuilder) TrimLeft() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(trimLeft(sb.value))
	return sb
}

// TrimRight removes trailing whitespace or specified characters from the end of the StringBuilder's value.
func (sb *StringBuilder) TrimRight() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(trimRight(sb.value))
	return sb
}

// TrimChars removes all leading and trailing characters specified in the input string from the StringBuilder's value.
func (sb *StringBuilder) TrimChars(chars string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(trimChars(sb.value, chars))
	return sb
}

// TrimCharsLeft removes all leading occurrences of the specified characters from the string
// and returns the updated StringBuilder.
func (sb *StringBuilder) TrimCharsLeft(chars string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(trimCharsLeft(sb.value, chars))
	return sb
}

// TrimCharsRight removes all occurrences of the specified characters from the end of the string
// and returns the StringBuilder.
func (sb *StringBuilder) TrimCharsRight(chars string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(trimCharsRight(sb.value, chars))
	return sb
}

// NormalizeDiacritics removes diacritical marks from the string and replaces them with their non-accented counterparts.
func (sb *StringBuilder) NormalizeDiacritics() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(normalizeDiacritics(sb.value))
	return sb
}

// Slugify converts the string into a URL-friendly slug with a maximum length specified by the parameter.
func (sb *StringBuilder) Slugify(length int) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(slugify(sb.value, length))
	return sb
}

// Truncate shortens the string to the specified length and appends the provided suffix if truncation occurs.
func (sb *StringBuilder) Truncate(length int, suffix string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(truncate(sb.value, length, suffix))
	return sb
}

// If applies the provided function to the StringBuilder's value if the condition is true and continues processing.
func (sb *StringBuilder) If(condition bool, fn func(string) string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if condition {
		sb.setValue(fn(sb.value))
	}
	return sb
}

// Transform applies a custom transformation function to the StringBuilder's value and returns the updated instance.
//
// Examples:
//
// s := New("Hello World!").Transform(strings.ToUpper)
//
// result: "HELLO WORLD!"
//
//	s := New("Hello World!).Transform(func(input string) string {
//		return input + " Goodbye!"
//	})
//
// result: "Hello World! Goodbye!"
func (sb *StringBuilder) Transform(fn func(string) string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(fn(sb.value))
	return sb
}

// NormalizeWhitespace collapses consecutive whitespace characters into a single space
// and trims leading and trailing spaces.
func (sb *StringBuilder) NormalizeWhitespace(whitespace rune) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(normalizeWhitespace(sb.value, whitespace))
	return sb
}

// NormalizeWhitespaceWithIgnore replaces whitespace characters in the StringBuilder's value with a specified rune,
// except those present in the ignoreChars. Returns the updated StringBuilder.
func (sb *StringBuilder) NormalizeWhitespaceWithIgnore(whitespace rune, ignoreChars string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(normalizeWhitespaceWithIgnore(sb.value, whitespace, ignoreChars))
	return sb
}

// CollapseWhitespace collapses consecutive whitespace characters in the StringBuilder's value into a single space
// and preserves leading and trailing spaces.
func (sb *StringBuilder) CollapseWhitespace() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(collapseWhitespace(sb.value))
	return sb
}

// CollapseWhitespaceWithIgnore reduces multiple sequential whitespace characters to a single
// instance, ignoring specified characters.
func (sb *StringBuilder) CollapseWhitespaceWithIgnore(ignoreChars string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(collapseWhitespaceWithIgnore(sb.value, ignoreChars))
	return sb
}

// ReplaceWhitespace replaces all whitespace characters in the StringBuilder's value
// with the specified replacement string.
func (sb *StringBuilder) ReplaceWhitespace(replacement string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(replaceWhitespace(sb.value, replacement))
	return sb
}

// ReplaceWhitespaceWithIgnore replaces all whitespace in the string with the specified replacement,
// ignoring specified characters.
// It modifies the current StringBuilder's value and skips processing if shouldContinueProcessing returns false.
func (sb *StringBuilder) ReplaceWhitespaceWithIgnore(replacement string, ignoreChars string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(replaceWhitespaceWithIgnore(sb.value, replacement, ignoreChars))
	return sb
}

// ReplaceSpaces replaces all spaces in the StringBuilder's value with the specified replacement string.
func (sb *StringBuilder) ReplaceSpaces(replacement string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(replaceSpaces(sb.value, replacement))
	return sb
}

// ReplaceNonAlpha replaces all alphabetical characters in the StringBuilder's value
// with the specified replacement string.
func (sb *StringBuilder) ReplaceNonAlpha(replacement string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(replaceNonAlpha(sb.value, replacement))
	return sb
}

// ReplaceNonAlphaWithIgnore replaces non-alphabetic characters in the current string with a
// replacement, excluding ignoreChars.
// It returns the updated StringBuilder after modification.
// The operation is skipped if processing is disabled for the StringBuilder.
func (sb *StringBuilder) ReplaceNonAlphaWithIgnore(replacement string, ignoreChars string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(replaceNonAlphaWithIgnore(sb.value, replacement, ignoreChars))
	return sb
}

// ReplaceNonAlphaNumeric replaces all alphanumeric characters in the StringBuilder value
// with the specified replacement string.
func (sb *StringBuilder) ReplaceNonAlphaNumeric(replacement string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(replaceNonAlphaNumeric(sb.value, replacement))
	return sb
}

// ReplaceNonAlphaNumericWithIgnore replaces non-alphanumeric characters in the string except for those in ignoreChars.
// Returns the modified StringBuilder instance.
func (sb *StringBuilder) ReplaceNonAlphaNumericWithIgnore(replacement string, ignoreChars string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(replaceNonAlphaNumericWithIgnore(sb.value, replacement, ignoreChars))
	return sb
}

// NormalizeUnicode normalizes the StringBuilder's string to the specified Unicode
// normalization form (NFC, NFD, NFKC, or NFKD).
func (sb *StringBuilder) NormalizeUnicode(form NormalizationFormat) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(normalizeUnicode(sb.value, form))
	return sb
}
