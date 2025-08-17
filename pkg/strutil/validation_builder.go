package strutil

import (
	"github.com/bmj2728/utils/pkg/internal/errors"
)

// IsEmail checks if the current StringBuilder value is a valid email format and
// returns true if valid, false otherwise.
func (sb *StringBuilder) IsEmail() bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return isEmail(sb.value)
}

// IsDomain checks if the current value of the StringBuilder is a valid domain name
// based on predefined validation rules.
func (sb *StringBuilder) IsDomain() bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return isDomain(sb.value)
}

// IsURL determines if the string value of the StringBuilder is a valid URL with a proper scheme and host.
func (sb *StringBuilder) IsURL() bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return isURL(sb.value)
}

// IsUUID checks if the StringBuilder's value is a valid UUID and returns true if valid, otherwise false.
func (sb *StringBuilder) IsUUID() bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return isUUID(sb.value)
}

// IsLengthInRange checks if the length of the StringBuilder's value is within the specified inclusive range [min, max].
func (sb *StringBuilder) IsLengthInRange(min, max int) bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return isLengthInRange(sb.value, min, max)
}

// IsEmpty determines whether the StringBuilder contains an empty string, returning true if empty, otherwise false.
func (sb *StringBuilder) IsEmpty() bool {
	if !sb.shouldContinueProcessing() {
		return true
	}
	return isEmpty(sb.value)
}

// IsEmptyNormalized determines if the normalized version of
// the StringBuilder's value is empty after trimming whitespace.
func (sb *StringBuilder) IsEmptyNormalized() bool {
	if !sb.shouldContinueProcessing() {
		return true
	}
	return isEmptyNormalized(sb.value)
}

// IsAlphaNumeric determines if the underlying string value consists only of
// alphanumeric characters (letters and digits).
func (sb *StringBuilder) IsAlphaNumeric() bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return isAlphaNumeric(sb.value)
}

// IsAlpha checks if the StringBuilder's current value contains only alphabetic characters.
func (sb *StringBuilder) IsAlpha() bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return isAlpha(sb.value)
}

func (sb *StringBuilder) IsNumeric(strict bool) bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return isNumeric(sb.value, strict)
}

// IsNormalizedUnicode checks if the StringBuilder's current value is normalized
// in the given Unicode normalization format.
func (sb *StringBuilder) IsNormalizedUnicode(format NormalizationFormat) bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return isNormalizedUnicode(sb.value, format)
}

// Contains checks if the specified substring exists within the StringBuilder's value.
// Returns true if found, otherwise false.
func (sb *StringBuilder) Contains(substr string) bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return contains(sb.value, substr)
}

// ContainsIgnoreCase checks if the StringBuilder's value contains the given substring, ignoring case sensitivity.
func (sb *StringBuilder) ContainsIgnoreCase(substr string) bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return containsIgnoreCase(sb.value, substr)
}

// ContainsAny checks if the StringBuilder's value contains any of the strings from the provided slice substrs.
// Returns true if at least one substring is found, otherwise false.
func (sb *StringBuilder) ContainsAny(substrs []string) bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return containsAny(sb.value, substrs)
}

// ContainsAnyIgnoreCase checks if the StringBuilder contains any of the given substrings, ignoring case sensitivity.
// Returns false if processing should not continue or the input is invalid.
func (sb *StringBuilder) ContainsAnyIgnoreCase(substrs []string) bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return containsAnyIgnoreCase(sb.value, substrs)
}

// ContainsAll checks if all strings in the provided slice are present in
// the StringBuilder's value and returns true or false.
func (sb *StringBuilder) ContainsAll(substrs []string) bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return containsAll(sb.value, substrs)
}

// ContainsAllIgnoreCase checks if all substrings in the given slice exist in the StringBuilder's value, ignoring case.
func (sb *StringBuilder) ContainsAllIgnoreCase(substrs []string) bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return containsAllIgnoreCase(sb.value, substrs)
}

// HasPrefix checks if the StringBuilder's value starts with the specified prefix and returns true or false accordingly.
func (sb *StringBuilder) HasPrefix(prefix string) bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return hasPrefix(sb.value, prefix)
}

