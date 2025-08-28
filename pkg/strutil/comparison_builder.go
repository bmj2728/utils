package strutil

import (
	"github.com/bmj2728/utils/pkg/internal/errors"
)

// CompareStringBuilderSlices compares two slices of StringBuilder for equality,
// optionally allowing nil slices to be considered equal.
// The order of elements in the slices does not affect the comparisonData,
// and the 'nulls' flag determines nil-handling behavior.
func CompareStringBuilderSlices(a, b []StringBuilder, nulls bool) bool {
	return compareStringBuilderSlices(a, b, nulls)
}

// LevenshteinDistance calculates the Levenshtein distance between the current string and another string.
// It measures the minimum number of edit operations needed to transform the current string into the other string.
// The result is stored in the StringBuilder's comparison manager if it exists.
func (sb *StringBuilder) LevenshteinDistance(other string) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	ld := levenshteinDistance(sb.value, other)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(LevDist, sb.value, other).WithScoreInt(ld)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return ld
}

// DamerauLevenshteinDistance computes the Damerau-Levenshtein distance between the StringBuilder
// value and another string.
// It calculates the minimum transformation operations including insertion, deletion, substitution, and transposition.
// The result is stored in the ComparisonManager of the StringBuilder instance.
func (sb *StringBuilder) DamerauLevenshteinDistance(other string) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	dld := damerauLevenshteinDistance(sb.value, other)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(DamLevDist, sb.value, other).WithScoreInt(dld)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return dld
}

// OSADamerauLevenshteinDistance calculates the edit distance between the StringBuilder's
// current value and another string.
// It uses the optimal string alignment Damerau-Levenshtein distance algorithm and stores the
// result in the ComparisonManager.
// Returns the StringBuilder instance for chaining or error handling if an error exists in the current object.
func (sb *StringBuilder) OSADamerauLevenshteinDistance(other string) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}

	osaDld := osaDamerauLevenshteinDistance(sb.value, other)

	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(OSADamLevDist, sb.value, other).WithScoreInt(osaDld)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return osaDld
}

// LCS computes the longest common subsequence (LCS) between the current string and the provided string.
// It updates the comparison manager with the result of the LCS computation and returns the updated StringBuilder.
// If an error exists in the StringBuilder, it returns itself without performing computations.
func (sb *StringBuilder) LCS(other string) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	lcs := lcs(sb.value, other)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(LCSLength, sb.value, other).WithScoreInt(lcs)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return lcs
}

// LCSEditDistance calculates the edit distance based on the Longest Common Subsequence (LCS) between two strings.
// It updates the ComparisonManager of the StringBuilder instance with the result and returns the StringBuilder.
func (sb *StringBuilder) LCSEditDistance(other string) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	l := lcsEditDistance(sb.value, other)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(LCSDist, sb.value, other).WithScoreInt(l)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return l
}

// LCSBacktrack computes the longest common subsequence (LCS) between the StringBuilder's value and another string.
// It updates the StringBuilder's ComparisonManager with the LCS result and handles potential errors during computation.
// Returns the updated StringBuilder instance.
func (sb *StringBuilder) LCSBacktrack(other string) (string, error) {
	if !sb.shouldContinueProcessing() {
		return "", sb.err
	}
	lb, err := lcsBacktrack(sb.value, other)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		lbSl := []string{lb}
		res := NewLCSResult(LCSBacktrackWord, sb.value, other, &lbSl, err)
		sb.WithComparisonManager(false).comparisonManager.AddLCSResult(*res)
	}
	if err != nil {
		sb.setError(err, false)
		return "", err
	}
	return lb, nil
}

// LCSBacktrackAll computes all longest common subsequences between the StringBuilder's value and another string.
// It updates the ComparisonManager with the computed LCS result.
// If an error occurs during computation, it propagates the error state to the StringBuilder.
func (sb *StringBuilder) LCSBacktrackAll(other string) ([]string, error) {
	if !sb.shouldContinueProcessing() {
		return nil, sb.err
	}
	lba, err := lcsBacktrackAll(sb.value, other)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewLCSResult(LCSBacktrackWordAll, sb.value, other, &lba, err)
		sb.WithComparisonManager(false).comparisonManager.AddLCSResult(*res)
	}
	if err != nil {
		sb.setError(err, false)
		return nil, err
	}
	return lba, nil
}

