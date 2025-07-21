package strutil

import "fmt"

// StringBuilder Type & Core Methods
type StringBuilder struct {
	value          string
	err            error
	comparisonData *ComparisonData
	similarities   *[]SimilarityResult
}

// TODO update to utilize ComparisonManager

// Print outputs the value stored in the StringBuilder if no error exists and returns the StringBuilder itself.
func (sb *StringBuilder) Print() *StringBuilder {
	if sb.err != nil {
		fmt.Printf("%s", sb.err.Error())
		return sb
	}
	fmt.Printf("%s", sb.value)
	return sb
}

// String returns the current string value of the StringBuilder instance.
func (sb *StringBuilder) String() string {
	return sb.value
}

// Error returns the error stored in the StringBuilder instance, or nil if no error is set.
func (sb *StringBuilder) Error() error {
	return sb.err
}

// ComparisonData returns the comparisonData object stored in the StringBuilder.
func (sb *StringBuilder) ComparisonData() *ComparisonData {
	return sb.comparisonData
}

// Build constructs the final string from the StringBuilder and returns it along with any encountered error.
func (sb *StringBuilder) Build() (string, error) {
	if sb.err != nil {
		return "", sb.err
	}
	return sb.value, nil
}

// Result returns the current value of the StringBuilder along with any associated error.
func (sb *StringBuilder) Result() (string, *ComparisonData, error) {
	return sb.value, sb.comparisonData, sb.err
}
