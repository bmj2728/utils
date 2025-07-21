package strutil

import "fmt"

// SimilarityTypeMap maps edlib.Algorithm constants to their corresponding string representations for display purposes.
var SimilarityTypeMap = map[string]string{
	"Levenshtein":           "Levenshtein",
	"DamerauLevenshtein":    "Damerau-Levenshtein",
	"OSADamerauLevenshtein": "OSA Damerau-Levenshtein",
	"Lcs":                   "LCS",
	"Hamming":               "Hamming",
	"Jaro":                  "Jaro",
	"JaroWinkler":           "Jaro-Winkler",
	"Cosine":                "Cosine",
	"Jaccard":               "Jaccard",
	"SorensenDice":          "Sorensen-Dice",
	"QGrams":                "Q-Gram",
}

// SimilarityResult represents the result of a similarity computation between two strings.
type SimilarityResult struct {
	algorithm  string   // the algorithm used
	string1    string   // input string/string builder value
	string2    string   // comparison value
	similarity *float32 // similarity result
	err        error    // error if it occurred
}

// NewSimilarityResult initializes and returns a new SimilarityResult instance with the provided parameters.
func NewSimilarityResult(algorithm string, str1 string, str2 string, similarity *float32, err error) *SimilarityResult {
	return &SimilarityResult{
		algorithm:  algorithm,
		string1:    str1,
		string2:    str2,
		similarity: similarity,
		err:        err,
	}
}

// GetAlgorithm retrieves the algorithm name used in the similarity computation.
func (sr *SimilarityResult) GetAlgorithm() string {
	return sr.algorithm
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
			fmt.Printf("Error during processing %s\nFirst String: %s\nSecond String: %s\nError: %s\n",
				SimilarityTypeMap[sr.algorithm], sr.string1, sr.string2, sr.err.Error())
			return
		} else {
			fmt.Printf("Comparison: %s\nFirst String: %s\nSecond String: %s\nScore: %f\n",
				SimilarityTypeMap[sr.algorithm], sr.string1, sr.string2, *sr.similarity)
			return
		}
	} else {
		if sr.err != nil {
			fmt.Printf("%s Error: %s\n",
				SimilarityTypeMap[sr.algorithm], sr.err.Error())
			return
		} else {
			fmt.Printf("%s: %f\n", SimilarityTypeMap[sr.algorithm], *sr.similarity)
			return
		}
	}
}
