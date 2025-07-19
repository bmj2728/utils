package strutil

import (
	"strings"

	"github.com/fatih/camelcase"
	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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

// toUpper converts the input string to uppercase and returns the result.
func toUpper(s string) string {
	return strings.ToUpper(s)
}

// toLower converts all characters in the input string to their lowercase equivalents and returns the resulting string.
func toLower(s string) string {
	return strings.ToLower(s)
}
