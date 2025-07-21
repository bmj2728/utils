package strutil

import (
	"regexp"

	"github.com/hbollon/go-edlib"
)

//// PhoneFormat represents the enumerated type for different phone number formatting styles.
//type PhoneFormat int
//
//// PhoneUS represents the US phone number format.
//// PhoneInternational represents the international phone number format.
//// PhoneDigitsOnly represents the phone number format containing only digits.
//const (
//	PhoneUS PhoneFormat = iota
//	PhoneInternational
//	PhoneDigitsOnly
//)

// CharacterSet defines a custom type representing sets of characters
// fit for various uses.
//
// Options: AlphaNumeric, Alpha, HexChars, URLSafe, WhiteSpace
type CharacterSet string

// CreateCharacterSet initializes and returns a new CharacterSet using the provided string representation.
func createCharacterSet(set string) CharacterSet {
	return CharacterSet(set)
}

const (
	// AlphaNumeric represents a string containing all uppercase letters, lowercase letters, and numeric digits (0-9).
	AlphaNumeric CharacterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// Alpha represents a string containing all uppercase and lowercase letters of the English alphabet.
	Alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// HexChars represents the hexadecimal character set containing digits 0-9 and letters a-f.
	HexChars = "0123456789abcdef"
	// URLSafe defines a set of characters considered safe for use in URLs, including alphanumeric characters and "-_".
	URLSafe = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	// WhiteSpace represents a string containing common whitespace characters: space, tab, newline, carriage return,
	// vertical tab, and form feed.
	WhiteSpace = " \t\n\r\v\f"
)

var (

	// LabelRegex defines a regular expression for validating domain labels, ensuring they meet DNS hostname requirements.
	LabelRegex = `[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?`

	// TLDRegex is a regular expression that matches valid top-level domain (TLD) strings
	// consisting of 2 or more alphabetic characters.
	TLDRegex = `[a-zA-Z]{2,}`

	// CamelCaseRegex is a regular expression used to identify transitions between lowercase
	// and uppercase characters or consecutive capitals in camelCase or PascalCase strings.
	CamelCaseRegex = regexp.MustCompile(`([a-z0-9])([A-Z])|([A-Z])([A-Z][a-z])`)
)

// AlgorithmMap is a mapping of algorithm names as keys to their
// corresponding edlib.Algorithm implementations as values.
var AlgorithmMap = map[string]edlib.Algorithm{
	"Levenshtein":           edlib.Levenshtein,
	"DamerauLevenshtein":    edlib.DamerauLevenshtein,
	"OSADamerauLevenshtein": edlib.OSADamerauLevenshtein,
	"Lcs":                   edlib.Lcs,
	"Hamming":               edlib.Hamming,
	"Jaro":                  edlib.Jaro,
	"JaroWinkler":           edlib.JaroWinkler,
	"Cosine":                edlib.Cosine,
	"Jaccard":               edlib.Jaccard,
	"SorensenDice":          edlib.SorensenDice,
	"QGrams":                edlib.Qgram,
}

// SimilarityTypeMap maps edlib.Algorithm constants to their corresponding string representations for display purposes.
var SimilarityTypeMap = map[string]string{
	"Levenshtein":           "Levenshtein",
	"DamerauLevenshtein":    "Damerau-Levenshtein",
	"OSADamerauLevenshtein": "OSA Damerau-Levenshtein",
	"Lcs":                   "LCS",
	"Hamming":               "Hamming",
	"Jaro":                  "Jaro",
	"JaroWinkler":           "Jaro-Winkler",
	"Cosine":                "Cosine",
	"Jaccard":               "Jaccard",
	"SorensenDice":          "Sorensen-Dice",
	"QGrams":                "Q-Gram",
}

// ShingleResultType is an enumerated type used to represent the type of shingle result, such as map or slice.
type ShingleResultType int

// ShinglesMap represents a result type where shingles are stored in a map.
// ShinglesSlice represents a result type where shingles are stored in a slice.
const (
	ShinglesMap ShingleResultType = iota
	ShinglesSlice
)

// ShingleResultTypeMap maps ShingleResultType constants to their corresponding descriptive string representations.
var ShingleResultTypeMap = map[ShingleResultType]string{
	ShinglesMap:   "Shingle Map",
	ShinglesSlice: "Shingle Slice",
}

// LCSResultType represents an enumerated type used to define different
// result types for Longest Common Subsequence calculations.
type LCSResultType int

// LCSBacktrackWord indicates backtracking for a single LCS word result.
// LCSBacktrackWordAll indicates backtracking for all LCS word results.
// LCSDiffSlice represents an LCS result in the form of a diff slice.
const (
	LCSBacktrackWord LCSResultType = iota
	LCSBacktrackWordAll
	LCSDiffSlice
)

// LCSResultTypeMap maps LCSResultType constants to their corresponding string representations.
var LCSResultTypeMap = map[LCSResultType]string{
	LCSBacktrackWord:    "LCS Backtrack",
	LCSBacktrackWordAll: "LCS Backtrack All",
	LCSDiffSlice:        "LCS Diff",
}
