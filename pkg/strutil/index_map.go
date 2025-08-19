package strutil

import "fmt"

// IndexMap is a data structure that maps IndexType to a nested map of strings
// and their corresponding StringIndex pointers.
type IndexMap map[IndexType]map[string]*StringIndex

// NewIndexMap initializes and returns a new, empty IndexMap instance.
func NewIndexMap() IndexMap {
	return make(map[IndexType]map[string]*StringIndex)
}

// Add inserts a StringIndex into the IndexMap under the appropriate IndexType and string key.
func (im IndexMap) Add(index *StringIndex) {
	if im[index.GetIndexType()] == nil {
		im[index.GetIndexType()] = make(map[string]*StringIndex)
	}
	im[index.GetIndexType()][index.GetString()] = index
}

// GetCopy creates and returns a deep copy of the current IndexMap, ensuring all map contents are cloned.
func (im IndexMap) GetCopy() IndexMap {
	cloned := NewIndexMap()
	for indexType, v := range im {
		cloned[indexType] = make(map[string]*StringIndex)
		for str, i := range v {
			// Dereference the pointer to create a copy of the value.
			iCopy := *i
			// Store the address of the new copy in the cloned map.
			cloned[indexType][str] = &iCopy
		}
	}
	return cloned
}

// Get retrieves the StringIndex associated with the given IndexType and string key. Returns nil if no match is found.
func (im IndexMap) Get(indexType IndexType, str string) *StringIndex {
	if im[indexType] == nil {
		return nil
	}
	return im[indexType][str]
}

// GetByIndexType retrieves a slice of StringIndex instances associated with the specified IndexType from the IndexMap.
func (im IndexMap) GetByIndexType(indexType IndexType) []StringIndex {
	if im[indexType] == nil {
		return nil
	}
	var results []StringIndex
	for _, v := range im[indexType] {
		results = append(results, *v)
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// FilterByType filters the IndexMap to include only entries of the specified IndexType.
// Returns a new IndexMap or nil if empty.
func (im IndexMap) FilterByType(indexType IndexType) IndexMap {
	if im[indexType] == nil {
		return nil
	}
	results := NewIndexMap()
	for _, v := range im[indexType] {
		results.Add(v)
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// GetByComparisonString retrieves all StringIndex instances matching the provided comparison
// string across all IndexTypes.
func (im IndexMap) GetByComparisonString(compStr string) []StringIndex {
	if len(im) == 0 {
		return nil
	}
	var results []StringIndex
	for _, v := range im {
		if v[compStr] != nil {
			results = append(results, *v[compStr])
		}
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// FilterByComparisonString filters the IndexMap by a comparison string and returns a
// new IndexMap with matched entries.
func (im IndexMap) FilterByComparisonString(compStr string) IndexMap {
	if len(im) == 0 {
		return nil
	}
	results := NewIndexMap()
	for _, v := range im {
		if v[compStr] != nil {
			results.Add(v[compStr])
		}
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

// TypeCount returns the number of distinct IndexType entries in the IndexMap.
func (im IndexMap) TypeCount() int {
	return len(im)
}

// EntryCount returns the total number of non-nil entries across all inner maps within the IndexMap.
func (im IndexMap) EntryCount() int {
	mapLength := 0
	for _, v := range im {
		for _, v2 := range v {
			if v2 != nil {
				mapLength++
			}
		}
	}
	return mapLength
}

// IsMatch determines if the current IndexMap is equivalent to
// another IndexMap by comparing types and entries recursively.
func (im IndexMap) IsMatch(other IndexMap) bool {
	if im.TypeCount() != other.TypeCount() || im.EntryCount() != other.EntryCount() {
		return false
	}
	for indexType, v := range im {
		if other[indexType] == nil {
			return false
		}
		for str, v2 := range v {
			if other[indexType][str] == nil {
				return false
			} else {
				if !v2.IsMatch(other[indexType][str], true) {
					return false
				}
			}
		}
	}
	return true
}

// Print outputs the IndexMap data to standard output in a formatted manner based on the verbosity flag.
func (im IndexMap) Print(verbose bool) {
	fmt.Println(formatIndexMapOutput(im, verbose))
}

// formatIndexMapOutput generates a formatted string representation of the IndexMap object based on the verbosity flag.
func formatIndexMapOutput(im IndexMap, verbose bool) string {
	var output string
	for indexType, v := range im {
		output += fmt.Sprintf("***Index Map for %s***\n\n", indexType.String())
		for _, v2 := range v {
			if v2 != nil {
				output += formatStringIndexOutput(v2, verbose) + "\n"
			}
		}
	}
	return output
}
