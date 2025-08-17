package strutil

// IndexType is an enumerated type used to define various indexing strategies.
type IndexType int

// Byte represents the standard byte index as returned by the standard library regexp string index.
// ByteNormalized represents a byte index after NFC normalization.
// Rune represents an index based on rune positions in a string.
// UTF32Normalized represents the 'absolute' index by normalizing all code points to UTF-32 after NFC normalization.
// This is useful for calculating the human-readable location of a substring within a string.
const (
	Byte IndexType = iota
	ByteNormalized
	Rune
	UTF32Normalized
)
