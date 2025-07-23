package strutil

import (
	"fmt"

	"github.com/hbollon/go-edlib"
)

// Algorithm represents a type of string similarity or distance algorithm used for comparison.
type Algorithm edlib.Algorithm

// String returns the string representation of the Algorithm type based on the corresponding value in AlgorithmTypeMap.
func (a Algorithm) String() string {
	return AlgorithmTypeMap[a]
}

// Levenshtein represents the Levenshtein distance algorithm for string similarity measurement.
// DamerauLevenshtein represents the Damerau-Levenshtein algorithm for string similarity measurement.
// OSADamerauLevenshtein represents the Optimal String Alignment (OSA) Damerau-Levenshtein algorithm.
// Lcs represents the Longest Common Subsequence algorithm for string similarity measurement.
// Hamming represents the Hamming distance algorithm for string similarity measurement.
// Jaro represents the Jaro distance algorithm for string similarity measurement.
// JaroWinkler represents the Jaro-Winkler distance algorithm for string similarity measurement.
// Cosine represents the Cosine similarity algorithm used for vector-based string comparison.
// Jaccard represents the Jaccard similarity algorithm for set-based string comparison.
// SorensenDice represents the Sørensen-Dice coefficient for string similarity measurement.
// QGram represents the Q-Gram algorithm for string similarity measurement.
const (
	Levenshtein           = Algorithm(edlib.Levenshtein)
	DamerauLevenshtein    = Algorithm(edlib.DamerauLevenshtein)
	OSADamerauLevenshtein = Algorithm(edlib.OSADamerauLevenshtein)
	Lcs                   = Algorithm(edlib.Lcs)
	Hamming               = Algorithm(edlib.Hamming)
	Jaro                  = Algorithm(edlib.Jaro)
	JaroWinkler           = Algorithm(edlib.JaroWinkler)
	Cosine                = Algorithm(edlib.Cosine)
	Jaccard               = Algorithm(edlib.Jaccard)
	SorensenDice          = Algorithm(edlib.SorensenDice)
	QGram                 = Algorithm(edlib.Qgram)
)

// AlgorithmTypeMap maps edlib.Algorithm constants to their corresponding string representations for display purposes.
var AlgorithmTypeMap = map[Algorithm]string{
	Levenshtein:           "Levenshtein",
	DamerauLevenshtein:    "Damerau-Levenshtein",
	OSADamerauLevenshtein: "OSA Damerau-Levenshtein",
	Lcs:                   "LCS",
	Hamming:               "Hamming",
	Jaro:                  "Jaro",
	JaroWinkler:           "Jaro-Winkler",
	Cosine:                "Cosine",
	Jaccard:               "Jaccard",
	SorensenDice:          "Sorensen-Dice",
	QGram:                 "Q-Gram",
}

// SimilarityResult represents the result of a similarity computation between two strings.
type SimilarityResult struct {
	algorithm  Algorithm // the algorithm used
	string1    string    // input string/string builder value
	string2    string    // comparison value
	similarity *float32  // similarity result
	err        error     // error if it occurred
}

// NewSimilarityResult initializes and returns a new SimilarityResult instance with the provided parameters.
func NewSimilarityResult(algorithm Algorithm,
	str1 string,
	str2 string,
	similarity *float32,
	err error) *SimilarityResult {
	return &SimilarityResult{
		algorithm:  algorithm,
		string1:    str1,
		string2:    str2,
		similarity: similarity,
		err:        err,
	}
}

// GetAlgorithm returns the algorithm used for the similarity computation.
func (sr *SimilarityResult) GetAlgorithm() Algorithm {
	return sr.algorithm
}

// GetAlgorithmName returns the string representation of the algorithm used in the similarity computation.
func (sr *SimilarityResult) GetAlgorithmName() string {
	return sr.algorithm.String()
}

// GetString1 retrieves the first string used in the similarity comparison.
func (sr *SimilarityResult) GetString1() string {
	return sr.string1
}

// GetString2 returns the second string used in the similarity comparison.
func (sr *SimilarityResult) GetString2() string {
	return sr.string2
}

func (sr *SimilarityResult) GetStrings() (string, string) {
	return sr.string1, sr.string2
}

// Error returns the error encountered during the similarity calculation, if any.
func (sr *SimilarityResult) Error() error {
	return sr.err
}

// GetSimilarity returns a pointer to the similarity score calculated between string1 and string2.
func (sr *SimilarityResult) GetSimilarity() *float32 {
	return sr.similarity
}

// Print outputs the comparison result or error of the SimilarityResult based on verbosity flag 'v'.
func (sr *SimilarityResult) Print(v bool) {
	if v {
		if sr.err != nil {
			fmt.Printf("GetError during processing %s\nFirst String: %s\nSecond String: %s\nGetError: %s\n",
				sr.algorithm.String(), sr.string1, sr.string2, sr.err.Error())
			return
		} else {
			fmt.Printf("Comparison: %s\nFirst String: %s\nSecond String: %s\nScore: %f\n",
				sr.algorithm.String(), sr.string1, sr.string2, *sr.similarity)
			return
		}
	} else {
		if sr.err != nil {
			fmt.Printf("%s GetError: %s\n",
				sr.algorithm.String(), sr.err.Error())
			return
		} else {
			fmt.Printf("%s: %f\n", sr.algorithm.String(), *sr.similarity)
			return
		}
	}
}

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
func (smr SimilarityResultsMap) Print(v bool) {
	for algorithm, results := range smr {
		fmt.Printf("Algorithm: %s\n", algorithm.String())
		for word, result := range results {
			fmt.Printf("Comparison Word: %s\n", word)
			result.Print(v)
		}
	}
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
