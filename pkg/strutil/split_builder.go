package strutil

// Substrings generates and returns a slice of substrings from the StringBuilder's value, split by whitespace.
// Returns nil if processing should not continue.
func (sb *StringBuilder) Substrings() []string {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return substrings(sb.value)
}

// Words processes the internal string value and returns a slice of words by removing non-alphanumerics.
// Returns nil if string processing is terminated.
func (sb *StringBuilder) Words() []string {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return words(sb.value)
}

// WordCount returns the number of words in the StringBuilder instance. If processing is halted, it returns -1.
func (sb *StringBuilder) WordCount() int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return len(sb.Words())
}

// WordMap generates a map with words from the StringBuilder's value as keys and their occurrences as values.
// Returns nil if processing should not continue.
func (sb *StringBuilder) WordMap() map[string]int {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return wordMap(sb.value)
}

// Split divides the StringBuilder's value into substrings separated by the specified
// separator and returns a slice of substrings.
func (sb *StringBuilder) Split(sep string, preserveSep bool) []string {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return split(sb.value, sep, preserveSep)
}

// SplitAt splits the StringBuilder's value into two substrings at the specified index and
// returns them as a slice of strings.
// Returns nil if the processing condition is not met.
func (sb *StringBuilder) SplitAt(index int) []string {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return splitAt(sb.value, index)
}

// SplitN splits the StringBuilder value into at most n substrings separated by the specified separator.
// If preserveSep is true, the separator is included in the result.
// Returns nil if the StringBuilder should not continue processing.
func (sb *StringBuilder) SplitN(sep string, n int, preserveSep bool) []string {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return splitN(sb.value, sep, n, preserveSep)
}

func (sb *StringBuilder) Lines() []string {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return lines(sb.value)
}

// LineCount returns the total number of lines stored in the StringBuilder. Returns -1 if
// processing should not continue.
func (sb *StringBuilder) LineCount() int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return len(sb.Lines())
}

// Sentences splits the content of the StringBuilder into sentences based on the specified language and returns them.
// Supported Languages: "en", "ja", "ru", "zh", "he". Default is "en".
func (sb *StringBuilder) Sentences(lang string) []string {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return sentences(sb.value, lang)
}

// SentenceCount returns the number of sentences in the StringBuilder. Returns -1 if processing should not continue.
// Supported Languages: "en", "ja", "ru", "zh", "he". Default is "en".
func (sb *StringBuilder) SentenceCount(lang string) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return len(sb.Sentences(lang))
}
