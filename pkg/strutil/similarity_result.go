package strutil

import (
	"fmt"
	"math"

	"github.com/bmj2728/utils/pkg/internal/errors"
	"github.com/bmj2728/utils/pkg/internal/types"

	"github.com/hbollon/go-edlib"
)

// Algorithm represents a type of string score or distance algorithm used for comparison.
type Algorithm edlib.Algorithm

// String returns the string representation of the Algorithm type based on the corresponding value in AlgorithmTypeMap.
func (a Algorithm) String() string {
	return AlgorithmTypeMap[a]
}

// Levenshtein represents the Levenshtein distance algorithm for string score measurement.
// DamerauLevenshtein represents the Damerau-Levenshtein algorithm for string score measurement.
// OSADamerauLevenshtein represents the Optimal String Alignment (OSA) Damerau-Levenshtein algorithm.
// Lcs represents the Longest Common Subsequence algorithm for string score measurement.
// Hamming represents the Hamming distance algorithm for string score measurement.
// Jaro represents the Jaro distance algorithm for string score measurement.
// JaroWinkler represents the Jaro-Winkler distance algorithm for string score measurement.
// Cosine represents the Cosine score algorithm used for vector-based string comparison.
// Jaccard represents the Jaccard score algorithm for set-based string comparison.
// SorensenDice represents the Sørensen-Dice coefficient for string score measurement.
// QGram represents the Q-Gram algorithm for string score measurement.
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

// SimilarityResult represents the result of a score computation between two strings.
type SimilarityResult struct {
	algorithm Algorithm // the algorithm used
	string1   string    // input string/string builder value
	string2   string    // comparison value
	score     *float32  // score result
	err       error     // error if it occurred
}

// NewSimilarityResult initializes and returns a new SimilarityResult instance with the provided parameters.
func NewSimilarityResult(algorithm Algorithm,
	str1 string,
	str2 string,
	similarity *float32,
	err error) *SimilarityResult {
	return &SimilarityResult{
		algorithm: algorithm,
		string1:   str1,
		string2:   str2,
		score:     similarity,
		err:       err,
	}
}

// GetAlgorithm returns the algorithm used for the score computation.
func (sr *SimilarityResult) GetAlgorithm() Algorithm {
	return sr.algorithm
}

// GetAlgorithmName returns the string representation of the algorithm used in the score computation.
func (sr *SimilarityResult) GetAlgorithmName() string {
	return sr.algorithm.String()
}

// GetString1 retrieves the first string used in the score comparison.
func (sr *SimilarityResult) GetString1() string {
	return sr.string1
}

// GetString2 returns the second string used in the score comparison.
func (sr *SimilarityResult) GetString2() string {
	return sr.string2
}

func (sr *SimilarityResult) GetStrings() (string, string) {
	return sr.string1, sr.string2
}

// GetError returns the error encountered during the score calculation, if any.
func (sr *SimilarityResult) GetError() error {
	return sr.err
}

// GetScore returns a pointer to the score calculated between string1 and string2.
func (sr *SimilarityResult) GetScore() (float32, error) {
	if sr.score == nil && sr.err == nil {
		return 0, errors.ErrUnknownError
	}
	if sr.err != nil {
		return 0, sr.err
	}
	if sr.score == nil {
		return 0, errors.ErrNilScore
	}
	return *sr.score, nil
}

// IsMatch compares the current SimilarityResult with another based on their algorithm, strings, and computed scores.
func (sr *SimilarityResult) IsMatch(other *SimilarityResult) bool {
	if sr.algorithm != other.algorithm || sr.string1 != other.string1 || sr.string2 != other.string2 {
		return false
	}
	cScore, CErr := sr.GetScore()
	oScore, OErr := other.GetScore()
	if !errors.CompareErrors(CErr, OErr) {
		return false
	}
	if math.Abs(float64(cScore-oScore)) > types.Float64EqualityThreshold {
		return false
	}
	return true
}

// Print outputs the formatted score result or error information based on the verbosity flag.
func (sr *SimilarityResult) Print(v bool) {
	fmt.Print(formatSimilarityResultOutput(sr, v))
}

// formatSimilarityResultOutput formats the string output of a similarity
// result based on the verbosity flag and error state.
func formatSimilarityResultOutput(sr *SimilarityResult, v bool) string {
	if v {
		if sr.err != nil {
			return fmt.Sprintf("GetError during processing %s\nFirst String: %s\nSecond String: %s\nGetError: %s\n",
				sr.algorithm.String(), sr.string1, sr.string2, sr.err.Error())
		} else {
			return fmt.Sprintf("Comparison: %s\nFirst String: %s\nSecond String: %s\nScore: %f\n",
				sr.algorithm.String(), sr.string1, sr.string2, *sr.score)
		}
	} else {
		if sr.err != nil {
			return fmt.Sprintf("%s GetError: %s\n",
				sr.algorithm.String(), sr.err.Error())
		} else {
			return fmt.Sprintf("%s: %f\n", sr.algorithm.String(), *sr.score)
		}
	}
}
