package strutil

import (
	"errors"
	"testing"

	"github.com/bmj2728/utils/pkg/internal/comparison"
	"github.com/bmj2728/utils/pkg/internal/types"
)

func TestComparisonResultPrint(t *testing.T) {
	tests := []struct {
		name    string
		resType ComparisonResultType
	}{
		{"PrintJaroWink", JaroWinklerSim},
		{"PrintLev", LevDist},
		{"PrintHam", HammingDist},
	}
	correct := "With great power there must also come great responsibility."
	common := "With great power comes great responsibility."

	sh := "Jaro-Winkler Similarity " +
		"(\"With great power there must also come great responsibility.\"/" +
		"\"With great power comes great responsibility.\"): 0.90369797\n"
	verb := "Comparison: Jaro-Winkler Similarity\n" +
		"First String: With great power there must also come great responsibility.\n" +
		"Second String: With great power comes great responsibility.\n" +
		"Score: 0.90369797\n"

	shHam := "Hamming Distance " +
		"(\"With great power there must also come great responsibility." +
		"\"/\"With great power comes great responsibility.\") Error: Undefined for strings of unequal length\n"
	verbHam := "Error during processing Hamming Distance\n" +
		"First String: With great power there must also come great responsibility.\n" +
		"Second String: With great power comes great responsibility.\n" +
		"Error: Undefined for strings of unequal length\n"

	shLev := "Levenshtein Distance " +
		"(\"With great power there must also come great responsibility." +
		"\"/\"With great power comes great responsibility.\"): 17\n"
	verbLev := "Comparison: Levenshtein Distance\n" +
		"First String: With great power there must also come great responsibility.\n" +
		"Second String: With great power comes great responsibility.\n" +
		"Score: 17\n"

	x := New(correct).WithComparisonManager(false)
	_ = x.JaroWinklerSimilarity(common)
	_ = x.LevenshteinDistance(common)
	_, _ = x.HammingDistance(common)
	ls := x.GetComparisonManager().GetComparisonResultsMap()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.resType == JaroWinklerSim {
				if formatComparisonResultOutput(ls.GetByType(tt.resType)[0], false) != sh {
					t.Errorf("ComparisonResult.formatComparisonResultOutput() = %s, want %s",
						formatComparisonResultOutput(ls.GetByType(tt.resType)[0], false), sh)
				}
				if formatComparisonResultOutput(ls.GetByType(tt.resType)[0], true) != verb {
					t.Errorf("ComparisonResult.formatComparisonResultOutput() = %s, want %s",
						formatComparisonResultOutput(ls.GetByType(tt.resType)[0], true), verb)
				}
			}
			if tt.resType == LevDist {
				if formatComparisonResultOutput(ls.GetByType(tt.resType)[0], false) != shLev {
					t.Errorf("ComparisonResult.formatComparisonResultOutput() = %s, want %s",
						formatComparisonResultOutput(ls.GetByType(tt.resType)[0], false), shLev)
				}
				if formatComparisonResultOutput(ls.GetByType(tt.resType)[0], true) != verbLev {
					t.Errorf("ComparisonResult.formatComparisonResultOutput() = %s, want %s",
						formatComparisonResultOutput(ls.GetByType(tt.resType)[0], true), verbLev)
				}
			}
			if tt.resType == HammingDist {
				if formatComparisonResultOutput(ls.GetByType(tt.resType)[0], false) != shHam {
					t.Errorf("ComparisonResult.formatComparisonResultOutput() = %s, want %s",
						formatComparisonResultOutput(ls.GetByType(tt.resType)[0], false), shHam)
				}
				if formatComparisonResultOutput(ls.GetByType(tt.resType)[0], true) != verbHam {
					t.Errorf("ComparisonResult.formatComparisonResultOutput() = %s, want %s",
						formatComparisonResultOutput(ls.GetByType(tt.resType)[0], true), verbHam)
				}
			}
		})
	}
}

