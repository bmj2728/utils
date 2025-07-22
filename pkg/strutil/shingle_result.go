package strutil

import "fmt"

// ShingleResultType is an enumerated type used to represent the type of shingle result, such as map or slice.
type ShingleResultType int

// String returns the string representation of a ShingleResultType by looking it up in the ShingleResultTypeMap.
func (s ShingleResultType) String() string {
	return ShingleResultTypeMap[s]
}

// ShinglesMap represents a result type where shingles are stored in a map.
// ShinglesSlice represents a result type where shingles are stored in a slice.
const (
	ShinglesMap ShingleResultType = iota
	ShinglesSlice
)

// ShingleResultTypeMap maps ShingleResultType constants to their corresponding descriptive string representations.
var ShingleResultTypeMap = map[ShingleResultType]string{
	ShinglesMap:   "Shingle Map",
	ShinglesSlice: "Shingle Slice",
}

// ShingleResult defines an interface for managing and retrieving shingle-related results and their metadata.
// GetType retrieves the shingle result type, indicating the structure of the result such as map or slice.
// GetInput retrieves the input string used to generate the shingle result.
// GetNgramLength retrieves the n-gram length used to compute the shingles.
// Error returns the error encountered during the shingle computation, if any.
// Print outputs the details of the shingle result or error based on the verbosity flag 'v'.
type ShingleResult interface {
	GetType() ShingleResultType
	GetTypeName() string
	GetInput() string
	GetNgramLength() int
	Error() error
	Print(v bool)
}

// ShingleSliceResult represents the result of a shingle operation stored as a slice, encapsulating related metadata.
type ShingleSliceResult struct {
	resultType ShingleResultType
	input      string
	ngram      int
	shingles   *[]string
	err        error
}

// NewShingleResultSlice initializes and returns a pointer to a ShingleSliceResult with provided parameters.
func NewShingleResultSlice(resultType ShingleResultType,
	input string,
	ngram int,
	shingles *[]string,
	err error) *ShingleSliceResult {
	return &ShingleSliceResult{
		resultType: resultType,
		input:      input,
		ngram:      ngram,
		shingles:   shingles,
		err:        err,
	}
}

// GetType returns the ShingleResultType associated with a ShingleSliceResult instance.
func (s ShingleSliceResult) GetType() ShingleResultType {
	return s.resultType
}

// GetTypeName returns the string representation of the result type associated with the ShingleSliceResult instance.
func (s ShingleSliceResult) GetTypeName() string {
	return s.resultType.String()
}

// GetInput returns the input string associated with the ShingleSliceResult.
func (s ShingleSliceResult) GetInput() string {
	return s.input
}

// GetNgramLength returns the n-gram length associated with the ShingleSliceResult instance.
func (s ShingleSliceResult) GetNgramLength() int {
	return s.ngram
}

// GetShinglesSlice returns a pointer to the slice of shingles contained in the ShingleSliceResult.
func (s ShingleSliceResult) GetShinglesSlice() *[]string {
	return s.shingles
}

// Error returns the error associated with the ShingleSliceResult, if any.
func (s ShingleSliceResult) Error() error {
	return s.err
}

// Print outputs shingle data or error information based on the verbose flag.
func (s ShingleSliceResult) Print(v bool) {
	if v {
		if s.err != nil {
			fmt.Printf("Error processing %s\nInput: %s\nN-Gram Length: %d\nError: %s\n",
				ShingleResultTypeMap[s.resultType], s.input, s.ngram, s.err.Error())
			return
		} else {
			fmt.Printf("Input: %s\nN-Gram Length: %d\nShingles:\n",
				s.input, s.ngram)
			for _, shingle := range *s.shingles {
				fmt.Printf("%s\n", shingle)
			}
			return
		}
	} else {
		if s.err != nil {
			fmt.Printf("%s Error: %s\n",
				ShingleResultTypeMap[s.resultType], s.err.Error())
			return
		} else {
			fmt.Printf("%s:\n", ShingleResultTypeMap[s.resultType])
			for _, shingle := range *s.shingles {
				fmt.Printf("%s\n", shingle)
			}
			return
		}
	}
}

