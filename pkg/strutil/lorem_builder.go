package strutil

// NewLoremWord creates a new StringBuilder instance with a randomly generated word as its initial value.
func NewLoremWord() *StringBuilder {
	s := loremWord()
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewLoremWords creates a new StringBuilder initialized with a string
// containing the specified number of lorem ipsum words.
func NewLoremWords(count int) *StringBuilder {
	s := loremWords(count)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewLoremSentence creates and returns a new StringBuilder initialized with a randomly generated lorem ipsum sentence.
func NewLoremSentence() *StringBuilder {
	s := loremSentence()
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewLoremSentenceCustom creates a new StringBuilder instance containing
// a lorem ipsum sentence with the specified word count.
func NewLoremSentenceCustom(length int) *StringBuilder {
	s := loremSentenceCustom(length)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewLoremSentences creates a new StringBuilder containing the given number of lorem ipsum sentences.
func NewLoremSentences(count int) *StringBuilder {
	s := loremSentences(count)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewLoremSentencesCustom creates a new StringBuilder instance
// containing lorem ipsum sentences based on the given count and length.
func NewLoremSentencesCustom(count int, length int) *StringBuilder {
	s := loremSentencesCustom(count, length)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewLoremSentencesVariable generates a string of lorem ipsum sentences
// with count, min, and max controlling quantity and length.
func NewLoremSentencesVariable(count int, min int, max int) *StringBuilder {
	s := loremSentencesVariable(count, min, max)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewLoremParagraph creates a new StringBuilder initialized with a random Lorem Ipsum paragraph.
func NewLoremParagraph() *StringBuilder {
	s := loremParagraph()
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewLoremParagraphs generates a StringBuilder containing the specified number of lorem ipsum paragraphs.
func NewLoremParagraphs(count int) *StringBuilder {
	s := loremParagraphs(count)
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewLoremDomain creates and returns a new StringBuilder initialized
// with a domain string from the loremDomain function.
func NewLoremDomain() *StringBuilder {
	s := loremDomain()
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewLoremURL creates and returns a new StringBuilder initialized with a lorem ipsum URL string.
func NewLoremURL() *StringBuilder {
	s := loremURL()
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}

// NewLoremEmail initializes a StringBuilder with a generated mock email value and returns the instance.
func NewLoremEmail() *StringBuilder {
	s := loremEmail()
	return &StringBuilder{
		value:         s,
		originalValue: s,
	}
}
