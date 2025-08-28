package strutil

import (
	"fmt"
	"math"

	"github.com/bmj2728/utils/pkg/internal/errors"
)

//// ComparisonResultScore represents a type constraint for a pointer to either an int or a float32.
//type ComparisonResultScore interface {
//	*int | *float32
//}

// ComparisonResultType represents various types of comparison results for string score and distance measurements.
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
// JaroSim represents the comparison result type using Jaro score.
// JaroWinklerSim represents the comparison result type using Jaro-Winkler score.
// JaccardSim represents the comparison result type using Jaccard score.
// CosineSim represents the comparison result type using Cosine score.
// SorensenDiceCo represents the comparison result type using Sørensen-Dice coefficient.
// QGramDist represents the comparison result type using Q-Gram distance with default q-gram size.
// QGramDistCust represents the comparison result type using Q-Gram distance with custom q-gram size.
// QGramSim represents the comparison result type using Q-Gram score.
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

// ComparisonResult represents the result of a comparison between two strings,
// including type, score, and error details.
type ComparisonResult struct {
	comparisonType ComparisonResultType
	string1        string
	string2        string
	splitLength    *int
	scoreInt       *int
	scoreFloat     *float32
	err            error
}

// NewComparisonResult creates a new ComparisonResult instance with the provided comparison type and two strings.
func NewComparisonResult(comparisonType ComparisonResultType, string1 string, string2 string) *ComparisonResult {
	return &ComparisonResult{
		comparisonType: comparisonType,
		string1:        string1,
		string2:        string2,
		splitLength:    nil,
		scoreInt:       nil,
		scoreFloat:     nil,
		err:            nil,
	}
}

// WithSplitLength sets the split length for the comparison result and returns the updated ComparisonResult instance.
func (c *ComparisonResult) WithSplitLength(splitLength int) *ComparisonResult {
	c.splitLength = &splitLength
	return c
}

// WithScoreInt sets the integer score for the ComparisonResult,
// adds the floating-point representation of the score, and returns the instance.
func (c *ComparisonResult) WithScoreInt(score int) *ComparisonResult {
	c.scoreInt = &score
	flScore := float32(score)
	c.scoreFloat = &flScore
	return c
}

// WithScoreFloat sets the float score for the ComparisonResult,
// adds the integer representation if it can be created without precision loss, and returns the instance.
func (c *ComparisonResult) WithScoreFloat(score float32) *ComparisonResult {
	c.scoreFloat = &score
	if float64(score) == math.Trunc(float64(score)) {
		iScore := int(score)
		c.scoreInt = &iScore
	}
	return c
}

// WithError sets the error for the ComparisonResult and returns the updated instance.
func (c *ComparisonResult) WithError(err error) *ComparisonResult {
	c.err = err
	return c
}

// GetType returns the ComparisonResultType associated with the ComparisonResult instance.
func (c *ComparisonResult) GetType() ComparisonResultType {
	return c.comparisonType
}

// GetTypeName retrieves the string representation of the comparison type for the ComparisonResult instance.
func (c *ComparisonResult) GetTypeName() string {
	return c.comparisonType.String()
}

// GetString1 retrieves the first string (string1) associated with the ComparisonResult instance.
func (c *ComparisonResult) GetString1() string {
	return c.string1
}

// GetString2 returns the second comparison string from the ComparisonResult instance.
func (c *ComparisonResult) GetString2() string {
	return c.string2
}

// GetStrings returns the two strings, string1 and string2, stored in the ComparisonResult instance.
func (c *ComparisonResult) GetStrings() (string, string) {
	return c.string1, c.string2
}

// GetSplitLength returns the split length of the comparison result as a pointer to an integer.
func (c *ComparisonResult) GetSplitLength() (int, error) {
	if c.splitLength == nil {
		return 0, errors.ErrNoSplitLengthSet
	}
	return *c.splitLength, nil
}

// GetError returns the error associated with the ComparisonResult, if any.
func (c *ComparisonResult) GetError() error {
	return c.err
}

