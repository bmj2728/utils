package strutil

import (
	"errors"
	"slices"

	errors2 "github.com/bmj2728/utils/pkg/internal/errors"

	"github.com/hbollon/go-edlib"
)

// levenshteinDistance computes the Levenshtein distance between two strings and returns a ComparisonResultInt.
// It calculates the minimum number of edits (insertions, deletions, substitutions)
// required to transform one string into another.
func levenshteinDistance(s1, s2 string) *ComparisonResultInt {
	ld := edlib.LevenshteinDistance(s1, s2)
	return NewComparisonResultInt(LevDist, s1, s2, nil, &ld, nil)
}

// damerauLevenshteinDistance computes the Damerau-Levenshtein distance between
// two strings and returns a comparison result.
// This distance measures the minimum number of operations required to transform one string into the other.
// Supported operations include insertion, deletion, substitution, and transposition of adjacent characters.
func damerauLevenshteinDistance(str1, str2 string) *ComparisonResultInt {
	dld := edlib.DamerauLevenshteinDistance(str1, str2)
	return NewComparisonResultInt(DamLevDist, str1, str2, nil, &dld, nil)
}

// osaDamerauLevenshteinDistance computes the optimal string alignment Damerau-Levenshtein distance between two strings.
// It returns a pointer to a ComparisonResultInt containing the result of the comparison between `str1` and `str2`.
// The returned structure includes the calculated distance, input strings, and any
// potential error encountered during computation.
func osaDamerauLevenshteinDistance(str1, str2 string) *ComparisonResultInt {
	osaDLD := edlib.OSADamerauLevenshteinDistance(str1, str2)
	return NewComparisonResultInt(OSADamLevDist, str1, str2, nil, &osaDLD, nil)
}

// lcs returns the length of the longest common subsequence between two input strings, string1 and string2.
func lcs(str1 string, str2 string) *ComparisonResultInt {
	l := edlib.LCS(str1, str2)
	return NewComparisonResultInt(LCSLength, str1, str2, nil, &l, nil)
}

// lcsEditDistance computes the edit distance between two strings based on their Longest Common Subsequence (LCS).
func lcsEditDistance(s1, s2 string) *ComparisonResultInt {
	l := edlib.LCSEditDistance(s1, s2)
	return NewComparisonResultInt(LCSDist, s1, s2, nil, &l, nil)
}

// lcsBacktrack computes the longest common subsequence (LCS) between two input strings using backtracking.
// It returns an LCSResult containing the LCS result or an error if the computation fails.
func lcsBacktrack(str1 string, str2 string) *LCSResult {
	var resultSlice []string
	result, err := edlib.LCSBacktrack(str1, str2)
	if err != nil {
		return NewLCSResult(LCSBacktrackWord, str1, str2, nil, errors.Join(errors2.ErrLCSBacktrackFailure, err))
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
		return NewLCSResult(LCSBacktrackWordAll, str1, str2, nil, errors.Join(errors2.ErrLCSBacktrackAllFailure, err))
	}
	return NewLCSResult(LCSBacktrackWordAll, str1, str2, &result, nil)
}

