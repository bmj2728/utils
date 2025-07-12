package strutil

// New creates and returns a new StringBuilder instance initialized with the provided string.
func New(s string) *StringBuilder {
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewRandom generates a new StringBuilder containing a random string of the specified length
// using the given CharacterSet.
func NewRandom(length int, charSet CharacterSet) *StringBuilder {
	s := randomFromCharset(length, charSet)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewRandomFromCustomCharSet generates a random string of the specified length using a custom character set.
func NewRandomFromCustomCharSet(length int, charSet string) *StringBuilder {
	s := randomFromCustomCharset(length, charSet)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewRandomAlphaNumeric creates a new StringBuilder with a random alphanumeric string of the specified length.
func NewRandomAlphaNumeric(length int) *StringBuilder {
	s := randomAlphaNumericString(length)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewRandomAlpha creates a StringBuilder with a random alphabetic string (A-Za-z) of the specified length.
func NewRandomAlpha(length int) *StringBuilder {
	s := randomAlphaString(length)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewRandomHex generates a new StringBuilder containing a random hexadecimal string of the specified length.
func NewRandomHex(length int) *StringBuilder {
	s := randomHex(length)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewRandomURLSafe generates a StringBuilder initialized with a random URL-safe string of the given length.
func NewRandomURLSafe(length int) *StringBuilder {
	s := randomURLSafe(length)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewUUID creates and returns a new StringBuilder instance with a generated UUID value.
func NewUUID() *StringBuilder {
	s := makeUUID()
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewUUIDV7 generates a new UUID version 7 and returns it wrapped in a StringBuilder instance.
func NewUUIDV7() *StringBuilder {
	s := makeUUIDV7()
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}
