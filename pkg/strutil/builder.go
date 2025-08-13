package strutil

import (
	"fmt"

	"github.com/bmj2728/utils/pkg/internal/errors"
)

// StringBuilder Type & Core Methods
type StringBuilder struct {
	value             string
	err               error
	originalValue     string
	comparisonManager *ComparisonManager
	history           *StringHistory
}

// TODO builder logic to be updated re: error handling
// fatal errors: transforms cause undefined behavior if incomplete and security functions
// non-fatal errors: other errors e.g. Hamming Dist on unequal lengths

// Print outputs the value stored in the StringBuilder if no error exists and returns the StringBuilder itself.
func (sb *StringBuilder) Print() *StringBuilder {
	if sb.err != nil {
		fmt.Printf("%s\n", sb.formatOutput())
		return sb
	}
	fmt.Printf("%s\n", sb.formatOutput())
	return sb
}

func (sb *StringBuilder) formatOutput() string {
	if !sb.shouldContinueProcessing() {
		return sb.err.Error()
	}
	return sb.value
}

// String returns the current string value of the StringBuilder instance.
func (sb *StringBuilder) String() string {
	return sb.value
}

func (sb *StringBuilder) GetOriginalValue() string {
	return sb.originalValue
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

// WithHistory initializes the history for the StringBuilder if it does not already exist and returns the instance.
func (sb *StringBuilder) WithHistory(limit int) *StringBuilder {
	if sb.history == nil {
		sb.history = NewStringHistory(limit)
	}
	sb.updateHistory(sb.value)
	return sb
}

// GetHistory returns the StringHistory associated with the StringBuilder, which tracks all string modifications.
func (sb *StringBuilder) GetHistory() *StringHistory {
	return sb.history
}

// updateHistory appends a string to the history if it exists and returns the updated StringBuilder instance.
func (sb *StringBuilder) updateHistory(s string) *StringBuilder {
	if sb.history != nil {
		sb.history.Add(s)
	}
	return sb
}

// GetComparisonManager returns the associated ComparisonManager instance of the StringBuilder.
func (sb *StringBuilder) GetComparisonManager() *ComparisonManager {
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
func (sb *StringBuilder) Result() (string, error) {
	return sb.value, sb.err
}

// setValue sets the value of the StringBuilder to the provided string and returns the updated StringBuilder instance.
func (sb *StringBuilder) setValue(value string) *StringBuilder {
	sb.value = value
	if sb.history != nil {
		sb.updateHistory(value)
	}
	return sb
}

// setError sets the error field in the StringBuilder and returns the updated StringBuilder instance.
func (sb *StringBuilder) setError(err error, fatal bool) *StringBuilder {
	sb.err = err
	if fatal {
		sb.setValue("")
	}
	return sb
}

// shouldContinueProcessing determines whether processing should continue based on the state of the error and the value.
func (sb *StringBuilder) shouldContinueProcessing() bool {
	if sb.err != nil {
		if sb.value == "" {
			return false
		}
	}
	return true
}

// RevertToOriginal restores the StringBuilder value to its initial
// state based on its history or sets an error if unresolved.
func (sb *StringBuilder) RevertToOriginal() *StringBuilder {
	sb.value = sb.originalValue
	return sb
}

// RevertToPrevious restores the StringBuilder's value to the previous entry in history or sets an error if unavailable.
func (sb *StringBuilder) RevertToPrevious() *StringBuilder {
	if sb.history != nil {
		prev, err := sb.history.GetPreviousValue()
		if err == nil {
			// this throws an error if there is only an original value
			sb.value = prev
			(*sb.history).transforms = (*sb.history).transforms[:sb.history.Len()-1]
		} else {
			// fatal - reversion failed
			sb.setError(err, true)
		}
	} else {
		sb.setError(errors.ErrHistoryNotInitialized, false)
	}
	return sb
}

// RevertToIndex reverts the StringBuilder's value to the state stored at the specified history index.
// Returns an error if the index is invalid or if the history is not initialized.
// Sets an error in the StringBuilder if issues occur during the operation.
// Invalid indexes will result in a fatal error
func (sb *StringBuilder) RevertToIndex(index int) *StringBuilder {
	if sb.history != nil {
		if index < 0 {
			sb.setError(errors.ErrInvalidHistoryIndex, true)
			return sb
		}
		ind, err := sb.history.GetByIndex(index)
		if err == nil {
			// throws error when invalid index
			sb.value = ind
			(*sb.history).transforms = (*sb.history).transforms[:index+1]
		} else {
			// fatal - reversion has failed
			sb.setError(err, true)
		}
	} else {
		sb.setError(errors.ErrHistoryNotInitialized, false)
	}
	return sb
}

// RevertWithFunction reverts the StringBuilder to a specific state based on the
// index returned by the provided function.
// The function receives the current StringBuilder instance and must return an index to revert to.
// If history is not initialized, it sets an error indicating the issue and takes no action.
// Returns the modified StringBuilder instance.
func (sb *StringBuilder) RevertWithFunction(fn func(history *StringHistory) int) *StringBuilder {
	if sb.history != nil {
		index := fn(sb.history)
		if index >= 0 {
			sb.RevertToIndex(index)
		} else {
			// fatal when expected revert fails
			sb.setError(errors.ErrInvalidHistoryIndex, true)
		}
	} else {
		sb.setError(errors.ErrHistoryNotInitialized, false)
	}
	return sb
}
