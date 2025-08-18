package strutil

// FuzzyMatch performs a fuzzy search to find the closest matching string to `str` in `compList` using the given `algo`.
// Returns the best match string and an error if the search fails.
func FuzzyMatch(str string, compList []string, algo Algorithm) (string, error) {
	return fuzzyMatch(str, compList, algo)
}

// FuzzyMatchWithThreshold performs a fuzzy matching operation on a string against a list of strings.
// Returns the best match meeting the specified similarity threshold or an error if no match qualifies.
// str is the input string to match, compList is the list of comparison strings, minSim is the similarity threshold,
// and algo specifies the string matching algorithm to use.
func FuzzyMatchWithThreshold(str string, compList []string, minSim float32, algo Algorithm) (string, error) {
	return fuzzyMatchWithThreshold(str, compList, minSim, algo)
}

// FuzzyMatchSet performs a fuzzy search on compList using the specified algorithm and returns the top quantity matches.
// str is the input string to be matched against compList.
// compList is the list of strings to compare with the input string.
// quantity specifies the maximum number of results to return.
// algo determines the algorithm used for string scoring or distance computation.
// Returns a slice of matching strings and an error if any processing issues occur.
func FuzzyMatchSet(str string, compList []string, quantity int, algo Algorithm) ([]string, error) {
	return fuzzyMatchSet(str, compList, quantity, algo)
}

// FuzzyMatchSetWithThreshold retrieves a set of strings from a list that match the input within a similarity threshold.
// str is the target string to fuzzy match against the comparison list.
// compList is the list of candidate strings for the fuzzy matching process.
// quantity specifies the maximum number of matched strings to return in the output.
// minSim is the minimum similarity threshold (0 to 1) required for a match to be included.
// algo defines the algorithm used for string similarity calculation, influencing the matching process.
// Returns a slice of matched strings based on the specified criteria or an error if the operation fails.
func FuzzyMatchSetWithThreshold(str string,
	compList []string,
	quantity int,
	minSim float32,
	algo Algorithm) ([]string, error) {
	return fuzzyMatchSetWithThreshold(str, compList, quantity, minSim, algo)
}
