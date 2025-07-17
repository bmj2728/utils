package strutil

import (
	"github.com/hbollon/go-edlib"
	"math/rand"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"unicode"

	godiacritics "github.com/Regis24GmbH/go-diacritics"
	lorelai "github.com/UltiRequiem/lorelai/pkg"
	"github.com/fatih/camelcase"
	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mrz1836/go-sanitize"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func makeUUID() string {
	return uuid.NewString()
}

func makeUUIDV7() string {
	u, err := uuid.NewV7()
	if err != nil {
		return ""
	}
	return u.String()
}

// cleanWhitespace removes all whitespace characters (spaces, tabs, newlines, etc.)
// from the input string and returns the result.
func cleanWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, c := range s {
		if !unicode.IsSpace(c) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// normalizeWhitespace collapses whitespace runs to single spaces and trims
func normalizeWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// collapseWhitespace reduces all consecutive whitespace characters in the input string to a single space preserving
// leading and trailing whitespace.
func collapseWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	prevWasSpace := false

	for _, r := range s {
		if unicode.IsSpace(r) {
			if !prevWasSpace {
				b.WriteRune(' ')
			}
			prevWasSpace = true
		} else {
			b.WriteRune(r)
			prevWasSpace = false
		}
	}
	return b.String()
}

// isValidEmail checks if the provided string s is a valid email address format
// and returns true if valid, false otherwise.
func isValidEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

// isValidURL checks if the provided string is a valid URL with a defined scheme and host.
// Returns true if valid, false otherwise.
func isValidURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}
	u, err := url.Parse(s)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

// isValidDomain validates whether a given string conforms to a valid domain name format based on specified rules.
func isValidDomain(domain string) bool {
	if domain == "" {
		return false
	}

	domainRegex := regexp.MustCompile(`^(` + LabelRegex + `\.)+` + TLDRegex + `$`)

	return domainRegex.MatchString(strings.TrimSpace(domain))
}

// isValidUUID checks if the provided string is a valid UUID. Returns true if valid, otherwise false.
func isValidUUID(s string) bool {
	err := uuid.Validate(s)
	return err == nil
}

// isLengthInRange checks if the length of the string s is within the inclusive range specified by min and max values.
func isLengthInRange(s string, min, max int) bool {
	if min < 0 || max < 0 {
		return false
	}
	if min > max {
		return false
	}
	return len(s) >= min && len(s) <= max
}

// isEmpty checks if the provided string is empty and returns true if it is, otherwise false.
func isEmpty(s string) bool {
	return len(s) == 0
}

// isEmptyNormalized checks if the input string is empty after normalizing and trimming whitespace.
func isEmptyNormalized(s string) bool {
	return len(normalizeWhitespace(s)) == 0
}

// toUpper converts the input string to uppercase and returns the result.
func toUpper(s string) string {
	return strings.ToUpper(s)
}

// toLower converts all characters in the input string to their lowercase equivalents and returns the resulting string.
func toLower(s string) string {
	return strings.ToLower(s)
}

// isAlphaNumericString checks whether a given string consists only of alphanumeric characters (letters and digits).
func isAlphaNumericString(s string) bool {
	for _, c := range s {
		if !isAlphaNumericRune(c) {
			return false
		}
	}
	return true
}

// isAlphaString checks if the given string consists only of alphabetic characters.
// Returns true if all characters are letters.
func isAlphaString(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

// isAlphaNumericRune determines if the given rune is an alphanumeric character (letter or digit).
func isAlphaNumericRune(r rune) bool {
	if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
		return false
	}
	return true
}

// isWhiteSpaceRune checks if the given rune is classified as a whitespace character based on Unicode standards.
func isWhiteSpaceRune(r rune) bool {
	return unicode.IsSpace(r)
}

// replaceWhitespace replaces all whitespace characters in the input string with the specified replacement string.
func replaceWhitespace(s string, replacement string) string {
	if isEmpty(s) {
		return s
	}

	for _, c := range s {
		if isWhiteSpaceRune(c) {
			s = strings.Replace(s, string(c), replacement, 1)
		}
	}
	return s
}

// replaceSpaces replaces all spaces in the input string with the specified replacement string.
func replaceSpaces(s string, replacement string) string {
	return strings.ReplaceAll(s, " ", replacement)
}

// randomFromCharset generates a random string of the specified length using characters from the provided charset.
func randomFromCharset(length int, charset string) string {
	if length < 1 {
		return ""
	}
	s := make([]byte, length)
	for i := range s {
		s[i] = charset[rand.Intn(len(charset))]
	}
	return string(s)
}

// alpha removes all non-alphabetic characters from the given string, optionally retaining whitespace if ws is true.
func alpha(s string, ws bool) string {
	return sanitize.Alpha(s, ws)
}

// alphaNumeric removes all non-alphanumeric characters from the input string,
// optionally preserving whitespace if ws is true.
func alphaNumeric(s string, ws bool) string {
	return sanitize.AlphaNumeric(s, ws)
}