// ShingleMapResult is a struct that holds the results of generating shingles, including metadata and possible errors.
// resultType defines the type of shingle result, e.g., map or slice.
// input holds the original string input used for generating shingles.
// ngram specifies the length of n-grams used in shingle generation.
// shingles is a pointer to a map containing shingles as keys and their frequencies as values.
// err represents a potential error encountered during shingle generation.
type ShingleMapResult struct {
	resultType ShingleResultType
	input      string
	ngram      int
	shingles   *map[string]int
	err        error
}

// NewShingleResultMap creates and returns a new instance of ShingleMapResult with the provided parameters.
func NewShingleResultMap(resultType ShingleResultType,
	input string,
	ngram int,
	shingles *map[string]int,
	err error) *ShingleMapResult {
	return &ShingleMapResult{
		resultType: resultType,
		input:      input,
		ngram:      ngram,
		shingles:   shingles,
		err:        err,
	}
}

// GetType returns the type of the shingle result as a ShingleResultType.
func (s ShingleMapResult) GetType() ShingleResultType {
	return s.resultType
}

// GetTypeName returns the string representation of the result type stored in the ShingleMapResult instance.
func (s ShingleMapResult) GetTypeName() string {
	return s.resultType.String()
}

// GetInput returns the input string associated with the ShingleMapResult.
func (s ShingleMapResult) GetInput() string {
	return s.input
}

// GetNgramLength returns the n-gram length associated with the ShingleMapResult instance.
func (s ShingleMapResult) GetNgramLength() int {
	return s.ngram
}

// GetShinglesMap returns a pointer to the map of shingles and their corresponding counts for the given input.
func (s ShingleMapResult) GetShinglesMap() *map[string]int {
	return s.shingles
}

// Error returns the error associated with the ShingleMapResult, if any.
func (s ShingleMapResult) Error() error {
	return s.err
}

// Print outputs the shingle result information based on the verbose flag.
// Handles errors and displays shingles if present.
func (s ShingleMapResult) Print(v bool) {
	if v {
		if s.err != nil {
			fmt.Printf("Error processing %s\nInput: %s\nN-Gram Length: %d\nError: %s\n",
				ShingleResultTypeMap[s.resultType], s.input, s.ngram, s.err.Error())
			return
		} else {
			fmt.Printf("Input: %s\nN-Gram Length: %d\nShingles:\n",
				s.input, s.ngram)
			for word, length := range *s.shingles {
				fmt.Printf("%s:%d\n", word, length)
			}
			return
		}
	} else {
		if s.err != nil {
			fmt.Printf("%s Error: %s\n",
				ShingleResultTypeMap[s.resultType], s.err.Error())
			return
		} else {
			fmt.Printf("%s:\n", ShingleResultTypeMap[s.resultType])
			for word, length := range *s.shingles {
				fmt.Printf("%s:%d\n", word, length)
			}
			return
		}
	}
}

// ShingleResultsMap is a nested map structure that organizes shingle results by their type and n-gram length.
type ShingleResultsMap map[ShingleResultType]map[int]*ShingleResult

// NewShingleResultsMap initializes and returns a new ShingleResultsMap as an empty nested map.
func NewShingleResultsMap() ShingleResultsMap {
	return make(map[ShingleResultType]map[int]*ShingleResult)
}

// Add inserts a ShingleResult into the ShingleResultsMap, organizing it by type and n-gram length.
func (srm ShingleResultsMap) Add(result ShingleResult) {
	if srm[result.GetType()] == nil {
		srm[result.GetType()] = make(map[int]*ShingleResult)
	}
	srm[result.GetType()][result.GetNgramLength()] = &result
}
