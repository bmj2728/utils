package strutil

import (
	"fmt"

	"utils/pkg/internal/comparison"
	"utils/pkg/internal/errors"
)

// LCSResultType represents an enumerated type used to define different
// result types for Longest Common Subsequence calculations.
type LCSResultType int

// String returns the string representation for the corresponding LCSResultType using LCSResultTypeMap.
func (l LCSResultType) String() string {
	return LCSResultTypeMap[l]
}

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
func NewLCSResult(resultType LCSResultType, string1 string, string2 string, result *[]string, err error) *LCSResult {
	return &LCSResult{
		resultType: resultType,
		string1:    string1,
		string2:    string2,
		result:     result,
		err:        err,
	}
}

// GetType returns the result type of the LCSResult, indicating the type of LCS computation performed.
func (lcs *LCSResult) GetType() LCSResultType {
	return lcs.resultType
}

func (lcs *LCSResult) GetTypeName() string {
	return lcs.resultType.String()
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

// GetError returns the error associated with the LCSResult instance, if any.
func (lcs *LCSResult) GetError() error {
	return lcs.err
}

// GetResult retrieves the pointer to the list of longest common subsequence results stored in the LCSResult instance.
func (lcs *LCSResult) GetResult() []string {
	if lcs.result == nil {
		return nil
	}
	return *lcs.result
}

func (lcs *LCSResult) IsMatch(other *LCSResult) bool {
	if lcs == nil || other == nil {
		return false
	}
	if lcs.resultType != other.resultType ||
		lcs.string1 != other.string1 ||
		lcs.string2 != other.string2 ||
		!errors.CompareErrors(lcs.err, other.err) ||
		!comparison.CompareStringSlices(lcs.GetResult(), other.GetResult(), false) {
		return false
	}
	return true
}

// Print outputs the LCSResult details to the console based on the provided verbosity flag.
func (lcs *LCSResult) Print(v bool) {
	fmt.Print(formatLCSResultOutput(lcs, v))
}

func formatLCSResultOutput(lcs *LCSResult, v bool) string {
	if lcs == nil {
		return ""
	}
	if lcs.GetError() != nil {
		if !v {
			return fmt.Sprintf("GetError calculating %s (%s/%s): %s\n",
				lcs.GetTypeName(), lcs.GetString1(), lcs.GetString2(), lcs.GetError().Error())
		} else {
			return fmt.Sprintf("GetError calculating %s:\nFirst Word: %s\nSecond Word:%s\nGetError: %s\n",
				lcs.GetTypeName(), lcs.GetString1(), lcs.GetString2(), lcs.GetError().Error())
		}
	}
	switch lcs.GetType() {
	case LCSBacktrackWord:
		if v {
			return fmt.Sprintf("%s:\nFirst: %s\nSecond: %s\nWord: %s\n",
				lcs.GetTypeName(), lcs.GetString1(), lcs.GetString2(), lcs.GetResult()[0])
		} else {
			return fmt.Sprintf("%s: %s\n", lcs.GetTypeName(), lcs.GetResult()[0])
		}
	case LCSBacktrackWordAll:
		if v {
			s := fmt.Sprintf("%s:\nFirst: %s\nSecond: %s\nWords:\n",
				lcs.GetTypeName(), lcs.GetString1(), lcs.GetString2())
			for _, word := range lcs.GetResult() {
				s += fmt.Sprintf("%s\n", word)
			}
			return s
		} else {
			return fmt.Sprintf("%s(%d):\nFirst: %s\n",
				lcs.GetTypeName(), len(lcs.GetResult()), lcs.GetResult()[0])
		}
	case LCSDiffSlice:
		if v {
			s := fmt.Sprintf("%s (%s/%s):\n",
				lcs.GetTypeName(), lcs.GetString1(), lcs.GetString2())
			for _, diff := range lcs.GetResult() {
				s += fmt.Sprintf("%s\n", diff)
			}
			return s
		} else {
			s := fmt.Sprintf("%s:\n", lcs.GetTypeName())
			for _, diff := range lcs.GetResult() {
				s += fmt.Sprintf("%s\n", diff)
			}
			return s
		}
	default:
		return "Unknown LCS Result Type"
	}
}
