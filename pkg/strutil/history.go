package strutil

// StringHistory represents a collection of string values used to track the history of
// string changes or operations.
type StringHistory []string

// NewStringHistory creates and returns a new, empty StringHistory instance.
func NewStringHistory() *StringHistory {
	sh := make(StringHistory, 0)
	return &sh
}

// Add appends a string value to the StringHistory collection, updating the history with the provided string.
func (sh *StringHistory) Add(s string) {
	*sh = append(*sh, s)
}
