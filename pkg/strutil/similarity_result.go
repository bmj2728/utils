package strutil

import "fmt"

// SimilarityResult represents the outcome of a similarity comparison between two data entities.
type SimilarityResult struct {
	Algorithm  string
	Str1       string
	Str2       string
	Similarity *float32
	Err        error
}

// NewSimilarityResult creates and returns a pointer to a new SimilarityResult
// with the specified algorithm, other, and similarity.
func NewSimilarityResult(algorithm string, str1 string, str2 string, similarity *float32, err error) *SimilarityResult {
	return &SimilarityResult{
		Algorithm:  algorithm,
		Str1:       str1,
		Str2:       str2,
		Similarity: similarity,
		Err:        err,
	}
}

func (sr *SimilarityResult) String() string {
	if sr.Similarity == nil {
		return fmt.Sprintf("%s / %s \n(%s - No Score): %s", sr.Str1, sr.Str2, sr.Algorithm, sr.Err.Error())
	} else {
		return fmt.Sprintf("%s / %s \n(%s - %.2f)", sr.Str1, sr.Str2, sr.Algorithm, *sr.Similarity)
	}
}

// Error returns the error associated with the similarity result, or nil if no error occurred.
func (sr *SimilarityResult) Error() error {
	return sr.Err
}

// Print generates a StringBuilder containing the formatted string representation of the SimilarityResult instance.
func (sr *SimilarityResult) Print() *StringBuilder {
	return New(sr.String()).Print()
}
