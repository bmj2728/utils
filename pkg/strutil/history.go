package strutil

import (
	"fmt"

	"github.com/bmj2728/utils/pkg/internal/errors"
)

// StringHistory represents a collection of string values used to track the history of
type StringHistory struct {
	transforms []string
	limit      int
}

// NewStringHistory creates and returns a new, empty StringHistory instance.
func NewStringHistory(limit int) *StringHistory {
	return &StringHistory{
		transforms: make([]string, 0, limit),
		limit:      limit,
	}
}

// Add appends a string value to the StringHistory collection, updating the history with the provided string.
func (sh *StringHistory) Add(s string) {
	if sh.Len() >= sh.limit {
		(*sh).transforms = (*sh).transforms[1:sh.limit]
	}
	(*sh).transforms = append((*sh).transforms, s)
}

// Len returns the number of items in the StringHistory collection.
func (sh *StringHistory) Len() int {
	if sh == nil {
		return 0
	}
	return len(sh.transforms)
}

// GetPreviousValue returns the second-to-last string value from the StringHistory collection.
func (sh *StringHistory) GetPreviousValue() (string, error) {
	if sh.Len() < 2 {
		return "", errors.ErrInvalidHistoryIndex
	}
	return (*sh).transforms[sh.Len()-2], nil
}

// GetByIndex retrieves the string at the specified index from the StringHistory collection.
// Returns an error if the index is out of bounds.
func (sh *StringHistory) GetByIndex(index int) (string, error) {
	if index < 0 || index >= sh.Len() {
		return "", errors.ErrInvalidHistoryIndex
	}
	return (*sh).transforms[index], nil
}

// formatHistoryOutput formats the StringHistory into a string representation,
// with verbosity controlled by the verbose flag.
func formatHistoryOutput(history StringHistory, verbose bool) string {
	output := ""
	if verbose {
		output += "\nHistory: \n"
		for seq, str := range history.transforms {
			output += fmt.Sprintf("%d: %s\n", seq+1, str)
		}
	} else {
		output += "\nHistory: \n"
		for i, str := range history.transforms {
			if i != history.Len()-1 {
				output += fmt.Sprintf("%s, ", str)
			} else {
				output += fmt.Sprintf("%s\n", str)
			}
		}
	}
	return output
}

// Print outputs the string history in either verbose or concise format based on the verbose parameter.
func (sh *StringHistory) Print(verbose bool) {
	fmt.Print(formatHistoryOutput(*sh, verbose))
}