// alphaReplace replaces all non-alphabetic characters in the input string with the specified replacement string.
func alphaReplace(s string, replacement string) string {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			s = strings.Replace(s, string(c), replacement, 1)
		}
	}
	return s
}

// alphaNumericReplace replaces all non-alphanumeric characters
// in the input string with the specified replacement string.
func alphaNumericReplace(s string, replacement string) string {
	for _, c := range s {
		if !isAlphaNumericRune(c) {
			s = strings.Replace(s, string(c), replacement, 1)
		}
	}
	return s
}

// stripHTML removes all HTML tags and sanitizes the input string to prevent potential security risks.
func stripHTML(s string) string {
	p := bluemonday.StrictPolicy()
	return p.Sanitize(s)
}

// sanitizeHTML sanitizes an input HTML string by removing potentially unsafe or harmful content.
func sanitizeHTML(s string) string {
	p := bluemonday.UGCPolicy()
	return p.Sanitize(s)
}

//TODO: extend implementation to better address complex options

// sanitizeHTMLCustom sanitizes the input HTML string by allowing only the specified elements in allowedElements.
func sanitizeHTMLCustom(s string, allowedElements []string) string {
	p := bluemonday.NewPolicy()
	p.AllowElements(allowedElements...)
	return p.Sanitize(s)
}

// escapeHTML escapes special HTML characters in a string to their corresponding HTML entity codes.
func escapeHTML(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		switch r {
		case '"':
			b.WriteString("&quot;")
		case '&':
			b.WriteString("&amp;")
		case '<':
			b.WriteString("&lt;")
		case '>':
			b.WriteString("&gt;")
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

// loremWord generates and returns a single random word as a string.
func loremWord() string {
	return lorelai.Word()
}

// loremWords generates a string containing the specified number of lorem ipsum words.
// If less than 1 returns a nil string
func loremWords(count int) string {
	if count < 1 {
		return ""
	}
	return lorelai.LoremWords(count)
}

// loremSentence returns a randomly generated lorem ipsum sentence as a string.
func loremSentence() string {
	return lorelai.Sentence()
}

// loremSentenceCustom generates a lorem ipsum sentence with the specified number of words
// determined by the length parameter.
func loremSentenceCustom(length int) string {
	if length < 1 {
		return ""
	}
	return lorelai.FormattedLoremWords(length)
}

// loremSentences generates a string containing the specified number of randomly generated lorem ipsum sentences.
func loremSentences(count int) string {
	if count < 1 {
		return ""
	}
	sentences := ""
	for i := 0; i < count; i++ {
		if sentences != "" {
			sentences += " "
		}
		sentences += loremSentence()
	}
	return sentences
}

// loremSentencesCustom generates multiple lorem ipsum sentences with a specified count and length per sentence.
func loremSentencesCustom(count int, length int) string {
	if count < 1 || length < 1 {
		return ""
	}
	sentences := ""
	for i := 0; i < count; i++ {
		if sentences != "" {
			sentences += " "
		}
		sentences += loremSentenceCustom(length)
	}
	return sentences
}

// loremSentenceVariable generates a specified number of lorem ipsum sentences
// with random lengths between min and max values.
func loremSentencesVariable(count, min, max int) string {
	if count < 1 || min > max {
		return ""
	}
	sentences := ""
	for i := 0; i < count; i++ {
		if sentences != "" {
			sentences += " "
		}
		sentences += loremSentenceCustom(rand.Intn(max-min) + min)
	}
	return sentences
}

// loremParagraph generates and returns a single Lorem Ipsum paragraph as a string.
func loremParagraph() string {
	return lorelai.Paragraph()
}

// loremParagraphs generates a specified number of lorem ipsum paragraphs and returns them as a single string.
// The parameter 'count' determines the number of paragraphs to generate.
// Returns an empty string if count is less than 1.
func loremParagraphs(count int) string {
	if count < 1 {
		return ""
	}
	paragraphs := ""
	for i := 0; i < count; i++ {
		if paragraphs != "" {
			paragraphs += "\n\n"
		}
		paragraphs += loremParagraph()
	}
	return paragraphs
}

// loremDomain returns a string representing the domain from the lorelai package.
func loremDomain() string {
	return lorelai.Domain()
}

// loremURL returns a URL string generated by the lorelai package.
func loremURL() string {
	return lorelai.URL()
}

// loremEmail generates and returns a mock email address using the lorelai package.
func loremEmail() string {
	return lorelai.Email()
}

// truncate shortens the input string s to the specified length and appends the given suffix if truncation occurs.
func truncate(s string, length int, suffix string) string {
	if len(s) <= length {
		return s
	}
	if length < 0 {
		return ""
	}
	return s[:length] + suffix
}

// trim removes all leading and trailing white spaces from the input string and returns the trimmed result.
func trim(s string) string {
	return strings.TrimSpace(s)
}

// trimLeft removes all leading whitespace characters from the input string `s` based on the defined `WhiteSpace` set.
func trimLeft(s string) string {
	return strings.TrimLeft(s, WhiteSpace)
}

// trimRight removes all trailing whitespace characters from the provided string.
func trimRight(s string) string {
	return strings.TrimRight(s, WhiteSpace)
}

// trimChars removes all leading and trailing occurrences of the characters specified in chars from the string s.
func trimChars(s string, chars string) string {
	return strings.Trim(s, chars)
}

// trimCharsLeft removes all leading characters found in 'chars' from the string 's'.
func trimCharsLeft(s string, chars string) string {
	return strings.TrimLeft(s, chars)
}

// trimCharsRight removes all occurrences of the specified characters from the end of the given string.
func trimCharsRight(s string, chars string) string {
	return strings.TrimRight(s, chars)
}

// normalizeDiacritics removes diacritical marks (accents) from the input string, returning the normalized version.
func normalizeDiacritics(s string) string {
	return godiacritics.Normalize(s)
}

// slugify converts a given string into a URL-friendly slug, ensuring lowercase, truncation, and hyphenation if needed.
func slugify(s string, length int) string {

	//early return if empty string
	if s == "" || length < 1 {
		return ""
	}

	// address camelCase/PascalCase
	if CamelCaseRegex.MatchString(s) {
		s = splitCamelCase(s)
	}

	//clean the diacritics
	s = normalizeDiacritics(s)

	// replace non-alphanumerics
	s = alphaNumericReplace(s, " ")

	// normalize whitespace
	s = normalizeWhitespace(s)

	// replace whitespace with "-"
	s = replaceSpaces(s, "-")

	// make lower
	s = toLower(s)

	// if a length is provided, truncate to that length
	s = truncate(s, length, "")

	// ensure no misbehaving "-"
	s = trimChars(s, "-")

	return s
}

// splitCamelCase splits a camelCase or PascalCase string into space-separated words using predefined rules.
func splitCamelCase(s string) string {
	entries := camelcase.Split(s)
	return strings.Join(entries, " ")
}

// splitPascalCase splits a PascalCase or camelCase
// string into space-separated words by leveraging the splitCamelCase function.
func splitPascalCase(s string) string {
	return splitCamelCase(s)
}

// toSnakeCase converts a string to snake_case or SCREAMING_SNAKE_CASE based on the scream parameter.
// It normalizes diacritical marks before formatting the string.
func toSnakeCase(s string, scream bool) string {
	s = normalizeDiacritics(s)
	if !scream {
		s = strcase.ToSnake(s)
	}
	if scream {
		s = strcase.ToScreamingSnake(s)
	}

	return s
}

// toSnakeCaseWithIgnore converts a string to snake_case format, optionally
// in uppercase, and ignores the specified ignore charset.
func toSnakeCaseWithIgnore(s string, scream bool, ignore string) string {
	s = normalizeDiacritics(s)
	if !scream {
		s = strcase.ToSnakeWithIgnore(s, ignore)
	}
	if scream {
		s = toUpper(strcase.ToSnakeWithIgnore(s, ignore))
	}

	return s
}

// capitalize returns the input string with the first character converted to uppercase while preserving the rest as is.
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return toUpper(s[:1]) + s[1:]
}

// uncapitalize converts the first character of the input string to lowercase and returns the modified string.
func uncapitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return toLower(s[:1]) + s[1:]
}

