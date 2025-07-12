package errors

import "errors"

var (

	// ErrInvalidEmail indicates that the input string is not in a valid email address format.
	ErrInvalidEmail = errors.New("invalid email address")

	// ErrInvalidURL indicates that the provided string is not a valid URL.
	ErrInvalidURL = errors.New("invalid URL")

	// ErrInvalidUUID represents an error message indicating that the provided value is not a valid UUID format.
	ErrInvalidUUID = errors.New("invalid UUID")

	// ErrInvalidLengthRange signifies that the length range provided is invalid,
	// such as when min is greater than max or negative.
	ErrInvalidLengthRange = errors.New("invalid length range")

	// ErrInvalidLength indicates that the length of the string is outside the allowed range.
	ErrInvalidLength = errors.New("invalid length")

	// ErrInvalidEmpty indicates that the provided string is empty and does not meet the required validation criteria.
	ErrInvalidEmpty = errors.New("empty string")

	// ErrInvalidEmptyAfterNormalization represents an error for strings that become empty after whitespace normalization.
	ErrInvalidEmptyAfterNormalization = errors.New("empty string after whitespace normalization")

	// ErrInvalidNotAlphaNumeric indicates that a string contains non-alphanumeric characters.
	ErrInvalidNotAlphaNumeric = errors.New("string contains non-alphanumeric characters")

	// ErrInvalidNotAlpha indicates that the string contains non-alphabetic characters.
	ErrInvalidNotAlpha = errors.New("string contains non-alphabetic characters")

	// ErrInvalidDomain indicates that the provided string is not a valid domain.
	ErrInvalidDomain = errors.New("invalid domain")

	// ErrLCSBacktrackFailure indicates a failure in the
	// backtracking process to determine the longest common substring.
	ErrLCSBacktrackFailure = errors.New("error backtracking longest common substring")

	// ErrLCSBacktrackAllFailure indicates an error occurred while attempting to
	// backtrack for a slice of common substrings.
	ErrLCSBacktrackAllFailure = errors.New("error backtracking for slice of common substrings")

	// ErrLCSDiffFailure represents an error occurring during the generation of Longest Common Subsequence (LCS) diff.
	ErrLCSDiffFailure = errors.New("error generating lcs diff")

	// ErrHammingDistanceFailure is returned when there is an error while
	// calculating the Hamming distance between two strings.
	ErrHammingDistanceFailure = errors.New("error calculating hamming distance")

	// ErrShingleLengthOutOfRange indicates that the shingle length provided is outside the acceptable range.
	ErrShingleLengthOutOfRange = errors.New("shingle length out of range")

	// ErrUnknownError represents an unknown error, typically used when no specific error or score applies.
	ErrUnknownError = errors.New("no score or error")

	// ErrNoSplitLengthSet indicates that no split length was provided or set for the operation.
	ErrNoSplitLengthSet = errors.New("no split length set")

	// ErrNotNormalizedUnicode indicates that a string is not in a normalized Unicode format.
	ErrNotNormalizedUnicode = errors.New("string is not normalized unicode")

	// ErrNilScore is returned when the expected score is nil but no specific error is associated.
	ErrNilScore = errors.New("score is nil")

	// ErrInvalidHistoryIndex indicates that the provided index is out of bounds for the history collection.
	ErrInvalidHistoryIndex = errors.New("invalid history index")

	// ErrHistoryNotInitialized indicates that the history has not been initialized prior to use.
	ErrHistoryNotInitialized = errors.New("history not initialized")

	// ErrHistoryIsEmpty indicates that a history operation was attempted on an empty history.
	ErrHistoryIsEmpty = errors.New("history is empty")

	// ErrInvalidNgramMap indicates that the provided ngram map is invalid.
	ErrInvalidNgramMap = errors.New("invalid ngram map")

	// ErrPatternNotFound indicates that the specified pattern could not be found.
	ErrPatternNotFound = errors.New("pattern not found")
)

// CompareErrors compares two error values for equality by checking their string representations.
// Returns true if both errors are nil or their messages are the same, otherwise false.
func CompareErrors(err1, err2 error) bool {
	if err1 == nil {
		return err2 == nil
	}
	if err2 == nil {
		return false
	}
	return err1.Error() == err2.Error()
}
