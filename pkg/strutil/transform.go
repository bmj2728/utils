package strutil

// AppendString concatenates the input string `s` with `suffix`, using `sep`
// as the separator, and returns the resulting string.
func AppendString(s string, suffix string, sep string) string {
	return appendString(s, suffix, sep)
}

// PrependString adds a prefix and separator to the beginning of the given string and returns the resulting string.
func PrependString(s string, prefix string, sep string) string {
	return prependString(s, prefix, sep)
}

// Trim removes all leading and trailing white spaces from the given string and returns the trimmed result.
func Trim(s string) string {
	return trim(s)
}

// TrimLeft removes all leading whitespace characters from the input string `s`.
func TrimLeft(s string) string {
	return trimLeft(s)
}

// TrimRight removes all trailing whitespace characters from the given string.
func TrimRight(s string) string {
	return trimRight(s)
}

// TrimChars removes all leading and trailing occurrences of specified characters from the input string.
func TrimChars(s string, chars string) string {
	return trimChars(s, chars)
}

// TrimCharsLeft removes all leading characters specified in 'chars' from the input string 's'.
func TrimCharsLeft(s string, chars string) string {
	return trimCharsLeft(s, chars)
}

// TrimCharsRight removes all specified characters from the end of the given string.
func TrimCharsRight(s string, chars string) string {
	return trimCharsRight(s, chars)
}

// Slugify converts a string to a URL-friendly slug, ensuring lowercase, trimming, truncation,
// and replacing non-alphanumerics.
func Slugify(s string, length int) string {
	return slugify(s, length)
}

// Truncate shortens the input string s to the specified length and appends the given suffix if truncation occurs.
func Truncate(s string, length int, suffix string) string {
	return truncate(s, length, suffix)
}

// NormalizeWhitespace removes excess whitespace by trimming and collapsing
// multiple whitespace characters into single spaces.
func NormalizeWhitespace(s string) string {
	return normalizeWhitespace(s)
}

// CollapseWhitespace reduces all consecutive whitespace characters in a string to a single space,
// preserving leading and trailing whitespace.
func CollapseWhitespace(s string) string {
	return collapseWhitespace(s)
}

// ReplaceWhitespace replaces all whitespace characters in the input string with the given replacement string.
func ReplaceWhitespace(s string, replacement string) string {
	return replaceWhitespace(s, replacement)
}

// ReplaceSpaces replaces all spaces in the input string with the specified replacement string.
func ReplaceSpaces(s string, replacement string) string {
	return replaceSpaces(s, replacement)
}

// AlphaReplace replaces all non-alphabetic characters in the input string with the specified replacement string.
func AlphaReplace(s string, replacement string) string {
	return alphaReplace(s, replacement)
}

// AlphaNumericReplace replaces all non-alphanumeric characters in the input string
// with the specified replacement string.
func AlphaNumericReplace(s string, replacement string) string {
	return alphaNumericReplace(s, replacement)
}

// NormalizeDiacritics removes diacritical marks (accents) from the input string and returns the normalized version.
func NormalizeDiacritics(s string) string {
	return normalizeDiacritics(s)
}

func NormalizeUnicode(s string) string {
	panic("Implement me!")
}
