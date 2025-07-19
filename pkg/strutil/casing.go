package strutil

// ToUpper converts the input string to uppercase and returns the result.
func ToUpper(s string) string {
	return toUpper(s)
}

// ToLower converts all characters in the input string to lowercase and returns the resulting string.
func ToLower(s string) string {
	return toLower(s)
}

// Capitalize returns the input string with the first character converted to uppercase while preserving the rest as is.
func Capitalize(s string) string {
	return capitalize(s)
}

// Uncapitalize takes a string and returns a new string with the first character converted to lowercase.
func Uncapitalize(s string) string {
	return uncapitalize(s)
}

// ToTitleCase converts the input string to title case, capitalizing
// the first letter of each word following English rules.
func ToTitleCase(s string) string {
	return toTitleCase(s)
}

// SplitCamelCase splits a CamelCase string into separate words with spaces between them.
func SplitCamelCase(s string) string {
	return splitCamelCase(s)
}

// SplitPascalCase splits a PascalCase or camelCase string into space-separated words.
func SplitPascalCase(s string) string {
	return splitPascalCase(s)
}

// ToSnakeCase converts a string to snake_case format. If scream is true, the output will be in SCREAMING_SNAKE_CASE.
func ToSnakeCase(s string, scream bool) string {
	return toSnakeCase(s, scream)
}

// ToSnakeCaseWithIgnore converts a string to snake_case format, optionally in uppercase, ignoring specified characters.
func ToSnakeCaseWithIgnore(s string, scream bool, ignore string) string {
	return toSnakeCaseWithIgnore(s, scream, ignore)
}

// ToKebabCase converts a string to kebab-case or screaming-kebab-case based on the scream flag.
// Diacritics in the input string are normalized before conversion.
func ToKebabCase(s string, scream bool) string {
	return toKebabCase(s, scream)
}

// ToCamelCase converts a string to camel case format where the first letter
// is lowercase and subsequent words are capitalized.
func ToCamelCase(s string) string {
	return toCamelCase(s)
}

// ToPascalCase converts a given string to PascalCase format, capitalizing the first letter of each word.
func ToPascalCase(s string) string {
	return toPascalCase(s)
}

// ToDelimited converts a string into a delimited format based on the specified delimiter and casing options.
func ToDelimited(s string, delim uint8, ignore string, scream bool) string {
	return toDelimited(s, delim, ignore, scream)
}
