package strutil

// New creates and returns a new StringBuilder instance initialized with the provided string.
func New(s string) *StringBuilder {
	return &StringBuilder{
		value: s,
	}
}

// NewRandom generates a new StringBuilder containing a random string of the specified length
// using the given CharacterSet.
func NewRandom(length int, charSet CharacterSet) *StringBuilder {
	return &StringBuilder{
		value: randomFromCharset(length, charSet),
	}
}

// NewRandomFromCustomCharSet generates a random string of the specified length using a custom character set.
func NewRandomFromCustomCharSet(length int, charSet string) *StringBuilder {
	return &StringBuilder{
		value: randomFromCustomCharset(length, charSet),
	}
}

// NewRandomAlphaNumeric creates a new StringBuilder with a random alphanumeric string of the specified length.
func NewRandomAlphaNumeric(length int) *StringBuilder {
	return &StringBuilder{
		value: randomAlphaNumericString(length),
	}
}

// NewRandomAlpha creates a StringBuilder with a random alphabetic string (A-Za-z) of the specified length.
func NewRandomAlpha(length int) *StringBuilder {
	return &StringBuilder{
		value: randomAlphaString(length),
	}
}

// NewRandomHex generates a new StringBuilder containing a random hexadecimal string of the specified length.
func NewRandomHex(length int) *StringBuilder {
	return &StringBuilder{
		value: randomHex(length),
	}
}

// NewRandomURLSafe generates a StringBuilder initialized with a random URL-safe string of the given length.
func NewRandomURLSafe(length int) *StringBuilder {
	return &StringBuilder{
		value: randomURLSafe(length),
	}
}

// NewUUID creates and returns a new StringBuilder instance with a generated UUID value.
func NewUUID() *StringBuilder {
	return &StringBuilder{
		value: makeUUID(),
	}
}

// NewUUIDV7 generates a new UUID version 7 and returns it wrapped in a StringBuilder instance.
func NewUUIDV7() *StringBuilder {
	return &StringBuilder{
		value: makeUUIDV7(),
	}
}
