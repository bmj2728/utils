package strutil

// PhoneFormat represents the enumerated type for different phone number formatting styles.
type PhoneFormat int

// PhoneUS represents the US phone number format.
// PhoneInternational represents the international phone number format.
// PhoneDigitsOnly represents the phone number format containing only digits.
const (
	PhoneUS PhoneFormat = iota
	PhoneInternational
	PhoneDigitsOnly
)

// CharacterSet defines a custom type representing different character encoding sets.
type CharacterSet int

// UTF8 represents the UTF-8 character set, which is the default and Go native encoding.
// ASCII represents the ASCII character set, which is safe and widely usable anywhere.
// Latin1 represents the extended ASCII character set, also known as ISO-8859-1.
const (
	UTF8   CharacterSet = iota // Default: UTF8 Go native
	ASCII                      // Safe, usable anywhere
	Latin1                     // Extended ASCII (ISO-8859-1)
)

const (

	// AlphaNumeric represents a string containing all uppercase letters, lowercase letters, and numeric digits (0-9).
	AlphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// HexChars represents the hexadecimal character set containing digits 0-9 and letters a-f.
	HexChars = "0123456789abcdef"

	// URLSafe defines a set of characters considered safe for use in URLs, including alphanumeric characters and "-_".
	URLSafe = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
)

const (

	// LabelRegex defines a regular expression for validating domain labels, ensuring they meet DNS hostname requirements.
	LabelRegex = `[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?`

	// TLDRegex is a regular expression that matches valid top-level domain (TLD) strings consisting of 2 or more alphabetic characters.
	TLDRegex = `[a-zA-Z]{2,}`
)
