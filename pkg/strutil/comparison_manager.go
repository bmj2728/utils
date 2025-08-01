package strutil

// ComparisonManager is a structure for managing comparison, score, shingle, and LCS results.
type ComparisonManager struct {
	ComparisonResults ComparisonResultsMap
	SimilarityResults SimilarityResultsMap
	ShingleResults    ShingleResultsMap
	LCSResults        LCSResultsMap
}

// NewComparisonManager initializes and returns a new instance of ComparisonManager with empty result maps.
func NewComparisonManager() *ComparisonManager {
	return &ComparisonManager{
		ComparisonResults: NewComparisonResultsMap(),
		SimilarityResults: NewSimilarityResultsMap(),
		ShingleResults:    NewShingleResultsMap(),
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

// GetShingleResultsMap retrieves the ShingleResultsMap from the ComparisonManager instance.
func (cm *ComparisonManager) GetShingleResultsMap() ShingleResultsMap {
	if cm.ShingleResults == nil {
		return nil
	}
	return cm.ShingleResults
}

// CopyShingleResultsMap returns a copy of the ShingleResultsMap if it exists; otherwise, it returns nil.
func (cm *ComparisonManager) CopyShingleResultsMap() ShingleResultsMap {
	if cm.ShingleResults == nil {
		return nil
	}
	return cm.ShingleResults.GetCopy()
}

// AddShingleResult inserts a ShingleResult into the ShingleResults map organized by type and n-gram length.
func (cm *ComparisonManager) AddShingleResult(result ShingleResult) {
	if cm.ShingleResults == nil {
		cm.ShingleResults = NewShingleResultsMap()
	}
	cm.ShingleResults.Add(result)
}

// GetShingleResult retrieves a ShingleResult based on the specified ShingleResultType and n-gram length.
// Returns nil if no matching result exists or if the ShingleResults map is uninitialized.
func (cm *ComparisonManager) GetShingleResult(resType ShingleResultType, ngramLength int) ShingleResult {
	if cm.ShingleResults == nil {
		return nil
	}
	return cm.ShingleResults.Get(resType, ngramLength)
}

// GetShingleResultsByType retrieves all ShingleResults of the specified ShingleResultType from the ShingleResults map.
func (cm *ComparisonManager) GetShingleResultsByType(resType ShingleResultType) []ShingleResult {
	if cm.ShingleResults == nil {
		return nil
	}
	return cm.ShingleResults.GetByType(resType)
}

// FilterShingleResultsByType filters the ShingleResultsMap to include entries matching the specified ShingleResultType.
func (cm *ComparisonManager) FilterShingleResultsByType(resType ShingleResultType) ShingleResultsMap {
	if cm.ShingleResults == nil {
		return nil
	}
	return cm.ShingleResults.FilterByType(resType)
}

// GetShingleResultsByNGramLength retrieves a slice of ShingleResult based on the specified n-gram length.
// Returns nil if no ShingleResults are available.
func (cm *ComparisonManager) GetShingleResultsByNGramLength(ngramLength int) []ShingleResult {
	if cm.ShingleResults == nil {
		return nil
	}
	return cm.ShingleResults.GetByNGramLength(ngramLength)
}

// FilterShingleResultsByNGramLength filters ShingleResultsMap to include entries matching the specified n-gram length.
func (cm *ComparisonManager) FilterShingleResultsByNGramLength(ngramLength int) ShingleResultsMap {
	if cm.ShingleResults == nil {
		return nil
	}
	return cm.ShingleResults.FilterByNGramLength(ngramLength)
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

// GetLCSResult retrieves the LCSResult object based on the provided LCSResultType and input string.
// Returns nil if the LCSResults map is uninitialized.
func (cm *ComparisonManager) GetLCSResult(lcsType LCSResultType, inputStr string) *LCSResult {
	if cm.LCSResults == nil {
		return nil
	}
	return cm.LCSResults.Get(lcsType, inputStr)
}

// GetLCSResultsByType retrieves LCSResults filtered by the specified LCSResultType from the ComparisonManager.
func (cm *ComparisonManager) GetLCSResultsByType(lcsType LCSResultType) []LCSResult {
	if cm.LCSResults == nil {
		return nil
	}
	return cm.LCSResults.GetByType(lcsType)
}

// FilterLCSResultsByType filters the LCS results based on the given LCSResultType and returns the filtered results map.
func (cm *ComparisonManager) FilterLCSResultsByType(lcsType LCSResultType) LCSResultsMap {
	if cm.LCSResults == nil {
		return nil
	}
	return cm.LCSResults.FilterByType(lcsType)
}

// GetLCSResultsByComparisonString retrieves LCS results that match the provided comparison string.
// It returns a slice of LCSResult or nil if no results are available.
func (cm *ComparisonManager) GetLCSResultsByComparisonString(compStr string) []LCSResult {
	if cm.LCSResults == nil {
		return nil
	}
	return cm.LCSResults.GetByComparisonString(compStr)
}

// FilterLCSResultsByComparisonString filters LCS results based on the provided
// comparison string and returns a filtered map.
func (cm *ComparisonManager) FilterLCSResultsByComparisonString(compStr string) LCSResultsMap {
	if cm.LCSResults == nil {
		return nil
	}
	return cm.LCSResults.FilterByComparisonString(compStr)
}
