package strutil

import "fmt"

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
	if len(results) == 0 {
		return nil
	}
	return results
}

// FilterByType filters the LCSResultsMap by the specified LCSResultType and returns a new map containing
// only matching results.
func (lrm LCSResultsMap) FilterByType(resType LCSResultType) LCSResultsMap {
	if lrm[resType] == nil {
		return nil
	}
	results := NewLCSResultsMap()
	for _, v := range lrm[resType] {
		results.Add(*v)
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// GetByComparisonString retrieves all LCSResult objects from the map that match the specified comparison string.
// Returns nil if no results are found or the map is empty.
func (lrm LCSResultsMap) GetByComparisonString(compStr string) []LCSResult {
	if len(lrm) == 0 {
		return nil
	}
	var results []LCSResult
	for _, v := range lrm {
		if v[compStr] != nil {
			results = append(results, *v[compStr])
		}
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// FilterByComparisonString filters the LCSResultsMap by a given comparison string and
// returns a new map with matching results.
func (lrm LCSResultsMap) FilterByComparisonString(compStr string) LCSResultsMap {
	if len(lrm) == 0 {
		return nil
	}
	results := NewLCSResultsMap()
	for _, v := range lrm {
		if v[compStr] != nil {
			results.Add(*v[compStr])
		}
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// TypeCount returns the number of LCSResultType keys in the LCSResultsMap.
func (lrm LCSResultsMap) TypeCount() int {
	return len(lrm)
}

// EntryCount returns the total number of non-nil LCSResult entries in the LCSResultsMap.
func (lrm LCSResultsMap) EntryCount() int {
	mapLength := 0
	for _, v := range lrm {
		for _, v2 := range v {
			if v2 != nil {
				mapLength++
			}
		}
	}
	return mapLength
}

// IsMatch compares the current LCSResultsMap with another,
// checking if they have identical structure and matching entries.
func (lrm LCSResultsMap) IsMatch(other LCSResultsMap) bool {
	if lrm.TypeCount() != other.TypeCount() || lrm.EntryCount() != other.EntryCount() {
		return false
	}
	for resType, v := range lrm {
		if other[resType] == nil {
			return false
		}
		for compStr, v2 := range v {
			if other[resType][compStr] == nil {
				return false
			}
			if !v2.IsMatch(other[resType][compStr]) {
				return false
			}
		}
	}
	return true
}

// Print outputs the contents of the LCSResultsMap to the console in a
// formatted manner based on the verbosity flag.
func (lrm LCSResultsMap) Print(v bool) LCSResultsMap {
	fmt.Print(formatLCSResultsMapOutput(lrm, v))
	return lrm
}

// formatLCSResultsMapOutput formats the output of an LCSResultsMap into a string,
// grouped by result type with verbosity control.
func formatLCSResultsMapOutput(lrm LCSResultsMap, verbose bool) string {
	var output string
	for resType, v := range lrm {
		output += fmt.Sprintf("***LCS Results for %s***\n\n", resType.String())
		for _, v2 := range v {
			if v2 != nil {
				output += formatLCSResultOutput(v2, verbose) + "\n"
			}
		}
	}
	return output
}
