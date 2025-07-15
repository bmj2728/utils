package strutil

import (
	lorelai "github.com/UltiRequiem/lorelai/pkg"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mrz1836/go-sanitize"
	"math/rand"
	"net/mail"
	"net/url"
	"strings"
	"unicode"
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

// cleanWhitespace removes all whitespace characters (spaces, tabs, newlines, etc.) from the input string and returns the result.
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

// collapseWhitespace reduces all consecutive whitespace characters in the input string to a single space preserving leading and trailing whitespace.
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

// isValidEmail checks if the provided string s is a valid email address format and returns true if valid, false otherwise.
func isValidEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

// isValidUrl checks if the provided string is a valid URL with a defined scheme and host. Returns true if valid, false otherwise.
func isValidUrl(s string) bool {
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

// isValidUUID checks if the provided string is a valid UUID. Returns true if valid, otherwise false.
func isValidUUID(s string) bool {
	err := uuid.Validate(s)
	if err != nil {
		return false
	}
	return true
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

// isAlphaString checks if the given string consists only of alphabetic characters. Returns true if all characters are letters.
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

//TODO: Needs tests

// alpha removes all non-alphabetic characters from the given string, optionally retaining whitespace if ws is true.
func alpha(s string, ws bool) string {
	return sanitize.Alpha(s, ws)
}

// alphaNumeric removes all non-alphanumeric characters from the input string, optionally preserving whitespace if ws is true.
func alphaNumeric(s string, ws bool) string {
	return sanitize.AlphaNumeric(s, ws)
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

// sanitizeHTMLCustom sanitizes the input HTML string by allowing only the specified elements in allowedElements.
func sanitizeHTMLCustom(s string, allowedElements []string) string {
	p := bluemonday.NewPolicy()
	//TODO: extend implementation to better address complex options
	p.AllowElements(allowedElements...)
	return p.Sanitize(s)
}

// loremWord generates and returns a single random word as a string.
func loremWord() string {
	return lorelai.Word()
}

// loremWords generates a string containing the specified number of lorem ipsum words. If less than 1 returns a nil string
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

// loremSentenceCustom generates a lorem ipsum sentence with the specified number of words determined by the length parameter.
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

// loremSentenceVariable generates a specified number of lorem ipsum sentences with random lengths between min and max values.
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
// The parameter 'count' determines the number of paragraphs to generate. Returns an empty string if count is less than 1.
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
