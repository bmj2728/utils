// Package strutil provides utilities for string manipulation and analysis.
// functions are available as part of a functional api as well as
// a fluent builder api
//
// This package includes functions for:
//   - String generation
//   - Text processing manipulation
//   - Text sanitizing functions
//   - Text comparison
//
// # Performance Considerations
//
// Most functions are designed for moderate-sized strings.
// For large text processing, consider streaming approaches.
package strutil

// Validation

// Basic Manipulation

// AppendString concatenates the input string `s` with `suffix`, using `sep`
// as the separator, and returns the resulting string.
func AppendString(s string, suffix string, sep string) string {
	return appendString(s, suffix, sep)
}

// PrependString adds a prefix and separator to the beginning of the given string and returns the resulting string.
func PrependString(s string, prefix string, sep string) string {
	return prependString(s, prefix, sep)
}

// CleanWhitespace removes all whitespace characters (spaces, tabs, newlines, etc.)
// from the input string and returns the result.
func CleanWhitespace(s string) string {
	return cleanWhitespace(s)
}

// CleanWhitespaceWithIgnore removes whitespace from the input string while
// ignoring whitespace characters specified in the charset string.
// Whitespace Chars: '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP)
func CleanWhitespaceWithIgnore(s string, charset string) string {
	return cleanWhitespaceWithIgnore(s, charset)
}

// NormalizeWhitespace removes excess whitespace by trimming and collapsing
// multiple whitespace characters into single spaces.
func NormalizeWhitespace(s string) string {
	return normalizeWhitespace(s)
}

// CollapseWhitespace reduces all consecutive whitespace characters in a string to a single space,
// preserving leading and trailing whitespace.
func CollapseWhitespace(s string) string {
	return collapseWhitespace(s)
}

// ReplaceWhitespace replaces all whitespace characters in the input string with the given replacement string.
func ReplaceWhitespace(s string, replacement string) string {
	return replaceWhitespace(s, replacement)
}

// ReplaceSpaces replaces all spaces in the input string with the specified replacement string.
func ReplaceSpaces(s string, replacement string) string {
	return replaceSpaces(s, replacement)
}

// Trim removes all leading and trailing white spaces from the given string and returns the trimmed result.
func Trim(s string) string {
	return trim(s)
}

// TrimLeft removes all leading whitespace characters from the input string `s`.
func TrimLeft(s string) string {
	return trimLeft(s)
}

// TrimRight removes all trailing whitespace characters from the given string.
func TrimRight(s string) string {
	return trimRight(s)
}

// TrimChars removes all leading and trailing occurrences of specified characters from the input string.
func TrimChars(s string, chars string) string {
	return trimChars(s, chars)
}

// TrimCharsLeft removes all leading characters specified in 'chars' from the input string 's'.
func TrimCharsLeft(s string, chars string) string {
	return trimCharsLeft(s, chars)
}

// TrimCharsRight removes all specified characters from the end of the given string.
func TrimCharsRight(s string, chars string) string {
	return trimCharsRight(s, chars)
}

// AlphaReplace replaces all non-alphabetic characters in the input string with the specified replacement string.
func AlphaReplace(s string, replacement string) string {
	return alphaReplace(s, replacement)
}

// AlphaNumericReplace replaces all non-alphanumeric characters in the input string
// with the specified replacement string.
func AlphaNumericReplace(s string, replacement string) string {
	return alphaNumericReplace(s, replacement)
}

// NormalizeDiacritics removes diacritical marks (accents) from the input string and returns the normalized version.
func NormalizeDiacritics(s string) string {
	return normalizeDiacritics(s)
}

// Slugify converts a string to a URL-friendly slug, ensuring lowercase, trimming, truncation,
// and replacing non-alphanumerics.
func Slugify(s string, length int) string {
	return slugify(s, length)
}

// Truncate shortens the input string s to the specified length and appends the given suffix if truncation occurs.
func Truncate(s string, length int, suffix string) string {
	return truncate(s, length, suffix)
}

// KeepAlpha removes all non-alphabetic characters from the input string, optionally keeping whitespace if ws is true.
func KeepAlpha(s string, ws bool) string {
	return alpha(s, ws)
}

// KeepAlphaNumeric removes all non-alphanumeric characters from the input string,
// optionally preserving whitespace if ws is true.
func KeepAlphaNumeric(s string, ws bool) string {
	return alphaNumeric(s, ws)
}

// HTML Sanitization

// StripHTML removes all HTML tags from the input string and sanitizes it to ensure safe content.
func StripHTML(s string) string {
	return stripHTML(s)
}

// SanitizeHTML removes potentially unsafe or harmful content from the input HTML string.
func SanitizeHTML(s string) string {
	return sanitizeHTML(s)
}

// SanitizeHTMLCustom sanitizes an HTML string, allowing only the tags specified in the allowedTags slice.
// Returns a sanitized string with disallowed HTML elements removed or escaped.
func SanitizeHTMLCustom(s string, allowedTags []string) string {
	return sanitizeHTMLCustom(s, allowedTags)
}

// EscapeHTML escapes special HTML characters in a string, replacing them with their corresponding HTML entity codes.
func EscapeHTML(s string) string {
	return escapeHTML(s)
}

// Case Conversion

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

//TODO additional sections

// File/Path Safety

func SanitizeFilename(s string) string {
	panic("Implement me!")
}

func SanitizePath(s string) string {
	panic("Implement me!")
}

// General Sanitization

func RemoveNonPrintable(s string) string {
	panic("Implement me!")
}

func NormalizeUnicode(s string) string {
	panic("Implement me!")
}

func StripAnsi(s string) string {
	panic("Implement me!")
}

// User Input Sanitization

func SanitizeEmail(s string) string {
	panic("Implement me!")
}

func SanitizeUsername(s string) string {
	panic("Implement me!")
}

func SanitizeSearchQuery(s string) string {
	panic("Implement me!")
}

// Format-Specific Sanitization

func SanitizePhone(s string) string {
	panic("Implement me!")
}

//func FormatPhone(s string, format PhoneFormat) string {
//	panic("Implement me!")
//}

func SanitizeCardNumber(s string) string {
	panic("Implement me!")
}

func FormatCardNumber(s string) string {
	panic("Implement me!")
}

func MaskCardNumber(s string) string {
	panic("Implement me!")
}

func SanitizePostalCode(s string) string {
	panic("Implement me!")
}

func SanitizeSSN(s string) string {
	panic("Implement me!")
}

func FormatSSN(s string) string {
	panic("Implement me!")
}

func MaskSSN(s string) string {
	panic("Implement me!")
}

// Security Sanitization

func SanitizeEnvValue(s string) string {
	panic("Implement me!")
}

func SanitizeCommandLineArg(s string) string {
	panic("Implement me!")
}

func SanitizeLogMessage(s string) string {
	panic("Implement me!")
}

func SanitizeLogMessageWithLevel(s string, level string) string {
	panic("Implement me!")
}

func SanitizeShellArg(s string) string {
	panic("Implement me!")
}

func EscapeSQL(s string) string {
	panic("Implement me!")
}

// Comparison Functions
