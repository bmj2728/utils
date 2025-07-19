package strutil

// Fluent StringBuilder API

// Constructor

// Manipulation Methods

// Append adds the given string `s` to the current value with the specified
// separator `sep` and returns the updated StringBuilder.
func (sb *StringBuilder) Append(s string, sep string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = appendString(sb.value, s, sep)
	return sb
}

// Prepend adds the specified string and separator to the beginning of the StringBuilder's value.
func (sb *StringBuilder) Prepend(s string, ser string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = prependString(sb.value, s, ser)
	return sb
}

// CleanWhitespace removes all whitespace characters from the StringBuilder's value
// and returns the updated StringBuilder.
func (sb *StringBuilder) CleanWhitespace() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = cleanWhitespace(sb.value)
	return sb
}

// CleanWhitespaceWithIgnore removes whitespace from the StringBuilder's value
// except those whitespace chars specified in the provided charset.
func (sb *StringBuilder) CleanWhitespaceWithIgnore(charset string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = cleanWhitespaceWithIgnore(sb.value, charset)
	return sb
}

// NormalizeWhitespace collapses consecutive whitespace characters into a single space
// and trims leading and trailing spaces.
func (sb *StringBuilder) NormalizeWhitespace() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = normalizeWhitespace(sb.value)
	return sb
}

// CollapseWhitespace collapses consecutive whitespace characters in the StringBuilder's value into a single space
// and preserves leading and trailing spaces.
func (sb *StringBuilder) CollapseWhitespace() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = collapseWhitespace(sb.value)
	return sb
}

// ReplaceWhitespace replaces all whitespace characters in the StringBuilder's value
// with the specified replacement string.
func (sb *StringBuilder) ReplaceWhitespace(replacement string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = replaceWhitespace(sb.value, replacement)
	return sb
}

// ReplaceSpaces replaces all spaces in the StringBuilder's value with the specified replacement string.
func (sb *StringBuilder) ReplaceSpaces(replacement string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = replaceSpaces(sb.value, replacement)
	return sb
}

// KeepAlpha removes all non-alphabetic characters from the StringBuilder's value,
// optionally retaining whitespace if ws is true.
func (sb *StringBuilder) KeepAlpha(ws bool) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = alpha(sb.value, ws)
	return sb
}

// KeepAlphaNumeric removes all non-alphanumeric characters from the StringBuilder's value,
// optionally preserving whitespace.
func (sb *StringBuilder) KeepAlphaNumeric(ws bool) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = alphaNumeric(sb.value, ws)
	return sb
}

// StripHTML removes all HTML tags from the StringBuilder's value and returns the updated StringBuilder.
func (sb *StringBuilder) StripHTML() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = stripHTML(sb.value)
	return sb
}

// EscapeHTML escapes special HTML characters in the StringBuilder's value and returns the updated StringBuilder.
func (sb *StringBuilder) EscapeHTML() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = escapeHTML(sb.value)
	return sb
}

// SanitizeHTML sanitizes the StringBuilder's value by removing potentially unsafe or harmful HTML content.
func (sb *StringBuilder) SanitizeHTML() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = sanitizeHTML(sb.value)
	return sb
}

// SanitizeHTMLCustom sanitizes the StringBuilder's value, allowing only specified HTML tags,
// and returns the updated StringBuilder.
func (sb *StringBuilder) SanitizeHTMLCustom(allowedTags []string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = sanitizeHTMLCustom(sb.value, allowedTags)
	return sb
}

// Trim removes leading and trailing whitespace or specific characters from the StringBuilder's value.
func (sb *StringBuilder) Trim() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trim(sb.value)
	return sb
}

// TrimLeft removes all leading whitespace characters from the string stored in the StringBuilder and updates its value.
func (sb *StringBuilder) TrimLeft() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trimLeft(sb.value)
	return sb
}

// TrimRight removes trailing whitespace or specified characters from the end of the StringBuilder's value.
func (sb *StringBuilder) TrimRight() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trimRight(sb.value)
	return sb
}

// TrimChars removes all leading and trailing characters specified in the input string from the StringBuilder's value.
func (sb *StringBuilder) TrimChars(chars string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trimChars(sb.value, chars)
	return sb
}

// TrimCharsLeft removes all leading occurrences of the specified characters from the string
// and returns the updated StringBuilder.
func (sb *StringBuilder) TrimCharsLeft(chars string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trimCharsLeft(sb.value, chars)
	return sb
}

// TrimCharsRight removes all occurrences of the specified characters from the end of the string
// and returns the StringBuilder.
func (sb *StringBuilder) TrimCharsRight(chars string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trimCharsRight(sb.value, chars)
	return sb
}