// toKebabCase converts a string to kebab-case or screaming-kebab-case depending on the scream flag.
// The string is first normalized to remove diacritics before converting to the desired case format.
func toKebabCase(s string, scream bool) string {
	s = normalizeDiacritics(s)
	if !scream {
		s = strcase.ToKebab(s)
	}
	if scream {
		s = strcase.ToScreamingKebab(s)
	}

	return s
}

// toTitleCase converts a string to title case, capitalizing the first
// letter of each word based on English language rules.
func toTitleCase(s string) string {
	s = Trim(s)
	if CamelCaseRegex.MatchString(s) {
		s = splitCamelCase(s)
	}
	return cases.Title(language.English).String(s)
}

// toCamelCase converts a string to camel case format where
// the first letter is lowercase and subsequent words are capitalized.
func toCamelCase(s string) string {
	s = normalizeDiacritics(s)
	return strcase.ToLowerCamel(s)
}

// toPascalCase converts a string to PascalCase format, where each word starts with an uppercase letter.
func toPascalCase(s string) string {
	s = normalizeDiacritics(s)
	return strcase.ToCamel(s)
}

// toDelimited converts a string to a delimited format using the specified delimiter and casing options.
func toDelimited(s string, delim uint8, ignore string, scream bool) string {
	s = normalizeDiacritics(s)
	return strcase.ToScreamingDelimited(s, delim, ignore, scream)
}

func levenshteinDistance(s1, s2 string) int {
	return edlib.LevenshteinDistance(s1, s2)
}
