package strutil

import punkt "github.com/neurosnap/sentences"

// PunktSentenceData tokenizes the input string into sentences using the Punkt tokenizer and
// returns a slice of sentences.
func PunktSentenceData(s string) []*punkt.Sentence {
	return punktSentenceData(s)
}

// PunktSentences tokenizes the input string into a slice of sentences based on the Punkt sentence tokenizer.
func PunktSentences(s string) []string {
	return punktSentences(s)
}

// PunktSentencesCleaned tokenizes the input string into sentences, cleans up excess
// punctuation, and removes whitespace.
func PunktSentencesCleaned(s string) []string {
	return punktSentencesCleaned(s)
}

// PunktSentenceCountCleaned returns the count of cleaned sentences in the input string using the Punkt tokenizer.
func PunktSentenceCountCleaned(s string) int {
	return punktSentenceCountCleaned(s)
}

// PunktSentenceCount counts and returns the number of sentences in the input string using a Punkt sentence tokenizer.
func PunktSentenceCount(s string) int {
	return punktSentenceCount(s)
}

// PunktAnnotatedTokensData tokenizes the given string and returns a
// slice of annotated tokens using the Punkt algorithm.
func PunktAnnotatedTokensData(s string) []*punkt.Token {
	return punktAnnotatedTokensData(s)
}

// PunktAnnotatedTokens applies the Punkt tokenizer to the input string and returns a slice of annotated token strings.
func PunktAnnotatedTokens(s string) []string {
	return punktAnnotatedTokens(s)
}

// PunktAnnotatedTokenCount returns the count of annotated tokens in the provided string using a Punkt tokenizer.
func PunktAnnotatedTokenCount(s string) int {
	return punktAnnotatedTokenCount(s)
}
