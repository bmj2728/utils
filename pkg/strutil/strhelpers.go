package strutil

import (
	"strings"
	"unicode"

	godiacritics "github.com/Regis24GmbH/go-diacritics"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mrz1836/go-sanitize"
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

// toUpper converts the input string to uppercase and returns the result.
func toUpper(s string) string {
	return strings.ToUpper(s)
}

// toLower converts all characters in the input string to their lowercase equivalents and returns the resulting string.
func toLower(s string) string {
	return strings.ToLower(s)
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
