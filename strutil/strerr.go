package strutil

const (

	// ErrInvalidEmail indicates that the input string is not in a valid email address format.
	ErrInvalidEmail = "invalid email address"

	// ErrInvalidURL indicates that the provided string is not a valid URL.
	ErrInvalidURL = "invalid URL"

	// ErrInvalidUUID represents an error message indicating that the provided value is not a valid UUID format.
	ErrInvalidUUID = "invalid UUID"

	// ErrInvalidLengthRange signifies that the length range provided is invalid, such as when min is greater than max or negative.
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
)
