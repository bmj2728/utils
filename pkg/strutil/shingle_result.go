package strutil

import (
	"fmt"
	"maps"

	"github.com/bmj2728/utils/pkg/internal/comparison"
	"github.com/bmj2728/utils/pkg/internal/errors"
)

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
	GetError() error
	IsMatch(other ShingleResult) bool
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

// NewShingleSliceResult initializes and returns a pointer to a ShingleSliceResult with provided parameters.
func NewShingleSliceResult(resultType ShingleResultType,
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
func (s *ShingleSliceResult) GetType() ShingleResultType {
	return s.resultType
}

// GetTypeName returns the string representation of the result type associated with the ShingleSliceResult instance.
func (s *ShingleSliceResult) GetTypeName() string {
	return s.resultType.String()
}

// GetInput returns the input string associated with the ShingleSliceResult.
func (s *ShingleSliceResult) GetInput() string {
	return s.input
}

// GetNgramLength returns the n-gram length associated with the ShingleSliceResult instance.
func (s *ShingleSliceResult) GetNgramLength() int {
	return s.ngram
}

// GetShinglesSlice returns a pointer to the slice of shingles contained in the ShingleSliceResult.
func (s *ShingleSliceResult) GetShinglesSlice() []string {
	if s.shingles == nil {
		return nil
	}
	return *s.shingles
}

// GetError returns the error associated with the ShingleSliceResult, if any.
func (s *ShingleSliceResult) GetError() error {
	return s.err
}

func (s *ShingleSliceResult) IsMatch(other ShingleResult) bool {
	casted, ok := other.(*ShingleSliceResult)
	if !ok {
		return false
	}
	return compareShingleInputFields(s, casted) &&
		comparison.CompareStringSlices(s.GetShinglesSlice(), casted.GetShinglesSlice(), false) &&
		errors.CompareErrors(s.GetError(), casted.GetError())
}

// Print outputs shingle data or error information based on the verbose flag.
func (s *ShingleSliceResult) Print(v bool) {
	fmt.Print(formatShingleResultOutput(s, v))
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
	shingles   map[string]int
	err        error
}

// NewShingleMapResult creates and returns a new instance of ShingleMapResult with the provided parameters.
func NewShingleMapResult(resultType ShingleResultType,
	input string,
	ngram int,
	shingles map[string]int,
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
func (s *ShingleMapResult) GetType() ShingleResultType {
	return s.resultType
}

// GetTypeName returns the string representation of the result type stored in the ShingleMapResult instance.
func (s *ShingleMapResult) GetTypeName() string {
	return s.resultType.String()
}

// GetInput returns the input string associated with the ShingleMapResult.
func (s *ShingleMapResult) GetInput() string {
	return s.input
}

// GetNgramLength returns the n-gram length associated with the ShingleMapResult instance.
func (s *ShingleMapResult) GetNgramLength() int {
	return s.ngram
}

// GetShinglesMap returns a pointer to the map of shingles and their corresponding counts for the given input.
func (s *ShingleMapResult) GetShinglesMap() map[string]int {
	if s.shingles == nil {
		return nil
	}
	return s.shingles
}

// GetError returns the error associated with the ShingleMapResult, if any.
func (s *ShingleMapResult) GetError() error {
	return s.err
}

// IsMatch compares the current ShingleMapResult with another ShingleResult for equality
// based on their fields and content.
func (s *ShingleMapResult) IsMatch(other ShingleResult) bool {
	casted, ok := other.(*ShingleMapResult)
	if !ok {
		return false
	}
	if !compareShingleInputFields(s, casted) {
		return false
	}
	if !errors.CompareErrors(s.GetError(), other.GetError()) {
		return false
	}
	if s.shingles == nil && casted.shingles == nil {
		return true
	}
	if s.shingles == nil || casted.shingles == nil {
		return false
	}
	// this function should work for comparing the resulting shingle maps
	// the map is expected in format [ngram]count, with no duplicate entries
	if !maps.Equal(s.shingles, casted.shingles) {
		return false
	}
	return true
}

// Print outputs the shingle result information based on the verbose flag.
// Handles errors and displays shingles if present.
func (s *ShingleMapResult) Print(v bool) {
	fmt.Print(formatShingleResultOutput(s, v))
}

// Utility Functions

// compareShingleInputFields compares two ShingleResult objects and returns true
// if their type, input, and n-gram length match.
func compareShingleInputFields(s1 ShingleResult, s2 ShingleResult) bool {
	if s1 == nil || s2 == nil {
		return false
	}
	if s1.GetType() != s2.GetType() ||
		s1.GetInput() != s2.GetInput() ||
		s1.GetNgramLength() != s2.GetNgramLength() {
		return false
	}
	return true
}

// CastShingleResult casts a generic ShingleResult to a concrete type like ShingleMapResult
// or ShingleSliceResult based on its type.
// Returns nil if the input is nil or cannot be cast to a valid ShingleResult implementation.
func CastShingleResult(raw *ShingleResult) ShingleResult {
	if raw == nil {
		return nil
	}
	switch (*raw).GetType() {
	case ShinglesMap:
		casted, ok := (*raw).(*ShingleMapResult)
		if !ok {
			return nil
		}
		return casted
	case ShinglesSlice:
		casted, ok := (*raw).(*ShingleSliceResult)
		if !ok {
			return nil
		}
		return casted
	default:
		return nil
	}
}

// formatShingleResultOutput formats the output of a ShingleResult based on its type, verbosity,
// and error state.
func formatShingleResultOutput(result ShingleResult, v bool) string {
	header := ""
	payload := ""
	resType := result.GetTypeName()
	input := result.GetInput()
	ngram := result.GetNgramLength()
	err := result.GetError()
	var shinglesSlice []string
	var shinglesMap map[string]int

	if result.GetError() != nil {
		if v {
			return fmt.Sprintf("Error processing %s (%s/%d):\n%s\n",
				resType, input, ngram, err.Error())
		} else {
			return fmt.Sprintf("%s (%s/%d): %s\n",
				resType, input, ngram, err.Error())
		}
	}

	if v {
		header = fmt.Sprintf("Results for %s\nWord: %s\nN-Gram Length: %d\n",
			resType, input, ngram)
	} else {
		header = fmt.Sprintf("%s (%s/%d):\n",
			resType, input, ngram)
	}

	switch s := result.(type) {
	case *ShingleMapResult:
		shinglesMap = s.GetShinglesMap()
		if v {
			payload = "Shingles Map:\n"
			for k, v := range shinglesMap {
				payload += fmt.Sprintf("%s: %d\n", k, v)
			}
		} else {
			payload = fmt.Sprintf("%d shingles found\n", len(shinglesMap))
		}
	case *ShingleSliceResult:
		shinglesSlice = s.GetShinglesSlice()
		if v {
			payload = "Shingles Slice:\n"
			for _, v := range shinglesSlice {
				payload += fmt.Sprintf("%s\n", v)
			}
		} else {
			payload = fmt.Sprintf("%d shingles found\n", len(shinglesSlice))
		}
	}
	formatted := header + payload
	return formatted
}
