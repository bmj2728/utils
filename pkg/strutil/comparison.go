package strutil

// CompareStringSlices compares two slices of strings for equality, with an option
// to treat nil slices as equal if nulls is true.
func CompareStringSlices(a, b []string, nulls bool) bool {
	return compareStringSlices(a, b, nulls)
}

// LevenshteinDistance calculates the Levenshtein distance between two strings s1 and s2.
// It represents the minimum number of edits needed to convert one string into the other.
// An edit is an insertion, deletion, or substitution of a single character.
//
// Additional information: https://en.wikipedia.org/wiki/Levenshtein_distance
func LevenshteinDistance(s1, s2 string) int {
	return levenshteinDistance(s1, s2)
}

// DamerauLevenshteinDistance computes the edit distance between two strings,
// including transpositions of adjacent characters.
//
// It represents the minimum number of operations to change one string to another.
// An operation is an insertion, deletion/substitution of a single character, or transposition of adjacent characters.
// Additional information: https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance
func DamerauLevenshteinDistance(s1, s2 string) int {
	return damerauLevenshteinDistance(s1, s2)
}

// OSADamerauLevenshteinDistance calculates the Damerau-Levenshtein
// distance between two strings for similarity measurement.
//
// This optimal string alignment variant of DamerauLevenshteinDistance
// does not allow multiple transformations on the same substring
func OSADamerauLevenshteinDistance(s1, s2 string) int {
	return osaDamerauLevenshteinDistance(s1, s2)
}

// LCS calculates the length of the longest common subsequence between two input strings s1 and s2.
func LCS(s1, s2 string) int {
	return lcs(s1, s2)
}

// LCSBacktrack calculates the longest common subsequence (LCS) of
// two input strings and returns the result or an error.
func LCSBacktrack(s1, s2 string) (string, error) {
	return lcsBacktrack(s1, s2)
}

// LCSBacktrackAll computes all longest common subsequences between two input strings s1 and s2.
// Returns a slice of subsequences and an error if the computation fails.
func LCSBacktrackAll(s1, s2 string) ([]string, error) {
	return lcsBacktrackAll(s1, s2)
}

func LCSDiff(str1, str2 string) ([]string, error) {
	return lcsDiff(str1, str2)
}

// LCSEditDistance computes the edit distance between two strings using the Longest Common Subsequence (LCS) method.
func LCSEditDistance(s1, s2 string) int {
	return lcsEditDistance(s1, s2)
}

// HammingDistance computes the Hamming distance between two strings s1 and s2, returning an integer
// result and an error if the strings are of unequal length.
//
// The Hamming distance between two equal-length strings of symbols is the number of
// positions at which the corresponding symbols are different.
//
// Additional Info: https://en.wikipedia.org/wiki/Hamming_distance
func HammingDistance(s1, s2 string) (*int, error) {
	return hammingDistance(s1, s2)
}

// JaroSimilarity calculates the Jaro similarity between two strings, returning a float32 value in the range [0, 1].
//
// The higher the value, the more similar the strings are.
// The score is normalized such that 0 equates to no similarities and 1 is an exact match
//
// Additional Info: https://rosettacode.org/wiki/Jaro_similarity
func JaroSimilarity(s1, s2 string) float32 {
	return jaroSimilarity(s1, s2)
}

// JaroWinklerSimilarity computes the Jaro-Winkler similarity between two strings
// and returns a float32 between 0 and 1.
//
// Uses Jaro similarity with a more favorable weighting for similar common prefixes.
//
// Additional Info: https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance#Jaro%E2%80%93Winkler_similarity
func JaroWinklerSimilarity(s1, s2 string) float32 {
	return jaroWinklerSimilarity(s1, s2)
}

// JaccardSimilarity computes the Jaccard similarity coefficient between two strings, using k-grams
// of the given split length.
// For splitLength = 0, the strings are split on whitespaces
//
// The Jaccard index is defined as the size of the intersection divided by the size of the union
// for two given finite, non-empty sets
//
// Additional Info: https://en.wikipedia.org/wiki/Jaccard_index
func JaccardSimilarity(s1, s2 string, splitLength int) *float32 {
	return jaccardSimilarity(s1, s2, splitLength)
}

// CosineSimilarity computes the cosine similarity between two strings using the specified n-gram split length.
// Returns a pointer to the similarity score or nil if the split length is negative.
// If the split length is zero, it splits the strings on whitespaces.
//
// Cosine similarity is the cosine of the angle between the vectors.
//
// Additional Info: https://en.wikipedia.org/wiki/Cosine_similarity/
func CosineSimilarity(s1, s2 string, splitLength int) *float32 {
	return cosineSimilarity(s1, s2, splitLength)
}

// SorensenDiceCoefficient computes the Sørensen–Dice coefficient for two strings using a given n-gram split length.
// Returns a pointer to the coefficient value or nil if the splitLength is negative.
//
// The Sørensen index equals twice the number of elements common to both sets
// divided by the sum of the number of elements in each set.
//
// Additional Info: https://en.wikipedia.org/wiki/Dice-S%C3%B8rensen_coefficient
func SorensenDiceCoefficient(s1, s2 string, splitLength int) *float32 {
	return sorensenDiceCoefficient(s1, s2, splitLength)
}

// QgramDistance calculates the q-gram distance between two strings s1 and s2 using the specified q-gram size q.
func QgramDistance(s1, s2 string, q int) *int {
	return qgramDistance(s1, s2, q)
}
