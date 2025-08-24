package strutil

import (
	"github.com/bmj2728/utils/pkg/internal/types"

	punkt "github.com/neurosnap/sentences"
)

// getPunktTokenizer returns an English sentence tokenizer based on the Punkt algorithm for sentence
// boundary detection.
func getPunktTokenizer() *punkt.DefaultSentenceTokenizer {
	return types.EnglishPunktTokenizer
}

// punktSentenceData tokenizes the input string into a slice of sentences using the Punkt sentence tokenizer.
func punktSentenceData(s string) []*punkt.Sentence {
	data := getPunktTokenizer().Tokenize(s)
	return data
}

// punktSentences takes a string and tokenizes it into a slice of sentences using a Punkt sentence tokenizer.
func punktSentences(s string) []string {
	data := getPunktTokenizer().Tokenize(s)
	sentences := make([]string, len(data))
	for i, sentence := range data {
		sentences[i] = trim(sentence.Text)
	}
	return sentences
}

// punktSentenceCount counts and returns the number of sentences in the provided string using the Punkt tokenizer.
func punktSentenceCount(s string) int {
	return len(punktSentences(s))
}

func punktSentencesCleaned(s string) []string {
	data := getPunktTokenizer().Tokenize(s)
	var sentences []string
	for _, sentence := range data {
		needsClean := PunktCleanupRegex.FindStringIndex(sentence.Text)
		if len(needsClean) > 0 {
			// sentence has more terminators to address
			// split sentence into substrings and append each substring to the slice
			for _, sub := range splitAt(sentence.Text, needsClean[1]-1) {
				sentences = append(sentences, trim(sub))
			}
		} else if len(needsClean) == 0 {
			sentences = append(sentences, trim(sentence.Text))
		}
	}
	return sentences
}

// punktSentenceCountCleaned takes a string input and returns the number of cleaned sentences identified
// by Punkt.
func punktSentenceCountCleaned(s string) int {
	return len(punktSentencesCleaned(s))
}

// punktAnnotatedTokensData processes a string using the Punkt algorithm and returns annotated tokens as a slice.
func punktAnnotatedTokensData(s string) []*punkt.Token {
	data := getPunktTokenizer().AnnotatedTokens(s)
	return data
}

// punktAnnotatedTokens processes a string input and returns a slice of annotated tokens as
// strings using a Punkt tokenizer.
func punktAnnotatedTokens(s string) []string {
	data := getPunktTokenizer().AnnotatedTokens(s)
	tokens := make([]string, len(data))
	for i, token := range data {
		tokens[i] = token.Tok
	}
	return tokens
}

// punktAnnotatedTokenCount returns the number of annotated tokens in the given string using a Punkt tokenizer.
func punktAnnotatedTokenCount(s string) int {
	return len(punktAnnotatedTokens(s))
}
