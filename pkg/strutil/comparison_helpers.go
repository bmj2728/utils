package strutil

import (
	"errors"
	"slices"

	"github.com/hbollon/go-edlib"
)

// levenshteinDistance computes the Levenshtein distance between two strings and returns a ComparisonResultInt.
// It calculates the minimum number of edits (insertions, deletions, substitutions)
// required to transform one string into another.
func levenshteinDistance(s1, s2 string) *ComparisonResultInt {
	ld := edlib.LevenshteinDistance(s1, s2)
	return NewComparisonResultInt(LevDist, s1, s2, &ld, nil)
}

// damerauLevenshteinDistance computes the Damerau-Levenshtein distance between
// two strings and returns a comparison result.
// This distance measures the minimum number of operations required to transform one string into the other.
// Supported operations include insertion, deletion, substitution, and transposition of adjacent characters.
func damerauLevenshteinDistance(str1, str2 string) *ComparisonResultInt {
	dld := edlib.DamerauLevenshteinDistance(str1, str2)
	return NewComparisonResultInt(DamLevDist, str1, str2, &dld, nil)
}

// osaDamerauLevenshteinDistance computes the optimal string alignment Damerau-Levenshtein distance between two strings.
// It returns a pointer to a ComparisonResultInt containing the result of the comparison between `str1` and `str2`.
// The returned structure includes the calculated distance, input strings, and any
// potential error encountered during computation.
func osaDamerauLevenshteinDistance(str1, str2 string) *ComparisonResultInt {
	osaDLD := edlib.OSADamerauLevenshteinDistance(str1, str2)
	return NewComparisonResultInt(OSADamLevDist, str1, str2, &osaDLD, nil)
}

// lcs returns the length of the longest common subsequence between two input strings, string1 and string2.
func lcs(str1 string, str2 string) int {
	return edlib.LCS(str1, str2)
}

// lcsEditDistance computes the edit distance between two strings based on their Longest Common Subsequence (LCS).
func lcsEditDistance(s1, s2 string) int {
	return edlib.LCSEditDistance(s1, s2)
}

// lcsBacktrack computes the longest common subsequence (LCS) between two input strings using backtracking.
// It returns an LCSResult containing the LCS result or an error if the computation fails.
func lcsBacktrack(str1 string, str2 string) *LCSResult {
	var resultSlice []string
	result, err := edlib.LCSBacktrack(str1, str2)
	if err != nil {
		return NewLCSResult(LCSBacktrackWord, str1, str2, nil, errors.Join(ErrLCSBacktrackFailure, err))
	}
	resultSlice = append(resultSlice, result)
	return NewLCSResult(LCSBacktrackWord, str1, str2, &resultSlice, nil)
}

// lcsBacktrackAll computes all possible longest common subsequences between two strings and returns an LCSResult.
// It utilizes the edlib.LCSBacktrackAll function for computation and wraps the result into an LCSResult instance.
// If an error occurs during computation, it is captured and included in the returned LCSResult.
func lcsBacktrackAll(str1 string, str2 string) *LCSResult {
	result, err := edlib.LCSBacktrackAll(str1, str2)
	if err != nil {
		return NewLCSResult(LCSBacktrackWordAll, str1, str2, nil, errors.Join(ErrLCSBacktrackAllFailure, err))
	}
	return NewLCSResult(LCSBacktrackWordAll, str1, str2, &result, nil)
}

// lcsDiff calculates the LCS difference between two strings and returns a
// pointer to an LCSResult containing the result.
// Returns nil for the result if an error occurs during the calculation.
func lcsDiff(str1, str2 string) *LCSResult {
	result, err := edlib.LCSDiff(str1, str2)
	if err != nil {
		return NewLCSResult(LCSDiffSlice, str1, str2, nil, errors.Join(ErrLCSDiffFailure, err))
	}
	return NewLCSResult(LCSDiffSlice, str1, str2, &result, nil)
}

// compareStringSlices checks if two slices of strings contain the same elements,
// regardless of order, with optional nil equality.
// If `nulls` is true, nil slices are treated as equal.
// It returns true if slices match, otherwise false.
func compareStringSlices(s1, s2 []string, nulls bool) bool {
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
		if !slices.Contains(remaining, s) {
			return false
		} else {
			for i, w := range remaining {
				if s == w {
					remaining = slices.Delete(remaining, i, i+1)
					break
				}
			}
		}
		continue
	}
	return true
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

