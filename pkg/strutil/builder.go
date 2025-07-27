package strutil

import "fmt"

// StringBuilder Type & Core Methods
type StringBuilder struct {
	value             string
	err               error
	comparisonManager *ComparisonManager
}

// Print outputs the value stored in the StringBuilder if no error exists and returns the StringBuilder itself.
func (sb *StringBuilder) Print() *StringBuilder {
	if sb.err != nil {
		fmt.Printf("%s", sb.formatOutput())
		return sb
	}
	fmt.Printf("%s", sb.formatOutput())
	return sb
}

func (sb *StringBuilder) formatOutput() string {
	if sb.err != nil {
		return sb.err.Error()
	}
	return sb.value
}

// String returns the current string value of the StringBuilder instance.
func (sb *StringBuilder) String() string {
	return sb.value
}

// Error returns the error stored in the StringBuilder instance, or nil if no error is set.
func (sb *StringBuilder) Error() error {
	return sb.err
}

// WithComparisonManager initializes a new ComparisonManager if it doesn't already exist and assigns it to the builder.
func (sb *StringBuilder) WithComparisonManager() *StringBuilder {
	if sb.comparisonManager == nil {
		sb.comparisonManager = NewComparisonManager()
	}
	return sb
}

// ComparisonManager returns the associated ComparisonManager instance of the StringBuilder.
func (sb *StringBuilder) ComparisonManager() *ComparisonManager {
	return sb.comparisonManager
}

// Build constructs the final string from the StringBuilder and returns it along with any encountered error.
func (sb *StringBuilder) Build() (string, error) {
	if sb.err != nil {
		return "", sb.err
	}
	return sb.value, nil
}

// Result returns the current value of the StringBuilder along with any associated error.
func (sb *StringBuilder) Result() (string, *ComparisonManager, error) {
	return sb.value, sb.comparisonManager, sb.err
}

// SetValue sets the value of the StringBuilder to the provided string and returns the updated StringBuilder instance.
func (sb *StringBuilder) SetValue(value string) *StringBuilder {
	sb.value = value
	return sb
}

// SetError sets the error field in the StringBuilder and returns the updated StringBuilder instance.
func (sb *StringBuilder) SetError(err error) *StringBuilder {
	sb.err = err
	return sb
}
