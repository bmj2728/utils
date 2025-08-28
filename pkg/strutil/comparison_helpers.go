package strutil

import (
	"slices"

	errors2 "github.com/bmj2728/utils/pkg/internal/errors"

	"github.com/hbollon/go-edlib"
)

// levenshteinDistance computes the Levenshtein distance between two strings and returns a ComparisonResult.
// It calculates the minimum number of edits (insertions, deletions, substitutions)
// required to transform one string into another.
func levenshteinDistance(s1, s2 string) int {
	ld := edlib.LevenshteinDistance(s1, s2)
	return ld
}

// damerauLevenshteinDistance computes the Damerau-Levenshtein distance between
// two strings and returns a comparison result.
// This distance measures the minimum number of operations required to transform one string into the other.
// Supported operations include insertion, deletion, substitution, and transposition of adjacent characters.
func damerauLevenshteinDistance(str1, str2 string) int {
	dld := edlib.DamerauLevenshteinDistance(str1, str2)
	return dld
}

// osaDamerauLevenshteinDistance computes the optimal string alignment Damerau-Levenshtein distance between two strings.
// It returns a pointer to a ComparisonResult containing the result of the comparison between `str1` and `str2`.
// The returned structure includes the calculated distance, input strings, and any
// potential error encountered during computation.
func osaDamerauLevenshteinDistance(str1, str2 string) int {
	osaDLD := edlib.OSADamerauLevenshteinDistance(str1, str2)
	return osaDLD
}

// lcs returns the length of the longest common subsequence between two input strings, string1 and string2.
func lcs(str1 string, str2 string) int {
	l := edlib.LCS(str1, str2)
	return l
}

// lcsEditDistance computes the edit distance between two strings based on their Longest Common Subsequence (LCS).
func lcsEditDistance(s1, s2 string) int {
	l := edlib.LCSEditDistance(s1, s2)
	return l
}

// lcsBacktrack computes the longest common subsequence (LCS) between two input strings using backtracking.
// It returns an LCSResult containing the LCS result or an error if the computation fails.
func lcsBacktrack(str1 string, str2 string) (string, error) {
	result, err := edlib.LCSBacktrack(str1, str2)
	return result, err
}

// lcsBacktrackAll computes all possible longest common subsequences between two strings and returns an LCSResult.
// It utilizes the edlib.LCSBacktrackAll function for computation and wraps the result into an LCSResult instance.
// If an error occurs during computation, it is captured and included in the returned LCSResult.
func lcsBacktrackAll(str1 string, str2 string) ([]string, error) {
	result, err := edlib.LCSBacktrackAll(str1, str2)
	return result, err
}

// lcsDiff calculates the LCS difference between two strings and returns a
// pointer to an LCSResult containing the result.
// Returns nil for the result if an error occurs during the calculation.
func lcsDiff(str1, str2 string) ([]string, error) {
	result, err := edlib.LCSDiff(str1, str2)
	return result, err
}

