package strutil

// FuzzyMatch performs a fuzzy matching process on compList using a specified algorithm and returns the best match.
// Returns an error if the operation encounters issues or if further processing is not allowed.
func (sb *StringBuilder) FuzzyMatch(compList []string, algo Algorithm) (string, error) {
	if !sb.shouldContinueProcessing() {
		return "", sb.err
	}
	return fuzzyMatch(sb.value, compList, algo)
}

// FuzzyMatchWithThreshold performs a fuzzy string match with a similarity threshold using the specified algorithm.
// compList is a list of strings to compare against the StringBuilder's value.
// minSim defines the minimum similarity threshold for a successful match.
// algo is the algorithm used to compute similarity.
func (sb *StringBuilder) FuzzyMatchWithThreshold(compList []string, minSim float32, algo Algorithm) (string, error) {
	if !sb.shouldContinueProcessing() {
		return "", sb.err
	}
	return fuzzyMatchWithThreshold(sb.value, compList, minSim, algo)
}

// FuzzyMatchSet performs a fuzzy search on a comparison list using the specified algorithm and returns top matches.
// compList is the list of strings to match against the StringBuilder's value.
// quantity specifies the maximum number of matching results to return.
// algo defines the string matching algorithm used for comparison.
// Returns a slice of matched strings and an error if the processing fails.
func (sb *StringBuilder) FuzzyMatchSet(compList []string, quantity int, algo Algorithm) ([]string, error) {
	if !sb.shouldContinueProcessing() {
		return nil, sb.err
	}
	return fuzzyMatchSet(sb.value, compList, quantity, algo)
}

// FuzzyMatchSetWithThreshold performs fuzzy matching on a list of strings with a similarity
// threshold and returns matches.
// compList is the comparison list of strings to match against.
// quantity determines the maximum number of matches to return.
// minSim is the minimum similarity threshold for a match to be considered valid.
// algo specifies the string comparison algorithm to use.
// Returns a slice of matched strings or an error if an issue occurs during execution.
func (sb *StringBuilder) FuzzyMatchSetWithThreshold(compList []string,
	quantity int,
	minSim float32,
	algo Algorithm) ([]string, error) {
	if !sb.shouldContinueProcessing() {
		return nil, sb.err
	}
	return fuzzyMatchSetWithThreshold(sb.value, compList, quantity, minSim, algo)
}
