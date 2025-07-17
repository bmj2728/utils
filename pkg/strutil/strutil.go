// Package strutil provides utilities for string manipulation and analysis.
// functions are available as part of a functional api as well as
// a fluent builder api
//
// This package includes functions for:
//   - String generation
//   - Text processing manipulation
//   - Text sanitizing functions
//
// # Usage
//
// Basic example to create a url-safe slug using the functional api:
//
//	 title := "The Fantasic Four: First Steps"
//		slug := strutil.Slugify(title, 100)
//		fmt.Print(slug) // Output: the-fantastic-four-first-steps
//
// Using the fluent api:
//
//		slug := strutil.New("The Fantasic Four: First Steps").Slugify(100).String()
//	 fmt.Print(slug) // Output: the-fantastic-four-first-steps
//
// # Performance Considerations
//
// Most functions are designed for moderate-sized strings.
// For large text processing, consider streaming approaches.
package strutil

// UUID Generation

// GenerateUUID generates and returns a new random UUID as a string.
func GenerateUUID() string {
	return makeUUID()
}

// GenerateUUIDV7 generates and returns a new random UUID as a string using UUID V7
// It is recommended to use V7 unless legacy compatibility is required
func GenerateUUIDV7() string {
	return makeUUIDV7()
}

// String Generation

// RandomString generates a random alphanumeric string of the specified length using the AlphaNumeric character set.
// This uses math/rand and is NOT cryptographically secure.
// For security-sensitive applications, use crypto/rand directly.
func RandomString(length int) string {
	return randomFromCharset(length, AlphaNumeric)
}

// RandomHex generates a random hexadecimal string of the specified length.
func RandomHex(length int) string {
	return randomFromCharset(length, HexChars)
}

// RandomUrlSafe generates a random URL-safe string of the specified length using characters suitable for URLs.
func RandomUrlSafe(length int) string {
	return randomFromCharset(length, URLSafe)
}

// LoremWord generates and returns a random lorem ipsum word as a string.
func LoremWord() string {
	return loremWord()
}

// LoremWords generates a string containing the specified number of lorem ipsum words.
func LoremWords(count int) string {
	return loremWords(count)
}

// LoremSentence generates and returns a placeholder sentence of 8 words using lorem ipsum text.
func LoremSentence() string {
	return loremSentence()
}

// LoremSentenceCustom generates a lorem ipsum sentence with the specified word count. Returns the generated string.
func LoremSentenceCustom(count int) string {
	return loremSentenceCustom(count)
}

// LoremSentences generates a string containing the specified number of 8 word lorem ipsum sentences.
func LoremSentences(count int) string {
	return loremSentences(count)
}

// LoremSentencesCustom generates multiple lorem ipsum sentences with
// the specified sentence count and word length per sentence.
func LoremSentencesCustom(count int, length int) string {
	return loremSentencesCustom(count, length)
}

// LoremSentencesVariable generates variable length lorem sentences with lengths between specified min and max values.
// The parameter 'count' specifies the number of sentences to generate.
func LoremSentencesVariable(count, min, max int) string {
	return loremSentencesVariable(count, min, max)
}

// LoremParagraph generates and returns a string containing a randomly generated Lorem Ipsum paragraph of 45 words.
func LoremParagraph() string {
	return loremParagraph()
}

// LoremParagraphs generates and returns a specified number of lorem ipsum paragraphs as a single string.
// The parameter 'count' specifies the number of paragraphs to generate.
func LoremParagraphs(count int) string {
	return loremParagraphs(count)
}

// LoremDomain generates and returns a placeholder domain name in string format.
func LoremDomain() string {
	return loremDomain()
}

// LoremURL generates and returns a string representing a placeholder or mock URL,
// intended for testing or default usage.
func LoremURL() string {
	return loremURL()
}

// LoremEmail generates and returns a placeholder or mock email address as a string.
func LoremEmail() string {
	return loremEmail()
}

// Validation

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

// Basic Manipulation

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

// CleanWhitespace removes all whitespace characters (spaces, tabs, newlines, etc.)
// from the input string and returns the result.
func CleanWhitespace(s string) string {
	return cleanWhitespace(s)
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

// Comparison Functions - edlib

// LevenshteinDistance calculates the Levenshtein distance between two strings s1 and s2.
// It represents the minimum number of edits needed to convert one string into the other.
// An edit is an insertion, deletion, or substitution of a single character.
//
// Additional information: https://en.wikipedia.org/wiki/Levenshtein_distance
func LevenshteinDistance(s1, s2 string) int {
	return levenshteinDistance(s1, s2)
}

// DamerauLevenshteinDistance computes the edit distance between two strings,
// including transpositions of adjacent characters.
//
// It represents the minimum number of operations to change one string to another.
// An operation is an insertion, deletion/substitution of a single character, or transposition of adjacent characters.
// Additional information: https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance
func DamerauLevenshteinDistance(s1, s2 string) int {
	return damerauLevenshteinDistance(s1, s2)
}

// OSADamerauLevenshteinDistance calculates the Damerau-Levenshtein
// distance between two strings for similarity measurement.
//
// This optimal string alignment variant of DamerauLevenshteinDistance
// does not allow multiple transformations on the same substring
func OSADamerauLevenshteinDistance(s1, s2 string) int {
	return osaDamerauLevenshteinDistance(s1, s2)
}

// LCS calculates the length of the longest common subsequence between two input strings s1 and s2.
func LCS(s1, s2 string) int {
	return lcs(s1, s2)
}