func TestComparisonResultGetters(t *testing.T) {

	int1 := 1
	int5 := 5

	tests := []struct {
		name        string
		resType     ComparisonResultType
		string1     string
		string2     string
		splitLength int
		scoreInt    *int
		scoreFloat  float32
		err         error
	}{
		{"GettersTest1", LevDist, "Hello", "Hello!", 0, &int1, 1.0, nil},
		{"GettersTest2", OSADamLevDist, "Hello", "Hello!", 0, &int1, 1.0, nil},
		{"GettersTest3", DamLevDist, "Hello", "Hello!", 0, &int1, 1.0, nil},
		{"GettersTest4", LCSLength, "Hello", "Hello!", 0, &int5, 5.0, nil},
		{"GettersTest4", QGramDist, "Hello", "Hello!", 3, &int1, 1.0, nil},
		{"GettersTest6", HammingDist,
			"Hello",
			"Hello!",
			0,
			nil,
			-1.0,
			errors.New("Undefined for strings of unequal length"),
		},
		{"GettersTest7", JaroSim, "Hello", "Hello!", 0, nil,
			0.944444, nil},
		{"GettersTest8", JaroWinklerSim, "Hello", "Hello!", 0, nil,
			0.966667, nil},
		{"GettersTest9", JaccardSim, "Hello", "Hello!", 0, nil,
			0.000000, nil},
		{"GettersTest10", CosineSim, "Hello", "Hello!", 0, nil,
			0.000000, nil},
		{"GettersTest11", SorensenDiceCo, "Hello", "Hello!", 0, nil,
			0.888889, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := New(tt.string1).WithComparisonManager(false)
			_ = x.LevenshteinDistance(tt.string2)
			_ = x.OSADamerauLevenshteinDistance(tt.string2)
			_ = x.DamerauLevenshteinDistance(tt.string2)
			_ = x.LCS(tt.string2)
			_ = x.LCSEditDistance(tt.string2)
			_ = x.QgramDistance(tt.string2, tt.splitLength)
			_ = x.JaroSimilarity(tt.string2)
			_ = x.JaroWinklerSimilarity(tt.string2)
			_ = x.JaccardSimilarity(tt.string2, tt.splitLength)
			_ = x.CosineSimilarity(tt.string2, tt.splitLength)
			_ = x.SorensenDiceCoefficient(tt.string2, tt.splitLength)
			_, _ = x.HammingDistance(tt.string2)
			res := x.GetComparisonManager().GetComparisonResultsMap()
			entry := res.GetByType(tt.resType)[0]

			if entry.GetString1() != tt.string1 {
				t.Errorf("ComparisonResult.GetString1() = %s, want %s",
					entry.GetString1(), tt.string1)
			}
			if entry.GetString2() != tt.string2 {
				t.Errorf("ComparisonResult.GetString2() = %s, want %s",
					entry.GetString2(), tt.string2)
			}
			s1, s2 := entry.GetStrings()
			if s1 != tt.string1 || s2 != tt.string2 {
				t.Errorf("ComparisonResult.GetStrings() = %s, %s, want %s, %s",
					s1, s2, tt.string1, tt.string2)
			}
			if tt.scoreInt != nil {
				scoreInt, err := entry.GetScoreInt()
				if (err != nil && scoreInt != -1) || scoreInt != *tt.scoreInt {
					t.Errorf("ComparisonResult.GetScoreInt() = %d, want %d",
						scoreInt, tt.scoreInt)
				}
			}
			scoreFloat, err := entry.GetScoreFloat()
			floatMatch := comparison.CompareFloats(float64(scoreFloat),
				float64(tt.scoreFloat),
				types.Float64EqualityThreshold)
			if (err != nil && scoreFloat != -1.0) ||
				!floatMatch {
				t.Errorf("ComparisonResult.GetScoreFloat() = %f, want %f",
					scoreFloat, tt.scoreFloat)
			}
			if entry.GetError() != nil && tt.err != nil && entry.GetError().Error() != tt.err.Error() {
				t.Errorf("ComparisonResult.GetScoreInt() = %s, want %s",
					err.Error(), tt.err.Error())
			}
			if entry.GetType() != tt.resType {
				t.Errorf("ComparisonResult.GetType() = %s, want %s",
					entry.GetType(), tt.resType)
			}
			if entry.GetTypeName() != tt.resType.String() {
				t.Errorf("ComparisonResult.GetTypeName() = %s, want %s",
					entry.GetTypeName(), tt.resType.String())
			}
			split, err := entry.GetSplitLength()
			if (tt.resType != QGramDist && split != 0) || split != tt.splitLength {
				t.Errorf("ComparisonResult.GetSplitLength() = %d, want %d",
					split, tt.splitLength)
			}
			if entry.GetError() != nil && tt.err != nil && entry.GetError().Error() != tt.err.Error() {
				t.Errorf("ComparisonResult.GetScoreInt() = %s, want %s",
					err.Error(), tt.err.Error())
			}
			entry.Print(false)
			entry.Print(true)

		})
	}
}

func TestComparisonResultIsMatch(t *testing.T) {
	tests := []struct {
		name   string
		input1 string
		comp1  string
		input2 string
		comp2  string
		result bool
	}{
		{"IsMatch1", "Hello", "Hello!", "Hello", "Hello!", true},
		{"IsMatch2", "Hello", "Hello!", "Hello", "Hello", false},
		{"IsMatch3", "Hello", "Hello!", "Hello", "", false},
		{"IsMatch4", "Hello", "", "Hello", "Hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := New(tt.input1).WithComparisonManager(false)
			_ = x.LevenshteinDistance(tt.comp1)
			res1 := x.GetComparisonManager().GetComparisonResultsMap()
			entry1 := res1.GetByType(LevDist)[0]

			x2 := New(tt.input2).WithComparisonManager(false)
			_ = x2.LevenshteinDistance(tt.comp2)
			res2 := x2.GetComparisonManager().GetComparisonResultsMap()
			entry2 := res2.GetByType(LevDist)[0]

			if entry1.IsMatch(entry2) != tt.result {
				t.Errorf("ComparisonResult.IsMatch() = %v, want %v",
					entry1.IsMatch(entry2), tt.result)
			}
		})
	}

}

func TestComparisonResultIsMatchError(t *testing.T) {
	tests := []struct {
		name   string
		input1 string
		comp1  string
		input2 string
		comp2  string
		result bool
	}{
		{"IsMatch1", "Hello!", "Hello!", "Hello!", "Hello!", true},
		{"IsMatch2", "Hello", "Hello!", "Hello", "Hello!", true},
		{"IsMatch3", "Hello", "Hello", "Hello", "world", false},
		{"IsMatch4", "Hello", "", "Hello", "Hello", false},
		{"IsMatch5", "Hello", "Hello!", "Hello", "Hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x1 := New(tt.input1).WithComparisonManager(false)
			_, _ = x1.HammingDistance(tt.comp1)
			res1 := x1.GetComparisonManager().GetComparisonResultsMap()
			entry1 := res1.GetByType(HammingDist)[0]

			x2 := New(tt.input2).WithComparisonManager(false)
			_, _ = x2.HammingDistance(tt.comp2)
			res2 := x2.GetComparisonManager().GetComparisonResultsMap()
			entry2 := res2.GetByType(HammingDist)[0]

			if entry1.IsMatch(entry2) != tt.result {
				t.Errorf("ComparisonResult.IsMatch() = %v, want %v",
					entry1.IsMatch(entry2), tt.result)
			}
		})
	}

}
