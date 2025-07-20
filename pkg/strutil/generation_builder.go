package strutil

// New creates and returns a new StringBuilder instance initialized with the provided string.
func New(s string) *StringBuilder {
	return &StringBuilder{
		value:          s,
		comparisonData: NewComparisonData(),
		similarities:   &[]SimilarityResult{},
	}
}

// NewRandom generates a new StringBuilder containing a random string of the specified length
// using the given CharacterSet.
func NewRandom(length int, charSet CharacterSet) *StringBuilder {
	return &StringBuilder{
		value:          randomFromCharset(length, charSet),
		comparisonData: NewComparisonData(),
		similarities:   &[]SimilarityResult{},
	}
}

// NewRandomFromCustomCharSet generates a random string of the specified length using a custom character set.
func NewRandomFromCustomCharSet(length int, charSet string) *StringBuilder {
	return &StringBuilder{
		value:          randomFromCustomCharset(length, charSet),
		comparisonData: NewComparisonData(),
		similarities:   &[]SimilarityResult{},
	}
}

// NewRandomAlphaNumeric creates a new StringBuilder with a random alphanumeric string of the specified length.
func NewRandomAlphaNumeric(length int) *StringBuilder {
	return &StringBuilder{
		value:          randomAlphaNumericString(length),
		comparisonData: NewComparisonData(),
		similarities:   &[]SimilarityResult{},
	}
}

// NewRandomAlpha creates a StringBuilder with a random alphabetic string (A-Za-z) of the specified length.
func NewRandomAlpha(length int) *StringBuilder {
	return &StringBuilder{
		value:          randomAlphaString(length),
		comparisonData: NewComparisonData(),
		similarities:   &[]SimilarityResult{},
	}
}

// NewRandomHex generates a new StringBuilder containing a random hexadecimal string of the specified length.
func NewRandomHex(length int) *StringBuilder {
	return &StringBuilder{
		value:          randomHex(length),
		comparisonData: NewComparisonData(),
		similarities:   &[]SimilarityResult{},
	}
}

// NewRandomURLSafe generates a StringBuilder initialized with a random URL-safe string of the given length.
func NewRandomURLSafe(length int) *StringBuilder {
	return &StringBuilder{
		value:          randomURLSafe(length),
		comparisonData: NewComparisonData(),
		similarities:   &[]SimilarityResult{},
	}
}

// NewUUID creates and returns a new StringBuilder instance with a generated UUID value.
func NewUUID() *StringBuilder {
	return &StringBuilder{
		value:          makeUUID(),
		comparisonData: NewComparisonData(),
		similarities:   &[]SimilarityResult{},
	}
}

// NewUUIDV7 generates a new UUID version 7 and returns it wrapped in a StringBuilder instance.
func NewUUIDV7() *StringBuilder {
	return &StringBuilder{
		value:          makeUUIDV7(),
		comparisonData: NewComparisonData(),
		similarities:   &[]SimilarityResult{},
	}
}
