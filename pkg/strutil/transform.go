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
func NormalizeWhitespace(s string, whitespace rune) string {
	return normalizeWhitespace(s, whitespace)
}

// NormalizeWhitespaceWithIgnore replaces all whitespace in a string with the specified rune, excluding ignored ones.
func NormalizeWhitespaceWithIgnore(s string, whitespace rune, ignoreWhitespace string) string {
	return normalizeWhitespaceWithIgnore(s, whitespace, ignoreWhitespace)
}

// CollapseWhitespace reduces all consecutive whitespace characters in a string to a single space,
// preserving leading and trailing whitespace.
func CollapseWhitespace(s string) string {
	return collapseWhitespace(s)
}

// CollapseWhitespaceWithIgnore reduces consecutive whitespace to a single space while ignoring
// whitespace in the 'ignore' string.
func CollapseWhitespaceWithIgnore(s string, ignore string) string {
	return collapseWhitespaceWithIgnore(s, ignore)
}

// ReplaceWhitespace replaces all whitespace characters in the input string with the given replacement string.
func ReplaceWhitespace(s string, replacement string) string {
	return replaceWhitespace(s, replacement)
}

// ReplaceWhitespaceWithIgnore replaces all whitespace in the input string with the specified replacement,
// ignoring characters present in the ignore string.
func ReplaceWhitespaceWithIgnore(s string, replacement string, ignore string) string {
	return replaceWhitespaceWithIgnore(s, replacement, ignore)
}

// ReplaceSpaces replaces all spaces in the input string with the specified replacement string.
func ReplaceSpaces(s string, replacement string) string {
	return replaceSpaces(s, replacement)
}

// ReplaceNonAlpha replaces all non-alphabetic characters in the input string with the specified replacement string.
func ReplaceNonAlpha(s string, replacement string) string {
	return replaceNonAlpha(s, replacement)
}

// ReplaceNonAlphaWithIgnore replaces non-alphabetic characters in a string with a replacement string,
// excluding specified ignores.
func ReplaceNonAlphaWithIgnore(s string, replacement string, ignore string) string {
	return replaceNonAlphaWithIgnore(s, replacement, ignore)
}

// ReplaceNonAlphaNumeric replaces all non-alphanumeric characters in the input string
// with the specified replacement string.
func ReplaceNonAlphaNumeric(s string, replacement string) string {
	return replaceNonAlphaNumeric(s, replacement)
}

// ReplaceNonAlphaNumericWithIgnore replaces non-alphanumeric characters in the
// input string with the replacement string,
// while preserving characters specified in the ignore string.
func ReplaceNonAlphaNumericWithIgnore(s string, replacement string, ignore string) string {
	return replaceNonAlphaNumericWithIgnore(s, replacement, ignore)
}

// NormalizeDiacritics removes diacritical marks (accents) from the input string and returns the normalized version.
func NormalizeDiacritics(s string) string {
	return normalizeDiacritics(s)
}

// NormalizeUnicode normalizes a string to the specified Unicode normalization form (NFC, NFD, NFKC, or NFKD).
func NormalizeUnicode(s string, form NormalizationFormat) string {
	return normalizeUnicode(s, form)
}

// RemovePrefix removes the specified prefix from a string and returns the resulting string.
func RemovePrefix(s string, prefix string) string {
	return removePrefix(s, prefix)
}

// RemovePrefixWithResult removes the specified prefix from the input string and
// returns the modified string and a boolean.
// The boolean indicates whether the prefix was found and removed.
func RemovePrefixWithResult(s string, prefix string) (string, bool) {
	return removePrefixWithResult(s, prefix)
}

// RemoveSuffix removes the specified suffix from the input string if it exists and
// returns the resulting string.
func RemoveSuffix(s string, suffix string) string {
	return removeSuffix(s, suffix)
}

// RemoveSuffixWithResult removes the specified suffix from the input string and
// returns the modified string and a boolean.
// The boolean indicates whether the suffix was found and removed.
func RemoveSuffixWithResult(s string, suffix string) (string, bool) {
	return removeSuffixWithResult(s, suffix)
}

// AddLeftPadding adds spaces to the left of a string until it reaches the specified total
// length and returns the result.
func AddLeftPadding(s string, length int) string {
	return addLeftPadding(s, length)
}

// AddRightPadding appends spaces to the right of a string until it reaches the specified total length.
func AddRightPadding(s string, length int) string {
	return addRightPadding(s, length)
}

// AddPadding adds equal padding of spaces to both sides of a string for the specified length and returns the result.
func AddPadding(s string, length int) string {
	return addPadding(s, length)
}

// LeftPadToLength pads the input string with spaces on the left to reach the specified length.
// If the string already meets or exceeds the length, it is returned unchanged.
func LeftPadToLength(s string, length int) string {
	return leftPadToLength(s, length)
}

// RightPadToLength pads the input string with spaces on the right until it reaches the specified length.
// Returns the input string unchanged if it is empty or if the specified length is less than 1.
func RightPadToLength(s string, length int) string {
	return rightPadToLength(s, length)
}

// PadToLength adjusts the given string to the specified length by padding spaces evenly on both sides.
// If the string exceeds the specified length, it returns the original string.
// The provided length must be non-negative for proper functionality.
func PadToLength(s string, length int, equalize bool) string {
	return padToLength(s, length, equalize)
}
