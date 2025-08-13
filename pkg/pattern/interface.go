package pattern

// BasicPatternGroup defines behavior for matching patterns and extracting substrings from input strings.
// Match checks if the input string satisfies the pattern, with an option for strict matching.
// Extract retrieves the first substring from the input string that matches the defined pattern.
// ExtractAll retrieves all substrings from the input string that match the defined pattern.
// Should be implemented as a struct containing strict, lenient, and extract patterns
// Example:
//
//	 type EmailPatternGroup struct {
//		strictPattern *regexp.Regexp
//		lenientPattern *regexp.Regexp
//		extractPattern *regexp.Regexp
//		}
type BasicPatternGroup interface {
	Match(input string, strict bool) bool
	Extract(input string) string
	ExtractAll(input string) []string
}

// ComplexPatternGroup defines an interface for managing and processing groups of complex string patterns.
// Match checks if the input string matches a custom pattern identified by its name in the provided pattern set.
// Extract returns the first substring from the input matching a custom pattern specified
// by its name in the pattern set.
// ExtractAll returns all substrings from the input matching a custom pattern specified by its name in the pattern set.
// Should be implemented as a struct containing a custom pattern set.
// Example:
//
//	 type PostalCodePatternGroup struct {
//		postalCodeSet *CustomPatternSet
//		}
type ComplexPatternGroup interface {
	Match(input string, name string) (bool, error)
	Extract(input string, name string) (string, error)
	ExtractAll(input string, name string) ([]string, error)
}