// compareStringBuilderSlices compares two slices of StringBuilder for equality,
// optionally considering nil slices as equal. The comparisonData ignores the order of elements
// and uses the 'nulls' flag to determine nil-handling behavior.
func compareStringBuilderSlices(s1, s2 []StringBuilder, nulls bool) bool {
	if nulls && s1 == nil && s2 == nil {
		return true
	}
	if s1 == nil || s2 == nil {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	remaining := slices.Clone(s2)
	for _, s := range s1 {
		if !slices.Contains(s2, s) {
			return false
		} else {
			for i, w := range remaining {
				if s == w {
					remaining = slices.Delete(remaining, i, i+1)
					break
				}
			}
		}
	}
	return true
}

// hammingDistance computes the Hamming distance between two strings and returns a ComparisonResult instance.
// Returns an error if the strings are not of equal length or if the distance calculation fails.
func hammingDistance(s1, s2 string) (int, error) {
	dist, err := edlib.HammingDistance(s1, s2)
	if err != nil {
		return -1, err
	}
	return dist, err
}

// jaroSimilarity computes the Jaro similarity between two strings
// and returns the result as a ComparisonResultFloat object.
func jaroSimilarity(s1, s2 string) float32 {
	js := edlib.JaroSimilarity(s1, s2)
	return js
}

// jaroWinklerSimilarity computes the Jaro-Winkler similarity between two input strings.
// Returns a ComparisonResultFloat containing the similarity score and input data.
func jaroWinklerSimilarity(s1, s2 string) float32 {
	js := edlib.JaroWinklerSimilarity(s1, s2)
	return js
}

// jaccardSimilarity calculates the Jaccard similarity coefficient between two strings
// using k-grams of a specified length.
// Returns a ComparisonResultFloat object containing the result or an error if the splitLength is invalid (< 0).
func jaccardSimilarity(s1, s2 string, splitLength int) float32 {
	if splitLength < 0 {
		splitLength = 0
	}
	js := edlib.JaccardSimilarity(s1, s2, splitLength)
	return js
}

// cosineSimilarity calculates the cosine similarity between two strings using n-gram splitting with the given length.
// Returns a pointer to a ComparisonResultFloat containing the similarity score or an error
// if the split length is invalid.
func cosineSimilarity(s1, s2 string, splitLength int) float32 {
	if splitLength < 0 {
		splitLength = 0
	}
	cs := edlib.CosineSimilarity(s1, s2, splitLength)
	return cs
}

// sorensenDiceCoefficient calculates the Sørensen–Dice coefficient of two strings based on a specified n-gram length.
// Returns a pointer to a ComparisonResultFloat, containing the coefficient
// value or an error if the splitLength is invalid.
func sorensenDiceCoefficient(s1, s2 string, splitLength int) float32 {
	if splitLength < 1 {
		splitLength = 1
	}
	sdc := edlib.SorensenDiceCoefficient(s1, s2, splitLength)
	return sdc
}

// qgramDistance computes the Q-Gram distance between two strings s1 and s2 using the specified q-gram size q.
// Returns a ComparisonResult containing the computed distance or an error if q is invalid.
func qgramDistance(s1, s2 string, q int) int {
	if q < 0 {
		q = 0
	}
	qd := edlib.QgramDistance(s1, s2, q)
	return qd
}

// qgramDistanceCustomNgram calculates the Q-Gram distance between two n-gram maps and returns a comparison result.
func qgramDistanceCustomNgram(nmap1, nmap2 map[string]int) int {
	qdc := edlib.QgramDistanceCustomNgram(nmap1, nmap2)
	return qdc
}

// qgramSimilarity calculates the q-gram similarity between two strings with a specified q size.
// Returns a ComparisonResultFloat containing the similarity score or an error if q is invalid.
func qgramSimilarity(s1, s2 string, q int) float32 {
	if q < 1 {
		q = 1
	}
	qs := edlib.QgramSimilarity(s1, s2, q)
	return qs
}

// shingle generates k-shingles for the given string and returns them
// as a pointer to a map of shingles and their counts.
// Takes `s` as the input string and `k` as the size of each shingle.
// Returns nil if `k` is less than 1.
func shingle(s string, k int) *ShingleMapResult {
	if k < 1 {
		return NewShingleMapResult(ShinglesMap, s, k, nil, errors2.ErrShingleLengthOutOfRange)
	}
	shingle := edlib.Shingle(s, k)
	smr := NewShingleMapResult(ShinglesMap, s, k, shingle, nil)
	return smr
}

// shingleSlice generates k-length shingles (substrings) from the input string and
// returns them as a pointer to a string slice.
// Returns nil if k is less than 1.
func shingleSlice(s string, k int) *ShingleSliceResult {
	if k < 1 {
		return NewShingleSliceResult(ShinglesSlice, s, k, nil, errors2.ErrShingleLengthOutOfRange)
	}
	shingle := edlib.ShingleSlice(s, k)
	slr := NewShingleSliceResult(ShinglesSlice, s, k, &shingle, nil)
	return slr
}

// similarity calculates the similarity between two strings using
// the specified algorithm and returns a SimilarityResult.
func similarity(s1, s2 string, algorithm Algorithm) *SimilarityResult {
	sim, err := edlib.StringsSimilarity(s1, s2, edlib.Algorithm(algorithm))
	if err != nil {
		// underlying returns 0.0, errors.New("Illegal argument for algorithm method")
		// the error case occurs when Algorithm is invalid
		return NewSimilarityResult(algorithm, s1, s2, nil, err)
	}
	return NewSimilarityResult(algorithm, s1, s2, &sim, err)
}
