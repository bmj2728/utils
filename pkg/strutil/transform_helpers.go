package strutil

import (
	"strings"
	"unicode"

	godiacritics "github.com/Regis24GmbH/go-diacritics"
	"golang.org/x/text/unicode/norm"
)

// truncate shortens the input string s to the specified length and appends the given suffix if truncation occurs.
func truncate(s string, length int, suffix string) string {
	if len(s) <= length {
		return s
	}
	if length < 0 {
		return ""
	}
	return s[:length] + suffix
}

// appendString concatenates the given string `s` with `suffix`,
// separated by the specified `sep`, and returns the result.
func appendString(s string, suffix string, sep string) string {
	if len(suffix) < 1 {
		return s
	}
	return s + sep + suffix
}

// prependString concatenates a prefix and a separator to the beginning of a string and returns the resulting string.
func prependString(s string, prefix string, sep string) string {
	if len(prefix) < 1 {
		return s
	}
	return prefix + sep + s
}

// trim removes all leading and trailing white spaces from the input string and returns the trimmed result.
func trim(s string) string {
	return strings.TrimSpace(s)
}

// trimLeft removes all leading whitespace characters from the input string `s`
// based on the defined `WhiteSpaceChars` set.
func trimLeft(s string) string {
	return strings.TrimLeft(s, WhiteSpaceChars)
}

// trimRight removes all trailing whitespace characters from the provided string.
func trimRight(s string) string {
	return strings.TrimRight(s, WhiteSpaceChars)
}

// trimChars removes all leading and trailing occurrences of the characters specified in chars from the string s.
func trimChars(s string, chars string) string {
	return strings.Trim(s, chars)
}

// trimCharsLeft removes all leading characters found in 'chars' from the string 's'.
func trimCharsLeft(s string, chars string) string {
	return strings.TrimLeft(s, chars)
}

// trimCharsRight removes all occurrences of the specified characters from the end of the given string.
func trimCharsRight(s string, chars string) string {
	return strings.TrimRight(s, chars)
}

// slugify converts a given string into a URL-friendly slug, ensuring lowercase, truncation, and hyphenation if needed.
func slugify(s string, length int) string {

	//early return if empty string
	if s == "" || length < 1 {
		return ""
	}

	// address camelCase/PascalCase
	if CamelCaseRegex.MatchString(s) {
		s = splitCamelCase(s)
	}

	//clean the diacritics
	s = normalizeDiacritics(s)

	// replace non-alphanumerics
	s = replaceNonAlphaNumeric(s, " ")

	s = collapseWhitespace(s)

	// normalize whitespace
	s = normalizeWhitespace(s, ' ')

	// replace whitespace with "-"
	s = replaceSpaces(s, "-")

	// make lower
	s = toLower(s)

	// if a length is provided, truncate to that length
	s = truncate(s, length, "")

	// ensure no misbehaving "-"
	s = trimChars(s, "-")

	return s
}

// normalizeDiacritics removes diacritical marks (accents) from the input string, returning the normalized version.
func normalizeDiacritics(s string) string {
	return godiacritics.Normalize(s)
}

// normalizeWhitespace replaces all whitespace characters in the input string with the
// specified rune and trims the result.
func normalizeWhitespace(s string, whitespace rune) string {
	if isEmpty(s) {
		return s
	}
	if !unicode.IsSpace(whitespace) {
		return s
	}
	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if unicode.IsSpace(r) {
			b.WriteRune(whitespace)
		} else {
			b.WriteRune(r)
		}
	}

	return trim(b.String())
}

// normalizeWhitespaceWithIgnore replaces all whitespace in the input string with a specified rune,
// except ignored whitespaces.
// It processes each character, preserving non-whitespace characters and ignored whitespaces as provided in the input.
// Returns the processed string with leading and trailing spaces trimmed.
func normalizeWhitespaceWithIgnore(s string, whitespace rune, ignoreWhitespace string) string {
	if isEmpty(s) {
		return s
	}
	if !unicode.IsSpace(whitespace) {
		return s
	}

	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if !unicode.IsSpace(r) {
			b.WriteRune(r)
		} else if strings.ContainsRune(ignoreWhitespace, r) {
			b.WriteRune(r)
		} else {
			b.WriteRune(whitespace)
		}
	}
	return trim(b.String())
}

