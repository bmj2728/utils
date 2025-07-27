package strutil

import (
	"fmt"
)

// ComparisonResultScore represents a type constraint for a pointer to either an int or a float32.
type ComparisonResultScore interface {
	*int | *float32
}

// ComparisonResultType represents various types of comparison results for string similarity and distance measurements.
type ComparisonResultType int

// String returns the string representation of the ComparisonResultType value using the ComparisonResultTypeMap.
func (c ComparisonResultType) String() string {
	return ComparisonResultTypeMap[c]
}

// LCSLength represents the comparison result type using Longest Common Subsequence length.
// LCSDist represents the comparison result type using Longest Common Subsequence distance.
// LevDist represents the comparison result type using Levenshtein distance.
// DamLevDist represents the comparison result type using Damerau-Levenshtein distance.
// OSADamLevDist represents the comparison result type using Optimal String Alignment distance.
// HammingDist represents the comparison result type using Hamming distance.
// JaroSim represents the comparison result type using Jaro similarity.
// JaroWinklerSim represents the comparison result type using Jaro-Winkler similarity.
// JaccardSim represents the comparison result type using Jaccard similarity.
// CosineSim represents the comparison result type using Cosine similarity.
// SorensenDiceCo represents the comparison result type using Sørensen-Dice coefficient.
// QGramDist represents the comparison result type using Q-Gram distance with default q-gram size.
// QGramDistCust represents the comparison result type using Q-Gram distance with custom q-gram size.
// QGramSim represents the comparison result type using Q-Gram similarity.
const (
	LCSLength ComparisonResultType = iota
	LCSDist
	LevDist
	DamLevDist
	OSADamLevDist
	HammingDist
	JaroSim
	JaroWinklerSim
	JaccardSim
	CosineSim
	SorensenDiceCo
	QGramDist
	QGramDistCust
	QGramSim
)

var ComparisonResultTypeMap = map[ComparisonResultType]string{
	LCSLength:      "LCS Length",
	LCSDist:        "LCS Distance",
	LevDist:        "Levenshtein Distance",
	DamLevDist:     "Damerau-Levenshtein Distance",
	OSADamLevDist:  "OSA Damerau-Levenshtein Distance",
	HammingDist:    "Hamming Distance",
	JaroSim:        "Jaro Similarity",
	JaroWinklerSim: "Jaro-Winkler Similarity",
	JaccardSim:     "Jaccard Similarity",
	CosineSim:      "Cosine Similarity",
	SorensenDiceCo: "Sorensen-Dice Coefficient",
	QGramDist:      "Q-Gram Distance",
	QGramDistCust:  "Q-Gram Distance Custom",
	QGramSim:       "Q-Gram Similarity",
}

// ComparisonResult defines an interface for comparing two strings and retrieving results, types, and errors.
type ComparisonResult interface {
	GetType() ComparisonResultType
	GetTypeName() string
	GetString1() string
	GetString2() string
	GetStrings() (string, string)
	GetSplitLength() (int, error)
	GetError() error
	Print(v bool)
	formatOutput(v bool) string
}

// ComparisonResultInt represents the result of a comparison between two strings,
// including type, score, and error details.
type ComparisonResultInt struct {
	comparisonType ComparisonResultType
	string1        string
	string2        string
	splitLength    *int
	score          *int
	err            error
}

func NewComparisonResultInt(comparisonType ComparisonResultType,
	string1 string,
	string2 string,
	splitLength *int,
	score *int,
	error error) *ComparisonResultInt {
	return &ComparisonResultInt{
		comparisonType: comparisonType,
		string1:        string1,
		string2:        string2,
		splitLength:    splitLength,
		score:          score,
		err:            error,
	}
}

// GetType returns the ComparisonResultType associated with the ComparisonResultInt instance.
func (c ComparisonResultInt) GetType() ComparisonResultType {
	return c.comparisonType
}

// GetTypeName retrieves the string representation of the comparison type for the ComparisonResultInt instance.
func (c ComparisonResultInt) GetTypeName() string {
	return c.comparisonType.String()
}

// GetString1 retrieves the first string (string1) associated with the ComparisonResultInt instance.
func (c ComparisonResultInt) GetString1() string {
	return c.string1
}

// GetString2 returns the second comparison string from the ComparisonResultInt instance.
func (c ComparisonResultInt) GetString2() string {
	return c.string2
}

// GetStrings returns the two strings, string1 and string2, stored in the ComparisonResultInt instance.
func (c ComparisonResultInt) GetStrings() (string, string) {
	return c.string1, c.string2
}

// GetSplitLength returns the split length of the comparison result as a pointer to an integer.
func (c ComparisonResultInt) GetSplitLength() (int, error) {
	if c.splitLength == nil {
		return 0, ErrNoSplitLengthSet
	}
	return *c.splitLength, nil
}

// GetError returns the error associated with the ComparisonResultInt, if any.
func (c ComparisonResultInt) GetError() error {
	return c.err
}

// GetScoreInt retrieves the comparison score as an integer and returns an error if no score or an error is present.
func (c ComparisonResultInt) GetScoreInt() (int, error) {
	if c.score == nil && c.err == nil {
		return 0, ErrUnknownError
	}
	if c.err != nil {
		return 0, c.err
	}
	return *c.score, nil
}

