package strutil

import "fmt"

// ShingleResultsMap is a nested map structure that organizes shingle results by their type and n-gram length.
type ShingleResultsMap map[ShingleResultType]map[int]*ShingleResult

// NewShingleResultsMap initializes and returns a new ShingleResultsMap as an empty nested map.
func NewShingleResultsMap() ShingleResultsMap {
	return make(map[ShingleResultType]map[int]*ShingleResult)
}

// Add inserts a ShingleResult into the ShingleResultsMap, organizing it by type and n-gram length.
func (srm ShingleResultsMap) Add(result ShingleResult) {
	if srm[result.GetType()] == nil {
		srm[result.GetType()] = make(map[int]*ShingleResult)
	}
	srm[result.GetType()][result.GetNgramLength()] = &result
}

// GetCopy creates and returns a deep copy of the ShingleResultsMap,
// preserving its nested structure and data integrity.
func (srm ShingleResultsMap) GetCopy() ShingleResultsMap {
	cloned := NewShingleResultsMap()
	for resType, v := range srm {
		cloned[resType] = make(map[int]*ShingleResult)
		for ngramLength, v2 := range v {
			if v2 != nil {
				resultCopy := *v2
				cloned[resType][ngramLength] = &resultCopy
			}
		}
	}
	return cloned
}

// Get retrieves a ShingleResult from the ShingleResultsMap based on the specified type and n-gram length.
func (srm ShingleResultsMap) Get(resType ShingleResultType, ngramLength int) ShingleResult {
	if srm[resType] == nil {
		return nil
	}
	if srm[resType][ngramLength] == nil {
		return nil
	}
	return CastShingleResult(srm[resType][ngramLength])
}

// GetByType retrieves all ShingleResult instances of the specified ShingleResultType from the ShingleResultsMap.
func (srm ShingleResultsMap) GetByType(resType ShingleResultType) []ShingleResult {
	if srm[resType] == nil {
		return nil
	}
	var results []ShingleResult
	for _, v := range srm[resType] {
		results = append(results, CastShingleResult(v))
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// FilterByType filters the ShingleResultsMap by the specified ShingleResultType and
// returns a new map with matching results.
func (srm ShingleResultsMap) FilterByType(resType ShingleResultType) ShingleResultsMap {
	if srm[resType] == nil {
		return nil
	}
	results := NewShingleResultsMap()
	for _, v := range srm[resType] {
		results.Add(*v)
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// GetByNGramLength retrieves all ShingleResult instances from the map that match the specified n-gram length.
func (srm ShingleResultsMap) GetByNGramLength(ngramLength int) []ShingleResult {
	if len(srm) == 0 || ngramLength < 1 {
		return nil
	}
	var results []ShingleResult
	for _, v := range srm {
		if v[ngramLength] != nil {
			results = append(results, CastShingleResult(v[ngramLength]))
		}
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// FilterByNGramLength filters the map by a specified n-gram length and returns a new map containing the results.
func (srm ShingleResultsMap) FilterByNGramLength(ngramLength int) ShingleResultsMap {
	if len(srm) == 0 || ngramLength < 1 {
		return nil
	}
	results := NewShingleResultsMap()
	for _, v := range srm {
		if v[ngramLength] != nil {
			results.Add(*v[ngramLength])
		}
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// TypeCount returns the total number of ShingleResultType keys in the ShingleResultsMap.
func (srm ShingleResultsMap) TypeCount() int {
	return len(srm)
}

// EntryCount returns the total number of non-nil entries across all nested maps within the ShingleResultsMap.
func (srm ShingleResultsMap) EntryCount() int {
	mapLength := 0
	for _, v := range srm {
		for _, v2 := range v {
			if v2 != nil {
				mapLength++
			}
		}
	}
	return mapLength
}

// IsMatch compares the current ShingleResultsMap to another and returns true
// if they are structurally and value-wise identical.
func (srm ShingleResultsMap) IsMatch(other ShingleResultsMap) bool {
	if srm.TypeCount() != other.TypeCount() || srm.EntryCount() != other.EntryCount() {
		return false
	}
	for resType, v := range srm {
		if other[resType] == nil {
			return false
		}
		for ngramLength, v2 := range v {
			if other[resType][ngramLength] == nil {
				return false
			} else {
				if !CastShingleResult(v2).IsMatch(CastShingleResult(other[resType][ngramLength])) {
					return false
				}
			}
		}
	}
	return true
}

// Print outputs the ShingleResultsMap as a formatted string and optionally includes verbose details if 'v' is true.
func (srm ShingleResultsMap) Print(v bool) ShingleResultsMap {
	fmt.Print(formatShingleResultsMapOutput(srm, v))
	return srm
}

// formatShingleResultsMapOutput formats the output of a ShingleResultsMap, optionally including verbose details.
// It iterates over result types and aggregates their formatted output using formatShingleResultOutput.
func formatShingleResultsMapOutput(srm ShingleResultsMap, verbose bool) string {
	var output string
	for resType, v := range srm {
		output += fmt.Sprintf("Shingle Results for %s\n", resType.String())
		for _, v2 := range v {
			output += formatShingleResultOutput(CastShingleResult(v2), verbose) + "\n"
		}
	}
	return output
}
