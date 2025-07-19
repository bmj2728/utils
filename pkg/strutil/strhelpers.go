package strutil

import (
	"errors"
	"net/mail"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"unicode"

	godiacritics "github.com/Regis24GmbH/go-diacritics"
	"github.com/fatih/camelcase"
	"github.com/google/uuid"
	"github.com/hbollon/go-edlib"
	"github.com/iancoleman/strcase"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mrz1836/go-sanitize"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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

// cleanWhitespaceWithIgnore removes all whitespace characters from the input string
// while ignoring whitespace characters in the given charset.
func cleanWhitespaceWithIgnore(s string, charset string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, c := range s {
		if !unicode.IsSpace(c) || (unicode.IsSpace(c) && strings.ContainsRune(charset, c)) {
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

// appendString concatenates the given string `s` with `suffix`,
// separated by the specified `sep`, and returns the result.
func appendString(s string, suffix string, sep string) string {
	if len(suffix) < 1 {
		return s
	}
	return s + sep + suffix
}

// prependString concatenates a prefix and a separator to the beginning of a string and returns the resulting string.
func prependString(s string, prefix string, sep string) string {
	if len(prefix) < 1 {
		return s
	}
	return prefix + sep + s
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

// levenshteinDistance calculates the Levenshtein distance between two strings s1 and s2.
// It determines the minimum number of single-character edits (insertions, deletions, or substitutions) required.
func levenshteinDistance(s1, s2 string) int {
	return edlib.LevenshteinDistance(s1, s2)
}

// damerauLevenshteinDistance calculates the Damerau-Levenshtein distance
// between two strings to measure edit similarity.
func damerauLevenshteinDistance(str1, str2 string) int {
	return edlib.DamerauLevenshteinDistance(str1, str2)
}

// osaDamerauLevenshteinDistance calculates the Damerau-Levenshtein distance between str1 and str2,
// considering adjacent transpositions.
//
// This optimal string alignment variant of damerauLevenshteinDistance
// does not allow multiple transformations on the same substring
func osaDamerauLevenshteinDistance(str1, str2 string) int {
	return edlib.OSADamerauLevenshteinDistance(str1, str2)
}

// lcs returns the length of the longest common subsequence between two input strings, str1 and str2.
func lcs(str1 string, str2 string) int {
	return edlib.LCS(str1, str2)
}

// lcsBacktrack computes the longest common subsequence (LCS) of two input strings and returns the result as a string.
func lcsBacktrack(str1 string, str2 string) (string, error) {
	result, err := edlib.LCSBacktrack(str1, str2)
	if err != nil {
		return "", errors.Join(ErrLCSBacktrackFailure, err)
	}
	return result, nil
}

// lcsBacktrackAll computes all longest common subsequences of two input strings.
// Returns a slice of strings representing the subsequences and an error if the computation fails.
func lcsBacktrackAll(str1 string, str2 string) ([]string, error) {
	result, err := edlib.LCSBacktrackAll(str1, str2)
	if err != nil {
		return nil, errors.Join(ErrLCSBacktrackAllFailure, err)
	}
	return result, nil
}

// lcsDiff calculates the Longest Common Subsequence (LCS) diff between
// two strings and returns the differences or an error.
func lcsDiff(str1, str2 string) ([]string, error) {
	result, err := edlib.LCSDiff(str1, str2)
	if err != nil {
		return nil, errors.Join(ErrLCSDiffFailure, err)
	}
	return result, nil
}

// lcsEditDistance computes the edit distance between two strings based on their Longest Common Subsequence (LCS).
func lcsEditDistance(s1, s2 string) int {
	return edlib.LCSEditDistance(s1, s2)
}

// compareStringSlices checks if two slices of strings contain the same elements,
// regardless of order, with optional nil equality.
// If `nulls` is true, nil slices are treated as equal.
// It returns true if slices match, otherwise false.
func compareStringSlices(s1, s2 []string, nulls bool) bool {
	if nulls && s1 == nil && s2 == nil {
		return true
	}
	if s1 == nil || s2 == nil {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	remaining := slices.Clone(s2)
	for _, s := range s1 {
		if !slices.Contains(remaining, s) {
			return false
		} else {
			for i, w := range remaining {
				if s == w {
					remaining = slices.Delete(remaining, i, i+1)
					break
				}
			}
		}
		continue
	}
	return true
}

// compareStringBuilderSlices compares two slices of StringBuilder for equality,
// optionally considering nil slices as equal. The comparison ignores the order of elements
// and uses the 'nulls' flag to determine nil-handling behavior.
func compareStringBuilderSlices(s1, s2 []StringBuilder, nulls bool) bool {
	if nulls && s1 == nil && s2 == nil {
		return true
	}
	if s1 == nil || s2 == nil {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	remaining := slices.Clone(s2)
	for _, s := range s1 {
		if !slices.Contains(s2, s) {
			return false
		} else {
			for i, w := range remaining {
				if s == w {
					remaining = slices.Delete(remaining, i, i+1)
					break
				}
			}
		}
	}
	return true
}

// hammingDistance computes the Hamming distance between two strings s1 and s2.
// Returns a pointer to the integer distance and an error if the strings have unequal lengths or a failure occurs.
func hammingDistance(s1, s2 string) (*int, error) {
	dist, err := edlib.HammingDistance(s1, s2)
	if err != nil {
		return nil, errors.Join(ErrHammingDistanceFailure, err)
	}
	return &dist, nil
}

// jaroSimilarity computes the Jaro similarity between
// two strings s1 and s2, returning a float32 value in the range [0, 1].
func jaroSimilarity(s1, s2 string) float32 {
	return edlib.JaroSimilarity(s1, s2)
}

// jaroWinklerSimilarity computes the Jaro-Winkler similarity between two strings s1 and s2.
// It returns a float32 value between 0 and 1, where 1 indicates exact similarity.
func jaroWinklerSimilarity(s1, s2 string) float32 {
	return edlib.JaroWinklerSimilarity(s1, s2)
}

// jaccardSimilarity computes the Jaccard similarity coefficient between two strings based on a given split length.
// If split length of zero, the string is split on whitespaces and returns an index
func jaccardSimilarity(s1, s2 string, splitLength int) *float32 {
	if splitLength < 0 {
		return nil
	}
	js := edlib.JaccardSimilarity(s1, s2, splitLength)
	return &js
}

// cosineSimilarity computes the cosine similarity between two strings using a specified n-gram split length.
// Returns a pointer to the similarity score or nil if the split length is negative.
// Split lengths of zero, split on whitespaces.
func cosineSimilarity(s1, s2 string, splitLength int) *float32 {
	if splitLength < 0 {
		return nil
	}
	cs := edlib.CosineSimilarity(s1, s2, splitLength)
	return &cs
}

// sorensenDiceCoefficient calculates the Sørensen–Dice coefficient between two strings with a specified n-gram length.
// Returns a pointer to the coefficient value or nil if the splitLength is negative.
func sorensenDiceCoefficient(s1, s2 string, splitLength int) *float32 {
	if splitLength < 0 {
		return nil
	}
	sdc := edlib.SorensenDiceCoefficient(s1, s2, splitLength)
	return &sdc
}
