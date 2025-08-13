package strutil

// GenerateUUID generates and returns a new random UUID as a string.
func GenerateUUID() string {
	return makeUUID()
}

// GenerateUUIDV7 generates and returns a new random UUID as a string using UUID V7
// It is recommended to use V7 unless legacy compatibility is required
func GenerateUUIDV7() string {
	return makeUUIDV7()
}

// RandomString generates a random string of the specified length using the provided CharacterSet.
func RandomString(length int, charSet CharacterSet) string {
	return randomFromCharset(length, charSet)
}

// RandomStringFromCustomCharset generates a random string of a given length using a specified custom character set.
func RandomStringFromCustomCharset(length int, customCharset string) string {
	return randomFromCustomCharset(length, customCharset)
}

// RandomAlphaNumericString generates a random alphanumeric string of the
// specified length using the AlphaNumericChars character set.
// This uses math/rand and is NOT cryptographically secure.
// For security-sensitive applications, use crypto/rand directly.
func RandomAlphaNumericString(length int) string {
	return randomAlphaNumericString(length)
}

// RandomAlphaString generates a random string of the specified length containing only alphabetic characters (A-Za-z).
func RandomAlphaString(length int) string {
	return randomAlphaString(length)
}

// RandomHex generates a random hexadecimal string of the specified length.
func RandomHex(length int) string {
	return randomHex(length)
}

// RandomUrlSafe generates a random URL-safe string of the specified length using characters suitable for URLs.
func RandomUrlSafe(length int) string {
	return randomURLSafe(length)
}
