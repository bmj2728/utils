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
