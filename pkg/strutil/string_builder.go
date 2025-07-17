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

// NewRandom creates a new StringBuilder with a random alphanumeric string of the specified length.
func NewRandom(length int) *StringBuilder {
	return &StringBuilder{
		value: randomFromCharset(length, AlphaNumeric),
	}
}

// NewRandomHex generates a new StringBuilder containing a random hexadecimal string of the specified length.
func NewRandomHex(length int) *StringBuilder {
	return &StringBuilder{
		value: randomFromCharset(length, HexChars),
	}
}

// NewRandomURLSafe generates a StringBuilder initialized with a random URL-safe string of the given length.
func NewRandomURLSafe(length int) *StringBuilder {
	return &StringBuilder{
		value: randomFromCharset(length, URLSafe),
	}
}

// NewUUID creates and returns a new StringBuilder instance with a generated UUID value.
func NewUUID() *StringBuilder {
	return &StringBuilder{
		value: makeUUID(),
	}
}

// NewUUIDV7 generates a new UUID version 7 and returns it wrapped in a StringBuilder instance.
func NewUUIDV7() *StringBuilder {
	return &StringBuilder{
		value: makeUUIDV7(),
	}
}

// NewLoremWord creates a new StringBuilder instance with a randomly generated word as its initial value.
func NewLoremWord() *StringBuilder {
	return &StringBuilder{
		value: loremWord(),
	}
}

// NewLoremWords creates a new StringBuilder initialized with a string
// containing the specified number of lorem ipsum words.
func NewLoremWords(count int) *StringBuilder {
	return &StringBuilder{
		value: loremWords(count),
	}
}

// NewLoremSentence creates and returns a new StringBuilder initialized with a randomly generated lorem ipsum sentence.
func NewLoremSentence() *StringBuilder {
	return &StringBuilder{
		value: loremSentence(),
	}
}

// NewLoremSentenceCustom creates a new StringBuilder instance containing
// a lorem ipsum sentence with the specified word count.
func NewLoremSentenceCustom(length int) *StringBuilder {
	return &StringBuilder{
		value: loremSentenceCustom(length),
	}
}

// NewLoremSentences creates a new StringBuilder containing the given number of lorem ipsum sentences.
func NewLoremSentences(count int) *StringBuilder {
	return &StringBuilder{
		value: loremSentences(count),
	}
}

// NewLoremSentencesCustom creates a new StringBuilder instance
// containing lorem ipsum sentences based on the given count and length.
func NewLoremSentencesCustom(count int, length int) *StringBuilder {
	return &StringBuilder{
		value: loremSentencesCustom(count, length),
	}
}

// NewLoremSentencesVariable generates a string of lorem ipsum sentences
// with count, min, and max controlling quantity and length.
func NewLoremSentencesVariable(count int, min int, max int) *StringBuilder {
	return &StringBuilder{
		value: loremSentencesVariable(count, min, max),
	}
}

// NewLoremParagraph creates a new StringBuilder initialized with a random Lorem Ipsum paragraph.
func NewLoremParagraph() *StringBuilder {
	return &StringBuilder{
		value: loremParagraph(),
	}
}

// NewLoremParagraphs generates a StringBuilder containing the specified number of lorem ipsum paragraphs.
func NewLoremParagraphs(count int) *StringBuilder {
	return &StringBuilder{
		value: loremParagraphs(count),
	}
}

// NewLoremDomain creates and returns a new StringBuilder initialized
// with a domain string from the loremDomain function.
func NewLoremDomain() *StringBuilder {
	return &StringBuilder{
		value: loremDomain(),
	}
}

// NewLoremURL creates and returns a new StringBuilder initialized with a lorem ipsum URL string.
func NewLoremURL() *StringBuilder {
	return &StringBuilder{
		value: loremURL(),
	}
}

// NewLoremEmail initializes a StringBuilder with a generated mock email value and returns the instance.
func NewLoremEmail() *StringBuilder {
	return &StringBuilder{
		value: loremEmail(),
	}
}

// Manipulation Methods

// CleanWhitespace removes all whitespace characters from the StringBuilder's value
// and returns the updated StringBuilder.
func (sb *StringBuilder) CleanWhitespace() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = cleanWhitespace(sb.value)
	return sb
}

// NormalizeWhitespace collapses consecutive whitespace characters into a single space
// and trims leading and trailing spaces.
func (sb *StringBuilder) NormalizeWhitespace() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = normalizeWhitespace(sb.value)
	return sb
}

