package strutil

import "errors"

// TODO ensure consistent use of error vars
// TODO refactor to define these as errors not strings
// TODO calls should foll pattern with these errors joined to call errors
var (

	// ErrInvalidEmail indicates that the input string is not in a valid email address format.
	ErrInvalidEmail = "invalid email address"

	// ErrInvalidURL indicates that the provided string is not a valid URL.
	ErrInvalidURL = "invalid URL"

	// ErrInvalidUUID represents an error message indicating that the provided value is not a valid UUID format.
	ErrInvalidUUID = "invalid UUID"

	// ErrInvalidLengthRange signifies that the length range provided is invalid,
	// such as when min is greater than max or negative.
	ErrInvalidLengthRange = "invalid length range"

	// ErrInvalidLength indicates that the length of the string is outside the allowed range.
	ErrInvalidLength = "invalid length"

	// ErrInvalidEmpty indicates that the provided string is empty and does not meet the required validation criteria.
	ErrInvalidEmpty = "empty string"

	// ErrInvalidEmptyAfterNormalization represents an error for strings that become empty after whitespace normalization.
	ErrInvalidEmptyAfterNormalization = "empty string after whitespace normalization"

	// ErrInvalidNotAlphaNumeric indicates that a string contains non-alphanumeric characters.
	ErrInvalidNotAlphaNumeric = "string contains non-alphanumeric characters"

	// ErrInvalidNotAlpha indicates that the string contains non-alphabetic characters.
	ErrInvalidNotAlpha = "string contains non-alphabetic characters"

	// ErrInvalidDomain indicates that the provided string is not a valid domain.
	ErrInvalidDomain = "invalid domain"

	// ErrLCSBacktrackFailure indicates a failure in the
	// backtracking process to determine the longest common substring.
	ErrLCSBacktrackFailure = "error backtracking longest common substring"

	// ErrLCSBacktrackAllFailure indicates an error occurred while attempting to
	// backtrack for a slice of common substrings.
	ErrLCSBacktrackAllFailure = "error backtracking for slice of common substrings"

	// ErrLCSDiffFailure represents an error occurring during the generation of Longest Common Subsequence (LCS) diff.
	ErrLCSDiffFailure = "error generating lcs diff"

	// ErrHammingDistanceFailure is returned when there is an error while
	// calculating the Hamming distance between two strings.
	ErrHammingDistanceFailure = errors.New("error calculating hamming distance")
)
