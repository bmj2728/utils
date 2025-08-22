package strutil

import (
	"bufio"
	"strings"
)

// substrings splits the input string `s` by whitespace and returns a slice of substrings.
// Returns an empty slice if the input string is empty.
func substrings(s string) []string {
	return strings.Fields(s)
}

// words extracts and returns all substrings, after stripping leading and trailing non-alphanumerics
// from the input string `s` as a slice of strings.
func words(s string) []string {
	subs := substrings(s)
	var words []string
	for _, sub := range subs {
		fa := firstAlphaNumericIndex(sub)
		la := lastAlphaNumericIndex(sub)
		if fa == -1 || la == -1 {
			continue
		}
		words = append(words, sub[fa:la+1])
	}
	return words
}

// wordCount returns the number of words in the input string `s`.
// Words are determined by splitting and cleaning the string.
func wordCount(s string) int {
	return len(words(s))
}

func wordMap(s string) map[string]int {
	words := words(s)
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	return wordMap
}

// split divides the string s into substrings separated by the sep and returns a slice of substrings.
// If preserveSep is true, the resulting substrings include the separator at the end.
func split(s string, sep string, preserveSep bool) []string {
	if preserveSep {
		return strings.SplitAfter(s, sep)
	}
	return strings.Split(s, sep)
}

// splitN splits a string into at most n substrings separated by sep, optionally preserving the separator in the result.
func splitN(s string, sep string, n int, preserveSep bool) []string {
	if preserveSep {
		return strings.SplitAfterN(s, sep, n)
	}
	return strings.SplitN(s, sep, n)
}

// lines splits the input string s into a slice of strings, each representing a line, and returns the resulting slice.
func lines(s string) []string {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
