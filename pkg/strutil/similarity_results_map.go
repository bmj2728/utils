package strutil

import "fmt"

// SimilarityResultsMap is keyed by algorithm and holds maps of score results keyed by comparison string.
//
//	map[Levenshtein]["comparison text"]SimilarityResult{
//		algorithm: Levenshtein,
//		string1: "original",
//		string2: "comparison text",
//		score: 0.12,
//		err: nil}
type SimilarityResultsMap map[Algorithm]map[string]*SimilarityResult

// NewSimilarityResultsMap initializes and returns a new SimilarityResultsMap as an empty map.
func NewSimilarityResultsMap() SimilarityResultsMap {
	return make(map[Algorithm]map[string]*SimilarityResult)
}

// Add inserts or updates a SimilarityResult in the SimilarityResultsMap based on its algorithm and comparison word.
func (smr SimilarityResultsMap) Add(result SimilarityResult) {
	// if there's not a map for this algorithm, add one
	if smr[result.GetAlgorithm()] == nil {
		smr[result.GetAlgorithm()] = make(map[string]*SimilarityResult)
	}
	// add or update score for this comparison word
	smr[result.GetAlgorithm()][result.GetString2()] = &result
}

// GetCopy creates and returns a deep copy of the SimilarityResultsMap, ensuring all nested maps and values are cloned.
func (smr SimilarityResultsMap) GetCopy() SimilarityResultsMap {
	cloned := NewSimilarityResultsMap()
	for algo, val := range smr {
		cloned[algo] = make(map[string]*SimilarityResult)
		for compStr, result := range val {
			if result != nil {
				resultCopy := *result
				cloned[algo][compStr] = &resultCopy
			}
		}
	}
	return cloned
}

// Get retrieves a pointer to the SimilarityResult for the given Algorithm and
// comparison string, or returns nil if not found.
func (smr SimilarityResultsMap) Get(algo Algorithm, compStr string) *SimilarityResult {
	if smr[algo] == nil {
		return nil
	}
	if smr[algo][compStr] == nil {
		return nil
	}
	return smr[algo][compStr]
}

// GetByType retrieves all SimilarityResult entries for a specified Algorithm from the SimilarityResultsMap.
func (smr SimilarityResultsMap) GetByType(algo Algorithm) []SimilarityResult {
	if smr[algo] == nil {
		return nil
	}
	var results []SimilarityResult
	for _, v := range smr[algo] {
		results = append(results, *v)
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// FilterByType filters the SimilarityResultsMap to include only results of the
// specified algorithm and returns a new map.
func (smr SimilarityResultsMap) FilterByType(algo Algorithm) SimilarityResultsMap {
	if smr[algo] == nil {
		return nil
	}
	results := NewSimilarityResultsMap()
	for _, v := range smr[algo] {
		results.Add(*v)
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// GetByComparisonString retrieves all SimilarityResult entries that match the specified
// comparison string across all algorithms.
func (smr SimilarityResultsMap) GetByComparisonString(compStr string) []SimilarityResult {
	if len(smr) == 0 {
		return nil
	}
	var results []SimilarityResult
	for _, v := range smr {
		if v[compStr] != nil {
			results = append(results, *v[compStr])
		}
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// FilterByComparisonString filters the map to include only results matching the given
// comparison string across all algorithms.
func (smr SimilarityResultsMap) FilterByComparisonString(compStr string) SimilarityResultsMap {
	if len(smr) == 0 {
		return nil
	}
	results := NewSimilarityResultsMap()
	for _, v := range smr {
		if v[compStr] != nil {
			results.Add(*v[compStr])
		}
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// TypeCount returns the number of algorithms present in the SimilarityResultsMap.
func (smr SimilarityResultsMap) TypeCount() int {
	return len(smr)
}

// EntryCount returns the total number of non-nil SimilarityResult entries within the SimilarityResultsMap.
func (smr SimilarityResultsMap) EntryCount() int {
	mapLength := 0
	for _, v := range smr {
		for _, v2 := range v {
			if v2 != nil {
				mapLength++
			}
		}
	}
	return mapLength
}

// IsMatch compares two SimilarityResultsMap objects for equality by checking type
// and entry counts, and all nested results.
func (smr SimilarityResultsMap) IsMatch(other SimilarityResultsMap) bool {
	if smr.TypeCount() != other.TypeCount() || smr.EntryCount() != other.EntryCount() {
		return false
	}
	for algo, v := range smr {
		if other[algo] == nil {
			return false
		}
		for compStr, v2 := range v {
			if other[algo][compStr] == nil {
				return false
			} else {
				if !v2.IsMatch(other[algo][compStr]) {
					return false
				}
			}
		}
	}
	return true
}

// Print iterates through the SimilarityResultsMap and prints score results for each algorithm and comparison word.
func (smr SimilarityResultsMap) Print(v bool) SimilarityResultsMap {
	fmt.Print(formatSimilarityMapOutput(smr, v))
	return smr
}

// formatSimilarityMapOutput generates a formatted string output for all similarity results organized by algorithm.
// It includes verbose details if the verbose flag is set to true.
func formatSimilarityMapOutput(smr SimilarityResultsMap, verbose bool) string {
	var output string
	for algo, v := range smr {
		output += fmt.Sprintf("***Similarity Results for %s***\n\n", algo.String())
		for _, v2 := range v {
			if v2 != nil {
				output += formatSimilarityResultOutput(v2, verbose) + "\n"
			}
		}
	}
	return output
}
