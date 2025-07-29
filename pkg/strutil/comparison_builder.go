package strutil

// CompareStringBuilderSlices compares two slices of StringBuilder for equality,
// optionally allowing nil slices to be considered equal.
// The order of elements in the slices does not affect the comparisonData,
// and the 'nulls' flag determines nil-handling behavior.
func CompareStringBuilderSlices(a, b []StringBuilder, nulls bool) bool {
	return compareStringBuilderSlices(a, b, nulls)
}

// LevenshteinDistance calculates the Levenshtein distance between the current string and another string.
// It tracks and stores the result using a ComparisonManager.
// Returns the StringBuilder instance, enabling method chaining.
func (sb *StringBuilder) LevenshteinDistance(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	ld := levenshteinDistance(sb.value, other)
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(ld)
	if ld.err != nil {
		sb.err = ld.err
	}
	return sb
}

// DamerauLevenshteinDistance computes the Damerau-Levenshtein distance between the StringBuilder
// value and another string.
// It calculates the minimum transformation operations including insertion, deletion, substitution, and transposition.
// The result is stored in the ComparisonManager of the StringBuilder instance.
func (sb *StringBuilder) DamerauLevenshteinDistance(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	dld := damerauLevenshteinDistance(sb.value, other)
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(dld)
	if dld.err != nil {
		sb.err = dld.err
	}
	return sb
}

// OSADamerauLevenshteinDistance calculates the edit distance between the StringBuilder's
// current value and another string.
// It uses the optimal string alignment Damerau-Levenshtein distance algorithm and stores the
// result in the ComparisonManager.
// Returns the StringBuilder instance for chaining or error handling if an error exists in the current object.
func (sb *StringBuilder) OSADamerauLevenshteinDistance(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	osadld := osaDamerauLevenshteinDistance(sb.value, other)
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(osadld)
	if osadld.err != nil {
		sb.err = osadld.err
	}
	return sb
}

// LCS computes the longest common subsequence (LCS) between the current string and the provided string.
// It updates the comparison manager with the result of the LCS computation and returns the updated StringBuilder.
// If an error exists in the StringBuilder, it returns itself without performing computations.
func (sb *StringBuilder) LCS(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	lcs := lcs(sb.value, other)
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(lcs)
	if lcs.err != nil {
		sb.err = lcs.err
	}
	return sb
}

// LCSEditDistance calculates the edit distance based on the Longest Common Subsequence (LCS) between two strings.
// It updates the ComparisonManager of the StringBuilder instance with the result and returns the StringBuilder.
func (sb *StringBuilder) LCSEditDistance(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	l := lcsEditDistance(sb.value, other)
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(l)
	if l.err != nil {
		sb.err = l.err
	}
	return sb
}

// LCSBacktrack computes the longest common subsequence (LCS) between the StringBuilder's value and another string.
// It updates the StringBuilder's ComparisonManager with the LCS result and handles potential errors during computation.
// Returns the updated StringBuilder instance.
func (sb *StringBuilder) LCSBacktrack(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	lb := lcsBacktrack(sb.value, other)
	sb.WithComparisonManager().comparisonManager.AddLCSResult(*lb)
	if lb.err != nil {
		sb.err = lb.err
	}
	return sb
}

// LCSBacktrackAll computes all longest common subsequences between the StringBuilder's value and another string.
// It updates the ComparisonManager with the computed LCS result.
// If an error occurs during computation, it propagates the error state to the StringBuilder.
func (sb *StringBuilder) LCSBacktrackAll(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	lba := lcsBacktrackAll(sb.value, other)
	sb.WithComparisonManager().comparisonManager.AddLCSResult(*lba)
	if lba.err != nil {
		sb.err = lba.err
	}
	return sb
}

// LCSDiff computes the Longest Common Subsequence (LCS) difference between the
// current StringBuilder value and another string.
// It updates the comparison manager with the LCS result and returns the updated StringBuilder instance.
// If an error occurs during the computation, it sets the error on the StringBuilder instance.
func (sb *StringBuilder) LCSDiff(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	ld := lcsDiff(sb.value, other)
	sb.WithComparisonManager().comparisonManager.AddLCSResult(*ld)
	if ld.err != nil {
		sb.err = ld.err
	}
	return sb
}

// HammingDistance computes the Hamming distance between the StringBuilder value and another
// string, updating comparison data.
// Returns the StringBuilder instance. If an error occurs, it sets the internal error and preserves the original state.
func (sb *StringBuilder) HammingDistance(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	dist := hammingDistance(sb.value, other)
	if dist.err != nil {
		sb.err = dist.err
	}
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(dist)
	if dist.err != nil {
		sb.err = dist.err
	}
	return sb
}

// JaroSimilarity computes the Jaro score between the StringBuilder's value and the provided string.
// It adds the result to the ComparisonManager and returns the updated StringBuilder instance.
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
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(js)
	if js.err != nil {
		sb.err = js.err
	}
	return sb
}

