package strutil

// LCSResultsMap is a map that organizes LCSResult instances by their LCSResultType and a string identifier.
type LCSResultsMap map[LCSResultType]map[string]*LCSResult

// NewLCSResultsMap initializes and returns a new LCSResultsMap as an empty nested map.
func NewLCSResultsMap() LCSResultsMap {
	return make(map[LCSResultType]map[string]*LCSResult)
}

// Add inserts an LCSResult into the LCSResultsMap organized by result type and the comparison input string.
func (lrm LCSResultsMap) Add(result LCSResult) {
	if lrm[result.GetType()] == nil {
		lrm[result.GetType()] = make(map[string]*LCSResult)
	}
	lrm[result.GetType()][result.GetString2()] = &result
}

//Add
//GetCopy
//Get
//GetByType
//FilterByType
//GetByComparisonString
//FilterByComparisonString
//TypeCount
//EntryCount
//IsMatch
//Print
//formatLCSResultsMapOutput

// GetCopy creates and returns a deep copy of the LCSResultsMap, duplicating all
// nested maps and their LCSResult values.
func (lrm LCSResultsMap) GetCopy() LCSResultsMap {
	cloned := NewLCSResultsMap()
	for resType, results := range lrm {
		cloned[resType] = make(map[string]*LCSResult)
		for compStr, result := range results {
			if result != nil {
				resultCopy := *result
				cloned[resType][compStr] = &resultCopy
			}
		}
	}
	return cloned
}

// Get retrieves the LCSResult for the given LCSResultType and comparison string from the LCSResultsMap.
// Returns nil if no result is found for the specified type or string.
func (lrm LCSResultsMap) Get(resType LCSResultType, compStr string) *LCSResult {
	if lrm[resType] == nil {
		return nil
	}
	if lrm[resType][compStr] == nil {
		return nil
	}
	return lrm[resType][compStr]
}

// GetByType returns a slice of LCSResult objects from the LCSResultsMap based on their LCSResultType
func (lrm LCSResultsMap) GetByType(resType LCSResultType) []LCSResult {
	if lrm[resType] == nil {
		return nil
	}
	var results []LCSResult
	for _, v := range lrm[resType] {
		results = append(results, *v)
	}
	return results
}
