package strutil

// ComparisonManager is a structure for managing comparison, similarity, shingle, and LCS results.
type ComparisonManager struct {
	ComparisonResults *map[string]map[string]ComparisonResult
	SimilarityResults *map[string]map[string]SimilarityResult
	ShingleData       *map[string]map[string]ShingleResult
	LCSData           *map[string]map[string]LCSResult
}

// AddSimilarity stores a SimilarityResult in the SimilarityResults map, categorized by
// algorithm and comparison string.
// e.g. (*ComparisonResults[LevDist)
func (cm *ComparisonManager) AddSimilarity(result SimilarityResult) {
	if (*cm.SimilarityResults)[result.GetAlgorithm()] == nil {
		(*cm.SimilarityResults)[result.GetAlgorithm()] = make(map[string]SimilarityResult)
	}
	(*cm.SimilarityResults)[result.GetAlgorithm()][result.GetString2()] = result
}