// JaroWinklerSimilarity calculates the Jaro-Winkler score between the StringBuilder value and another string.
// It adds the result to the ComparisonManager if no internal error is present.
// Returns the updated StringBuilder instance.
//
// Uses Jaro score with a more favorable weighting for similar common prefixes.
//
// Additional Info: https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance#Jaro%E2%80%93Winkler_similarity
func (sb *StringBuilder) JaroWinklerSimilarity(other string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	jws := jaroWinklerSimilarity(sb.value, other)
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(jws)
	if jws.err != nil {
		sb.err = jws.err
	}
	return sb
}

// JaccardSimilarity computes the Jaccard score coefficient between two strings, using k-grams
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
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(js)
	if js.err != nil {
		sb.err = js.err
	}
	return sb
}

// CosineSimilarity computes the cosine score between the StringBuilder value and
// another string with n-gram splitting. Updates the ComparisonManager with
// the computed score and returns the modified StringBuilder. When an error exists in the StringBuilder,
// it skips computation and returns itself.
//
// Cosine score is the cosine of the angle between the vectors.
//
// Additional Info: https://en.wikipedia.org/wiki/Cosine_similarity/
func (sb *StringBuilder) CosineSimilarity(other string, splitLength int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	cs := cosineSimilarity(sb.value, other, splitLength)
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(cs)
	if cs.err != nil {
		sb.err = cs.err
	}
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
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(sdc)
	if sdc.err != nil {
		sb.err = sdc.err
	}
	return sb
}

// QgramDistance calculates the Q-Gram distance between the current string and another string using a specified q value.
// Stores the result in the ComparisonManager for further access or management.
// Returns the updated StringBuilder instance.
func (sb *StringBuilder) QgramDistance(other string, q int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	qd := qgramDistance(sb.value, other, q)
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(qd)
	if qd.err != nil {
		sb.err = qd.err
	}
	return sb
}

// QgramDistanceCustomNgram computes the Q-gram distance between the current StringBuilder's n-grams
// and a given n-gram map.
// It uses a custom comparison name, storing the result in the ComparisonManager if applicable,
// and returns the StringBuilder.
func (sb *StringBuilder) QgramDistanceCustomNgram(nmapOther map[string]int, customName string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	var k int
	for n := range nmapOther {
		k = len(n)
		break
	}
	if sb.comparisonManager == nil || sb.comparisonManager.ShingleData[ShinglesMap][k] == nil {
		return sb.WithComparisonManager().Shingle(k).QgramDistanceCustomNgram(nmapOther, customName)
	} else {
		sr := sb.comparisonManager.ShingleData[ShinglesMap][k]
		if shingleMap, ok := (*sr).(*ShingleMapResult); ok {
			qdc := qgramDistanceCustomNgram(shingleMap.shingles, nmapOther, customName)
			sb.WithComparisonManager().comparisonManager.AddComparisonResult(qdc)
			if qdc.err != nil {
				sb.err = qdc.err
			}
		}
	}
	return sb
}

// QgramSimilarity calculates the q-gram score between the builder's string
// and the given string using a specified q size.
// It updates the comparison data with the calculated score and returns the StringBuilder instance.
func (sb *StringBuilder) QgramSimilarity(other string, q int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	qs := qgramSimilarity(sb.value, other, q)
	sb.WithComparisonManager().comparisonManager.AddComparisonResult(qs)
	if qs.err != nil {
		sb.err = qs.err
	}
	return sb
}

// Shingle generates k-shingles for the current string value and stores the result using the ComparisonManager.
// Updates the error state if shingle generation fails. Returns the updated StringBuilder instance.
func (sb *StringBuilder) Shingle(k int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	shingle := shingle(sb.value, k)
	sb.WithComparisonManager().comparisonManager.AddShingleResult(shingle)
	if shingle.err != nil {
		sb.err = shingle.err
	}
	return sb
}

// ShingleSlice processes the string to generate k-length shingles, manages errors, and updates the ComparisonManager.
func (sb *StringBuilder) ShingleSlice(k int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	shingle := shingleSlice(sb.value, k)
	sb.WithComparisonManager().comparisonManager.AddShingleResult(shingle)
	if shingle.err != nil {
		sb.err = shingle.err
	}
	return sb
}

// Similarity computes the score between the current string and another string using the specified algorithm.
// Updates the ComparisonManager with the resulting score data and maintains the chainable state of StringBuilder.
// If an error occurs during computation, it sets the error state in the StringBuilder and returns itself.
func (sb *StringBuilder) Similarity(other string, algorithm Algorithm) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sr := similarity(sb.value, other, algorithm)
	sb.WithComparisonManager().comparisonManager.AddSimilarityResult(*sr)
	if sr.err != nil {
		sb.err = sr.err
	}
	return sb
}
