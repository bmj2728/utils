package strutil

import (
	"github.com/bmj2728/utils/pkg/internal/errors"
)

// RequireEmail validates if the StringBuilder's value is a valid email format,
// sets an error if invalid, and returns the instance.
func (sb *StringBuilder) RequireEmail() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !isValidEmail(sb.value) {
		return sb.setError(errors.ErrInvalidEmail, true)
	}
	return sb
}

// RequireDomain ensures that the value of the StringBuilder is a valid domain, setting an error if validation fails.
func (sb *StringBuilder) RequireDomain() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !isValidDomain(sb.value) {
		return sb.setError(errors.ErrInvalidDomain, true)
	}
	return sb
}

// RequireURL validates if the StringBuilder's value is a properly formatted URL,
// sets an error if invalid, and returns the instance.
func (sb *StringBuilder) RequireURL() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !isValidURL(sb.value) {
		return sb.setError(errors.ErrInvalidURL, true)
	}
	return sb
}

// RequireUUID validates whether the StringBuilder's value conforms to a valid UUID format,
// sets an error if invalid, and returns the instance.
func (sb *StringBuilder) RequireUUID() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !isValidUUID(sb.value) {
		return sb.setError(errors.ErrInvalidUUID, true)
	}
	return sb
}

// RequireLength validates that the StringBuilder's value length is within the specified min and max range.
// Sets an error if invalid.
func (sb *StringBuilder) RequireLength(min, max int) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if min < 0 || max < 0 {
		return sb.setError(errors.ErrInvalidLengthRange, true)

	} else if min > max {
		return sb.setError(errors.ErrInvalidLengthRange, true)
	} else if !isLengthInRange(sb.value, min, max) {
		return sb.setError(errors.ErrInvalidLength, true)
	}
	return sb
}

// RequireNotEmpty ensures the StringBuilder's value is not empty, sets an error if it is, and returns the instance.
func (sb *StringBuilder) RequireNotEmpty() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if isEmpty(sb.value) {
		return sb.setError(errors.ErrInvalidEmpty, true)
	}
	return sb
}

// RequireNotEmptyNormalized ensures the StringBuilder's value is not empty after normalizing whitespace,
// setting an error otherwise.
func (sb *StringBuilder) RequireNotEmptyNormalized() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if isEmptyNormalized(sb.value) {
		return sb.setError(errors.ErrInvalidEmptyAfterNormalization, true)
	}
	return sb
}

// RequireAlphaNumeric ensures the StringBuilder's value contains only alphanumeric characters,
// setting an error if invalid.
func (sb *StringBuilder) RequireAlphaNumeric() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !isAlphaNumericString(sb.value) {
		return sb.setError(errors.ErrInvalidNotAlphaNumeric, true)
	}
	return sb
}

// RequireAlpha ensures the StringBuilder's value contains only alphabetic characters, setting an error if invalid.
func (sb *StringBuilder) RequireAlpha() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !isAlphaString(sb.value) {
		return sb.setError(errors.ErrInvalidNotAlpha, true)
	}
	return sb
}

// RequireNormalizedUnicode ensures the string is normalized according to the specified Unicode normalization format.
// If not, it sets an error state in the StringBuilder and returns itself for chaining.
func (sb *StringBuilder) RequireNormalizedUnicode(format NormalizationFormat) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !isNormalizedUnicode(sb.value, format) {
		return sb.setError(errors.ErrNotNormalizedUnicode, true)
	}
	return sb
}
