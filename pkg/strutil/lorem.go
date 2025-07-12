package strutil

// LoremWord generates and returns a random lorem ipsum word as a string.
func LoremWord() string {
	return loremWord()
}

// LoremWords generates a string containing the specified number of lorem ipsum words.
func LoremWords(count int) string {
	return loremWords(count)
}

// LoremSentence generates and returns a placeholder sentence of 8 words using lorem ipsum text.
func LoremSentence() string {
	return loremSentence()
}

// LoremSentenceCustom generates a lorem ipsum sentence with the specified word count. Returns the generated string.
func LoremSentenceCustom(count int) string {
	return loremSentenceCustom(count)
}

// LoremSentences generates a string containing the specified number of 8 word lorem ipsum sentences.
func LoremSentences(count int) string {
	return loremSentences(count)
}

// LoremSentencesCustom generates multiple lorem ipsum sentences with
// the specified sentence count and word length per sentence.
func LoremSentencesCustom(count int, length int) string {
	return loremSentencesCustom(count, length)
}

// LoremSentencesVariable generates variable length lorem sentences with lengths between specified min and max values.
// The parameter 'count' specifies the number of sentences to generate.
func LoremSentencesVariable(count, min, max int) string {
	return loremSentencesVariable(count, min, max)
}

// LoremParagraph generates and returns a string containing a randomly generated Lorem Ipsum paragraph of 45 words.
func LoremParagraph() string {
	return loremParagraph()
}

// LoremParagraphs generates and returns a specified number of lorem ipsum paragraphs as a single string.
// The parameter 'count' specifies the number of paragraphs to generate.
func LoremParagraphs(count int) string {
	return loremParagraphs(count)
}

// LoremDomain generates and returns a placeholder domain name in string format.
func LoremDomain() string {
	return loremDomain()
}

// LoremURL generates and returns a string representing a placeholder or mock URL,
// intended for testing or default usage.
func LoremURL() string {
	return loremURL()
}

// LoremEmail generates and returns a placeholder or mock email address as a string.
func LoremEmail() string {
	return loremEmail()
}
