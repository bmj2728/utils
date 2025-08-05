package comparison

import "slices"

// CompareStringSlices checks if two slices of strings contain the same elements,
// regardless of order, with optional nil equality.
// If `nulls` is true, nil slices are treated as equal.
// It returns true if slices match, otherwise false.
func CompareStringSlices(s1, s2 []string, nulls bool) bool {
	if nulls && s1 == nil && s2 == nil {
		return true
	}
	if s1 == nil || s2 == nil {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	remaining := slices.Clone(s2)
	for _, s := range s1 {
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
