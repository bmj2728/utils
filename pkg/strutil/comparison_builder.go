package strutil

import (
	"errors"
)

// CompareStringBuilderSlices compares two slices of StringBuilder for equality,
// optionally allowing nil slices to be considered equal.
// The order of elements in the slices does not affect the comparisonData,
// and the 'nulls' flag determines nil-handling behavior.
func CompareStringBuilderSlices(a, b []StringBuilder, nulls bool) bool {
	return compareStringBuilderSlices(a, b, nulls)
}

// LevenshteinDistance calculates the Levenshtein distance between the StringBuilder's value and the provided string.
//
// It represents the minimum number of edits needed to convert one string into the other.
// An edit is an insertion, deletion, or substitution of a single character.
//
// Additional information: https://en.wikipedia.org/wiki/Levenshtein_distance
func (sb *StringBuilder) LevenshteinDistance(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	ld := levenshteinDistance(sb.value, other)
	sb.comparisonData.SetLevenshteinDist(&ld)
	return sb
}

// DamerauLevenshteinDistance computes the edit distance between two strings,
// including transpositions of adjacent characters.
//
// It represents the minimum number of operations to change one string to another.
// An operation is an insertion, deletion/substitution of a single character, or transposition of adjacent characters.
//
// Additional information: https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance
func (sb *StringBuilder) DamerauLevenshteinDistance(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	dld := damerauLevenshteinDistance(sb.value, other)
	sb.comparisonData.SetDamerauLevDist(&dld)
	return sb
}

// OSADamerauLevenshteinDistance calculates the optimal string alignment Damerau-Levenshtein distance
// with the given string.
// Updates the comparisonData field with the computed distance value.
func (sb *StringBuilder) OSADamerauLevenshteinDistance(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	osadld := osaDamerauLevenshteinDistance(sb.value, other)
	sb.comparisonData.SetOSADamerauLevDist(&osadld)
	return sb
}

// LCS calculates and returns the length of the longest common subsequence (LCS) between
// the StringBuilder value and another string.
func (sb *StringBuilder) LCS(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	lcs := lcs(sb.value, other)
	sb.comparisonData.SetLCS(&lcs)
	return sb
}

// LCSBacktrack computes the longest common subsequence (LCS) between the StringBuilder's value and another string.
// Updates the internal state with the computed LCS or sets an error if the operation fails.
// Returns the StringBuilder instance for method chaining.
func (sb *StringBuilder) LCSBacktrack(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	s, err := lcsBacktrack(sb.value, other)
	if err != nil {
		sb.err = errors.Join(ErrLCSBacktrackFailure, err)
	}
	sb.comparisonData.SetLCSBacktrack(&s)
	return sb
}

// LCSBacktrackAll computes all longest common subsequences between the current StringBuilder value and another string.
// Updates the comparisonData field with all subsequences if successful or records an error if the operation fails.
// Returns the StringBuilder instance.
func (sb *StringBuilder) LCSBacktrackAll(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	seqs, err := lcsBacktrackAll(sb.value, other)
	if err != nil {
		sb.err = errors.Join(ErrLCSBacktrackAllFailure, err)
	}
	sb.comparisonData.SetLCSBacktrackAll(&seqs)
	return sb
}

// LCSDiff calculates and sets the Longest Common Subsequence diff between
// the StringBuilder's value and the given string.
// It modifies the StringBuilder by updating its comparisonData state with the LCS diff.
// If an error occurs during the calculation, the error is stored in the StringBuilder's error field.
func (sb *StringBuilder) LCSDiff(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	diff, err := lcsDiff(sb.value, other)
	if err != nil {
		sb.err = errors.Join(ErrLCSDiffFailure, err)
	}
	sb.comparisonData.SetLCSDiff(&diff)
	return sb
}

// LCSEditDistance calculates the edit distance between the StringBuilder's value and another string using LCS.
func (sb *StringBuilder) LCSEditDistance(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	i := lcsEditDistance(sb.value, other)
	sb.comparisonData.SetLCSEditDistance(&i)
	return sb
}

func (sb *StringBuilder) HammingDistance(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	dist, err := hammingDistance(sb.value, other)
	if err != nil {
		sb.err = errors.Join(ErrHammingDistanceFailure, err)
		return sb
	}
	sb.comparisonData.SetHammingDist(dist)
	return sb
}

// JaroSimilarity computes the Jaro similarity between the StringBuilder's value and the provided string.
// It updates the Jaro similarity value in the comparisonData field and returns the updated StringBuilder instance.
//
// The higher the value, the more similar the strings are.
// The score is normalized such that 0 equates to no similarities and 1 is an exact match
//
// Additional Info: https://rosettacode.org/wiki/Jaro_similarity
func (sb *StringBuilder) JaroSimilarity(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	js := jaroSimilarity(sb.value, other)
	sb.comparisonData.SetJaroSimilarity(&js)
	return sb
}

