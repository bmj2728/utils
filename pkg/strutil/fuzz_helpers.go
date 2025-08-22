package strutil

import "github.com/hbollon/go-edlib"

// fuzzyMatch performs a fuzzy search on the compList to match the given str using the specified algorithm.
// Returns the closest matching string and an error if the operation fails.
func fuzzyMatch(str string, compList []string, algo Algorithm) (string, error) {
	if algo == Hamming {
		compList = hammingCleanse(str, compList)
	}
	return edlib.FuzzySearch(str, compList, edlib.Algorithm(algo))
}

// fuzzyMatchWithThreshold performs a fuzzy match on a string against a list, returning the best
// match above a similarity threshold.
// str is the input string to match, compList is the comparison list, minSim is the similarity threshold,
// algo is the algorithm used.
func fuzzyMatchWithThreshold(str string, compList []string, minSim float32, algo Algorithm) (string, error) {
	if algo == Hamming {
		compList = hammingCleanse(str, compList)
	}
	return edlib.FuzzySearchThreshold(str, compList, minSim, edlib.Algorithm(algo))
}

// fuzzyMatchSet performs a fuzzy search on compList using the specified Algorithm and returns the top quantity matches.
// str is the input string to be matched against compList.
// compList is a collection of strings to compare with the input string.
// quantity specifies the maximum number of results to return.
// algo defines the algorithm used for string distance or scoring.
// Returns a list of matching strings and an error if any occurs during processing.
func fuzzyMatchSet(str string, compList []string, quantity int, algo Algorithm) ([]string, error) {
	if algo == Hamming {
		compList = hammingCleanse(str, compList)
	}
	return edlib.FuzzySearchSet(str, compList, quantity, edlib.Algorithm(algo))
}

// fuzzyMatchSetWithThreshold finds a set of strings from the comparison list that match the input string
// within a similarity threshold.
// str is the input string to match.
// compList is the list of strings to compare against.
// quantity specifies the maximum number of matches to return.
// minSim defines the minimum similarity threshold for matches.
// algo determines the string matching algorithm to use.
// Returns a slice of matched strings or an error if the operation fails.
func fuzzyMatchSetWithThreshold(str string,
	compList []string,
	quantity int,
	minSim float32,
	algo Algorithm) ([]string, error) {
	if algo == Hamming {
		compList = hammingCleanse(str, compList)
	}
	return edlib.FuzzySearchSetThreshold(str, compList, quantity, minSim, edlib.Algorithm(algo))
}

// hammingCleanse filters the compWords slice to include only strings with the same length as the input str.
func hammingCleanse(str string, compWords []string) []string {
	needs := len(str)
	cleansed := make([]string, 0)
	for _, word := range compWords {
		if len(word) == needs {
			cleansed = append(cleansed, word)
		}
	}
	return cleansed
}
