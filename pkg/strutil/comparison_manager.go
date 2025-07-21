package strutil

// ComparisonManager is a structure for managing comparison, similarity, shingle, and LCS results.
type ComparisonManager struct {
	ComparisonResults *map[string]map[string]ComparisonResult
	SimilarityResults *map[Algorithm]map[string]SimilarityResult
	ShingleData       *map[string]map[string]ShingleResult
	LCSData           *map[string]map[string]LCSResult
}

// AddSimilarity adds a similarity result in the ComparisonManager's similarity results map.
// If the algorithm's map has not been initialized, one is created
// If there is an existing result for the Algorithm/comparison word combination, it is replaced
func (cm *ComparisonManager) AddSimilarity(result SimilarityResult) {
	// if there's not a map for this algorithm, add one
	if (*cm.SimilarityResults)[result.GetAlgorithm()] == nil {
		(*cm.SimilarityResults)[result.GetAlgorithm()] = make(map[string]SimilarityResult)
	}
	// add or update similarity for this comparison word
	(*cm.SimilarityResults)[result.GetAlgorithm()][result.GetString2()] = result
}
