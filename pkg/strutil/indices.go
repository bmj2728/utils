package strutil

// Indices provides a structure for managing a collection of indexed data using an IndexMap.
type Indices struct {
	IdxMap IndexMap
}

// NewIndices initializes and returns a new instance of Indices with a fresh IndexMap.
func NewIndices() *Indices {
	return &Indices{
		IdxMap: NewIndexMap(),
	}
}

// Reset reinitializes the IdxMap field in the Indices struct with a new, empty IndexMap.
// this is useful after transforming a string for reindexing purposes.
func (i *Indices) Reset() {
	i.IdxMap = NewIndexMap()
}