// CollapseWhitespace collapses consecutive whitespace characters in the StringBuilder's value into a single space
// and preserves leading and trailing spaces.
func (sb *StringBuilder) CollapseWhitespace() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = collapseWhitespace(sb.value)
	return sb
}

// ReplaceWhitespace replaces all whitespace characters in the StringBuilder's value
// with the specified replacement string.
func (sb *StringBuilder) ReplaceWhitespace(replacement string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = replaceWhitespace(sb.value, replacement)
	return sb
}

// ReplaceSpaces replaces all spaces in the StringBuilder's value with the specified replacement string.
func (sb *StringBuilder) ReplaceSpaces(replacement string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = replaceSpaces(sb.value, replacement)
	return sb
}

// KeepAlpha removes all non-alphabetic characters from the StringBuilder's value,
// optionally retaining whitespace if ws is true.
func (sb *StringBuilder) KeepAlpha(ws bool) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = alpha(sb.value, ws)
	return sb
}

// KeepAlphaNumeric removes all non-alphanumeric characters from the StringBuilder's value,
// optionally preserving whitespace.
func (sb *StringBuilder) KeepAlphaNumeric(ws bool) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = alphaNumeric(sb.value, ws)
	return sb
}

// StripHTML removes all HTML tags from the StringBuilder's value and returns the updated StringBuilder.
func (sb *StringBuilder) StripHTML() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = stripHTML(sb.value)
	return sb
}

// EscapeHTML escapes special HTML characters in the StringBuilder's value and returns the updated StringBuilder.
func (sb *StringBuilder) EscapeHTML() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = escapeHTML(sb.value)
	return sb
}

// SanitizeHTML sanitizes the StringBuilder's value by removing potentially unsafe or harmful HTML content.
func (sb *StringBuilder) SanitizeHTML() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = sanitizeHTML(sb.value)
	return sb
}

// SanitizeHTMLCustom sanitizes the StringBuilder's value, allowing only specified HTML tags,
// and returns the updated StringBuilder.
func (sb *StringBuilder) SanitizeHTMLCustom(allowedTags []string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = sanitizeHTMLCustom(sb.value, allowedTags)
	return sb
}

// Trim removes leading and trailing whitespace or specific characters from the StringBuilder's value.
func (sb *StringBuilder) Trim() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trim(sb.value)
	return sb
}

// TrimLeft removes all leading whitespace characters from the string stored in the StringBuilder and updates its value.
func (sb *StringBuilder) TrimLeft() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trimLeft(sb.value)
	return sb
}

// TrimRight removes trailing whitespace or specified characters from the end of the StringBuilder's value.
func (sb *StringBuilder) TrimRight() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trimRight(sb.value)
	return sb
}

// TrimChars removes all leading and trailing characters specified in the input string from the StringBuilder's value.
func (sb *StringBuilder) TrimChars(chars string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trimChars(sb.value, chars)
	return sb
}

// TrimCharsLeft removes all leading occurrences of the specified characters from the string
// and returns the updated StringBuilder.
func (sb *StringBuilder) TrimCharsLeft(chars string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trimCharsLeft(sb.value, chars)
	return sb
}

// TrimCharsRight removes all occurrences of the specified characters from the end of the string
// and returns the StringBuilder.
func (sb *StringBuilder) TrimCharsRight(chars string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = trimCharsRight(sb.value, chars)
	return sb
}

// AlphaReplace replaces all alphabetical characters in the StringBuilder's value with the specified replacement string.
func (sb *StringBuilder) AlphaReplace(replacement string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = alphaReplace(sb.value, replacement)
	return sb
}

// AlphaNumericReplace replaces all alphanumeric characters in the StringBuilder value
// with the specified replacement string.
func (sb *StringBuilder) AlphaNumericReplace(replacement string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = alphaNumericReplace(sb.value, replacement)
	return sb
}

// NormalizeDiacritics removes diacritical marks from the string and replaces them with their non-accented counterparts.
func (sb *StringBuilder) NormalizeDiacritics() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = normalizeDiacritics(sb.value)
	return sb
}

// Slugify converts the string into a URL-friendly slug with a maximum length specified by the parameter.
func (sb *StringBuilder) Slugify(length int) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = slugify(sb.value, length)
	return sb
}

