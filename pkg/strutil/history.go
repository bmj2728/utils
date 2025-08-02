package strutil

import "fmt"

// StringHistory represents a collection of string values used to track the history of
// string changes or operations.
/*
[
"Hello
]
*/
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

// Len returns the number of items in the StringHistory collection.
func (sh *StringHistory) Len() int {
	return len(*sh)
}

// GetOriginalValue returns the first string value in the StringHistory, representing its original value.
func (sh *StringHistory) GetOriginalValue() (string, error) {
	if sh.Len() < 1 {
		return "", ErrHistoryIsEmpty
	}
	return (*sh)[0], nil
}

// GetPreviousValue returns the second-to-last string value from the StringHistory collection.
func (sh *StringHistory) GetPreviousValue() (string, error) {
	if sh.Len() < 2 {
		return "", ErrInvalidHistoryIndex
	}
	return (*sh)[sh.Len()-2], nil
}

// GetByIndex retrieves the string at the specified index from the StringHistory collection.
// Returns an error if the index is out of bounds.
func (sh *StringHistory) GetByIndex(index int) (string, error) {
	if index < 0 || index >= len(*sh) {
		return "", ErrInvalidHistoryIndex
	}
	return (*sh)[index], nil
}

// formatHistoryOutput formats the StringHistory into a string representation,
// with verbosity controlled by the verbose flag.
func formatHistoryOutput(history StringHistory, verbose bool) string {
	output := ""
	if verbose {
		output += fmt.Sprintf("\nHistory: \n")
		for seq, str := range history {
			output += fmt.Sprintf("%d: %s\n", seq+1, str)
		}
	} else {
		output += fmt.Sprintf("\nHistory: \n")
		for i, str := range history {
			if i != len(history)-1 {
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
