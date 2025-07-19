package strutil

// NewLoremWord creates a new StringBuilder instance with a randomly generated word as its initial value.
func NewLoremWord() *StringBuilder {
	return &StringBuilder{
		value: loremWord(),
	}
}

// NewLoremWords creates a new StringBuilder initialized with a string
// containing the specified number of lorem ipsum words.
func NewLoremWords(count int) *StringBuilder {
	return &StringBuilder{
		value: loremWords(count),
	}
}

// NewLoremSentence creates and returns a new StringBuilder initialized with a randomly generated lorem ipsum sentence.
func NewLoremSentence() *StringBuilder {
	return &StringBuilder{
		value: loremSentence(),
	}
}

// NewLoremSentenceCustom creates a new StringBuilder instance containing
// a lorem ipsum sentence with the specified word count.
func NewLoremSentenceCustom(length int) *StringBuilder {
	return &StringBuilder{
		value: loremSentenceCustom(length),
	}
}

// NewLoremSentences creates a new StringBuilder containing the given number of lorem ipsum sentences.
func NewLoremSentences(count int) *StringBuilder {
	return &StringBuilder{
		value: loremSentences(count),
	}
}

// NewLoremSentencesCustom creates a new StringBuilder instance
// containing lorem ipsum sentences based on the given count and length.
func NewLoremSentencesCustom(count int, length int) *StringBuilder {
	return &StringBuilder{
		value: loremSentencesCustom(count, length),
	}
}

// NewLoremSentencesVariable generates a string of lorem ipsum sentences
// with count, min, and max controlling quantity and length.
func NewLoremSentencesVariable(count int, min int, max int) *StringBuilder {
	return &StringBuilder{
		value: loremSentencesVariable(count, min, max),
	}
}

// NewLoremParagraph creates a new StringBuilder initialized with a random Lorem Ipsum paragraph.
func NewLoremParagraph() *StringBuilder {
	return &StringBuilder{
		value: loremParagraph(),
	}
}

// NewLoremParagraphs generates a StringBuilder containing the specified number of lorem ipsum paragraphs.
func NewLoremParagraphs(count int) *StringBuilder {
	return &StringBuilder{
		value: loremParagraphs(count),
	}
}

// NewLoremDomain creates and returns a new StringBuilder initialized
// with a domain string from the loremDomain function.
func NewLoremDomain() *StringBuilder {
	return &StringBuilder{
		value: loremDomain(),
	}
}

// NewLoremURL creates and returns a new StringBuilder initialized with a lorem ipsum URL string.
func NewLoremURL() *StringBuilder {
	return &StringBuilder{
		value: loremURL(),
	}
}

// NewLoremEmail initializes a StringBuilder with a generated mock email value and returns the instance.
func NewLoremEmail() *StringBuilder {
	return &StringBuilder{
		value: loremEmail(),
	}
}