// HasSuffix determines whether the StringBuilder's value ends with the specified suffix and returns true if it does.
func (sb *StringBuilder) HasSuffix(suffix string) bool {
	if !sb.shouldContinueProcessing() {
		return false
	}
	return hasSuffix(sb.value, suffix)
}

// RequireEmail validates if the StringBuilder's value is a valid email format,
// sets an error if invalid, and returns the instance.
func (sb *StringBuilder) RequireEmail() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !isEmail(sb.value) {
		return sb.setError(errors.ErrInvalidEmail, true)
	}
	return sb
}

// RequireDomain ensures that the value of the StringBuilder is a valid domain, setting an error if validation fails.
func (sb *StringBuilder) RequireDomain() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !isDomain(sb.value) {
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
	if !isURL(sb.value) {
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
	if !isUUID(sb.value) {
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
	if !isAlphaNumeric(sb.value) {
		return sb.setError(errors.ErrInvalidNotAlphaNumeric, true)
	}
	return sb
}

// RequireNumeric validates if the StringBuilder's value contains only numeric characters, with optional strict mode.
// If validation fails, it sets an error and halts further processing in the builder.
func (sb *StringBuilder) RequireNumeric(strict bool) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !isNumeric(sb.value, strict) {
		return sb.setError(errors.ErrInvalidNotNumeric, true)
	}
	return sb
}

// RequireAlpha ensures the StringBuilder's value contains only alphabetic characters, setting an error if invalid.
func (sb *StringBuilder) RequireAlpha() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !isAlpha(sb.value) {
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

// RequireContains ensures that the StringBuilder's value contains the specified substring.
// If the substring is not found, it sets an error and halts further processing.
func (sb *StringBuilder) RequireContains(substr string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !contains(sb.value, substr) {
		return sb.setError(errors.ErrDoesNotContainSubstring, true)
	}
	return sb
}

// RequireContainsIgnoreCase verifies that the StringBuilder's value contains the given substring, ignoring case.
// Returns the StringBuilder itself, or sets an error if the substring is not found.
func (sb *StringBuilder) RequireContainsIgnoreCase(substr string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !containsIgnoreCase(sb.value, substr) {
		return sb.setError(errors.ErrDoesNotContainSubstring, true)
	}
	return sb
}

// RequireContainsAny ensures the string contains at least one of the specified substrings, else sets an error.
func (sb *StringBuilder) RequireContainsAny(substrs []string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !containsAny(sb.value, substrs) {
		return sb.setError(errors.ErrDoesNotContainSubstring, true)
	}
	return sb
}

// RequireContainsAnyIgnoreCase ensures the string contains at least one of the specified substrings, disregarding case.
// If none of the substrings are found, it sets an error and halts further processing.
// Returns the updated StringBuilder instance.
func (sb *StringBuilder) RequireContainsAnyIgnoreCase(substrs []string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !containsAnyIgnoreCase(sb.value, substrs) {
		return sb.setError(errors.ErrDoesNotContainSubstring, true)
	}
	return sb
}

// RequireContainsAll ensures the StringBuilder's value contains all provided substrings; sets an error if not.
func (sb *StringBuilder) RequireContainsAll(substrs []string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !containsAll(sb.value, substrs) {
		return sb.setError(errors.ErrDoesNotContainSubstring, true)
	}
	return sb
}

// RequireContainsAllIgnoreCase checks if all substrings in the slice exist in the current value,
// ignoring case sensitivity.
// Returns the StringBuilder instance with an error set if any substring is missing.
func (sb *StringBuilder) RequireContainsAllIgnoreCase(substrs []string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !containsAllIgnoreCase(sb.value, substrs) {
		return sb.setError(errors.ErrDoesNotContainSubstring, true)
	}
	return sb
}

// RequireHasPrefix ensures the builder's string value has the specified prefix or sets an error
// if the prefix is missing.
func (sb *StringBuilder) RequireHasPrefix(prefix string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !hasPrefix(sb.value, prefix) {
		return sb.setError(errors.ErrMissingPrefix, true)
	}
	return sb
}

// RequireHasSuffix ensures the StringBuilder's value ends with the specified suffix, setting an error if not.
func (sb *StringBuilder) RequireHasSuffix(suffix string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	if !hasSuffix(sb.value, suffix) {
		return sb.setError(errors.ErrMissingSuffix, true)
	}
	return sb
}
