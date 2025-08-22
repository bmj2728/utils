package strutil

import (
	"regexp"
	"strings"

	"github.com/bmj2728/utils/pkg/internal/convert"
	"github.com/bmj2728/utils/pkg/internal/errors"
)

// index returns the index for the first occurrence of `subStr` in the string `s`.
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

// indexAll returns a slice of starting indices of all non-overlapping occurrences of subStr in s.
// If normalized is true, both s and subStr are normalized to NFC before searching.
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
	if isEmpty(s) || r < 1 {
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
	if isEmpty(s) || r < 1 {
		return indices
	}
	for i, c := range s {
		if c == r {
			indices = append(indices, i)
		}
	}
	return indices
}

// lastIndex finds the last occurrence of subStr in the string s.
// If normalized is true, both s and subStr are Unicode normalized before searching.
// Returns -1 if subStr is not found.
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
	if isEmpty(s) || r < 1 {
		return idx
	}
	for i, c := range s {
		if c == r {
			idx = i
		}
	}
	return idx
}

// convertToUTF32 is a helper function that converts the input string to UTF-32 encoding and normalizes it to NFC form.
// It returns the converted strings and any errors that occur during conversion or normalization.
// If either the input string or the substring is empty, the function returns an error.
func convertToUTF32(s string, substr string) (sUTF32 string, subUTF32 string, err error) {
	if isEmpty(s) || isEmpty(substr) {
		return s, substr, errors.ErrInvalidEmpty
	}
	s, err = convert.ToUTF32(normalizeUnicode(s, NFC))
	if err != nil {
		return s, substr, err
	}
	substr, err = convert.ToUTF32(normalizeUnicode(substr, NFC))
	if err != nil {
		return s, substr, err
	}
	return s, substr, nil
}

// utf32NormalizedIndex finds the normalized index of subStr in s using UTF-32 encoding;
// this represents the human-readable location of the substring in the string.
// returns -1 if not found or errors occur.
func utf32NormalizedIndex(s string, substr string) int {
	s, substr, err := convertToUTF32(s, substr)
	if err != nil {
		return -1
	}
	return strings.Index(s, substr) / 4
}

// utf32NormalizedLastIndex returns the last index of a UTF-32 normalized substring within a UTF-32 normalized string.
// Normalization uses NFC form. Returns -1 if normalization or UTF-32 conversion fails.
func utf32NormalizedLastIndex(s string, substr string) int {
	s, substr, err := convertToUTF32(s, substr)
	if err != nil {
		return -1
	}
	return strings.LastIndex(s, substr) / 4
}

// utf32NormalizedIndexAll returns a slice of starting indices of all occurrences of a
// normalized substring in a string.
// The function normalizes both the input string and the substring to
// the NFC Unicode normalization form before processing.
// Indices in the returned slice correspond to character positions in the UTF-32 representation of the input string.
// Returns nil if either the input string or the substring is empty, or if
// any error occurs during normalization or conversion.
func utf32NormalizedIndexAll(s string, substr string) []int {
	var indices []int
	s, substr, err := convertToUTF32(s, substr)
	if err != nil {
		return indices
	}

	// find the first index
	index := strings.Index(s, substr)

	// if it's not -1, add it to the slice and keep going'
	for index > -1 {
		// add the index to the slice, dividing by 4 to account for UTF-32 encoding
		indices = append(indices, index/4)
		// find the next index
		prevIndex := index
		newStr := s[index+len(substr):]
		index = strings.Index(newStr, substr)
		// if it's not -1, adjust it to account for the string manipulation
		if index > -1 {
			index += prevIndex + len(substr)
		}
		// if -1 the loop ends, indicating no more matches
	}

	return indices
}

// firstAlphaNumericIndex returns the index of the first alphanumeric character in the input string `s`.
// If no alphanumeric character is found, it returns -1.
func firstAlphaNumericIndex(s string) int {
	return index(s, "[a-zA-Z0-9]", true)
}

// lastAlphaNumericIndex returns the index of the last alphanumeric character in the string s or -1 if none is found.
func lastAlphaNumericIndex(s string) int {
	return lastIndex(s, "[a-zA-Z0-9]", true)
}