// collapseWhitespace reduces consecutive whitespace characters in a string to a
// single instance while retaining single spaces.
func collapseWhitespace(s string) string {
	if isEmpty(s) {
		return s
	}

	var b strings.Builder
	b.Grow(len(s))

	var prevRune rune

	for _, r := range s {
		// if the current rune is whitespace and is not the previous rune, write it to the buffer
		if unicode.IsSpace(r) {
			if r != prevRune {
				b.WriteRune(r)
			}
			prevRune = r
		} else {
			// otherwise, write the current rune to the buffer
			b.WriteRune(r)
			prevRune = r
		}
	}
	return b.String()
}

// collapseWhitespaceWithIgnore reduces sequential whitespace to a single instance,
// except for characters in ignoreChars.
func collapseWhitespaceWithIgnore(s string, ignoreChars string) string {
	if isEmpty(s) {
		return s
	}

	var b strings.Builder
	b.Grow(len(s))

	var prevRune rune

	for _, r := range s {
		if unicode.IsSpace(r) {
			// if the current rune is whitespace and is not the previous rune, write it to the buffer
			// or if the current rune is in the ignoreChars string, write it to the buffer
			if r != prevRune || strings.ContainsRune(ignoreChars, r) {
				b.WriteRune(r)
			}
			prevRune = r
		} else {
			// otherwise, write the current rune to the buffer
			b.WriteRune(r)
			prevRune = r
		}
	}
	return b.String()
}

// replaceWhitespace replaces all whitespace characters in the input string with the specified replacement string.
func replaceWhitespace(s string, replacement string) string {
	if isEmpty(s) {
		return s
	}
	var b strings.Builder
	b.Grow(len(s) * 2)

	for _, c := range s {
		if unicode.IsSpace(c) {
			b.WriteString(replacement)
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

func replaceWhitespaceWithIgnore(s string, replacement string, ignoreChars string) string {
	if isEmpty(s) {
		return s
	}
	var b strings.Builder
	b.Grow(len(s) * 2)

	for _, c := range s {
		if unicode.IsSpace(c) && !strings.ContainsRune(ignoreChars, c) {
			b.WriteString(replacement)
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// replaceSpaces replaces all spaces in the input string with the specified replacement string.
func replaceSpaces(s string, replacement string) string {
	if isEmpty(s) {
		return s
	}
	var b strings.Builder
	b.Grow(len(s) * 2)
	for _, c := range s {
		if c == ' ' {
			b.WriteString(replacement)
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// replaceNonAlpha replaces all non-alphabetic characters in the input string with the specified replacement string.
func replaceNonAlpha(s string, replacement string) string {
	if isEmpty(s) {
		return s
	}
	var b strings.Builder
	b.Grow(len(s) * 2)
	for _, c := range s {
		if !unicode.IsLetter(c) {
			b.WriteString(replacement)
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// replaceNonAlphaWithIgnore replaces non-alphabetic characters in a string with a replacement string,
// excluding specified ignoreChars.
func replaceNonAlphaWithIgnore(s string, replacement string, ignoreChars string) string {
	if isEmpty(s) {
		return s
	}
	var b strings.Builder
	b.Grow(len(s) * 2)
	for _, c := range s {
		if !unicode.IsLetter(c) && !strings.ContainsRune(ignoreChars, c) {
			b.WriteString(replacement)
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// replaceNonAlphaNumeric replaces all non-alphanumeric characters
// in the input string with the specified replacement string.
func replaceNonAlphaNumeric(s string, replacement string) string {
	if isEmpty(s) {
		return s
	}
	var b strings.Builder
	b.Grow(len(s) * 2)
	for _, c := range s {
		if !IsAlphaNumericRune(c) {
			b.WriteString(replacement)
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// replaceNonAlphaNumericWithIgnore replaces non-alphanumeric characters in a string, excluding specified ignoreChars.
func replaceNonAlphaNumericWithIgnore(s string, replacement string, ignoreChars string) string {
	if isEmpty(s) {
		return s
	}
	var b strings.Builder
	b.Grow(len(s) * 2)
	for _, c := range s {
		if !IsAlphaNumericRune(c) && !strings.ContainsRune(ignoreChars, c) {
			b.WriteString(replacement)
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// normalizeUnicode converts a string to normalized Unicode form specified by the norm.Form argument.
// NFC Unicode Normalization Form C
// NFD Unicode Normalization Form D
// NFKC Unicode Normalization Form KC
// NFKD Unicode Normalization Form KD
func normalizeUnicode(s string, format NormalizationFormat) string {
	return string(norm.Form(format).Bytes([]byte(s)))
}
