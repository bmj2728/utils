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

// GetSimilarityResultsMap retrieves the SimilarityResultsMap, containing score
// results organized by algorithm and comparison string.
func (cm *ComparisonManager) GetSimilarityResultsMap() SimilarityResultsMap {
	if cm.SimilarityResults == nil {
		return nil
	}
	return cm.SimilarityResults
}

// CopySimilarityResultsMap returns a copy of the SimilarityResultsMap held by the ComparisonManager instance.
// If the SimilarityResults map is nil, it returns nil.
func (cm *ComparisonManager) CopySimilarityResultsMap() SimilarityResultsMap {
	if cm.SimilarityResults == nil {
		return nil
	}
	return cm.SimilarityResults.GetCopy()
}

// AddSimilarityResult adds a given SimilarityResult to the SimilarityResultsMap of the ComparisonManager instance.
func (cm *ComparisonManager) AddSimilarityResult(result SimilarityResult) {
	if cm.SimilarityResults == nil {
		cm.SimilarityResults = NewSimilarityResultsMap()
	}
	cm.SimilarityResults.Add(result)
}

// GetSimilarityResult retrieves a SimilarityResult for the specified algorithm and comparison string.
// Returns nil if no results are found.
func (cm *ComparisonManager) GetSimilarityResult(algo Algorithm, compStr string) *SimilarityResult {
	if cm.SimilarityResults == nil {
		return nil
	}
	return cm.SimilarityResults.Get(algo, compStr)
}

// GetSimilarityResultsByType retrieves all SimilarityResults corresponding to the specified Algorithm type.
func (cm *ComparisonManager) GetSimilarityResultsByType(algo Algorithm) []SimilarityResult {
	if cm.SimilarityResults == nil {
		return nil
	}
	return cm.SimilarityResults.GetByType(algo)
}

// FilterSimilarityResultsByType filters the SimilarityResultsMap to include entries
// matching the specified algorithm type.
func (cm *ComparisonManager) FilterSimilarityResultsByType(algo Algorithm) SimilarityResultsMap {
	if cm.SimilarityResults == nil {
		return nil
	}
	return cm.SimilarityResults.FilterByType(algo)
}

// GetSimilarityResultsByComparisonString retrieves a slice of SimilarityResult associated with
// the given comparison string.
func (cm *ComparisonManager) GetSimilarityResultsByComparisonString(compStr string) []SimilarityResult {
	if cm.SimilarityResults == nil {
		return nil
	}
	return cm.SimilarityResults.GetByComparisonString(compStr)
}

// FilterSimilarityResultsByComparisonString filters the SimilarityResultsMap based on a given comparison string.
func (cm *ComparisonManager) FilterSimilarityResultsByComparisonString(compStr string) SimilarityResultsMap {
	if cm.SimilarityResults == nil {
		return nil
	}
	return cm.SimilarityResults.FilterByComparisonString(compStr)
}

// Shingle Data

// GetShingleData retrieves the ShingleResultsMap from the ComparisonManager instance.
func (cm *ComparisonManager) GetShingleData() ShingleResultsMap {
	if cm.ShingleData == nil {
		return nil
	}
	return cm.ShingleData
}

// CopyShingleData returns a copy of the ShingleResultsMap if it exists; otherwise, it returns nil.
func (cm *ComparisonManager) CopyShingleData() ShingleResultsMap {
	if cm.ShingleData == nil {
		return nil
	}
	return cm.ShingleData.GetCopy()
}

// AddShingleResult inserts a ShingleResult into the ShingleData map organized by type and n-gram length.
func (cm *ComparisonManager) AddShingleResult(result ShingleResult) {
	if cm.ShingleData == nil {
		cm.ShingleData = NewShingleResultsMap()
	}
	cm.ShingleData.Add(result)
}

// LCS Data

// GetLCSResultsMap retrieves the LCSResultsMap containing LCS results organized by type and input string.
func (cm *ComparisonManager) GetLCSResultsMap() LCSResultsMap {
	if cm.LCSResults == nil {
		return nil
	}
	return cm.LCSResults
}

// CopyLCSResultsMap returns a copy of the LCSResultsMap from the ComparisonManager,
// or nil if LCSResults is uninitialized.
func (cm *ComparisonManager) CopyLCSResultsMap() LCSResultsMap {
	if cm.LCSResults == nil {
		return nil
	}
	return cm.LCSResults.GetCopy()
}

// AddLCSResult adds an LCSResult to the LCSResults map in the ComparisonManager,
// organizing it by type and input string.
func (cm *ComparisonManager) AddLCSResult(result LCSResult) {
	if cm.LCSResults == nil {
		cm.LCSResults = NewLCSResultsMap()
	}
	cm.LCSResults.Add(result)
}
