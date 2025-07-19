package strutil

// RequireEmail validates if the StringBuilder's value is a valid email format,
// sets an error if invalid, and returns the instance.
func (sb *StringBuilder) RequireEmail() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isValidEmail(sb.value) {
		sb.err = ErrInvalidEmail
	}
	return sb
}

// RequireDomain ensures that the value of the StringBuilder is a valid domain, setting an error if validation fails.
func (sb *StringBuilder) RequireDomain() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isValidDomain(sb.value) {
		sb.err = ErrInvalidDomain
	}
	return sb
}

// RequireURL validates if the StringBuilder's value is a properly formatted URL,
// sets an error if invalid, and returns the instance.
func (sb *StringBuilder) RequireURL() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isValidURL(sb.value) {
		sb.err = ErrInvalidURL
	}
	return sb
}

// RequireUUID validates whether the StringBuilder's value conforms to a valid UUID format,
// sets an error if invalid, and returns the instance.
func (sb *StringBuilder) RequireUUID() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isValidUUID(sb.value) {
		sb.err = ErrInvalidUUID
	}
	return sb
}

// RequireLength validates that the StringBuilder's value length is within the specified min and max range.
// Sets an error if invalid.
func (sb *StringBuilder) RequireLength(min, max int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if min < 0 || max < 0 {
		sb.err = ErrInvalidLengthRange
		return sb
	} else if min > max {
		sb.err = ErrInvalidLengthRange
	} else if !isLengthInRange(sb.value, min, max) {
		sb.err = ErrInvalidLength
	}
	return sb
}

// RequireNotEmpty ensures the StringBuilder's value is not empty, sets an error if it is, and returns the instance.
func (sb *StringBuilder) RequireNotEmpty() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if isEmpty(sb.value) {
		sb.err = ErrInvalidEmpty
	}
	return sb
}

// RequireNotEmptyNormalized ensures the StringBuilder's value is not empty after normalizing whitespace,
// setting an error otherwise.
func (sb *StringBuilder) RequireNotEmptyNormalized() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if isEmptyNormalized(sb.value) {
		sb.err = ErrInvalidEmptyAfterNormalization
	}
	return sb
}

// RequireAlphaNumeric ensures the StringBuilder's value contains only alphanumeric characters,
// setting an error if invalid.
func (sb *StringBuilder) RequireAlphaNumeric() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isAlphaNumericString(sb.value) {
		sb.err = ErrInvalidNotAlphaNumeric
	}
	return sb
}

// RequireAlpha ensures the StringBuilder's value contains only alphabetic characters, setting an error if invalid.
func (sb *StringBuilder) RequireAlpha() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isAlphaString(sb.value) {
		sb.err = ErrInvalidNotAlpha
	}
	return sb
}
