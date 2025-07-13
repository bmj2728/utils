package strutil

import "errors"

// Fluent StringBuilder API

// StringBuilder Type & Core Methods
type StringBuilder struct {
	value string
	err   error
}

// Constructor

// New creates and returns a new StringBuilder instance initialized with the provided string.
func New(s string) *StringBuilder {
	return &StringBuilder{
		value: s,
	}
}

// Manipulation Methods

// CleanWhitespace removes all whitespace characters from the StringBuilder's value and returns the updated StringBuilder.
func (sb *StringBuilder) CleanWhitespace() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = cleanWhitespace(sb.value)
	return sb
}

// NormalizeWhitespace collapses consecutive whitespace characters into a single space and trims leading and trailing spaces.
func (sb *StringBuilder) NormalizeWhitespace() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = normalizeWhitespace(sb.value)
	return sb
}

// CollapseWhitespace collapses consecutive whitespace characters in the StringBuilder's value into a single space and preserves leading and trailing spaces.
func (sb *StringBuilder) CollapseWhitespace() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = collapseWhitespace(sb.value)
	return sb
}

func (sb *StringBuilder) StripHTML() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) EscapeHTML() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) SanitizeHTML(allowedTags []string) *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) Slugify() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) Truncate(length int, suffix string) *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) RemoveNonPrintable() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) NormalizeUnicode() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) StripAnsi() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) SanitizeFilename() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) SanitizePath() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) SanitizeEmail() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) SanitizeUsername() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) ToSnakeCase() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) ToCamelCase() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) ToKebabCase() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) ToLower() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) ToUpper() *StringBuilder {
	panic("Implement me!")
}

// Validation Methods (can set error)

// RequireEmail validates if the StringBuilder's value is a valid email format, sets an error if invalid, and returns the instance.
func (sb *StringBuilder) RequireEmail() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isValidEmail(sb.value) {
		sb.err = errors.New("invalid email address")
	}
	return sb
}

// RequireURL validates if the StringBuilder's value is a properly formatted URL, sets an error if invalid, and returns the instance.
func (sb *StringBuilder) RequireURL() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isValidUrl(sb.value) {
		sb.err = errors.New("invalid URL")
	}
	return sb
}

// RequireUUID validates whether the StringBuilder's value conforms to a valid UUID format, sets an error if invalid, and returns the instance.
func (sb *StringBuilder) RequireUUID() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isValidUUID(sb.value) {
		sb.err = errors.New("invalid UUID")
	}
	return sb
}

// RequireLength validates that the StringBuilder's value length is within the specified min and max range. Sets an error if invalid.
func (sb *StringBuilder) RequireLength(min, max int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if min < 0 || max < 0 {
		sb.err = errors.New("invalid length range")
		return sb
	} else if min > max {
		sb.err = errors.New("invalid length range")
	} else if !isLengthInRange(sb.value, min, max) {
		sb.err = errors.New("invalid length")
	}
	return sb
}

// RequireNotEmpty ensures the StringBuilder's value is not empty, sets an error if it is, and returns the instance.
func (sb *StringBuilder) RequireNotEmpty() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if isEmpty(sb.value) {
		sb.err = errors.New("empty string")
	}
	return sb
}

// RequireNotEmptyNormalized ensures the StringBuilder's value is not empty after normalizing whitespace, setting an error otherwise.
func (sb *StringBuilder) RequireNotEmptyNormalized() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if isEmptyNormalized(sb.value) {
		sb.err = errors.New("empty string after whitespace normalization")
	}
	return sb
}

// Control Flow

// If conditionally applies the provided function to the StringBuilder if the condition is true and no error exists.
func (sb *StringBuilder) If(condition bool, fn func(*StringBuilder) *StringBuilder) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if condition {
		return fn(sb)
	}
	return sb
}

// Transform applies a custom transformation function to the StringBuilder's value and returns the updated instance.
func (sb *StringBuilder) Transform(fn func(string) string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = fn(sb.value)
	return sb
}

// Terminal Methods

// String returns the current string value of the StringBuilder instance.
func (sb *StringBuilder) String() string {
	return sb.value
}

// Error returns the error stored in the StringBuilder instance, or nil if no error is set.
func (sb *StringBuilder) Error() error {
	return sb.err
}

// Must returns the final string value or panics if an error is present in the StringBuilder instance.
func (sb *StringBuilder) Must() string {
	if sb.err != nil {
		panic(sb.err)
	}
	return sb.value
}

// Result returns the current value of the StringBuilder along with any associated error.
func (sb *StringBuilder) Result() (string, error) {
	return sb.value, sb.err
}
