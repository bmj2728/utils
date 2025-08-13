package strutil

import (
	"strings"
	"sync"
	"unicode"
)

type RuneSet int

// IsAlphaNumericRune determines if the given rune is an alphanumeric character (letter or digit).
func IsAlphaNumericRune(r rune) bool {
	if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
		return false
	}
	return true
}

// Letter represents a set of letter runes.
// Digit represents a set of digit runes.
// Number represents a set of numeric runes.
// AlphaNumericChars represents a set of alphanumeric runes.
// WhiteSpaceChars represents a set of whitespace runes.
// Punctuation represents a set of punctuation runes.
// Symbol represents a set of symbol runes.
// Upper represents a set of uppercase letter runes.
// Lower represents a set of lowercase letter runes.
// Title represents a set of title-case letter runes.
// Control represents a set of control character runes.
// Graphic represents a set of graphic character runes.
// Mark represents a set of mark character runes.
// Printable represents a set of printable character runes.
const (
	Letter RuneSet = iota
	Digit
	Number
	AlphaNumeric
	WhiteSpace
	Punctuation
	Symbol
	Upper
	Lower
	Title
	Control
	Graphic
	Mark
	Printable
)

// RuneSets maps predefined RuneSet constants to their respective rune classification functions.
var RuneSets = map[RuneSet]func(rune) bool{
	Letter:       unicode.IsLetter,
	Digit:        unicode.IsDigit,
	Number:       unicode.IsNumber,
	AlphaNumeric: IsAlphaNumericRune,
	WhiteSpace:   unicode.IsSpace,
	Punctuation:  unicode.IsPunct,
	Symbol:       unicode.IsSymbol,
	Upper:        unicode.IsUpper,
	Lower:        unicode.IsLower,
	Title:        unicode.IsTitle,
	Control:      unicode.IsControl,
	Graphic:      unicode.IsGraphic,
	Mark:         unicode.IsMark,
	Printable:    unicode.IsPrint,
}

// RuneSetNames is a map associating string keys with predefined RuneSet constants for categorizing runes.
var RuneSetNames = map[string]RuneSet{
	"Letter":            Letter,
	"Digit":             Digit,
	"Number":            Number,
	"AlphaNumericChars": AlphaNumeric,
	"WhiteSpaceChars":   WhiteSpace,
	"Punctuation":       Punctuation,
	"Symbol":            Symbol,
	"Upper":             Upper,
	"Lower":             Lower,
	"Title":             Title,
	"Control":           Control,
	"Graphic":           Graphic,
	"Mark":              Mark,
	"Printable":         Printable,
}

// GetRuneSet retrieves a RuneSet from the RuneSetNames map by its name.
// Returns the RuneSet and a boolean indicating if the name was found.
func GetRuneSet(name string) RuneSet {
	set, ok := RuneSetNames[name]
	if !ok {
		return -1
	}
	return set
}

// runeMu is a mutex used to provide safe access to shared resources involving runes.
// customSetIndex represents a rune set with a defined maximum value of 99.
var (
	runeMu         sync.Mutex
	customSetIndex = RuneSet(99)
)

// RegisterRuneSet registers a new rune classification function with a unique name
// and returns its associated RuneSet constant.
func RegisterRuneSet(name string, fn func(rune) bool) RuneSet {
	runeMu.Lock()
	defer runeMu.Unlock()
	customSetIndex++
	RuneSets[customSetIndex] = fn
	RuneSetNames[name] = customSetIndex
	return customSetIndex
}

// CheckRunesBySetName checks if a rune in the string belongs to the rune set identified by the set name.
// Returns a boolean result.
func CheckRunesBySetName(s string, set string) bool {
	runeMu.Lock()
	fn, exists := RuneSets[GetRuneSet(set)]
	runeMu.Unlock()
	if !exists {
		return false
	}
	return strings.ContainsFunc(s, fn)
}

// CheckRunes determines if a string contains any rune that matches the function associated with the specified RuneSet.
func CheckRunes(s string, set RuneSet) bool {
	runeMu.Lock()
	fn, exists := RuneSets[set]
	runeMu.Unlock()
	if !exists {
		return false
	}
	return strings.ContainsFunc(s, fn)
}
