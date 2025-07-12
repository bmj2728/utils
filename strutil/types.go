package strutil

type PhoneFormat int

const (
	PhoneUS PhoneFormat = iota
	PhoneInternational
	PhoneDigitsOnly
)

// CharacterSet represents an Enumeration of available character sets.
// ASCII will be addressed first and serve as the default value.
// UTF8 will be added in a later version.
// Latin1 will be added in a later version.
// Additional options will be added as the module matures.
//type CharacterSet int
//
//const (
//	ASCII  CharacterSet = iota // Default: ASCII-safe, works everywhere
//	UTF8 // Unicode-preserving
//	Latin1 // Extended ASCII (ISO-8859-1)
//)

const (
	AlphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	HexChars     = "0123456789abcdef"
	UrlSafe      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
)
