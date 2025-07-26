package strutil

// ComparisonResultsMap maps a ComparisonResultType to a nested map of string
// keys and pointers to ComparisonResult objects.
type ComparisonResultsMap map[ComparisonResultType]map[string]*ComparisonResult

// NewComparisonResultsMap initializes and returns an empty ComparisonResultsMap for storing comparison results.
func NewComparisonResultsMap() ComparisonResultsMap {
	return make(map[ComparisonResultType]map[string]*ComparisonResult)
}

// Add inserts a ComparisonResult into the map, organizing it by its type and
// second string, creating sub-maps as needed.
func (crm ComparisonResultsMap) Add(result ComparisonResult) {
	if crm[result.GetType()] == nil {
		crm[result.GetType()] = make(map[string]*ComparisonResult)
	}
	crm[result.GetType()][result.GetString2()] = &result
}

// CastComparisonResult converts a raw ComparisonResult pointer into a specific type based on
// the ComparisonResultType input.
// Returns ComparisonResultInt for integer-based types, ComparisonResultFloat for float-based
// types, or the raw result otherwise.
// Returns nil if the input raw ComparisonResult pointer is nil.
func (crm ComparisonResultsMap) CastComparisonResult(raw *ComparisonResult,
	compResType ComparisonResultType) ComparisonResult {
	if raw == nil {
		return nil
	}
	switch compResType {
	case LevDist, DamLevDist, OSADamLevDist, LCSLength, LCSDist, HammingDist, QGramDist, QGramDistCust:
		casted, ok := (*raw).(*ComparisonResultInt)
		if !ok {
			return *raw
		}
		return casted
	case JaroSim, JaroWinklerSim, JaccardSim, CosineSim, SorensenDiceCo, QGramSim:
		casted, ok := (*raw).(*ComparisonResultFloat)
		if !ok {
			return *raw
		}
		return casted
	default:
		return *raw
	}
}

// GetCopy creates and returns a deep copy of the current ComparisonResultsMap.
func (crm ComparisonResultsMap) GetCopy() ComparisonResultsMap {
	cloned := NewComparisonResultsMap()
	for compType, v := range crm {
		cloned[compType] = make(map[string]*ComparisonResult)
		for compStr, v2 := range v {
			if v2 != nil {
				// Dereference the pointer to create a copy of the value.
				resultCopy := *v2
				// Store the address of the new copy in the cloned map.
				cloned[compType][compStr] = &resultCopy
			}
		}
	}
	return cloned
}

// Get retrieves a ComparisonResult from the map using the specified ComparisonResultType and comparison string key.
func (crm ComparisonResultsMap) Get(compResType ComparisonResultType, compStr string) ComparisonResult {
	if crm[compResType] == nil {
		return nil
	}
	return crm.CastComparisonResult(crm[compResType][compStr], compResType)
}

// 7/25/25 - added filter by methods

// GetByType retrieves a slice of ComparisonResult for the specified ComparisonResultType or
// nil if no results are found.
func (crm ComparisonResultsMap) GetByType(compResType ComparisonResultType) []ComparisonResult {
	if crm[compResType] == nil {
		return nil
	}
	var results []ComparisonResult
	for _, v := range crm[compResType] {
		results = append(results, crm.CastComparisonResult(v, compResType))
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// FilterByType filters the ComparisonResultsMap by the specified ComparisonResultType
// and returns a new map with the results.
func (crm ComparisonResultsMap) FilterByType(compResType ComparisonResultType) ComparisonResultsMap {
	if crm[compResType] == nil {
		return nil
	}
	results := NewComparisonResultsMap()
	for _, v := range crm[compResType] {
		results.Add(*v)
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// GetByComparisonString retrieves a list of ComparisonResult objects associated with the provided comparison string.
func (crm ComparisonResultsMap) GetByComparisonString(compStr string) []ComparisonResult {
	if len(crm) == 0 {
		return nil
	}
	var results []ComparisonResult
	for compType, v := range crm {
		if v[compStr] != nil {
			results = append(results, crm.CastComparisonResult(v[compStr], compType))
		}
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// FilterByComparisonString filters the map, returning a new map with results matching the given comparison string key.
func (crm ComparisonResultsMap) FilterByComparisonString(compStr string) ComparisonResultsMap {
	if len(crm) == 0 {
		return nil
	}
	results := NewComparisonResultsMap()
	for _, v := range crm {
		if v[compStr] != nil {
			results.Add(*v[compStr])
		}
	}
	if len(results) == 0 {
		return nil
	}
	return results
}