// GetScoreInt retrieves the integer score from the ComparisonResult.
// Returns -1 and an error if the score is nil or non-convertible to an integer.
func (c *ComparisonResult) GetScoreInt() (int, error) {
	// if there is an error, return -1 and the error
	if c.GetError() != nil {
		return -1, c.GetError()
	}
	// check if we have an int
	if c.scoreInt == nil {
		// if we don't, check if we have a float
		if c.scoreFloat == nil {
			// if we don't, return -1 and an error
			return -1, errors.ErrNilScore
		} else if float64(*c.scoreFloat) != math.Trunc(float64(*c.scoreFloat)) {
			// if we do, and it can't cast to an int without truncation, return -1 and an error
			return -1, errors.ErrNilScore
		} else {
			// if we do and it got this far cast to int and return
			return int(*c.scoreFloat), nil
		}
	}
	// if we have an int, return it
	return *c.scoreInt, nil
}

// GetScoreFloat retrieves the float score from the comparison result.
// Returns an error if the score is nil or an error exists.
func (c *ComparisonResult) GetScoreFloat() (float32, error) {
	if c.err != nil {
		return -1.00, c.err
	}
	if c.scoreFloat == nil {
		return -1.00, errors.ErrNilScore
	}
	return *c.scoreFloat, nil
}

// IsMatch checks if two ComparisonResult objects are equivalent by
// comparing their type, strings, score, split length, and errors.
func (c *ComparisonResult) IsMatch(other *ComparisonResult) bool {
	// check type, strings, split length
	if !compareInputFields(c, other) {
		return false
	}
	// GetScoreInt will return a score/0 + error/nil
	cScore, cErr := c.GetScoreInt()
	oScore, oErr := other.GetScoreInt()
	// check errors first - this is the underlying ComparisonResult error if not null
	if !errors.CompareErrors(cErr, oErr) {
		return false
	}
	// check int score
	if cScore != oScore {
		return false
	}
	// no returns, we must match
	return true
}

// Print outputs the formatted result of the comparison, optionally including
// verbose error details based on the input flag v.
func (c *ComparisonResult) Print(v bool) {
	fmt.Print(formatComparisonResultOutput(c, v))
}

// Utility Functions

// compareInputFields checks if two ComparisonResult objects are equivalent
// by comparing their types, strings, and split lengths.
func compareInputFields(c1, c2 *ComparisonResult) bool {
	if c1.GetType() != c2.GetType() || c1.GetString1() != c2.GetString1() || c1.GetString2() != c2.GetString2() {
		return false
	}
	cSplit, cErr := c1.GetSplitLength()
	oSplit, oErr := c2.GetSplitLength()
	if !errors.CompareErrors(cErr, oErr) {
		return false
	}
	if cSplit != oSplit {
		return false
	}
	return true
}

// formatComparisonResultOutput formats the comparison result output as a string, with optional verbose error details.
func formatComparisonResultOutput(result *ComparisonResult, v bool) string {
	// Type assert to get the concrete fields we need
	var comparisonType ComparisonResultType
	var string1, string2 string
	var err error
	var score interface{}

	r := result
	comparisonType = r.comparisonType
	string1 = r.string1
	string2 = r.string2
	err = r.err

	if r.scoreInt != nil {
		score = *r.scoreInt
	} else if r.scoreFloat != nil {
		score = *r.scoreFloat
	} else {
		score = -1
	}
	if v {
		if err != nil {
			return fmt.Sprintf("Error during processing %s\nFirst String: %s\nSecond String: %s\nError: %s\n",
				ComparisonResultTypeMap[comparisonType], string1, string2, err.Error())
		} else {
			return fmt.Sprintf("Comparison: %s\nFirst String: %s\nSecond String: %s\nScore: %v\n",
				ComparisonResultTypeMap[comparisonType], string1, string2, score)
		}
	} else {
		if err != nil {
			return fmt.Sprintf("%s (%q/%q) Error: %s\n",
				ComparisonResultTypeMap[comparisonType], string1, string2, err.Error())
		} else {
			return fmt.Sprintf("%s (%q/%q): %v\n", ComparisonResultTypeMap[comparisonType], string1, string2, score)
		}
	}
}
