package strutil

import "fmt"

// SimilarityResultsMap is keyed by algorithm and holds maps of similarity results keyed by comparison string.
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

// Print iterates through the SimilarityResultsMap and prints similarity results for each algorithm and comparison word.
func (smr SimilarityResultsMap) Print(v bool) SimilarityResultsMap {
	for algorithm, results := range smr {
		fmt.Printf("Algorithm: %s\n", algorithm.String())
		for word, result := range results {
			fmt.Printf("Comparison Word: %s\n", word)
			result.Print(v)
		}
	}
	return smr
}

// Add inserts or updates a SimilarityResult in the SimilarityResultsMap based on its algorithm and comparison word.
func (smr SimilarityResultsMap) Add(result SimilarityResult) {
	// if there's not a map for this algorithm, add one
	if smr[result.GetAlgorithm()] == nil {
		smr[result.GetAlgorithm()] = make(map[string]*SimilarityResult)
	}
	// add or update similarity for this comparison word
	smr[result.GetAlgorithm()][result.GetString2()] = &result
}
