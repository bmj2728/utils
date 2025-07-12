package strutil

// ToLower converts all characters in the StringBuilder's value to lowercase and returns the updated StringBuilder.
func (sb *StringBuilder) ToLower() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(toLower(sb.value))
	return sb
}

// ToUpper converts the StringBuilder's current value to uppercase and returns the updated StringBuilder.
func (sb *StringBuilder) ToUpper() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(toUpper(sb.value))
	return sb
}

// Capitalize converts the first character of the StringBuilder's value to uppercase while preserving the rest as is.
func (sb *StringBuilder) Capitalize() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(capitalize(sb.value))
	return sb
}

// Uncapitalize converts the first character of the StringBuilder's value to lowercase if no error is present.
func (sb *StringBuilder) Uncapitalize() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(uncapitalize(sb.value))
	return sb
}

// ToTitleCase converts the string value of the StringBuilder to title case
// and returns the updated StringBuilder instance.
func (sb *StringBuilder) ToTitleCase() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(toTitleCase(sb.value))
	return sb
}

// SplitCamelCase splits the string stored in the StringBuilder into separate words based on camel case boundaries.
func (sb *StringBuilder) SplitCamelCase() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(splitCamelCase(sb.value))
	return sb
}

// SplitPascalCase splits a PascalCase string into separate words, modifying the StringBuilder's value in-place.
func (sb *StringBuilder) SplitPascalCase() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(splitPascalCase(sb.value))
	return sb
}

// ToSnakeCase converts the current string to snake_case or SCREAMING_SNAKE_CASE based on the scream parameter.
func (sb *StringBuilder) ToSnakeCase(scream bool) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(toSnakeCase(sb.value, scream))
	return sb
}

// ToSnakeCaseWithIgnore converts the StringBuilder's value to snake_case,
// optionally in uppercase, ignoring specified characters.
func (sb *StringBuilder) ToSnakeCaseWithIgnore(scream bool, ignore string) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(toSnakeCaseWithIgnore(sb.value, scream, ignore))
	return sb
}

// ToKebabCase converts the string in the StringBuilder to kebab-case or screaming-kebab-case based on the scream flag.
func (sb *StringBuilder) ToKebabCase(scream bool) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(toKebabCase(sb.value, scream))
	return sb
}

// ToCamelCase converts the current string value to camel case format and updates the StringBuilder instance.
func (sb *StringBuilder) ToCamelCase() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(toCamelCase(sb.value))
	return sb
}

// ToPascalCase converts the current string value of the StringBuilder
// to PascalCase format and updates the StringBuilder.
func (sb *StringBuilder) ToPascalCase() *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(toPascalCase(sb.value))
	return sb
}

// ToDelimited converts the string in the StringBuilder to a delimited format using the specified delimiter and options.
func (sb *StringBuilder) ToDelimited(delim uint8, ignore string, scream bool) *StringBuilder {
	if !sb.shouldContinueProcessing() {
		return sb
	}
	sb.setValue(toDelimited(sb.value, delim, ignore, scream))
	return sb
}
