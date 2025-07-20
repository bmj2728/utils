package strutil

import (
	"errors"
	"slices"

	"github.com/hbollon/go-edlib"
)

// levenshteinDistance calculates the Levenshtein distance between two strings s1 and s2.
// It determines the minimum number of single-character edits (insertions, deletions, or substitutions) required.
func levenshteinDistance(s1, s2 string) int {
	return edlib.LevenshteinDistance(s1, s2)
}

// damerauLevenshteinDistance calculates the Damerau-Levenshtein distance
// between two strings to measure edit similarity.
func damerauLevenshteinDistance(str1, str2 string) int {
	return edlib.DamerauLevenshteinDistance(str1, str2)
}

// osaDamerauLevenshteinDistance calculates the Damerau-Levenshtein distance between str1 and str2,
// considering adjacent transpositions.
//
// This optimal string alignment variant of damerauLevenshteinDistance
// does not allow multiple transformations on the same substring
func osaDamerauLevenshteinDistance(str1, str2 string) int {
	return edlib.OSADamerauLevenshteinDistance(str1, str2)
}

// lcs returns the length of the longest common subsequence between two input strings, str1 and str2.
func lcs(str1 string, str2 string) int {
	return edlib.LCS(str1, str2)
}

// lcsBacktrack computes the longest common subsequence (LCS) of two input strings and returns the result as a string.
func lcsBacktrack(str1 string, str2 string) (string, error) {
	result, err := edlib.LCSBacktrack(str1, str2)
	if err != nil {
		return "", errors.Join(ErrLCSBacktrackFailure, err)
	}
	return result, nil
}

// lcsBacktrackAll computes all longest common subsequences of two input strings.
// Returns a slice of strings representing the subsequences and an error if the computation fails.
func lcsBacktrackAll(str1 string, str2 string) ([]string, error) {
	result, err := edlib.LCSBacktrackAll(str1, str2)
	if err != nil {
		return nil, errors.Join(ErrLCSBacktrackAllFailure, err)
	}
	return result, nil
}

// lcsDiff calculates the Longest Common Subsequence (LCS) diff between
// two strings and returns the differences or an error.
func lcsDiff(str1, str2 string) ([]string, error) {
	result, err := edlib.LCSDiff(str1, str2)
	if err != nil {
		return nil, errors.Join(ErrLCSDiffFailure, err)
	}
	return result, nil
}

// lcsEditDistance computes the edit distance between two strings based on their Longest Common Subsequence (LCS).
func lcsEditDistance(s1, s2 string) int {
	return edlib.LCSEditDistance(s1, s2)
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
func shingle(s string, k int) *map[string]int {
	if k < 1 {
		return nil
	}
	shingle := edlib.Shingle(s, k)
	return &shingle
}

// shingleSlice generates k-length shingles (substrings) from the input string and
// returns them as a pointer to a string slice.
// Returns nil if k is less than 1.
func shingleSlice(s string, k int) *[]string {
	if k < 1 {
		return nil
	}
	shingle := edlib.ShingleSlice(s, k)
	return &shingle
}
