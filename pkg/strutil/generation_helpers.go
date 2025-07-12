package strutil

import (
	"math/rand"

	"github.com/google/uuid"
)

// makeUUID generates and returns a new random UUID as a string.
func makeUUID() string {
	return uuid.NewString()
}

// makeUUIDV7 generates and returns a new UUID string using UUID version 7 specification.
// Falls back to generating a standard UUID if UUID V7 generation fails.
// UUID V7 failures result from manipulating the randomness pool
// Since we do not manipulate randomness in the uuid package, this should never fire
func makeUUIDV7() string {
	u, err := uuid.NewV7()
	if err != nil {
		return makeUUID()
	}
	return u.String()
}

// randomFromCharset generates a random string of the specified length using characters from the provided charset.
func randomFromCharset(length int, charset CharacterSet) string {
	if length < 1 {
		return ""
	}
	s := make([]byte, length)
	for i := range s {
		s[i] = charset[rand.Intn(len(charset))]
	}
	return string(s)
}

// randomFromCustomCharset generates a random string of the specified length
// using a custom character set provided as a string.
func randomFromCustomCharset(length int, charset string) string {
	custom := createCharacterSet(charset)
	return randomFromCharset(length, custom)
}

// randomAlphaNumericString generates a random alphanumeric string of the specified length.
func randomAlphaNumericString(length int) string {
	return randomFromCharset(length, AlphaNumericChars)
}

// randomAlphaString generates a random string of the specified length containing only alphabetic characters (A-Za-z).
func randomAlphaString(length int) string {
	return randomFromCharset(length, Alpha)
}

// randomHex generates a random hexadecimal string of the specified length.
func randomHex(length int) string {
	return randomFromCharset(length, HexChars)
}

// randomURLSafe generates a random URL-safe string of the specified length using alphanumeric characters and "-_".
func randomURLSafe(length int) string {
	return randomFromCharset(length, URLSafe)
}
