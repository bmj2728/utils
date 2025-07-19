package strutil

import "fmt"

// StringBuilder Type & Core Methods
type StringBuilder struct {
	value      string
	err        error
	comparison *EdLibData
}

// Print outputs the value stored in the StringBuilder if no error exists and returns the StringBuilder itself.
func (sb *StringBuilder) Print() *StringBuilder {
	if sb.err != nil {
		fmt.Printf("%s", sb.err)
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

func (sb *StringBuilder) Comparison() *EdLibData {
	return sb.comparison
}

// Must returns the constructed string or an error if one occurred during the building process.
func (sb *StringBuilder) Must() (string, error) {
	if sb.err != nil {
		return "", sb.err
	}
	return sb.value, nil
}

// Result returns the current value of the StringBuilder along with any associated error.
func (sb *StringBuilder) Result() (string, *EdLibData, error) {
	return sb.value, sb.comparison, sb.err
}