// JaroWinklerSimilarity calculates the Jaro-Winkler similarity between the StringBuilder value and another string.
// It updates the comparisonData field with the computed similarity value if no internal error is present.
// Returns the updated StringBuilder instance.
//
// Uses Jaro similarity with a more favorable weighting for similar common prefixes.
//
// Additional Info: https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance#Jaro%E2%80%93Winkler_similarity
func (sb *StringBuilder) JaroWinklerSimilarity(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	jws := jaroWinklerSimilarity(sb.value, other)
	sb.comparisonData.SetJaroWinklerSim(&jws)
	return sb
}

// JaccardSimilarity computes the Jaccard similarity coefficient between two strings, using k-grams
// of the given split length.
// For splitLength = 0, the strings are split on whitespaces. Negative split lengths return nil
//
// The Jaccard index is defined as the size of the intersection divided by the size of the union
// for two given finite, non-empty sets
//
// Additional Info: https://en.wikipedia.org/wiki/Jaccard_index
func (sb *StringBuilder) JaccardSimilarity(other string, splitLength int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	js := jaccardSimilarity(sb.value, other, splitLength)
	sb.comparisonData.SetJaccardSim(js)
	return sb
}

// CosineSimilarity computes the cosine similarity between the StringBuilder value and
// another string with n-gram splitting. Updates the comparisonData state with
// the computed similarity and returns the modified StringBuilder. When an error exists in the StringBuilder,
// it skips computation and returns itself.
//
// Cosine similarity is the cosine of the angle between the vectors.
//
// Additional Info: https://en.wikipedia.org/wiki/Cosine_similarity/
func (sb *StringBuilder) CosineSimilarity(other string, splitLength int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	cs := cosineSimilarity(sb.value, other, splitLength)
	sb.comparisonData.SetCosineSimilarity(cs)
	return sb
}

// SorensenDiceCoefficient computes the Sørensen–Dice coefficient for two strings using a given n-gram split length.
// Returns a pointer to the coefficient value or nil if the splitLength is negative.
//
// The Sørensen index equals twice the number of elements common to both sets
// divided by the sum of the number of elements in each set.
//
// Additional Info: https://en.wikipedia.org/wiki/Dice-S%C3%B8rensen_coefficient
func (sb *StringBuilder) SorensenDiceCoefficient(other string, splitLength int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sdc := sorensenDiceCoefficient(sb.value, other, splitLength)
	sb.comparisonData.SetSorensenDiceCo(sdc)
	return sb
}

// QgramDistance computes the q-gram distance between the StringBuilder's value
// and another string with the specified q value.
// It updates the comparisonData's QGramDist field with the computed distance.
// Returns the StringBuilder instance for method chaining.
func (sb *StringBuilder) QgramDistance(other string, q int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	qd := qgramDistance(sb.value, other, q)
	sb.comparisonData.SetQGramDist(qd)
	return sb
}

// QgramDistanceCustomNgram calculates the q-gram distance between the current string value and another n-gram map.
// If the StringBuilder doesn't have an existing ComparisonData.Shingle map, one is generated using the n-gram size
// of the comparison map.
func (sb *StringBuilder) QgramDistanceCustomNgram(nmapOther map[string]int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	var k int
	for n := range nmapOther {
		k = len(n)
		break
	}
	if sb.comparisonManager == nil || sb.comparisonManager.ShingleData[ShinglesMap][k] == nil {
		return sb.WithComparisonManager().Shingle(k).QgramDistanceCustomNgram(nmapOther)
	} else {
		sr := sb.comparisonManager.ShingleData[ShinglesMap][k]
		if shingleMap, ok := (*sr).(*ShingleMapResult); ok {
			qdc := qgramDistanceCustomNgram(*shingleMap.shingles, nmapOther)
			sb.comparisonData.SetQGramDistCustom(&qdc)
		}
	}
	return sb
}

// QgramSimilarity calculates the q-gram similarity between the builder's string
// and the given string using a specified q size.
// It updates the comparison data with the calculated similarity and returns the StringBuilder instance.
func (sb *StringBuilder) QgramSimilarity(other string, q int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	qs := qgramSimilarity(sb.value, other, q)
	sb.comparisonData.SetQGramSim(qs)
	return sb
}

// Shingle generates k-shingles from the StringBuilder's value and sets them as comparison data for further operations.
func (sb *StringBuilder) Shingle(k int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	shingle := shingle(sb.value, k)
	sb.comparisonManager.AddShingleResult(shingle)
	return sb
}

// ShingleSlice generates k-length shingles from the StringBuilder's value and updates the comparison data.
// Returns the StringBuilder instance for method chaining.
func (sb *StringBuilder) ShingleSlice(k int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	shingle := shingleSlice(sb.value, k)
	sb.comparisonManager.AddShingleResult(shingle)
	return sb
}

// Similarity computes the similarity between the current string and another
// provided string using the specified algorithm.
// It appends the result to the similarities list and returns the updated StringBuilder.
// If an error is already set in the StringBuilder, it skips computation and returns itself.
func (sb *StringBuilder) Similarity(other string, algorithm Algorithm) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sr := similarity(sb.value, other, algorithm)
	sb.ComparisonManager().AddSimilarityResult(*sr)
	return sb
}
