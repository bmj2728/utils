package strutil

import "fmt"

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

// Get retrieves a ComparisonResult from the map using the specified ComparisonResultType and comparison string.
// The underlying returned objects are pointers to either ComparisonResultInt or ComparisonResultFloat
func (crm ComparisonResultsMap) Get(compResType ComparisonResultType, compStr string) ComparisonResult {
	if crm[compResType] == nil {
		return nil
	}
	if crm[compResType][compStr] == nil {
		return nil
	}
	return CastComparisonResult(crm[compResType][compStr])
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
		results = append(results, CastComparisonResult(v))
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
	for _, v := range crm {
		if v[compStr] != nil {
			results = append(results, CastComparisonResult(v[compStr]))
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

// TypeCount returns the number of distinct ComparisonResultType keys in the ComparisonResultsMap.
func (crm ComparisonResultsMap) TypeCount() int {
	return len(crm)
}

// EntryCount returns the total number of non-nil ComparisonResult entries stored in the nested maps of the structure.
func (crm ComparisonResultsMap) EntryCount() int {
	mapLength := 0
	for _, v := range crm {
		for _, v2 := range v {
			if v2 != nil {
				mapLength++
			}
		}
	}
	return mapLength
}

// IsMatch compares the current ComparisonResultsMap with another map for structural and value equality.
func (crm ComparisonResultsMap) IsMatch(other ComparisonResultsMap) bool {
	// return quickly if counts don't match
	if crm.TypeCount() != other.TypeCount() || crm.EntryCount() != other.EntryCount() {
		return false
	}
	//iterate through the type maps
	for compType, v := range crm {
		//if the type map doesn't exist, we return
		if other[compType] == nil {
			return false
		}
		for compStr, v2 := range v {
			// if we don't have a match on the comp string, return
			if other[compType][compStr] == nil {
				return false
			} else {
				// cast the result to the appropriate type, then run the match
				if !CastComparisonResult(v2).IsMatch(*other[compType][compStr]) {
					return false
				}
			}
		}
	}
	return true
}

// Print iterates through the ComparisonResultsMap and prints the comparison results, optionally in verbose mode.
func (crm ComparisonResultsMap) Print(verbose bool) ComparisonResultsMap {
	fmt.Println(formatComparisonMapOutput(crm, verbose))
	return crm
}

// formatComparisonMapOutput generates a formatted string output for a ComparisonResultsMap with optional verbosity.
func formatComparisonMapOutput(crm ComparisonResultsMap, verbose bool) string {
	var output string
	for compType, v := range crm {
		output += fmt.Sprintf("***Comparison Results for %s***\n\n", compType.String())
		for _, v2 := range v {
			if v2 != nil {
				output += formatComparisonResultOutput(CastComparisonResult(v2), verbose) + "\n"
			}
		}
	}
	return output
}