// LCSDiff computes the Longest Common Subsequence (LCS) difference between the
// current StringBuilder value and another string.
// It updates the comparison manager with the LCS result and returns the updated StringBuilder instance.
// If an error occurs during the computation, it sets the error on the StringBuilder instance.
func (sb *StringBuilder) LCSDiff(other string) ([]string, error) {
	if !sb.shouldContinueProcessing() {
		return nil, sb.err
	}
	ld, err := lcsDiff(sb.value, other)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewLCSResult(LCSDiffSlice, sb.value, other, &ld, err)
		sb.WithComparisonManager(false).comparisonManager.AddLCSResult(*res)
	}
	if err != nil {
		sb.setError(err, false)
		return nil, err
	}
	return ld, nil
}

// HammingDistance computes the Hamming distance between the StringBuilder value and another
// string, updating comparison data.
// Returns the StringBuilder instance. If an error occurs, it sets the internal error and preserves the original state.
func (sb *StringBuilder) HammingDistance(other string) (int, error) {
	if !sb.shouldContinueProcessing() {
		return -1, sb.err
	}
	dist, err := hammingDistance(sb.value, other)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(HammingDist, sb.value, other).WithScoreInt(dist).WithError(err)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	if err != nil {
		sb.setError(err, false)
		return -1, err
	}
	return dist, nil
}

// JaroSimilarity computes the Jaro score between the StringBuilder's value and the provided string.
// It adds the result to the ComparisonManager and returns the updated StringBuilder instance.
//
// The higher the value, the more similar the strings are.
// The score is normalized such that 0 equates to no similarities and 1 is an exact match
//
// Additional Info: https://rosettacode.org/wiki/Jaro_similarity
func (sb *StringBuilder) JaroSimilarity(other string) float32 {
	if !sb.shouldContinueProcessing() {
		return -1.0
	}
	js := jaroSimilarity(sb.value, other)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(JaroSim, sb.value, other).WithScoreFloat(js)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return js
}

// JaroWinklerSimilarity calculates the Jaro-Winkler score between the StringBuilder value and another string.
// It adds the result to the ComparisonManager if no internal error is present.
// Returns the updated StringBuilder instance.
//
// Uses Jaro score with a more favorable weighting for similar common prefixes.
//
// Additional Info: https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance#Jaro%E2%80%93Winkler_similarity
func (sb *StringBuilder) JaroWinklerSimilarity(other string) float32 {
	if !sb.shouldContinueProcessing() {
		return -1.0
	}
	jws := jaroWinklerSimilarity(sb.value, other)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(JaroWinklerSim, sb.value, other).WithScoreFloat(jws)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return jws
}

// JaccardSimilarity computes the Jaccard score coefficient between two strings, using k-grams
// of the given split length.
// For splitLength = 0, the strings are split on whitespaces. Negative split lengths return nil
//
// The Jaccard index is defined as the size of the intersection divided by the size of the union
// for two given finite, non-empty sets
//
// Additional Info: https://en.wikipedia.org/wiki/Jaccard_index
func (sb *StringBuilder) JaccardSimilarity(other string, splitLength int) float32 {
	if !sb.shouldContinueProcessing() {
		return -1.0
	}
	js := jaccardSimilarity(sb.value, other, splitLength)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(JaccardSim, sb.value, other).WithSplitLength(splitLength).WithScoreFloat(js)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return js
}

// CosineSimilarity computes the cosine score between the StringBuilder value and
// another string with n-gram splitting. Updates the ComparisonManager with
// the computed score and returns the modified StringBuilder. When an error exists in the StringBuilder,
// it skips computation and returns itself.
//
// Cosine score is the cosine of the angle between the vectors.
//
// Additional Info: https://en.wikipedia.org/wiki/Cosine_similarity/
func (sb *StringBuilder) CosineSimilarity(other string, splitLength int) float32 {
	if !sb.shouldContinueProcessing() {
		return -1.0
	}
	cs := cosineSimilarity(sb.value, other, splitLength)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(CosineSim, sb.value, other).WithSplitLength(splitLength).WithScoreFloat(cs)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return cs
}

// SorensenDiceCoefficient computes the Sørensen–Dice coefficient for two strings using a given n-gram split length.
// Returns a pointer to the coefficient value or nil if the splitLength is negative.
//
// The Sørensen index equals twice the number of elements common to both sets
// divided by the sum of the number of elements in each set.
//
// Additional Info: https://en.wikipedia.org/wiki/Dice-S%C3%B8rensen_coefficient
func (sb *StringBuilder) SorensenDiceCoefficient(other string, splitLength int) float32 {
	if !sb.shouldContinueProcessing() {
		return -1.0
	}
	sdc := sorensenDiceCoefficient(sb.value, other, splitLength)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(SorensenDiceCo, sb.value, other).WithSplitLength(splitLength).WithScoreFloat(sdc)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return sdc
}

