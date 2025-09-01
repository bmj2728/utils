package strutil

import "github.com/bmj2728/utils/pkg/internal/comparison"

// CompareStringSlices compares two slices of strings for equality, with an option
// to treat nil slices as equal if nulls is set to true.
func CompareStringSlices(a, b []string, nulls bool) bool {
	return comparison.CompareSlices(a, b, nulls)
}
