package pattern

import (
	"regexp"

	"utils/pkg/internal/errors"
)

// CustomPatternSet is a type that maps string keys to compiled regular expressions.
// it's strongly recommended to compile regex patterns once and use throughout the application.
type CustomPatternSet map[string]*regexp.Regexp

// NewCustomPatternSet initializes and returns an empty CustomPatternSet for storing named regex patterns.
func NewCustomPatternSet() CustomPatternSet {
	return make(map[string]*regexp.Regexp)
}

// Add adds a new pattern to the CustomPatternSet with the given name as the key and pattern as its value.
func (cps CustomPatternSet) Add(name string, pattern *regexp.Regexp) {

	cps[name] = pattern
}

// Remove deletes the pattern associated with the given name from the CustomPatternSet.
func (cps CustomPatternSet) Remove(name string) {
	delete(cps, name)
}

// Get retrieves a compiled regular expression associated with the provided name.
// Returns an error if the name is not found.
func (cps CustomPatternSet) Get(name string) (*regexp.Regexp, error) {
	pattern, ok := cps[name]
	if !ok {
		return nil, errors.ErrPatternNotFound
	}
	return pattern, nil
}

// GetNames returns a slice of all keys (names) present in the CustomPatternSet.
func (cps CustomPatternSet) GetNames() []string {
	var names []string
	for name := range cps {
		names = append(names, name)
	}
	return names
}

// Exists checks if a pattern with the given name exists in the CustomPatternSet and
// returns true if found, otherwise false.
func (cps CustomPatternSet) Exists(name string) bool {
	_, ok := cps[name]
	return ok
}