// IsMatch checks if two ComparisonResultInt objects are equivalent by
// comparing their type, strings, score, split length, and errors.
func (c ComparisonResultInt) IsMatch(other ComparisonResultInt) bool {
	// check required fields
	if c.GetType() != other.GetType() || c.GetString1() != other.GetString1() || c.GetString2() != other.GetString2() {
		return false
	}
	// GetScoreInt will return a score/0 + error/nil
	cScore, cErr := c.GetScoreInt()
	oScore, oErr := other.GetScoreInt()
	// check errors first - this is the underlying ComparisonResult error if not null
	if cErr != nil || oErr != nil {
		if cErr == nil || oErr == nil {
			return false
		}
		if cErr.Error() != oErr.Error() {
			return false
		}
	}
	// check int score
	if cScore != oScore {
		return false

	}
	// skip redundant error check - we returned already if there was an error on the builder
	cSplit, _ := c.GetSplitLength()
	oSplit, _ := other.GetSplitLength()
	if cSplit != oSplit {
		return false
	}
	// no returns, we must match
	return true
}

// Print outputs the formatted result of the comparison, optionally including
// verbose error details based on the input flag v.
func (c ComparisonResultInt) Print(v bool) {
	fmt.Print(c.formatOutput(v))
}

// formatOutput generates a formatted string representation of the comparison result
// based on error presence and verbosity.
func (c ComparisonResultInt) formatOutput(v bool) string {
	if c.err != nil && v {
		return fmt.Sprintf("Error during processing: %s\nFirst String: %s\nSecond String: %s\nError: %s\n",
			ComparisonResultTypeMap[c.comparisonType], c.string1, c.string2, c.err.Error())
	} else if c.err != nil && !v {
		return fmt.Sprintf("%s Error: %s\n", ComparisonResultTypeMap[c.comparisonType], c.err.Error())
	}
	if v {
		return fmt.Sprintf("Comparison: %s\nFirst String: %s\nSecond String: %s\nScore: %d\n",
			ComparisonResultTypeMap[c.comparisonType], c.string1, c.string2, *c.score)
	} else {
		return fmt.Sprintf("%s: %d\n", ComparisonResultTypeMap[c.comparisonType], *c.score)
	}
}

// ComparisonResultFloat represents the result of a comparison operation,
// including its type, input strings, score, and an optional error.
type ComparisonResultFloat struct {
	comparisonType ComparisonResultType
	string1        string
	string2        string
	splitLength    *int
	score          *float32
	err            error
}

func NewComparisonResultFloat(comparisonType ComparisonResultType,
	string1 string,
	string2 string,
	splitLength *int,
	score *float32,
	error error) *ComparisonResultFloat {
	return &ComparisonResultFloat{
		comparisonType: comparisonType,
		string1:        string1,
		string2:        string2,
		splitLength:    splitLength,
		score:          score,
		err:            error,
	}
}

// GetType returns the type of comparison as a value of ComparisonResultType.
func (c ComparisonResultFloat) GetType() ComparisonResultType {
	return c.comparisonType
}

// GetTypeName returns the string representation of the comparison type from the comparisonType field.
func (c ComparisonResultFloat) GetTypeName() string {
	return c.comparisonType.String()
}

// GetString1 returns the first string (string1) stored in the ComparisonResultFloat instance.
func (c ComparisonResultFloat) GetString1() string {
	return c.string1
}

// GetString2 returns the second string associated with the ComparisonResultFloat instance.
func (c ComparisonResultFloat) GetString2() string {
	return c.string2
}

// GetStrings returns the two strings stored in the ComparisonResultFloat instance.
func (c ComparisonResultFloat) GetStrings() (string, string) {
	return c.string1, c.string2
}

// GetSplitLength retrieves the value of the splitLength field as a pointer to an integer.
func (c ComparisonResultFloat) GetSplitLength() (int, error) {
	if c.splitLength == nil {
		return 0, ErrNoSplitLengthSet
	}
	return *c.splitLength, nil
}

// GetError returns the error encountered during the comparison, or nil if no error occurred.
func (c ComparisonResultFloat) GetError() error {
	return c.err
}

// GetScoreFloat retrieves the comparison score as a float32 and any associated error.
// Returns 0.00 and an error if unavailable.
func (c ComparisonResultFloat) GetScoreFloat() (float32, error) {
	if c.score == nil && c.err == nil {
		return 0.00, ErrUnknownError
	}
	if c.err != nil {
		return 0.00, c.err
	}
	return *c.score, nil
}

// Print outputs the formatted comparison result or error details depending on the verbosity flag (v).
func (c ComparisonResultFloat) Print(v bool) {
	fmt.Print(c.formatOutput(v))
}

// formatOutput formats the comparison result or error into a string based on the verbosity level (v).
func (c ComparisonResultFloat) formatOutput(v bool) string {
	if v {
		if c.err != nil {
			return fmt.Sprintf("Error during processing %s\nFirst String: %s\nSecond String: %s\nError: %s\n",
				ComparisonResultTypeMap[c.comparisonType], c.string1, c.string2, c.err.Error())

		} else {
			return fmt.Sprintf("Comparison: %s\nFirst String: %s\nSecond String: %s\nScore: %f\n",
				ComparisonResultTypeMap[c.comparisonType], c.string1, c.string2, *c.score)
		}
	} else {
		if c.err != nil {
			return fmt.Sprintf("%s Error: %s\n",
				ComparisonResultTypeMap[c.comparisonType], c.err.Error())

		} else {
			return fmt.Sprintf("%s: %f\n", ComparisonResultTypeMap[c.comparisonType], *c.score)
		}
	}
}
