package strutil

import (
	"errors"
	"testing"
)

func TestNewComparisonManager(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"CM1"},
		{"CM2"},
		{"CM3"},
		{"CM4"},
		{"CM5"},
		{"CM6"},
		{"CM7"},
		{"CM8"},
		{"CM9"},
		{"CM10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := NewComparisonManager()
			if cm.ComparisonResults == nil ||
				cm.SimilarityResults == nil ||
				cm.ShingleData == nil ||
				cm.LCSData == nil {
				t.Errorf("NewComparisonManager() = %v, want %v", cm, nil)
			}
		})
	}
}

func TestComparisonManager_AddComparisonResult(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"AddComparisonResult1"},
		{"AddComparisonResult2"},
		{"AddComparisonResult3"},
	}
	sb := NewLoremSentence().WithComparisonManager()
	i := 0
	for i < 100 {
		comp := randomAlphaNumericString(10)
		sb = sb.LevenshteinDistance(comp)
		i++
	}
	for range tests {
		if len(sb.ComparisonManager().ComparisonResults.GetByType(LevDist)) != 100 {
			t.Errorf("ComparisonManager.AddComparisonResult() = %v, want %v", len(sb.ComparisonManager().ComparisonResults), 100)
		}
	}
}

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

	sh := "Jaro-Winkler Similarity: 0.903698\n"
	verb := "Comparison: Jaro-Winkler Similarity\n" +
		"First String: With great power there must also come great responsibility.\n" +
		"Second String: With great power comes great responsibility.\n" +
		"Score: 0.903698\n"

	shHam := "Hamming Distance Error: error calculating hamming distance\n" +
		"Undefined for strings of unequal length\n"
	verbHam := "Error during processing: Hamming Distance\n" +
		"First String: With great power there must also come great responsibility.\n" +
		"Second String: With great power comes great responsibility.\n" +
		"Error: error calculating hamming distance\n" +
		"Undefined for strings of unequal length\n"

	shLev := "Levenshtein Distance: 17\n"
	verbLev := "Comparison: Levenshtein Distance\n" +
		"First String: With great power there must also come great responsibility.\n" +
		"Second String: With great power comes great responsibility.\n" +
		"Score: 17\n"

	ls := New(correct).
		WithComparisonManager().
		JaroWinklerSimilarity(common).
		LevenshteinDistance(common).
		HammingDistance(common).
		ComparisonManager().
		ComparisonResultsMap()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.resType == JaroWinklerSim {
				if ls.GetByType(tt.resType)[0].formatOutput(false) != sh {
					t.Errorf("ComparisonResult.formatOutput() = %s, want %s",
						ls.GetByType(tt.resType)[0].formatOutput(false), sh)
				}
				if ls.GetByType(tt.resType)[0].formatOutput(true) != verb {
					t.Errorf("ComparisonResult.formatOutput() = %s, want %s",
						ls.GetByType(tt.resType)[0].formatOutput(true), verb)
				}
			}
			if tt.resType == LevDist {
				if ls.GetByType(tt.resType)[0].formatOutput(false) != shLev {
					t.Errorf("ComparisonResult.formatOutput() = %s, want %s",
						ls.GetByType(tt.resType)[0].formatOutput(false), shLev)
				}
				if ls.GetByType(tt.resType)[0].formatOutput(true) != verbLev {
					t.Errorf("ComparisonResult.formatOutput() = %s, want %s",
						ls.GetByType(tt.resType)[0].formatOutput(true), verbLev)
				}
			}
			if tt.resType == HammingDist {
				if ls.GetByType(tt.resType)[0].formatOutput(false) != shHam {
					t.Errorf("ComparisonResult.formatOutput() = %s, want %s",
						ls.GetByType(tt.resType)[0].formatOutput(false), shHam)
				}
				if ls.GetByType(tt.resType)[0].formatOutput(true) != verbHam {
					t.Errorf("ComparisonResult.formatOutput() = %s, want %s",
						ls.GetByType(tt.resType)[0].formatOutput(true), verbHam)
				}
			}
		})
	}
}

func TestComparisonResultGettersInt(t *testing.T) {
	tests := []struct {
		name        string
		resType     ComparisonResultType
		string1     string
		string2     string
		splitLength int
		score       int
		err         error
	}{
		{"GettersTest1", LevDist, "Hello", "Hello!", 0, 1, nil},
		{"GettersTest2", OSADamLevDist, "Hello", "Hello!", 0, 1, nil},
		{"GettersTest3", DamLevDist, "Hello", "Hello!", 0, 1, nil},
		{"GettersTest4", LCSLength, "Hello", "Hello!", 0, 5, nil},
		{"GettersTest4", QGramDist, "Hello", "Hello!", 3, 1, nil},
		{"GettersTest6", HammingDist,
			"Hello",
			"Hello!",
			0,
			0,
			errors.Join(ErrHammingDistanceFailure, errors.New("Undefined for strings of unequal length")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := New(tt.string1).
				WithComparisonManager().
				LevenshteinDistance(tt.string2).
				OSADamerauLevenshteinDistance(tt.string2).
				DamerauLevenshteinDistance(tt.string2).
				LCS(tt.string2).
				LCSEditDistance(tt.string2).
				QgramDistance(tt.string2, tt.splitLength).
				HammingDistance(tt.string2).
				ComparisonManager().
				ComparisonResultsMap()
			entry, ok := res.GetByType(tt.resType)[0].(*ComparisonResultInt)
			if ok {
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
				score, err := entry.GetScoreInt()
				if (err != nil && score != 0) || score != tt.score {
					t.Errorf("ComparisonResult.GetScoreInt() = %d, want %d",
						score, tt.score)
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
				}
				if err != nil && tt.err != nil && entry.GetError().Error() != tt.err.Error() {
					t.Errorf("ComparisonResult.GetScoreInt() = %s, want %s",
						err.Error(), tt.err.Error())
				}
			}
		})
	}
}
