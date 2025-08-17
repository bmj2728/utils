package strutil

import (
	"regexp"

	"github.com/bmj2728/utils/pkg/internal/convert"
)

// index returns the index for the first occurrence of `substr` in the string `s`.
// If normalized is true, the string is first normalized to NFC.
// This represents the byte index.
func index(s string, substr string, normalized bool) int {
	if isEmpty(s) || isEmpty(substr) {
		return -1
	}
	if normalized {
		subPat := regexp.MustCompile(normalizeUnicode(substr, NFC))
		idx := subPat.FindStringIndex(normalizeUnicode(s, NFC))
		if idx == nil {
			return -1
		}
		return idx[0]
	} else {
		subPat := regexp.MustCompile(substr)
		idx := subPat.FindStringIndex(s)
		if idx == nil {
			return -1
		}
		return idx[0]
	}
}

// indexAll returns a slice of starting indices of all non-overlapping occurrences of substr in s.
// If normalized is true, both s and substr are normalized to NFC before searching.
func indexAll(s string, substr string, normalized bool) []int {
	var indices []int
	if isEmpty(s) || isEmpty(substr) {
		return indices
	}
	if normalized {
		subPat := regexp.MustCompile(normalizeUnicode(substr, NFC))
		res := subPat.FindAllStringIndex(normalizeUnicode(s, NFC), -1)
		for _, idx := range res {
			indices = append(indices, idx[0])
		}
		return indices
	} else {
		subPat := regexp.MustCompile(substr)
		res := subPat.FindAllStringIndex(s, -1)
		for _, idx := range res {
			indices = append(indices, idx[0])
		}
		return indices
	}
}

// indexRune returns the index of the first occurrence of the rune r in the string s, or -1 if r is not present in s.
func indexRune(s string, r rune) int {
	idx := -1
	if isEmpty(s) {
		return idx
	}
	for i, c := range s {
		if c == r {
			idx = i
			break
		}
	}
	return idx
}

// indexAllRune returns a slice of all indices where the specified rune `r` appears in the string `s`.
func indexAllRune(s string, r rune) []int {
	var indices []int
	if isEmpty(s) {
		return indices
	}
	for i, c := range s {
		if c == r {
			indices = append(indices, i)
		}
	}
	return indices
}

// lastIndex finds the last occurrence of substr in the string s.
// If normalized is true, both s and substr are Unicode normalized before searching.
// Returns -1 if substr is not found.
func lastIndex(s string, substr string, normalized bool) int {
	idx := -1
	if isEmpty(s) || isEmpty(substr) {
		return idx
	}
	if normalized {
		subPat := regexp.MustCompile(normalizeUnicode(substr, NFC))
		res := subPat.FindAllStringIndex(normalizeUnicode(s, NFC), -1)
		for _, i := range res {
			if i[0] > idx {
				idx = i[0]
			}
		}
		return idx
	} else {
		subPat := regexp.MustCompile(substr)
		res := subPat.FindAllStringIndex(s, -1)
		for _, i := range res {
			if i[0] > idx {
				idx = i[0]
			}
		}
		return idx
	}
}

// lastIndexRune returns the index of the last occurrence of the rune r in the string s, or -1 if r is not present in s.
func lastIndexRune(s string, r rune) int {
	idx := -1
	if isEmpty(s) {
		return idx
	}
	for i, c := range s {
		if c == r {
			idx = i
		}
	}
	return idx
}

// utf32NormalizedIndex finds the normalized index of substr in s using UTF-32 encoding;
// this represents the human readable location of the substring in the string.
// returns -1 if not found or errors occur.
func utf32NormalizedIndex(s string, substr string) int {
	if isEmpty(s) || isEmpty(substr) {
		return -1
	}
	s, err := convert.ToUTF32(normalizeUnicode(s, NFC))
	if err != nil {
		return -1
	}
	substr, err = convert.ToUTF32(normalizeUnicode(substr, NFC))
	if err != nil {
		return -1
	}
	return index(s, substr, false) / 4
}

// utf32NormalizedLastIndex returns the last index of a UTF-32 normalized substring within a UTF-32 normalized string.
// Normalization uses NFC form. Returns -1 if normalization or UTF-32 conversion fails.
func utf32NormalizedLastIndex(s string, substr string) int {
	if isEmpty(s) || isEmpty(substr) {
		return -1
	}
	s, err := convert.ToUTF32(normalizeUnicode(s, NFC))
	if err != nil {
		return -1
	}
	substr, err = convert.ToUTF32(normalizeUnicode(substr, NFC))
	if err != nil {
		return -1
	}
	return lastIndex(s, substr, false) / 4
}