// QgramDistance calculates the Q-Gram distance between the current string and another string using a specified q value.
// Stores the result in the ComparisonManager for further access or management.
// Returns the updated StringBuilder instance.
func (sb *StringBuilder) QgramDistance(other string, q int) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	qd := qgramDistance(sb.value, other, q)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(QGramDist, sb.value, other).WithSplitLength(q).WithScoreInt(qd)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return qd
}

// QgramDistanceCustomNgram computes the Q-gram distance between the current StringBuilder's n-grams
// and a given n-gram map.
// It uses a custom comparison name, storing the result in the ComparisonManager if applicable,
// and returns the StringBuilder.
func (sb *StringBuilder) QgramDistanceCustomNgram(nmapOther map[string]int, customName string) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	// get ngram length of other map
	var k int
	for n := range nmapOther {
		k = len(n)
		if k < 1 {
			sb.setError(errors.ErrInvalidNgramMap, false)
			return -1
		}
		break
	}
	//TODO: refactor when shingles done - I hate this
	if sb.comparisonManager != nil || sb.lazyCompMan {
		//if we have a manager, we can use it. If it's nil and we have a lazy manager, we can init/use it.
		if sb.lazyCompMan {
			sb.WithComparisonManager(false)
		}
		// check if we have the shingles for this ngram length, if not, we generate them.
		if sb.GetComparisonManager().GetShingleResult(ShinglesMap, k) == nil {
			sb.Shingle(k)
		}
		shi := sb.GetComparisonManager().GetShingleResult(ShinglesMap, k)
		shiCasted := shi.(*ShingleMapResult)
		dist := qgramDistanceCustomNgram(shiCasted.shingles, nmapOther)
		res := NewComparisonResult(QGramDistCust, sb.value, customName).WithSplitLength(k).WithScoreInt(dist)
		sb.GetComparisonManager().AddComparisonResult(res)
		return dist
	} else if sb.comparisonManager == nil {
		// there's no manager and we do not lazy load, so we just don't cache anything.
		// refactor when shingles done
		_ = sb.Shingle(k)
		shi := sb.GetComparisonManager().GetShingleResult(ShinglesMap, k)
		shiCasted := shi.(*ShingleMapResult)
		return qgramDistanceCustomNgram(shiCasted.shingles, nmapOther)
	}
	return -1
}

// QgramSimilarity calculates the q-gram score between the builder's string
// and the given string using a specified q size.
// It updates the comparison data with the calculated score and returns the StringBuilder instance.
func (sb *StringBuilder) QgramSimilarity(other string, q int) float32 {
	if !sb.shouldContinueProcessing() {
		return -1.0
	}
	qs := qgramSimilarity(sb.value, other, q)
	if sb.comparisonManager != nil || sb.lazyCompMan {
		res := NewComparisonResult(QGramSim, sb.value, other).WithSplitLength(q).WithScoreFloat(qs)
		sb.WithComparisonManager(false).comparisonManager.AddComparisonResult(res)
	}
	return qs
}

// Shingle generates k-shingles for the current string value and stores the result using the ComparisonManager.
// Updates the error state if shingle generation fails. Returns the updated StringBuilder instance.
func (sb *StringBuilder) Shingle(k int) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	shingle := shingle(sb.value, k)
	sb.WithComparisonManager(false).comparisonManager.AddShingleResult(shingle)
	if shingle.err != nil {
		return sb.setError(shingle.err, false)
	}
	return sb
}

// ShingleSlice processes the string to generate k-length shingles, manages errors, and updates the ComparisonManager.
func (sb *StringBuilder) ShingleSlice(k int) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	shingle := shingleSlice(sb.value, k)
	sb.WithComparisonManager(false).comparisonManager.AddShingleResult(shingle)
	if shingle.err != nil {
		return sb.setError(shingle.err, false)
	}
	return sb
}

// Similarity computes the score between the current string and another string using the specified algorithm.
// Updates the ComparisonManager with the resulting score data and maintains the chainable state of StringBuilder.
// If an error occurs during computation, it sets the error state in the StringBuilder and returns itself.
func (sb *StringBuilder) Similarity(other string, algorithm Algorithm) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sr := similarity(sb.value, other, algorithm)
	sb.WithComparisonManager(false).comparisonManager.AddSimilarityResult(*sr)
	if sr.err != nil {
		return sb.setError(sr.err, false)
	}
	return sb
}
