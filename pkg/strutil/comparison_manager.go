package strutil

// ComparisonManager is a structure for managing comparison, score, shingle, and LCS results.
type ComparisonManager struct {
	ComparisonResults ComparisonResultsMap
	SimilarityResults SimilarityResultsMap
	ShingleData       ShingleResultsMap
	LCSResults        LCSResultsMap
}

// NewComparisonManager initializes and returns a new instance of ComparisonManager with empty result maps.
func NewComparisonManager() *ComparisonManager {
	return &ComparisonManager{
		ComparisonResults: NewComparisonResultsMap(),
		SimilarityResults: NewSimilarityResultsMap(),
		ShingleData:       NewShingleResultsMap(),
		LCSResults:        NewLCSResultsMap(),
	}
}

// Comparison Results

// GetComparisonResultsMap returns the ComparisonResultsMap from the ComparisonManager instance.
func (cm *ComparisonManager) GetComparisonResultsMap() ComparisonResultsMap {
	if cm.ComparisonResults == nil {
		return nil
	}
	return cm.ComparisonResults
}

// CopyComparisonResultsMap returns a copy of the ComparisonResultsMap from the ComparisonManager instance.
func (cm *ComparisonManager) CopyComparisonResultsMap() ComparisonResultsMap {
	if cm.ComparisonResults == nil {
		return nil
	}
	return cm.ComparisonResults.GetCopy()
}

// AddComparisonResult inserts a ComparisonResult into the ComparisonResultsMap, organizing it by type and first string.
func (cm *ComparisonManager) AddComparisonResult(result ComparisonResult) {
	if cm.ComparisonResults == nil {
		cm.ComparisonResults = NewComparisonResultsMap()
	}
	cm.ComparisonResults.Add(result)
}

// GetComparisonResult retrieves a ComparisonResult from the ComparisonResults map based on the
// provided type and string key.
func (cm *ComparisonManager) GetComparisonResult(compResType ComparisonResultType, compStr string) ComparisonResult {
	if cm.ComparisonResults == nil {
		return nil
	}
	return cm.ComparisonResults.Get(compResType, compStr)
}

// FilterComparisonResultsByType filters and returns a map of ComparisonResults matching the specified type.
func (cm *ComparisonManager) FilterComparisonResultsByType(compType ComparisonResultType) ComparisonResultsMap {
	if cm.ComparisonResults == nil {
		return nil
	}
	return cm.ComparisonResults.FilterByType(compType)
}

// GetComparisonResultsByType retrieves all ComparisonResults of the
// specified ComparisonResultType from the ComparisonResults map.
func (cm *ComparisonManager) GetComparisonResultsByType(compResType ComparisonResultType) []ComparisonResult {
	if cm.ComparisonResults == nil {
		return nil
	}
	return cm.ComparisonResults.GetByType(compResType)
}

// FilterComparisonResultsByComparisonString filters and returns a ComparisonResultsMap containing
// entries matching compStr.
func (cm *ComparisonManager) FilterComparisonResultsByComparisonString(compStr string) ComparisonResultsMap {
	if cm.ComparisonResults == nil {
		return nil
	}
	return cm.ComparisonResults.FilterByComparisonString(compStr)
}

// GetComparisonResultsByString retrieves all ComparisonResult objects
// associated with the given string key from ComparisonResults.
func (cm *ComparisonManager) GetComparisonResultsByString(compStr string) []ComparisonResult {
	if cm.ComparisonResults == nil {
		return nil
	}
	return cm.ComparisonResults.GetByComparisonString(compStr)
}

// Similarity Results

// GetSimilarityResults retrieves the SimilarityResultsMap, containing score
// results organized by algorithm and comparison string.
func (cm *ComparisonManager) GetSimilarityResults() SimilarityResultsMap {
	return cm.SimilarityResults
}

// AddSimilarityResult adds a given SimilarityResult to the SimilarityResultsMap of the ComparisonManager instance.
func (cm *ComparisonManager) AddSimilarityResult(result SimilarityResult) {
	cm.SimilarityResults.Add(result)
}

// Shingle Data

// GetShingleData retrieves the ShingleResultsMap from the ComparisonManager instance.
func (cm *ComparisonManager) GetShingleData() ShingleResultsMap {
	return cm.ShingleData
}

// AddShingleResult inserts a ShingleResult into the ShingleData map organized by type and n-gram length.
func (cm *ComparisonManager) AddShingleResult(result ShingleResult) {
	cm.ShingleData.Add(result)
}

// LCS Data

// GetLCSResults retrieves the LCSResultsMap containing LCS results organized by type and input string.
func (cm *ComparisonManager) GetLCSResults() LCSResultsMap {
	return cm.LCSResults
}

// AddLCSResult adds an LCSResult to the LCSResults map in the ComparisonManager,
// organizing it by type and input string.
func (cm *ComparisonManager) AddLCSResult(result LCSResult) {
	cm.LCSResults.Add(result)
}
