package strutil

// ComparisonManager is a structure for managing comparison, similarity, shingle, and LCS results.
type ComparisonManager struct {
	ComparisonResults ComparisonResultsMap
	SimilarityResults SimilarityResultsMap
	ShingleData       ShingleResultsMap
	LCSData           LCSResultsMap
}

func NewComparisonManager() *ComparisonManager {
	return &ComparisonManager{
		ComparisonResults: NewComparisonResultsMap(),
		SimilarityResults: NewSimilarityResultsMap(),
		ShingleData:       NewShingleResultsMap(),
		LCSData:           NewLCSResultsMap(),
	}
}
