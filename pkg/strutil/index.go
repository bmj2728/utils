package strutil

import (
	"fmt"
	"unicode/utf8"
)

// TODO implement new struct across info files, create map, implement on builder

// IndexType is an enumerated type used to define various indexing strategies.
type IndexType int

// Byte represents the standard byte index as returned by the standard library regexp string index.
// ByteNormalized represents a byte index after NFC normalization.
// Rune represents an index based on rune positions in a string.
// UTF32Normalized represents the 'absolute' index by normalizing all code points to UTF-32 after NFC normalization.
// This is useful for calculating the human-readable location of a substring within a string.
const (
	Byte IndexType = iota
	ByteNorm
	Rune
	UTF32Norm
)

func (i IndexType) String() string {
	return IndexTypeMap[i]
}

var IndexTypeMap = map[IndexType]string{
	Byte:      "Byte",
	ByteNorm:  "ByteNorm",
	Rune:      "Rune",
	UTF32Norm: "UTF32Norm",
}

// StringIndex represents a structure holding information about a string, its substring,
// an index type, and an index value.
type StringIndex struct {
	str       string
	subStr    string
	indexType IndexType
	index     int
}

// NewStringIndex creates and returns a new StringIndex instance initialized with the given string, substring,
// index type, and index.
func NewStringIndex(str string, subStr string, indexType IndexType, idx int) *StringIndex {
	return &StringIndex{
		str:       str,
		subStr:    subStr,
		indexType: indexType,
		index:     idx,
	}
}

// NewStringRuneIndex creates a new StringIndex instance with the given string, rune, index type, and index values.
func NewStringRuneIndex(str string, r rune, idx int) *StringIndex {
	var subStr string
	if r > 0 {
		subStr = string(r)
	}
	return &StringIndex{
		str:       str,
		subStr:    subStr,
		indexType: Rune,
		index:     idx,
	}
}

// GetString returns the string value stored in the StringIndex instance.
func (s *StringIndex) GetString() string {
	return s.str
}

// GetRune returns the first rune of the substring if the index type is Rune and
// the substring length is 1. Returns 0 otherwise.
func (s *StringIndex) GetRune() rune {
	if utf8.RuneCountInString(s.subStr) != 1 {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(s.subStr)
	return r
}

// GetSubstr returns the substring associated with the StringIndex instance. If the index type is Rune,
// the substring will be a string representation of the rune. GetRune() should be used instead.
func (s *StringIndex) GetSubstr() string {
	return s.subStr
}

// GetIndexType retrieves the index type associated with the StringIndex instance.
func (s *StringIndex) GetIndexType() IndexType {
	return s.indexType
}

// GetIndex returns the index field of the StringIndex structure.
func (s *StringIndex) GetIndex() int {
	return s.index
}

// IsMatch compares the current StringIndex with another for equality based on string, index, and optionally indexType.
func (s *StringIndex) IsMatch(other *StringIndex, matchType bool) bool {
	if s == nil || other == nil {
		return false
	}
	if s.str != other.str {
		return false
	}
	if s.index != other.index {
		return false
	}
	if matchType && s.indexType != other.indexType {
		return false
	}
	return true
}

// Print outputs the StringIndex details to standard output based on the verbosity flag.
func (s *StringIndex) Print(v bool) {
	fmt.Print(formatStringIndexOutput(s, v))
}

// formatStringIndexOutput generates a formatted string representation of a StringIndex object based on verbosity.
func formatStringIndexOutput(si *StringIndex, v bool) string {
	if si == nil {
		return ""
	}
	if v {
		return fmt.Sprintf("String: %s\nSubstring: %s\nIndexType: %s\nIndex: %d\n",
			si.str, si.subStr, si.indexType, si.index)
	}
	return fmt.Sprintf("Index Type: %s, Index: %d\n", si.indexType, si.index)
}