// AlphaReplace replaces all alphabetical characters in the StringBuilder's value with the specified replacement string.
func (sb *StringBuilder) AlphaReplace(replacement string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = alphaReplace(sb.value, replacement)
	return sb
}

// AlphaNumericReplace replaces all alphanumeric characters in the StringBuilder value
// with the specified replacement string.
func (sb *StringBuilder) AlphaNumericReplace(replacement string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = alphaNumericReplace(sb.value, replacement)
	return sb
}

// NormalizeDiacritics removes diacritical marks from the string and replaces them with their non-accented counterparts.
func (sb *StringBuilder) NormalizeDiacritics() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = normalizeDiacritics(sb.value)
	return sb
}

// Slugify converts the string into a URL-friendly slug with a maximum length specified by the parameter.
func (sb *StringBuilder) Slugify(length int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = slugify(sb.value, length)
	return sb
}

// Truncate shortens the string to the specified length and appends the provided suffix if truncation occurs.
func (sb *StringBuilder) Truncate(length int, suffix string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = truncate(sb.value, length, suffix)
	return sb
}

func (sb *StringBuilder) RemoveNonPrintable() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) NormalizeUnicode() *StringBuilder {
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

// ToLower converts all characters in the StringBuilder's value to lowercase and returns the updated StringBuilder.
func (sb *StringBuilder) ToLower() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toLower(sb.value)
	return sb
}

// ToUpper converts the StringBuilder's current value to uppercase and returns the updated StringBuilder.
func (sb *StringBuilder) ToUpper() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toUpper(sb.value)
	return sb
}

// Capitalize converts the first character of the StringBuilder's value to uppercase while preserving the rest as is.
func (sb *StringBuilder) Capitalize() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = capitalize(sb.value)
	return sb
}

// Uncapitalize converts the first character of the StringBuilder's value to lowercase if no error is present.
func (sb *StringBuilder) Uncapitalize() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = uncapitalize(sb.value)
	return sb
}

// ToTitleCase converts the string value of the StringBuilder to title case
// and returns the updated StringBuilder instance.
func (sb *StringBuilder) ToTitleCase() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toTitleCase(sb.value)
	return sb
}

// SplitCamelCase splits the string stored in the StringBuilder into separate words based on camel case boundaries.
func (sb *StringBuilder) SplitCamelCase() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = splitCamelCase(sb.value)
	return sb
}

// SplitPascalCase splits a PascalCase string into separate words, modifying the StringBuilder's value in-place.
func (sb *StringBuilder) SplitPascalCase() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = splitPascalCase(sb.value)
	return sb
}

// ToSnakeCase converts the current string to snake_case or SCREAMING_SNAKE_CASE based on the scream parameter.
func (sb *StringBuilder) ToSnakeCase(scream bool) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toSnakeCase(sb.value, scream)
	return sb
}

// ToSnakeCaseWithIgnore converts the StringBuilder's value to snake_case,
// optionally in uppercase, ignoring specified characters.
func (sb *StringBuilder) ToSnakeCaseWithIgnore(scream bool, ignore string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toSnakeCaseWithIgnore(sb.value, scream, ignore)
	return sb
}

// ToKebabCase converts the string in the StringBuilder to kebab-case or screaming-kebab-case based on the scream flag.
func (sb *StringBuilder) ToKebabCase(scream bool) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toKebabCase(sb.value, scream)
	return sb
}

// ToCamelCase converts the current string value to camel case format and updates the StringBuilder instance.
func (sb *StringBuilder) ToCamelCase() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toCamelCase(sb.value)
	return sb
}

// ToPascalCase converts the current string value of the StringBuilder
// to PascalCase format and updates the StringBuilder.
func (sb *StringBuilder) ToPascalCase() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toPascalCase(sb.value)
	return sb
}

// ToDelimited converts the string in the StringBuilder to a delimited format using the specified delimiter and options.
func (sb *StringBuilder) ToDelimited(delim uint8, ignore string, scream bool) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toDelimited(sb.value, delim, ignore, scream)
	return sb
}

// Comparison Methods

// Validation Methods (can set error)

// Control Flow

// If conditionally applies the provided function to the StringBuilder if the condition is true and no error exists.
func (sb *StringBuilder) If(condition bool, fn func(*StringBuilder) *StringBuilder) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if condition {
		return fn(sb)
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
	if sb.err != nil {
		return sb
	}
	sb.value = fn(sb.value)
	return sb
}

// Terminal Methods
