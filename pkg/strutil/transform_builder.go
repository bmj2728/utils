package strutil

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
