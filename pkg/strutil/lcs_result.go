package strutil

import "fmt"

type LCSResult struct {
	resultType LCSResultType
	string1    string
	string2    string
	result     *[]string
	err        error
}

func NewLCSResult(resultType LCSResultType, string1 string, string2 string, result *[]string) *LCSResult {
	return &LCSResult{
		resultType: resultType,
		string1:    string1,
		string2:    string2,
		result:     result,
	}
}

func (lcs *LCSResult) GetType() LCSResultType {
	return lcs.resultType
}

func (lcs *LCSResult) GetString1() string {
	return lcs.string1
}

func (lcs *LCSResult) GetString2() string {
	return lcs.string2
}

func (lcs *LCSResult) GetStrings() (string, string) {
	return lcs.string1, lcs.string2
}

func (lcs *LCSResult) Error() error {
	return lcs.err
}

func (lcs *LCSResult) GetResult() *[]string {
	return lcs.result
}

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