// hammingDistance computes the Hamming distance between two strings s1 and s2.
// Returns a pointer to the integer distance and an error if the strings have unequal lengths or a failure occurs.
func hammingDistance(s1, s2 string) (*int, error) {
	dist, err := edlib.HammingDistance(s1, s2)
	if err != nil {
		return nil, errors.Join(ErrHammingDistanceFailure, err)
	}
	return &dist, nil
}

// jaroSimilarity computes the Jaro similarity between
// two strings s1 and s2, returning a float32 value in the range [0, 1].
func jaroSimilarity(s1, s2 string) float32 {
	return edlib.JaroSimilarity(s1, s2)
}

// jaroWinklerSimilarity computes the Jaro-Winkler similarity between two strings s1 and s2.
// It returns a float32 value between 0 and 1, where 1 indicates exact similarity.
func jaroWinklerSimilarity(s1, s2 string) float32 {
	return edlib.JaroWinklerSimilarity(s1, s2)
}

// jaccardSimilarity computes the Jaccard similarity coefficient between two strings based on a given split length.
// If split length of zero, the string is split on whitespaces and returns an index
func jaccardSimilarity(s1, s2 string, splitLength int) *float32 {
	if splitLength < 0 {
		return nil
	}
	js := edlib.JaccardSimilarity(s1, s2, splitLength)
	return &js
}

// cosineSimilarity computes the cosine similarity between two strings using a specified n-gram split length.
// Returns a pointer to the similarity score or nil if the split length is negative.
// Split lengths of zero, split on whitespaces.
func cosineSimilarity(s1, s2 string, splitLength int) *float32 {
	if splitLength < 0 {
		return nil
	}
	cs := edlib.CosineSimilarity(s1, s2, splitLength)
	return &cs
}

// sorensenDiceCoefficient calculates the Sørensen–Dice coefficient between two strings with a specified n-gram length.
// Returns a pointer to the coefficient value or nil if the splitLength is negative.
func sorensenDiceCoefficient(s1, s2 string, splitLength int) *float32 {
	if splitLength < 0 {
		return nil
	}
	sdc := edlib.SorensenDiceCoefficient(s1, s2, splitLength)
	return &sdc
}

// qgramDistance computes the q-gram distance between two strings s1 and s2 using a specified q-gram size q.
func qgramDistance(s1, s2 string, q int) *int {
	if q < 1 {
		return nil
	}
	qd := edlib.QgramDistance(s1, s2, q)
	return &qd
}

// qgramDistanceCustomNgram calculates the q-gram distance using two custom n-gram frequency maps.
// It returns the distance as an integer based on the differences between the input n-gram maps.
func qgramDistanceCustomNgram(nmap1, nmap2 map[string]int) int {
	return edlib.QgramDistanceCustomNgram(nmap1, nmap2)
}

// qgramSimilarity calculates the q-gram similarity between two strings using the specified q-gram size.
// Returns a pointer to the similarity score or nil if the q-gram size is less than 1.
func qgramSimilarity(s1, s2 string, q int) *float32 {
	if q < 1 {
		return nil
	}
	qs := edlib.QgramSimilarity(s1, s2, q)
	return &qs
}

// shingle generates k-shingles for the given string and returns them
// as a pointer to a map of shingles and their counts.
// Takes `s` as the input string and `k` as the size of each shingle.
// Returns nil if `k` is less than 1.
func shingle(s string, k int) *ShingleMapResult {
	if k < 1 {
		return NewShingleMapResult(ShinglesMap, s, k, nil, ErrShingleLengthOutOfRange)
	}
	shingle := edlib.Shingle(s, k)
	smr := NewShingleMapResult(ShinglesMap, s, k, &shingle, nil)
	return smr
}

// shingleSlice generates k-length shingles (substrings) from the input string and
// returns them as a pointer to a string slice.
// Returns nil if k is less than 1.
func shingleSlice(s string, k int) *ShingleSliceResult {
	if k < 1 {
		return NewShingleSliceResult(ShinglesSlice, s, k, nil, ErrShingleLengthOutOfRange)
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
		return NewSimilarityResult(algorithm, s1, s2, nil, err)
	}
	return NewSimilarityResult(algorithm, s1, s2, &sim, err)
}

// TODO update functions to return result objects
