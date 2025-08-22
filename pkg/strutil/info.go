package strutil

// Index finds the index of the first occurrence of `subStr` in `s`.
// If `normalized` is true, both strings are normalized to NFC before comparison.
// Returns the byte index of the match or -1 if not found.
func Index(s string, substr string, normalized bool) int {
	return index(s, substr, normalized)
}

// IndexAll returns a slice of starting indices of all non-overlapping occurrences of substrs in s.
// If normalized is true, both s and substrs are normalized to NFC before searching.
func IndexAll(s string, substr string, normalized bool) []int {
	return indexAll(s, substr, normalized)
}

// IndexRune returns the index of the first occurrence of the rune r in the string s, or -1 if r is not present in s.
func IndexRune(s string, r rune) int {
	return indexRune(s, r)
}

// IndexAllRune returns a slice of all indices where the specified rune appears in the input string.
func IndexAllRune(s string, r rune) []int {
	return indexAllRune(s, r)
}

// LastIndex returns the last index of `subStr` in `s`. If `normalized` is true,
// strings are NFC-normalized before searching.
func LastIndex(s string, substr string, normalized bool) int {
	return lastIndex(s, substr, normalized)
}

// LastIndexRune returns the index of the last occurrence of the rune r in the
// string s, or -1 if r is not present in s.
func LastIndexRune(s string, r rune) int {
	return lastIndexRune(s, r)
}

// UTF32NormalizedIndex finds the normalized index of subStr in s using UTF-32 encoding, returning -1 if not found.
func UTF32NormalizedIndex(s string, substr string) int {
	return utf32NormalizedIndex(s, substr)
}

// UTF32NormalizedLastIndex returns the last index of a UTF-32 normalized substring within a UTF-32 normalized string.
// It uses NFC normalization form, and returns -1 if UTF-32 conversion or normalization fails.
func UTF32NormalizedLastIndex(s string, substr string) int {
	return utf32NormalizedLastIndex(s, substr)
}

// UTF32NormalizedIndexAll returns the starting indices of all occurrences of subStr in s,
// normalized to NFC form. This represents the human-readable location of the substring in the string.
func UTF32NormalizedIndexAll(s string, substr string) []int {
	return utf32NormalizedIndexAll(s, substr)
}

// FirstAlphaNumericIndex returns the index of the first alphanumeric character in
// the string `s` or -1 if none is found.
func FirstAlphaNumericIndex(s string) int {
	return firstAlphaNumericIndex(s)
}

// LastAlphaNumericIndex returns the index of the last alphanumeric character in the string s or -1 if none is found.
func LastAlphaNumericIndex(s string) int {
	return lastAlphaNumericIndex(s)
}
