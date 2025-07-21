package strutil

import "fmt"

// LCSResultType represents an enumerated type used to define different
// result types for Longest Common Subsequence calculations.
type LCSResultType int

// LCSBacktrackWord indicates backtracking for a single LCS word result.
// LCSBacktrackWordAll indicates backtracking for all LCS word results.
// LCSDiffSlice represents an LCS result in the form of a diff slice.
const (
	LCSBacktrackWord LCSResultType = iota
	LCSBacktrackWordAll
	LCSDiffSlice
)

// LCSResultTypeMap maps LCSResultType constants to their corresponding string representations.
var LCSResultTypeMap = map[LCSResultType]string{
	LCSBacktrackWord:    "LCS Backtrack",
	LCSBacktrackWordAll: "LCS Backtrack All",
	LCSDiffSlice:        "LCS Diff",
}

// LCSResult encapsulates the result of a Longest Common Subsequence (LCS) computation between two strings.
type LCSResult struct {
	resultType LCSResultType
	string1    string
	string2    string
	result     *[]string
	err        error
}

// NewLCSResult creates and returns a pointer to an LCSResult with the specified type, strings, and result slice.
func NewLCSResult(resultType LCSResultType, string1 string, string2 string, result *[]string) *LCSResult {
	return &LCSResult{
		resultType: resultType,
		string1:    string1,
		string2:    string2,
		result:     result,
	}
}

// GetType returns the result type of the LCSResult, indicating the type of LCS computation performed.
func (lcs *LCSResult) GetType() LCSResultType {
	return lcs.resultType
}

// GetString1 returns the value of string1 from the LCSResult instance.
func (lcs *LCSResult) GetString1() string {
	return lcs.string1
}

// GetString2 retrieves the second string (string2) associated with the LCSResult.
func (lcs *LCSResult) GetString2() string {
	return lcs.string2
}

// GetStrings returns the two input strings stored in the LCSResult instance.
func (lcs *LCSResult) GetStrings() (string, string) {
	return lcs.string1, lcs.string2
}

// Error returns the error associated with the LCSResult instance, if any.
func (lcs *LCSResult) Error() error {
	return lcs.err
}

// GetResult retrieves the pointer to the list of longest common subsequence results stored in the LCSResult instance.
func (lcs *LCSResult) GetResult() *[]string {
	return lcs.result
}

// Print outputs the LCSResult details to the console based on the provided verbosity flag.
func (lcs *LCSResult) Print(v bool) {
	if v {
		if lcs.err != nil {
			fmt.Printf("Error processing %s\nString 1: %s\nString 2: %s\nError: %s\n",
				LCSResultTypeMap[lcs.resultType], lcs.string1, lcs.string2, lcs.err.Error())
			return
		} else {
			switch lcs.resultType {
			case LCSBacktrackWord:
				fmt.Printf("%s\nString 1: %s\nString 2: %s\nLongest Common Substring: %s\n",
					LCSResultTypeMap[lcs.resultType], lcs.string1, lcs.string2, (*lcs.result)[0])
				return
			case LCSBacktrackWordAll:
				fmt.Printf("%s\nString 1: %s\nString 2: %s\nLongest Common Substrings:\n",
					LCSResultTypeMap[lcs.resultType], lcs.string1, lcs.string2)
				for _, lcs := range *lcs.result {
					fmt.Printf("%s\n", lcs)
				}
				return
			case LCSDiffSlice:
				fmt.Printf("%s\nString 1: %s\nString 2: %s\n",
					LCSResultTypeMap[lcs.resultType], lcs.string1, lcs.string2)
				for _, lcs := range *lcs.result {
					fmt.Printf("%s\n", lcs)
				}
				return
			default:
				fmt.Printf("Error processing %s\nString 1: %s\nString 2: %s\nError: %s\n",
					LCSResultTypeMap[lcs.resultType], lcs.string1, lcs.string2, lcs.err.Error())
				return
			}
		}
	} else {
		if lcs.err != nil {
			fmt.Printf("%s\nError: %s\n",
				LCSResultTypeMap[lcs.resultType], lcs.err.Error())
			return
		} else {
			switch lcs.resultType {
			case LCSBacktrackWord:
				fmt.Printf("%s\nLCS: %s\n",
					LCSResultTypeMap[lcs.resultType], (*lcs.result)[0])
				return
			case LCSBacktrackWordAll:
				fmt.Printf("%s\nLCS List:\n",
					LCSResultTypeMap[lcs.resultType])
				for _, lcs := range *lcs.result {
					fmt.Printf("%s\n", lcs)
				}
				return
			case LCSDiffSlice:
				fmt.Printf("%s\n",
					LCSResultTypeMap[lcs.resultType])
				for _, lcs := range *lcs.result {
					fmt.Printf("%s\n", lcs)
				}
				return
			default:
				fmt.Printf("%s\nError: %s\n",
					LCSResultTypeMap[lcs.resultType], lcs.err.Error())
				return
			}
		}
	}
}
