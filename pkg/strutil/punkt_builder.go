package strutil

import punkt "github.com/neurosnap/sentences"

// PunktSentenceData tokenizes the internal string value into a slice of sentences using the Punkt tokenizer.
func (sb *StringBuilder) PunktSentenceData() []*punkt.Sentence {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return punktSentenceData(sb.value)
}

// PunktSentences tokenizes the current string value into a slice of sentences using a Punkt sentence tokenizer.
func (sb *StringBuilder) PunktSentences() []string {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return punktSentences(sb.value)
}

// PunktSentenceCount returns the count of sentences identified by the Punkt tokenizer within the StringBuilder.
func (sb *StringBuilder) PunktSentenceCount() int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return len(sb.PunktSentences())
}

// PunktSentencesCleaned processes the StringBuilder's value to extract and clean sentences using the Punkt algorithm.
// Returns a slice of cleaned sentences or nil if processing is not allowed.
func (sb *StringBuilder) PunktSentencesCleaned() []string {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return punktSentencesCleaned(sb.value)
}

// PunktSentenceCountCleaned returns the count of cleaned sentences derived using the Punkt tokenizer logic.
func (sb *StringBuilder) PunktSentenceCountCleaned() int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return len(sb.PunktSentencesCleaned())
}

// PunktAnnotatedTokensData processes the string value of the StringBuilder using Punkt to return annotated tokens.
func (sb *StringBuilder) PunktAnnotatedTokensData() []*punkt.Token {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return punktAnnotatedTokensData(sb.value)
}

// PunktAnnotatedTokens retrieves a slice of annotated tokens derived from the StringBuilder's
// value using a Punkt tokenizer.
func (sb *StringBuilder) PunktAnnotatedTokens() []string {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return punktAnnotatedTokens(sb.value)
}

// PunktAnnotatedTokenCount returns the count of tokens annotated using the Punkt tokenizer in
// the StringBuilder's value.
func (sb *StringBuilder) PunktAnnotatedTokenCount() int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return punktAnnotatedTokenCount(sb.value)
}