// Truncate shortens the string to the specified length and appends the provided suffix if truncation occurs.
func (sb *StringBuilder) Truncate(length int, suffix string) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = truncate(sb.value, length, suffix)
	return sb
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

// ToLower converts all characters in the StringBuilder's value to lowercase and returns the updated StringBuilder.
func (sb *StringBuilder) ToLower() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toLower(sb.value)
	return sb
}

// ToUpper converts the StringBuilder's current value to uppercase and returns the updated StringBuilder.
func (sb *StringBuilder) ToUpper() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toUpper(sb.value)
	return sb
}

// ToTitleCase converts the string value of the StringBuilder to title case and returns the updated StringBuilder instance.
func (sb *StringBuilder) ToTitleCase() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toTitleCase(sb.value)
	return sb
}

// SplitCamelCase splits the string stored in the StringBuilder into separate words based on camel case boundaries.
func (sb *StringBuilder) SplitCamelCase() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = splitCamelCase(sb.value)
	return sb
}

// SplitPascalCase splits a PascalCase string into separate words, modifying the StringBuilder's value in-place.
func (sb *StringBuilder) SplitPascalCase() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = splitPascalCase(sb.value)
	return sb
}

// ToSnakeCase converts the string to snake_case. `norm` controls normalization
// and `scream` enables uppercase snake case.
func (sb *StringBuilder) ToSnakeCase(norm bool, scream bool) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toSnakeCase(sb.value, norm, scream)
	return sb
}

// ToKebabCase converts the string value to kebab-case format. Optionally normalizes and handles uppercase if specified.
func (sb *StringBuilder) ToKebabCase(norm bool, scream bool) *StringBuilder {
	if sb.err != nil {
		return sb
	}
	sb.value = toKebabCase(sb.value, norm, scream)
	return sb
}

func (sb *StringBuilder) ToCamelCase() *StringBuilder {
	panic("Implement me!")
}

func (sb *StringBuilder) ToPascalCase() *StringBuilder {
	panic("Implement me!")
}

// Validation Methods (can set error)

// RequireEmail validates if the StringBuilder's value is a valid email format,
// sets an error if invalid, and returns the instance.
func (sb *StringBuilder) RequireEmail() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isValidEmail(sb.value) {
		sb.err = errors.New(ErrInvalidEmail)
	}
	return sb
}

// RequireDomain ensures that the value of the StringBuilder is a valid domain, setting an error if validation fails.
func (sb *StringBuilder) RequireDomain() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isValidDomain(sb.value) {
		sb.err = errors.New(ErrInvalidDomain)
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
		sb.err = errors.New(ErrInvalidURL)
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
		sb.err = errors.New(ErrInvalidUUID)
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
		sb.err = errors.New(ErrInvalidLengthRange)
		return sb
	} else if min > max {
		sb.err = errors.New(ErrInvalidLengthRange)
	} else if !isLengthInRange(sb.value, min, max) {
		sb.err = errors.New(ErrInvalidLength)
	}
	return sb
}

// RequireNotEmpty ensures the StringBuilder's value is not empty, sets an error if it is, and returns the instance.
func (sb *StringBuilder) RequireNotEmpty() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if isEmpty(sb.value) {
		sb.err = errors.New(ErrInvalidEmpty)
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
		sb.err = errors.New(ErrInvalidEmptyAfterNormalization)
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
		sb.err = errors.New(ErrInvalidNotAlphaNumeric)
	}
	return sb
}

// RequireAlpha ensures the StringBuilder's value contains only alphabetic characters, setting an error if invalid.
func (sb *StringBuilder) RequireAlpha() *StringBuilder {
	if sb.err != nil {
		return sb
	}
	if !isAlphaString(sb.value) {
		sb.err = errors.New(ErrInvalidNotAlpha)
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
//
// Examples:
//
// s := New("Hello World!").Transform(strings.ToUpper)
//
// result: "HELLO WORLD!"
//
//	s := New("Hello World!).Transform(func(input string) string {
//		return input + " Goodbye!"
//	})
//
// result: "Hello World! Goodbye!"
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

// Must returns the final string value or nil string if an error is present in the StringBuilder instance.
func (sb *StringBuilder) Must() string {
	if sb.err != nil {
		return ""
	}
	return sb.value
}

// Result returns the current value of the StringBuilder along with any associated error.
func (sb *StringBuilder) Result() (string, error) {
	return sb.value, sb.err
}
