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
