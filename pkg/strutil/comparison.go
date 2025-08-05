package strutil

import "utils/pkg/internal/comparison"

// CompareSlices compares two slices of strings for equality, with an option
// to treat nil slices as equal if nulls is set to true.
func CompareSlices(a, b []string, nulls bool) bool {
	return comparison.CompareStringSlices(a, b, nulls)
}

// LevenshteinDistance calculates the Levenshtein distance between two strings and returns a ComparisonResultInt.
// It determines the minimum number of single-character edits (insertions, deletions, or substitutions) required.
func LevenshteinDistance(s1, s2 string) *ComparisonResultInt {
	return levenshteinDistance(s1, s2)
}

// DamerauLevenshteinDistance calculates the Damerau-Levenshtein distance between
// two strings to measure their score.
func DamerauLevenshteinDistance(s1, s2 string) *ComparisonResultInt {
	return damerauLevenshteinDistance(s1, s2)
}

// OSADamerauLevenshteinDistance calculates the optimal string alignment variant of Damerau-Levenshtein distance.
// It returns a ComparisonResultInt containing the computed distance and details for the input strings.
func OSADamerauLevenshteinDistance(s1, s2 string) *ComparisonResultInt {
	return osaDamerauLevenshteinDistance(s1, s2)
}

// LCS calculates the length of the longest common subsequence between two input strings s1 and s2.
func LCS(s1, s2 string) *ComparisonResultInt {
	return lcs(s1, s2)
}

// LCSEditDistance computes the edit distance between two strings using the Longest Common Subsequence (LCS) method.
func LCSEditDistance(s1, s2 string) *ComparisonResultInt {
	return lcsEditDistance(s1, s2)
}

// LCSBacktrack computes the longest common subsequence (LCS) between two input strings and returns an LCSResult.
func LCSBacktrack(s1, s2 string) *LCSResult {
	return lcsBacktrack(s1, s2)
}

// LCSBacktrackAll computes all longest common subsequences of two input strings and returns them as an LCSResult.
func LCSBacktrackAll(s1, s2 string) *LCSResult {
	return lcsBacktrackAll(s1, s2)
}

// LCSDiff computes the Longest Common Subsequence (LCS) difference
// between two strings and returns an LCSResult instance.
func LCSDiff(str1, str2 string) *LCSResult {
	return lcsDiff(str1, str2)
}

// HammingDistance computes the Hamming distance between two strings s1 and s2, returning a ComparisonResultInt
// and an error if the strings are of unequal length.
//
// The Hamming distance between two equal-length strings of symbols is the number of
// positions at which the corresponding symbols are different.
//
// Additional Info: https://en.wikipedia.org/wiki/Hamming_distance
func HammingDistance(s1, s2 string) *ComparisonResultInt {
	return hammingDistance(s1, s2)
}

// JaroSimilarity calculates the Jaro score between two strings, returns a ComparisonResultFloat pointer
//
// The higher the value, the more similar the strings are.
// The score is normalized such that 0 equates to no similarities and 1 is an exact match
//
// Additional Info: https://rosettacode.org/wiki/Jaro_similarity
func JaroSimilarity(s1, s2 string) *ComparisonResultFloat {
	return jaroSimilarity(s1, s2)
}

// JaroWinklerSimilarity computes the Jaro-Winkler score between two strings
// and returns a ComparisonResultFloat pointer.
//
// Uses Jaro score with a more favorable weighting for similar common prefixes.
//
// Additional Info: https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance#Jaro%E2%80%93Winkler_similarity
func JaroWinklerSimilarity(s1, s2 string) *ComparisonResultFloat {
	return jaroWinklerSimilarity(s1, s2)
}

// JaccardSimilarity computes the Jaccard score coefficient between two strings, using k-grams
// of the given split length.
// For splitLength = 0, the strings are split on whitespaces
//
// The Jaccard index is defined as the size of the intersection divided by the size of the union
// for two given finite, non-empty sets
//
// Additional Info: https://en.wikipedia.org/wiki/Jaccard_index
func JaccardSimilarity(s1, s2 string, splitLength int) *ComparisonResultFloat {
	return jaccardSimilarity(s1, s2, splitLength)
}

// CosineSimilarity computes the cosine score between two strings using the specified n-gram split length.
// Returns a pointer to the score score or nil if the split length is negative.
// If the split length is zero, it splits the strings on whitespaces.
//
// Cosine score is the cosine of the angle between the vectors.
//
// Additional Info: https://en.wikipedia.org/wiki/Cosine_similarity/
func CosineSimilarity(s1, s2 string, splitLength int) *ComparisonResultFloat {
	return cosineSimilarity(s1, s2, splitLength)
}

// SorensenDiceCoefficient computes the Sørensen–Dice coefficient for two strings using a given n-gram split length.
// Returns a pointer to the coefficient value or nil if the splitLength is negative.
//
// The Sørensen index equals twice the number of elements common to both sets
// divided by the sum of the number of elements in each set.
//
// Additional Info: https://en.wikipedia.org/wiki/Dice-S%C3%B8rensen_coefficient
func SorensenDiceCoefficient(s1, s2 string, splitLength int) *ComparisonResultFloat {
	return sorensenDiceCoefficient(s1, s2, splitLength)
}

// QgramDistance calculates the q-gram distance between two strings s1 and s2 using the specified q-gram size q.
func QgramDistance(s1, s2 string, q int) *ComparisonResultInt {
	return qgramDistance(s1, s2, q)
}

// QgramDistanceCustomNgram computes the q-gram distance between two n-gram
// frequency maps represented as string-int maps.
// It compares the input frequency maps and returns an integer representing the calculated distance.
func QgramDistanceCustomNgram(nmap1, nmap2 map[string]int, customName string) *ComparisonResultInt {
	return qgramDistanceCustomNgram(nmap1, nmap2, customName)
}

// QgramSimilarity calculates the q-gram score between two strings using the specified q-gram size.
// Returns a pointer to the score score or nil if the q-gram size is less than 1.
func QgramSimilarity(s1, s2 string, q int) *ComparisonResultFloat {
	return qgramSimilarity(s1, s2, q)
}

// Shingle generates k-shingles from the input string and returns a pointer to a map with shingles and their counts.
func Shingle(s string, k int) *ShingleMapResult {
	return shingle(s, k)
}

// ShingleSlice generates k-length shingles (substrings) from the input string and
// returns them as a pointer to a string slice.
// Returns nil if k is less than 1.
func ShingleSlice(s string, k int) *ShingleSliceResult {
	return shingleSlice(s, k)
}

// Similarity computes the score between two input strings using the specified algorithm.
// Returns a SimilarityResult containing the score score or any error encountered during computation.
func Similarity(s1, s2 string, algorithm Algorithm) *SimilarityResult {
	return similarity(s1, s2, algorithm)
}
