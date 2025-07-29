package strutil

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
