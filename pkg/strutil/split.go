package strutil

// Substrings takes a string as input and returns a slice of substrings split by whitespace.
func Substrings(s string) []string {
	return substrings(s)
}

// Words takes an input string `s` and returns a slice of substrings with leading and
// trailing non-alphanumerics removed.
func Words(s string) []string {
	return words(s)
}

// WordCount returns the number of words in the input string `s`. Words are determined by
// splitting and cleaning the string.
func WordCount(s string) int {
	return wordCount(s)
}

// WordMap takes a string `s` and returns a map where each key is a unique word in `s`, and the value is its frequency.
func WordMap(s string) map[string]int {
	return wordMap(s)
}

// Split divides the input string s into substrings separated by the separator sep and returns a slice of substrings.
func Split(s string, sep string, preserveSep bool) []string {
	return split(s, sep, preserveSep)
}

// SplitAt splits the input string into two substrings at the specified index and returns them as a slice of strings.
// If the index is out of range, it returns a slice with two empty strings.
func SplitAt(s string, index int) []string {
	return splitAt(s, index)
}

// SplitN splits the input string s into at most n substrings separated by sep, optionally preserving the separator.
func SplitN(s string, sep string, n int, preserveSep bool) []string {
	return splitN(s, sep, n, preserveSep)
}

// Lines takes a string input and returns a slice of strings where each string represents a line from the input.
func Lines(s string) []string {
	return lines(s)
}

// LineCount returns the number of lines in the given string. It utilizes an internal helper function for computation.
func LineCount(s string) int {
	return lineCount(s)
}

// Sentences splits the input string `s` into a slice of sentences based on the specified `language`.
// Supported Languages: "en", "ja", "ru", "zh", "he". Default is "en".
func Sentences(s string, language string) []string {
	return sentences(s, language)
}

// SentenceCount returns the number of sentences in the provided string `s` based on the specified `language`.
// Supported Languages: "en", "ja", "ru", "zh", "he". Default is "en".
func SentenceCount(s string, language string) int {
	return sentenceCount(s, language)
}
