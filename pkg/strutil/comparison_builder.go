package strutil

// CompareStringBuilderSlices compares two slices of StringBuilder for equality,
// optionally allowing nil slices to be considered equal.
// The order of elements in the slices does not affect the comparisonData,
// and the 'nulls' flag determines nil-handling behavior.
func CompareStringBuilderSlices(a, b []StringBuilder, nulls bool) bool {
	return compareStringBuilderSlices(a, b, nulls)
}
