package strutil

// IsEmail checks if the input string is in a valid email address format and returns true if valid, false otherwise.
func IsEmail(s string) bool {
	return isValidEmail(s)
}

// IsURL determines whether the input string is a valid URL with a scheme and host.
// Returns true if valid, otherwise false.
func IsURL(s string) bool {
	return isValidURL(s)
}

// IsDomain checks if a given string is a valid domain name format as per defined rules.
func IsDomain(domain string) bool {
	return isValidDomain(domain)
}

// IsUUID verifies if the provided string has a valid UUID format. Returns true if valid, false otherwise.
func IsUUID(s string) bool {
	return isValidUUID(s)
}

// IsValidLength checks if the length of the given string s is within the inclusive range defined by min and max values.
func IsValidLength(s string, min, max int) bool {
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

// IsNormalizedUnicode checks if the given string is normalized according to the specified Unicode normalization format.
func IsNormalizedUnicode(s string, format NormalizationFormat) bool {
	return isNormalizedUnicode(s, format)
}

// input validation

//func FormatCardNumber(s string) string {
//	panic("Implement me!")
//}

//func FormatSSN(s string) string {
//	panic("Implement me!")
//}

//func FormatPhone(s string, format PhoneFormat) string {
//	panic("Implement me!")
//}
