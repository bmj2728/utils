package strutil

// ComparisonManager is a structure for managing comparison, similarity, shingle, and LCS results.
type ComparisonManager struct {
	ComparisonResults ComparisonResultsMap
	SimilarityResults SimilarityResultsMap
	ShingleData       ShingleResultsMap
	LCSData           LCSResultsMap
}

// NewComparisonManager initializes and returns a new instance of ComparisonManager with empty result maps.
func NewComparisonManager() *ComparisonManager {
	return &ComparisonManager{
		ComparisonResults: NewComparisonResultsMap(),
		SimilarityResults: NewSimilarityResultsMap(),
		ShingleData:       NewShingleResultsMap(),
		LCSData:           NewLCSResultsMap(),
	}
}

// AddComparisonResult inserts a ComparisonResult into the ComparisonResultsMap, organizing it by type and first string.
func (cm *ComparisonManager) AddComparisonResult(result ComparisonResult) {
	cm.ComparisonResults.Add(result)
}

// AddSimilarityResult adds a given SimilarityResult to the SimilarityResultsMap of the ComparisonManager instance.
func (cm *ComparisonManager) AddSimilarityResult(result SimilarityResult) {
	cm.SimilarityResults.Add(result)
}

// AddShingleResult inserts a ShingleResult into the ShingleData map organized by type and n-gram length.
func (cm *ComparisonManager) AddShingleResult(result ShingleResult) {
	cm.ShingleData.Add(result)
}

// AddLCSResult adds an LCSResult to the LCSData map in the ComparisonManager, organizing it by type and input string.
func (cm *ComparisonManager) AddLCSResult(result LCSResult) {
	cm.LCSData.Add(result)
}
