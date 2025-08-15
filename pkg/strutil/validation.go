package strutil

// IsEmail checks if the input string is in a valid email address format and
// returns true if valid, false otherwise.
func IsEmail(s string) bool {
	return isEmail(s)
}

// IsURL determines whether the input string is a valid URL with a scheme and host.
// Returns true if valid, otherwise false.
func IsURL(s string) bool {
	return isURL(s)
}

// IsDomain checks if a given string is a valid domain name format as per defined rules.
func IsDomain(domain string) bool {
	return isDomain(domain)
}

// IsUUID verifies if the provided string has a valid UUID format. Returns true if valid, false otherwise.
func IsUUID(s string) bool {
	return isUUID(s)
}

// IsLengthInRange checks if the length of the given string s is within the inclusive range
// defined by min and max values.
func IsLengthInRange(s string, min, max int) bool {
	return isLengthInRange(s, min, max)
}

// IsEmpty checks if the provided string is empty and returns true if it is, otherwise false.
func IsEmpty(s string) bool {
	return isEmpty(s)
}

// IsEmptyNormalized checks if the normalized version of the input string is empty
// after trimming and collapsing whitespace.
func IsEmptyNormalized(s string) bool {
	return isEmptyNormalized(s)
}

// IsAlphaNumericString checks if the input string consists only of alphanumeric characters (letters and digits).
func IsAlphaNumericString(s string) bool {
	return isAlphaNumericString(s)
}

// IsAlphaString checks if the given string contains only alphabetic characters.
// Returns true if all characters are letters.
func IsAlphaString(s string) bool {
	return isAlphaString(s)
}

// IsNormalizedUnicode checks if the given string is normalized according to the specified
// Unicode normalization format.
func IsNormalizedUnicode(s string, format NormalizationFormat) bool {
	return isNormalizedUnicode(s, format)
}

// Contains checks if the substring `substr` is present within the string `s` and
// returns true if found, otherwise false.
func Contains(s string, substr string) bool {
	return contains(s, substr)
}

// ContainsIgnoreCase checks if substr is present in s, ignoring case sensitivity. Returns true if found.
func ContainsIgnoreCase(s string, substr string) bool {
	return containsIgnoreCase(s, substr)
}

// ContainsAny checks if any substring in the slice `substrs` exists in the string `s` and returns true if found.
func ContainsAny(s string, substrs []string) bool {
	return containsAny(s, substrs)
}

// ContainsAnyIgnoreCase checks if the input string contains any of the provided substrings,
// ignoring case sensitivity.
func ContainsAnyIgnoreCase(s string, substrs []string) bool {
	return containsAnyIgnoreCase(s, substrs)
}

// ContainsAll checks if all strings in the slice substrs are present within the string s, returning true if found.
func ContainsAll(s string, substrs []string) bool {
	return containsAll(s, substrs)
}

// ContainsAllIgnoreCase checks if all substrings in the provided slice are found in the string,
// ignoring case sensitivity.
func ContainsAllIgnoreCase(s string, substrs []string) bool {
	return containsAllIgnoreCase(s, substrs)
}

// HasPrefix reports whether the string s begins with the specified prefix. Returns true if it does, otherwise false.
func HasPrefix(s string, prefix string) bool {
	return hasPrefix(s, prefix)
}

// HasSuffix reports whether the string `s` ends with the specified `suffix`.
func HasSuffix(s string, suffix string) bool {
	return hasSuffix(s, suffix)
}
