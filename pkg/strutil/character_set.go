package strutil

// CharacterSet defines a custom type representing sets of characters
// fit for various uses.
//
// Options: AlphaNumericChars, Alpha, HexChars, URLSafe, WhiteSpaceChars
type CharacterSet string

// CreateCharacterSet initializes and returns a new CharacterSet using the provided string representation.
func createCharacterSet(set string) CharacterSet {
	return CharacterSet(set)
}

const (
	// AlphaNumericChars represents a string containing all uppercase letters, lowercase letters, and numeric digits (0-9).
	AlphaNumericChars CharacterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// Alpha represents a string containing all uppercase and lowercase letters of the English alphabet.
	Alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// HexChars represents the hexadecimal character set containing digits 0-9 and letters a-f.
	HexChars = "0123456789abcdef"
	// URLSafe defines a set of characters considered safe for use in URLs, including alphanumeric characters and "-_".
	URLSafe = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	// WhiteSpaceChars represents a string containing common whitespace characters: space, tab, newline, carriage return,
	// vertical tab, and form feed.
	WhiteSpaceChars = " \t\n\r\v\f"
)
