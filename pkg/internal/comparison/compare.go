package comparison

import (
	"reflect"
	"slices"
)

// CompareSlices checks if two slices contain the same elements, regardless of order,
// with an option to treat nil slices as equal.
// It requires the elements to be of a comparable type and optionally considers
// both slices nil as equal if `nulls` is true.
// Returns true if the slices are equivalent; otherwise, returns false.
func CompareSlices[T comparable](sl1, sl2 []T, nulls bool) bool {
	if nulls && sl1 == nil && sl2 == nil {
		return true
	}
	if sl1 == nil || sl2 == nil {
		return false
	}
	if reflect.TypeOf(sl1[0]) != reflect.TypeOf(sl2[0]) {
		return false
	}
	if len(sl1) != len(sl2) {
		return false
	}
	remaining := slices.Clone(sl2)
	for _, s := range sl1 {
		if !slices.Contains(remaining, s) {
			return false
		} else {
			for i, w := range remaining {
				if s == w {
					remaining = slices.Delete(remaining, i, i+1)
					break
				}
			}
		}
		continue
	}
	return true
}