// lcsDiff calculates the LCS difference between two strings and returns a
// pointer to an LCSResult containing the result.
// Returns nil for the result if an error occurs during the calculation.
func lcsDiff(str1, str2 string) *LCSResult {
	result, err := edlib.LCSDiff(str1, str2)
	if err != nil {
		return NewLCSResult(LCSDiffSlice, str1, str2, nil, errors.Join(errors2.ErrLCSDiffFailure, err))
	}
	return NewLCSResult(LCSDiffSlice, str1, str2, &result, nil)
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

// hammingDistance computes the Hamming distance between two strings and returns a ComparisonResultInt instance.
// Returns an error if the strings are not of equal length or if the distance calculation fails.
func hammingDistance(s1, s2 string) *ComparisonResultInt {
	dist, err := edlib.HammingDistance(s1, s2)
	if err != nil {
		return NewComparisonResultInt(HammingDist, s1, s2, nil, nil, errors.Join(errors2.ErrHammingDistanceFailure, err))
	}
	return NewComparisonResultInt(HammingDist, s1, s2, nil, &dist, nil)
}

// jaroSimilarity computes the Jaro similarity between two strings
// and returns the result as a ComparisonResultFloat object.
func jaroSimilarity(s1, s2 string) *ComparisonResultFloat {
	js := edlib.JaroSimilarity(s1, s2)
	return NewComparisonResultFloat(JaroSim, s1, s2, nil, &js, nil)
}

// jaroWinklerSimilarity computes the Jaro-Winkler similarity between two input strings.
// Returns a ComparisonResultFloat containing the similarity score and input data.
func jaroWinklerSimilarity(s1, s2 string) *ComparisonResultFloat {
	js := edlib.JaroWinklerSimilarity(s1, s2)
	return NewComparisonResultFloat(JaroWinklerSim, s1, s2, nil, &js, nil)
}

// jaccardSimilarity calculates the Jaccard similarity coefficient between two strings
// using k-grams of a specified length.
// Returns a ComparisonResultFloat object containing the result or an error if the splitLength is invalid (< 0).
func jaccardSimilarity(s1, s2 string, splitLength int) *ComparisonResultFloat {
	if splitLength < 0 {
		return NewComparisonResultFloat(JaccardSim, s1, s2, &splitLength, nil, errors2.ErrInvalidLengthRange)
	}
	js := edlib.JaccardSimilarity(s1, s2, splitLength)
	return NewComparisonResultFloat(JaccardSim, s1, s2, &splitLength, &js, nil)
}

// cosineSimilarity calculates the cosine similarity between two strings using n-gram splitting with the given length.
// Returns a pointer to a ComparisonResultFloat containing the similarity score or an error
// if the split length is invalid.
func cosineSimilarity(s1, s2 string, splitLength int) *ComparisonResultFloat {
	if splitLength < 0 {
		return NewComparisonResultFloat(CosineSim, s1, s2, &splitLength, nil, errors2.ErrInvalidLengthRange)
	}
	cs := edlib.CosineSimilarity(s1, s2, splitLength)
	return NewComparisonResultFloat(CosineSim, s1, s2, &splitLength, &cs, nil)
}

// sorensenDiceCoefficient calculates the Sørensen–Dice coefficient of two strings based on a specified n-gram length.
// Returns a pointer to a ComparisonResultFloat, containing the coefficient
// value or an error if the splitLength is invalid.
func sorensenDiceCoefficient(s1, s2 string, splitLength int) *ComparisonResultFloat {
	if splitLength < 0 {
		return NewComparisonResultFloat(SorensenDiceCo, s1, s2, &splitLength, nil, errors2.ErrInvalidLengthRange)
	}
	sdc := edlib.SorensenDiceCoefficient(s1, s2, splitLength)
	return NewComparisonResultFloat(SorensenDiceCo, s1, s2, &splitLength, &sdc, nil)
}

// qgramDistance computes the Q-Gram distance between two strings s1 and s2 using the specified q-gram size q.
// Returns a ComparisonResultInt containing the computed distance or an error if q is invalid.
func qgramDistance(s1, s2 string, q int) *ComparisonResultInt {
	if q < 0 {
		return NewComparisonResultInt(QGramDist, s1, s2, &q, nil, errors2.ErrInvalidLengthRange)
	}
	qd := edlib.QgramDistance(s1, s2, q)
	return NewComparisonResultInt(QGramDist, s1, s2, &q, &qd, nil)
}

// qgramDistanceCustomNgram calculates the Q-Gram distance between two n-gram maps and returns a comparison result.
func qgramDistanceCustomNgram(nmap1, nmap2 map[string]int, customName string) *ComparisonResultInt {
	qdc := edlib.QgramDistanceCustomNgram(nmap1, nmap2)
	return NewComparisonResultInt(QGramDistCust, "", customName, nil, &qdc, nil)
}

// qgramSimilarity calculates the q-gram similarity between two strings with a specified q size.
// Returns a ComparisonResultFloat containing the similarity score or an error if q is invalid.
func qgramSimilarity(s1, s2 string, q int) *ComparisonResultFloat {
	if q < 1 {
		return NewComparisonResultFloat(QGramSim, s1, s2, &q, nil, errors2.ErrInvalidLengthRange)
	}
	qs := edlib.QgramSimilarity(s1, s2, q)
	return NewComparisonResultFloat(QGramSim, s1, s2, &q, &qs, nil)
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
